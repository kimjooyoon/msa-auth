package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"msa-auth/database"
	"msa-auth/members"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Run("integration scenario 1", func(t *testing.T) {
		database.AllDeleteRows()
		// sign-on
		b, err := json.Marshal(members.SignOnDto{
			"test@test.test", "test",
			"tester", "test-user", "01012341234",
		})
		if err != nil {
			log.Panicf("%v", err)
		}
		q, status := ClientE("/sign-on", "POST", b)
		assert.Contains(t, q, "message")
		assert.Equal(t, status, "200 OK")

		// sign-in
		b, err = json.Marshal(members.SignInDto{
			"test@test.test", "test",
		})
		if err != nil {
			log.Panicf("%v", err)
		}
		q, status = ClientE("/sign-in", "POST", b)
		assert.Contains(t, q, "token")
		assert.Equal(t, status, "200 OK")

		type TypeToken struct {
			Token string `json:"token"`
		}
		var bToken TypeToken
		err1 := json.Unmarshal([]byte(q), &bToken)
		if err1 != nil {
			log.Panicf("%v", err1)
		}
		token := bToken.Token

		// my-info
		b, err = json.Marshal(members.SignInDto{
			"test@test.test", "test",
		})
		if err != nil {
			log.Panicf("%v", err)
		}
		q, status = ClientToken("/my/info", "GET", b, token)
		assert.Contains(t, q, "result")
		assert.Contains(t, q, "id")
		assert.Contains(t, q, "exp")
		assert.Equal(t, status, "200 OK")

		// logout
		b, err = json.Marshal(members.SignInDto{
			"test@test.test", "test",
		})
		if err != nil {
			log.Panicf("%v", err)
		}
		q, status = ClientToken("/logout", "POST", b, token)
		assert.Contains(t, q, "success")
		assert.Equal(t, status, "200 OK")

		// my-info
		b, err = json.Marshal(members.SignInDto{
			"test@test.test", "test",
		})
		if err != nil {
			log.Panicf("%v", err)
		}
		q, status = ClientToken("/my/info", "GET", b, token)
		assert.Contains(t, q, "error")
		assert.Contains(t, q, "token in black list")
		assert.Contains(t, q, "message")
		assert.Contains(t, q, "fail")
		assert.Equal(t, status, "500 Internal Server Error")

	})

}
