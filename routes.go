package main

func intializeRoutes() {

	r.GET("/", loadLogin)
	r.POST("/auth", auth)
	r.Static("/css", "static/css")

	uRoutes := r.Group("/u")
	uRoutes.Use(isLogin)

	{
		uRoutes.GET("/chat", chatpage)
		uRoutes.POST("/chat", postmsg)

		uRoutes.GET("/msglist", jsonmsg)
		uRoutes.Static("/src", "static/src")
		uRoutes.Static("/css", "static/css")
	}
}
