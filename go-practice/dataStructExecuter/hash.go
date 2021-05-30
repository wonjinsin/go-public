package main

import (
	"fmt"

	"github.com/wonjinsin/go-practice/dataStruct"
)

func main() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcdf = ", dataStruct.Hash("abcdf"))
	fmt.Println("tbcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("abcdefdfdddfdf = ", dataStruct.Hash("abcdefdfdddfdf"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CDEFRGTEFVDF", "0111111111")
	m.Add("CCC", "017575757575")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
	fmt.Println("CDEFRGTEFVDF = ", m.Get("CDEFRGTEFVDF"))
}
