package storage

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSaveInstanceToDB(t *testing.T) {
	var key []string
	var value []string
	str := "{\"Nodes\":[1006,1007,1008,1009],\"Edges\":[1006,1007,1008,1009]}"
	for i := 0; i < 500; i++ {
		key = append(key, strconv.Itoa(i))
		value = append(value, strconv.Itoa(i)+"-"+str)
	}
	SaveInstanceToDB(key, value)
	GetInstanceFromDB()
}

type user struct {
	Id   int
	Name string
}

func TestEncode(t *testing.T) {
	data := user{
		Id:   0,
		Name: "test",
	}

	code := Encode(data)
	//fmt.Println(code)
	res := user{}
	Decode(&code, &res)
	fmt.Println(res)

}

func BenchmarkFormatUint(b *testing.B) {
	var num uint64 = 1024
	for i := 0; i < b.N; i++ {
		strconv.FormatUint(num, 10)
	}
}

func BenchmarkSprintf(b *testing.B) {
	var num uint64 = 1024
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}
