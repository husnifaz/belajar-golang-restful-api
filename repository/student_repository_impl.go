package repository

import (
	"context"
	"database/sql"
	"errors"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/domain"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepository() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (repository *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student domain.Student) domain.Student {
	SQL := "insert into student(name, address) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, student.Name, student.Address)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.Id = int(id)
	return student
}

func (repository *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, student domain.Student) {
	SQL := "delete from student where id = ?"
	_, err := tx.ExecContext(ctx, SQL, student.Id)
	helper.PanicIfError(err)
}

func (repository *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, studentId int) (domain.Student, error) {
	SQL := "select id, name from student where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := domain.Student{}
	if rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Address)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

func (repository *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Student {
	SQL := "select id, name from student"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []domain.Student
	for rows.Next() {
		student := domain.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Address)
		helper.PanicIfError(err)
		students = append(students, student)
	}
	return students
}
