package endpoints

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/shysudo/url_shortner/models"
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
	fmt.Println(payload)
}
