package echgo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

type evalFunc func(string)

// Given a test evaluation function, prepare a test engine with the root route set to check the evaluation function.
func createTestEngine() (router *gin.Engine, resp *httptest.ResponseRecorder) {
	resp = httptest.NewRecorder()
	router = gin.Default()
	router.GET("/", Echgo)
	return
}

// Prepare a test request to represent specific parameters
func createRequest() *http.Request {
	body := new(bytes.Buffer)
	req, _ := http.NewRequest("GET", "/?hello=world", body)
	return req
}

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Echgo", func() {
		g.It("Should repeat one query string parameter", func() {
			router, resp := createTestEngine()
			req := createRequest()
			router.ServeHTTP(resp, req)
			g.Assert(resp.Body.String()).Equal("{\"hello\":[\"world\"]}\n")
		})
	})
}
