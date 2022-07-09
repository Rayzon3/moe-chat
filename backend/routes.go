package main

func intializeRoutes() {

	r.GET("/", loadLogin)
	r.POST("/auth", auth)
	r.Static("/css", "frontend/css")

	uRoutes := r.Group("/u")
	uRoutes.Use(isLogin)

	{
		uRoutes.GET("/moe", chatPage)
		uRoutes.POST("/moe", postMsg)

		uRoutes.GET("/msglist", jsonMsg)
		uRoutes.Static("/ts", "frontend/src")
		uRoutes.Static("/css", "frontend/css")
	}
}
