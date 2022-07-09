package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	username = "user"
	ucolor   = "color"
)

func loadLogin(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"login.html",
		gin.H{},
	)
}

func auth(ctx *gin.Context) {
	user := ctx.PostForm("user")
	color := ctx.PostForm("color")

	if strings.Contains(user, "<") || strings.Contains(user, ">") || user == "" {
		ctx.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"error": "Use a valid username !!",
			},
		)
		return
	}

	session := sessions.Default(ctx)
	session.Set(username, user)
	session.Set(ucolor, color)

	if err := session.Save(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	} else {
		ctx.Redirect(http.StatusFound, "/moe/chat")
	}
}

func isLogin(ctx *gin.Context) {
	sessions := sessions.Default(ctx)
	user := sessions.Get(username)

	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}

	ctx.Next()
}

func chatPage(ctx *gin.Context) {

	session := sessions.Default(ctx)

	user := session.Get(username)
	color := session.Get(ucolor)
	msgList := getAllMsgs()

	ctx.HTML(
		http.StatusOK,
		"chat.html",
		gin.H{
			"user":    user,
			"color":   color,
			"msgList": msgList,
		},
	)

}

func postMsg(ctx *gin.Context) {

	session := sessions.Default(ctx)
	user := session.Get(username)
	color := session.Get(ucolor)
	time := time.Now()

	umessage := ctx.PostForm("usermessage")

	data, err := addMsgtoDB(db, fmt.Sprint(user), umessage, fmt.Sprint(color), time.Format("[15:04:05] "))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Unable to add message to database"})
	}

	addMsg(*data)

	ctx.Redirect(http.StatusFound, "/u/chat")

}

func jsonMsg(ctx *gin.Context) {
	ctx.IndentedJSON(
		http.StatusOK,
		globalMsgList,
	)
}
