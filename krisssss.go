package main

import (
	"fmt"
)

 const sl = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

func IterateOverASCIIStringLiteral(sl string) {
	//sl = []string {"x80", "x81", "x82", "x83", "x84", "x85", "x86", "x87", "x88", "x89", "x8A", "x8B", "x8C", "x8D", "x8E",
	//"x8F", "x90", "x91", "x92", "x93", "x94", "x95", "x96", "x97", "x98", "x99", "x9A", "x9C", "x9D", "x9E", "x9F", "xA0",
	//"xA1", "xA2", "xA3", "xA4", "xA5", "xA6"}
	

	
	for i := 0; i < len(sl); i++ {
	fmt.Printf("%X %c %b \n", sl[i], sl[i], sl[i])
	}
}

func main(){
IterateOverASCIIStringLiteral(sl)
}
