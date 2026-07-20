package main

import (
	"fmt"
	com "github.com/matesu777/Mafrana/internal/components"
	"log"
	"time"
)

func main() {
	teste()
}

func teste() {
	var mem com.Memory

	cpu := com.NewCpu()
	var disk com.Disk

	if err := mem.Scan(); err != nil {
		log.Fatal(err)
	}
	if err := disk.Scan(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", mem)
	fmt.Printf("%+v\n", disk)

	cpu.Scan()
	time.Sleep(time.Second)

	cpu.Scan()

	fmt.Printf("usage: %.2f%% cores: %d\n", cpu.Usage, cpu.Cores)
}
