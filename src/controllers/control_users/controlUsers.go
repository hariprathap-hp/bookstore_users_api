package control_users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hari/bookstore_oauth_go/oauth"
	"github.com/hari/bookstore_users_api/src/domains/domain_user"
	"github.com/hari/bookstore_users_api/src/log/zapper"
	"github.com/hari/bookstore_users_api/src/services/serviceUser"
	"github.com/hari/bookstore_users_api/src/utils/errors"
)

func getUserID(c *gin.Context) int64 {
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		fmt.Println("error getting user-id")
		return -1
	}
	return user_id
}

func GetUserByID(c *gin.Context) {
	id := getUserID(c)
	if err := oauth.AuthenticateAccessToken(c.Request); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	result, getErr := serviceUser.GetByID(id)
	if getErr != nil {
		c.JSON(int(getErr.Status), getErr)
	}

	if oauth.GetCallerId(c.Request) == result.Id {
		c.JSON(http.StatusOK, result.Marshal(false))
		return
	}
	c.JSON(http.StatusOK, result)
}

func SaveUser(c *gin.Context) {
	user := domain_user.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, saveErr := serviceUser.SaveUser(&user)
	if saveErr != nil {
		zapper.Error("saving of the user failed with",
			zapper.Field("error", saveErr.Error))
		c.JSON(int(saveErr.Status), saveErr)
		return
	}
	zapper.Info("user created is successful with",
		zapper.Field("user_id", result.Id),
		zapper.Field("firstname", result.FirstName))
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	user := domain_user.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user.Id = getUserID(c)
	isPatch := c.Request.Method == http.MethodPatch
	result, updateErr := serviceUser.UpdateUser(isPatch, &user)
	if updateErr != nil {
		c.JSON(int(updateErr.Status), updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	id := getUserID(c)
	_, getErr := serviceUser.DeleteByID(id)
	if getErr != nil {
		c.JSON(int(getErr.Status), getErr)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func SearchUsers(c *gin.Context) {
	status := c.Query("status")
	users, searchErr := serviceUser.Search(status)
	if searchErr != nil {
		c.JSON(int(searchErr.Status), searchErr)
		return
	}
	c.JSON(http.StatusOK, users)
}

func LoginUser(c *gin.Context) {
	fmt.Println("Login User Entered")
	request := domain_user.LoginUser{}
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewHTTPInternalServerError("invalid json body")
		c.JSON(int(restErr.Status), restErr)
	}
	result, loginErr := serviceUser.LoginUser(&request)
	if loginErr != nil {
		c.JSON(int(loginErr.Status), loginErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
