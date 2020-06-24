package ysx

import (
	"fmt"
	"time"

	"github.com/imroc/req"
)

type (
	// 虚拟手机号与ID
	VirtualUserData struct {
		VirtualMobile string `json:"vitualMobile"`
		UserId        string `json:"userId"`
	}

	deleteResult struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	virtualUserResult struct {
		Code    int             `json:"code"`
		Message string          `json:"msg"`
		Data    VirtualUserData `json:"data"`
	}

	tokenResult struct {
		Code         int    `json:"code"`
		Message      string `json:"msg"`
		EnterpriseID string `json:"enterpriseId"`
		UserID       string `json:"userId"`
		Username     string `json:"username"`
		Token        string `jsn:"token"`
	}

	virtualUser struct {
		host     string
		identity string
		orgCode  int
		ecid     string
		key      string
	}
)

func newUser(host, identity, ecid, key string, orgCode int) *virtualUser {
	return &virtualUser{
		host:     host,
		identity: identity,
		orgCode:  orgCode,
		ecid:     ecid,
		key:      key,
	}
}

// 创建
func (u *virtualUser) Create(mobile string, name string, isTrial int8) (data *VirtualUserData, err error) {
	url := fmt.Sprintf("%s/access/rest/v200/createVirtualUser", u.host)
	params := req.Param{
		"mobile":    mobile,
		"isTrial":   isTrial,
		"name":      name,
		"identity":  u.identity,
		"orgCode":   u.orgCode,
		"ECID":      u.ecid,
		"timestamp": time.Now().UnixNano() / 1e6,
	}

	return u.postReq(url, params)
}

//更新
func (u *virtualUser) Update(mobile string, name string) (data *VirtualUserData, err error) {
	url := fmt.Sprintf("%s/access/rest/v200/modifyVirtualUserInfo", u.host)
	params := req.Param{
		"mobile":    mobile,
		"name":      name,
		"identity":  u.identity,
		"timestamp": time.Now().UnixNano() / 1e6,
	}

	return u.postReq(url, params)
}

// 删除
func (u *virtualUser) Delete(virtualMobile string) error {
	url := fmt.Sprintf("%s/access/rest/v200/deleteVirtualUserInfo", u.host)
	params := req.Param{
		"virtualMobile": virtualMobile,
		"identity":      u.identity,
		"orgCode":       u.orgCode,
		"timestamp":     time.Now().UnixNano() / 1e6,
	}

	return u.deleteReq(url, params)
}

// 获取token
func (u *virtualUser) GetToken(virtualMobile string) (token string, err error) {
	var data tokenResult

	url := fmt.Sprintf("%s/access/rest/v200/token/getByVitualMobile", u.host)
	params := req.Param{
		"mobile":    virtualMobile,
		"identity":  u.identity,
		"timestamp": time.Now().UnixNano() / 1e6,
	}
	if err = u.getReq(url, params, &data); nil != err {
		return
	}

	if 200 != data.Code {
		return "", fmt.Errorf("{code=%d, msg=%s}", data.Code, data.Message)
	}
	token = data.Token

	return
}

func (u *virtualUser) postReq(url string, params req.Param) (data *VirtualUserData, err error) {
	var (
		resp   *req.Resp
		result virtualUserResult
	)

	params["sign"] = getUserSign(params, u.key)

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(&result); err != nil {
		return
	}

	if 200 != result.Code {
		return nil, fmt.Errorf("{code=%d, msg=%s}", result.Code, result.Message)
	}

	return &result.Data, nil
}

func (u *virtualUser) deleteReq(url string, params req.Param) (err error) {
	var (
		resp   *req.Resp
		result deleteResult
	)
	params["sign"] = getUserSign(params, u.key)

	if resp, err = req.Delete(url, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(&result); err != nil {
		return
	}

	if 200 != result.Code {
		return fmt.Errorf("{code=%d, msg=%s}", result.Code, result.Msg)
	}

	return
}

func (u *virtualUser) getReq(url string, params req.Param, data interface{}) (err error) {
	var resp *req.Resp

	params["sign"] = getUserSign(params, u.key)

	if resp, err = req.Get(url, params); nil != err {
		return
	}

	return resp.ToJSON(data)
}
