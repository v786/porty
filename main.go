package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"
)

type GalleryGrid struct {
	GalleryObjects []GalleryObj `json:"galleryGrid"`
}

type GalleryObj struct {
	Title string `json: "title"`
	Link  string `json: "link"`
	Img   string `json: "img"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", home)

	router.GET("/mark", mark)

	router.Run(":" + port)
}

func getData(fileName string) GalleryGrid {
	file, _ := ioutil.ReadFile(fileName)

	data := GalleryGrid{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}

func home(c *gin.Context) {
	data := make(map[string]interface{})
	gallreyList := getData("templates/test.json")
	time := time.Now()
	data["time"] = time.Unix()
	data["galleryList"] = gallreyList
	c.HTML(http.StatusOK, "index.tmpl.html", data)
}

func mark(c *gin.Context) {
	c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
}
