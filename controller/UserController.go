package controller

import (
	"Gin/common"
	"Gin/dto"
	"Gin/model"
	"Gin/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(context *gin.Context) {
	// 获取参数

	DB := common.GetDB()

	username := context.PostForm("username")
	nickname := context.PostForm("nickname")
	password := context.PostForm("password")

	// 数据验证

	if len(nickname) > 32 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "昵称长度不能大于32位")
		return
	}
	if len(nickname) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "昵称不能为空")
		return
	}

	if len(password) > 32 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能大于32位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能小于6位")
		return
	}

	if len(username) > 32 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名长度不能大于32位")
		return
	}
	if len(username) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}

	log.Println(username, nickname, password)

	// 判断是否存在

	if isUserExist(DB, username) {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密过程出现错误")
	}

	newUser := model.User{
		Username: username,
		Nickname: nickname,
		Password: string(hashedPassword),
	}

	DB.Create(&newUser)

	// 返回结果
	response.Success(context, nil, "注册成功")
}

func Login(context *gin.Context) {

	DB := common.GetDB()

	// 获取参数

	username := context.PostForm("username")
	password := context.PostForm("password")

	// 数据验证

	if len(password) > 32 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能大于32位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码长度不能小于6位")
		return
	}

	if len(username) > 32 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名长度不能大于32位")
		return
	}
	if len(username) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}

	// 判断用户是否存在

	var user model.User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户名不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) ; err != nil {
		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Success(context, gin.H{"token": token}, "登陆成功!")

}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	response.Response(context, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func isUserExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
