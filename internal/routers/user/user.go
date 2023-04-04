package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"noir.github.com/aws_billing/internal/models"
	"noir.github.com/aws_billing/internal/modules/logger"
	"noir.github.com/aws_billing/internal/modules/utils"
	"strconv"
	"time"
)

const tokenDuration = 2 * time.Minute

func Remove(c *gin.Context) {}

func Disable(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)

	userModel := &models.User{}
	userModel.Disable(int(userID))

	jr := utils.JsonResponse{}
	c.JSON(200, jr.SuccessJson("操作成功", nil))
	return
}

func Enable(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)

	userModel := &models.User{}
	userModel.Enable(int(userID))

	jr := utils.JsonResponse{}
	c.JSON(200, jr.SuccessJson("操作成功", nil))
	return
}

func List(c *gin.Context) {
	jr := &utils.JsonResponse{}
	userModel := &models.User{}
	UserList, err := userModel.List()
	if err != nil {
		c.JSON(200, jr.CommonFailureJson("获取用户列表失败"))
		return
	}
	c.JSON(200, jr.SuccessJson("获取用户列表成功", UserList))
	return
}

func UpdatePassword(c *gin.Context) {}

func Forget(c *gin.Context) {}

func Registry(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	checkPassword := c.PostForm("checkPassword")
	// 验证
	jr := &utils.JsonResponse{}
	if password != checkPassword {
		c.JSON(200, jr.CommonFailureJson("密码不匹配"))
		return
	}

	userModel := &models.User{
		Username: username,
		Password: password,
	}

	_, err := userModel.Create()
	if err != nil {
		c.JSON(200, jr.CommonFailureJson(username+" 创建失败"))
		return
	}

	c.JSON(200, jr.SuccessJson("注册成功", map[string]interface{}{
		"username": username,
		"password": password,
	}))
	return
}

func CheckUserStatus(c *gin.Context) {
	username := c.PostForm("username")
	jr := &utils.JsonResponse{}
	userModel := &models.User{}
	userModel.Find(username)
	if userModel.Status != 1 {
		c.JSON(200, jr.CommonFailureJson("用户已被禁用，请联系管理员"))
		c.Abort()
	}
	c.Next()
}

func ValidateLogin(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 数据验证
	jr := &utils.JsonResponse{}
	if username == "" || password == "" {
		c.JSON(200, jr.CommonFailureJson("用户名、密码不能为空"))
		return
	}
	userModel := &models.User{}
	if !userModel.Match(username, password) {
		c.JSON(200, jr.CommonFailureJson("用户名或密码错误"))
		return
	}
	// 记录日志
	loginLogModel := &models.LoginLog{}
	loginLogModel.Username = userModel.Username
	loginLogModel.IP = c.RemoteIP()
	_, err := loginLogModel.Create()
	if err != nil {
		logger.Error("记录用户登录日志失败", err)
	}
	token, err := generateToken(userModel)
	if err != nil {
		c.JSON(200, jr.CommonFailureJson("生成jwt失败"))
		return
	}
	// 返回
	c.JSON(200, jr.SuccessJson("登录成功", map[string]interface{}{
		"token":    token,
		"uid":      userModel.ID,
		"username": userModel.Username,
		"is_admin": userModel.IsAdmin,
	}))
	return
}

// 生成jwt
func generateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(tokenDuration).Unix()
	claims["uid"] = user.ID
	claims["iat"] = time.Now().Unix()
	claims["username"] = user.Username
	claims["is_admin"] = user.IsAdmin
	token.Claims = claims

	// 加密密钥
	return token.SignedString([]byte("qtoagl159"))
}

// 还原jwt
func RestoreToken(c *gin.Context) {
	jr := utils.JsonResponse{}
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(200, jr.CommonFailureJson("auth不能为空"))
		c.Abort()
		return
	}
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("qtoagl159"), nil
	})
	if err != nil {
		c.JSON(200, jr.CommonFailureJson(err.Error()))
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(200, jr.CommonFailureJson("invalid claims"))
		c.Abort()
		return
	}
	c.Set("uid", int(claims["uid"].(float64)))
	c.Set("username", claims["username"])
	c.Set("is_admin", int(claims["is_admin"].(float64)))
	c.Next()
	return
}
