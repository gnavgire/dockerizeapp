package danalyse

import (
	"fmt"
	"os"
	"archive/zip"
//	"appDockerize/dapproute"
	//"io/ioutil"
	"os/exec"
	"regexp"
	"log"
	//"github.com/gin-gonic/gin"
)

func Analysebin(binname string)(){

	out, err := exec.Command("scripts/file.sh", binname).Output()
	if err != nil {
		log.Fatal("Error while executing file command")
	}
	fmt.Printf("File is %s \n", out)

	//ELF 64-bit LSB executable, x86-64,
	re := regexp.MustCompile(`(\w+) (.*?) (\w\d\d-\d\d)?(\w+)?`)
	result := re.FindStringSubmatch(string(out))
	fmt.Println(result[1], result[2], result[3])
	switch result[3] {
		case "x86-64" :
			fmt.Println("This is a x86-64 image")
			buildx86Container(binname)
		case "object" :
			fmt.Println("This needs a ARM image")
			//buildarmhfContainer(binname)
		default :
			fmt.Println("No handling for this image")
	}
}

func buildx86Container(binname string)(){

	fmt.Println("Received binary name is ", binname)
	out, err := exec.Command("scripts/createx86-64dockerfile.sh", binname).Output()
	if err != nil {
		log.Fatal("Error while executing createx86-64dockerfile.sh command", out)
	}

	out1, err1 := exec.Command("scripts/createDockercontainer.sh").Output()
	if err1 != nil {
		log.Fatal("Error while executing createDockercontainer.sh command", out1)
	}
}
