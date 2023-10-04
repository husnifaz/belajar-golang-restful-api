package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"programmerzamannow/belajar-golang-restful-api/exception"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/repository"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewStudentService(studentRepository repository.StudentRepository, DB *sql.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *StudentServiceImpl) Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student := domain.Student{
		Name:    request.Name,
		Address: request.Address,
	}

	student = service.StudentRepository.Save(ctx, tx, student)

	return helper.ToStudentResponse(student)
}

func (service *StudentServiceImpl) FindById(ctx context.Context, studentId int) web.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStudentResponse(student)
}

func (service *StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := service.StudentRepository.FindAll(ctx, tx)

	return helper.ToStudentResponses(students)
}
