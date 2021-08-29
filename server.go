//一个普通的 Hello World
//流量测试命令如下：
// hey -z 10m -q 10 -cpus 1 http://localhost:8000/hello/world

//第一步：使用 Elastic APM 的 Go 语言 Agent 埋点
// git checkout step1
// hey -z 10m -q 10 -cpus 1 http://localhost:8000/hello/world
// 查看 APM 的监控界面
//文档： https://pkg.go.dev/go.elastic.co/apm/module/apmgorilla

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	//第一步
	"go.elastic.co/apm/module/apmgorilla"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {
	r := mux.NewRouter()
	r.Use(apmgorilla.Middleware())
	r.HandleFunc("/hello/{name}", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
