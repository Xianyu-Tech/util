package netutil

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type HttpClient struct {
	_client *http.Client
}

func NewHttpClient(timeout int) *HttpClient {
	httpClient := &HttpClient{}

	transport := &http.Transport{
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	httpClient._client = &http.Client{
		Transport: transport,
	}

	if timeout > 0 {
		httpClient._client.Timeout = time.Millisecond * time.Duration(timeout)
	}

	return httpClient
}

func NewHttpClientWithMaxIdleConns(maxidleconns, timeout int) *HttpClient {
	httpClient := &HttpClient{}

	transport := &http.Transport{
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},

		MaxIdleConns: maxidleconns,
	}

	httpClient._client = &http.Client{
		Transport: transport,
	}

	if timeout > 0 {
		httpClient._client.Timeout = time.Millisecond * time.Duration(timeout)
	}

	return httpClient
}

func (this *HttpClient) Request(method, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	retry := 3
	var resp *http.Response

	for i := 0; i < retry; i++ {
		resp, err = this._client.Do(req)

		if err != nil {
			time.Sleep(100 * time.Millisecond)

			continue
		} else {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (this *HttpClient) HttpHead(url string, refer string, headers map[string]string) (int, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("HEAD", url, headers, nil)

	if err != nil {
		return -1, -1, err
	}

	srcContentLen := resp.Header.Get("Content-Length")

	if srcContentLen == "" {
		return resp.StatusCode, 0, nil
	}

	contentLen, err := strconv.ParseInt(srcContentLen, 10, 32)

	if err != nil {
		return -1, -1, err
	}

	return resp.StatusCode, int(contentLen), nil
}

func (this *HttpClient) HttpGet(url string, refer string, headers map[string]string) (int, string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("GET", url, headers, nil)

	if err != nil {
		return -1, "", err
	}

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	defer resp.Body.Close()
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

func (this *HttpClient) HttpGetLen(url string, refer string, headers map[string]string) (int64, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("GET", url, headers, nil)

	if err != nil {
		return 0, nil
	}

	contentLen := resp.Header.Get("Content-Length")

	length, err := strconv.ParseInt(contentLen, 10, 64)

	if err != nil {
		return 0, err
	}

	return length, nil
}

//Post
func (this *HttpClient) HttpPost(url string, refer string, headers map[string]string, params string) (int, string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("POST", url, headers, bytes.NewBuffer([]byte(params)))

	if err != nil {
		return -1, "", err
	}

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	defer resp.Body.Close()
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

func (this *HttpClient) HttpPut(url string, refer string, headers map[string]string, params string) (int, string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("PUT", url, headers, bytes.NewBuffer([]byte(params)))

	if err != nil {
		return -1, "", err
	}

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	defer resp.Body.Close()
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}

// Delete
func (this *HttpClient) HttpDelete(url string, refer string, headers map[string]string) (int, string, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	if refer != "" {
		headers["Referer"] = refer
	}

	resp, err := this.Request("DELETE", url, headers, nil)

	if err != nil {
		return -1, "", err
	}

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	defer resp.Body.Close()
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return -1, "", err
	}

	return resp.StatusCode, string(body), nil
}
