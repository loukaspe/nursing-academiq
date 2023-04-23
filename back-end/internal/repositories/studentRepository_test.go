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
	"time"
)

func TestStudentRepository_CreateStudent(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		student *domain.Student
	}
	tests := []struct {
		name                          string
		args                          args
		mockSqlUserQueryExpected      string
		mockSqlStudentQueryExpected   string
		mockInsertedStudentIdReturned int
		mockInsertedUserIdReturned    int
		expectedStudentUid            uint
	}{
		{
			name: "valid",
			args: args{
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentQueryExpected:   `INSERT INTO "students" ("created_at","updated_at","deleted_at","registration_number","user_id") VALUES ($1,$2,$3,$4,$5)`,
			mockSqlUserQueryExpected:      `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned:    1,
			mockInsertedStudentIdReturned: 2,
			expectedStudentUid:            2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.Username, tt.args.student.Password, tt.args.student.FirstName,
					tt.args.student.LastName, tt.args.student.Email, tt.args.student.BirthDate, tt.args.student.PhoneNumber,
					tt.args.student.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.RegistrationNumber, tt.mockInsertedUserIdReturned,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedStudentIdReturned))
			mockDb.ExpectCommit()

			actual, err := repo.CreateStudent(context.Background(), tt.args.student)
			if err != nil {
				t.Errorf("CreateStudent() error = %v", err)
			}

			assert.Equal(t, tt.expectedStudentUid, actual)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestStudentRepository_CreateStudentReturnsErrorOnStudentCreate(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		student *domain.Student
	}
	tests := []struct {
		name                        string
		args                        args
		mockSqlUserQueryExpected    string
		mockInsertedUserIdReturned  int
		mockSqlStudentQueryExpected string
		mockSqlStudentQueryError    error
		expectedError               error
	}{
		{
			name: "random error",
			args: args{
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlUserQueryExpected:    `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned:  1,
			mockSqlStudentQueryExpected: `INSERT INTO "students" ("created_at","updated_at","deleted_at","registration_number","user_id") VALUES ($1,$2,$3,$4,$5)`,
			mockSqlStudentQueryError:    errors.New("random error"),
			expectedError:               errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.Username, tt.args.student.Password, tt.args.student.FirstName,
					tt.args.student.LastName, tt.args.student.Email, tt.args.student.BirthDate, tt.args.student.PhoneNumber,
					tt.args.student.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.RegistrationNumber, tt.mockInsertedUserIdReturned,
				).
				WillReturnError(tt.mockSqlStudentQueryError)
			mockDb.ExpectRollback()

			_, actualError := repo.CreateStudent(context.Background(), tt.args.student)

			assert.Equal(t, tt.expectedError, actualError)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestStudentRepository_CreateStudentReturnsErrorOnUserCreate(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		student *domain.Student
	}
	tests := []struct {
		name                     string
		args                     args
		mockSqlUserQueryExpected string
		mockSqlUserQueryError    error
		expectedError            error
	}{
		{
			name: "random error",
			args: args{
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlUserQueryExpected: `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockSqlUserQueryError:    errors.New("random error"),
			expectedError:            errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.Username, tt.args.student.Password, tt.args.student.FirstName,
					tt.args.student.LastName, tt.args.student.Email, tt.args.student.BirthDate, tt.args.student.PhoneNumber,
					tt.args.student.Photo,
				).
				WillReturnError(tt.mockSqlUserQueryError)
			mockDb.ExpectRollback()

			_, actualError := repo.CreateStudent(context.Background(), tt.args.student)

			assert.Equal(t, tt.expectedError, actualError)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestStudentRepository_CreateStudentReturnsErrorOnUserValidation(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		student *domain.Student
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "missing username",
			args: args{
				student: &domain.Student{
					User: domain.User{
						Username:    "",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			expectedError: apierrors.UserValidationError{
				ReturnedStatusCode: http.StatusBadRequest,
				OriginalError:      errors.New("Required Username"),
			},
		},
		{
			name: "malformed email",
			args: args{
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "malformedemail",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			expectedError: apierrors.UserValidationError{
				ReturnedStatusCode: http.StatusBadRequest,
				OriginalError:      errors.New("Invalid Email"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			_, actualError := repo.CreateStudent(context.Background(), tt.args.student)

			assert.Equal(t, tt.expectedError, actualError)
		})
	}
}

func TestStudentRepository_DeleteStudent(t *testing.T) {
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
		name                              string
		args                              args
		mockSqlSelectStudentQueryExpected string
		mockSqlDeleteStudentQueryExpected string
		mockStudentReturned               *Student
		mockDeletedStudentIdReturned      int
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
			},
			mockSqlSelectStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
				User: User{
					Model:       gorm.Model{ID: 3},
					Username:    "username",
					Password:    "password",
					FirstName:   "firstName",
					LastName:    "lastName",
					Email:       "email@email",
					BirthDate:   time.Time{},
					PhoneNumber: "phoneNumber",
					Photo:       "photo",
				},
			},
			// The query is UPDATE as the deletion is SOFT in the gorm action I've set
			mockSqlDeleteStudentQueryExpected: `UPDATE "students" SET "deleted_at"=$1 WHERE id = $2 AND "students"."deleted_at" IS NULL`,
			mockDeletedStudentIdReturned:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)
			mockDb.ExpectBegin()
			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteStudentQueryExpected)).
				WithArgs(sqlmock.AnyArg(), tt.args.uid).
				WillReturnResult(sqlmock.NewResult(int64(tt.mockDeletedStudentIdReturned), 1))
			mockDb.ExpectCommit()

			err := repo.DeleteStudent(context.Background(), tt.args.uid)

			assert.NoError(t, err)

			mockDb.ExpectationsWereMet()
		})
	}
}

// The delete action in the gorm is done by selecting the requested student and then
// soft delete it, so we have two possible error sources: the select and the update
func TestStudentRepository_DeleteStudentWithSelectError(t *testing.T) {
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
		name                              string
		args                              args
		mockSqlSelectStudentQueryExpected string
		mockSqlErrorReturned              error
		expectedError                     error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlSelectStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlErrorReturned:              errors.New("random error"),
			expectedError:                     errors.New("random error"),
		},
		{
			name: "student not found",
			args: args{
				uid: 2,
			},
			mockSqlSelectStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlErrorReturned:              gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnError(tt.mockSqlErrorReturned)

			actual := repo.DeleteStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_DeleteStudentWithUpdateError(t *testing.T) {
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
		name                              string
		args                              args
		mockSqlSelectStudentQueryExpected string
		mockSqlDeleteStudentQueryExpected string
		mockStudentReturned               *Student
		mockSqlErrorReturned              error
		expectedError                     error
	}{
		{
			name: "random error",
			args: args{
				uid: 666,
			},
			mockSqlSelectStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
				User: User{
					Model:       gorm.Model{ID: 3},
					Username:    "username",
					Password:    "password",
					FirstName:   "firstName",
					LastName:    "lastName",
					Email:       "email@email",
					BirthDate:   time.Time{},
					PhoneNumber: "phoneNumber",
					Photo:       "photo",
				},
			},
			// The query is UPDATE as the deletion is SOFT in the gorm action I've set
			mockSqlDeleteStudentQueryExpected: `UPDATE "students" SET "deleted_at"=$1 WHERE id = $2 AND "students"."deleted_at" IS NULL`,
			mockSqlErrorReturned:              errors.New("random error"),
			expectedError:                     errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)
			mockDb.ExpectBegin()
			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteStudentQueryExpected)).
				WithArgs(sqlmock.AnyArg(), tt.args.uid).
				WillReturnError(tt.mockSqlErrorReturned)
			mockDb.ExpectRollback()

			actual := repo.DeleteStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_GetStudent(t *testing.T) {
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
		name                        string
		args                        args
		mockSqlUserQueryExpected    string
		mockUserReturned            *User
		mockSqlStudentQueryExpected string
		mockStudentReturned         *Student
		expected                    *domain.Student
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
			},
			mockSqlStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockUserReturned: &User{
				Model:       gorm.Model{ID: 3},
				Username:    "username",
				Password:    "password",
				FirstName:   "firstName",
				LastName:    "lastName",
				Email:       "email@email",
				BirthDate:   time.Time{},
				PhoneNumber: "phoneNumber",
				Photo:       "photo",
			},
			expected: &domain.Student{
				User: domain.User{
					Username:    "username",
					Password:    "password",
					FirstName:   "firstName",
					LastName:    "lastName",
					Email:       "email@email",
					BirthDate:   time.Time{},
					PhoneNumber: "phoneNumber",
					Photo:       "photo",
				},
				RegistrationNumber: "registrationNumber",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(tt.mockStudentReturned.UserId).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"id", "username", "password", "first_name",
							"last_name", "email", "birth_date", "phone_number", "photo",
						},
					).AddRow(
						tt.mockUserReturned.ID, tt.mockUserReturned.Username, tt.mockUserReturned.Password,
						tt.mockUserReturned.FirstName, tt.mockUserReturned.LastName, tt.mockUserReturned.Email,
						tt.mockUserReturned.BirthDate, tt.mockUserReturned.PhoneNumber, tt.mockUserReturned.Photo,
					),
				)

			actual, err := repo.GetStudent(context.Background(), tt.args.uid)
			if err != nil {
				t.Errorf("GetStudent() error = %v", err)
				return
			}

			assert.Equal(t, tt.expected, actual)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_GetStudentReturnsErrorOnUserSelect(t *testing.T) {
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
		name                        string
		args                        args
		mockSqlStudentQueryExpected string
		mockStudentReturned         *Student
		mockSqlUserQueryExpected    string
		mockSqlUserErrorReturned    error
		expectedError               error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned: errors.New("random error"),
			expectedError:            errors.New("random error"),
		},
		{
			name: "user not found",
			args: args{
				uid: 2,
			},
			mockSqlStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned: gorm.ErrRecordNotFound,
			// The error is that we didnt find a uid 2 even though the Student with
			// id 2 exists. It's because the Student had not an associated user with uid 2
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(tt.mockStudentReturned.UserId).
				WillReturnError(tt.mockSqlUserErrorReturned)

			_, actual := repo.GetStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_GetStudentReturnsErrorOnStudentSelect(t *testing.T) {
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
		name                        string
		args                        args
		mockSqlStudentQueryExpected string
		mockSqlStudentErrorReturned error
		expectedError               error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlStudentErrorReturned: errors.New("random error"),
			expectedError:               errors.New("random error"),
		},
		{
			name: "student not found",
			args: args{
				uid: 2,
			},
			mockSqlStudentQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlStudentErrorReturned: gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnError(tt.mockSqlStudentErrorReturned)

			_, actual := repo.GetStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_UpdateExistingStudent(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid     uint32
		student *domain.Student
	}
	tests := []struct {
		name                              string
		args                              args
		mockSqlStudentSelectQueryExpected string
		mockStudentReturned               *Student
		mockSqlUserSelectQueryExpected    string
		mockUserReturned                  *User
		mockSqlUserInsertQueryExpected    string
		mockInsertedUserIdReturned        int
		mockUpdatedStudentIdReturned      int
		mockSqlStudentUpdateQueryExpected string
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE "students"."deleted_at" IS NULL ORDER BY "students"."id" LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockUserReturned: &User{
				Model:       gorm.Model{ID: 3},
				Username:    "username",
				Password:    "password",
				FirstName:   "firstName",
				LastName:    "lastName",
				Email:       "email@email",
				BirthDate:   time.Time{},
				PhoneNumber: "phoneNumber",
				Photo:       "photo",
			},
			mockSqlStudentUpdateQueryExpected: `UPDATE "students" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"registration_number"=$4,"user_id"=$5 WHERE "students"."deleted_at" IS NULL AND "id" = $6`,
			mockSqlUserInsertQueryExpected:    `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned:        1,
			mockUpdatedStudentIdReturned:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentSelectQueryExpected)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
				WithArgs(tt.mockStudentReturned.UserId).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"id", "username", "password", "first_name",
							"last_name", "email", "birth_date", "phone_number", "photo",
						},
					).AddRow(
						tt.mockUserReturned.ID, tt.mockUserReturned.Username, tt.mockUserReturned.Password,
						tt.mockUserReturned.FirstName, tt.mockUserReturned.LastName, tt.mockUserReturned.Email,
						tt.mockUserReturned.BirthDate, tt.mockUserReturned.PhoneNumber, tt.mockUserReturned.Photo,
					),
				)

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserInsertQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.Username, tt.args.student.Password, tt.args.student.FirstName,
					tt.args.student.LastName, tt.args.student.Email, tt.args.student.BirthDate, tt.args.student.PhoneNumber,
					tt.args.student.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlStudentUpdateQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.student.RegistrationNumber, tt.mockInsertedUserIdReturned,
					tt.args.uid,
				).
				WillReturnResult(
					sqlmock.NewResult(
						int64(tt.mockUpdatedStudentIdReturned),
						1,
					),
				)
			mockDb.ExpectCommit()

			err := repo.UpdateStudent(
				context.Background(),
				tt.args.uid,
				tt.args.student,
			)

			assert.NoError(t, err)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_UpdateExistingStudentReturnsStudentSelectError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid     uint32
		student *domain.Student
	}
	tests := []struct {
		name                              string
		args                              args
		mockSqlStudentSelectQueryExpected string
		mockSqlStudentErrorReturned       error
		expectedError                     error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlStudentErrorReturned:       errors.New("random error"),
			expectedError:                     errors.New("random error"),
		},
		{
			name: "student not found",
			args: args{
				uid: 2,
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockSqlStudentErrorReturned:       gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentSelectQueryExpected)).
				WithArgs().
				WillReturnError(tt.mockSqlStudentErrorReturned)

			_, actual := repo.GetStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestStudentRepository_UpdateExistingStudentReturnsUserSelectError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid     uint32
		student *domain.Student
	}
	tests := []struct {
		name                              string
		args                              args
		mockSqlStudentSelectQueryExpected string
		mockStudentReturned               *Student
		mockSqlUserSelectQueryExpected    string
		mockSqlUserErrorReturned          error
		expectedError                     error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned:       errors.New("random error"),
			expectedError:                  errors.New("random error"),
		},
		{
			name: "user not found",
			args: args{
				uid: 2,
				student: &domain.Student{
					User: domain.User{
						Username:    "username",
						Password:    "password",
						FirstName:   "firstName",
						LastName:    "lastName",
						Email:       "email@email",
						BirthDate:   time.Time{},
						PhoneNumber: "phoneNumber",
						Photo:       "photo",
					},
					RegistrationNumber: "registrationNumber",
				},
			},
			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE id = $1 AND "students"."deleted_at" IS NULL LIMIT 1`,
			mockStudentReturned: &Student{
				Model:              gorm.Model{ID: 2},
				RegistrationNumber: "registrationNumber",
				UserId:             3,
			},
			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned:       gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &StudentRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentSelectQueryExpected)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "registration_number", "user_id"},
					).AddRow(
						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
						tt.mockStudentReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
				WithArgs(tt.mockStudentReturned.UserId).
				WillReturnError(tt.mockSqlUserErrorReturned)

			_, actual := repo.GetStudent(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

// TODO: test cases for updateExistingStudent errors for updating student and user queries
////func TestStudentRepository_UpdateExistingStudentReturnsStudentUpdateError(t *testing.T) {
////	db, mockDb, err := sqlmock.New()
////	if err != nil {
////		t.Error(err.Error())
////	}
////	defer db.Close()
////
////	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
////
////	type args struct {
////		uid     uint32
////		student *domain.Student
////	}
////	tests := []struct {
////		name                              string
////		args                              args
////		mockSqlStudentSelectQueryExpected string
////		mockStudentReturned               *Student
////		mockSqlUserSelectQueryExpected    string
////		mockUserReturned                  *User
////		mockSqlStudentUpdateQueryExpected string
////		mockSqlStudentUpdateErrorReturned error
////		expectedError                     error
////	}{
////		{
////			name: "random error",
////			args: args{
////				uid: 2,
////				student: &domain.Student{
////					User: domain.User{
////						Username:    "username",
////						Password:    "password",
////						FirstName:   "firstName",
////						LastName:    "lastName",
////						Email:       "email@email",
////						BirthDate:   time.Time{},
////						PhoneNumber: "phoneNumber",
////						Photo:       "photo",
////					},
////					RegistrationNumber: "registrationNumber",
////				},
////			},
////			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE "students"."deleted_at" IS NULL ORDER BY "students"."id" LIMIT 1`,
////			mockStudentReturned: &Student{
////				Model:              gorm.Model{ID: 2},
////				RegistrationNumber: "registrationNumber",
////				UserId:             3,
////			},
////			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
////			mockUserReturned: &User{
////				Model:       gorm.Model{ID: 3},
////				Username:    "username",
////				Password:    "password",
////				FirstName:   "firstName",
////				LastName:    "lastName",
////				Email:       "email@email",
////				BirthDate:   time.Time{},
////				PhoneNumber: "phoneNumber",
////				Photo:       "photo",
////			},
////			mockSqlStudentUpdateQueryExpected: `UPDATE "students" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"registration_number"=$4,"user_id"=$5 WHERE "students"."deleted_at" IS NULL AND "id" = $6`,
////			mockSqlStudentUpdateErrorReturned: errors.New("random error"),
////		},
////	}
////
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			repo := &StudentRepository{
////				db: gormDb,
////			}
////
////			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentSelectQueryExpected)).
////				WithArgs().
////				WillReturnRows(
////					sqlmock.NewRows(
////						[]string{"id", "registration_number", "user_id"},
////					).AddRow(
////						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
////						tt.mockStudentReturned.UserId,
////					),
////				)
////
////			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
////				WithArgs(tt.mockStudentReturned.UserId).
////				WillReturnRows(
////					sqlmock.NewRows(
////						[]string{
////							"id", "username", "password", "first_name",
////							"last_name", "email", "birth_date", "phone_number", "photo",
////						},
////					).AddRow(
////						tt.mockUserReturned.ID, tt.mockUserReturned.Username, tt.mockUserReturned.Password,
////						tt.mockUserReturned.FirstName, tt.mockUserReturned.LastName, tt.mockUserReturned.Email,
////						tt.mockUserReturned.BirthDate, tt.mockUserReturned.PhoneNumber, tt.mockUserReturned.Photo,
////					),
////				)
////
////			mockDb.ExpectBegin()
////			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlStudentUpdateQueryExpected)).
////				WithArgs(
////					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
////					tt.args.student.RegistrationNumber, sqlmock.AnyArg(),
////					tt.args.uid,
////				).
////				WillReturnError(tt.mockSqlStudentUpdateErrorReturned)
////			mockDb.ExpectRollback()
////
////			err := repo.UpdateStudent(
////				context.Background(),
////				tt.args.uid,
////				tt.args.student,
////			)
////
////			assert.NoError(t, err)
////
////			mockDb.ExpectationsWereMet()
////		})
////	}
////}
////
//func TestStudentRepository_UpdateExistingStudentReturnsUserUpdateError(t *testing.T) {
//	db, mockDb, err := sqlmock.New()
//	if err != nil {
//		t.Error(err.Error())
//	}
//	defer db.Close()
//
//	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
//
//	type args struct {
//		uid     uint32
//		student *domain.Student
//	}
//	tests := []struct {
//		name                              string
//		args                              args
//		mockSqlStudentSelectQueryExpected string
//		mockStudentReturned               *Student
//		mockSqlUserSelectQueryExpected    string
//		mockUserReturned                  *User
//		mockSqlUserInsertQueryExpected    string
//		mockSqlUserErrorReturned          error
//		expectedError                     error
//	}{
//		{
//			name: "random error",
//			args: args{
//				uid: 2,
//				student: &domain.Student{
//					User: domain.User{
//						Username:    "username",
//						Password:    "password",
//						FirstName:   "firstName",
//						LastName:    "lastName",
//						Email:       "email@email",
//						BirthDate:   time.Time{},
//						PhoneNumber: "phoneNumber",
//						Photo:       "photo",
//					},
//					RegistrationNumber: "registrationNumber",
//				},
//			},
//			mockSqlStudentSelectQueryExpected: `SELECT * FROM "students" WHERE "students"."deleted_at" IS NULL ORDER BY "students"."id" LIMIT 1`,
//			mockStudentReturned: &Student{
//				Model:              gorm.Model{ID: 2},
//				RegistrationNumber: "registrationNumber",
//				UserId:             3,
//			},
//			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
//			mockUserReturned: &User{
//				Model:       gorm.Model{ID: 3},
//				Username:    "username",
//				Password:    "password",
//				FirstName:   "firstName",
//				LastName:    "lastName",
//				Email:       "email@email",
//				BirthDate:   time.Time{},
//				PhoneNumber: "phoneNumber",
//				Photo:       "photo",
//			},
//			mockSqlUserInsertQueryExpected: `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
//			mockSqlUserErrorReturned:       errors.New("random error"),
//			expectedError:                  errors.New("random error"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			repo := &StudentRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlStudentSelectQueryExpected)).
//				WithArgs().
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{"id", "registration_number", "user_id"},
//					).AddRow(
//						tt.mockStudentReturned.ID, tt.mockStudentReturned.RegistrationNumber,
//						tt.mockStudentReturned.UserId,
//					),
//				)
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
//				WithArgs(tt.mockStudentReturned.UserId).
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{
//							"id", "username", "password", "first_name",
//							"last_name", "email", "birth_date", "phone_number", "photo",
//						},
//					).AddRow(
//						tt.mockUserReturned.ID, tt.mockUserReturned.Username, tt.mockUserReturned.Password,
//						tt.mockUserReturned.FirstName, tt.mockUserReturned.LastName, tt.mockUserReturned.Email,
//						tt.mockUserReturned.BirthDate, tt.mockUserReturned.PhoneNumber, tt.mockUserReturned.Photo,
//					),
//				)
//
//			mockDb.ExpectBegin()
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserInsertQueryExpected)).
//				WithArgs(
//					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
//					tt.args.student.Username, tt.args.student.Password, tt.args.student.FirstName,
//					tt.args.student.LastName, tt.args.student.Email, tt.args.student.BirthDate, tt.args.student.PhoneNumber,
//					tt.args.student.Photo,
//				).
//				WillReturnError(tt.mockSqlUserErrorReturned)
//			mockDb.ExpectRollback()
//
//			err := repo.UpdateStudent(
//				context.Background(),
//				tt.args.uid,
//				tt.args.student,
//			)
//
//			assert.NoError(t, err)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
