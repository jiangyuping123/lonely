package routers

import (
	"colasoft.rd.go.backend/middleware/cors"
	"colasoft.rd.go.backend/middleware/myjwt"
	"colasoft.rd.go.backend/pkg/setting"
	v1 "colasoft.rd.go.backend/routers/api/v1"
	v2 "colasoft.rd.go.backend/routers/api/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(setting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())

	//////////////////////////////////////////////////////////////////////////
	var authMiddleware = myjwt.GinJWTMiddlewareInit(myjwt.AllUserAuthorizator)

	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), myjwt.NoRouteHandler)

	auth := r.Group("/auth")
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	api := r.Group("/user")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/info", v1.GetUserInfo)
		api.POST("/logout", v1.Logout)
	}

	//////////////////////////////////////////////////////////////////////////
	var adminMiddleware = myjwt.GinJWTMiddlewareInit(myjwt.AdminAuthorizator)
	apiv1 := r.Group("/api/v1")

	apiv1.Use(adminMiddleware.MiddlewareFunc())
	{
		apiv1.GET("/table/list", v2.GetArticles)
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	//////////////////////////////////////////////////////////////////////////
	var testMiddleware = myjwt.GinJWTMiddlewareInit(myjwt.TestAuthorizator)

	apiv2 := r.Group("/api/v2")
	apiv2.Use(testMiddleware.MiddlewareFunc())
	{
		apiv2.GET("/articles", v2.GetArticles)
	}

	return r
}
