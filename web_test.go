package musdk

import (
	"testing"
)

func TestWeb(t *testing.T) {
	baseUrl := "http://x5.dev/mu"
	nodeId := 1
	token := "123"

	client := NewClient(baseUrl, token, nodeId)

	users, err := client.GetUsers()

	t.Log(err)
	t.Log(users)

}
