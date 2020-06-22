package ysxsdk

import (
	"fmt"
	"time"

	"github.com/imroc/req"
)

type user struct {
	host     string
	identity string
	orgCode  int
	ecid     string
	key      string
}

func newUser(host, identity, ecid, key string, orgCode int) *user {
	return &user{
		host:     host,
		identity: identity,
		orgCode:  orgCode,
		ecid:     ecid,
		key:      key,
	}
}

type VirtualUserData struct {
	VirtualMobile string `json:"vitualMobile"`
	UserID        string `json:"userId"`
}

type VirtualUserResult struct {
	Code    int             `json:"code"`
	Message string          `json:"msg"`
	Data    VirtualUserData `json:"data"`
}

func (u *user) Create(mobile string, name string, isTrial int8) (data *VirtualUserData, err error) {
	var (
		resp   *req.Resp
		result VirtualUserResult
		url    = fmt.Sprintf("%s/access/rest/v200/createVirtualUser", u.host)

		params = req.Param{
			"mobile":    mobile,
			"isTrial":   isTrial,
			"name":      name,
			"identity":  u.identity,
			"orgCode":   u.orgCode,
			"ECID":      u.ecid,
			"timestamp": time.Now().UnixNano() / 1e6,
		}
		key = u.key
	)

	params["sign"] = getUserSign(params, key)

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

func (u *user) Update(mobile string, name string) (data *VirtualUserData, err error) {
	var (
		resp   *req.Resp
		result VirtualUserResult
		url    = fmt.Sprintf("%s/access/rest/v200/modifyVirtualUserInfo", u.host)

		params = req.Param{
			"mobile":    mobile,
			"name":      name,
			"identity":  u.identity,
			"timestamp": time.Now().UnixNano() / 1e6,
		}
	)

	params["sign"] = getUserSign(params, u.key)

	if resp, err = req.Post(url, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(&result); err != nil {
		return nil, err
	}

	if 200 != result.Code {
		return nil, fmt.Errorf("{code=%d, msg=%s}", result.Code, result.Message)
	}

	return &result.Data, nil
}

type tokenResult struct {
	Code         int    `json:"code"`
	Message      string `json:"msg"`
	EnterpriseID string `json:"enterpriseId"`
	UserID       string `json:"userId"`
	Username     string `json:"username"`
	Token        string `jsn:"token"`
}

func (u *user) GetToken(virtualMobile string) (token string, err error) {
	var (
		resp *req.Resp
		data tokenResult
	)

	url := fmt.Sprintf("%s/access/rest/v200/token/getByVitualMobile", u.host)
	params := req.Param{
		"mobile":    virtualMobile,
		"identity":  u.identity,
		"timestamp": time.Now().UnixNano() / 1e6,
	}

	key := u.key
	params["sign"] = getUserSign(params, key)

	if resp, err = req.Get(url, params); nil != err {
		return
	}

	if err = resp.ToJSON(&data); err != nil {
		return
	}

	if 200 != data.Code {
		return "", fmt.Errorf("{code=%d, msg=%s}", data.Code, data.Message)
	}
	token = data.Token

	return
}

type DeleteResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (u *user) DeleteVirtualMobile(mobile string) (*DeleteResult, error) {
	var (
		data DeleteResult
	)

	url := fmt.Sprintf("%s/access/rest/v200/deleteVirtualUserInfo", u.host)
	params := req.Param{
		"virtualMobile": mobile,
		"identity":      u.identity,
		"orgCode":       u.orgCode,
		"timestamp":     time.Now().UnixNano() / 1e6,
	}

	key := u.key
	params["sign"] = getUserSign(params, key)

	resp, err := req.Delete(url, req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	if err = resp.ToJSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
