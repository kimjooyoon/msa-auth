package members

import "github.com/gin-gonic/gin"

func (r Controller) Mapping(server *gin.Engine) {
	server.POST("sign-on", r.SignOn)
	server.POST("sign-in", r.SignIn)
	server.GET("my/info", r.MyInfo)
	server.GET("members", r.SearchUser)
	server.GET("emails", r.SearchByEmail)
}
