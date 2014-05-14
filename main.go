package main

import "fmt"
import "debug/pe"
import "os"
import "log"

func main() {
	exepath := os.Args[1]
	exef, e := pe.Open(exepath)
	die(e)
	defer exef.Close()
	switch exef.FileHeader.Machine {
	case 0x14c:
		fmt.Println("i386")
	case 0x8664:
		fmt.Println("amd64")
	case 0x1c0:
		fmt.Println("arm")
	case 0x200:
		fmt.Println("ia64")
	case 0x0:
		fallthrough
	default:
		fmt.Println("unknown")

	}

}

func die(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
