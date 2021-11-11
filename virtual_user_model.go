package ysx

import (
	`github.com/storezhang/gox`
)

type (
	// VirtualUser 虚拟用户数据
	VirtualUser struct {
		gox.BaseStruct  `xorm:"extends"`
		BaseVirtualUser `xorm:"extends"`

		// 会议用虚拟手机号
		VirtualPhone string `xorm:"varchar(64) notnull default('')" json:"virtualPhone"`
		// 虚拟用户id
		VirtualUserId string `xorm:"varchar(64) notnull default('')" json:"virtualUserId"`
	}

	// BaseVirtualUser 虚拟用户数据
	BaseVirtualUser struct {
		// UserId 用户Id
		UserId int64 `json:"userId,string" validate:"required"`
		// NickName 用户昵称
		Name string `json:"name" validate:"required,min=1,max=64"`
		// AppId 产品Id
		AppId int64 `json:"appId,string" validate:"required"`
		// CourseTimeId 课程时刻Id
		CourseTimeId int64 `json:"courseTimeId" validate:"omitempty"`
	}
)
