package main

import "io/ioutil"

func main() {
	buffer, err := ioutil.ReadFile("zork1.dat")
	if err != nil {
		panic(err)
	}

	var header ZHeader
	header.read(buffer)

	if header.version != 3 {
		panic("Only Version 3 files supported")
	}

	zm := New(buffer, header)
	for !zm.done {
		zm.InterpretInstruction()
	}
}
