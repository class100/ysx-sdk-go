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

	// GetUserReq 查询用户的请求
	GetUserReq struct {
		// 用户编号
		UserId int64 `json:"userId,string" validate:"required"`
		// 课节编号
		CourseTimeId int64 `json:"courseTimeId,string" validate:"required"`
	}

	// GetUserRsp 查询用户的响应
	GetUserRsp struct {
		*VirtualUser
	}
)
