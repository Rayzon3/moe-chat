package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var r *gin.Engine
var db *sql.DB

func main() {
	url := "postgres://auqvywneajawrw:dadf8a14f2f6489f8c9289784041d4548812299b2afcab7f14a591aa78fe4180@ec2-34-239-241-121.compute-1.amazonaws.com:5432/d1lum6mbittgat"

	// if !ok {
	// 	log.Fatalln("Database url is required.")
	// }

	var err error
	db, err = connDB(url)

	if err != nil {
		log.Fatalf("Error connecting database: %s", err.Error())
	}

	// port := os.Getenv("PORT")

	msglist, err := getAllMsgsFromDB(db)

	if err != nil {
		fmt.Println("No messages in DB")
	}

	globalMsgList = append(globalMsgList, msglist...)

	r = gin.Default()
	r.LoadHTMLGlob("../frontend/template/*")

	r.Use(sessions.Sessions("chatsession", sessions.NewCookieStore([]byte("secret"))))

	intializeRoutes()

	r.Run()
}
