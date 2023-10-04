package repository

import (
	"context"
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type StudentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student
	Delete(ctx context.Context, tx *sql.Tx, student domain.Student)
	FindById(ctx context.Context, tx *sql.Tx, studentId int) (domain.Student, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Student
}
