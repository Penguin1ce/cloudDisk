package utils

import (
	"cloudDisk/core/models"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var user models.UserBasic
	user.Id = 5
	user.Identity = "User"
	user.Password = "111"

	token, err := GenerateToken(5, "User", "111")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestMd5(t *testing.T) {
	pass := Md5("111")
	t.Log(pass)
}

func TestRandString(t *testing.T) {
	s := RandomCode()
	t.Log(s)
}
