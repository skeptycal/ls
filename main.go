package main

import (
	"fmt"
	"log"
)

func main() {

	// fmt.Printf("listing for %s\n", pathFlag)
	dirs, err := Dirs(pathFlag, true, false)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range dirs {
		s := ""
		if fi.IsDir() {
			s = fi.Name()

			if sizeFlag {
				s += fmt.Sprintf(" %d", fi.Size())
			}

			s += outputSEP

			// fmt.Printf("%v %s %d\n", fi.Mode(), fi.Name(), fi.Size())
		}
		fmt.Println(s)

	}

}
