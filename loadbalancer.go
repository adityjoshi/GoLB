package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []*httputil.ReverseProxy{
		createHost("http://127.0.0.1:5000"),
		createHost("http://127.0.0.1:5001"),
		createHost("http://127.0.0.1:5002"),
		createHost("http://127.0.0.1:5003"),
		createHost("http://127.0.0.1:5004"),
	}
	lastServerIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	fmt.Print("server is running\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server := getServer()
	server.ServeHTTP(res, req)

}

func getServer() *httputil.ReverseProxy {
	serveIndex := (lastServerIndex + 1) % 5
	server := serverList[serveIndex]
	lastServerIndex = serveIndex
	return server
}

func createHost(urlStr string) *httputil.ReverseProxy {
	u, _ := url.Parse(urlStr)
	return httputil.NewSingleHostReverseProxy(u)

}
