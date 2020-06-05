package search

// ResultItem .
type ResultItem struct {
	Title    string `json:"title"`
	Snippet  string `json:"snippet"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

// Result .
type Result struct {
	Items []*ResultItem `json:"items"`
}
