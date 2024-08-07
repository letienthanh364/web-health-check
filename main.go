package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/component/tokenprovider/jwt"
	"github.com/teddlethal/web-health-check/linkchecker"
	"github.com/teddlethal/web-health-check/middleware"
	storageuser "github.com/teddlethal/web-health-check/modules/user/storage"
	ginuser "github.com/teddlethal/web-health-check/modules/user/transport/gin"
	"github.com/teddlethal/web-health-check/modules/website/biz"
	"github.com/teddlethal/web-health-check/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
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

	// Config the link checker
	webConfigs := bizwebsite.FetchWebsites(db)
	alertEmail := "letienthanh364@gmail.com"
	checkInterval := 10 * time.Minute
	alertThreshold := 24 * time.Hour
	lc := linkchecker.NewLinkChecker(webConfigs, alertEmail, checkInterval, alertThreshold)

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.Recover())

	r.Use(middleware.Recover())

	v1 := r.Group("/v1")
	{
		v1.POST("/register", ginuser.Register(db))
		v1.POST("/login", ginuser.Login(db, tokenProvider))
	}
	routes.CustomerRoutes(v1, db, middlewareAuth)
	routes.WebsiteRoutes(v1, db, middlewareAuth, lc)
	routes.ContactRoutes(v1, db, middlewareAuth)
	routes.CheckerRoutes(v1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "web health check",
		})
	})

	//Set up the website business logic
	// Start the link checker
	lc.Start()

	// Ensure the cron job is stopped gracefully on program exit
	defer lc.Stop()

	errRun := r.Run(":2000")

	if errRun != nil {
		return
	}
}
