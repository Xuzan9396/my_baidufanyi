package token_ck

import (
	"github.com/Albert-Zhan/httpc"
	"log"
	"my_baidufanyi/common"
	"os"
	"regexp"
)

//func GetCookie() string {
//	//新建一个请求和http客户端
//	req:=httpc.NewRequest(httpc.NewHttpClient())
//	//get请求,返回string类型的body
//	resp,body,err:=req.SetUrl("https://fanyi.baidu.com/").Send().End()
//	if err!=nil {
//		log.Println(err)
//	}else{
//		//log.Println(resp)
//		//log.Println(body)
//		_ = body
//	}
//	val := resp.Header.Get("Set-Cookie")
//	arr := strings.Split(val," ")
//	for _, s := range arr {
//		if strings.Contains(s,"BAIDUID="){
//			log.Println(s)
//			return s
//		}
//	}
//	return ""
//}

func GetCookieTokenV2() (tokens string ) {
	//新建http客户端
	//client:=httpc.NewHttpClient()
	////新建一个cookie管理器,后面所有请求的cookie将保存在这
	//cookieJar:=httpc.NewCookieJar()
	////设置cookie管理器,
	//client.SetCookieJar(cookieJar)
	//新建一个请求
	req:=httpc.NewRequest(common.Client)
	//req.SetMethod("post").SetUrl("http://127.0.0.1")
	//设置头信息,返回byte类型的body
	_,_,err:=req.SetUrl("https://fanyi.baidu.com/").Send().End()
	//WriteContent("test1.html",body)

	if err!=nil {
		log.Println(err)
		return
	}else{
		//从cookie管理器中获取当前访问url保存的cookie
		//u, _ := url.Parse("https://fanyi.baidu.com")
		//cookies:=common.CookieJar.Cookies(u)
		//for _, cookie := range cookies {
		//	if cookie.Name == "BAIDUID"{
		//		cookiestr = cookie.Name + "=" + cookie.Value
		//	}
		//}
		//log.Println(cookiestr) // [BAIDUID=0F068BB05FFC3B40E42A243263195B2E:FG=1; Domain=baidu.com]
		_,body,err:=req.SetUrl("https://fanyi.baidu.com/").Send().End()
		if err != nil {
			log.Println(err)
			return
		}
		mustReg := regexp.MustCompile(`token: '(.+)',`)
		arr := mustReg.FindStringSubmatch(body)
		log.Println(arr)
		if len(arr) >= 2 {
			return arr[1]
		}
		//WriteContent("test2.html",body)

	}
	return
}


func WriteContent(fileName, content string) {

	if !isExist(fileName) {
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			// 创建文件失败处理

		} else {

			_, err = f.Write([]byte(content))
			if err != nil {
				// 写入失败处理

			}
		}

	} else {
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0755)
		defer f.Close()
		if err != nil {
			return
		} else {
			_, err = f.Write([]byte(content))
		}
	}

}

func isExist(path string) bool {
	//递归创建文件夹

	//err := os.MkdirAll("./test/1/2", os.ModePerm)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsExist(err) {

			return true

		}

		if os.IsNotExist(err) {

			return false

		}


		return false

	}

	return true

}