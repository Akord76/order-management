package routes

import (
	"order-management/handler"
	"order-management/middleware"

	"github.com/gin-gonic/gin"
)

// Handlers bundles every HTTP handler the router needs. Built once in
// main.go and passed in here so routing stays decoupled from wiring.
type Handlers struct {
	Auth      *handler.AuthHandler
	User      *handler.UserHandler
	Category  *handler.CategoryHandler
	Product   *handler.ProductHandler
	Customer  *handler.CustomerHandler
	Employee  *handler.EmployeeHandler
	Supplier  *handler.SupplierHandler
	Order     *handler.OrderHandler
	OrderSale *handler.OrderSaleHandler
}

// SetupRouter wires every route and applies the role hierarchy:
//
//	JWT Authentication -> JWT Middleware -> User Identity -> Role Authorization
//	                                                          /     |      \
//	                                                     ADMIN  MANAGER   USER
//
//	ADMIN   : Create / Update / Delete / View Users
//	MANAGER : Create / Update / View Order
//	USER    : View Order
func SetupRouter(jwtSecret string, h *Handlers) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// ---- Public routes (no JWT required) ----
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", h.Auth.Register)
		authGroup.POST("/login", h.Auth.Login)
	}

	// ---- Everything below requires a valid JWT ----
	protected := api.Group("")
	protected.Use(middleware.JWTAuth(jwtSecret))

	// ============================================================
	// ADMIN — User Management
	// ============================================================
	users := protected.Group("/users")
	users.Use(middleware.RequireRoles(middleware.RoleAdmin))
	{
		users.POST("", h.User.Create)       // Create User
		users.GET("", h.User.GetAll)        // View Users
		users.GET("/:id", h.User.GetByID)   // View Users (single)
		users.PUT("/:id", h.User.Update)    // Update User
		users.DELETE("/:id", h.User.Delete) // Delete User
	}

	// ============================================================
	// Orders — MANAGER can create/update, MANAGER + USER + ADMIN can view
	// ============================================================
	orders := protected.Group("/orders")
	{
		orders.GET("", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.Order.GetAll)
		orders.GET("/:orderID/:orderNo", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.Order.GetByID)
		orders.GET("/:orderID/:orderNo/details", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.Order.GetDetails)

		orders.POST("", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.Order.Create)                     // Create Order
		orders.PUT("/:orderID/:orderNo", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.Order.Update)    // Update Order
		orders.DELETE("/:orderID/:orderNo", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.Order.Delete) // Update/remove Order

		// FIXED: Updated the prefix from "/:orderNo/..." to "/:orderID/:orderNo/..." to prevent wildcard conflicts
		orders.POST("/:orderID/:orderNo/details", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.Order.AddDetail)
		orders.DELETE("/:orderID/:orderNo/details/:orderDetailNo", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.Order.DeleteDetail)
	}

	// Order Sales follow the same MANAGER-write / everyone-read pattern.
	orderSales := protected.Group("/order-sales")
	{
		orderSales.GET("", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.OrderSale.GetAll)
		orderSales.GET("/:orderSaleNo/:orderSaleID", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.OrderSale.GetByID)

		// FIXED: Standardized prefix to include both parameters (avoids static vs wildcard conflicts at the second segment)
		orderSales.GET("/:orderSaleNo/:orderSaleID/details", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager, middleware.RoleUser), h.OrderSale.GetDetails)

		orderSales.POST("", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.OrderSale.Create)
		orderSales.PUT("/:orderSaleNo/:orderSaleID", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.OrderSale.Update)
		orderSales.DELETE("/:orderSaleNo/:orderSaleID", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.OrderSale.Delete)

		// FIXED: Standardized prefix to include both parameters
		orderSales.POST("/:orderSaleNo/:orderSaleID/details", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.OrderSale.AddDetail)
		orderSales.DELETE("/:orderSaleNo/:orderSaleID/details/:orderSaleDetailNo/:orderSaleDetailID", middleware.RequireRoles(middleware.RoleAdmin, middleware.RoleManager), h.OrderSale.DeleteDetail)
	}
	// ============================================================
	// Master data — any authenticated user may read/write.
	// Tighten with middleware.RequireRoles(...) per-route if the
	// business rules ever need it (e.g. restrict writes to ADMIN).
	// ============================================================
	categories := protected.Group("/categories")
	{
		categories.POST("", h.Category.Create)
		categories.GET("", h.Category.GetAll)
		categories.GET("/:id", h.Category.GetByID)
		categories.PUT("/:id", h.Category.Update)
		categories.DELETE("/:id", h.Category.Delete)
	}

	products := protected.Group("/products")
	{
		products.POST("", h.Product.Create)
		products.GET("", h.Product.GetAll)
		products.GET("/:productNumber/:productID", h.Product.GetByID)
		products.PUT("/:productNumber/:productID", h.Product.Update)
		products.DELETE("/:productNumber/:productID", h.Product.Delete)
	}

	customers := protected.Group("/customers")
	{
		customers.POST("", h.Customer.Create)
		customers.GET("", h.Customer.GetAll)
		customers.GET("/:custNumber/:custID", h.Customer.GetByID)
		customers.PUT("/:custNumber/:custID", h.Customer.Update)
		customers.DELETE("/:custNumber/:custID", h.Customer.Delete)
	}

	employees := protected.Group("/employees")
	{
		employees.POST("", h.Employee.Create)
		employees.GET("", h.Employee.GetAll)
		employees.GET("/:employeeID/:cardNumber", h.Employee.GetByID)
		employees.PUT("/:employeeID/:cardNumber", h.Employee.Update)
		employees.DELETE("/:employeeID/:cardNumber", h.Employee.Delete)
	}

	suppliers := protected.Group("/suppliers")
	{
		suppliers.POST("", h.Supplier.Create)
		suppliers.GET("", h.Supplier.GetAll)
		suppliers.GET("/:supplierNumber/:supplierID", h.Supplier.GetByID)
		suppliers.PUT("/:supplierNumber/:supplierID", h.Supplier.Update)
		suppliers.DELETE("/:supplierNumber/:supplierID", h.Supplier.Delete)
	}

	return r
}
