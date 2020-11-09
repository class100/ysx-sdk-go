package ysx

import (
	`strings`

	`github.com/class100/core`
)

// Client 云视课堂客户端
type Client interface {
	meeting
	token
}

// NewClient 创建云视课堂客户端
func NewClient(options ...Option) (client Client, err error) {
	var hsc *core.HttpSignatureClient

	appliedOptions := defaultOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	if "" == strings.TrimSpace(appliedOptions.Endpoint) {
		err = ErrMustSetEndpoint

		return
	}

	if hsc, err = core.NewHttpSignatureClient(appliedOptions.options...); nil != err {
		return
	}
	client = &httpSignatureClient{
		client:  hsc,
		options: appliedOptions,
	}

	return
}
