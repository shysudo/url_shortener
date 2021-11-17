package endpoints

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/shysudo/url_shortner/services"
	"github.com/stretchr/testify/assert"
)

func TestShortenUrlSuccess(t *testing.T) {
	httpRecorder := httptest.NewRecorder()
	res := &restful.Response{ResponseWriter: httpRecorder}
	body := ioutil.NopCloser(strings.NewReader(`{"url":"gireeshkademan"}`))
	req := restful.NewRequest(httptest.NewRequest("POST", "/user", body))
	req.Request.Header.Set("content-type", restful.MIME_JSON)
	UrlResource.ShortenUrl(req, res)
	assert.NotEmpty(t, httpRecorder.Body.String())
	assert.Equal(t, http.StatusOK, res.StatusCode())
}

func TestShortenUrlFetchShortUrlSuccess(t *testing.T) {
	httpRecorder := httptest.NewRecorder()
	services.Store["http://google.com/"] = "short://somevalue"
	res := &restful.Response{ResponseWriter: httpRecorder}
	body := ioutil.NopCloser(strings.NewReader(`{"url":"http://google.com/"}`))
	req := restful.NewRequest(httptest.NewRequest("POST", "/user", body))
	req.Request.Header.Set("content-type", restful.MIME_JSON)
	UrlResource.ShortenUrl(req, res)
	assert.Equal(t, "short://somevalue", httpRecorder.Body.String())
	assert.Equal(t, http.StatusOK, res.StatusCode())
}
func TestShortenUrlFailled(t *testing.T) {
	httpRecorder := httptest.NewRecorder()
	res := &restful.Response{ResponseWriter: httpRecorder}
	body := ioutil.NopCloser(strings.NewReader(`{"url":"http://google.com/"}`))
	req := restful.NewRequest(httptest.NewRequest("POST", "/user", body))
	UrlResource.ShortenUrl(req, res)
	// assert.NotEmpty(t, httpRecorder.Body.String())
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode())
}
