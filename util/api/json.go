package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok() (int, gin.H) {
	return http.StatusOK, gin.H{
		"message": "success",
	}
}

func OkWithMessage(message string) (int, gin.H) {
	return http.StatusOK, gin.H{
		"message": message,
	}
}

func OkWithId(id int64) (int, gin.H) {
	return http.StatusOK, gin.H{
		"id": id,
	}
}

func OkWithStringList(list []string) (int, gin.H) {
	return http.StatusOK, gin.H{
		"results": list,
	}
}

func OkWithObject(obj interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{
		"result": obj,
	}
}

func OkWithEmail(email string) (int, gin.H) {
	return http.StatusOK, gin.H{
		"email": email,
	}
}

func OkWithToken(token string) (int, gin.H) {
	return http.StatusOK, gin.H{
		"token": token,
	}
}

func ServerError() (int, gin.H) {
	return http.StatusInternalServerError, gin.H{
		"message": "fail",
	}
}

func ServerErrorWithError(e error) (int, gin.H) {
	return http.StatusInternalServerError, gin.H{
		"message": "fail",
		"error":   e.Error(),
	}
}
