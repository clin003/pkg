package oauth2

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func HttpGetByte(geturl string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Transport: tr, Jar: cookieJar}

	resp, err := client.Get(geturl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func HttpGet(geturl string) (string, error) {
	if body, err := HttpGetByte(geturl); err != nil {
		return "", err
	} else {
		return string(body), nil
	}

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// //http cookie接口
	// cookieJar, _ := cookiejar.New(nil)
	// client := &http.Client{Transport: tr, Jar: cookieJar}

	// resp, err := client.Get(geturl)
	// if err != nil {
	// 	return "", err
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }

	// return string(body), nil
}
func HttpPost(geturl string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest("POST", geturl, nil)

	if err != nil {
		return "", err
	}

	response, _ := client.Do(reqest)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
