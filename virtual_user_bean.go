package ysx

type (
	// CreateTokenReq 创建Token请求
	CreateTokenReq struct {
		// UserId 用户Id
		UserId int64 `json:"userId,string" validate:"required"`
		// NickName 用户昵称
		NickName string `json:"name" validate:"required,min=1,max=64"`
		// AppId 产品Id
		AppId int64 `json:"appId,string" validate:"required"`
		// CourseTimeId 课程时刻Id
		CourseTimeId int64 `json:"courseTimeId" validate:"omitempty"`
	}

	CreateTokenRsp struct {
		*VirtualUser

		Token string `json:"token"`
	}
)
