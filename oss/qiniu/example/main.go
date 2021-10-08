package main

import (
	"fmt"

	. "gitee.com/lyhuilin/pkg/oss/qiniu"
)

func main() {
	filename := "/Users/baicailin/Downloads/143c3bbdd9dcb9ff591b225c2c784af3.mp4"
	buckername := "xxx"
	accessKey := "xxx"
	secretKey := "xxx"
	domain := "http://xxx.xxx"
	isOnly := true
	u, k, h, err := UploadFile(domain, filename, buckername, accessKey, secretKey, isOnly)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file1", u, k, h, err)

	filename = "http://xx.xx/xx.mp4"
	u2, k, h, err := UploadFileUrl(domain, filename, buckername, accessKey, secretKey, isOnly)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file2", u2, k, h, err)
}
