package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// func gormConnect() *gorm.DB {
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=username dbname=gin_rest_api password=pass sslmode=disable")

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println("db connected: ", &db)
// 	return db
// }

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", func(c *gin.Context) {
		// セッションの作成
		session := sessions.Default(c)
		session.Set("loginUser", c.PostForm("userId"))
		session.Save()
		c.String(http.StatusOK, "ログイン完了")
	})

	r.GET("/logout", func(c *gin.Context) {
		// セッションの破棄
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.String(http.StatusOK, "ログアウトしました")
	})

	r.Run(":8080")
}
