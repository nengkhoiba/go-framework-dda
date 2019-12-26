package main

import (
	"github.com/bradfitz/gomemcache/memcache"

	"log"
	"net/http"
	"os"
	"sync"
	"time"

	micro_middleware "github.com/nengkhoiba/go-framework-dda/middleware"

	//microservices apps
	auth "github.com/nengkhoiba/go-framework-dda/apps/auth"
	//microservices apps

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	memCached  *memcache.Client
	setupOnce  sync.Once
	Middleware *micro_middleware.Middleware
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	initializeDB()
	initializeMiddleware()
	initializeCache()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})
	// mount the microservices to the main router
	//Mounting the auth in the router
	oa := auth.App{
		Middleware: Middleware,
	}
	r.Mount("/auth", oa.Routes())

	http.DefaultClient.Timeout = time.Minute * 10
	log.Fatal(http.ListenAndServe(os.Getenv("API_PORT"), r))

}

func initializeDB() {
	// add database connection code
}
func initializeCache() {
	//memCached = memcache.New(os.Getenv("AUTH_MEMCACHE_ADDRESS"))
	// add any cache
}
func initializeMiddleware() {
	Middleware = &micro_middleware.Middleware{}
}
