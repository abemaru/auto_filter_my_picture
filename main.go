package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	// ToDo: Support multiple files
	files, err := ioutil.ReadDir("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fname := "input/" + file.Name()
		f, err := os.Open(fname)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		exif.RegisterParsers(mknote.All...)
		fmt.Println(f.Name())

		// ToDo: create file when sorting pictures

		// fileInfo, err := os.Lstat("./")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fileMode := fileInfo.Mode()
		// unixPerms := fileMode & os.ModePerm
		// if err := os.MkdirAll("case/case01", unixPerms); err != nil {
		// 	log.Fatal(err)
		// }

	}
}
