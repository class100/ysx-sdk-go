package ysxsdk

import (
	"fmt"
	"github.com/imroc/req"
	"strings"
	"time"
)

type meeting struct {
	host      string
	apiKey    string
	apiSecret string
	apiEcid   string
}

type MeetingData struct {
	Id        string `json:"Id"`
	MeetingNo uint64 `json:"MeetingNo"`
}

type MeetingResult struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Data    MeetingData `json:"Data"`
}

func newMeeting(host, apiKey, apiSecret, apiEcid string) *meeting {
	return &meeting{
		host:      host,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		apiEcid:   apiEcid,
	}
}

func (m *meeting) Create(startTime int64, topic, hostId string) (data *MeetingData, err error) {
	var (
		url    = fmt.Sprintf("%s/v20/meeting/createScheduledMeeting", m.host)
		header = req.Header{
			"Authorization": fmt.Sprintf("Bearer %s", getAPIToken(m.apiKey, m.apiSecret, m.apiEcid)),
			"Content-Type":  "application/json; charset=utf-8",
		}
		params = req.Param{
			"Topic":                 topic,
			"Agenda":                "",
			"Duration":              360,
			"UTCStartTime":          time.Unix(startTime, 0).In(time.Local).Format("2006-01-02 15:04:05"),
			"LocalStartTime":        time.Unix(startTime, 0).In(time.Local).UTC().Format("2006-01-02 15:04:05"),
			"HostId":                hostId,
			"OpenHostVideo":         true,
			"OpenParticipantsVideo": true,
		}

		resp   *req.Resp
		result = new(MeetingResult)
	)

	if resp, err = req.Post(url, header, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(result); nil != err {
		return
	}

	if 0 != result.Code {
		return nil, fmt.Errorf("{code=%d, msg=%s}", result.Code, result.Message)
	}

	return &result.Data, nil
}

func (m *meeting) UpdateScheduleMeeting(startTime int64, ID, topic, hostID string, participants []string) (data *MeetingData, err error) {
	var (
		url    = fmt.Sprintf("%s/v20/meeting/update", m.host)
		header = req.Header{
			"Authorization": fmt.Sprintf("Bearer %s", getAPIToken(m.apiKey, m.apiSecret, m.apiEcid)),
			"Content-Type":  "application/json; charset=utf-8",
		}

		params = req.Param{
			"id":                    ID,
			"Topic":                 topic,
			"Agenda":                "",
			"Duration":              360,
			"UTCStartTime":          time.Unix(startTime, 0).In(time.Local).Format("2006-01-02 15:04:05"),
			"LocalStartTime":        time.Unix(startTime, 0).In(time.Local).UTC().Format("2006-01-02 15:04:05"),
			"HostId":                hostID,
			"Participants":          strings.Join(participants, ","),
			"OpenHostVideo":         true,
			"OpenParticipantsVideo": true,
		}

		resp   *req.Resp
		result = new(MeetingResult)
	)

	if resp, err = req.Post(url, header, req.BodyJSON(params)); nil != err {
		return
	}

	if err = resp.ToJSON(&data); err != nil {
		return
	}

	if 0 != result.Code {
		return nil, fmt.Errorf("{code=%d, msg=%s}", result.Code, result.Message)
	}

	return &result.Data, nil
}

func (m *meeting) DeleteMeeting(id, hostID string) (*MeetingResult, error) {
	var (
		data MeetingResult
	)

	url := fmt.Sprintf("%s/v20/meeting/delete", m.host)
	header := req.Header{
		"Authorization": fmt.Sprintf("Bearer %s", getAPIToken(m.apiKey, m.apiSecret, m.apiEcid)),
		"Content-Type":  "application/json; charset=utf-8",
	}
	params := req.Param{
		"Id":     id,
		"HostId": hostID,
	}
	resp, err := req.Post(url, header, req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	if err = resp.ToJSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *meeting) EndMeeting(id, hostID string) (*MeetingResult, error) {
	var (
		data MeetingResult
	)

	url := fmt.Sprintf("%s/v20/meeting/end", m.host)
	header := req.Header{
		"Authorization": fmt.Sprintf("Bearer %s", getAPIToken(m.apiKey, m.apiSecret, m.apiEcid)),
		"Content-Type":  "application/json; charset=utf-8",
	}
	params := req.Param{
		"Id":     id,
		"HostId": hostID,
	}

	resp, err := req.Post(url, header, req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	if err = resp.ToJSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *meeting) GetMeeting(id string) (*MeetingResult, error) {
	var (
		data MeetingResult
	)

	url := fmt.Sprintf("%s/v20/meeting/get", m.host)
	header := req.Header{
		"Authorization": fmt.Sprintf("Bearer %s", getAPIToken(m.apiKey, m.apiSecret, m.apiEcid)),
		"Content-Type":  "application/json; charset=utf-8",
	}
	params := req.Param{
		"Id": id,
	}

	req.Debug = true
	resp, err := req.Post(url, header, req.BodyJSON(params))
	if err != nil {
		return nil, err
	}

	if err = resp.ToJSON(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
