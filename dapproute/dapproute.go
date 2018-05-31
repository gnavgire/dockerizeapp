package dapproute

import (
	"fmt"
	"os/exec"
	"log"
	"regexp"
	"github.com/gin-gonic/gin"
)

var Docstatus = 0
type Dappcontext gin.Context

func (c *Dappcontext) Uploadhandle()  {
	Out := "/root/go/src/GAutoDocker/dest1"
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, Out)

	c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))

	out, err := exec.Command("scripts/file.sh", Out).Output()
	if err != nil {
		log.Fatal("Error while executing file command")
	}
	fmt.Printf("File is %s \n", out)

	//ELF 64-bit LSB executable, x86-64,
	re := regexp.MustCompile(`(\w+) (.*?) (\w+)`)
	result := re.FindStringSubmatch(string(out))
	fmt.Println(result[1], result[2], result[3])
	Docstatus = 1
}


func (c *Dappcontext) Getstatus()  {
		//user := c.Params.ByName("name")
		//value, ok := DB[user]
		if Docstatus {
			c.JSON(200, gin.H{"user": "Ganesh", "value": "GoodJob"})
		} else {
			c.JSON(200, gin.H{"user": "Ganesh", "status": "Not nice "})
		}
}


