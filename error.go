package ysx

import (
	`github.com/storezhang/gox`
)

var (
	// ErrMustSetEndpoint 必须设置Endpoint服务器端点
	ErrMustSetEndpoint = &gox.CodeError{ErrorCode: 1000, Message: "必须设置Endpoint服务器端点"}
)
