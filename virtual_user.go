package ysxsdk

import (
	`fmt`

	`github.com/imroc/req`
	`github.com/storezhang/gox`
)

type (
	VirtualUser struct {
		gox.BaseStruct `xorm:"extends"`

		// 手机号
		Mobile string `xorm:"varchar(15) notnull default('') unique(uidx_phone)" json:"phone" validate:"omitempty,alphanum,max=15"`
		// 名称
		Name string `xorm:"varchar(32) notnull default('')" json:"name"`
		// 会议用虚拟手机号
		VirtualMobile string `xorm:"varchar(64) notnull default('')" json:"VirtualMobile"`
		// 虚拟用户id
		VirtualUserId string `xorm:"varchar(64) notnull default('')" json:"VirtualUserId"`
	}

	CreateTokenReq struct {
		// 手机号
		Phone string `json:"mobile" validate:"required,alphanum,max=15"`
		// 名称
		Name string `json:"name" validate:"required,min=1,max=64"`
	}

	CreateTokenResp struct {
		VirtualUser

		Token string `json:"token"`
	}
)

func CreateTokenBy(phone string, name string, meetingHost string) (tk *CreateTokenResp, err error) {
	var (
		resp *req.Resp
	)
	getTokenUrl := fmt.Sprintf("%s/api/virtual/users/token", meetingHost)
	getTokenParams := req.Param{
		"mobile": phone,
		"name":   name,
	}

	if resp, err = req.Post(getTokenUrl, req.BodyJSON(getTokenParams)); nil != err {
		return
	}

	if err = resp.ToJSON(tk); err != nil {
		return
	}
	return
}

