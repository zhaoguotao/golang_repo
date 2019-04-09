package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fin := "C:\\Users\\Administrator\\go\\src\\github.com\\zhaoguotao\\Rename\\in\\go.gif"
	fout := "C:\\Users\\Administrator\\go\\src\\github.com\\zhaoguotao\\Rename\\out\\go.gif"

	// Rename or move file from one location to another.
	ret := os.Rename(fin, fout)
	if ret != nil {
		fmt.Println(ret)
		return
	} else {
		fmt.Println("sucess...")
		return
	}
}
