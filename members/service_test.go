package members

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type mockCommand struct{}

func (m mockCommand) Update(members Members) error {
	return nil
}

func (m mockCommand) Create(Members) (int64, error) {
	return 1, nil
}

type mockQuery struct{}

func (q mockQuery) FindById(id int64) (m *Members, e error) {
	return &Members{}, nil
}

func (q mockQuery) CountByEmail(string) (int64, error) {
	return 0, nil
}
func (q mockQuery) FindByEmail(_ string) (m *Members, e error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 10)

	return &Members{
		Email:    "test@email.com",
		Password: string(hashedPassword),
	}, nil
}

func TestMemberService_SignOn(t *testing.T) {

	type fields struct {
		command Command
		query   Query
	}
	type args struct {
		dto SignOnDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			"success",
			fields{
				command: mockCommand{},
				query:   mockQuery{},
			},
			args{dto: SignOnDto{
				Email:    "asketeddy@gmail.com",
				Password: "test1234",
				Name:     "tester",
				NickName: "nick",
				Call:     "010-0000-0000",
			}},
			1,
			false,
		},
		{
			"fail, empty password",
			fields{
				command: mockCommand{},
			},
			args{dto: SignOnDto{
				Email:    "asketeddy@gmail.com",
				Password: "",
				Name:     "",
				NickName: "",
				Call:     "",
			}},
			0,
			true,
		},
		{
			"fail, empty email",
			fields{
				command: mockCommand{},
			},
			args{dto: SignOnDto{
				Email:    "",
				Password: "test1234",
				Name:     "",
				NickName: "",
				Call:     "",
			}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MemberServiceImpl{
				command: tt.fields.command,
				query:   tt.fields.query,
			}
			got, err := s.SignOn(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignOn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_signOnValid(t *testing.T) {
	type args struct {
		dto SignOnDto
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"성공: 이메일과 비밀번호 값이 공백이 아닌 경우 nil 을 반환합니다.",
			args{SignOnDto{
				Email:    "asketeddy@gmail.com",
				Password: "test1234",
				Name:     "test user name",
				NickName: "tester",
				Call:     "010-0000-0000",
			}},
			false,
		},
		{
			"실패: 입력 모두가 공백이라면 error 를 반환합니다.",
			args{SignOnDto{
				Email:    "",
				Password: "",

				Name:     "",
				NickName: "",
				Call:     "",
			}},
			true,
		},
		{
			"실패: 비밀번호 입력만 공백이라면 error 를 반환합니다.",
			args{SignOnDto{
				Email:    "asketeddy@gmail.com",
				Password: "",
				Name:     "",
				NickName: "",
				Call:     "",
			}},
			true,
		},
		{
			"실패: 이메일 입력만 공백이라면 error 를 반환합니다.",
			args{SignOnDto{
				Email:    "",
				Password: "test1234",
				Name:     "",
				NickName: "",
				Call:     "",
			}},
			true,
		},
		{
			"failed, name is empty'",
			args{SignOnDto{
				Email:    "test",
				Password: "test1234",
				Name:     "",
				NickName: "test",
				Call:     "test",
			}},
			true,
		},
		{
			"failed, nick-name is empty'",
			args{SignOnDto{
				Email:    "test",
				Password: "test1234",
				Name:     "test",
				NickName: "",
				Call:     "test",
			}},
			true,
		},
		{
			"failed, call is empty'",
			args{SignOnDto{
				Email:    "test",
				Password: "test1234",
				Name:     "test",
				NickName: "test",
				Call:     "",
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := signOnValid(tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("signOnValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockQueryFailed struct{}

func (q mockQueryFailed) FindById(id int64) (m *Members, e error) {
	return &Members{}, nil
}

func (q mockQueryFailed) CountByEmail(email string) (int64, error) {
	if email == "admin@test.test" {
		return 0, errors.New("CountByEmail error")
	}
	return 100, nil
}
func (q mockQueryFailed) FindByEmail(_ string) (m *Members, e error) {
	return &Members{}, nil
}

func TestMemberServiceImpl_SignOn_Failed(t *testing.T) {
	type fields struct {
		command Command
		query   Query
		rds     R
	}
	type args struct {
		dto SignOnDto
	}

	var bigstring = "aaaaaaaaaaaaaaaaaadsnjfnasdjfnsadfjnsdafjnsadfjdsnfjdsnfjdsfnjadsfnsdjfndsjfnsdjfn"

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"fail, count > 0",
			fields{
				command: mockCommand{},
				query:   mockQueryFailed{},
				rds:     nil,
			},
			args{dto: SignOnDto{
				Email:    "asketeddy@gmail.com",
				Password: "test1234",
				Name:     "tester",
				NickName: "nick",
				Call:     "010-0000-0000",
			}},
			0,
			true,
		},
		{"fail, query error",
			fields{
				command: mockCommand{},
				query:   mockQueryFailed{},
				rds:     nil,
			},
			args{dto: SignOnDto{
				Email:    "admin@test.test",
				Password: "test1234",
				Name:     "tester",
				NickName: "nick",
				Call:     "010-0000-0000",
			}},
			0,
			true,
		},
		{"fail, len(password) > 72",
			fields{
				command: mockCommand{},
				query:   mockQuery{},
				rds:     nil,
			},
			args{dto: SignOnDto{
				Email:    "admin@test.test",
				Password: bigstring,
				Name:     "tester",
				NickName: "nick",
				Call:     "010-0000-0000",
			}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MemberServiceImpl{
				command: tt.fields.command,
				query:   tt.fields.query,
				rds:     tt.fields.rds,
			}
			got, err := s.SignOn(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignOn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemberServiceImpl_GetTokenBySignIn(t *testing.T) {

	type fields struct {
		command Command
		query   Query
		rds     R
	}
	type args struct {
		dto SignInDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"success", fields{
			command: mockCommand{},
			query:   mockQuery{},
			rds:     nil,
		},
			args{SignInDto{
				Email:    "test@email.com",
				Password: "test",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MemberServiceImpl{
				command: tt.fields.command,
				query:   tt.fields.query,
				rds:     tt.fields.rds,
			}
			got, err := s.GetTokenBySignIn(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenBySignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GetTokenBySignIn() return is empty")
			}
		})
	}
}
