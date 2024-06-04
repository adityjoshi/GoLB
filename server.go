package main

import "net/http/httputil"

type server struct {
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}
