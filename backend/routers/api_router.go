package routers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupAPIRouter sets up the API router
func SetupAPIRouter(api *gin.Engine, db *sql.DB) {
	api.Use(middleware.ApiMiddleware(db))

	// Setup Redis session store
	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}
	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		log.Fatal("Unable to create save session with redis session: ", storeError)
	}
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	api.Use(sessions.Sessions("itsmsession", store))

	// Set up middleware
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// Define API routes
	api.GET("/", handlers.IndexHandler)
	api.GET("/index", handlers.IndexHandler)
	api.GET("/login", handlers.LoginHandler)
	api.GET("/auth", handlers.AuthHandler)
	api.GET("/logout", handlers.LogoutHandler)
	api.GET("/login/admin", handlers.GetAgentHandler)

	///////////////////////////////////////////////GET////////////////////////////////////////////////

	// Define API router group for authorized routes
	authorized := api.Group("/api")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/tickets", handlers.ItsmHandler)
		authorized.GET("/assets", handlers.ItsmHandler)
		authorized.GET("/staff", handlers.ItsmHandler)
		authorized.GET("/agent", handlers.ItsmHandler)
		authorized.GET("/admin", handlers.ItsmHandler)
		authorized.GET("/procurement", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	ticket := api.Group("/api/ticket")
	ticket.Use(middleware.AuthorizeRequest())
	{
		ticket.GET("/status", handlers.ItsmHandler)
		ticket.GET("/assets", handlers.ItsmHandler)
		ticket.GET("/category", handlers.ItsmHandler)
		ticket.GET("/sub-category", handlers.ItsmHandler)
		ticket.GET("/staff", handlers.ItsmHandler)
		ticket.GET("/agent", handlers.ItsmHandler)
		ticket.GET("/admin", handlers.ItsmHandler)
		ticket.GET("/priority", handlers.ItsmHandler)
		ticket.GET("/sla", handlers.ItsmHandler)
		ticket.GET("/satisfaction", handlers.ItsmHandler)
		ticket.GET("/policy", handlers.ItsmHandler)
		ticket.GET("/related", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	agent := api.Group("/api/agent")
	agent.Use(middleware.AuthorizeRequest())
	{
		agent.GET("/role", handlers.ItsmHandler)
		agent.GET("/unit", handlers.ItsmHandler)
		agent.GET("/username", handlers.ItsmHandler)
		agent.GET("/supervisor", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	staff := api.Group("/api/staff")
	staff.Use(middleware.AuthorizeRequest())
	{
		staff.GET("/role", handlers.ItsmHandler)
		staff.GET("/unit", handlers.ItsmHandler)
		staff.GET("/username", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	asset := api.Group("/api/assets")
	asset.Use(middleware.AuthorizeRequest())
	{
		asset.GET("/creator", handlers.ItsmHandler)
	}

	///////////////////////////////////////////////POST////////////////////////////////////////////////
	// Define API router group for authorized routes
	apiRoute := api.Group("/api")
	apiRoute.Use(middleware.AuthorizeAdminRequest())
	{
		apiRoute.POST("/tickets", handlers.ItsmHandler)
		apiRoute.POST("/assets", handlers.ItsmHandler)
		apiRoute.POST("/staff", handlers.ItsmHandler)
		apiRoute.POST("/agent", handlers.ItsmHandler)
		apiRoute.POST("/admin", handlers.ItsmHandler)
		apiRoute.POST("/procurement", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiTicket := api.Group("/api/ticket")
	apiTicket.Use(middleware.AuthorizeRequest())
	{
		apiTicket.POST("/status", handlers.ItsmHandler)
		apiTicket.POST("/assets", handlers.ItsmHandler)
		apiTicket.POST("/category", handlers.ItsmHandler)
		apiTicket.POST("/sub-category", handlers.ItsmHandler)
		apiTicket.POST("/staff", handlers.ItsmHandler)
		apiTicket.POST("/agent", handlers.ItsmHandler)
		apiTicket.POST("/admin", handlers.ItsmHandler)
		apiTicket.POST("/priority", handlers.ItsmHandler)
		apiTicket.POST("/sla", handlers.ItsmHandler)
		apiTicket.POST("/satisfaction", handlers.ItsmHandler)
		apiTicket.POST("/policy", handlers.ItsmHandler)
		apiTicket.POST("/related", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAgent := api.Group("/api/agent")
	apiAgent.Use(middleware.AuthorizeAdminRequest())
	{
		apiAgent.POST("/role", handlers.ItsmHandler)
		apiAgent.POST("/unit", handlers.ItsmHandler)
		apiAgent.POST("/username", handlers.ItsmHandler)
		apiAgent.POST("/supervisor", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiStaff := api.Group("/api/staff")
	apiStaff.Use(middleware.AuthorizeAdminRequest())
	{
		apiStaff.POST("/role", handlers.ItsmHandler)
		apiStaff.POST("/unit", handlers.ItsmHandler)
		apiStaff.POST("/username", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAsset := api.Group("/api/assets")
	apiAsset.Use(middleware.AuthorizeAdminRequest())
	{
		apiAsset.POST("/creator", handlers.ItsmHandler)
	}

	///////////////////////////////////////////////PUT////////////////////////////////////////////////

	// Define API router group for authorized routes
	apiRouteU := api.Group("/api")
	apiRouteU.Use(middleware.AuthorizeAdminRequest())
	{
		apiRouteU.PUT("/tickets", handlers.ItsmHandler)
		apiRouteU.PUT("/assets", handlers.ItsmHandler)
		apiRouteU.PUT("/staff", handlers.ItsmHandler)
		apiRouteU.PUT("/agent", handlers.ItsmHandler)
		apiRouteU.PUT("/admin", handlers.ItsmHandler)
		apiRouteU.PUT("/procurement", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiTicketU := api.Group("/api/ticket")
	apiTicketU.Use(middleware.AuthorizeRequest())
	{
		apiTicketU.PUT("/status", handlers.ItsmHandler)
		apiTicketU.PUT("/assets", handlers.ItsmHandler)
		apiTicketU.PUT("/category", handlers.ItsmHandler)
		apiTicketU.PUT("/sub-category", handlers.ItsmHandler)
		apiTicketU.PUT("/staff", handlers.ItsmHandler)
		apiTicketU.PUT("/agent", handlers.ItsmHandler)
		apiTicketU.PUT("/admin", handlers.ItsmHandler)
		apiTicketU.PUT("/priority", handlers.ItsmHandler)
		apiTicketU.PUT("/sla", handlers.ItsmHandler)
		apiTicketU.PUT("/satisfaction", handlers.ItsmHandler)
		apiTicketU.PUT("/policy", handlers.ItsmHandler)
		apiTicketU.PUT("/related", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAgentU := api.Group("/api/agent")
	apiAgentU.Use(middleware.AuthorizeAdminRequest())
	{
		apiAgentU.PUT("/role", handlers.ItsmHandler)
		apiAgentU.PUT("/unit", handlers.ItsmHandler)
		apiAgentU.PUT("/username", handlers.ItsmHandler)
		apiAgentU.PUT("/supervisor", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiStaffU := api.Group("/api/staff")
	apiStaffU.Use(middleware.AuthorizeAdminRequest())
	{
		apiStaffU.PUT("/role", handlers.ItsmHandler)
		apiStaffU.PUT("/unit", handlers.ItsmHandler)
		apiStaffU.PUT("/username", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAssetU := api.Group("/api/assets")
	apiAssetU.Use(middleware.AuthorizeAdminRequest())
	{
		apiAssetU.PUT("/creator", handlers.ItsmHandler)
	}

	///////////////////////////////////////////////PUT////////////////////////////////////////////////

	// Define API router group for authorized routes
	apiRouteD := api.Group("/api")
	apiRouteD.Use(middleware.AuthorizeAdminRequest())
	{
		apiRouteD.DELETE("/tickets", handlers.ItsmHandler)
		apiRouteU.DELETE("/assets", handlers.ItsmHandler)
		apiRouteU.DELETE("/staff", handlers.ItsmHandler)
		apiRouteU.DELETE("/agent", handlers.ItsmHandler)
		apiRouteU.DELETE("/admin", handlers.ItsmHandler)
		apiRouteU.DELETE("/procurement", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiTicketD := api.Group("/api/ticket")
	apiTicketD.Use(middleware.AuthorizeRequest())
	{
		apiTicketU.DELETE("/status", handlers.ItsmHandler)
		apiTicketU.DELETE("/assets", handlers.ItsmHandler)
		apiTicketU.DELETE("/category", handlers.ItsmHandler)
		apiTicketU.DELETE("/sub-category", handlers.ItsmHandler)
		apiTicketU.DELETE("/staff", handlers.ItsmHandler)
		apiTicketU.DELETE("/agent", handlers.ItsmHandler)
		apiTicketU.DELETE("/admin", handlers.ItsmHandler)
		apiTicketU.DELETE("/priority", handlers.ItsmHandler)
		apiTicketU.DELETE("/sla", handlers.ItsmHandler)
		apiTicketU.DELETE("/satisfaction", handlers.ItsmHandler)
		apiTicketU.DELETE("/policy", handlers.ItsmHandler)
		apiTicketU.DELETE("/related", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAgentD := api.Group("/api/agent")
	apiAgentD.Use(middleware.AuthorizeAdminRequest())
	{
		apiAgentU.DELETE("/role", handlers.ItsmHandler)
		apiAgentU.DELETE("/unit", handlers.ItsmHandler)
		apiAgentU.DELETE("/username", handlers.ItsmHandler)
		apiAgentU.DELETE("/supervisor", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiStaffD := api.Group("/api/staff")
	apiStaffD.Use(middleware.AuthorizeAdminRequest())
	{
		apiStaffD.DELETE("/role", handlers.ItsmHandler)
		apiStaffD.DELETE("/unit", handlers.ItsmHandler)
		apiStaffD.DELETE("/username", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
	apiAssetD := api.Group("/api/assets")
	apiAssetD.Use(middleware.AuthorizeAdminRequest())
	{
		apiAssetD.DELETE("/creator", handlers.ItsmHandler)
	}
}
