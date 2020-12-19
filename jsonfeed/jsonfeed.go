package jsonfeed

import (
	"encoding/json"
	"io"
)

func Parse(feed io.Reader) (*Feed, error) {

	jFeed := &Feed{}

	decoder := json.NewDecoder(feed)

	if err := decoder.Decode(jFeed); err != nil {
		return nil, err
	}

	return jFeed, nil
}
