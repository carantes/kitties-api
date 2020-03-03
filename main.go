package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/carantes/kitties-api/app/bundles/kittiesbundle"
	"github.com/carantes/kitties-api/app/core"
)

func main() {
	// TODO: Dont need to init db every time app run
	err := initDB()

	if err != nil {
		log.Fatal(err)
	}

	startServer(":8080")
}

func initBundles(db *gorm.DB) []core.Bundle {
	return []core.Bundle{
		kittiesbundle.NewKittiesBundle(db),
	}
}

func initDB() error {
	cfg := loadConfig()
	db, err := gorm.Open(cfg.DBType, cfg.DBConnection)

	if err != nil {
		return err
	}

	db.AutoMigrate(&kittiesbundle.Kitty{})

	//Clear table and insert some sample data
	db.Delete(&kittiesbundle.Kitty{})
	db.Create(kittiesbundle.NewKitty("Gaspar", "British", "2016-07-05"))
	db.Create(kittiesbundle.NewKitty("Marcel", "European", "2014-05-02"))

	return nil
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

	db, err := gorm.Open(c.DBType, c.DBConnection)
	defer db.Close()

	if err != nil {
		return err
	}

	r := mux.NewRouter()
	s := r.PathPrefix(c.APIPrefix).Subrouter()

	for _, b := range initBundles(db) {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
