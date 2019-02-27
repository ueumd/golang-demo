package handler

import (
	"github.com/gin-gonic/gin"
	"myapiserver/api/model"
			"myapiserver/pkg/token"
	. "myapiserver/pkg/result"
	"myapiserver/pkg/errno"
	"fmt"
	"net/http"
	"strconv"
	"errors"
	"log"
)

func Login(c *gin.Context)  {
	var u model.User

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Println(u.Username, u.Password)

	d, err := u.GetUser(u.Username, u.Password)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	t, err := token.CreateToken(token.CustomClaims{Id: d.ID, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, errno.OK, model.Token{Token: t})
}

func LoginBind(c *gin.Context)  {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {

	} else {
		log.Println(u.Username)
		log.Println(u.Password)
		log.Println(u.ID)
	}

	list := map[string]interface{} {"username": u.Username, "password": u.Password, "id": u.ID}

	SendResponse(c, errors.New("ok"), list)
}

func LoginBind2(c *gin.Context)  {
	var u model.User
	if err := c.Bind(&u); err != nil {

	} else {
		log.Println(u.ID, u.Username, u.Password)
	}

	list := map[string]interface{} {"username": u.Username, "password": u.Password, "id": u.ID}

	SendResponse(c, errors.New("ok"), list)
}


/**
表单和Body参数（Multipart/Urlencoded Form）
典型的如 POST 提交的数据，
multipart/form-data
application/x-www-form-urlencoded
格式，都可以使用 c.PostForm获取到参数
 */
func LoginTest(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username, password)

	list := make(map[string]string)
	list["username"] = username
	list["password"] = password

	SendResponse(c, errors.New("ok"), list)
}

func LoginGet(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	love := c.QueryArray("love")

	fmt.Println(username, password, love)

	list := map[string]interface{}{"username": username, "password": password, "love": love}

	SendResponse(c, errors.New("ok"), list)

}

//列表数据
func GetUser(c *gin.Context)  {
	var user model.User

	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	result, err := user.Users()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"message":  "抱歉未找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func GetAllUser(c *gin.Context)  {
	var user model.User

	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")

	result, err := user.Users()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"message":  "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//添加数据
func AddUser(c *gin.Context)  {
	var user model.User
	//user.Username = c.Request.FormValue("username")
	//user.Password = c.Request.FormValue("password")

	if err := c.Bind(&user); err != nil {
	//	SendResponse(c, errno.ErrBind, nil)
		return
	}


	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "添加成功",
		"data": id,
	})
}

//修改数据

func Update(c *gin.Context)  {
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "修改成功",
	})
}


//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
}
