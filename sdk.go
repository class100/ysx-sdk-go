package ysx

import (
	"crypto/tls"
	"net/http"

	"github.com/imroc/req"
)

func init() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req.SetClient(client)
}
