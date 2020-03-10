package submail

import (
	"fmt"
	"net/url"

	"github.com/json-iterator/go"
)

type SmsMultiXSend struct {
	appId    string
	appKey   string
	signType string

	project string
	multi   []map[string]*SmsMulti
	tag     string
}

type SmsMulti struct {
	to   string
	vars map[string]string
}

func NewSmsMulti() *SmsMulti {
	return &SmsMulti{"", make(map[string]string)}
}

func (this *SmsMulti) SetTo(to string) {
	this.to = to
}

func (this *SmsMulti) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *SmsMulti) Get() map[string]interface{} {
	item := make(map[string]interface{})

	item["to"] = this.to
	item["vars"] = this.vars

	return item
}

func NewSmsMultiXSend(appid int, appkey, signtype string) *SmsMultiXSend {
	return &SmsMultiXSend{
		appId:    fmt.Sprintf("%d", appid),
		appKey:   appkey,
		signType: signtype,
	}
}

func (this *SmsMultiXSend) SetProject(project string) {
	this.project = project
}

func (this *SmsMultiXSend) AddMulti(multi map[string]*SmsMulti) {
	this.multi = append(this.multi, multi)
}

func (this *SmsMultiXSend) SetTag(tag string) {
	this.tag = tag
}

func (this *SmsMultiXSend) MultiXSend() ([]string, error) {
	config := make(map[string]string)

	config["appid"] = this.appId
	config["appkey"] = this.appKey
	config["signType"] = this.signType

	params := url.Values{}

	params.Set("appid", this.appId)
	params.Set("project", this.project)

	if this.signType != "normal" {
		timestamp, err := GetTimestamp()

		if err != nil {
			return nil, err
		}

		params.Set("sign_type", this.signType)
		params.Set("timestamp", fmt.Sprintf("%d", timestamp))
		params.Set("sign_version", "2")
	}

	if this.tag != "" {
		params.Set("tag", this.tag)
	}

	signature := caculSign(params, config)

	params.Set("signature", signature)

	//v2 数字签名 multi 不参与计算
	data, err := jsoniter.Marshal(this.multi)

	if err != nil {
		return nil, err
	}

	params.Set("multi", string(data))

	respData, err := HttpPost(SUBMAIL_SMS_XSEND_URL, params.Encode())

	if err != nil {
		return nil, err
	}

	smsResps := make([]*SmsResp, 0)

	err = jsoniter.Unmarshal([]byte(respData), &smsResps)

	if err != nil {
		return nil, err
	}

	recvs := make([]string, 0)

	for _, smsResp := range smsResps {
		if smsResp.Status == "success" {
			recvs = append(recvs, smsResp.To)
		}
	}

	return recvs, nil
}
