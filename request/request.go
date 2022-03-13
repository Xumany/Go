package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(Addr string, data, header map[string]string) string {
	queryData := make(url.Values)
	for i, v := range data {
		queryData.Add(i, v)
	}
	p := queryData.Encode()
	r, _ := http.NewRequest("POST", Addr, strings.NewReader(p))
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
	fmt.Println(string(b))
	return string(b)
}
func Get(addr string) []byte {
	r, _ := http.NewRequest("POST", addr, nil)
	r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req := http.Client{}
	resp, err := req.Do(r)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}
