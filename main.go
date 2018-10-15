package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/agungdwiprasetyo/agungkiki-backend/config"
	env "github.com/joho/godotenv"
)

func main() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("APP_PATH", appPath)

	if err := env.Load(fmt.Sprintf("%s/.env", appPath)); err != nil {
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
