package submail

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"
)

type SmsResp struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`

	SendId     string `json:"send_id"`
	To         string `json:"to"`
	Fee        int    `json:"fee"`
	SmsCredits string `json:"sms_credits"`
}

func HttpGet(request string) (string, error) {
	resp, err := http.Get(request)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", nil
	}

	return string(body), nil
}

func HttpPost(request string, params string) (string, error) {
	resp, err := http.Post(request, "application/x-www-form-urlencoded;charset=utf-8", bytes.NewBuffer([]byte(params)))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", nil
	}

	return string(body), nil
}

func GetTimestamp() (int64, error) {
	resp, err := HttpGet(SUBMAIL_SVC_TIMESTAMP)

	if err != nil {
		return 0, err
	}

	var params map[string]interface{}

	err = jsoniter.Unmarshal([]byte(resp), &params)

	if err != nil {
		return 0, nil
	}

	return int64(params["timestamp"].(float64)), nil
}
