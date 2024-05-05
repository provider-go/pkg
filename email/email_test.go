package email

import (
	"github.com/provider-go/pkg/email/typeemail"
	"testing"
)

func TestName(t *testing.T) {
	c := typeemail.ConfigSendEmail{
		Host:     "smtp.qq.com",
		Port:     465,
		User:     "69401295@qq.com",
		AuthCode: "ebqkvnfwvpumbjgh",
	}

	client := NewSendEmail("qq", c)
	err := client.SMTP("biwow@qq.com",
		"你擅长什么编程语言？",
		"我擅长Java编程语言。\nJava是一种广泛应用的编程语言，"+
			"特别适合于企业级应用和Android移动应用开发。"+
			"Java具有强类型、面向对象和跨平台的特点，能够确保程序的安全性和可靠性。"+
			"此外，Java拥有丰富的第三方库和工具，以及庞大的开发社区和生态系统，"+
			"这些特点使得Java成为开发人员的重要选择。")
	if err != nil {
		t.Log(err)
	}
	t.Log("发送成功")
}
