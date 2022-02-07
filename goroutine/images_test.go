package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	imageUrlsChan chan string
	wg            = sync.WaitGroup{}
	TaskChan      chan string
	reImgUrl      = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
	taskNum       = 10
	downLoadNum   = 5
	imageChanNum  = 10000
)

func TestImages(t *testing.T) {
	start := time.Now()
	imageUrlsChan = make(chan string, imageChanNum)
	TaskChan = make(chan string, taskNum)

	// 爬虫协程
	for i := 1; i < taskNum+1; i++ {
		wg.Add(1)
		go getImages("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	wg.Add(1)
	go checkOut()

	for i := 0; i < downLoadNum; i++ {
		wg.Add(1)
		go DownLoadImages()
	}
	wg.Wait()
	fmt.Println("over:", time.Since(start))
}

func DownLoadImages() {
	for url := range imageUrlsChan {
		filename := getFilename(url)
		ok := DownLoadFile(url, filename)
		if !ok {
			fmt.Println("下载失败", url)
		} else {
			fmt.Println("下载成功", url)
		}
	}
	wg.Done()
}

func DownLoadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError("http get url", err)
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError("read page-bytes fail", err)

	filename = "C:/Project/image/" + filename
	err = ioutil.WriteFile(filename, pageBytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

func getFilename(url string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

func checkOut() {
	count := 0
	for {
		url := <-TaskChan
		count++
		fmt.Println(url, "完成了爬取任务")
		if count == taskNum {
			fmt.Println("爬取结束")
			close(imageUrlsChan)
			break
		}
	}
	//close(TaskChan)
	wg.Done()
}

func getImages(url string) {
	pageStr := getPageStr(url)
	re := regexp.MustCompile(reImgUrl)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果", len(results))
	for _, result := range results {
		url := result[0]
		imageUrlsChan <- url
	}
	TaskChan <- url
	wg.Done()
}

func getPageStr(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError("http get url", err)
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError("read page-bytes fail", err)
	return string(pageBytes)
}
