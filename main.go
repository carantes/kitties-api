package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/carantes/kitties-api/app/bundles/kittiesbundle"
	"github.com/carantes/kitties-api/app/core"
)

func main() {
	startServer(":8080")
}

func initBundles() []core.Bundle {
	return []core.Bundle{
		kittiesbundle.NewKittiesBundle(),
	}
}

func loadConfig() *core.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	c := &core.Config{}
	c.Load()
	return c
}

func startServer(addr string) error {
	c := loadConfig()
	r := mux.NewRouter()
	s := r.PathPrefix(c.APIPrefix).Subrouter()

	for _, b := range initBundles() {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
