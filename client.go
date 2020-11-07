package ysx

import (
	"fmt"

	"github.com/class100/core"
	`github.com/storezhang/gox`
)

type httpSignatureClient struct {
	client *core.HttpSignatureClient

	options options
}

func (hsc *httpSignatureClient) requestApi(
	path ApiPath,
	method core.HttpMethod,
	data interface{},
	version ApiVersion,
	rsp interface{},
	params ...gox.HttpParameter,
) (err error) {
	var url string
	if ApiVersionDefault == version {
		url = fmt.Sprintf("%s/api/%s", hsc.options.Endpoint, path)
	} else {
		url = fmt.Sprintf("%s/api/%s/%s", hsc.options.Endpoint, version, path)
	}

	if rsp == nil {
		rsp = new(interface{})
	}

	return hsc.client.RequestApi(url, method, data, rsp, params...)
}
