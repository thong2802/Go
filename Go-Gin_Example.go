package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"path/filepath"
)
// gin-gonic handler,
//middleware,
//group,
//upload single,
//multiple file
// uuid
// path/filepath

func main() {
	router := gin.Default()

	router.Static("/static-file", "./assets")
	//router.Use(AcustomMiddleWare)
	router.Use(AcustomMiddleWare())

	router.GET("/ping", func(context *gin.Context) {
		//context.String(http.StatusOK, "Ping")
		//context.JSONP(http.StatusOK, gin.H{
		//	"name" : "Thong",
		//	"Age"  : "21",
		//})
		var data = map[string]interface{}{
			"Name" : "Thong",
			"Age" : 21,
			"company" : "Grab Singapore",
		}
		context.JSONP(http.StatusOK, data)
	})

	router.GET("/pingdetail:id", postPing)
	router.GET("/pingQuery", QueryPing)
	router.GET("/postform", PostFrom)


	//Group api
	// api/v1/ping
	// api/v1/pong

	// api/v2/a
	// api/v2/b
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// initialization middleWare for V1
			v1.Use(AcustomMiddleWareV1())
			v1.GET("/ping", func(context *gin.Context) {
				context.JSONP(http.StatusOK, gin.H{
						"message" : "Group api ping",
				})
			})
			v1.GET("/pong", func(context *gin.Context) {
				context.JSONP(http.StatusOK, gin.H{
					"message" : "Group api pong",
				})
			})
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/a", func(context *gin.Context) {
				context.JSONP(http.StatusOK, gin.H{
					"message" : "Group api a",
				})
			})
			v2.GET("/b", func(context *gin.Context) {
				context.JSONP(http.StatusOK, gin.H{
					"message" : "Group api b",
				})
			})
		}

	}

	//uploadFile

	//Single file
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./assets/upload/" + file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Multiple files
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload_mutiple_file", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			log.Println("hi")
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "./assets/upload/" + uuid.New().String() + filepath.Ext(file.Filename)) // uuid render ra name ngau nhien va dung filepath lay ra duoi file va tranh trung ten va bao mat file
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":2802")
}

	func PostFrom(context *gin.Context) {
		// get dataa form
		address := context.DefaultPostForm("address", "Viet Nam")
		context.JSONP(http.StatusOK, gin.H{
			"message" : "Hello " + address + " Post ping",
		})
	}

	func QueryPing(context *gin.Context) {
		//get data Query
		name := context.DefaultQuery("name", "Guest")
		var data = map[string]interface{}{
			"message": "Hello from " + name + " Get ping",
		}
		context.JSONP(http.StatusOK, data)
	}

	func postPing(context *gin.Context) {
		id := context.Param("id") // get Param
		var data = map[string]interface{}{
			"id" : id,
			"Name" : "Thong",
			"Age" : 21,
			"company" : "Grab Singapore",
		}
		context.JSONP(http.StatusOK, data)
	}


	// middleWare // sẽ xử lí trung gian, khi có một request về server, middleWare sẽ kiểm tra dữ liệu hợp lệ hay không sau đó chuyển về server,
	// Ngược lại trước khi response về lại client thì middleware cũng kiểm tra dữ liệu hợp lệ không rồi chuyển qua client.

	// MiddeleWare for v1
	func AcustomMiddleWareV1 () gin.HandlerFunc {
		return func(context *gin.Context) {
			log.Println("I'm in v1 middleware")
			context.Next()
		}
	}

	func  AcustomMiddleWare() gin.HandlerFunc {
		return func(context *gin.Context) {
			if true {
				log.Println("I'm in global middleware")
			}
			context.Next()
		}
	}

	//func AcustomMiddleWare( context *gin.Context){
	//	//jwt
	//	//auth
	//	log.Println("I'm in global middleware")
	//	context.Next()
	//}
