package members

import (
	"msa-auth/util"
)

type Members struct {
	util.Model
	Email    string
	Password string

	Name     string
	NickName string
	Call     string
}
