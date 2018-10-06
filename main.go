package main

import (
	"log"
	"sync"

	"github.com/agungdwiprasetyo/agungkiki-backend/config"
	env "github.com/joho/godotenv"
)

func main() {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}

	conf := config.New()
	service := NewService(conf)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.ServeHTTP(8008)
	}()

	wg.Wait()
}
