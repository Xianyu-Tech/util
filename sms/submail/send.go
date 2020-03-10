package submail

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/json-iterator/go"
)

type SmsSend struct {
	appId    string
	appKey   string
	signType string

	to      string
	content string
	tag     string
}

func NewSmsSend(appid int, appkey, signtype string) *SmsSend {
	return &SmsSend{
		appId:    fmt.Sprintf("%d", appid),
		appKey:   appkey,
		signType: signtype,
	}
}

func (this *SmsSend) SetTo(to string) {
	this.to = to
}

func (this *SmsSend) SetContent(content string) {
	this.content = content
}

func (this *SmsSend) SetTag(tag string) {
	this.tag = tag
}

func (this *SmsSend) Send() (string, error) {
	config := make(map[string]string)

	config["appid"] = this.appId
	config["appkey"] = this.appKey
	config["signType"] = this.signType

	params := url.Values{}

	params.Set("appid", this.appId)
	params.Set("to", this.to)

	if this.signType != "normal" {
		timestamp, err := GetTimestamp()

		if err != nil {
			return "", err
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

	//v2 数字签名 content 不参与计算
	params.Set("content", this.content)

	respData, err := HttpPost(SUBMAIL_SMS_XSEND_URL, params.Encode())

	if err != nil {
		return "", err
	}

	smsResp := &SmsResp{}

	err = jsoniter.Unmarshal([]byte(respData), smsResp)

	if err != nil {
		return "", err
	}

	if smsResp.Status != "success" {
		return "", errors.New(smsResp.Msg)
	}

	return this.to, nil
}
