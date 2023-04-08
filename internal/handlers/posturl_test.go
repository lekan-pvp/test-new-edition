package handlers

import (
	"github.com/go-chi/chi"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/storage/memrepo"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func ExamplePostURL() {
	router := chi.NewRouter()
	config.New()
	serverAddress := config.Cfg.ServerAddress
	repo := memrepo.New(config.Cfg.FileStoragePath)
	router.Post("/", PostURL(repo))

	log.Fatal(http.ListenAndServe(serverAddress, router))
}

func BenchmarkPostURL(b *testing.B) {
	data := "http://yandex.ru"
	r, _ := http.NewRequest("POST", "/", strings.NewReader(data))
	w := httptest.NewRecorder()
	repo := memrepo.New(config.Cfg.FileStoragePath)
	handler := PostURL(repo)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(w, r)
	}
}
