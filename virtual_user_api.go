package ysx

import (
	"github.com/class100/core"
)

type token interface {
	// CreateToken 创建Token
	CreateToken(ctr *CreateTokenReq) (tk *CreateTokenRsp, err error)
	// GetUser 查询用户
	GetUser(req *GetUserReq) (rsp *GetUserRsp, err error)
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

func (hsc *httpSignatureClient) GetUser(req *GetUserReq) (rsp *GetUserRsp, err error) {
	rsp = new(GetUserRsp)
	err = hsc.requestApi(
		ApiPathGetUser,
		core.HttpMethodGet,
		req,
		ApiVersionDefault,
		&rsp,
	)

	return
}
