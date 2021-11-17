package service

type urlService interface {
	ShortenUrl(url string) string
	FetchShortenUrl(url string) string
}

type urlServiceImpl struct{}

var UrlService urlService = urlServiceImpl{}

func (service urlServiceImpl) ShortenUrl(url string) string {
	return ""
}

func (service urlServiceImpl) FetchShortenUrl(url string) string {
	return ""
}
