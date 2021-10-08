package main

import (
	"testing"

	. "gitee.com/lyhuilin/pkg/oss/qiniu"
)

func TestUploadFile(t *testing.T) {
	filename := "/Users/baicailin/Downloads/_b_Bfc53d0507e49291bdf5c0a8700571222.mp4"
	buckername := "xxx"
	accessKey := "xxx"
	secretKey := "xxx"
	isOnly := true
	u, k, h, err := UploadFile(domain, filename, buckername, accessKey, secretKey, isOnly)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(u, k, h)
}
