package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/vi/post"
	r.POST(url, CreatePostHandler)
	body := `{"community_id": 1, "title": "Test", "content": "Test"}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// NOTE: method 1
	// assert.Contains(t, w.Body.String(), CodeNeedLogin.Msg())

	// NOTE: method 2
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json unmarshal w.body failed. err%s", err)
	}
	assert.Equal(t, res.Msg, CodeNeedLogin.Msg())
}
