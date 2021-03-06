package test

import (
	"creeper/orm_model"
	"creeper/service"
	"log"
	"testing"
)

func Test_create_access_token(t *testing.T) {
	ot := service.CreateAccessToken(1, "24cd1a8fa6da73dc78e2f56619cc70de")
	if ot.Error != nil {
		t.Error(ot.Error)
		t.Fail()
		return
	}
	log.Println(ot.Data.(*orm_model.AccessToken).Token)
	//第二次创建
	ot1 := service.CreateAccessToken(1, "24cd1a8fa6da73dc78e2f56619cc70de")
	if ot1.Error != nil {
		t.Error(ot1.Error)
		t.Fail()
		return
	}
	log.Println(ot1.Data.(*orm_model.AccessToken).Token)
}
