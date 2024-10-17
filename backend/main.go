package main

import (
	"backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	port := os.Getenv("PORT")
	http.HandleFunc("/", routes.GetRoot)
	Port := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
