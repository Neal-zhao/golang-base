package blog

import (
	"blogger/controller"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/",controller.IndexHandler)
	r.Run()
}