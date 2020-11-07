package ysx

import (
	`github.com/storezhang/gox`
)

type VirtualUser struct {
	gox.BaseStruct `xorm:"extends"`

	// UserId 用户编号
	UserId int64 `xorm:"bigint default(0)" json:"userId,string"`
	// 名称
	NickName string `xorm:"varchar(32) notnull default('')" json:"nickName"`
	// AppId 产品编号
	AppId int64 `xorm:"bigint default(0)" json:"appId,string"`
	// 课程时刻id
	CourseTimeId int64 `xorm:"bigint(20) notnull default(1)" json:"courseTimeId,string"`
	// 会议用虚拟手机号
	VirtualPhone string `xorm:"varchar(64) notnull default('')" json:"virtualPhone"`
	// 虚拟用户id
	VirtualUserId string `xorm:"varchar(64) notnull default('')" json:"VirtualUserId"`
}
