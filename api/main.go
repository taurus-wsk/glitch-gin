package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	app *gin.Engine
)

// CREATE ENDPOIND
type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB

func myRoute(r *gin.RouterGroup) {
	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		log.Print("hh3")
		err := DB.Find(&todos).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})
}

func init() {
	sqltext := "eckert:21fc0da1b5e38e32@tcp(mysql.sqlpub.com:3306)/docker_test?timeout=60s&parseTime=True&charset=utf8mb4,utf8"
	db, err := gorm.Open(mysql.Open(sqltext))
	if err != nil {
		log.Fatal(err)
	}

	DB, err := db.DB()
	DB.SetMaxOpenConns(256)
	DB.SetMaxIdleConns(8)
	DB.SetConnMaxLifetime(360 * time.Second)
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)

}

// ADD THIS SCRIPT
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
