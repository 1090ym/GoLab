package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

func TestReptile(t *testing.T) {
	// 从网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError("http get url", err)
	defer resp.Body.Close()

	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError("read page-bytes fail", err)

	pageStr := string(pageBytes)

	// 过滤数据，-1代表取全部
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Println(results)

	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("name:", result[1])
	}
}

func HandleError(why string, err error) {
	if err != nil {
		fmt.Println(why, err)
	}
}
