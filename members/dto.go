package members

type SignOnDto struct {
	Email    string
	Password string

	Name     string
	NickName string
	Call     string
}

type SignInDto struct {
	Email    string
	Password string
}

type FindDto struct {
	Id    int64
	Email string

	Name     string
	NickName string
	Call     string
}
