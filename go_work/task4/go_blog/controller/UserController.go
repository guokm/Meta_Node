package controller

import (
	"github/metanode/go_blog/config"
	"github/metanode/go_blog/middleware"
	"github/metanode/go_blog/model/dto"
	"github/metanode/go_blog/model/entity"
	"github/metanode/go_blog/model/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var regUser dto.RegUser
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}
	// 检查用户是否已存在
	var checkUser entity.User
	if err := config.DB.Where("username = ?", regUser.User).First(&checkUser); err == nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "User already exists"))
		return
	}
	// 检查邮箱是否已存在
	if err := config.DB.Where("email = ?", regUser.Email).First(&checkUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "Email already exists"))
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "Failed to hash password"))
		return
	}
	regUser.Password = string(hashedPassword)

	user := entity.User{
		Username: regUser.User,
		Password: regUser.Password,
		Email:    regUser.Email,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// 返回注册成功的用户信息
	c.JSON(http.StatusOK, vo.Success(user))
}

func Login(c *gin.Context) {
	var user dto.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}

	var checkUser entity.User
	if err := config.DB.Where("username = ?", user.User).Take(&checkUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "查询用户失败"))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "用户名或密码错误"))
		return
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&middleware.CustomClaims{
			ID:       checkUser.ID,
			Username: checkUser.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * config.GlobalConfig.Jwt.TokenExpire)),
				Issuer:    config.GlobalConfig.Jwt.Issuer,
			}},
	)

	tokenString, err := token.SignedString([]byte(config.GlobalConfig.Jwt.Secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, vo.Success(gin.H{"token": tokenString}))
}
