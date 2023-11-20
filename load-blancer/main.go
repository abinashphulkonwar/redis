package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	rp := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
		},
		ModifyResponse: func(r *http.Response) error {
			return r.Write(bytes.NewBufferString("hiiii"))
		},
	}
	s := http.Server{
		Addr:    "localhost:3000",
		Handler: rp,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}

}
