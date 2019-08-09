package router

import (
	"github.com/gin-gonic/gin"
	"go-music/controllers"
	"go-music/repositories"
	"net/http"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	AuthenLevel int
}

// Routes holds configurations related to API of this project
type Routes struct {
	v1       []route
}

func (r *Routes) Init(ctx *repositories.MusicContext) http.Handler {
	productRepository := repositories.NewProductRepository(ctx)
	customerRepository := repositories.NewCustomerRepository(ctx)
	productController := controllers.NewProductController(productRepository)
	customerController := controllers.NewCustomerController(customerRepository)


	r.v1 = []route{
		{
			Name:        "Get Products",
			Description: "Get Products",
			Method:      http.MethodGet,
			Pattern:     "/products",
			Endpoint:    productController.GetAllProducts,
			//AuthenLevel: ValidateNone,
		},
		{
			Name: "Sign in",
			Description: "Sign in",
			Method: http.MethodPost,
			Pattern: "/signin",
			Endpoint: customerController.SignIn,
		},
		{
			Name: "Get Customer By Id",
			Description: "Get Customer By Id",
			Method: http.MethodGet,
			Pattern: "/customers/:id",
			Endpoint: customerController.GetCustomerById,
		},
	}

	router := gin.New()
	v1 := router.Group("v1")
	for _, e := range r.v1 {
		v1.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return router
}

//type Route struct {
//	Name        string
//	Method      string
//	Pattern     string
//	HandlerFunc http.HandlerFunc
//	AuthMethod  int
//}
//
//type Routes []Route


