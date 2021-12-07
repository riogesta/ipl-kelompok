package main

import (
	"fmt"
	"time"
)

type Single struct {
	Conn string
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		singleInstance = &Single{"Koneksi Database"}
		fmt.Println("Instansiasi sudah terbuat sekarang")
	}

	return singleInstance
}

func main() {
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Println(*getInstance(), " - ", i)
		}
	}()

	go func() {
		fmt.Println(*getInstance())
	}()

	fmt.Scanln()
}