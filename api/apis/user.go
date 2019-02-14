package apis

import (
	"github.com/gin-gonic/gin"
	"myapiserver/api/model"
	"net/http"
	"strconv"
	"myapiserver/pkg/token"
	. "myapiserver/pkg/core"
	"myapiserver/pkg/errno"
)

func Login(c *gin.Context)  {
	var u model.User

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := u.GetUser(u.Username)

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
