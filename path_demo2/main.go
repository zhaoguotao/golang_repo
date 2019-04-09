package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(filepath.Dir(os.Args[0]))
	fmt.Println(filepath.Abs(os.Args[0])) //返回绝对路径
	dir, _ := filepath.Abs(os.Args[0])    //返回绝对路径
	fmt.Println(filepath.Dir(dir))        //去除最后一个元素的路径

	f := GetCurrentDirectory()
	fmt.Println(f)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
