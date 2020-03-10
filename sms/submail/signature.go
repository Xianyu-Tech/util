package submail

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strings"
)

func caculSign(params url.Values, config map[string]string) string {
	keys := make([]string, 0, 32)

	for key, _ := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	values := make([]string, 0, 32)

	for _, key := range keys {
		if len(params[key]) > 0 {
			values = append(values, fmt.Sprintf("%s=%s", key, params[key][0]))
		}
	}

	signature := strings.Join(values, "&")

	appKey := config["appkey"]
	appId := config["appid"]
	signType := config["signType"]

	signature = appId + appKey + signature + appId + appKey

	if signType == "md5" {
		mymd5 := md5.New()
		io.WriteString(mymd5, signature)

		return hex.EncodeToString(mymd5.Sum(nil))
	} else if signType == "sha1" {
		mysha1 := sha1.New()
		io.WriteString(mysha1, signature)

		return hex.EncodeToString(mysha1.Sum(nil))
	}

	return appKey
}
