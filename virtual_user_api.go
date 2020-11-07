package ysx

import (
	`github.com/class100/core`
)

type token interface {
	// CreateToken 创建Token
	CreateToken(ctr *CreateTokenReq) (tk *CreateTokenRsp, err error)
}

func (hsc *httpSignatureClient) CreateToken(ctr *CreateTokenReq) (rsp *CreateTokenRsp, err error) {
	rsp = new(CreateTokenRsp)
	err = hsc.requestApi(
		ApiPathCreateToken,
		core.HttpMethodPost,
		ctr,
		ApiVersionDefault,
		&rsp,
	)

	return
}
