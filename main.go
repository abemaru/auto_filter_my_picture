package main

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	// ToDo: Support multiple files
	fname := "input/PXL_20220702_091647779.jpg"

	f, err := os.Open(fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	exif.RegisterParsers(mknote.All...)
	fmt.Println(f)

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
