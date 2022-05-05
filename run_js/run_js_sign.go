package run_js

import (
	"io/ioutil"
	"log"
)
import 	"github.com/robertkrimen/otto"
import   "github.com/dop251/goja"

func GetSignV2(search string ) (sign string,err error ) {

	var  script = `
    function fib(n) {
        if (n === 1 || n === 2) {
            return 1 
        }
        return fib(n - 1) + fib(n - 2)
    }
    `
	vm := goja.New()
	filePath := "./run_js/sign.js"
	//filePath := "./sign.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return

	}

	script = string(bytes)
	_, err = vm.RunString(script)
	if err != nil {
		log.Println("JS代码有问题！")
		return
	}
	var fn func(string ) string
	err = vm.ExportTo(vm.Get("e"), &fn)
	if err != nil {
		log.Println("Js函数映射到 Go 函数失败！")
		return
	}
	sign =  fn(search)
	return
}

func GetSign(search string ) (sign string,err error )  {
	//filePath := "./run_js/sign.js"
	filePath := "./sign.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return

	}
	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err!=nil {
		log.Println(err,222222)
		return
	}
	//encodeInp是JS函数的函数名
	//var method, _ = vm.Get("getSend")
	//s, _ := vm.Compile("sign.js", `getSend("zzz")`)
	//value, err := vm.Run(s)
	//if err != nil {
	//	log.Println(err, 11111)
	//
	//	return
	//}
	value, err := vm.Call("e",nil, search)
	if err != nil {
		log.Println(err,11111)

		return

	}
	sign =   value.String()
	log.Println(sign)
	return
}