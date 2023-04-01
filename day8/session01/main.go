package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func vipHandler(c *gin.Context)  {

}
// sessionId 处理
func sessionMiddleware(mgr Mgr) gin.HandlerFunc  {
	return func(c *gin.Context) {
		//get sessionid
		//var sessionId string
		var sd SessionData
		sessionId,ok := c.Cookie(sessionCookieName)
		//取不到 新生成
		if ok != nil {
			//sd = SessionNew("redis")
			sd = mgr.CreateSessionData()
		}
		sd,ok2 := mgr.GetSessionData(sessionId)
		if !ok2 {
			//sd := SessionNew("redis")
			sd = mgr.CreateSessionData()
			sessionId = sd.GetID()
		}
		c.Set(sessionCookieName,sd)
		//回写cookie要在请求函数返回之前 Next之前
		c.SetCookie(sessionCookieName,sessionId,600,"/","127.0.0.1",false,true)
		c.Next()
	}
}
//鉴权
func AuthMiddleWare(c *gin.Context)  {
	tmpSd,_ := c.Get(sessionCookieName)
	sd := tmpSd.(SessionData)
	val,err := sd.Get("isLogin")
	if err != nil {
		c.Redirect(http.StatusNotFound,"/login")
		return
	}
	isLogin,ok := val.(bool)//类型断言
	if !ok {
		c.Redirect(http.StatusNotFound,"/login")
		return
	}
	if !isLogin {
		c.Redirect(http.StatusNotFound,"/login")
		return
	}
	c.Next()
}
func main()  {
	r := gin.Default()
	InitMgr("redis","127.0.0.1","")
	r.Use(sessionMiddleware(MgrObj))
	r.GET("/vip",AuthMiddleWare,vipHandler)
	r.NoRoute(func(c *gin.Context) {
		c.XML(404,nil)
	})
	r.Run()
}
