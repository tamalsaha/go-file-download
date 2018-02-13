package main

import (
	// "io"
	"log"
	"net/http"
	// "os"

	"fmt"

	"github.com/moul/http2curl"
	"github.com/tamalsaha/go-oneliners"
)

func main() {
	url := "<chart-url>"
	username := ""
	password := ""

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	reqCopy := &http.Request{}
	*reqCopy = *req
	reqCopy.Body = nil
	cmd, _ := http2curl.GetCurlCommand(reqCopy)
	fmt.Println(cmd)

	oneliners.DumpHttpRequestOut(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	oneliners.FILE("file downloaded")

	//_, err = io.Copy(os.Stdout, resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
