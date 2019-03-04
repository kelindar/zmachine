package main

import (
	"io/ioutil"
	"os"
)

func main() {
	buffer, err := ioutil.ReadFile("zork1.dat")
	if err != nil {
		panic(err)
	}

	var header ZHeader
	header.read(buffer)

	if header.Version != 3 {
		panic("Only Version 3 files supported")
	}

	var zm *ZMachine
	if save, err := os.Open("save.bin"); err == nil {
		zm, _ = Load(save)
	} else {
		zm = New(buffer, header)
	}

	for !zm.Done {
		zm.InterpretInstruction()
	}
}
