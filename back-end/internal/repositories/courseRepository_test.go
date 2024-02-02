package repositories

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"testing"
)

func TestCourseRepository_GetCourse(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid uint32
	}
	tests := []struct {
		name                       string
		args                       args
		mockSqlCourseQueryExpected string
		mockCourseReturned         *Course
		expected                   *domain.Course
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
			},
			mockSqlCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
			mockCourseReturned: &Course{
				Model:       gorm.Model{ID: 2},
				Title:       "title",
				Description: "description",
			},
			expected: &domain.Course{
				Title:       "title",
				Description: "description",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CourseRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "title", "description"},
					).AddRow(
						tt.mockCourseReturned.ID, tt.mockCourseReturned.Title, tt.mockCourseReturned.Description,
					),
				)

			actual, err := repo.GetCourse(context.Background(), tt.args.uid)
			if err != nil {
				t.Errorf("GetCourse() error = %v", err)
				return
			}

			assert.Equal(t, tt.expected, actual)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestCourseRepository_GetCourseReturnsError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid uint32
	}
	tests := []struct {
		name                       string
		args                       args
		mockSqlCourseQueryExpected string
		mockSqlCourseErrorReturned error
		expectedError              error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
			mockSqlCourseErrorReturned: errors.New("random error"),
			expectedError:              errors.New("random error"),
		},
		{
			name: "course not found",
			args: args{
				uid: 2,
			},
			mockSqlCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
			mockSqlCourseErrorReturned: gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("courseID 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CourseRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnError(tt.mockSqlCourseErrorReturned)

			_, actual := repo.GetCourse(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestCourseRepository_GetCourseByStudentID(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		studentID uint32
	}
	tests := []struct {
		name                       string
		args                       args
		mockSqlCourseQueryExpected string
		mockCoursesReturned        []Course
		expected                   []domain.Course
	}{
		{
			name: "valid",
			args: args{
				studentID: 2,
			},
			mockSqlCourseQueryExpected: `SELECT "courses"."id","courses"."created_at","courses"."updated_at","courses"."deleted_at","courses"."title","courses"."description","courses"."tutor_id" FROM "courses" JOIN "student_takes_course" ON "student_takes_course"."course_id" = "courses"."id" AND "student_takes_course"."student_id" IN (NULL) WHERE student_id = $1 AND "courses"."deleted_at" IS NULL`,
			mockCoursesReturned: []Course{
				{
					Model:       gorm.Model{ID: 2},
					Title:       "title",
					Description: "description",
				},
			},
			expected: []domain.Course{
				{
					Title:       "title",
					Description: "description",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CourseRepository{
				db: gormDb,
			}

			expectedRows := sqlmock.NewRows(
				[]string{"id", "title", "description"},
			)
			for _, course := range tt.mockCoursesReturned {
				expectedRows.AddRow(
					course.ID, course.Title, course.Description,
				)
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
				WithArgs(tt.args.studentID).
				WillReturnRows(expectedRows)

			actual, err := repo.GetCourseByStudentID(context.Background(), tt.args.studentID)
			if err != nil {
				t.Errorf("GetCourse() error = %v", err)
				return
			}

			assert.Equal(t, tt.expected, actual)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestCourseRepository_GetCourseByStudentIDReturnsError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		studentID uint32
	}
	tests := []struct {
		name                       string
		args                       args
		mockSqlCourseQueryExpected string
		mockSqlCourseErrorReturned error
		expectedError              error
	}{
		{
			name: "random error",
			args: args{
				studentID: 2,
			},
			mockSqlCourseQueryExpected: `SELECT "courses"."id","courses"."created_at","courses"."updated_at","courses"."deleted_at","courses"."title","courses"."description","courses"."tutor_id" FROM "courses" JOIN "student_takes_course" ON "student_takes_course"."course_id" = "courses"."id" AND "student_takes_course"."student_id" IN (NULL) WHERE student_id = $1 AND "courses"."deleted_at" IS NULL`,
			mockSqlCourseErrorReturned: errors.New("random error"),
			expectedError:              errors.New("random error"),
		},
		{
			name: "course not found",
			args: args{
				studentID: 2,
			},
			mockSqlCourseQueryExpected: `SELECT "courses"."id","courses"."created_at","courses"."updated_at","courses"."deleted_at","courses"."title","courses"."description","courses"."tutor_id" FROM "courses" JOIN "student_takes_course" ON "student_takes_course"."course_id" = "courses"."id" AND "student_takes_course"."student_id" IN (NULL) WHERE student_id = $1 AND "courses"."deleted_at" IS NULL`,
			mockSqlCourseErrorReturned: gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("studentID 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &CourseRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
				WithArgs(tt.args.studentID).
				WillReturnError(tt.mockSqlCourseErrorReturned)

			_, actual := repo.GetCourseByStudentID(context.Background(), tt.args.studentID)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

//func TestCourseRepository_UpdateExistingCourse(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		courseID uint
//		course   *domain.Course
//	}
//	tests := []struct {
//		name                             string
//		args                             args
//		mockSqlCourseSelectQueryExpected string
//		mockCourseReturned               *Course
//		mockUpdatedCourseIdReturned      int
//		mockSqlCourseUpdateQueryExpected string
//	}{
//		{
//			name: "valid",
//			args: args{
//				courseID: 2,
//				course: &domain.Course{
//					Title:       "title",
//					Description: "description",
//				},
//			},
//			mockSqlCourseSelectQueryExpected: `SELECT * FROM "courses" WHERE "courses"."deleted_at" IS NULL ORDER BY "courses"."id" LIMIT 1`,
//			mockCourseReturned: &Course{
//				Model: gorm.Model{ID: 2},
//			},
//			mockSqlCourseUpdateQueryExpected: `UPDATE "courses" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"title"=$4,"description"=$5,"tutor_id"=$6 WHERE "courses"."deleted_at" IS NULL AND "id" = $7`,
//			mockUpdatedCourseIdReturned:      2,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseSelectQueryExpected)).
//				WithArgs().
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{"id", "title", "description"},
//					).AddRow(
//						tt.mockCourseReturned.ID, tt.mockCourseReturned.Title, tt.mockCourseReturned.Description,
//					),
//				)
//
//			mockDb.ExpectBegin()
//			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlCourseUpdateQueryExpected)).
//				WithArgs(
//					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
//					tt.args.course.Title, tt.args.course.Description,
//					sqlmock.AnyArg(), tt.mockCourseReturned.ID,
//				).
//				WillReturnResult(
//					sqlmock.NewResult(
//						int64(tt.mockUpdatedCourseIdReturned),
//						1,
//					),
//				)
//			mockDb.ExpectCommit()
//
//			err := repo.UpdateCourse(
//				context.Background(),
//				uint32(tt.args.courseID),
//				tt.args.course,
//			)
//
//			assert.NoError(t, err)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
//
//func TestCourseRepository_UpdateExistingCourseReturnsCourseSelectError(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		courseID uint
//		course   *domain.Course
//	}
//	tests := []struct {
//		name                             string
//		args                             args
//		mockSqlCourseSelectQueryExpected string
//		mockSqlCourseErrorReturned       error
//		expectedError                    error
//	}{
//		{
//			name: "random error",
//			args: args{
//				courseID: 2,
//				course: &domain.Course{
//					Title:       "title",
//					Description: "description",
//				},
//			},
//			mockSqlCourseSelectQueryExpected: `SELECT * FROM "courses" WHERE "courses"."deleted_at" IS NULL ORDER BY "courses"."id" LIMIT 1`,
//			mockSqlCourseErrorReturned:       errors.New("random error"),
//			expectedError:                    errors.New("random error"),
//		},
//		{
//			name: "course not found",
//			args: args{
//				courseID: 2,
//				course: &domain.Course{
//					Title:       "title",
//					Description: "description",
//				},
//			},
//			mockSqlCourseSelectQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
//			mockSqlCourseErrorReturned:       gorm.ErrRecordNotFound,
//			expectedError: apierrors.DataNotFoundErrorWrapper{
//				ReturnedStatusCode: http.StatusNotFound,
//				OriginalError:      errors.New("courseID 2 not found"),
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseSelectQueryExpected)).
//				WithArgs().
//				WillReturnError(tt.mockSqlCourseErrorReturned)
//
//			actual := repo.UpdateCourse(context.Background(), uint32(tt.args.courseID), tt.args.course)
//
//			assert.Equal(t, tt.expectedError, actual)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
//
//func TestCourseRepository_CreateCourse(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		course  *domain.Course
//		tutorID uint
//	}
//	tests := []struct {
//		name                         string
//		args                         args
//		mockSqlCourseQueryExpected   string
//		mockInsertedCourseIdReturned int
//		expectedCourseUid            uint
//	}{
//		{
//			name: "valid",
//			args: args{
//				course: &domain.Course{
//					Title:       "title",
//					Description: "description",
//				},
//				tutorID: 1,
//			},
//			mockSqlCourseQueryExpected:   `INSERT INTO "courses" ("created_at","updated_at","deleted_at","title","description","tutor_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`,
//			mockInsertedCourseIdReturned: 2,
//			expectedCourseUid:            2,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectBegin()
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
//				WithArgs(
//					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
//					tt.args.course.Title, tt.args.course.Description, tt.args.tutorID,
//				).
//				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedCourseIdReturned))
//
//			mockDb.ExpectCommit()
//
//			actual, err := repo.CreateCourse(context.Background(), tt.args.course, tt.args.tutorID)
//			if err != nil {
//				t.Errorf("CreateCourse() error = %v", err)
//			}
//
//			assert.Equal(t, tt.expectedCourseUid, actual)
//
//			if err = mockDb.ExpectationsWereMet(); err != nil {
//				t.Errorf("there were unfulfilled expections: %s", err)
//			}
//		})
//	}
//}
//
//func TestCourseRepository_CreateCourseReturnsErrorOnCourseCreate(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		course  *domain.Course
//		tutorID uint
//	}
//	tests := []struct {
//		name                       string
//		args                       args
//		mockSqlCourseQueryExpected string
//		mockSqlCourseQueryError    error
//		expectedError              error
//	}{
//		{
//			name: "random error",
//			args: args{
//				course: &domain.Course{
//					Title:       "title",
//					Description: "description",
//				},
//				tutorID: 1,
//			},
//			mockSqlCourseQueryExpected: `INSERT INTO "courses" ("created_at","updated_at","deleted_at","title","description","tutor_id") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`,
//			mockSqlCourseQueryError:    errors.New("random error"),
//			expectedError:              errors.New("random error"),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectBegin()
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlCourseQueryExpected)).
//				WithArgs(
//					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
//					tt.args.course.Title, tt.args.course.Description, tt.args.tutorID,
//				).
//				WillReturnError(tt.mockSqlCourseQueryError)
//			mockDb.ExpectRollback()
//
//			_, actualError := repo.CreateCourse(context.Background(), tt.args.course, tt.args.tutorID)
//
//			assert.Equal(t, tt.expectedError, actualError)
//
//			if err = mockDb.ExpectationsWereMet(); err != nil {
//				t.Errorf("there were unfulfilled expections: %s", err)
//			}
//		})
//	}
//}
//
//func TestCourseRepository_DeleteCourse(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		uid uint32
//	}
//	tests := []struct {
//		name                             string
//		args                             args
//		mockSqlSelectCourseQueryExpected string
//		mockSqlDeleteCourseQueryExpected string
//		mockCourseReturned               *Course
//		mockDeletedCourseIdReturned      int
//	}{
//		{
//			name: "valid",
//			args: args{
//				uid: 2,
//			},
//			mockSqlSelectCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
//			mockCourseReturned: &Course{
//				Model:       gorm.Model{ID: 2},
//				Title:       "title",
//				Description: "description",
//			},
//			// The query is UPDATE as the deletion is SOFT in the gorm action I've set
//			mockSqlDeleteCourseQueryExpected: `UPDATE "courses" SET "deleted_at"=$1 WHERE id = $2 AND "courses"."deleted_at" IS NULL`,
//			mockDeletedCourseIdReturned:      2,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectCourseQueryExpected)).
//				WithArgs(tt.args.uid).
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{"id", "title", "description"},
//					).AddRow(
//						tt.mockCourseReturned.ID, tt.mockCourseReturned.Title,
//						tt.mockCourseReturned.Description,
//					),
//				)
//			mockDb.ExpectBegin()
//			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteCourseQueryExpected)).
//				WithArgs(sqlmock.AnyArg(), tt.args.uid).
//				WillReturnResult(sqlmock.NewResult(int64(tt.mockDeletedCourseIdReturned), 1))
//			mockDb.ExpectCommit()
//
//			err := repo.DeleteCourse(context.Background(), tt.args.uid)
//
//			assert.NoError(t, err)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
//
//// The delete action in the gorm is done by selecting the requested course and then
//// soft delete it, so we have two possible error sources: the select and the update
//func TestCourseRepository_DeleteCourseWithSelectError(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		uid uint32
//	}
//	tests := []struct {
//		name                             string
//		args                             args
//		mockSqlSelectCourseQueryExpected string
//		mockSqlErrorReturned             error
//		expectedError                    error
//	}{
//		{
//			name: "random error",
//			args: args{
//				uid: 2,
//			},
//			mockSqlSelectCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
//			mockSqlErrorReturned:             errors.New("random error"),
//			expectedError:                    errors.New("random error"),
//		},
//		{
//			name: "course not found",
//			args: args{
//				uid: 2,
//			},
//			mockSqlSelectCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
//			mockSqlErrorReturned:             gorm.ErrRecordNotFound,
//			expectedError: apierrors.DataNotFoundErrorWrapper{
//				ReturnedStatusCode: http.StatusNotFound,
//				OriginalError:      errors.New("courseID 2 not found"),
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectCourseQueryExpected)).
//				WithArgs(tt.args.uid).
//				WillReturnError(tt.mockSqlErrorReturned)
//
//			actual := repo.DeleteCourse(context.Background(), tt.args.uid)
//
//			assert.Equal(t, actual, tt.expectedError)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
//
//func TestCourseRepository_DeleteCourseWithUpdateError(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		uid uint32
//	}
//	tests := []struct {
//		name                             string
//		args                             args
//		mockSqlSelectCourseQueryExpected string
//		mockSqlDeleteCourseQueryExpected string
//		mockCourseReturned               *Course
//		mockSqlErrorReturned             error
//		expectedError                    error
//	}{
//		{
//			name: "random error",
//			args: args{
//				uid: 666,
//			},
//			mockSqlSelectCourseQueryExpected: `SELECT * FROM "courses" WHERE id = $1 AND "courses"."deleted_at" IS NULL LIMIT 1`,
//			mockCourseReturned: &Course{
//				Model:       gorm.Model{ID: 2},
//				Title:       "title",
//				Description: "description",
//			},
//			// The query is UPDATE as the deletion is SOFT in the gorm action I've set
//			mockSqlDeleteCourseQueryExpected: `UPDATE "courses" SET "deleted_at"=$1 WHERE id = $2 AND "courses"."deleted_at" IS NULL`,
//			mockSqlErrorReturned:             errors.New("random error"),
//			expectedError:                    errors.New("random error"),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &CourseRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectCourseQueryExpected)).
//				WithArgs(tt.args.uid).
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{"id", "title", "description"},
//					).AddRow(
//						tt.mockCourseReturned.ID, tt.mockCourseReturned.Title,
//						tt.mockCourseReturned.Description,
//					),
//				)
//			mockDb.ExpectBegin()
//			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteCourseQueryExpected)).
//				WithArgs(sqlmock.AnyArg(), tt.args.uid).
//				WillReturnError(tt.mockSqlErrorReturned)
//			mockDb.ExpectRollback()
//
//			actual := repo.DeleteCourse(context.Background(), tt.args.uid)
//
//			assert.Equal(t, actual, tt.expectedError)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
