package ysx

import (
	"fmt"
	"net/http"

	"github.com/imroc/req"
	"github.com/storezhang/gox"
)

type (
	response struct {
		ErrorCode int         `json:"errorCode"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}

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
		// 课程时刻id
		CourseTimeId int64 `xorm:"bigint(20) notnull default(1)" json:"courseTimeId,string"`
	}

	CreateTokenReq struct {
		// 手机号
		Phone string `json:"mobile" validate:"required,alphanum,max=15"`
		// 名称
		Name string `json:"name" validate:"required,min=1,max=64"`
		// 课程时刻Id
		CourseTimeId int64 `json:"courseTimeId" validate:"omitempty"`
	}

	CreateTokenResp struct {
		*VirtualUser

		Token string `json:"token"`
	}
)

func getErr(resp *req.Resp) (err error) {
	var v *response
	if err = resp.ToJSON(&v); nil != err {
		return
	}
	err = fmt.Errorf(v.Message)

	return
}

func CreateTokenBy(phone string, name string, courseTimeId int64, meetingHost string) (tk *CreateTokenResp, err error) {
	var (
		resp *req.Resp
	)
	getTokenUrl := fmt.Sprintf("%s/api/virtual/users/token", meetingHost)
	getTokenParams := req.Param{
		"mobile":       phone,
		"name":         name,
		"courseTimeId": courseTimeId,
	}

	if resp, err = req.Post(getTokenUrl, req.BodyJSON(getTokenParams)); nil != err {
		return
	}

	if resp.Response().StatusCode != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = resp.ToJSON(&tk); err != nil {
		return
	}

	return
}
