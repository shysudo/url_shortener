package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/shysudo/url_shortner/endpoints"
	"github.com/shysudo/url_shortner/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	restful.DefaultContainer.Add(webService())
	log.Println("Service is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func webService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/url").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/").To(endpoints.UrlResource.ShortenUrl).
		Doc("create short url").
		Reads(models.UrlPayload{}).
		Returns(201, "Created", models.UrlPayload{}))
	return ws
}
