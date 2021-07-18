package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
)

type Rest struct {
	*gin.Engine
}

// New creates a application instance
func New(rateLimit float64) *Rest {
	engine := gin.Default()
	engine.Use(RateLimit(rateLimit))
	engine.Use(gin.Recovery())
	engine.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	return &Rest{Engine: engine}
}

// Controller Register Controller application
func (a *Rest) Controller(c IController) {
	g := a.Group(c.GetName())
	c.Register(a, g)
}

// ResponseHandler gin.Context handler with IResponse
type ResponseHandler func(*gin.Context) IResponse