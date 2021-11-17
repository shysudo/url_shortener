package service

import (
	"math/rand"

	"github.com/hashicorp/go-uuid"
)

type urlService interface {
	ShortenUrl(url string) (string, error)
	FetchShortenUrl(url string) string
}

type urlServiceImpl struct{}

var UrlService urlService = urlServiceImpl{}

func (service urlServiceImpl) ShortenUrl(url string) (string, error) {
	shortUrl, err := generateShortLink(url)
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (service urlServiceImpl) FetchShortenUrl(url string) string {
	return ""
}

func generateShortLink(url string) (string, error) {
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return "", err
	}
	s := randString(10)

	return "short://" + s + "-" + uuid[:15], nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
