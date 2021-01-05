package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsController struct {
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

func (this *SnsController) SendMessage() {
	client, err := GetClient(beego.AppConfig.String("AccessKey"), beego.AppConfig.String("SecretKey"), beego.AppConfig.String("Region"))
	if err != nil {
		log.Println(err)
	}
	msg := createInputMessage("これはてすとめっせ", "+81")

	reasult, err := client.Publish(msg)

	if err != nil {
		log.Println(err)
	}

	log.Println(reasult)
}
