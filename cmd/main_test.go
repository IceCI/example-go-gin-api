package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebEndpoints(t *testing.T) {
	a := assert.New(t)

	dbConfig, err := loadDbConfig()
	if err != nil {
		t.Fatal(err)
	}

	db, err := setupDb(dbConfig)
	if err != nil {
		t.Fatal(err)
	}

	router, err := setupRouter(db)
	if err != nil {
		t.Fatal(err)
	}

	setupFixtures(db)

	t.Run("health endpoint should return status 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		a.Equal(200, w.Code)
		a.Equal("ok", w.Body.String())
	})

	t.Run("quotes endpoint should return an element", func(t *testing.T) {
		var result []map[string]interface{}

		req, _ := http.NewRequest("GET", "/quote", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil {
			t.Fatal(fmt.Sprintf("failed parsing endpoint response, %s", err))
		}

		a.Equal(200, w.Code)
		a.Equal(1, len(result))
	})

	teardownFixtures(db)
}

func setupFixtures(db *gorm.DB) {
	quote := Quote{
		Quote:  "A day without sunshine is like, you know, night",
		Author: "Steve Martin",
	}

	db.Create(&quote)
}

func teardownFixtures(db *gorm.DB) {
	db.Delete(Quote{}, "author = ?", "Steve Martin")
}
