package endpoints

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/shysudo/url_shortner/models"
	service "github.com/shysudo/url_shortner/services"
)

type urlResource interface {
	ShortenUrl(req *restful.Request, res *restful.Response)
}

type urlResourceImpl struct{}

var UrlResource urlResource = urlResourceImpl{}

func (resource urlResourceImpl) ShortenUrl(req *restful.Request, res *restful.Response) {
	payload := models.UrlPayload{}
	err := req.ReadEntity(&payload)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}
	if shortUrl := service.UrlService.FetchShortenUrl(payload.Url); shortUrl != "" {
		res.ResponseWriter.WriteHeader(http.StatusOK)
		res.ResponseWriter.Write([]byte(shortUrl))
		return
	}
	shortUrl, err := service.UrlService.ShortenUrl(payload.Url)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}
	service.Store[payload.Url] = shortUrl
	res.ResponseWriter.WriteHeader(http.StatusCreated)
	res.ResponseWriter.Write([]byte(shortUrl))
}
