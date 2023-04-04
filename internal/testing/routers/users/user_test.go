package users

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"noir.github.com/aws_billing/internal/models"
	"noir.github.com/aws_billing/internal/modules/httpclient"
	"noir.github.com/aws_billing/internal/routers/user"
	"testing"
)

func TestRun(t *testing.T) {
	models.DB = models.CreateDB()
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		msg := user.ValidateLogin(c)
		var result *gin.H
		_ = json.Unmarshal([]byte(msg), result)
		fmt.Println(result)

		c.JSON(200, result)
	})
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
}

func TestLogin(t *testing.T) {

	resp := httpclient.PostParams("http://127.0.0.1:8888/login", "username=test@highcloud.cn&password=000000", 30)
	fmt.Println(resp)

}
