package main

import (
	"fmt"
	"appDockerize/danalyse"
	"io/ioutil"
	//"os/exec"
	//"regexp"
	"log"
	"github.com/gin-gonic/gin"
)

var Docstatus = false

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.POST("/upload", func(c *gin.Context) {

	dstfile, err := ioutil.TempFile("", "appbin")
	if err != nil {
		log.Fatal(err)
	}
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dstfile.Name())

	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))

	danalyse.Analysebin(dstfile.Name())
	Docstatus = true
	})

	// Get user value
	r.GET("/status", func(c *gin.Context) {
		if Docstatus {
			c.JSON(200, gin.H{"user": "Ganesh", "value": "GoodJob"})
		} else {
			c.JSON(200, gin.H{"user": "Ganesh", "status": "Not nice "})
		}
	})
	return r
}


func main() {
	fmt.Println("Starting GAutoDocker....")
	router := setupRouter()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB

	router.Run(":8080")

}
