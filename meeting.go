package ysx

import (
	"fmt"
	"net/http"

	"github.com/imroc/req"
	"github.com/storezhang/gox"
)

type (
	// 会议数据
	MeetingData struct {
		MeetingId string `json:"Id"`
		MeetingNo uint64 `json:"MeetingNo"`
	}

	JoinMeetingReq struct {
		// 开始时间
		StartTime gox.Timestamp `json:"startTime" validate:"required"`
		// 持续时间
		// 单位：分钟
		Duration int64 `json:"duration" validate:"required"`
		// 主题
		Topic string `json:"topic" validate:"required,omitempty,min=1,max=64"`
		// 主持人手机号
		HostMobile string `json:"hostMobile" validate:"required,omitempty,alphanum,max=15"`
		// 主持人名称
		HostName string `json:"hostName" validate:"required,omitempty,min=2,max=32"`
		// 课程时刻Id
		CourseTimeId int64 `json:"courseTimeId" validate:"omitempty"`
	}

	EndMeetingReq struct {
		// 用户ID
		UserID string `json:"userId" validate:"required"`
		// 会议Id
		MeetingId string `json:"meetingId" validate:"required"`
	}
)

// startTime 课程开始时间
// topic 课程名字
// teacherNickName
// 至少两种类型的组合
// 密码长度至少8位 在外边验证
func JoinMeeting(
	courseTimeId int64,
	startTime gox.Timestamp,
	duration int64,
	topic string,
	hostName, hostMobile string,
	meetingHost string,
) (data *MeetingData, err error) {
	var (
		resp *req.Resp
	)

	url := fmt.Sprintf("%s/api/meetings/join", meetingHost)
	params := req.Param{
		"startTime":    startTime,
		"duration":     duration,
		"topic":        topic,
		"hostMobile":   hostMobile,
		"hostName":     hostName,
		"courseTimeId": courseTimeId,
	}

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}

	if resp.Response().StatusCode != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = resp.ToJSON(&data); nil != err {
		return
	}

	return
}

func EndMeeting(userId int64, meetingId string, meetingHost string) (data *MeetingData, err error) {
	var (
		resp *req.Resp
	)
	url := fmt.Sprintf("%s/api/meetings/ends", meetingHost)
	params := req.Param{
		"userId":    userId,
		"meetingId": meetingId,
	}

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}

	if resp.Response().StatusCode != http.StatusOK {
		err = getErr(resp)
		return
	}

	if err = resp.ToJSON(&data); err != nil {
		return
	}
	return
}
