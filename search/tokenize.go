package search

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func jumpToID(z *html.Tokenizer, tag, id string) (a attrs, eof bool) {
	for {
		token := z.Next()
		if token == html.ErrorToken {
			eof = true
			return
		}

		name, hasAttr := z.TagName()
		if string(name) != tag || !hasAttr {
			continue
		}

		a = getAttrs(z)
		if val, exists := a["id"]; exists && val == id {
			break
		}
	}
	return
}

func jumpToClass(z *html.Tokenizer, tag, class string) (a attrs, eof bool) {
	for {
		token := z.Next()
		if token == html.ErrorToken {
			eof = true
			return
		}

		name, hasAttr := z.TagName()
		if string(name) != tag || !hasAttr {
			continue
		}

		a = getAttrs(z)
		val, exists := a["class"]
		if !exists {
			continue
		}

		for _, c := range strings.Split(val, " ") {
			if c == class {
				return
			}
		}
	}
}

func expandToken(z *html.Tokenizer) (*bytes.Buffer, bool) {
	buffer := new(bytes.Buffer)
	depth := 1

	for {
		switch z.Next() {
		case html.ErrorToken:
			return buffer, true

		case html.StartTagToken:
			name, _ := z.TagName()
			if computeDepth(string(name)) {
				depth++
			}

		case html.EndTagToken:
			name, _ := z.TagName()
			if computeDepth(string(name)) {
				depth--
			}
			if depth == 0 {
				return buffer, false
			}
		}
		buffer.Write(z.Raw())
	}
}

func computeDepth(name string) bool {
	// When computing depth, we only care about div, span & a.
	allow := []string{"div", "span", "a"}
	for _, a := range allow {
		if name == a {
			return true
		}
	}
	return false
}

func getAttrs(z *html.Tokenizer) (o attrs) {
	o = attrs{}
	for {
		key, val, more := z.TagAttr()
		o[string(key)] = string(val)
		if !more {
			break
		}
	}
	return
}
