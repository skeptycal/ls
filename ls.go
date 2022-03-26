package main

import (
	"os"
	"path/filepath"
)

const (
	NL  rune = '\n'
	NUL rune = 0
)

// var (
// 	pathFlag, excludeList                    string
// 	sizeFlag, oneFlag, zeroFlag, versionFlag bool
// 	home                                     string = getHome()
// 	pwd                                      string = getPWD()
// 	outputSEP                                string = " "
// )

// func init() {
// 	flag.StringVar(&pathFlag, "root", "", "root path to use for searches")
// 	flag.StringVar(&pathFlag, "r", "", "root path to use for searches (short)")
// 	flag.StringVar(&excludeList, "x", "", "list of excluded patterns")

// 	flag.BoolVar(&versionFlag, "version", false, "version of the package")
// 	flag.BoolVar(&versionFlag, "v", false, "version of the package (short)")
// 	flag.BoolVar(&sizeFlag, "size", false, "print the allocated size of each file, in blocks")
// 	flag.BoolVar(&zeroFlag, "0", false, "end each output line with NUL, not newline")
// 	flag.BoolVar(&oneFlag, "1", false, "list one file per line")

// 	flag.Parse()

// 	// files := flag.Args()[0]
// 	// fmt.Printf("files: %s\n", files)

// 	flag.PrintDefaults()

// 	if versionFlag {
// 		fmt.Printf("Version: %s\n", "fakeVersion 1.0.0")
// 		os.Exit(0)
// 	}

// 	if oneFlag {
// 		outputSEP = string(NL)
// 	}

// 	if zeroFlag {
// 		outputSEP = string(NUL)
// 	}

// 	if len(os.Args) > 2 {
// 		pathFlag = os.Args[1]
// 	} else {
// 		pathFlag = getPWD()
// 	}
// }

func getHome() (d string) {
	d, _ = os.UserHomeDir()
	return
}

func Env(s string) string {
	return os.ExpandEnv(s)
}

func getPWD() string {
	return Env(`echo $PWD`)
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
	return list, err
}
