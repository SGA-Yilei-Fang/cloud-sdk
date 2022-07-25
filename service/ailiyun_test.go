package service

import (
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {

}

func TestCreateClient(t *testing.T) {
	accessKeyId := "mock_id"
	accessKeySecret := "mock_secret"
	if ans, _err := CreateClient(&accessKeyId, &accessKeySecret); ans == nil {
		fmt.Println(_err)
	}
}
