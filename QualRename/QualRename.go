package main

import (
	//	"flag"
	"fmt"
	//	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fin := os.Args[1]
		fmt.Println("Input folder: ", fin)
		fmt.Println("-----------------------------------------------")
		BatchRun(fin)
	} else {
		fmt.Println("No input folder, stopped...\n")
		fmt.Println("Usage:")
		fmt.Println("QualRename.exe <input folder>")
	}

}

func RenameFile(fname, new_name string) {
	// Rename or move file from one location to another.
	ret := os.Rename(fname, new_name)
	if ret != nil {
		fmt.Println(ret)
		return
	} else {
		fmt.Printf("rename %s to %s\n", fname, new_name)
		return
	}
}

func BatchRun(path string) {
	r, _ := regexp.Compile("20\\d{6}_\\d{6}|20\\d{12}")
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// Only handle RAWPLAIN16 file
		ok := strings.HasSuffix(path, ".RAWPLAIN16")
		if ok {
			if r.FindString(path) != "" {
				fpre := filepath.Dir(path) //去除最后一个元素的路径
				new_name := fpre + "/IMG_" + r.FindString(path) + ".raw"
				RenameFile(path, new_name)
				//fmt.Println(path + ">> " + new_name)
			} else {
				fmt.Println(path, ": NO match....")
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
