package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenUrlSuccess(t *testing.T) {
	shortUrl, err := UrlService.ShortenUrl("http://google.com")
	assert.Nil(t, err)
	assert.NotNil(t, shortUrl)
}

func TestShortenUrlError(t *testing.T) {
	shortUrl, err := UrlService.ShortenUrl("")
	assert.NotNil(t, err)
	assert.Equal(t, "", shortUrl)
	assert.Equal(t, "url should not be an empty", err.Error())
}

func TestFetchShortenUrlAvailable(t *testing.T) {
	Store["http://google.com"] = "short://sometext-0"
	urls := UrlService.FetchShortenUrl("http://google.com")
	assert.Equal(t, "short://sometext-0", urls)
}
func TestFetchShortenUrlNotAvailable(t *testing.T) {
	urls := UrlService.FetchShortenUrl("http://google.com1")
	assert.Equal(t, "", urls)
}
