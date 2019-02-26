package gohttpclient

import (
	"testing"
)

func TestClientGet(t *testing.T) {
	header := make(map[string]string)
	header["aa"] = "bb"
	resp := NewClient(3).SetHeader(header).Get("https://www.baidu.com")
	if len(resp) == 0 {
		t.Errorf("get resp content : %s", resp)
	}
}

func TestClientPost(t *testing.T) {
	header := make(map[string]string)
	header["Cookie"] = "username=test; _guid=R50603bc-4cc3-94d8-01af-9e382f116440; new_uv=1; account=hengzhao.wei; account_chk=e7336e32ba25c998bb8dcdfc1f168761"

	form := make(map[string][]string)
	form["content"] = []string{"goDev"}
	form["fmtType"] = []string{"1"}

	resp := NewClient(10).SetHeader(header).Post("http://api.bejson.com/btools/tools/convert/camelUnderscore", form)
	if string(resp) != "go_dev" {
		t.Errorf("post err")
	}
}
