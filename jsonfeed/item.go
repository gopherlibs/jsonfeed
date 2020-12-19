package jsonfeed

type Item struct {
	ID            string   `json:"id"`
	URL           string   `json:"url"`
	ExternalURL   string   `json:"external_url"`
	Title         string   `json:"title"`
	ContentHTML   string   `json:"content_html"`
	ContentText   string   `json:"content_text"`
	Summary       string   `json:"summary"`
	Image         string   `json:"image"`
	BannerImage   string   `json:"banner_image"`
	DatePublished string   `json:"date_published"`
	DateModified  string   `json:"date_modified"`
	Authors       []Author `json:"authors"`
	Tags          []string `json:"tags"`
	Language      string   `json:"language"`
}
