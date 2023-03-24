package main

import (
	controller "contentShareManage/controllers"
	router "contentShareManage/router"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

//	@title			ContentShareManage Api
//	@version		1.0
//	@description	通过访问Api将数据保存在内存中。
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@schemes	http https
//	@host		localhost:8000
//	@BasePath	/

//	@securitydefinitions.apikey
//	@in		header
//	@name	Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	args := os.Args
	controller.PassCount = 1
	if len(args) > 1 {
		count, _ := strconv.Atoi(args[1])
		controller.PassCount = count
	}
	r := gin.Default()
	router.RegisterRouter(r)
	r.Run(":8000")
}
