package main

import (
	"flag"
	"github.com/AdiSaripuloh/online-store/modules/product/handlers"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type Service struct {
}

var (
	migration *bool
	seed      *bool
)

func getFlags() {
	migration = flag.Bool("migration", false, "Migrate Database")
	seed = flag.Bool("seed", false, "Seed Database")
	flag.Parse()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	getFlags()
}

func main() {
	handler := handlers.NewHandler()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.Http
	}()
	wg.Wait()
}
