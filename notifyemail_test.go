package main

import ( // 导入包
	"testing" // 导入testing包
)

func init() {

	LoadConfig()
}

func Test_SendMail(t *testing.T) {

	err := SendMail([]string{"ke_dong@126.com"}, "subject", "body")
	if err != nil {
		t.Error(err)
	}
}
