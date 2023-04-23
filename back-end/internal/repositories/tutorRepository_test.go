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

func TestTutorRepository_CreateTutor(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		tutor *domain.Tutor
	}
	tests := []struct {
		name                        string
		args                        args
		mockSqlUserQueryExpected    string
		mockSqlTutorQueryExpected   string
		mockInsertedTutorIdReturned int
		mockInsertedUserIdReturned  int
		expectedTutorUid            uint
	}{
		{
			name: "valid",
			args: args{
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorQueryExpected:   `INSERT INTO "tutors" ("created_at","updated_at","deleted_at","academic_rank","user_id") VALUES ($1,$2,$3,$4,$5)`,
			mockSqlUserQueryExpected:    `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned:  1,
			mockInsertedTutorIdReturned: 2,
			expectedTutorUid:            2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.Username, tt.args.tutor.Password, tt.args.tutor.FirstName,
					tt.args.tutor.LastName, tt.args.tutor.Email, tt.args.tutor.BirthDate, tt.args.tutor.PhoneNumber,
					tt.args.tutor.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.AcademicRank, tt.mockInsertedUserIdReturned,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedTutorIdReturned))
			mockDb.ExpectCommit()

			actual, err := repo.CreateTutor(context.Background(), tt.args.tutor)
			if err != nil {
				t.Errorf("CreateTutor() error = %v", err)
			}

			assert.Equal(t, tt.expectedTutorUid, actual)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestTutorRepository_CreateTutorReturnsErrorOnTutorCreate(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		tutor *domain.Tutor
	}
	tests := []struct {
		name                       string
		args                       args
		mockSqlUserQueryExpected   string
		mockInsertedUserIdReturned int
		mockSqlTutorQueryExpected  string
		mockSqlTutorQueryError     error
		expectedError              error
	}{
		{
			name: "random error",
			args: args{
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlUserQueryExpected:   `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned: 1,
			mockSqlTutorQueryExpected:  `INSERT INTO "tutors" ("created_at","updated_at","deleted_at","academic_rank","user_id") VALUES ($1,$2,$3,$4,$5)`,
			mockSqlTutorQueryError:     errors.New("random error"),
			expectedError:              errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.Username, tt.args.tutor.Password, tt.args.tutor.FirstName,
					tt.args.tutor.LastName, tt.args.tutor.Email, tt.args.tutor.BirthDate, tt.args.tutor.PhoneNumber,
					tt.args.tutor.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.AcademicRank, tt.mockInsertedUserIdReturned,
				).
				WillReturnError(tt.mockSqlTutorQueryError)
			mockDb.ExpectRollback()

			_, actualError := repo.CreateTutor(context.Background(), tt.args.tutor)

			assert.Equal(t, tt.expectedError, actualError)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestTutorRepository_CreateTutorReturnsErrorOnUserCreate(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		tutor *domain.Tutor
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
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlUserQueryExpected: `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockSqlUserQueryError:    errors.New("random error"),
			expectedError:            errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectBegin()
			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.Username, tt.args.tutor.Password, tt.args.tutor.FirstName,
					tt.args.tutor.LastName, tt.args.tutor.Email, tt.args.tutor.BirthDate, tt.args.tutor.PhoneNumber,
					tt.args.tutor.Photo,
				).
				WillReturnError(tt.mockSqlUserQueryError)
			mockDb.ExpectRollback()

			_, actualError := repo.CreateTutor(context.Background(), tt.args.tutor)

			assert.Equal(t, tt.expectedError, actualError)

			if err = mockDb.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expections: %s", err)
			}
		})
	}
}

func TestTutorRepository_CreateTutorReturnsErrorOnUserValidation(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		tutor *domain.Tutor
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "missing username",
			args: args{
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
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
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
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
			repo := &TutorRepository{
				db: gormDb,
			}

			_, actualError := repo.CreateTutor(context.Background(), tt.args.tutor)

			assert.Equal(t, tt.expectedError, actualError)
		})
	}
}

func TestTutorRepository_DeleteTutor(t *testing.T) {
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
		name                            string
		args                            args
		mockSqlSelectTutorQueryExpected string
		mockSqlDeleteTutorQueryExpected string
		mockTutorReturned               *Tutor
		mockDeletedTutorIdReturned      int
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
			},
			mockSqlSelectTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			mockSqlDeleteTutorQueryExpected: `UPDATE "tutors" SET "deleted_at"=$1 WHERE id = $2 AND "tutors"."deleted_at" IS NULL`,
			mockDeletedTutorIdReturned:      2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)
			mockDb.ExpectBegin()
			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteTutorQueryExpected)).
				WithArgs(sqlmock.AnyArg(), tt.args.uid).
				WillReturnResult(sqlmock.NewResult(int64(tt.mockDeletedTutorIdReturned), 1))
			mockDb.ExpectCommit()

			err := repo.DeleteTutor(context.Background(), tt.args.uid)

			assert.NoError(t, err)

			mockDb.ExpectationsWereMet()
		})
	}
}

// The delete action in the gorm is done by selecting the requested tutor and then
// soft delete it, so we have two possible error sources: the select and the update
func TestTutorRepository_DeleteTutorWithSelectError(t *testing.T) {
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
		name                            string
		args                            args
		mockSqlSelectTutorQueryExpected string
		mockSqlErrorReturned            error
		expectedError                   error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlSelectTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlErrorReturned:            errors.New("random error"),
			expectedError:                   errors.New("random error"),
		},
		{
			name: "tutor not found",
			args: args{
				uid: 2,
			},
			mockSqlSelectTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlErrorReturned:            gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnError(tt.mockSqlErrorReturned)

			actual := repo.DeleteTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_DeleteTutorWithUpdateError(t *testing.T) {
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
		name                            string
		args                            args
		mockSqlSelectTutorQueryExpected string
		mockSqlDeleteTutorQueryExpected string
		mockTutorReturned               *Tutor
		mockSqlErrorReturned            error
		expectedError                   error
	}{
		{
			name: "random error",
			args: args{
				uid: 666,
			},
			mockSqlSelectTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			mockSqlDeleteTutorQueryExpected: `UPDATE "tutors" SET "deleted_at"=$1 WHERE id = $2 AND "tutors"."deleted_at" IS NULL`,
			mockSqlErrorReturned:            errors.New("random error"),
			expectedError:                   errors.New("random error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlSelectTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)
			mockDb.ExpectBegin()
			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlDeleteTutorQueryExpected)).
				WithArgs(sqlmock.AnyArg(), tt.args.uid).
				WillReturnError(tt.mockSqlErrorReturned)
			mockDb.ExpectRollback()

			actual := repo.DeleteTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_GetTutor(t *testing.T) {
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
		name                      string
		args                      args
		mockSqlUserQueryExpected  string
		mockUserReturned          *User
		mockSqlTutorQueryExpected string
		mockTutorReturned         *Tutor
		expected                  *domain.Tutor
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
			},
			mockSqlTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			expected: &domain.Tutor{
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
				AcademicRank: "academicRank",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(tt.mockTutorReturned.UserId).
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

			actual, err := repo.GetTutor(context.Background(), tt.args.uid)
			if err != nil {
				t.Errorf("GetTutor() error = %v", err)
				return
			}

			assert.Equal(t, tt.expected, actual)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_GetTutorReturnsErrorOnUserSelect(t *testing.T) {
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
		name                      string
		args                      args
		mockSqlTutorQueryExpected string
		mockTutorReturned         *Tutor
		mockSqlUserQueryExpected  string
		mockSqlUserErrorReturned  error
		expectedError             error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			mockSqlTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
			},
			mockSqlUserQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned: gorm.ErrRecordNotFound,
			// The error is that we didnt find a uid 2 even though the Tutor with
			// id 2 exists. It's because the Tutor had not an associated user with uid 2
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserQueryExpected)).
				WithArgs(tt.mockTutorReturned.UserId).
				WillReturnError(tt.mockSqlUserErrorReturned)

			_, actual := repo.GetTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_GetTutorReturnsErrorOnTutorSelect(t *testing.T) {
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
		name                      string
		args                      args
		mockSqlTutorQueryExpected string
		mockSqlTutorErrorReturned error
		expectedError             error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
			},
			mockSqlTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlTutorErrorReturned: errors.New("random error"),
			expectedError:             errors.New("random error"),
		},
		{
			name: "tutor not found",
			args: args{
				uid: 2,
			},
			mockSqlTutorQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlTutorErrorReturned: gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorQueryExpected)).
				WithArgs(tt.args.uid).
				WillReturnError(tt.mockSqlTutorErrorReturned)

			_, actual := repo.GetTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_UpdateExistingTutor(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid   uint32
		tutor *domain.Tutor
	}
	tests := []struct {
		name                            string
		args                            args
		mockSqlTutorSelectQueryExpected string
		mockTutorReturned               *Tutor
		mockSqlUserSelectQueryExpected  string
		mockUserReturned                *User
		mockSqlUserInsertQueryExpected  string
		mockInsertedUserIdReturned      int
		mockUpdatedTutorIdReturned      int
		mockSqlTutorUpdateQueryExpected string
	}{
		{
			name: "valid",
			args: args{
				uid: 2,
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE "tutors"."deleted_at" IS NULL ORDER BY "tutors"."id" LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			mockSqlTutorUpdateQueryExpected: `UPDATE "tutors" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"academic_rank"=$4,"user_id"=$5 WHERE "tutors"."deleted_at" IS NULL AND "id" = $6`,
			mockSqlUserInsertQueryExpected:  `INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","first_name","last_name","email","birth_date","phone_number","photo") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			mockInsertedUserIdReturned:      1,
			mockUpdatedTutorIdReturned:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorSelectQueryExpected)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
				WithArgs(tt.mockTutorReturned.UserId).
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
					tt.args.tutor.Username, tt.args.tutor.Password, tt.args.tutor.FirstName,
					tt.args.tutor.LastName, tt.args.tutor.Email, tt.args.tutor.BirthDate, tt.args.tutor.PhoneNumber,
					tt.args.tutor.Photo,
				).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(tt.mockInsertedUserIdReturned))

			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlTutorUpdateQueryExpected)).
				WithArgs(
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					tt.args.tutor.AcademicRank, tt.mockInsertedUserIdReturned,
					tt.args.uid,
				).
				WillReturnResult(
					sqlmock.NewResult(
						int64(tt.mockUpdatedTutorIdReturned),
						1,
					),
				)
			mockDb.ExpectCommit()

			err := repo.UpdateTutor(
				context.Background(),
				tt.args.uid,
				tt.args.tutor,
			)

			assert.NoError(t, err)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_UpdateExistingTutorReturnsTutorSelectError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid   uint32
		tutor *domain.Tutor
	}
	tests := []struct {
		name                            string
		args                            args
		mockSqlTutorSelectQueryExpected string
		mockSqlTutorErrorReturned       error
		expectedError                   error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlTutorErrorReturned:       errors.New("random error"),
			expectedError:                   errors.New("random error"),
		},
		{
			name: "tutor not found",
			args: args{
				uid: 2,
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockSqlTutorErrorReturned:       gorm.ErrRecordNotFound,
			expectedError: apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("uid 2 not found"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorSelectQueryExpected)).
				WithArgs().
				WillReturnError(tt.mockSqlTutorErrorReturned)

			_, actual := repo.GetTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

func TestTutorRepository_UpdateExistingTutorReturnsUserSelectError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}
	defer db.Close()

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))

	type args struct {
		uid   uint32
		tutor *domain.Tutor
	}
	tests := []struct {
		name                            string
		args                            args
		mockSqlTutorSelectQueryExpected string
		mockTutorReturned               *Tutor
		mockSqlUserSelectQueryExpected  string
		mockSqlUserErrorReturned        error
		expectedError                   error
	}{
		{
			name: "random error",
			args: args{
				uid: 2,
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
			},
			mockSqlUserSelectQueryExpected: `SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL`,
			mockSqlUserErrorReturned:       errors.New("random error"),
			expectedError:                  errors.New("random error"),
		},
		{
			name: "user not found",
			args: args{
				uid: 2,
				tutor: &domain.Tutor{
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
					AcademicRank: "academicRank",
				},
			},
			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE id = $1 AND "tutors"."deleted_at" IS NULL LIMIT 1`,
			mockTutorReturned: &Tutor{
				Model:        gorm.Model{ID: 2},
				AcademicRank: "academicRank",
				UserId:       3,
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
			repo := &TutorRepository{
				db: gormDb,
			}

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorSelectQueryExpected)).
				WithArgs().
				WillReturnRows(
					sqlmock.NewRows(
						[]string{"id", "academic_rank", "user_id"},
					).AddRow(
						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
						tt.mockTutorReturned.UserId,
					),
				)

			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
				WithArgs(tt.mockTutorReturned.UserId).
				WillReturnError(tt.mockSqlUserErrorReturned)

			_, actual := repo.GetTutor(context.Background(), tt.args.uid)

			assert.Equal(t, actual, tt.expectedError)

			mockDb.ExpectationsWereMet()
		})
	}
}

// TODO: test cases for updateExistingTutor errors for updating tutor and user queries
////func TestTutorRepository_UpdateExistingTutorReturnsTutorUpdateError(t *testing.T) {
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
////		tutor *domain.Tutor
////	}
////	tests := []struct {
////		name                              string
////		args                              args
////		mockSqlTutorSelectQueryExpected string
////		mockTutorReturned               *Tutor
////		mockSqlUserSelectQueryExpected    string
////		mockUserReturned                  *User
////		mockSqlTutorUpdateQueryExpected string
////		mockSqlTutorUpdateErrorReturned error
////		expectedError                     error
////	}{
////		{
////			name: "random error",
////			args: args{
////				uid: 2,
////				tutor: &domain.Tutor{
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
////					AcademicRank: "academicRank",
////				},
////			},
////			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE "tutors"."deleted_at" IS NULL ORDER BY "tutors"."id" LIMIT 1`,
////			mockTutorReturned: &Tutor{
////				Model:              gorm.Model{ID: 2},
////				AcademicRank: "academicRank",
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
////			mockSqlTutorUpdateQueryExpected: `UPDATE "tutors" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"academic_rank"=$4,"user_id"=$5 WHERE "tutors"."deleted_at" IS NULL AND "id" = $6`,
////			mockSqlTutorUpdateErrorReturned: errors.New("random error"),
////		},
////	}
////
////	for _, tt := range tests {
////		t.Run(tt.name, func(t *testing.T) {
////			repo := &TutorRepository{
////				db: gormDb,
////			}
////
////			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorSelectQueryExpected)).
////				WithArgs().
////				WillReturnRows(
////					sqlmock.NewRows(
////						[]string{"id", "academic_rank", "user_id"},
////					).AddRow(
////						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
////						tt.mockTutorReturned.UserId,
////					),
////				)
////
////			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
////				WithArgs(tt.mockTutorReturned.UserId).
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
////			mockDb.ExpectExec(regexp.QuoteMeta(tt.mockSqlTutorUpdateQueryExpected)).
////				WithArgs(
////					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
////					tt.args.tutor.AcademicRank, sqlmock.AnyArg(),
////					tt.args.uid,
////				).
////				WillReturnError(tt.mockSqlTutorUpdateErrorReturned)
////			mockDb.ExpectRollback()
////
////			err := repo.UpdateTutor(
////				context.Background(),
////				tt.args.uid,
////				tt.args.tutor,
////			)
////
////			assert.NoError(t, err)
////
////			mockDb.ExpectationsWereMet()
////		})
////	}
////}
////
//func TestTutorRepository_UpdateExistingTutorReturnsUserUpdateError(t *testing.T) {
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
//		tutor *domain.Tutor
//	}
//	tests := []struct {
//		name                              string
//		args                              args
//		mockSqlTutorSelectQueryExpected string
//		mockTutorReturned               *Tutor
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
//				tutor: &domain.Tutor{
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
//					AcademicRank: "academicRank",
//				},
//			},
//			mockSqlTutorSelectQueryExpected: `SELECT * FROM "tutors" WHERE "tutors"."deleted_at" IS NULL ORDER BY "tutors"."id" LIMIT 1`,
//			mockTutorReturned: &Tutor{
//				Model:              gorm.Model{ID: 2},
//				AcademicRank: "academicRank",
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
//			repo := &TutorRepository{
//				db: gormDb,
//			}
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlTutorSelectQueryExpected)).
//				WithArgs().
//				WillReturnRows(
//					sqlmock.NewRows(
//						[]string{"id", "academic_rank", "user_id"},
//					).AddRow(
//						tt.mockTutorReturned.ID, tt.mockTutorReturned.AcademicRank,
//						tt.mockTutorReturned.UserId,
//					),
//				)
//
//			mockDb.ExpectQuery(regexp.QuoteMeta(tt.mockSqlUserSelectQueryExpected)).
//				WithArgs(tt.mockTutorReturned.UserId).
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
//					tt.args.tutor.Username, tt.args.tutor.Password, tt.args.tutor.FirstName,
//					tt.args.tutor.LastName, tt.args.tutor.Email, tt.args.tutor.BirthDate, tt.args.tutor.PhoneNumber,
//					tt.args.tutor.Photo,
//				).
//				WillReturnError(tt.mockSqlUserErrorReturned)
//			mockDb.ExpectRollback()
//
//			err := repo.UpdateTutor(
//				context.Background(),
//				tt.args.uid,
//				tt.args.tutor,
//			)
//
//			assert.NoError(t, err)
//
//			mockDb.ExpectationsWereMet()
//		})
//	}
//}
