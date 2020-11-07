package ysx

type (
	// CreateTokenReq 创建Token请求
	CreateTokenReq struct {
		BaseVirtualUser
	}

	// CreateTokenRsp  创建Token响应
	CreateTokenRsp struct {
		*VirtualUser

		// Token
		Token string `json:"token"`
	}
)
