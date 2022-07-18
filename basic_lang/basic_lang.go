package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/gops/agent"
	"github.com/siddontang/go-log/log"
	"net/http"
	"sync"
	"sync/atomic"
)

var once Once
var tstStruct = &TstStruct{}

type TstStruct struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
}

func main() {
	//make_rs := make([]int, 0)
	//new_rs := new([]int)
	//direct := []int{}
	//fmt.Printf("make:%+v\n", make_rs)
	//fmt.Printf("new:%+v\n", new_rs)
	//direct = append(direct, 1)
	//fmt.Printf("[]int :%+v\n", direct)

	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < 100000; i++ {
	//	wg.Add(1)
	//	go func() {
	//		once.Do(func() {
	//			tstStruct = &TstStruct{
	//				A: "string",
	//				B: "string",
	//				C: "string",
	//				D: test(),
	//				E: "string",
	//				F: "string",
	//				G: "string",
	//				H: test(),
	//			}
	//
	//		})
	//		if tstStruct.A == "" || tstStruct.B == "" || tstStruct.C == "" || tstStruct.D == "" || tstStruct.E == "" ||
	//			tstStruct.F == "" || tstStruct.G == "" || tstStruct.H == "" {
	//			fmt.Printf("tstStruct :%+v\n", tstStruct)
	//		}
	//	}()
	//	wg.Done()
	//}
	//wg.Wait()

	//value := unsafe.Sizeof(uint8(1))
	//fmt.Printf("rs: %d\n", value)

	//slice_value := new([5]int)
	//a := cap(*slice_value)
	//fmt.Printf("rs: %d\n", a)

	// 创建并监听 gops agent，gops 命令会通过连接 agent 来读取进程信息
	// 若需要远程访问，可配置 agent.Options{Addr: "0.0.0.0:6060"}，否则默认仅允许本地访问
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatalf("agent.Listen err: %v", err)
	}

	http.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		fmt.Printf("req: %+v\n", req)

		str := "Have a  nice data, boy!"
		value, err := rsp.Write([]byte(str))
		fmt.Printf("value:%v, err:%v", value, err)

	})
	err := http.ListenAndServe(":9091", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
	//err := http.ListenAndServe(":9091", &HttpHandler{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("serve start success!")
}

type HttpHandler struct {
}

func (*HttpHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	fmt.Printf("req: %+v\n", req)

	str := "Have a  nice data, boy!"
	value, err := rsp.Write([]byte(str))
	fmt.Printf("value:%v, err:%v", value, err)

}

type Once struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/386),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		//if o.done == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)

		//o.done = 1

		f()
	}
}

func test() string {
	rs, _ := json.Marshal(&TstStruct{
		A: "string",
		B: "string",
		C: "string",
		D: "fdhsjakfhsafhausdifahdfpioahfsdjafosaj",
		E: "string",
		F: "string",
		G: "string",
		H: "string",
	})
	return string(rs)
}
