package members

import (
	"encoding/json"
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
	jsonData, err1 := io.ReadAll(closer)
	if err1 != nil {
		return nil, err1
	}

	var dto SignOnDto
	err2 := json.Unmarshal(jsonData, &dto)
	if err2 != nil {
		return nil, err2
	}

	return &dto, nil
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

func getSignInDto(closer io.ReadCloser) (*SignInDto, error) {
	jsonData, err1 := io.ReadAll(closer)
	if err1 != nil {
		return nil, err1
	}

	var dto SignInDto
	err2 := json.Unmarshal(jsonData, &dto)
	if err2 != nil {
		return nil, err2
	}

	return &dto, nil
}

func (r Controller) MyInfo(c *gin.Context) {
	hToken := c.Request.Header["Token"]
	if len(hToken) == 0 {
		c.JSON(api.ServerError())
		return
	}
	tkn := hToken[0]
	m, err1 := jwt.GetClaimsByTokenString(tkn)

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
