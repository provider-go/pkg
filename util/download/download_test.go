package util

import "testing"

func Test_DownloadFile(t *testing.T) {
	err := DownloadFile("http://120.53.243.73/tencent_sms.zip", "centos.iso")
	if err != nil {
		t.Log(err)
	}
}
