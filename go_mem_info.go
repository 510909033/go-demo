package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 获取程序当前占用内存（单位为字节）
	fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	// 获取程序最大可用内存（单位为字节）
	fmt.Printf("TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	// 获取程序已申请但未分配的内存（单位为字节）
	fmt.Printf("Sys = %v MiB\n", bToMb(m.Sys))
	// 获取程序堆内存占用情况（单位为字节）
	fmt.Printf("HeapAlloc = %v MiB\n", bToMb(m.HeapAlloc))
	fmt.Printf("HeapSys = %v MiB\n", bToMb(m.HeapSys))
	fmt.Printf("HeapIdle = %v MiB\n", bToMb(m.HeapIdle))
	fmt.Printf("HeapInuse = %v MiB\n", bToMb(m.HeapInuse))
}

// 将字节数转换为MiB
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
