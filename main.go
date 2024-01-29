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
		file, err := os.Open("data.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		data := gin.H{}
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "resume.html", data)
	})

	log.Println("Starting server on :3001")
	if err := r.Run(":3001"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
