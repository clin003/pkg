package qiniu

import (
	"testing"
	"time"
)

func TestTokenTime(t *testing.T) {
	t1 := getUploadToken("xxx", "xxx", "xxx")
	t.Log(t1)
	// time.Sleep(72000 * time.Second)
	t2 := getUploadToken("xxx", "xxx", "xxx")
	t.Log(t2)
}
