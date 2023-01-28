package members

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"msa-auth/util/jwt"
	"time"
)

type MemberService interface {
	SignOn(dto SignOnDto) (int64, error)
	GetTokenBySignIn(dto SignInDto) (string, error)
	FindMember(id int64) (FindDto, error)
	FindByEmail(email string) (FindDto, error)
	FindByAll() ([]FindDto, error)

	ValidToken(token string) error
	Logout(token string) error

	UpdateMyInfo(id int64, dto UpdateMyInfoDto) error
}

type MemberServiceImpl struct {
	command Command
	query   Query
	rds     R
}

func NewService(command Command, query Query, r R) MemberService {
	return MemberServiceImpl{command, query, r}
}

func (s MemberServiceImpl) SignOn(dto SignOnDto) (int64, error) {
	var err = signOnValid(dto)
	if err != nil {
		return 0, err
	}

	count, err1 := s.query.CountByEmail(dto.Email)
	if count > 0 {
		return 0, errors.New("already email")
	}
	if err1 != nil {
		return 0, err1
	}

	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err2 != nil {
		return 0, err2
	}
	return s.command.Create(Members{
		Email:    dto.Email,
		Password: string(hashedPassword),
		Name:     dto.Name,
		NickName: dto.NickName,
		Call:     dto.Call,
	})
}

func signOnValid(dto SignOnDto) error {
	if dto.Email == "" {
		return errors.New("not valid password")
	}
	if dto.Password == "" {
		return errors.New("not valid email")
	}
	if dto.Name == "" {
		return errors.New("not valid name")
	}
	if dto.NickName == "" {
		return errors.New("not valid nick-name")
	}
	if dto.Call == "" {
		return errors.New("not valid call")
	}
	return nil
}

func (s MemberServiceImpl) GetTokenBySignIn(dto SignInDto) (string, error) {
	m, err1 := s.query.FindByEmail(dto.Email)
	if err1 != nil {
		return "", err1
	}
	if err2 := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(dto.Password)); err2 != nil {
		return "", err2
	}
	return jwt.CreateToken(m.ID, m.Email)
}

func (s MemberServiceImpl) ValidToken(token string) error {
	return s.rds.Valid(token)
}

func (s MemberServiceImpl) Logout(token string) error {
	return s.rds.Logout(token)
}

func (s MemberServiceImpl) FindMember(id int64) (FindDto, error) {
	m, err1 := s.query.FindById(id)
	if err1 != nil {
		return FindDto{}, err1
	}

	return FindDto{
		Id:       m.ID,
		Email:    m.Email,
		Name:     m.Name,
		NickName: m.NickName,
		Call:     m.Call,
	}, nil
}

func (s MemberServiceImpl) FindByAll() ([]FindDto, error) {
	m, err1 := s.query.FindAll()
	r := make([]FindDto, len(*m))

	for i, member := range *m {
		r[i] = FindDto{
			Id:       member.ID,
			Email:    member.Email,
			Name:     member.Name,
			NickName: member.NickName,
			Call:     member.Call,
		}
	}

	return r, err1
}

func (s MemberServiceImpl) FindByEmail(email string) (FindDto, error) {
	m, err1 := s.query.FindByEmail(email)
	if err1 != nil {
		return FindDto{}, err1
	}

	return FindDto{
		Id:       m.ID,
		Email:    m.Email,
		Name:     m.Name,
		NickName: m.NickName,
		Call:     m.Call,
	}, nil
}

func (s MemberServiceImpl) UpdateMyInfo(id int64, dto UpdateMyInfoDto) error {
	member, err1 := s.query.FindById(id)
	if err1 != nil {
		return err1
	}

	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err2 != nil {
		return err2
	}

	member.Password = string(hashedPassword)
	member.Name = dto.Name
	member.NickName = dto.NickName
	member.Call = dto.Call
	member.UpdatedAt = time.Now()

	err3 := s.command.Update(*member)
	if err3 != nil {
		return err3
	}
	return nil
}
