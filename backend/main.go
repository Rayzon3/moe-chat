package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("postgres", "postgresql://username:password@localhost:5432/database")
	if err != nil {
		fmt.Println("Error while connecting the to database!!")
	}

	store, err := postgres.NewStore(db, []byte("secret"))
	if err != nil {
		fmt.Println("Error while creating store!!")
	}

	r.Use(sessions.Sessions("mySession", store))

	r.GET("/incr", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		ctx.JSON(200, gin.H{"count": count})

		r.Run(":8080")
	})
}
