package jsonfeed

type Attachment struct {
	URL               string `json:"url"`
	MIMEType          string `json:"mime_type"`
	Title             string `json:"title"`
	SizeInBytes       uint   `json:"size_in_bytes"`
	DurationInSeconds uint   `json:"duration_in_seconds"`
}
