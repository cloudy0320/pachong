package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/robertkrimen/otto"
	"strings"
)

func main() {
	vm := otto.New()
	ss := `     eval(function (p, a, c, k, e, d) {
        e = function (c) {
            return (c < a ? '' : e(parseInt(c / a))) + ((c = c % a) > 35 ? String.fromCharCode(c + 29) : c.toString(36))
        };
        if (!''.replace(/^/, String)) {
            while (c--) {
                d[e(c)] = k[c] || e(c)
            }
            k = [function (e) {
                return d[e]
            }];
            e = function () {
                return '\\w+'
            };
            c = 1
        }
        ;
        while (c--) {
            if (k[c]) {
                p = p.replace(new RegExp('\\b' + e(c) + '\\b', 'g'), k[c])
            }
        }
        return p
    }('p m=\'{"l":"j","q":"0","k":"i\\/4\\/3\\/2\\/n.5\\r\\6\\/4\\/3\\/2\\/h.5\\r\\6\\/4\\/3\\/2\\/g.5\\r\\6\\/4\\/3\\/2\\/a.5\\r\\6\\/4\\/3\\/2\\/9.5\\r\\6\\/4\\/3\\/2\\/8.5\\r\\6\\/4\\/3\\/2\\/7.5\\r\\6\\/4\\/3\\/2\\/b.5\\r\\6\\/4\\/3\\/2\\/c.5\\r\\6\\/4\\/3\\/2\\/f.5\\r\\6\\/4\\/3\\/2\\/e.5\\r\\6\\/4\\/3\\/2\\/d.5\\r\\6\\/4\\/3\\/2\\/o.5\\r\\6\\/4\\/3\\/2\\/v.5\\r\\6\\/4\\/3\\/2\\/J.5\\r\\6\\/4\\/3\\/2\\/s.5\\r\\6\\/4\\/3\\/2\\/F.5\\r\\6\\/4\\/3\\/2\\/E.5\\r\\6\\/4\\/3\\/2\\/G.5\\r\\6\\/4\\/3\\/2\\/H.5\\r\\6\\/4\\/3\\/2\\/I.5\\r\\6\\/4\\/3\\/2\\/D.5\\r\\6\\/4\\/3\\/2\\/B.5\\r\\6\\/4\\/3\\/2\\/C.5\\r\\6\\/4\\/3\\/2\\/u.5\\r\\6\\/4\\/3\\/2\\/t.5","w":"x","A":"1","z":"\\y\\K"}\';', 47, 47, '||29378|9897|chapterpic|jpg|nimg|14527578267048|14527578264801|1452757826227|14527578258772|14527578269175|14527578271228|14527578277821|14527578275448|1452757827326|14527578256189|14527578254117|img|45120|page_url|id|pages|14527578247495|14527578280205|var|hidden||14527578287565|14527578344148|14527578337993|14527578282921|sum_pages|26|u5e8f|chapter_name|chapter_order|14527578329076|14527578332895|14527578321738|14527578293242|14527578289773|14527578296316|14527578308707|14527578310716|14527578284941|u7ae0'.split('|'), 0, {}))`
	ss = "var url\n" + ss
	index := strings.Index(ss, "return p")
	ss = ss[:index] + "url=p\n" + ss[index:]
	vm.Run(ss)
	value, _ := vm.Get("url")
	s := value.String()[strings.Index(value.String(), "{") : strings.Index(value.String(), "}")+1]
	js, err := simplejson.NewJson([]byte(s))
	if err != nil {
		fmt.Println(err)
	}
	str, _ := js.Get("page_url").String()
	arr := strings.Split(str, "\r\n")
	for _, v := range arr {
		fmt.Println(v)
	}
}