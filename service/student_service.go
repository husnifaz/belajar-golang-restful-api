package service

import (
	"context"
	"programmerzamannow/belajar-golang-restful-api/model/web"
)

type StudentService interface {
	Create(ctx context.Context, request web.StudentCreateRequest) web.StudentResponse
	FindById(ctx context.Context, StudentId int) web.StudentResponse
	FindAll(ctx context.Context) []web.StudentResponse
}
