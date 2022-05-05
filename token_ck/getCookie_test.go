package token_ck

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestGetCookie(t *testing.T)  {
	//新建一个请求和http客户端
	GetCookie()
}


func TestGetCookieV2(t *testing.T)  {
	//新建一个请求和http客户端
	c,token := GetCookieTokenV2()
	t.Log(c,token)
}

func Test_regexp_Match(t *testing.T)  {
	filePath := "./test2.html"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return

	}
	html := string(bytes)
	mustReg := regexp.MustCompile(`token: '(.+)',`)
	arr := mustReg.FindStringSubmatch(html)
	t.Log(arr)
}