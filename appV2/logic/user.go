package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"libraryManagementSystem/appV2/model"
	"libraryManagementSystem/appV2/tools"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GetCode godoc
//
// @Tags		public
// @Summary		登录验证码
// @Description	用户登录页获取验证码操作
// @Produce		json
// @Success 200 {object} tools.Response{data=map[string]string}
// @Failure 500 {object} tools.Response
// @Router			/getCode [GET]
func GetCode(ctx *gin.Context) {
	fileName := func() string {

		// 设置图片大小
		width, height := 100, 50
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		// 随机种子
		rand.Seed(time.Now().Unix())

		// 随机生成4位验证码
		code := fmt.Sprintf("%04d", rand.Intn(10000))
		fmt.Println("验证码:", code)
		//验证码存到redis

		var redisClient *redis.Client = model.RedisConn
		err := redisClient.Set(ctx, "captcha", code, 5*time.Minute).Err()
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusOK, tools.Response{
				Code:    tools.OK,
				Message: err.Error(),
				Data:    nil,
			})
			return ""
		}
		// 设置字体大小
		fontSize := 30

		// 设置字体颜色
		fontColor := color.RGBA{255, 0, 0, 255}

		// 设置背景颜色
		bgColor := color.RGBA{255, 255, 255, 255}

		// 绘制背景
		draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

		// 绘制验证码
		for i, c := range code {
			// 计算字体位置
			x := (width / 4) * i
			y := (height - fontSize) / 2

			// 绘制字体
			func(img *image.RGBA, s string, x, y int, c color.Color, size int) {
				f := basicfont.Face7x13
				d := &font.Drawer{
					Dst:  img,
					Src:  image.NewUniform(c),
					Face: f,
					Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
				}
				d.DrawString(s)
			}(img, string(c), x, y, fontColor, fontSize)

		}
		// 将图像写入文件
		file, err := os.Create("appV2/resource/static/img/captcha.png")
		fileUrlArr := strings.Split(file.Name(), "/")
		fileName := fileUrlArr[len(fileUrlArr)-1]
		if err != nil {
			panic(err)
			ctx.JSON(http.StatusInternalServerError, tools.Response{
				Code:    tools.InternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}
		defer file.Close()
		png.Encode(file, img)
		return fileName
	}()
	ctx.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "获取验证码",
		Data:    map[string]string{"imgName": fileName},
	})
	return
}

// UserLogin godoc
//
// @Summary		用户登录
// @Description	会执行用户登录操作
// @Tags		public
// @Accept		multipart/form-data
// @Produce		json
// @Param userName formData string true "用户名"
// @Param password formData string true "密码"
// @Param captcha formData string true "验证码"
// @Success 200 {object} tools.Response{data=Token}
// @Failed 406,500 {object} tools.Response
// @Router			/userLogin [POST]
func UserLogin(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.BadRequest,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return
	}
	//校验验证码
	formCode := context.PostForm("captcha")
	var redisClient *redis.Client = model.RedisConn
	redisCode, _ := redisClient.Get(context, "captcha").Result()
	if redisCode != formCode {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.CaptchaError,
			Message: "验证码错误",
		})
		return
	}
	//校验
	dbUser := model.UserCheck(user.UserName, user.Password)
	if dbUser.Id > 0 {
		context.SetCookie("id", strconv.FormatInt(dbUser.Id, 10), 3600, "/", "", false, true)
		a, r, errT := tools.Token.GetToken(dbUser.Id, dbUser.UserName)
		log.Printf("atoken:%s\n", a)
		log.Printf("rtoken:%s\n", r)
		if errT != nil {
			context.JSON(http.StatusInternalServerError, tools.Response{
				Code:    tools.InternalServerError,
				Message: "Token生成失败:" + errT.Error(),
				Data:    nil,
			})
			return
		}
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "登陆成功",
			Data: Token{
				AccessToken:  a,
				RefreshToken: r,
			},
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.UserInfoError,
		Message: "用户信息错误",
		Data:    nil,
	})
}

// GetUser godoc
//
// @Summary		用户查看信息
// @Description	获取用户信息
// @Tags		user/users
// @Produce		json
// @CookieParam id string true "用户id"
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} tools.Response{data=model.User}
// @Router			/user/users [GET]
func GetUser(context *gin.Context) {
	idString, _ := context.Cookie("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	dbUser := model.GetUser(id)
	dbUser.Password = ""
	if dbUser.Id > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "用户存在",
			Data:    dbUser,
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.UserIsNotExist,
		Message: "用户不存在",
	})
}

// GetUserById godoc
//
// @Summary		管理员获取用户信息
// @Description	管理员获取用户信息
// @Tags		admin/users
// @Produce		json
// @Param id path int64 true "用户id"
// @Success 200 {object} tools.Response{data=model.User}
// @Router			/admin/users/{id} [GET]
func GetUserById(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	dbUser := model.GetUser(id)
	dbUser.Password = ""
	if dbUser.Id > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "用户存在",
			Data:    dbUser,
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.UserIsNotExist,
		Message: "用户不存在",
	})
}

// SearchUser godoc
//
// @Summary		搜索用户
// @Description	搜索获取用户信息
// @Tags		admin/users
// @Produce		json
// @Param q  query string false "查询条件"
// @Success 200 {object} tools.Response{data=[]model.User{}}
// @Router			/admin/users [GET]
func SearchUser(context *gin.Context) {
	query := context.Query("q")
	dbUsers := model.SearchUser(query)
	for i, _ := range dbUsers {
		dbUsers[i].Password = ""
	}
	if len(dbUsers) > 0 {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "用户存在",
			Data:    dbUsers,
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.UserIsNotExist,
		Message: "用户不存在",
		Data:    nil,
	})
}

// AddUser godoc
//
// @Summary		新增一个用户
// @Description	用户注册或管理员添加用户
// @Tags		public
// @Accept		multipart/form-data
// @Produce		json
// @Param userName formData string true "用户名"
// @Param password formData string true "密码"
// @Param name formData string true "性名"
// @Param sex formData string true "性别"
// @Param phone formData string true "电话"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/users [POST]
func AddUser(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return
	}
	model.AddUser(user)
	context.JSON(http.StatusCreated, tools.Response{
		Code:    tools.OK,
		Message: "注册成功",
		Data:    nil,
	})
}

// UpdateUser godoc
//
// @Summary		用户修改信息
// @Description	用户修改自己的信息
// @Tags		user/users
// @Accept		multipart/form-data
// @Produce		json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param userName formData string true "用户名"
// @Param password formData string true "旧密码"
// @Param newPassword formData string true "新密码"
// @Param phone formData string true "电话"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/user/users [PUT]
func UpdateUser(context *gin.Context) {
	var updateUser model.User
	if err := context.ShouldBind(&updateUser); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return
	}
	userIdString, _ := context.Cookie("id")
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	user := model.GetUser(userId)
	updateUser.Id = userId
	//校验用户名密码，正确则更新用户
	//updateUser密码与user密码匹配，则将updateUser密码设置为newPassword，再更新
	if user.Password == updateUser.Password {
		updateUser.Password = context.PostForm("newPassword")
		model.UpdateUser(updateUser)
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "更新成功",
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.UserInfoError,
		Message: "输入的旧密码错误",
	})
}

// UpdateUserByAdmin godoc
//
// @Summary		更新用户信息
// @Description	管理员更新用户信息
// @Tags		admin/users
// @Accept		multipart/form-data
// @Produce		json
// @Param id path int64 true "用户id"
// @Param userName formData string true "用户名"
// @Param password formData string true "密码"
// @Param phone formData string true "电话"
// @Param status formData int true "用户帐号状态"
// @Success 200 {object} tools.Response
// @Failed 406,500 {object} tools.Response
// @Router			/admin/users/{id} [PUT]
func UpdateUserByAdmin(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	var updateUser model.User
	if err := context.ShouldBind(&updateUser); err != nil {
		context.JSON(http.StatusNotAcceptable, tools.Response{
			Code:    tools.UserInfoError,
			Message: "绑定失败" + err.Error(),
			Data:    nil,
		})
		return
	}
	updateUser.Id = id
	err := model.UpdateUser(updateUser)
	if err != nil {
		context.JSON(http.StatusOK, tools.Response{
			Code:    tools.OK,
			Message: "更新失败" + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "更新成功",
	})
}

// DeleteUser godoc
//
// @Summary		删除用户
// @Description 管理员通过id删除用户
// @Tags		admin/users
// @Produce		json
// @Param id path int64 true "用户id"
// @Success 200 {object} tools.Response
// @Router			/user/users/{id} [DELETE]
func DeleteUser(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	model.DeleteUser(id)
	context.JSON(http.StatusOK, tools.Response{
		Code:    tools.OK,
		Message: "删除成功",
	})
}
