package ysxsdk

import (
	`fmt`

	`github.com/imroc/req`
	`github.com/storezhang/gox`
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
		// 主题
		Topic string `json:"topic" validate:"required,omitempty,min=1,max=64"`
		// 主持人手机号
		HostMobile string `json:"hostMobile" validate:"required,omitempty,alphanum,max=15"`
		// 主持人名称
		HostName string `json:"hostName" validate:"required,omitempty,min=2,max=32"`
	}

	JoinMeetingResp struct {
		// 用户ID
		UserID string `json:"userId"`
		// 虚拟手机号
		VirtualMobile string `json:"virtualMobile"`
		// 用户token
		UserToken string `json:"userToken"`
		// 会议Id
		MeetingId string `json:"meetingId"`
		// 会议No
		MeetingNo uint64 `json:"meetingNo"`
	}

	EndMeetingReq struct {
		// 开始时间
		StartTime gox.Timestamp `json:"startTime" validate:"required"`
		// 用户ID
		UserID string `json:"userId"`
		// 会议Id
		MeetingId string `json:"meetingId"`
		// 主持人手机号
		HostMobile string `json:"hostMobile" validate:"required,alphanum,max=15"`
	}

	EndMeetingResp struct {
		Data *MeetingData `json:"data"`
	}
)

// startTime 课程开始时间
// topic 课程名字
// teacherNickName
// 至少两种类型的组合
// 密码长度至少8位 在外边验证
func JoinMeeting(startTime gox.Timestamp, topic string, hostName string, hostMobile string, meetingHost string) (data *JoinMeetingResp, err error) {
	var (
		resp *req.Resp
	)
	url := fmt.Sprintf("%s/api/meetings/join", meetingHost)
	params := req.Param{
		"startTime":  startTime,
		"topic":      topic,
		"hostMobile": hostMobile,
		"hostName":   hostName,
	}

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(&data); nil != err {
		return
	}
	return
}

func EndMeeting(startTime gox.Timestamp, userId int64, meetingId string,
	hostMobile string, meetingHost string) (data *EndMeetingResp, err error) {
	var (
		resp *req.Resp
	)
	url := fmt.Sprintf("%s/api/meetings/ends", meetingHost)
	params := req.Param{
		"startTime":  startTime,
		"userId":     userId,
		"meetingId":  meetingId,
		"hostMobile": hostMobile,
	}

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}
	if err = resp.ToJSON(&data); err != nil {
		return
	}
	return
}