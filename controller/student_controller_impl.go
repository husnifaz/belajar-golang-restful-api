package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/service"
	"strconv"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentCreateRequest := web.StudentCreateRequest{}
	helper.ReadFromRequestBody(request, &studentCreateRequest)

	studentResponse := controller.StudentService.Create(request.Context(), studentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentId := params.ByName("StudentId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	studentResponse := controller.StudentService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	studentResponses := controller.StudentService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
