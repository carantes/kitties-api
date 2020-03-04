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
	// Load envs
	cfg := loadConfig()

	// Init DB
	db, err := initDB(cfg.DBType, cfg.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init Bundles
	bundles := initBundles(db)

	// Start Server
	startServer(":8080", cfg.APIPrefix, bundles)
}

func initBundles(db *gorm.DB) []core.Bundle {
	return []core.Bundle{
		kittiesbundle.NewKittiesBundle(db),
	}
}

func initDB(dbType string, dbConnection string) (*gorm.DB, error) {
	db, err := gorm.Open(dbType, dbConnection)
	if err != nil {
		return &gorm.DB{}, err
	}

	db.AutoMigrate(&kittiesbundle.Kitty{})

	//Clear table and insert some sample data
	db.Delete(&kittiesbundle.Kitty{})
	db.Create(kittiesbundle.NewKitty("Gaspar", "British", "2016-07-05"))
	db.Create(kittiesbundle.NewKitty("Marcel", "European", "2014-05-02"))

	return db, nil
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

func startServer(addr string, apiPrefix string, bundles []core.Bundle) error {
	r := mux.NewRouter()
	s := r.PathPrefix(apiPrefix).Subrouter()

	for _, b := range bundles {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, nil))

	return nil
}
