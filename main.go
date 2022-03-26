package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	pathFlag := "."
	sizeFlag := true
	outputSEP := "\n"

	if len(os.Args) > 1 {
		pathFlag = os.Args[1]
	}

	// fmt.Printf("listing for %s\n", pathFlag)
	dirs, err := GetDirs(pathFlag, true, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("dir list:")
	for _, dir := range dirs {
		fmt.Printf(" %v", dir.Name())
	}
	fmt.Println()

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
		fmt.Print(s)

	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		// a := new(ansiColorCode)
		var v ansiColorCode = 1
		fmt.Println(v)
	}

}

func GetDirs(root string, dirsOnly bool, showHidden bool) ([]os.FileInfo, error) {
	dirs, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	list := make([]os.FileInfo, 0, len(dirs))
	for _, fi := range dirs {
		if !showHidden && fi.Name()[0] == '.' {
			continue
		}

		if dirsOnly && !fi.IsDir() {
			continue
		}
		list = append(list, fi)
	}
	return list, err
}
