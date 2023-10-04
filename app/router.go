package app

import (
	"github.com/julienschmidt/httprouter"
	"programmerzamannow/belajar-golang-restful-api/controller"
	"programmerzamannow/belajar-golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController, studentController controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/students", studentController.FindAll)
	router.GET("/api/students/:studentId", studentController.FindById)
	router.POST("/api/students", studentController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
