package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user_id struct { //models
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

var db *gorm.DB

func main() {
	//config
	connectionString := "root:12345@tcp(localhost:3306)/alta_user?parseTime=true"
	var err error
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&user_id{})
	fmt.Println(db)

	e := echo.New()
	e.GET("/pengguna", GetUsers) //router
	e.GET("/pengguna/:id", GetOneUser)
	e.POST("/pengguna", CreateUsers)
	e.PUT("/pengguna/:id", EditUser)
	e.DELETE("/pengguna/:id", DeleteUser)
	e.Start(":8080")
}

func GetUsers(c echo.Context) error { //controller
	users := []user_id{} //database
	if tx := db.Find(&users); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUsers(c echo.Context) error {
	users := user_id{}
	c.Bind(&users)
	if tx := db.Save(&users); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func GetOneUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var users user_id
	if tx := db.Find(&users, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func EditUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var users user_id
	if tx := db.Find(&users, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	c.Bind(&users)
	if tx := db.Save(&users); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var users user_id
	if tx := db.Find(&users, "id=?", id); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	if tx := db.Delete(&users); tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}
