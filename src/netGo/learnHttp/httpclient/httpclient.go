package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情 ...
	// 发起HTTP请求
	// 添加一些域到req.Header中 (作为代理、网关时有用)
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t} //Transport对象为RoundTripper接口
}

func main() {
	t := &OurCustomTransport{
	//...
	}

	reqest, _ := http.NewRequest("GET", "http://localhost:8080/foo", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	c := t.Client()
	//	c.Do(reqest)
	resp, err := c.Do(reqest)
	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	//	fmt.Println(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
}
