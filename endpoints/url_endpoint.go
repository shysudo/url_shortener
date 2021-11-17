package endpoints

import (
	"fmt"
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

// var urls []models.UrlPayload
var store = map[string]string{}

func (resource urlResourceImpl) ShortenUrl(req *restful.Request, res *restful.Response) {
	payload := models.UrlPayload{}
	err := req.ReadEntity(&payload)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}
	shortUrl, err := service.UrlService.ShortenUrl(payload.Url)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
		return
	}
	store[payload.Url] = shortUrl
	fmt.Println(store)
	res.ResponseWriter.WriteHeader(http.StatusCreated)
	res.ResponseWriter.Write([]byte(shortUrl))
}
