package controllers

import (
	"log"

	"math/rand"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SmsController struct {
	beego.Controller
}

//接続クライアントの生成
func GetClient(AccessKey, SecretKey, Region string) (*sns.SNS, error) {
	creds := credentials.NewStaticCredentials(AccessKey, SecretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(Region),
	})

	if err != nil {
		log.Println("Error!")
		log.Println(err)
		return nil, err
	}

	return sns.New(sess), nil
}

func createInputMessage(Message, PhoneNumber string) *sns.PublishInput {
	pin := &sns.PublishInput{}
	pin.SetMessage(Message)
	pin.SetPhoneNumber(PhoneNumber)
	return pin
}

func (this *SmsController) SendMessage() {
	userPhoneNumber := this.GetString("phoneNumber")
	client, _ := GetClient(beego.AppConfig.String("AccessKey"), beego.AppConfig.String("SecretKey"), beego.AppConfig.String("Region"))

	randamNum := RandNum(6)
	msgContent := "TeDeの確認コードは次の通りです : " + randamNum
	msg := createInputMessage(msgContent, "+81"+userPhoneNumber)

	client.Publish(msg)

	this.Data["json"] = randamNum
	this.ServeJSON()
}

var ints = []rune("1234567890")

func RandNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = ints[rand.Intn(len(ints))]
	}
	return string(b)
}
