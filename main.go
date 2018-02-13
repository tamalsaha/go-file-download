package main

import (
	// "io"
	"log"
	"net/http"
	// "os"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/moul/http2curl"
	"github.com/tamalsaha/go-oneliners"
)

func main() {
	proxyURL := ""
	u := "<chart-url>"
	username := ""
	password := ""

	// ref: https://github.com/golang/go/blob/release-branch.go1.9/src/net/http/transport.go#L40
	var transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	if proxyURL == "" {
		transport.Proxy = http.ProxyFromEnvironment
	} else {
		pu, _ := url.Parse(proxyURL)
		transport.Proxy = http.ProxyURL(pu)
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		log.Fatalln(err)
	}
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	reqCopy := &http.Request{}
	*reqCopy = *req
	reqCopy.Body = nil
	cmd, _ := http2curl.GetCurlCommand(reqCopy)
	fmt.Println(cmd)

	oneliners.DumpHttpRequestOut(req)

	resp, err := client.Do(req)
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
