package main

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hola World", string(body))
}

func TestGetArticles(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/article", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var articles []Article
	err = json.NewDecoder(resp.Body).Decode(&articles)
	assert.NoError(t, err)
	assert.Empty(t, articles)
}

func TestCreateArticle(t *testing.T) {
	app := setupApp()

	articleJSON := `{"Title":"Test Article","Text":"This is a test article"}`
	req := httptest.NewRequest("POST", "/article", strings.NewReader(articleJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var createdArticle Article
	err = json.NewDecoder(resp.Body).Decode(&createdArticle)
	assert.NoError(t, err)
	assert.Equal(t, "Test Article", createdArticle.Title)
	assert.Equal(t, "This is a test article", createdArticle.Text)

	// Verify the article was added
	req = httptest.NewRequest("GET", "/articles", nil)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var articles []Article
	err = json.NewDecoder(resp.Body).Decode(&articles)
	assert.NoError(t, err)
	assert.Len(t, articles, 1)
	assert.Equal(t, "Test Article", articles[0].Title)
	assert.Equal(t, "This is a test article", articles[0].Text)
}

func TestCreateInvalidArticle(t *testing.T) {
	app := setupApp()

	invalidJSON := `{"Title":123,"Text":456}`
	req := httptest.NewRequest("POST", "/article", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}
