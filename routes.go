// routes.go

package main

func initializeRoutes() {

    // Handle the route:    /
    router.GET("/", showIndexPage)

    // Handle GET requests at /article/view/some_article_id
    router.GET("/article/view/:article_id", getArticle)

}



