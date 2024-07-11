package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/component/tokenprovider/jwt"
	"github.com/teddlethal/web-health-check/middleware"
	ginuser "github.com/teddlethal/web-health-check/modules/user/transport/gin"
	"github.com/teddlethal/web-health-check/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONNECTION")
	systemSecret := os.Getenv("SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	///////////////////////////////
	//authStore := storage.NewSqlStore(db)
	//middlewareAuth := middleware.RequireAuthen(authStore, tokenProvider)
	tokenProvider := jwt.NewTokenJwtProvider("jwt", systemSecret)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB Connection: ", db)

	r := gin.Default()
	r.Use(middleware.Recover())

	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginuser.Register(db))
		v1.POST("/login", ginuser.Login(db, tokenProvider))
	}
	routes.ConfigRoutes(v1, db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "web health check",
		})
	})
	errRun := r.Run(":2000")
	if errRun != nil {
		return
	}
}
