package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type attrs map[string]string

// Search does baidu search for up to 10*maxPage results.
func Search(ctx context.Context, kw string, maxPage int) (*Result, error) {
	result := &Result{
		Items: []*ResultItem{},
	}
	session := httpSession{}

	for p := 0; p < maxPage; p++ {
		if ctx.Err() != nil {
			return result, ctx.Err()
		}

		body, err := session.get(ctx, kw, p*10)
		if err != nil {
			return nil, err
		}
		defer body.Close()

		res, err := parseSearchResultPage(body)
		if err != nil {
			return nil, err
		}

		result.Items = append(result.Items, res...)

		// It seems baidu server needs a few time to setup the session with the
		// given BAIDUID, so here sleep seconds according to experience.
		if p == 0 {
			time.Sleep(2 * time.Second)
		}
	}

	return result, nil
}

type httpSession struct {
	cookie string
}

func (s *httpSession) get(
	ctx context.Context, kw string, pn int) (io.ReadCloser, error) {

	url := "http://www.baidu.com/s?wd=" + url.QueryEscape(kw)
	if pn > 0 {
		url = fmt.Sprintf("%v&pn=%v", url, pn)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)

	if len(s.cookie) > 0 {
		// Fill BAIDUID.
		// "BAIDUID=BE8E7632B8D5D3686917110E3E133F9B:FG=1; max-age=31536000;
		//  expires=Sat, 05-Jun-21 07:30:44 GMT; domain=.baidu.com; path=/;
		//  version=1; comment=bd"
		req.Header.Add("Cookie", s.cookie)
	}
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	client := http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if c := rsp.Header.Get("Content-Type"); !strings.Contains(c, "text/html") {
		return nil, fmt.Errorf("unexpected content type: %v", c)
	}

	for _, v := range rsp.Header["Set-Cookie"] {
		// When page number(pn) > 0, baidu search requires BAIDUID in the req.
		// cookie, so here stores the given BAIDUID from the resp. of the 1st
		// request.
		if len(s.cookie) == 0 &&
			strings.HasPrefix(v, "BAIDUID=") &&
			strings.Contains(v, "max-age=31536000") {
			s.cookie = v
		}
	}

	return rsp.Body, nil
}

func parseSearchResultPage(r io.Reader) ([]*ResultItem, error) {
	o := []*ResultItem{}

	z := html.NewTokenizer(r)
	z.NextIsNotRawText()

	// Jump to <div id="content_left">
	if _, eof := jumpToID(z, "div", "content_left"); eof {
		return o, nil
	}

	for {
		if _, eof := jumpToClass(z, "div", "c-container"); eof {
			break
		}

		b, eof := expandToken(z)
		if eof {
			break
		}

		if !isCommonResult(b.String()) {
			continue
		}

		sz := html.NewTokenizer(bytes.NewReader(b.Bytes()))
		imgURL := getImageURL(sz)

		sz = html.NewTokenizer(bytes.NewReader(b.Bytes()))
		if item := parseCcontainer(sz); item != nil {
			item.ImageURL = imgURL
			o = append(o, item)
		}
	}

	return o, nil
}

func isCommonResult(content string) bool {
	return strings.Contains(content, "c-abstract") &&
		strings.Contains(content, "c-tools")
}

func getImageURL(z *html.Tokenizer) (url string) {
	attr, eof := jumpToClass(z, "img", "c-img")
	if eof {
		return
	}
	if val, exists := attr["src"]; exists {
		url = val
	}
	return
}

func parseCcontainer(z *html.Tokenizer) *ResultItem {
	if _, eof := jumpToClass(z, "div", "c-abstract"); eof {
		return nil
	}

	b, eof := expandToken(z)
	if eof {
		return nil
	}
	snippet, _ := parseSnippet(b.String())

	attrs, eof := jumpToClass(z, "div", "c-tools")
	if eof {
		return nil
	}

	dataTools, exists := attrs["data-tools"]
	if !exists {
		return nil
	}

	data := struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}{}

	if err := json.Unmarshal([]byte(dataTools), &data); err != nil {
		return nil
	}

	return &ResultItem{
		Title:   data.Title,
		Snippet: snippet,
		URL:     data.URL,
	}
}

// parseSnippet returns snippet & time(if existed).
func parseSnippet(raw string) (string, string) {
	var (
		tf       = ""
		tfPrefix = "<span class=\" newTimeFactor_before_abs m\">"
	)
	// Some snippet start with <span class=" newTimeFactor_before_abs m">
	if strings.HasPrefix(raw, tfPrefix) {
		slice := strings.Split(raw, "</span>")
		raw = slice[1]
		tf = slice[0][len(tfPrefix):]
	}

	removeList := []string{"<em>", "</em>"}
	for _, i := range removeList {
		raw = strings.Replace(raw, i, "", -1)
	}
	return raw, tf
}
