package httpcode

import (
	"testing"
)

func TestNewRequestGet(t *testing.T) {
	var (
		url  = "https://www.baidu.com"
		rGet = NewRequestGet(url)
	)
	rGet.AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := rGet.Call()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("resp: %s", resp)

}

func TestNewRequestPostJson(t *testing.T) {
	var (
		url  = "http://www.baidu.com"
		data = map[string]interface{}{
			"name": "light",
		}
		pJson = NewRequestPostJson(url, data)
	)
	pJson.AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := pJson.Call()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("resp: %s", resp)
}

func TestNewRequestPostForm(t *testing.T) {
	var (
		url  = "https://www.baidu.com"
		data = map[string]string{
			"name": "light",
		}
		pForm = NewRequestPostForm(url, data)
	)
	pForm.AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := pForm.Call()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("resp: %s", resp)
}

func TestNewRequestPostFormWithFile(t *testing.T) {
	var (
		url  = "https://www.baidu.com"
		data = map[string]string{
			"name": "light",
		}
		files = map[string][]byte{
			"content": []byte("this is content"),
		}
		pFile = NewRequestPostFormWithFile(url, data, files)
	)
	pFile.AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := pFile.Call()
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("resp: %s", resp)
}
