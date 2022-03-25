package gofile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	NL  = '\n'
	NUL = 0
)

var (
	pathFlag, excludeList                    string
	sizeFlag, oneFlag, zeroFlag, versionFlag bool
	home                                     string = getHome()
	pwd                                      string = getPWD()
	outputSEP                                string = " "
)

func init() {
	flag.StringVar(&pathFlag, "root", "", "root path to use for searches")
	flag.StringVar(&pathFlag, "r", "", "root path to use for searches (short)")
	flag.StringVar(&excludeList, "x", "", "list of excluded patterns")

	flag.BoolVar(&versionFlag, "version", false, "version of the package")
	flag.BoolVar(&versionFlag, "v", false, "version of the package (short)")
	flag.BoolVar(&sizeFlag, "size", false, "print the allocated size of each file, in blocks")
	flag.BoolVar(&zeroFlag, "0", false, "end each output line with NUL, not newline")
	flag.BoolVar(&oneFlag, "1", false, "list one file per line")

	flag.Parse()

	// files := flag.Args()[0]
	// fmt.Printf("files: %s\n", files)

	flag.PrintDefaults()

	if versionFlag {
		fmt.Printf("Version: %s\n", "fakeVersion 1.0.0")
		os.Exit(0)
	}

	if oneFlag {
		outputSEP = string(NL)
	}

	if zeroFlag {
		outputSEP = string(NUL)
	}

	if len(os.Args) > 2 {
		pathFlag = os.Args[1]
	} else {
		pathFlag = getPWD()
	}

}

func getHome() (d string) {
	d, _ = os.UserHomeDir()
	return
}

func getPWD() (d string) {
	d, _ = os.Getwd()
	return
}

func Ls(root, pattern string) (list []os.FileInfo, err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			list = append(list, info)
			return nil
		})
	return
}

func DirsRecursive(root, pattern string) (list []os.FileInfo, err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				list = append(list, info)
			}
			return nil
		})
	return
}

func Dirs(root string, dirsOnly bool, showHidden bool) ([]os.FileInfo, error) {
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
	return list, nil
}

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
