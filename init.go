package ysxsdk

import (
	"github.com/imroc/req"
)

type Cfg struct {
	ApiKey       string `yaml:"apiKey"`
	ApiSecret    string `yaml:"apiSecret"`
	Key          string `yaml:"key"`
	EnterpriseId string `yaml:"enterpriseId"`
	Identity     string `yaml:"identity"`
	EcId         string `yaml:"ecId"`
	OrgCode      int    `yaml:"orgCode"`
	UserHost     string `yaml:"userHost"`
	MeetingHost  string `yaml:"meetingHost"`
}

var (
	User    *user
	Meeting *meeting
)

func Init(cfg *Cfg) {
	req.SetFlags(req.LrespHead | req.LrespBody | req.Lcost)
	req.Debug = true

	User = newUser(
		cfg.UserHost,
		cfg.Identity,
		cfg.EcId,
		cfg.Key,
		cfg.OrgCode,
	)
	Meeting = newMeeting(
		cfg.MeetingHost,
		cfg.ApiKey,
		cfg.ApiSecret,
		cfg.EnterpriseId,
	)
}
