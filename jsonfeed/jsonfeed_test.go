package jsonfeed

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	file, _ := os.Open("../testdata/test-feed-1.json")

	feed, err := Parse(file)
	if err != nil {
		println("Failed to parse feed.")
	}

	assert.Equal(t, "https://jsonfeed.org/version/1.1", feed.Version)
	assert.Equal(t, "Brent Simmons’s Microblog", feed.Title)
	assert.Equal(t, "Brent Simmons", feed.Authors[0].Name)

	file, _ = os.Open("../testdata/test-feed-2.json")

	feed, err = Parse(file)
	if err != nil {
		println("Failed to parse feed.")
	}

	assert.Equal(t, "https://jsonfeed.com/version/1.1", feed.Version)
	assert.Equal(t, "Brent Simmons’s Microblog", feed.Title)
	assert.Equal(t, "Brent Simmons", feed.Authors[0].Name)
}
