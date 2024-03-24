package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"server/common"
	"server/middleware"
	ginitem "server/modules/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(middleware.Recovery())
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.POST("", ginitem.CreateItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.GET("/ping", func(context *gin.Context) {
		go func() {
			defer common.Recovery()

			fmt.Println([]int{}[0])
		}()

		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	_ = r.Run()
}
