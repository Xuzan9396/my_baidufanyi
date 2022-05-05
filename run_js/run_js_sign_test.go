package run_js

import "testing"

func Test_sign(t *testing.T)  {
	sign,err := GetSignV2("徐赞") // 419476.165285
	t.Log(sign,err)

	return
	sign,err = GetSign("123")
	t.Log(sign,err)
}
