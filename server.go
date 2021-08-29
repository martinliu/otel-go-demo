//一个普通的 Hello World
//流量测试命令如下：
// hey -z 10m -q 10 -cpus 1 http://localhost:8000/hello/world

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
