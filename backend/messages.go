package main

type message struct {
	User  string `json: "user"`
	Msg   string `json: "msg"`
	Color string `json: "color"`
	Time  string `json: "time"`
}
