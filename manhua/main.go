package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"github.com/robertkrimen/otto"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	s1 string
	s2 string
)

func main() {
	fmt.Println("输入你要爬的漫画的网站")
	url := ""
	fmt.Scanln(&url)
	if url == "" {
		url = "https://www.dmzj.com/info/benghuai3rd.html"
	}
	download(url)
}

func GetResponse(url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("连接" + url + "出错" + "-----" + err.Error())
		return resp, err
	}
	return resp, err
}

func download(url string) {
	l := 1
	vm := otto.New()
	resp, err := GetResponse(url)
	if err != nil {
		os.Exit(1)
	}
	d, _ := goquery.NewDocumentFromResponse(resp)
	s1 := d.Find("h1").Text()
	os.Mkdir(s1, os.ModePerm)
	n := d.Find("div.tab-content").Eq(1).Find("a").Size()
	fmt.Println("一共有" + strconv.Itoa(n) + "话")
	d.Find("div.tab-content").Eq(1).Find("a").Each(func(i int, selection *goquery.Selection) {
		title, _ := selection.Attr("title")
		s2 = s1 + "/" + strconv.Itoa(l) + "：" + title
		os.Mkdir(s2, os.ModePerm)
		src, _ := selection.Attr("href")
		fmt.Println(src)
		resp, _ = GetResponse(src)
		d, _ = goquery.NewDocumentFromResponse(resp)
		d.Find("script").Each(func(i int, selection *goquery.Selection) {
			s := selection.Text()
			if strings.Contains(s, "eval") {
				s = s[strings.Index(s, "eval"):]
				s = "var url\n" + s
				index := strings.Index(s, "return p")
				s = s[:index] + "url=p\n" + s[index:]
				vm.Run(s)
				value, _ := vm.Get("url")
				s := value.String()[strings.Index(value.String(), "{") : strings.Index(value.String(), "}")+1]
				js, err := simplejson.NewJson([]byte(s))
				if err != nil {
					fmt.Println(err)
				}
				str, _ := js.Get("page_url").String()
				arr := strings.Split(str, "\r\n")
				for i, v := range arr {
					res, err := GetResponse("https://images.dmzj.com/" + v)
					if err == nil {
						img := s2 + "/" + strconv.Itoa(i) + ".jpg"
						comic, _ := os.Create(img)
						_, err = io.Copy(comic, res.Body)
					}
				}
			}
		})
		l++
		time.Sleep(3 * time.Second)
	})
}
