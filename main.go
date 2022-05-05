package main

import (
	"bytes"
	"github.com/Albert-Zhan/httpc"
	"github.com/Albert-Zhan/httpc/body"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"mime/multipart"
	"my_baidufanyi/common"
	"my_baidufanyi/run_js"
	"my_baidufanyi/token_ck"
	"net/http"
	"net/url"
	"time"
)

func init()  {
	//新建http客户端
	common.Client =httpc.NewHttpClient()
	//新建一个cookie管理器,后面所有请求的cookie将保存在这
	common.CookieJar =httpc.NewCookieJar()
	//设置cookie管理器,
	common.Client.SetCookieJar(common.CookieJar)
}

func main()  {
	token := token_ck.GetCookieTokenV2()
	req := httpc.NewRequest(common.Client)
	urls := "https://fanyi.baidu.com/v2transapi?from=zh&to=en"
	req.SetMethod("post").SetUrl(urls)

	searchNameList := []string{"工作","劳动节日","活动"}
	for _, searchName := range searchNameList {
		// 翻译搜索词
		sign,err := run_js.GetSignV2(searchName)
		if err != nil {
			log.Println(err)
			return
		}
		b:=body.NewUrlEncode()
		b.SetData("to", "en")
		b.SetData("query", searchName)
		b.SetData("transtype", "translang")
		b.SetData("simple_means_flag", "3")
		b.SetData("sign", sign)
		//log.Println("sign:",sign)
		b.SetData("token", token)
		b.SetData("domain", "common")
		_,bodys,err := req.SetBody(b).Send().End()
		if err != nil {
			log.Println(err)
			return
		}
		dst := gjson.Get(bodys,"trans_result.data.0.dst").String()
		log.Println("翻译前:",searchName,"----->翻译后:",dst)
		time.Sleep(5*time.Second)

	}


}




// 原生请求
func main2() {

	token := token_ck.GetCookieTokenV2()

	urls := "https://fanyi.baidu.com/v2transapi?from=zh&to=en"

	u, _ := url.Parse("https://fanyi.baidu.com")
	cookies:=common.CookieJar.Cookies(u)
	var cookiess []*http.Cookie
	var cookiestr string

	for _, cookie := range cookies {
		if cookie.Name == "BAIDUID"{
			cookiestr = cookie.Name + "=" + cookie.Value
			log.Println(cookiestr)
			cookiei:=&http.Cookie{Name:cookie.Name,Value:cookie.Value}
			cookiess= append(cookiess, cookiei)
			//return
		}
	}
	sign,err := run_js.GetSignV2("工作")
	if err != nil {
		log.Println(err)
		return
	}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("to", "en")
	_ = writer.WriteField("query", "工作")
	_ = writer.WriteField("transtype", "translang")
	_ = writer.WriteField("simple_means_flag", "3")
	_ = writer.WriteField("sign", sign)
	_ = writer.WriteField("token", token)
	_ = writer.WriteField("domain", "common")
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client {
	}
	req, err := http.NewRequest("POST", urls, payload)

	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Cookie", cookiestr)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//req.Header.Set("Content-Type", "multipart/form-data")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	dst := gjson.Get(string(body),"trans_result.data.0.dst").String()
	log.Println("翻以前:","工作","----->翻译后:",dst)
}