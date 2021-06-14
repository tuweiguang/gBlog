package main

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	GetMemory()
	GetCPU()
	GetDisk()
	GetGoroutineNum()
}

func GetMemory() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}

func GetCPU() {
	v, _ := cpu.Info()

	fmt.Println(v)
}

func GetDisk() {
	v, _ := disk.Partitions(true)

	fmt.Println(v)

	vv, _ := disk.Usage("/")
	fmt.Println(vv)
}

func GetGoroutineNum() {
	fmt.Println(runtime.NumGoroutine())
}
