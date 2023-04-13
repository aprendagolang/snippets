package db

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repo *repository
	person *Person
}

func (rs *RepositorySuite) SetupSuite() {
	var (
		err error
	)

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN: "sqlmock_db_0",
		DriverName: "postgres",
		Conn: rs.conn,
		PreferSimpleProtocol: true,
	})

	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.repo = NewRepository(rs.DB)
	assert.IsType(rs.T(), &repository{}, rs.repo)

	rs.person = &Person{
		ID: uuid.New(),
		Name: "Tiago",
		Age: 32,
	}
}

func (rs *RepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *RepositorySuite) TestInsert() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`INSERT INTO "people" ("id","name","age") VALUES ($1,$2,$3)`)).
		WithArgs(
			rs.person.ID,
			rs.person.Name,
			rs.person.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	p, err := rs.repo.Insert(rs.person)
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.person, p)
}

func (rs *RepositorySuite) TestUpdate() {
	rs.person.Name = "Tiago Temporin"

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE "people" SET "name"=$1 WHERE "id" = $2`)).
		WithArgs(
			rs.person.Name,
			rs.person.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	p, err := rs.repo.Update(rs.person)
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.person, p)
}

func (rs *RepositorySuite) TestDelete() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(
		regexp.QuoteMeta(`DELETE FROM "people" WHERE "people"."id" = $1`)).
		WithArgs(rs.person.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	rs.mock.ExpectCommit()

	err := rs.repo.Delete(rs.person.ID)
	assert.NoError(rs.T(), err)
}

func (rs *RepositorySuite) TestPaginate() {
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).
		AddRow(
			rs.person.ID,
			rs.person.Name,
			rs.person.Age).
		AddRow(
			uuid.New(),
			"Maria Silva",
			27)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "people"`)).
		WithArgs().
		WillReturnRows(rows)

	people, err := rs.repo.Paginate()
	assert.NoError(rs.T(), err)
	assert.Contains(rs.T(), people, *rs.person)
}

func (rs *RepositorySuite) TestFindyByID() {
	rows := sqlmock.NewRows([]string{"id", "name", "age"}).
		AddRow(
			rs.person.ID,
			rs.person.Name,
			rs.person.Age)

	rs.mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "people" WHERE "people"."id" = $1`)).
		WithArgs(rs.person.ID).
		WillReturnRows(rows)

	p, err := rs.repo.FindByID(rs.person.ID)
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.person, p)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}