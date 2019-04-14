package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 4 {
		image_src := os.Args[1]
		image_width, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[convert] fatal error: bad width.\nConvertion terminated.")
			return
		}
		image_height, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("[convert] fatal error: bad height.\nConvertion terminated.")
			return
		}
		start_time := time.Now() // get current time
		BatchConvert(image_src, image_width, image_height)
		elapsed_time := time.Since(start_time)
		fmt.Printf("[covnert] all covert time elapsed:%s\n", elapsed_time)
	} else if len(os.Args) == 1 {
		fmt.Println("[convert] no input")
		Usage()
	} else {
		fmt.Println("[convert] fatal error: bad param number.\nConvertion terminated.")
		return
	}
}

func Usage() {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Internal name: goconvert.exe")
	fmt.Println("Version: v0.02_build20190414")
	fmt.Println("Author: zhaoguotao(guotao.zhao@vivo.com)")
	fmt.Println("Convert nv12 to nv21 or nv21 to nv12 and save the conversion to <convert> sub-folder.")
	fmt.Println("Usage:\n	convert.exe <input folder> w h")
	fmt.Println("----------------------------------------------------------")
}

func NV12ToNV21(image_src, image_dst string, image_width, image_height int) (length int, err error) {
	var y_size int = image_width * image_height
	var j int                                     //for count
	yuv_buffer, err := ioutil.ReadFile(image_src) //yuv_buffer type: []byte
	if err != nil {
		fmt.Printf("[covnert] error: %v\n", err)
		return 0, err
	}
	// swap u/v
	for j = 0; j < y_size/2; j += 2 {
		temp := yuv_buffer[j+y_size]
		yuv_buffer[y_size+j] = yuv_buffer[j+y_size+1]
		yuv_buffer[y_size+j+1] = temp
	}
	err = ioutil.WriteFile(image_dst, yuv_buffer, 0644)
	if err != nil {
		fmt.Printf("[covnert] error: %v\n", err)
		return 0, err
	}
	return len(yuv_buffer), nil
}

func BatchConvert(fpath string, w, h int) {
	f_path := strings.Replace(fpath, "\\", "/", -1)
	convert_folder, err1 := mkdirSubConvert(f_path)
	if err1 != nil {
		fmt.Printf("[covnert] error: %v\n", err1)
		return
	}

	files, err := ioutil.ReadDir(f_path) //read folder
	if err != nil {
		fmt.Printf("[covnert] error: %v\n", err)
	}

	for _, f := range files {
		if !f.IsDir() {
			fyuv := f.Name()
			ok := strings.HasSuffix(fyuv, ".yuv") // Only handle yuv file
			if ok {
				image_src := f_path + "/" + fyuv
				image_dst := convert_folder + "/" + fyuv
				start_time := time.Now() // get current time
				size, err := NV12ToNV21(image_src, image_dst, w, h)
				elapsed_time := time.Since(start_time)
				if err == nil {
					fmt.Printf("[covnert] %s: (Width:%d; Height:%d; Size:%08x B; TimeElapsed:%s;)\n", image_src, w, h, size, elapsed_time)
				}
			}
		}
	}
}

func mkdirSubConvert(ppath string) (fpath string, err error) {
	convert_foler := ppath + "/convert"
	exist, err := PathExists(convert_foler)
	if err != nil {
		fmt.Printf("[convert] error: %v\n", err)
		return "", err
	}
	if !exist {
		err := os.Mkdir(convert_foler, os.ModePerm)
		if err != nil {
			return "", err
		} else {
			return convert_foler, nil
		}
	}
	return convert_foler, nil
}

// folder exist?
func PathExists(fpath string) (bool, error) {
	_, err := os.Stat(fpath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
