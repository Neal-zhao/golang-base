package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name string `db:"name" form:"name"`
	Password string `db:"password" form:"password"`
}
func vipHandler(c *gin.Context)  {
	v,ok := c.Get("cookieKey")
	if !ok {
		c.Redirect(302,"/index")
		return
	}
	fmt.Println(v)
	c.HTML(http.StatusOK,"vip.html",gin.H{
		"userName":v,
	})
}
func indexHandler(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}
func loginHandler(c *gin.Context)  {
	var UserInfo UserInfo
	c.ShouldBind(&UserInfo)
	fmt.Println(UserInfo)

	c.SetCookie("name",UserInfo.Name,600,"/","127.0.0.1",false,true)
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}
func middleDeal(c *gin.Context)  {
	val,err := c.Cookie("name")
	if err != nil {
		c.Redirect(302,"/index")
		return
	}
	c.Set("cookieKey",val)
	fmt.Println(val)
	c.Next()

}
func main()  {
	r := gin.Default()
	r.Use()
	r.LoadHTMLGlob("template/*")
	r.GET("/index",indexHandler)
	r.GET("/vip",indexHandler)
	r.POST("/login",loginHandler)
	r.Run()
}
