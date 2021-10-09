package qiniu

import (
	"bytes"
	"context"
	"path/filepath"
	"sync"
	"time"

	"gitee.com/lyhuilin/util"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var tokenMap map[string]time.Time
var mu sync.Mutex

func init() {
	tokenMap = make(map[string]time.Time)
}
func getUploadToken(buckername string, accessKey, secretKey string) string {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	for k, v := range tokenMap {
		if now.Before(v) {
			return k
		}
	}

	bucket := buckername
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200
	mac := auth.New(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	tokenMap[upToken] = now.Add(7200 * time.Second)
	return upToken
}
func UploadFile(domain, file string, buckername string, accessKey, secretKey string, isOnly bool) (publicAccessURL, key, hash string, err error) {
	localFile := file
	// 	ct := mime.TypeByExtension(filepath.Ext(u))
	basename := filepath.Base(file)
	// basename := util.EncryptMd5(file)

	upToken := getUploadToken(buckername, accessKey, secretKey)

	cfg := storage.Config{}
	// 空间对应的机房 可以指定空间对应的Zone(如不指定Zone则会使用自动判断区域)以及其他的一些影响上传的参数。
	// cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": basename,
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, basename, localFile, &putExtra)
	if err != nil {
		// fmt.Println(err)
		return "", "", "", err
	}
	// fmt.Println(ret.Key, ret.Hash)

	publicAccessURL = storage.MakePublicURL(domain, basename)
	// fmt.Printf("publicAccessURL:%s\n", publicAccessURL)

	if !isOnly {
		return publicAccessURL, ret.Key, ret.Hash, nil
	}

	hashMd5 := util.EncryptMd5(ret.Hash)
	if err := renameBucketFile(accessKey, secretKey, buckername, ret.Key, hashMd5, true); err != nil {
		return publicAccessURL, ret.Key, ret.Hash, err
	} else {
		publicAccessURL = storage.MakePublicURL(domain, hashMd5)
		return publicAccessURL, hashMd5, ret.Hash, nil
	}
}
func UploadFileByte(domain, fileName, buckername, accessKey, secretKey string, data []byte, isOnly bool) (pubURL, key, hash string, err error) {
	// 	ct := mime.TypeByExtension(filepath.Ext(u))
	basename := filepath.Base(fileName)
	if isOnly {
		basename = util.EncryptMd5(fileName)
	} else {
		basename = util.EncryptMd5Byte(data)
	}

	upToken := getUploadToken(buckername, accessKey, secretKey)

	cfg := storage.Config{}
	// 空间对应的机房 可以指定空间对应的Zone(如不指定Zone则会使用自动判断区域)以及其他的一些影响上传的参数。
	// cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": basename,
		},
	}

	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, upToken, basename, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		// fmt.Println(err)
		return "", "", "", err
	}
	// fmt.Println("k2:", ret.Key, ret.Hash)

	pubURL = storage.MakePublicURL(domain, basename)
	// fmt.Printf("pubURL:%s\n", pubURL)

	if !isOnly {
		return pubURL, ret.Key, ret.Hash, nil
	}

	hashMd5 := util.EncryptMd5(ret.Hash)
	if err := renameBucketFile(accessKey, secretKey, buckername, ret.Key, hashMd5, true); err != nil {
		return pubURL, ret.Key, ret.Hash, err
	} else {
		pubURL = storage.MakePublicURL(domain, hashMd5)
		return pubURL, hashMd5, ret.Hash, nil
	}
}

func UploadFileUrl(domain, fileUrl string, buckername string, accessKey, secretKey string, isOnly bool) (publicAccessURL, key, hash string, err error) {
	data, err := util.GetUrlToByte(fileUrl)
	if err != nil {
		return "", "", "", err
	}
	return UploadFileByte(domain, fileUrl, buckername, accessKey, secretKey, data, isOnly)
}

func UploadFileUrlEx(domain, fileUrl string, buckername string, accessKey, secretKey string, isOnly bool, timeOut time.Duration) (publicAccessURL, key, hash string, err error) {
	data, err := util.GetUrlToByteEx(fileUrl, timeOut)
	if err != nil {
		return "", "", "", err
	}
	return UploadFileByte(domain, fileUrl, buckername, accessKey, secretKey, data, isOnly)
}
func renameBucketFile(accessKey, secretKey, buckername, srcKey, destKey string, force bool) error {
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	//如果目标文件存在，是否强制覆盖，如果不覆盖，默认返回614 file exists
	// force := true
	// err := bucketManager.Move(srcBucket, srcKey, destBucket, destKey, force)
	err := bucketManager.Move(buckername, srcKey, buckername, destKey, force)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}
