package musdk

import (
	"testing"
)

func TestWeb(t *testing.T) {
	baseUrl := "http://x5.dev/mu"
	nodeId := 1
	token := "123"
	sType := TypeSs

	client := NewClient(baseUrl, token, nodeId, sType, nil)

	users, err := client.GetUsers()

	t.Log(err)
	t.Log(users)

}
