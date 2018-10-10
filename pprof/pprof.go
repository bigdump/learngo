package main

import (
	"fmt"
	"runtime"
)

func main() {
	var su = [100]int{100, 200}
	m2 := make(map[int]int)
	for i := 0; i < 1000; i++ {
		su[0]++
		m2[i] = i
	}
	s := new(runtime.MemStats)
	runtime.ReadMemStats(s)
	fmt.Printf("\n# runtime.MemStats\n")
	fmt.Printf("# Alloc = %d\n", s.Alloc)
	fmt.Printf("# TotalAlloc = %d\n", s.TotalAlloc)
	fmt.Printf("# Sys = %d\n", s.Sys)
	fmt.Printf("# Lookups = %d\n", s.Lookups)
	fmt.Printf("# Mallocs = %d\n", s.Mallocs)
	fmt.Printf("# Frees = %d\n", s.Frees)

	fmt.Printf("# HeapAlloc = %d\n", s.HeapAlloc)
	fmt.Printf("# HeapSys = %d\n", s.HeapSys)
	fmt.Printf("# HeapIdle = %d\n", s.HeapIdle)
	fmt.Printf("# HeapInuse = %d\n", s.HeapInuse)
	fmt.Printf("# HeapReleased = %d\n", s.HeapReleased)
	fmt.Printf("# HeapObjects = %d\n", s.HeapObjects)

	fmt.Print("# Stack = %d / %d\n", s.StackInuse, s.StackSys)
	fmt.Print("# MSpan = %d / %d\n", s.MSpanInuse, s.MSpanSys)
	fmt.Print("# MCache = %d / %d\n", s.MCacheInuse, s.MCacheSys)
	fmt.Print("# BuckHashSys = %d\n", s.BuckHashSys)
	fmt.Print("# GCSys = %d\n", s.GCSys)
	fmt.Print("# OtherSys = %d\n", s.OtherSys)

	fmt.Print("# NextGC = %d\n", s.NextGC)
	fmt.Print("# LastGC = %d\n", s.LastGC)
	fmt.Print("# PauseNs = %d\n", s.PauseNs)
	fmt.Print("# PauseEnd = %d\n", s.PauseEnd)
	fmt.Print("# NumGC = %d\n", s.NumGC)
	fmt.Print("# NumForcedGC = %d\n", s.NumForcedGC)
	fmt.Print("# GCCPUFraction = %v\n", s.GCCPUFraction)
	fmt.Print("# DebugGC = %v\n", s.DebugGC)
}
