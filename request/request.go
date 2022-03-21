package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Post(Addr string, data, header map[string]string) []byte {
	queryData := make(url.Values)
	for i, v := range data {
		queryData.Add(i, v)
	}
	p := queryData.Encode()
	r, _ := http.NewRequest(http.MethodPost, Addr, strings.NewReader(p))
	for i, v := range header {
		r.Header.Set(i, v)
	}
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}
func Get(addr string) []byte {
	r, _ := http.NewRequest(http.MethodGet, addr, nil)
	r.Header.Set("content-type", "application/json")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req := http.Client{}
	resp, err := req.Do(r)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}

func PostJson(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Notice(title, body string) {
	var sendKey = "SCT131115T7zdUhOf2deIWEAUo2xwgyR21"
	var addr = "https://sctapi.ftqq.com/"
	Get(addr + sendKey + ".send?title=" + title + "&desp=" + body)
}
