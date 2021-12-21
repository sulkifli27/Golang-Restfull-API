package main

import (
	"net/http"
	"sulkifli27/belajar-golang-restful-api/app"
	"sulkifli27/belajar-golang-restful-api/controller"
	"sulkifli27/belajar-golang-restful-api/helper"
	"sulkifli27/belajar-golang-restful-api/middleware"
	"sulkifli27/belajar-golang-restful-api/repository"
	"sulkifli27/belajar-golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router) ,
	}

	err := server.ListenAndServe()
	helper.PanicIferror(err)
}