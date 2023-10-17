package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Start() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&Todo{})

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
	err = router.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {}
