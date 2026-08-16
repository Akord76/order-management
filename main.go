package main

import (
	"log"

	"order-management/config"
	"order-management/handler"
	"order-management/repository"
	"order-management/routes"
	"order-management/service"
)

func main() {
	// ---- Config & Database ----
	cfg := config.LoadConfig()
	db := config.ConnectDatabase(cfg)

	// ---- Repositories ----
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	orderSaleRepo := repository.NewOrderSaleRepository(db)

	// ---- Services ----
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpireHour)
	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)
	customerService := service.NewCustomerService(customerRepo)
	employeeService := service.NewEmployeeService(employeeRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	orderService := service.NewOrderService(orderRepo)
	orderSaleService := service.NewOrderSaleService(orderSaleRepo)

	// ---- Handlers ----
	h := &routes.Handlers{
		Auth:      handler.NewAuthHandler(authService),
		User:      handler.NewUserHandler(userService),
		Category:  handler.NewCategoryHandler(categoryService),
		Product:   handler.NewProductHandler(productService),
		Customer:  handler.NewCustomerHandler(customerService),
		Employee:  handler.NewEmployeeHandler(employeeService),
		Supplier:  handler.NewSupplierHandler(supplierService),
		Order:     handler.NewOrderHandler(orderService),
		OrderSale: handler.NewOrderSaleHandler(orderSaleService),
	}

	// ---- Router ----
	router := routes.SetupRouter(cfg.JWTSecret, h)
	log.Printf("server starting on :%s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
