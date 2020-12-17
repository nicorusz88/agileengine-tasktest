package builder

import (
	"github.com/agileengine-tasktest/src/api/controller"
	"github.com/agileengine-tasktest/src/api/dependencies"
	"github.com/gin-gonic/gin"
)

const (
	APP_NAME = "accounting"
)

var (
	Router *gin.Engine
)

type Builder struct{}

func Build() Builder {
	return Builder{}
}

func (_ Builder) ConfigureRouter(container *dependencies.Container) {

	Router = gin.New()
	gin.SetMode(gin.ReleaseMode)

	v0Group := Router.Group("/" + APP_NAME)
	{
		// create account
		v0Group.POST("/account/transaction", handleRequestWrapper(controller.PostTransaction, container))

		// list transactions
		v0Group.GET("/account/transaction", handleRequestWrapper(controller.GetTransactions, container))

		// get transaction
		v0Group.GET("/account/transaction/:id", handleRequestWrapper(controller.GetTransaction, container))

		// get account balance
		v0Group.GET("/account/balance", handleRequestWrapper(controller.GetBalance, container))

	}

	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
}

func handleRequestWrapper(controller func(*gin.Context, *dependencies.Container), container *dependencies.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		controller(c, container)
	}
}

// Run Runs the API
func (api *Builder) Run() {
	Router.Run(":8080")
}
