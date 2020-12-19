package jsonfeed

import "errors"

type Feed struct {
	Version     string       `json:"version"`
	Title       string       `json:"title"`
	HomepageURL string       `json:"home_page_url"`
	FeedURL     string       `json:"feed_url"`
	Description string       `json:"description"`
	UserComment string       `json:"user_comment"`
	NextURL     string       `json:"next_url"`
	Icon        string       `json:"icon"`
	Favicon     string       `json:"favicon"`
	Authors     []Author     `json:"authors"`
	Language    string       `json:"language"`
	Expired     bool         `json:"expired"`
	Hubs        []Hub        `json:"hubs"`
	Items       []Item       `json:"items"`
	Attachments []Attachment `json:"attachments"`
}

func (this *Feed) Validate() []error {

	var valErrors []error

	if this.Version != "https://jsonfeed.org/version/1.1" {
		valErrors = append(valErrors, errors.New("Invalid version."))

		return valErrors
	}

	return valErrors
}
