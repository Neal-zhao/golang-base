package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func addBook(c *gin.Context)  {
	if c.Request.Method == "POST" {
		title := c.DefaultPostForm("title","无名")
		price := c.DefaultPostForm("price","0.1")
		price64,err := strconv.ParseFloat(price,64)
		if err != nil {
			c.JSON(http.StatusPartialContent,gin.H{
				"msg":"价格参数错误",
			})
		}
		err = insertBook(title,price64)
		if err != nil {
			//添加成功 跳转到 列表页
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently,"/book/list")
		return
	}
	c.HTML(http.StatusOK,"book/add_book.html",nil)
}
func deleteBook(c *gin.Context)  {
	id := c.Query("id")
	fmt.Println(id)
	id64,_ := strconv.ParseInt(id,10,10)
	fmt.Println(id64)
	deleteBook2(id64)
	c.Redirect(http.StatusMovedPermanently,"/book/list")
}
func bookList(c *gin.Context)  {
	allBook,err := QueryAllBook()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":err,
		})
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"code" :0,
	//	"msg" : "success",
	//	"data" :allBook,
	//})
	//return
	c.HTML(http.StatusOK,"book/book_list.html",gin.H{
		"data":allBook,
	})
}
func uploadFile(c *gin.Context)  {
	if c.Request.Method == "POST" {
		fileHeader,err := c.FormFile("ycy")
		fmt.Println(fileHeader.Filename)
		idstr := "ycy01.jpeg"
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":fmt.Sprintf("file err :%s",err),
			})
			return
		}
		err = c.SaveUploadedFile(fileHeader,"./image/"+idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":fmt.Sprintf("SaveUploadedFile err :%s",err),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"上传成功",
		})
		return
	}
	c.HTML(http.StatusOK,"book/upload_file.html",nil)
}
func middleFun(c *gin.Context)  {
	startTime := time.Now()
	time.Sleep(time.Millisecond * 5)
	duration := time.Since(startTime)
	fmt.Println("middle time fun：",duration.String())
}
func main()  {
	err:= initDB()
	if err != nil {
		panic(err)
	}
	//建立一个http serve
	//路由 方法
	router := gin.Default()
	router.Use(middleFun)
	v1 := router.Group("/file")
	{
		v1.POST("/upload",uploadFile)
		v1.GET("/upload",uploadFile)
	}

	//加载html文件
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/book/new",addBook)
	router.POST("/book/add",addBook)
	//router.DELETE("/book/delete",deleteBook)
	router.GET("/book/delete",deleteBook)
	router.GET("/book/list",bookList)
	router.Run(":9090")
}