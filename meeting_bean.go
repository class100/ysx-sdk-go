package ysx

import (
	`github.com/storezhang/gox`
)

type (
	// JoinMeetingReq 加入会议请求
	JoinMeetingReq struct {
		BaseVirtualUser

		// StartTime 开始时间
		StartTime gox.Timestamp `json:"startTime" validate:"required"`
		// Duration 持续时间
		// Topic 单位：分钟
		Duration int64 `json:"duration" validate:"required"`
		// Topic 主题
		Topic string `json:"topic" validate:"required,omitempty,min=1,max=64"`
	}

	// EndMeetingReq 结束会议请求
	EndMeetingReq struct {
		// UserId 用户ID
		UserId int64 `json:"userId" validate:"required"`
		// MeetingId 会议Id
		MeetingId string `json:"meetingId" validate:"required"`
		// AppId 产品编号
		AppId int64 `json:"appId" validate:"required"`
	}
)
