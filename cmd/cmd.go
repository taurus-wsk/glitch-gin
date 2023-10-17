package cmd

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Start() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	sqltext := "eckert:21fc0da1b5e38e32@tcp(mysql.sqlpub.com:3306)/docker_test?timeout=60s&parseTime=True&charset=utf8mb4,utf8"
	db, err := gorm.Open(mysql.Open(sqltext))
	if err != nil {
		log.Fatal(err)
	}

	DB, err := db.DB()
	DB.SetMaxOpenConns(256)
	DB.SetMaxIdleConns(8)
	DB.SetConnMaxLifetime(360 * time.Second)

	//db1, _ :=db.DB()
	//defer db1.Close()
	log.Print("hh")
	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)

		err := db.Create(&todo).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, todo)
	})

	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo

		err := db.Find(&todos).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo

		err := db.Where("id = ?", id).First(&todo).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.BindJSON(&todo)
		db.Save(&todo)

		c.JSON(http.StatusOK, todo)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var todo Todo

		err := db.Where("id = ?", id).Delete(&todo).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": "Todo deleted"})
	})

	// 监听指定端口
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
