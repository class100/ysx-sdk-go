package ysx

import (
	`github.com/class100/core`
)

type meeting interface {
	// JoinMeeting 加入会议
	JoinMeeting(jmr *JoinMeetingReq) (rsp *MeetingData, err error)
	// EndMeeting 结束会议
	EndMeeting(emr *EndMeetingReq) (rsp *MeetingData, err error)
}

func (hsc *httpSignatureClient) JoinMeeting(jmr *JoinMeetingReq) (rsp *MeetingData, err error) {
	rsp = new(MeetingData)
	err = hsc.requestApi(
		ApiPathMeetingJoin,
		core.HttpMethodPost,
		jmr,
		ApiVersionDefault,
		&rsp,
	)

	return
}

func (hsc *httpSignatureClient) EndMeeting(emr *EndMeetingReq) (rsp *MeetingData, err error) {
	rsp = new(MeetingData)
	err = hsc.requestApi(
		ApiPathMeetingEnd,
		core.HttpMethodPost,
		emr,
		ApiVersionDefault,
		&rsp,
	)

	return
}
