package routers

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func Init()  {
	gin.SetMode("debug")
	router = gin.Default()
}

func Run() (err error) {
	router.GET("/test", GetMethod)
	router.POST("/test", PostMethod)
	router.DELETE("/test", DeleteMethod)
	router.PUT("/test", PutMethod)
	router.SetTrustedProxies(nil)
	return router.Run(":8080")
}

func GetMethod(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "this is a test response",
	})
}

func PutMethod(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "this is a test response",
	})
}

func DeleteMethod(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "this is a test response",
	})
}

func PostMethod(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "this is a test response",
	})
}



