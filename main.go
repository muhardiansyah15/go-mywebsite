package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", getHome)

	// Route for about.html
	router.GET("/about", getAbout)

	// Route for resume.html
	router.GET("/resume", getResume)

	// Route for certifications.html
	router.GET("/certifications", getCertifications)

	// Route for publications.html
	router.GET("/publications", getPublications)

	// Route for contactus.html
	router.GET("/contact", getContactUs)

	router.Run(":8888")
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func getAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

func getResume(c *gin.Context) {
	c.HTML(http.StatusOK, "resume.html", gin.H{})
}

func getCertifications(c *gin.Context) {
	c.HTML(http.StatusOK, "certifications.html", gin.H{})
}

func getPublications(c *gin.Context) {
	c.HTML(http.StatusOK, "certifications.html", gin.H{})
}

func getContactUs(c *gin.Context) {
	c.HTML(http.StatusOK, "contactus.html", gin.H{})
}
