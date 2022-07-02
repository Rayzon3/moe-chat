package main

type message struct {
	User  string `json: "user"`
	Msg   string `json: "msg"`
	Color string `json: "color"`
	Time  string `json: "time"`
}

var globalMsgList = []message{{"Mod", "Welcome to Moe Chat UwU", "red", "[00:00:00]"}}

func getAllMsgs() []message {
	return globalMsgList
}
