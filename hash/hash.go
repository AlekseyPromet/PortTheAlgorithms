package main

import "fmt"

func main() {
	var char string = "name"
	fmt.Println("SDBM", sdbm(char))
	fmt.Println("DJB2", djb2(char))
	fmt.Println("xor8", xor8(char))
	fmt.Println("adler32", adler32(char))

}

func sdbm(s string) (hash int32) {

	for _, char := range s {
		hash = char + (hash << 6) + (hash << 16) - hash

	}
	return hash
}

func djb2(s string) (hash int32) {

	hash = 5381 // init value
	for _, char := range s {
		hash = ((hash << 5) + hash) + char
	}
	return hash
}

func xor8(s string) (hash int32) {

	for _, char := range s {
		hash = (hash + char) & 0xff
	}
	hash = ((hash ^ 0xff) + 1) & 0xff
	return hash
}

func adler32(s string) (hash int32) {
	var a int32 = 1
	var b int32 = 0

	const MODADLER int32 = 65521
	for _, char := range s {
		a = (a + char) % MODADLER
		b = (b + a) % MODADLER
	}
	return (b << 16) | a
}
