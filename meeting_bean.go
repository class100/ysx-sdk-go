package ysx

import (
	`github.com/storezhang/gox`
)

type (
	// JoinMeetingReq 加入会议请求
	JoinMeetingReq struct {
		// UserId 用户Id
		UserId int64 `json:"userId,string" validate:"required"`
		// NickName 用户昵称
		NickName string `json:"nickName" validate:"required,min=2,max=32"`
		// StartTime 开始时间
		StartTime gox.Timestamp `json:"startTime" validate:"required"`
		// Duration 持续时间
		// Topic 单位：分钟
		Duration int64 `json:"duration" validate:"required"`
		// Topic 主题
		Topic string `json:"topic" validate:"required,omitempty,min=1,max=64"`
		// CourseTimeId 课程时刻Id
		CourseTimeId int64 `json:"courseTimeId" validate:"omitempty"`
	}

	// EndMeetingReq 结束会议请求
	EndMeetingReq struct {
		// UserID 用户ID
		UserID string `json:"userId" validate:"required"`
		// MeetingId 会议Id
		MeetingId string `json:"meetingId" validate:"required"`
	}
)
