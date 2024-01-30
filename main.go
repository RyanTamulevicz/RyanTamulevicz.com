package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/resume")
	})

	r.GET("/resume", func(c *gin.Context) {
		data := gin.H{}
		file, err := os.Open("data.json")

		if err != nil {
			resume := os.Getenv("RESUME")
			if resume == "" {
				log.Fatal("No data.json file and no RESUME environment variable set")
			}

			err := json.Unmarshal([]byte(resume), &data)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			defer file.Close()
			decoder := json.NewDecoder(file)
			if err := decoder.Decode(&data); err != nil {
				log.Fatal(err)
			}
		}

		c.HTML(http.StatusOK, "resume.html", data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
