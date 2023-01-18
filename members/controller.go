package members

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"msa-auth/util/api"
	"msa-auth/util/jwt"
	"strconv"
)

type Controller struct {
	service MemberService
}

func NewHandler(memberService MemberService) Controller {
	return Controller{memberService}
}

func (r Controller) SignOn(c *gin.Context) {
	dto, err1 := getSignOnDto(c.Request.Body)
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}

	id, err2 := r.service.SignOn(*dto)
	if err2 != nil {
		c.JSON(api.ServerErrorWithError(err2))
		return
	}

	c.JSON(api.OkWithMessage(strconv.FormatInt(id, 10)))
	return
}

func getSignOnDto(closer io.ReadCloser) (*SignOnDto, error) {
	var dto SignOnDto

	return getDto(closer, dto)
}

func (r Controller) SignIn(c *gin.Context) {
	dto, err1 := getSignInDto(c.Request.Body)
	if err1 != nil {
		c.JSON(api.ServerError())
		return
	}

	tkn, err2 := r.service.GetTokenBySignIn(*dto)
	if err2 != nil {
		c.JSON(api.ServerError())
		return
	}
	c.Header("token", tkn)
	c.JSON(api.OkWithToken(tkn))
	return
}

func (r Controller) Logout(c *gin.Context) {
	hToken := c.Request.Header["Token"]
	if len(hToken) == 0 {
		c.JSON(api.ServerError())
		return
	}

	err1 := r.service.Logout(hToken[0])
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}
	c.JSON(api.Ok())
	return
}

func (r Controller) getClaims(c *gin.Context) (*jwt.AuthTokenClaims, error) {
	hToken := c.Request.Header["Token"]
	if len(hToken) == 0 {
		return nil, errors.New("not exists token")
	}
	tkn := hToken[0]

	err1 := r.service.ValidToken(tkn)
	if err1 != nil {
		return nil, err1
	}

	return jwt.GetClaimsByTokenString(tkn)
}

func getSignInDto(closer io.ReadCloser) (*SignInDto, error) {
	var dto SignInDto

	return getDto(closer, dto)
}

func getDto[T dtoType](closer io.ReadCloser, dto T) (*T, error) {
	jsonData, err1 := io.ReadAll(closer)
	if err1 != nil {
		return nil, err1
	}

	err2 := json.Unmarshal(jsonData, &dto)
	if err2 != nil {
		return nil, err2
	}

	return getDto(closer, dto)
}

func (r Controller) MyInfo(c *gin.Context) {
	m, err1 := r.getClaims(c)
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}
	if err1 != nil {
		c.JSON(api.ServerError())
		return
	}
	c.JSON(api.OkWithObject(m))
	return
}

func (r Controller) SearchUser(c *gin.Context) {
	member := c.Query("member")
	id, err1 := strconv.ParseInt(member, 10, 64)
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}

	dto, err2 := r.service.FindMember(id)
	if err2 != nil {
		c.JSON(api.ServerErrorWithError(err2))
		return
	}

	c.JSON(api.OkWithObject(dto))
	return
}

func (r Controller) SearchByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(api.ServerError())
		return
	}

	dto, err1 := r.service.FindByEmail(email)
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}

	c.JSON(api.OkWithObject(dto))
	return
}

func (r Controller) MyInfoUpdate(c *gin.Context) {
	m, err1 := r.getClaims(c)
	if err1 != nil {
		c.JSON(api.ServerErrorWithError(err1))
		return
	}
	dto, err2 := getUpdateMyInfoDto(c.Request.Body)
	if err2 != nil {
		c.JSON(api.ServerErrorWithError(err2))
		return
	}

	err3 := r.service.UpdateMyInfo(m.UserID, *dto)
	if err3 != nil {
		c.JSON(api.ServerErrorWithError(err3))
		return
	}

	c.JSON(api.Ok())
	return
}

func getUpdateMyInfoDto(closer io.ReadCloser) (*UpdateMyInfoDto, error) {
	var dto UpdateMyInfoDto

	return getDto(closer, dto)
}
