package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var r *gin.Engine
var db *sql.DB

func main() {

	url := "postgres://auqvywneajawrw:dadf8a14f2f6489f8c9289784041d4548812299b2afcab7f14a591aa78fe4180@ec2-34-239-241-121.compute-1.amazonaws.com:5432/d1lum6mbittgat"

	var err error
	db, err = connectDB(url)

	if err != nil {
		log.Fatalf("Error connecting database: %s", err.Error())
	}

	port := os.Getenv("PORT")

	fmt.Println("PORT = " + port)

	msglist, err := getAllMsgsDB(db)

	if err != nil {
		log.Fatalln("Unable to retrieve messages from database.")
	}

	globalmsgList = append(globalmsgList, msglist...)

	r = gin.Default()
	r.LoadHTMLGlob("static/template/*")

	r.Use(sessions.Sessions("chatsession", sessions.NewCookieStore([]byte("secret"))))

	intializeRoutes()

	r.Run()
}
