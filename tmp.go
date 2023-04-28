package main

import (
	// "arena"
	"errors"
	"fmt"
	"runtime"
)

// type student struct {
// Name string
// }

// func zhoujielun(v interface{}) {
// switch msg := v.(type) {
// case *student, student:
// 	// msg.Name
// }
// }

// type People struct {
// 	name string `json:"name"`
// }

// func main() {
// 	js := `{
// 		"name":"11"
// 	}`
// 	var p People
// 	err := json.Unmarshal([]byte(js), &p)
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}
// 	fmt.Println("people: ", p)
// }

// type People struct{}

// func (p *People) ShowA() {
// 	fmt.Println("showA")
// 	p.ShowB()
// }
// func (p *People) ShowB() {
// 	fmt.Println("showB")
// }

// type Teacher struct {
// 	People
// }

// func (t *Teacher) ShowB() {
// 	fmt.Println("teacher showB")
// }

// func main() {
// 	t := Teacher{}
// 	t.ShowA()
// 	t.ShowB()
// 	t.People.ShowB()
// }

// type People interface {
// 	Speak(string) string
// }

// type Student struct{}

// func (stu *Student) Speak(think string) (talk string) {
// 	if think == "bitch" {
// 		talk = "You are a good boy"
// 	} else {
// 		talk = "hi"
// 	}
// 	return
// }

// func main() {
// 	var peo People = &Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// func live() People {
// 	var stu Student
// 	return stu
// }

// func main() {
// 	out := make(chan int, 10)
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			out <- rand.Intn(5)
// 		}
// 		close(out)
// 		fmt.Println("closed..", len(out), cap(out))
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		// for i := range out {
// 		// 	fmt.Println(i)
// 		// }
// 		for {
// 			select {
// 			case i, ok := <-out:
// 				if !ok {
// 					return
// 				}
// 				fmt.Println(i, ok)
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}()
// 	wg.Wait()
// }

type T struct {
	Foo string
	Bar [16]byte
}

/*
提供的 arena API

NewArena：创建一个新的 arena 内存空间。
Free：释放 arena 及其关联对象。
New：基于 arena，创建新对象。
MakeSlice：基于 arena，创建新切片。
Clone：克隆一个 arena 的对象，并移动到内存堆上。
*/
func main() {
	getmem()
	// 在函数开头创建一个 arena
	// mem := arena.NewArena()
	// // 在函数结束时释放 arena
	// defer mem.Free()

	// // 从申请的 arena 中申请一些对象
	// for i := 0; i < 1000000; i++ {
	// 	obj := arena.New[T](mem)
	// 	// fmt.Println("obj", obj)
	// 	obj.Foo = fmt.Sprintf("foo=%d", i)
	// 	// fmt.Println("obj", obj.Foo)
	// }

	// //arena.MakeSlice[string](mem, length, capacity "string")
	// // 从申请的 arena 中申请切片对象（指定长度和容量）
	// // slice := arena.MakeSlice[T](mem, 100, 200)
	// // fmt.Println("slice", slice)

	// src := "脑子进煎鱼了"
	// bs := arena.MakeSlice[byte](mem, len(src), len(src))
	// copy(bs, src)
	// str := unsafe.String(&bs[0], len(bs))
	// fmt.Println("str=", str)

	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
}

//	func Stop(stop <-chan bool) {
//		close(stop)
//	}
func getmem() {
	runtime.GOMAXPROCS(0)
	// // 获取当前进程的基本信息
	// pid := os.Getpid()

	// // 获取当前进程的内存信息
	// goInfo := runtime.GOMAXPROCS(0)
	// var memInfo runtime.MemInfo
	// runtime.ReadMemInfo(&memInfo)

	// // 输出自身占用内存大小
	// fmt.Printf("My goroutine pid: %d, goroutine size: %d bytes\n", pid, memInfo.TotalSize)

	// // 使用 malloc 申请自身内存
	// mem := make([]byte, 1)

	// // 输出自身占用内存大小
	// fmt.Printf("My goroutine pid: %d, malloc size: %d bytes\n", pid, mem)
}
