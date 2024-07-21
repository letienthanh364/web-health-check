package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/component/tokenprovider/jwt"
	"github.com/teddlethal/web-health-check/middleware"
	storageuser "github.com/teddlethal/web-health-check/modules/user/storage"
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
	authStore := storageuser.NewSqlStore(db)
	tokenProvider := jwt.NewTokenJwtProvider("jwt", systemSecret)
	middlewareAuth := middleware.RequireAuthen(authStore, tokenProvider)

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
	routes.ConfigRoutes(v1, db, middlewareAuth)
	routes.CustomerRoutes(v1, db, middlewareAuth)
	routes.WebsiteRoutes(v1, db, middlewareAuth)
	routes.EmailRoutes(v1, db, middlewareAuth)
	routes.CheckerRoutes(v1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "web health check",
		})
	})

	// Set up the website business logic

	// Fetch the list of websites and start the link checker
	//configs := linkchecker.FetchWebsites(db)
	//lc := linkchecker.NewLinkChecker(configs)
	//lc.Start()
	//
	//// Ensure the cron job is stopped gracefully on program exit
	//defer lc.Stop()

	errRun := r.Run(":2000")

	if errRun != nil {
		return
	}
}
