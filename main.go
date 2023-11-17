package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rprtr258/scuf"
)

func main() {
	log.SetFlags(log.Lshortfile)

	var b []byte
	switch len(os.Args) {
	case 1:
		var err error
		b, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Printf("read stdin: %s\n", err.Error())
			os.Exit(1)
		}
	case 2:
		var err error
		b, err = os.ReadFile(os.Args[1])
		if err != nil {
			log.Printf("read file: %s\n", err.Error())
			os.Exit(1)
		}
	default:
		log.Println("Usage: xxd <filename>")
		os.Exit(1)
	}

	_specialStyle := scuf.FgRGB(255, 0, 0)
	for _, r := range b {
		switch {
		case '!' <= r && r <= '~':
			fmt.Printf("%c", r)
		case r == '\x00':
			fmt.Print(scuf.String(`\0`, _specialStyle))
		case r == ' ':
			fmt.Print(scuf.String(" ", _specialStyle))
		case r == '\n':
			fmt.Print(scuf.String(`\n`, _specialStyle))
		case r == '\t':
			fmt.Print(scuf.String(`\t`, _specialStyle))
		default:
			fmt.Print(scuf.String(fmt.Sprintf("%02x ", r), scuf.FgRGB(255-r, r, 0)))
		}
	}
}
