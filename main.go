package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type img struct {
	title string
	url   string
	fan   string
	date  string
	kind  string
}

var (
	imgs []img

	url string

	count = 0
)

func GetContents(url string) (string, error) {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	//增加header选项
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	//request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	resp, _ := client.Do(request)

	//resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get content failed status code is %d ", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func main() {
	fmt.Println("欢迎使用我的爬虫小软件，XX一时爽，一直XX一直爽")
	ch := make(chan int, 5)
	err := os.MkdirAll("学习资料/All", os.ModePerm)
	if err != nil {
		fmt.Println("创建文件出错")
		return
	}
	for i := 1; i < 2; i++ {
		url = "https://avmask.com/cn/page/" + strconv.Itoa(i)
		go download(url, ch)
	}
	for {
		count += <-ch
		if count == 4 {
			break
		}
	}
	//body, err := GetContents(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf(body)

}

//https://jp.netcdn.space/digital/video/urvrsp00030/urvrsp00030ps.jpg
func download(url string, ch chan int) {
	//var setu img
	d, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("连接" + url + "出错" + "-----" + err.Error())
		return
	}
	d.Find(".movie-box").Each(func(i int, selection *goquery.Selection) {
		u, _ := selection.Attr("href")
		d1, err := goquery.NewDocument(u)
		if err != nil {
			fmt.Println("连接" + u + "出错" + "-----" + err.Error())
			return
		}

		title := d1.Find("h3").Text()
		setu, _ := os.Create("学习资料/All/" + title + ".jpg")
		bigImage, _ := d1.Find(".bigImage").Attr("href")
		res, _ := http.Get(bigImage)
		_, err = io.Copy(setu, res.Body)
		os.Mkdir("学习资料/All/"+title, os.ModePerm)
		d1.Find(".sample-box").Each(func(i int, selection *goquery.Selection) {
			setu, _ := os.Create("学习资料/All/" + title + "/" + strconv.Itoa(i) + ".jpg")
			src, _ := selection.Attr("href")
			res, _ := http.Get(src)
			_, err = io.Copy(setu, res.Body)
			if err == nil {
				fmt.Println(title + "预览图" + strconv.Itoa(i) + "   下载成功")
			}
			time.Sleep(2 * time.Second)
		})
		//title:=d1.Find("h3").Text()
		//setu.url, _ = selection.Find("img").Attr("src")             //图片地址
		//setu.title, _ = selection.Find("img").Attr("title")         //标题
		//setu.fan = selection.Find("span date:first-of-type").Text() //番号
		//setu.date = selection.Find("span date:last-of-type").Text() //日期
		//os.Mkdir(setu.title+"-"+setu.fan+"-"+setu.date,)
		//src, _ := selection.Attr("src")
		//title, _ := selection.Attr("title")
		//res, err := http.Get(src)
		//if err != nil {
		//	fmt.Println("链接出错")
		//	return
		//}
		//f, err := os.Create(title + ".jpg")
		//if err != nil {
		//	fmt.Println("创造文件" + title + ".jpg出错")
		//	return
		//}
		//_, err = io.Copy(f, res.Body)
		//if err == nil {
		//	fmt.Println(title + "下载完成")
		//}

	})
	//d.Find("span date:first-of-type").Each(func(i int, selection *goquery.Selection) {
	//	os.Mkdir("abc", os.ModePerm)
	//	fmt.Println("番号是：" + selection.Text())
	//	fmt.Println("日期是" + selection.Next().Text())
	//})
	ch <- 1
}

//<title>AVMOO - 你的线上日本成人影片情报站。管理你的影片并分享你的想法。</title>
//<meta name="author" content="AVMOO">
//<meta name="keywords" content="">
//<meta name="description" content="AVMOO - 你的线上日本成人影片情报站。管理你的影片并分享你的想法。">
//<link rel="apple-touch-icon" href="https://avmask.com/app/jav/View/img/apple-touch-icon.png">
//<link rel="shortcut Icon" href="https://avmask.com/app/jav/View/img/favicon.ico">
//<link rel="bookmark" href="https://avmask.com/app/jav/View/img/favicon.ico">
//<link rel="dns-prefetch" href="https://jp.netcdn.space" />
//<link rel="dns-prefetch" href="https://us.netcdn.space" />
//<link rel="dns-prefetch" href="https://ads.exoclick.com" />
//<link rel="dns-prefetch" href="https://syndication.exoclick.com" />
//<link rel="dns-prefetch" href="https://adserver.juicyads.com" />
//<link rel="dns-prefetch" href="https://j.traffichunt.com" />
//<script language="javascript">function $ROOT_URL(){return "https://avmask.com"}function $APP(){return "jav"}function $APP_URL(){return "/index.php?app=jav&"}function $APP_INFO_URL(){return "/index.php/jav"}function $APP_REWRITE_URL(){return "/jav"}function $APP_VIEW_URL(){return "https://avmask.com/app/jav/View"}function $APP_UPLOAD_URL(){return "/app/jav/Upload"}</script><link rel='stylesheet' type='text/css' href='https://avmask.com/app/jav/View/css/app.min.css?v=1476953808'>
//<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries --><!--[if lt IE 9]><script src='https://avmask.com/script/html5shiv/3.7.2/html5shiv.min.js?v=1476953808'></script>
//<script src='https://avmask.com/script/respond/1.4.2/respond.min.js?v=1476953808'></script>
//<![endif]--><!--[if lt IE 8]><link rel='stylesheet' type='text/css' href='https://avmask.com/script/bootstrap.ie7/1.0/bootstrap.ie7.min.css?v=1476953808' >
//<![endif]--><script>
//(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
//(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
//m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
//})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
//ga('create', 'UA-74041965-4', {'sampleRate': 50});
//ga('send', 'pageview');
//</script>
//</head>
//
//  <body>
//    <nav class="navbar navbar-default navbar-fixed-top top-bar">
//      <div class="container-fluid">
//        <div class="navbar-header">
//                <a href="https://avmask.com/cn" class="logo"></a>
//            <div class="btn-group pull-right visible-xs-inline" role="group" style="margin-top:8px;margin-right:8px;">
//                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
//                  <span class="glyphicon glyphicon-globe"></span> 简体中文                  <span class="caret"></span>
//                </button>
//                <ul class="dropdown-menu" role="menu">
//                    <li><a href="https://avmask.com/en/page/1">English</a></li>
//                    <li><a href="https://avmask.com/ja/page/1">日本语</a></li>
//                    <li><a href="https://avmask.com/tw/page/1">正體中文</a></li>
//                    <li><a href="https://avmask.com/cn/page/1">简体中文</a></li>
//                </ul>
//            </div>
//        </div>
//
//        <div id="navbar" class="collapse navbar-collapse">
//          <form class="navbar-form navbar-left fullsearch-form" action="https://avmask.com/cn/search" onsubmit="return false">
//            <div class="input-group">
//              <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
//              <span class="input-group-btn">
//                <button class="btn btn-default" type="submit">搜寻</button>
//              </span>
//            </div>
//          </form>
//          <ul class="nav navbar-nav">
//            <li class="active"><a href="https://avmask.com/cn">全部</a></li>
//                        <li ><a href="https://avmask.com/cn/released">已发布</a></li>
//                        <li ><a href="https://avmask.com/cn/popular">热门</a></li>
//            <li ><a href="https://avmask.com/cn/actresses">女优</a></li>
//            <li ><a href="https://avmask.com/cn/genre">类别</a></li>
//          </ul>
//          <ul class="nav navbar-nav navbar-right">
//            <li class="dropdown">
//              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><span class="glyphicon glyphicon-globe" style="font-size:12px;"></span> <span class="hidden-sm">简体中文</span> <span class="caret"></span></a>
//              <ul class="dropdown-menu" role="menu">
//                <li><a href="https://avmask.com/en/page/1">English</a></li>
//                <li><a href="https://avmask.com/ja/page/1">日本语</a></li>
//                <li><a href="https://avmask.com/tw/page/1">正體中文</a></li>
//                <li><a href="https://avmask.com/cn/page/1">简体中文</a></li>
//              </ul>
//            </li>
//          </ul>
//        </div><!--/.nav-collapse -->
//      </div>
//    </nav>
//<div class="row visible-xs-inline footer-bar">
//    <div class="col-xs-3 text-center">
//        <a id="menu" class="btn btn-default trigger-overlay"><i class="glyphicon glyphicon-align-justify"></i></a>
//    </div>
//        <div class="col-xs-3 text-center">
//            </div>
//    <div class="col-xs-3 text-center">
//                <a id="prev" class="btn btn-default" href="/cn/page/2" style="display:none"><i class="glyphicon glyphicon-chevron-right"></i></a>
//            </div>
//        <div class="col-xs-3 text-center">
//        <a id="back" class="btn btn-default" href="javascript:window.history.back()"><i class="glyphicon glyphicon-share-alt flipx"></i></a>
//    </div>
//</div>
//    <div class="container-fluid">
//                <div class="row">
//                        <div id="waterfall">
//                                                                                                                                <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/69967e39f219e3fe">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/urvrsp00030/urvrsp00030ps.jpg" title="【VR】隣に住む憧れのお姉さんが無防備すぎて僕は、、、 佐藤ゆか">
//                        </div>
//                        <div class="photo-info">
//                               <span>【VR】隣に住む憧れのお姉さんが無防備すぎて僕は、、、 佐藤ゆか <br><date>URVRSP-030</date> / <date>2019-12-01</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/71efa47a8f27d447">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001ps.jpg" title="純潔ロリィタと性交 永瀬ゆい">
//                        </div>
//                        <div class="photo-info">
//                               <span>純潔ロリィタと性交 永瀬ゆい <br><date>BLD-001</date> / <date>2019-11-22</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/9b5cc05c0ae2ad80">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/ovr00004/ovr00004ps.jpg" title="【VR】私のうんち見てくれる？">
//                        </div>
//                        <div class="photo-info">
//                               <span>【VR】私のうんち見てくれる？ <br><date>OVR-004</date> / <date>2019-11-28</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/61ee9912360d6652">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/urvrsp00028/urvrsp00028ps.jpg" title="【VR】アニメ声でちょっとメンヘラな可愛すぎる僕の彼女が‘声我慢するからエッチしたい’と目をウルウルさせながら健気に僕を求める姿に大興奮！ゆう">
//                        </div>
//                        <div class="photo-info">
//                               <span>【VR】アニメ声でちょっとメンヘラな可愛すぎる僕の彼女が‘声我慢するからエッチしたい’と目をウルウルさせながら健気に僕を求める姿に大興奮！ゆう <br><date>URVRSP-028</date> / <date>2019-11-22</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/0c0e147fe7ad686f">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/dgcesd00842/dgcesd00842ps.jpg" title="★配信限定！特典映像付★新人の部下に調教されて快楽に溺れていった女上司 大槻ひびき">
//                        </div>
//                        <div class="photo-info">
//                               <span>★配信限定！特典映像付★新人の部下に調教されて快楽に溺れていった女上司 大槻ひびき <br><date>DGCESD-842</date> / <date>2019-12-08</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/b6eda87fe7f541b4">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/dgcesd00846/dgcesd00846ps.jpg" title="★配信限定！特典映像付★黒人生中出しNTR 浮気相手は宿泊客の外国人…そのデカマラに魅せられた人妻は夫に隠れて不倫SEX！！ みひな">
//                        </div>
//                        <div class="photo-info">
//                               <span>★配信限定！特典映像付★黒人生中出しNTR 浮気相手は宿泊客の外国人…そのデカマラに魅せられた人妻は夫に隠れて不倫SEX！！ みひな <br><date>DGCESD-846</date> / <date>2019-12-08</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/e7ef1a29689706ab">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00036/1svoks00036ps.jpg" title="由香里ちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>由香里ちゃん <br><date>SVOKS-036</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/9eaa50a3a23286d2">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00033/1svoks00033ps.jpg" title="もえちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>もえちゃん <br><date>SVOKS-033</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/8a5e2951fc624e4f">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00040/1svoks00040ps.jpg" title="実玖ちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>実玖ちゃん <br><date>SVOKS-040</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/0be340542b58ccdf">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00039/1svoks00039ps.jpg" title="Mao">
//                        </div>
//                        <div class="photo-info">
//                               <span>Mao <br><date>SVOKS-039</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/fb9154f157b46224">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00047/1svoks00047ps.jpg" title="ナナちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ナナちゃん <br><date>SVOKS-047</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/4c71991b87d7e453">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00035/1svoks00035ps.jpg" title="三奈">
//                        </div>
//                        <div class="photo-info">
//                               <span>三奈 <br><date>SVOKS-035</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/e1d1dba8a9071b53">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00056/1svoks00056ps.jpg" title="ゆりなちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ゆりなちゃん <br><date>SVOKS-056</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/1efdd692a01cff70">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00054/1svoks00054ps.jpg" title="みことちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>みことちゃん <br><date>SVOKS-054</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/fb90bde30ae974c6">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00050/1svoks00050ps.jpg" title="いちかちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>いちかちゃん <br><date>SVOKS-050</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/0603ee0f19f28496">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00046/1svoks00046ps.jpg" title="しおりちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>しおりちゃん <br><date>SVOKS-046</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/b58e6c24fcc229ef">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00037/1svoks00037ps.jpg" title="和歌ちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>和歌ちゃん <br><date>SVOKS-037</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/ff7f38db69fed59e">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00049/1svoks00049ps.jpg" title="カンナちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>カンナちゃん <br><date>SVOKS-049</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/15d815fd774fafd6">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00038/1svoks00038ps.jpg" title="菜水留ちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>菜水留ちゃん <br><date>SVOKS-038</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/e20d72aab5fa2bf5">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00043/1svoks00043ps.jpg" title="ウミちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ウミちゃん <br><date>SVOKS-043</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/0071803ffceb590a">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00041/1svoks00041ps.jpg" title="アリス">
//                        </div>
//                        <div class="photo-info">
//                               <span>アリス <br><date>SVOKS-041</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/e4d0040073392282">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00045/1svoks00045ps.jpg" title="みおちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>みおちゃん <br><date>SVOKS-045</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/ff793f1230fdd71a">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00051/1svoks00051ps.jpg" title="ののちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ののちゃん <br><date>SVOKS-051</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/c4b1e0b8fcb6ba8c">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00042/1svoks00042ps.jpg" title="ハルカちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ハルカちゃん <br><date>SVOKS-042</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/4db1a037ed8d4a83">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00044/1svoks00044ps.jpg" title="リナちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>リナちゃん <br><date>SVOKS-044</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/f7cb0406dcad23ed">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00034/1svoks00034ps.jpg" title="みことちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>みことちゃん <br><date>SVOKS-034</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/437d6c21106abe2f">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00055/1svoks00055ps.jpg" title="えなちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>えなちゃん <br><date>SVOKS-055</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/8b3f75dd4a52dcd5">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00052/1svoks00052ps.jpg" title="ゆずかちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>ゆずかちゃん <br><date>SVOKS-052</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/ea32fcf63f413e07">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00053/1svoks00053ps.jpg" title="もあちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>もあちゃん <br><date>SVOKS-053</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                        <div class="item">
//                        <a class="movie-box" href="https://avmask.com/cn/movie/10445288c9c7611a">
//                        <div class="photo-frame">
//                            <img src="https://jp.netcdn.space/digital/video/1svoks00048/1svoks00048ps.jpg" title="サトコちゃん">
//                        </div>
//                        <div class="photo-info">
//                               <span>サトコちゃん <br><date>SVOKS-048</date> / <date>2019-11-26</date></span>
//                            </div>
//                        </a>
//                                </div>
//                                                                        </div>
//                </div>
//    </div>
//             <div class="hidden-xs mtb-20 text-center">
//             <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTF83Mjh4OTAiLCJhZHNwb3QiOiJqX0xfNzI4eDkwIiwid2VpZ2h0IjoiMiIsImZjYXAiOmZhbHNlLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOiI3NjgiLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOmZhbHNlLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiPlxyXG52YXIgYWRfaWR6b25lID0gXCI4MTMzMDhcIixcclxuXHQgYWRfd2lkdGggPSBcIjcyOFwiLFxyXG5cdCBhZF9oZWlnaHQgPSBcIjkwXCI7XHJcbjxcL3NjcmlwdD5cclxuPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiIHNyYz1cImh0dHBzOlwvXC9hZHMuZXhvY2xpY2suY29tXC9hZHMuanNcIj48XC9zY3JpcHQ+XHJcbjxub3NjcmlwdD48YSBocmVmPVwiaHR0cDpcL1wvbWFpbi5leG9jbGljay5jb21cL2ltZy1jbGljay5waHA/aWR6b25lPTgxMzMwOFwiIHRhcmdldD1cIl9ibGFua1wiPjxpbWcgc3JjPVwiaHR0cHM6XC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9ODEzMzA4Jm91dHB1dD1pbWcmdHlwZT03Mjh4OTBcIiB3aWR0aD1cIjcyOFwiIGhlaWdodD1cIjkwXCI+PFwvYT48XC9ub3NjcmlwdD4ifSx7ImlkIjoianVpY19qX0xfNzI4eDkwIiwiYWRzcG90Ijoial9MXzcyOHg5MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTczNiBoZWlnaHQ9OTggc2Nyb2xsaW5nPW5vIGFsbG93dHJhbnNwYXJlbmN5PXRydWUgc3JjPVwvXC9hZHNlcnZlci5qdWljeWFkcy5jb21cL2Fkc2hvdy5waHA/YWR6b25lPTM3MTcwOD48XC9pZnJhbWU+In1d</div>             </div>
//
//             <div class="visible-xs-block text-center">
//             <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTV8zMDB4MjUwIiwiYWRzcG90Ijoial9NXzMwMHgyNTAiLCJ3ZWlnaHQiOiIyIiwiZmNhcCI6ZmFsc2UsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOiI3NjgiLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0PlxyXG5hZF9pZHpvbmUgPSBcIjEwMzEwNDJcIjtcclxuYWRfd2lkdGggPSBcIjMwMFwiO1xyXG5hZF9oZWlnaHQgPSBcIjI1MFwiO1xyXG5pZih0b3A9PT1zZWxmKSB2YXIgcD1kb2N1bWVudC5VUkw7IGVsc2UgdmFyIHA9ZG9jdW1lbnQucmVmZXJyZXI7dmFyIGR0PW5ldyBEYXRlKCkuZ2V0VGltZSgpO1xyXG52YXIgZXhvRG9jdW1lbnRQcm90b2NvbCA9IChkb2N1bWVudC5sb2NhdGlvbi5wcm90b2NvbCAhPSBcImh0dHBzOlwiICYmIGRvY3VtZW50LmxvY2F0aW9uLnByb3RvY29sICE9IFwiaHR0cDpcIikgPyBcImh0dHBzOlwiIDogZG9jdW1lbnQubG9jYXRpb24ucHJvdG9jb2w7XHJcbmlmKHR5cGVvZihhZF9zdWIpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfc3ViID0gXCJcIjtcclxuaWYodHlwZW9mKGFkX3RhZ3MpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfdGFncyA9IFwiXCI7XHJcbnZhciBhZF90eXBlID0gYWRfd2lkdGggKyAneCcgKyBhZF9oZWlnaHQ7XHJcbmlmKGFkX3dpZHRoID09ICcxMDAlJyAmJiBhZF9oZWlnaHQgPT0gJzEwMCUnKSBhZF90eXBlID0gJ2F1dG8nO1xyXG52YXIgYWRfc2NyZWVuX3Jlc29sdXRpb24gPSBzY3JlZW4ud2lkdGggKyAneCcgKyBzY3JlZW4uaGVpZ2h0O1xyXG5kb2N1bWVudC53cml0ZSgnPGlmcmFtZSBmcmFtZWJvcmRlcj1cIjBcIiBzY3JvbGxpbmc9XCJub1wiIHdpZHRoPVwiJyArIGFkX3dpZHRoICsgJ1wiIGhlaWdodD1cIicgKyBhZF9oZWlnaHQgKyAnXCIgc3JjPVwiJyArIGV4b0RvY3VtZW50UHJvdG9jb2wgKyAnXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9JyArIGFkX2lkem9uZSArICcmdHlwZT0nICsgYWRfdHlwZSArICcmcD0nICsgZXNjYXBlKHApICsgJyZkdD0nICsgZHQgKyAnJnN1Yj0nICsgYWRfc3ViICsgJyZ0YWdzPScgKyBhZF90YWdzICsgJyZzY3JlZW5fcmVzb2x1dGlvbj0nICsgYWRfc2NyZWVuX3Jlc29sdXRpb24gKyAnXCI+PFwvaWZyYW1lPicpO1xyXG48XC9zY3JpcHQ+In0seyJpZCI6Imp1aWNfal9NXzMwMHgyNTAiLCJhZHNwb3QiOiJqX01fMzAweDI1MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTMwOCBoZWlnaHQ9MjU4IHNjcm9sbGluZz1ubyBhbGxvd3RyYW5zcGFyZW5jeT10cnVlIHNyYz1cL1wvYWRzZXJ2ZXIuanVpY3lhZHMuY29tXC9hZHNob3cucGhwP2Fkem9uZT0zNzE3MjY+PFwvaWZyYW1lPiJ9XQ==</div>             </div>
//
//            <div class="text-center hidden-xs mtb-20">
//                            <ul class="pagination pagination-lg mtb-0">
//                    <li class="active"><a name="numbar"  href="/cn/page/1">1</a></li><li><a name="numbar"  href="/cn/page/2">2</a></li><li><a name="numbar"  href="/cn/page/3">3</a></li><li><a name="numbar"  href="/cn/page/4">4</a></li><li><a name="numbar"  href="/cn/page/5">5</a></li><li><a name="numbar"  href="/cn/page/6">6</a></li><li><a name="numbar"  href="/cn/page/7">7</a></li><li><a name="numbar"  href="/cn/page/8">8</a></li><li><a name="numbar"  href="/cn/page/9">9</a></li><li><a name="numbar"  href="/cn/page/10">10</a></li><li><a name="nextpage"  href="/cn/page/2">下一页 <span class="glyphicon phicon-chevron-right"></span></a></li>                </ul>
//                        </div>
//<div type="text/data-position" style="display:none">W3siaWQiOiJhZHN0X2pfUE9QVU5ERVIiLCJhZHNwb3QiOiJqX1BPUFVOREVSIiwid2VpZ2h0IjoiNSIsImZjYXAiOiIyIiwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxzY3JpcHQgdHlwZT0ndGV4dFwvamF2YXNjcmlwdCcgc3JjPSdodHRwczpcL1wvcGwxMTA1MjUucHVodG1sLmNvbVwvNzBcLzgyXC85Y1wvNzA4MjljMzgyMTZlMWMwNDYxNmFkYjQ2NzJlZGEzNDIuanMnPjxcL3NjcmlwdD4ifSx7ImlkIjoiY2xpY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjciLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IGRhdGEtY2Zhc3luYz1cImZhbHNlXCIgdHlwZT1cInRleHRcL2phdmFzY3JpcHRcIiBzcmM9XCJcL1wvY2xjYXNzZC5jb21cL3RcLzlcL2ZyZXRcL21lb3c0XC8zNjkwODlcL2JydC5qc1wiPjxcL3NjcmlwdD4ifSx7ImlkIjoiZXhvY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjYiLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6Ijc2OCIsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IHNyYz1cIlwvXC9zeW5kaWNhdGlvbi5leG9jbGljay5jb21cL3NwbGFzaC5waHA/aWR6b25lPTEwMDgwOTQmdHlwZT0zXCI+PFwvc2NyaXB0PiJ9LHsiaWQiOiJleG9tX2pfUE9QVU5ERVIiLCJhZHNwb3QiOiJqX1BPUFVOREVSIiwid2VpZ2h0IjoiNiIsImZjYXAiOiIyIiwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxzY3JpcHQgc3JjPVwiXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvc3BsYXNoLnBocD9pZHpvbmU9MTAyNjMxMiZ0eXBlPTExXCI+PFwvc2NyaXB0PlxyXG48c2NyaXB0PlxyXG4kKGRvY3VtZW50KS5yZWFkeShmdW5jdGlvbigpIHtcclxuICAgIGlmICh0eXBlb2YgZXhvVXJsICE9IFwidW5kZWZpbmVkXCIpIHtcclxuICAgICAgICBleG9VcmwgPSBleG9VcmwucmVwbGFjZSgnaHR0cDpcL1wvJywgJ2h0dHBzOlwvXC8nKTtcclxuICAgIH1cclxuICAgICQoXCJhXCIpLmJpbmQoXCJjbGlja1wiLCBmdW5jdGlvbihldmVudCkge1xyXG4gICAgICAgIGV4b01vYmlsZVBvcCgpO1xyXG4gICAgICAgICQodGhpcykudW5iaW5kKFwiY2xpY2tcIik7XHJcbiAgICB9KTtcclxufSk7XHJcbjxcL3NjcmlwdD4ifSx7ImlkIjoianVpY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjIiLCJmY2FwIjoiMSIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6Ijc2OCIsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6WyJ6aC1jbiJdLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPCEtLSBKdWljeUFkcyBQb3BVbmRlcnMgdjMgU3RhcnQgLS0+XHJcbjxzY3JpcHQgdHlwZT1cInRleHRcL2phdmFzY3JpcHRcIiBzcmM9XCJodHRwczpcL1wvanMuanVpY3lhZHMuY29tXC9qcC5waHA/Yz03NDU0MzN6MnQyNTZyMnEydzI4NDYzODQmdT1odHRwcyUzQSUyRiUyRmJ0c28ucHdcIj48XC9zY3JpcHQ+XHJcbjwhLS0gSnVpY3lBZHMgUG9wVW5kZXJzIHYzIEVuZCAtLT4ifV0=</div>
//<!-- Modal -->
//<div class="modal fade" id="advertisingModal" tabindex="-1" role="dialog" aria-labelledby="advertisingModalLabel" aria-hidden="true">
//  <div class="modal-dialog">
//    <div class="modal-content">
//      <div class="modal-header">
//        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
//        <h4 class="modal-title" id="advertisingModalLabel">Advertising</h4>
//      </div>
//      <div class="modal-body">
//            <p>Please contact following agents for advertising on AVMOO</p>
//            <p><a href="https://www.exoclick.com/?login=james666" target="_blank" style="color: #D80456;">ExoClick</a> / <a href="https://manage.juicyads.com/juicysites.php?id=128293" target="_blank" style="color: #D80456;">JuicyAds</a> / <a href="http://www.clickadu.com/?rfd=0l1" target="_blank" style="color: #D80456;">ClickADu</a></p>
//      </div>
//      <div class="modal-footer">
//        <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
//      </div>
//    </div>
//  </div>
//</div>
//
//<footer class="footer hidden-xs">
//        <div class="container-fluid">
//<p><a href="https://avmask.com/cn/terms">Terms</a> / <a href="https://avmask.com/cn/privacy">Privacy</a> / <a href="https://avmask.com/cn/usc">2257</a> / <a href="http://www.rtalabel.org/" target="_blank" rel="external nofollow">RTA</a> / <a href="#advertisingModal" role="button" data-toggle="modal">Advertising</a> / <a class="contactus" href="javascript:;" role="button" data-toggle="modal">Contact</a> / <a href="https://tellme.pw/avmoo" target="_blank">Guide</a> | Links: <a href="https://avmask.com" target="_blank">AVMOO</a> / <a href="https://avsox.asia" target="_blank">AVSOX</a> / <a href="https://avmemo.asia" target="_blank">AVMEMO</a><br>Copyright © 2013 AVMOO. All Rights Reserved. All other trademarks and copyrights are the property of their respective holders. The reviews and comments expressed at or through this website are the opinions of the individual author and do not reflect the opinions or views of AVMOO. AVMOO is not responsible for the accuracy of any of the information supplied here.</p>
//        </div>
//</footer>
//
//<div class="visible-xs-block footer-bar-placeholder"></div>
//
//<!-- ////////////////////////////////////////////////// -->
//<div class="overlay overlay-contentscale">
//    <nav>
//        <ul>
//            <li>
//            <form class="fullsearch-form" action="https://avmask.com/cn/search" onsubmit="return false">
//               <div class="input-group col-xs-offset-2 col-xs-8">
//                  <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
//                  <span class="input-group-btn">
//                    <button class="btn btn-default" type="submit">搜寻</button>
//                  </span>
//               </div>
//            </form>
//            </li>
//            <li><a href="https://avmask.com/cn/released">已发布</a></li>
//            <li><a href="https://avmask.com/cn/popular">热门</a></li>
//            <li><a href="https://avmask.com/cn/actresses">女优</a></li>
//            <li><a href="https://avmask.com/cn/genre">类别</a></li>
//        </ul>
//    </nav>
//    <div class="row overlay-close"><i class="glyphicon glyphicon-remove" style="color:#fff;font-size: 24px;margin:30px;"></i></div>
//</div>
//<script src='https://avmask.com/app/jav/View/js/app.min.js?v=1476953808'></script>  </body>
//</html>

//<!DOCTYPE html>
//<html lang="en">
//<head>
//<meta charset="utf-8">
//<meta name="trafficjunky-site-verification" content="445w343y8" />
//<meta http-equiv="X-UA-Compatible" content="IE=edge">
//<meta http-equiv="x-dns-prefetch-control" content="on">
//<meta name="renderer" content="webkit">
//<meta name="viewport" content="width=device-width, initial-scale=1">
//<title>BLD-001 純潔ロリィタと性交 永瀬ゆい - AVMOO</title>
//<meta name="author" content="AVMOO">
//<meta name="keywords" content="BLD-001,BLD001,永瀬ゆい,单体作品,贫乳・微乳,口交,美少女,娇小的,女上位,高画质,ドリームチケット,純潔ロリィタと性交">
//<meta name="description" content="BLD-001 純潔ロリィタと性交 永瀬ゆい - AVMOO - 你的线上日本成人影片情报站。管理你的影片并分享你的想法。">
//<link rel="apple-touch-icon" href="https://avmask.com/app/jav/View/img/apple-touch-icon.png">
//<link rel="shortcut Icon" href="https://avmask.com/app/jav/View/img/favicon.ico">
//<link rel="bookmark" href="https://avmask.com/app/jav/View/img/favicon.ico">
//<link rel="dns-prefetch" href="https://jp.netcdn.space" />
//<link rel="dns-prefetch" href="https://us.netcdn.space" />
//<link rel="dns-prefetch" href="https://ads.exoclick.com" />
//<link rel="dns-prefetch" href="https://syndication.exoclick.com" />
//<link rel="dns-prefetch" href="https://adserver.juicyads.com" />
//<link rel="dns-prefetch" href="https://j.traffichunt.com" />
//<script language="javascript">function $ROOT_URL(){return "https://avmask.com"}function $APP(){return "jav"}function $APP_URL(){return "/index.php?app=jav&"}function $APP_INFO_URL(){return "/index.php/jav"}function $APP_REWRITE_URL(){return "/jav"}function $APP_VIEW_URL(){return "https://avmask.com/app/jav/View"}function $APP_UPLOAD_URL(){return "/app/jav/Upload"}</script><link rel='stylesheet' type='text/css' href='https://avmask.com/app/jav/View/css/app.min.css?v=1476953808'>
//<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries --><!--[if lt IE 9]><script src='https://avmask.com/script/html5shiv/3.7.2/html5shiv.min.js?v=1476953808'></script>
//<script src='https://avmask.com/script/respond/1.4.2/respond.min.js?v=1476953808'></script>
//<![endif]--><!--[if lt IE 8]><link rel='stylesheet' type='text/css' href='https://avmask.com/script/bootstrap.ie7/1.0/bootstrap.ie7.min.css?v=1476953808' >
//<![endif]--><script>
//(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
//(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
//m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
//})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
//ga('create', 'UA-74041965-4', {'sampleRate': 50});
//ga('send', 'pageview');
//</script>
//</head>
//
//  <body>
//    <nav class="navbar navbar-default navbar-fixed-top top-bar">
//      <div class="container">
//        <div class="navbar-header">
//                <a href="https://avmask.com/cn" class="logo"></a>
//            <div class="btn-group pull-right visible-xs-inline" role="group" style="margin-top:8px;margin-right:8px;">
//                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-expanded="false">
//                  <span class="glyphicon glyphicon-globe"></span> 简体中文                  <span class="caret"></span>
//                </button>
//                <ul class="dropdown-menu" role="menu">
//                    <li><a href="https://avmask.com/en/movie/71efa47a8f27d447">English</a></li>
//                    <li><a href="https://avmask.com/ja/movie/71efa47a8f27d447">日本语</a></li>
//                    <li><a href="https://avmask.com/tw/movie/71efa47a8f27d447">正體中文</a></li>
//                    <li><a href="https://avmask.com/cn/movie/71efa47a8f27d447">简体中文</a></li>
//                </ul>
//            </div>
//        </div>
//
//        <div id="navbar" class="collapse navbar-collapse">
//          <form class="navbar-form navbar-left fullsearch-form" action="https://avmask.com/cn/search" onsubmit="return false">
//            <div class="input-group">
//              <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
//              <span class="input-group-btn">
//                <button class="btn btn-default" type="submit">搜寻</button>
//              </span>
//            </div>
//          </form>
//          <ul class="nav navbar-nav">
//            <li ><a href="https://avmask.com/cn">全部</a></li>
//                        <li ><a href="https://avmask.com/cn/released">已发布</a></li>
//                        <li ><a href="https://avmask.com/cn/popular">热门</a></li>
//            <li ><a href="https://avmask.com/cn/actresses">女优</a></li>
//            <li ><a href="https://avmask.com/cn/genre">类别</a></li>
//          </ul>
//          <ul class="nav navbar-nav navbar-right">
//            <li class="dropdown">
//              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false"><span class="glyphicon glyphicon-globe" style="font-size:12px;"></span> <span class="hidden-sm">简体中文</span> <span class="caret"></span></a>
//              <ul class="dropdown-menu" role="menu">
//                <li><a href="https://avmask.com/en/movie/71efa47a8f27d447">English</a></li>
//                <li><a href="https://avmask.com/ja/movie/71efa47a8f27d447">日本语</a></li>
//                <li><a href="https://avmask.com/tw/movie/71efa47a8f27d447">正體中文</a></li>
//                <li><a href="https://avmask.com/cn/movie/71efa47a8f27d447">简体中文</a></li>
//              </ul>
//            </li>
//          </ul>
//        </div><!--/.nav-collapse -->
//      </div>
//    </nav>
//<div class="row visible-xs-inline footer-bar">
//    <div class="col-xs-3 text-center">
//        <a id="menu" class="btn btn-default trigger-overlay"><i class="glyphicon glyphicon-align-justify"></i></a>
//    </div>
//        <div class="col-xs-6"></div>
//        <div class="col-xs-3 text-center">
//        <a id="back" class="btn btn-default" href="javascript:window.history.back()"><i class="glyphicon glyphicon-share-alt flipx"></i></a>
//    </div>
//</div>
//    <div class="container">
//        <h3>BLD-001 純潔ロリィタと性交 永瀬ゆい</h3>
//        <div class="row movie">
//            <div class="col-md-9 screencap">
//                <a class="bigImage" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001pl.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい"><img src="https://jp.netcdn.space/digital/video/24d00001pl.jpg"></a>
//            </div>
//            <div class="col-md-3 info">
//                <p><span class="header">识别码:</span> <span style="color:#CC0000;">BLD-001</span></p>
//                <p><span class="header">发行时间:</span> 2019-11-22</p>
//                                <p><span class="header">长度:</span> 119分钟</p>
//                                                                <p class="header">制作商: </p>
//                <p><a href="https://avmask.com/cn/studio/0a67a6d1aaa31d40">ドリームチケット</a></p>
//                                                                <p class="header">系列:</p>
//                <p><a href="https://avmask.com/cn/series/8866d9866b3673d3">純潔ロリィタと性交</a></p>
//                                                                <p class="header">类别:</p>
//                                <p><span class="genre"><a href="https://avmask.com/cn/genre/c4145926405d550f">单体作品</a></span><span class="genre"><a href="https://avmask.com/cn/genre/eaa8cebf3db8b605"微乳</a></span><span class="genre"><a href="https://avmask.com/cn/genre/428875c0f04324d6">口交</a></span><span class="genre"><a href="https://avmask.com/cn/genre/b0eaad139052cec8">美少女</a></span><span genre"><a href="https://avmask.com/cn/genre/00f15d8f74f7fb95">娇小的</a></span><span class="genre"><a href="https://avmask.com/cn/genre/8cb503f3126ac8dd">女上位</a></span><span class="genre"><a href="httvmask.com/cn/genre/5f9f62d40baa77cf">高画质</a></span></p>
//                            </div>
//        </div>
//        <div id="movie-share" class="row" style="margin:20px 0px 20px 0px;text-align:center;">
//            <a class="share-facebook"><span class="icon"></span></a>
//            <a class="share-twitter"><span class="icon"></span></a>
//            <a class="share-tumblr hidden-xxs"><span class="icon"></span></a>
//            <a class="share-reddit"><span class="icon"></span></a>
//            <a class="share-pinterest"><span class="icon"></span></a>
//            <a class="share-google_plusone_share hidden-xxs"><span class="icon"></span></a>
//            <a class="share-blogger hidden-xxs"><span class="icon"></span></a>
//            <a class="share-google hidden-xxs"><span class="icon"></span></a>
//            <a class="share-qrcode"><span class="icon"></span></a>
//            <a class="share-email"><span class="icon"></span></a>
//            <a class="share-favorites"><span class="icon"></span></a>
//        </div>
//
//        <h4>推荐</h4>
//        <div class="row hidden-xs ptb-10 text-center">
//        <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTV83Mjh4OTAiLCJhZHNwb3QiOiJqX01fNzI4eDkwIiwid2VpZ2h0IjoiMiIsImZjYXAiOmZhbHNlLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOiI3NjgiLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOmZhbHNlLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiPlxyXG52YXIgYWRfaWR6b25lID0gXCI4MTMzMDhcIixcclxuXHQgYWRfd2lkdGggPSBcIjcyOFwiLFxyXG5cdCBhZF9oZWlnaHQgPSBcIjkwXCI7XHJcbjxcL3NjcmlwdD5cclxuPHNjcmlwdCB0eXBlPVwidGV4dFwvamF2YXNjcmlwdFwiIHNyYz1cImh0dHBzOlwvXC9hZHMuZXhvY2xpY2suY29tXC9hZHMuanNcIj48XC9zY3JpcHQ+XHJcbjxub3NjcmlwdD48YSBocmVmPVwiaHR0cDpcL1wvbWFpbi5leG9jbGljay5jb21cL2ltZy1jbGljay5waHA/aWR6b25lPTgxMzMwOFwiIHRhcmdldD1cIl9ibGFua1wiPjxpbWcgc3JjPVwiaHR0cHM6XC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9ODEzMzA4Jm91dHB1dD1pbWcmdHlwZT03Mjh4OTBcIiB3aWR0aD1cIjcyOFwiIGhlaWdodD1cIjkwXCI+PFwvYT48XC9ub3NjcmlwdD4ifSx7ImlkIjoianVpY19qX01fNzI4eDkwIiwiYWRzcG90Ijoial9NXzcyOHg5MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTczNiBoZWlnaHQ9OTggc2Nyb2xsaW5nPW5vIGFsbG93dHJhbnNwYXJlbmN5PXRydWUgc3JjPVwvXC9hZHNlcnZlci5qdWljeWFkcy5jb21cL2Fkc2hvdy5waHA/YWR6b25lPTM3MTcwOD48XC9pZnJhbWU+In1d</div>        </div>
//
//        <div class="visible-xs-block pt-10 text-center">
//        <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTV8zMDB4MTAwIiwiYWRzcG90Ijoial9NXzMwMHgxMDAiLCJ3ZWlnaHQiOiIzIiwiZmNhcCI6ZmFsc2UsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOiI3NjgiLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0PlxyXG5hZF9pZHpvbmUgPSBcIjEwMjYzMTBcIjtcclxuYWRfd2lkdGggPSBcIjMwMFwiO1xyXG5hZF9oZWlnaHQgPSBcIjEwMFwiO1xyXG5pZih0b3A9PT1zZWxmKSB2YXIgcD1kb2N1bWVudC5VUkw7IGVsc2UgdmFyIHA9ZG9jdW1lbnQucmVmZXJyZXI7dmFyIGR0PW5ldyBEYXRlKCkuZ2V0VGltZSgpO1xyXG52YXIgZXhvRG9jdW1lbnRQcm90b2NvbCA9IChkb2N1bWVudC5sb2NhdGlvbi5wcm90b2NvbCAhPSBcImh0dHBzOlwiICYmIGRvY3VtZW50LmxvY2F0aW9uLnByb3RvY29sICE9IFwiaHR0cDpcIikgPyBcImh0dHBzOlwiIDogZG9jdW1lbnQubG9jYXRpb24ucHJvdG9jb2w7XHJcbmlmKHR5cGVvZihhZF9zdWIpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfc3ViID0gXCJcIjtcclxuaWYodHlwZW9mKGFkX3RhZ3MpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfdGFncyA9IFwiXCI7XHJcbnZhciBhZF90eXBlID0gYWRfd2lkdGggKyAneCcgKyBhZF9oZWlnaHQ7XHJcbmlmKGFkX3dpZHRoID09ICcxMDAlJyAmJiBhZF9oZWlnaHQgPT0gJzEwMCUnKSBhZF90eXBlID0gJ2F1dG8nO1xyXG52YXIgYWRfc2NyZWVuX3Jlc29sdXRpb24gPSBzY3JlZW4ud2lkdGggKyAneCcgKyBzY3JlZW4uaGVpZ2h0O1xyXG5kb2N1bWVudC53cml0ZSgnPGlmcmFtZSBmcmFtZWJvcmRlcj1cIjBcIiBzY3JvbGxpbmc9XCJub1wiIHdpZHRoPVwiJyArIGFkX3dpZHRoICsgJ1wiIGhlaWdodD1cIicgKyBhZF9oZWlnaHQgKyAnXCIgc3JjPVwiJyArIGV4b0RvY3VtZW50UHJvdG9jb2wgKyAnXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9JyArIGFkX2lkem9uZSArICcmdHlwZT0nICsgYWRfdHlwZSArICcmcD0nICsgZXNjYXBlKHApICsgJyZkdD0nICsgZHQgKyAnJnN1Yj0nICsgYWRfc3ViICsgJyZ0YWdzPScgKyBhZF90YWdzICsgJyZzY3JlZW5fcmVzb2x1dGlvbj0nICsgYWRfc2NyZWVuX3Jlc29sdXRpb24gKyAnXCI+PFwvaWZyYW1lPicpO1xyXG48XC9zY3JpcHQ+In0seyJpZCI6Imp1aWNfal9NXzMwMHgxMDAiLCJhZHNwb3QiOiJqX01fMzAweDEwMCIsIndlaWdodCI6IjUiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTMwMCBoZWlnaHQ9MTAwIHNjcm9sbGluZz1ubyBhbGxvd3RyYW5zcGFyZW5jeT10cnVlIHNyYz1cL1wvYWRzZXJ2ZXIuanVpY3lhZHMuY29tXC9hZHNob3cucGhwP2Fkem9uZT00NTE3Mzk+PFwvaWZyYW1lPiJ9XQ==</div>        </div>
//
//                <h4>演员</h4>
//        <div id="avatar-waterfall">
//                        <a class="avatar-box" href="https://avmask.com/cn/star/63a1765e84c06799">
//                                <div class="photo-frame">
//                    <img src="https://jp.netcdn.space/mono/actjpgs/nowprinting.gif" title="">
//                </div>
//                                <span>永瀬ゆい</span>
//                        </a>
//                    </div>
//
//                <div class="clearfix"></div>
//                <h4>样品图像</h4>
//        <div id="sample-waterfall">
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-1.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 1">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-1.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-2.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 2">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-2.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-3.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 3">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-3.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-4.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 4">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-4.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-5.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 5">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-5.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-6.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 6">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-6.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-7.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 7">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-7.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-8.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 8">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-8.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-9.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 9">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-9.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-10.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 10">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-10.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-11.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 11">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-11.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-12.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 12">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-12.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-13.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 13">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-13.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-14.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 14">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-14.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-15.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 15">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-15.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-16.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 16">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-16.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-17.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 17">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-17.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-18.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 18">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-18.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-19.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 19">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-19.jpg">
//                </div>
//                        </a>
//                                <a class="sample-box" href="https://jp.netcdn.space/digital/video/24bld00001/24bld00001jp-20.jpg" title="BLD-001 純潔ロリィタと性交 永瀬ゆい - 样品图像 - 20">
//                <div class="photo-frame">
//                                <img src="https://jp.netcdn.space/digital/video/24bld00001/24bld00001-20.jpg">
//                </div>
//                        </a>
//                    </div>
//                <div class="clearfix"></div>
//        <div class="visible-xs-block">
//            <h4>推荐</h4>
//            <div class="row ptb-10 text-center">
//                <div type="text/data-position" style="display:none">W3siaWQiOiJleG9jX2pfTV8zMDB4MjUwIiwiYWRzcG90Ijoial9NXzMwMHgyNTAiLCJ3ZWlnaHQiOiIyIiwiZmNhcCI6ZmFsc2UsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOiI3NjgiLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0PlxyXG5hZF9pZHpvbmUgPSBcIjEwMzEwNDJcIjtcclxuYWRfd2lkdGggPSBcIjMwMFwiO1xyXG5hZF9oZWlnaHQgPSBcIjI1MFwiO1xyXG5pZih0b3A9PT1zZWxmKSB2YXIgcD1kb2N1bWVudC5VUkw7IGVsc2UgdmFyIHA9ZG9jdW1lbnQucmVmZXJyZXI7dmFyIGR0PW5ldyBEYXRlKCkuZ2V0VGltZSgpO1xyXG52YXIgZXhvRG9jdW1lbnRQcm90b2NvbCA9IChkb2N1bWVudC5sb2NhdGlvbi5wcm90b2NvbCAhPSBcImh0dHBzOlwiICYmIGRvY3VtZW50LmxvY2F0aW9uLnByb3RvY29sICE9IFwiaHR0cDpcIikgPyBcImh0dHBzOlwiIDogZG9jdW1lbnQubG9jYXRpb24ucHJvdG9jb2w7XHJcbmlmKHR5cGVvZihhZF9zdWIpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfc3ViID0gXCJcIjtcclxuaWYodHlwZW9mKGFkX3RhZ3MpID09ICd1bmRlZmluZWQnKSB2YXIgYWRfdGFncyA9IFwiXCI7XHJcbnZhciBhZF90eXBlID0gYWRfd2lkdGggKyAneCcgKyBhZF9oZWlnaHQ7XHJcbmlmKGFkX3dpZHRoID09ICcxMDAlJyAmJiBhZF9oZWlnaHQgPT0gJzEwMCUnKSBhZF90eXBlID0gJ2F1dG8nO1xyXG52YXIgYWRfc2NyZWVuX3Jlc29sdXRpb24gPSBzY3JlZW4ud2lkdGggKyAneCcgKyBzY3JlZW4uaGVpZ2h0O1xyXG5kb2N1bWVudC53cml0ZSgnPGlmcmFtZSBmcmFtZWJvcmRlcj1cIjBcIiBzY3JvbGxpbmc9XCJub1wiIHdpZHRoPVwiJyArIGFkX3dpZHRoICsgJ1wiIGhlaWdodD1cIicgKyBhZF9oZWlnaHQgKyAnXCIgc3JjPVwiJyArIGV4b0RvY3VtZW50UHJvdG9jb2wgKyAnXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvYWRzLWlmcmFtZS1kaXNwbGF5LnBocD9pZHpvbmU9JyArIGFkX2lkem9uZSArICcmdHlwZT0nICsgYWRfdHlwZSArICcmcD0nICsgZXNjYXBlKHApICsgJyZkdD0nICsgZHQgKyAnJnN1Yj0nICsgYWRfc3ViICsgJyZ0YWdzPScgKyBhZF90YWdzICsgJyZzY3JlZW5fcmVzb2x1dGlvbj0nICsgYWRfc2NyZWVuX3Jlc29sdXRpb24gKyAnXCI+PFwvaWZyYW1lPicpO1xyXG48XC9zY3JpcHQ+In0seyJpZCI6Imp1aWNfal9NXzMwMHgyNTAiLCJhZHNwb3QiOiJqX01fMzAweDI1MCIsIndlaWdodCI6IjMiLCJmY2FwIjpmYWxzZSwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxpZnJhbWUgYm9yZGVyPTAgZnJhbWVib3JkZXI9MCBtYXJnaW5oZWlnaHQ9MCBtYXJnaW53aWR0aD0wIHdpZHRoPTMwOCBoZWlnaHQ9MjU4IHNjcm9sbGluZz1ubyBhbGxvd3RyYW5zcGFyZW5jeT10cnVlIHNyYz1cL1wvYWRzZXJ2ZXIuanVpY3lhZHMuY29tXC9hZHNob3cucGhwP2Fkem9uZT0zNzE3MjY+PFwvaWZyYW1lPiJ9XQ==</div>            </div>
//        </div>
//                <div class="row visible-xs-block">
//            <div class="col-xs-6 text-center">
//                <a href="https://btos.pw/search/BLD-001" class="btn btn-lg btn-primary" target="_blank"><span class="glyphicon glyphicon-save"></span> 下载</a>
//            </div>
//            <div class="col-xs-6 text-center">
//                <a href="https://btos.pw/video/BLD-001%E7%!B(MISSING)4%E6%!B(MISSING)D%E3%AD%!E(MISSING)3%AA%!E(MISSING)3%A3%!E(MISSING)3%BF%!E(MISSING)3%A8%!E(MISSING)6%A7%!E(MISSING)4%!B(MISSING)A%!A(MISSING)4%E6%!B(MISSING)0%!B(MISSING)8%!E(MISSING)7%AC%!E(MISSING)3%86%!E(MISSING)3%84" class="btn btn-lg btn-warning" target="_blank"><span class="glyphicon glyphicon-play"></span> 播放</a>
//            </div>
//        </div>
//        <div class="hidden-xs">
//            <h4>下载</h4>
//            <div class="row ptb-10 text-center">
//                <a href="https://btos.pw/search/BLD-001" target="_blank"><img src="https://avmask.com/app/jav/View/img/download_zh.png"></a>
//            </div>
//        </div>
//        <div class="hidden-xs">
//            <h4>推荐</h4>
//            <div class="row ptb-20 text-center">
//                <div type="text/data-position" style="display:none">W3siaWQiOiJqYXZ1X2pfUF83Mjh4OTAiLCJhZHNwb3QiOiJqX1BfNzI4eDkwIiwid2VpZ2h0IjoiMSIsImZjYXAiOmZhbHNlLCJzY2hlZHVsZSI6ZmFsc2UsIm1heFdpZHRoIjpmYWxzZSwibWluV2lkdGgiOmZhbHNlLCJ0aW1lem9uZSI6ZmFsc2UsImV4Y2x1ZGUiOmZhbHNlLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPGEgaHJlZj1cImh0dHBzOlwvXC90ZWxsbWUucHdcL2dvXC9qYXZ1XCIgdGFyZ2V0PVwiX2JsYW5rXCI+PGltZyBzcmM9XCJcL2FwcFwvamF2dVwvVmlld1wvaW1nXC9iNzI4OTAuanBnXCIgd2lkdGg9XCI3MjhcIiBoZWlnaHQ9XCI5MFwiIGJvcmRlcj1cIjBcIj48XC9hPiJ9LHsiaWQiOiJ3YXZfal9QXzcyOHg5MCIsImFkc3BvdCI6ImpfUF83Mjh4OTAiLCJ3ZWlnaHQiOiIzIiwiZmNhcCI6ZmFsc2UsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8YSBocmVmPVwiaHR0cHM6XC9cL2F2eG8uY2x1YlwvXCIgdGFyZ2V0PVwiX2JsYW5rXCI+PGltZyBzcmM9XCJcL2FwcFwvd2F2XC9WaWV3XC9pbWdcL2I3Mjg5MC5qcGdcIiB3aWR0aD1cIjcyOFwiIGhlaWdodD1cIjkwXCIgYm9yZGVyPVwiMFwiPjxcL2E+In1d</div>              </div>
//        </div>
//    </div>
//<div type="text/data-position" style="display:none">W3siaWQiOiJhZHN0X2pfUE9QVU5ERVIiLCJhZHNwb3QiOiJqX1BPUFVOREVSIiwid2VpZ2h0IjoiNSIsImZjYXAiOiIyIiwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6ZmFsc2UsIm1pbldpZHRoIjoiNzY4IiwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxzY3JpcHQgdHlwZT0ndGV4dFwvamF2YXNjcmlwdCcgc3JjPSdodHRwczpcL1wvcGwxMTA1MjUucHVodG1sLmNvbVwvNzBcLzgyXC85Y1wvNzA4MjljMzgyMTZlMWMwNDYxNmFkYjQ2NzJlZGEzNDIuanMnPjxcL3NjcmlwdD4ifSx7ImlkIjoiY2xpY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjciLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6ZmFsc2UsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IGRhdGEtY2Zhc3luYz1cImZhbHNlXCIgdHlwZT1cInRleHRcL2phdmFzY3JpcHRcIiBzcmM9XCJcL1wvY2xjYXNzZC5jb21cL3RcLzlcL2ZyZXRcL21lb3c0XC8zNjkwODlcL2JydC5qc1wiPjxcL3NjcmlwdD4ifSx7ImlkIjoiZXhvY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjYiLCJmY2FwIjoiMiIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6Ijc2OCIsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6ZmFsc2UsImRvbWFpbiI6ZmFsc2UsImNvZGUiOiI8c2NyaXB0IHNyYz1cIlwvXC9zeW5kaWNhdGlvbi5leG9jbGljay5jb21cL3NwbGFzaC5waHA/aWR6b25lPTEwMDgwOTQmdHlwZT0zXCI+PFwvc2NyaXB0PiJ9LHsiaWQiOiJleG9tX2pfUE9QVU5ERVIiLCJhZHNwb3QiOiJqX1BPUFVOREVSIiwid2VpZ2h0IjoiNiIsImZjYXAiOiIyIiwic2NoZWR1bGUiOmZhbHNlLCJtYXhXaWR0aCI6Ijc2OCIsIm1pbldpZHRoIjpmYWxzZSwidGltZXpvbmUiOmZhbHNlLCJleGNsdWRlIjpmYWxzZSwiZG9tYWluIjpmYWxzZSwiY29kZSI6IjxzY3JpcHQgc3JjPVwiXC9cL3N5bmRpY2F0aW9uLmV4b2NsaWNrLmNvbVwvc3BsYXNoLnBocD9pZHpvbmU9MTAyNjMxMiZ0eXBlPTExXCI+PFwvc2NyaXB0PlxyXG48c2NyaXB0PlxyXG4kKGRvY3VtZW50KS5yZWFkeShmdW5jdGlvbigpIHtcclxuICAgIGlmICh0eXBlb2YgZXhvVXJsICE9IFwidW5kZWZpbmVkXCIpIHtcclxuICAgICAgICBleG9VcmwgPSBleG9VcmwucmVwbGFjZSgnaHR0cDpcL1wvJywgJ2h0dHBzOlwvXC8nKTtcclxuICAgIH1cclxuICAgICQoXCJhXCIpLmJpbmQoXCJjbGlja1wiLCBmdW5jdGlvbihldmVudCkge1xyXG4gICAgICAgIGV4b01vYmlsZVBvcCgpO1xyXG4gICAgICAgICQodGhpcykudW5iaW5kKFwiY2xpY2tcIik7XHJcbiAgICB9KTtcclxufSk7XHJcbjxcL3NjcmlwdD4ifSx7ImlkIjoianVpY19qX1BPUFVOREVSIiwiYWRzcG90Ijoial9QT1BVTkRFUiIsIndlaWdodCI6IjIiLCJmY2FwIjoiMSIsInNjaGVkdWxlIjpmYWxzZSwibWF4V2lkdGgiOmZhbHNlLCJtaW5XaWR0aCI6Ijc2OCIsInRpbWV6b25lIjpmYWxzZSwiZXhjbHVkZSI6WyJ6aC1jbiJdLCJkb21haW4iOmZhbHNlLCJjb2RlIjoiPCEtLSBKdWljeUFkcyBQb3BVbmRlcnMgdjMgU3RhcnQgLS0+XHJcbjxzY3JpcHQgdHlwZT1cInRleHRcL2phdmFzY3JpcHRcIiBzcmM9XCJodHRwczpcL1wvanMuanVpY3lhZHMuY29tXC9qcC5waHA/Yz03NDU0MzN6MnQyNTZyMnEydzI4NDYzODQmdT1odHRwcyUzQSUyRiUyRmJ0c28ucHdcIj48XC9zY3JpcHQ+XHJcbjwhLS0gSnVpY3lBZHMgUG9wVW5kZXJzIHYzIEVuZCAtLT4ifV0=</div>
//<!-- Modal -->
//<div class="modal fade" id="advertisingModal" tabindex="-1" role="dialog" aria-labelledby="advertisingModalLabel" aria-hidden="true">
//  <div class="modal-dialog">
//    <div class="modal-content">
//      <div class="modal-header">
//        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
//        <h4 class="modal-title" id="advertisingModalLabel">Advertising</h4>
//      </div>
//      <div class="modal-body">
//            <p>Please contact following agents for advertising on AVMOO</p>
//            <p><a href="https://www.exoclick.com/?login=james666" target="_blank" style="color: #D80456;">ExoClick</a> / <a href="https://manage.juicyads.com/juicysites.php?id=128293" target="_blank" style="color: #D80456;">JuicyAds</a> / <a href="http://www.clickadu.com/?rfd=0l1" target="_blank" style="color: #D80456;">ClickADu</a></p>
//      </div>
//      <div class="modal-footer">
//        <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
//      </div>
//    </div>
//  </div>
//</div>
//
//<footer class="footer hidden-xs">
//        <div class="container">
//<p><a href="https://avmask.com/cn/terms">Terms</a> / <a href="https://avmask.com/cn/privacy">Privacy</a> / <a href="https://avmask.com/cn/usc">2257</a> / <a href="http://www.rtalabel.org/" target="_blank" rel="external nofollow">RTA</a> / <a href="#advertisingModal" role="button" data-toggle="modal">Advertising</a> / <a class="contactus" href="javascript:;" role="button" data-toggle="modal">Contact</a> / <a href="https://tellme.pw/avmoo" target="_blank">Guide</a> | Links: <a href="https://avmask.com" target="_blank">AVMOO</a> / <a href="https://avsox.asia" target="_blank">AVSOX</a> / <a href="https://avmemo.asia" target="_blank">AVMEMO</a><br>Copyright © 2013 AVMOO. All Rights Reserved. All other trademarks and copyrights are the property of their respective holders. The reviews and comments expressed at or through this website are the opinions of the individual author and do not reflect the opinions or views of AVMOO. AVMOO is not responsible for the accuracy of any of the information supplied here.</p>
//        </div>
//</footer>
//
//<div class="visible-xs-block footer-bar-placeholder"></div>
//
//<!-- ////////////////////////////////////////////////// -->
//<div class="overlay overlay-contentscale">
//    <nav>
//        <ul>
//            <li>
//            <form class="fullsearch-form" action="https://avmask.com/cn/search" onsubmit="return false">
//               <div class="input-group col-xs-offset-2 col-xs-8">
//                  <input name="keyword" type="text" class="form-control" placeholder="搜寻 识别码, 影片, 演员">
//                  <span class="input-group-btn">
//                    <button class="btn btn-default" type="submit">搜寻</button>
//                  </span>
//               </div>
//            </form>
//            </li>
//            <li><a href="https://avmask.com/cn/released">已发布</a></li>
//            <li><a href="https://avmask.com/cn/popular">热门</a></li>
//            <li><a href="https://avmask.com/cn/actresses">女优</a></li>
//            <li><a href="https://avmask.com/cn/genre">类别</a></li>
//        </ul>
//    </nav>
//    <div class="row overlay-close"><i class="glyphicon glyphicon-remove" style="color:#fff;font-size: 24px;margin:30px;"></i></div>
//</div>
//<script src='https://avmask.com/app/jav/View/js/app.min.js?v=1476953808'></script>  </body>
//</html>
//Process finished with exit code 0
