package members

import "github.com/gin-gonic/gin"

func (r Controller) Mapping(server *gin.Engine) {
	server.POST("sign-on", r.SignOn)
	server.POST("sign-in", r.SignIn)
	server.POST("logout", r.Logout)
	server.GET("my/info", r.MyInfo)
	server.GET("members", r.SearchUser)
	server.GET("emails", r.SearchByEmail)
}
