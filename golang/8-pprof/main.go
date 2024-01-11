/*
 * @Author:       Kit-Hung
 * @Date:         2024/1/11 0:31
 * @Description： 性能测试相关示例
 */
package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	runtimePporf "runtime/pprof"
)

func main() {
	// 1. 通过文件的形式
	// 2. 通过 http 接口
	pprofUsingHttp()
}

func pprofUsingHttp() {
	// localhost/debug/pprof
	// localhost/debug/pprof/goroutine?debug=2      查看 goroutine 情况
	// http://localhost/debug/pprof/allocs?debug=2  查看内存分配情况

	err := flag.Set("v", "4")
	if err != nil {
		log.Fatal(err)
	}

	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err = http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func pprofUsingFile() {
	// 运行完通过 go tool pprof /tmp/cpuprofile 查看
	cpuProfile := flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")
	flag.Parse()

	if *cpuProfile != "" {
		file, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}

		err = runtimePporf.StartCPUProfile(file)
		if err != nil {
			log.Fatal(err)
		}
		defer runtimePporf.StopCPUProfile()

		var result int
		for i := 0; i < 100000000; i++ {
			result += i
		}
		log.Println("result: ", result)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	writeString(&w, "ok")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler...")
	user := r.URL.Query().Get("user")
	if user == "" {
		user = "stranger"
	}
	writeString(&w, fmt.Sprintf("hello %s\n", user))

	writeString(&w, "================= Details of the http request header =================\n")
	for k, v := range r.Header {
		writeString(&w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func writeString(w *http.ResponseWriter, s string) {
	str, err := io.WriteString(*w, s)
	if err != nil {
		fmt.Println("write string error: ", err)
	}
	fmt.Println("write string: ", str)
}
