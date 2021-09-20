package userApp

import "github.com/hari/bookstore_users_api/src/controllers/control_users"

func mapUrls() {
	router.POST("/users/", control_users.SaveUser)
	router.GET("/users/:user_id", control_users.GetUserByID)
	router.PUT("/users/:user_id", control_users.UpdateUser)
	router.PATCH("/users/:user_id", control_users.UpdateUser)
	router.DELETE("/users/:user_id", control_users.DeleteUser)
	router.GET("/internal/users/search", control_users.SearchUsers)
	router.POST("/users/login", control_users.LoginUser)
}
