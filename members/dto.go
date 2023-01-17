package members

type dtoType interface {
	SignOnDto | SignInDto |
		FindDto | UpdateMyInfoDto
}
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

type UpdateMyInfoDto struct {
	Password string
	Name     string
	NickName string
	Call     string
}

type UpdatePasswordDto struct {
	Password string
}
