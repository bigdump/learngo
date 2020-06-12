package main

import (
	"fmt"
	"strconv"
	"time"
)

func init() {
	regStruct = make(map[string]interface{})
	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{Str: "fasfdsa"}
}

type Foo struct {
}

type Bar struct {
	Str string
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

var mp = make(map[string][]byte)

func main() {
	//str := "Bar"
	//if regStruct[str] != nil {
	//	t := reflect.ValueOf(regStruct[str]).Type()
	//	v := reflect.New(t)
	//	fmt.Println(v)
	//}
	c := make([]byte, 100)
	c[1] = 13
	c[9] = 18

	number := 10
	var key []string
	var value [][]byte

	for i := 0; i < number; i++ {
		key = append(key, strconv.Itoa(i))
		value = append(value, intToBytes(i))

	}

	start0 := time.Now().Nanosecond()
	for i := 0; i < number; i++ {
		//mp[strconv.Itoa(i)] = append(c, intToBytes(i)...)
		mp[key[i]] = value[i]
	}

	start := time.Now().Nanosecond()
	ss := DeepCopy(mp)
	after := time.Now().Nanosecond()
	fmt.Println(ss["100"])
	fmt.Printf("create Use time : %v ns\n", start-start0)
	fmt.Printf("Use time : %v ns\n", after-start)
	fmt.Printf("one Use time : %v ns\n", (after-start)/10)
	fmt.Printf("Use time : %v us\n", (after-start)/1000)
	fmt.Printf("Use time : %v ms\n", (after-start)/1000/1000)
}

func intToBytes(value int) []byte {
	src := make([]byte, 4)
	src[3] = (byte)((value >> 24) & 0xFF)
	src[2] = (byte)((value >> 16) & 0xFF)
	src[1] = (byte)((value >> 8) & 0xFF)
	src[0] = (byte)(value & 0xFF)
	return src
}

func DeepCopy(mp map[string][]byte) map[string][]byte {
	copy := make(map[string][]byte)
	for k, v := range mp {
		copy[k] = v
	}
	return copy
}
