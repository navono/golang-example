package repository

import (
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	mocket "github.com/selvatico/go-mocket"

	// "github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/suite"

	"golang-example/misc/gorm2/model"

	. "github.com/smartystreets/goconvey/convey"
)

// type Suite struct {
// 	suite.Suite
// 	DB *gorm.DB
//
// 	repo   Repository
// 	person *model.Person
// }

type (
	TestSuit struct {
		DB *gorm.DB

		repo   Repository
		person *model.Person
	}
)

func SetupTests() *TestSuit { // or *sql.DB
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	// GORM
	db, _ := gorm.Open(mocket.DriverName, "") // Can be any connection string
	// require.NoError(s.T(), err)

	// OR
	// Regular sql package usage
	// db, err := sql.Open(mocket.DriverName, "")

	return &TestSuit{
		DB:   db,
		repo: CreateRepository(db),
	}
}

func TestPerson(t *testing.T) {
	ts := SetupTests()
	Convey("People", t, func() {
		id := uuid.NewV4()
		Convey("Create", func() {
			var (
				name = "test-name"
			)

			commonReply := []map[string]interface{}{{"id": id, "name": name}}
			mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"person\"").WithID(int64(10)).WithReply(commonReply)
			// mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO　FOO").WithID(mockedId).WithReply(commonReply)

			err := ts.repo.Create(id, name)

			// id, err := InsertRecord(ts.DB)
			// So(id, ShouldEqual, mockedId)
			So(err, ShouldBeEmpty)
		})

		Convey("Get no record", func() {
			var (
				id = uuid.NewV4()
				// name = "test-name"
			)

			res, err := ts.repo.Get(id)
			mocket.Catcher.Reset().NewMock().WithQuery(regexp.QuoteMeta(
				`SELECT * FROM "person" WHERE (id = $1)`)).WithArgs(id.String())

			So(err, ShouldBeError, gorm.ErrRecordNotFound)
			So(res.Name, ShouldBeEmpty)
		})

		Convey("Get", func() {
			var (
				name = "test-name"
			)

			res, err := ts.repo.Get(id)

			// commonReply := []map[string]interface{}{{"id": id, "name": name}}
			mocket.Catcher.Reset().NewMock().WithQuery(regexp.QuoteMeta(
				`SELECT * FROM "person"`)).WithID(int64(10))

			So(err, ShouldBeEmpty)
			So(res.Name, ShouldEqual, name)
		})
	})
}

//
// func (s *Suite) SetupTests() *gorm.DB { // or *sql.DB
// 	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
// 	mocket.Catcher.Logging = true
// 	// GORM
// 	db, err := gorm.Open(mocket.DriverName, "") // Can be any connection string
// 	require.NoError(s.T(), err)
//
// 	s.DB = db
//
// 	// OR
// 	// Regular sql package usage
// 	// db, err := sql.Open(mocket.DriverName, "")
//
// 	s.repo = CreateRepository(s.DB)
//
// 	return db
// }

// func (s *Suite) AfterTest(_, _ string) {
// 	require.NoError(s.T(), s.mock.ExpectationsWereMet())
// }

// func (s *Suite) Test_repository_Get() {
// 	var (
// 		id   = uuid.NewV4()
// 		name = "test-name"
// 	)
//
// 	res, err := s.repo.Get(id)
//
// 	// s.mock.ExpectQuery(regexp.QuoteMeta(
// 	// 	`SELECT * FROM "person" WHERE (id = $1)`)).
// 	// 	WithArgs(id.String()).
// 	// 	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
// 	// 		AddRow(id.String(), name))
//
// 	require.NoError(s.T(), err)
// 	require.Nil(s.T(), deep.Equal(&model.Person{ID: id, Name: name}, res))
// }
//
// func (s *Suite) Test_repository_Create() {
// 	var (
// 		id   = uuid.NewV4()
// 		name = "test-name"
// 	)
//
// 	// s.mock.ExpectQuery(regexp.QuoteMeta(
// 	// 	`INSERT INTO "person" ("id","name")
// 	// 		VALUES ($1,$2) RETURNING "person"."id"`)).
// 	// 	WithArgs(id, name).
// 	// 	WillReturnRows(
// 	// 		sqlmock.NewRows([]string{"id"}).AddRow(id.String()))
//
// 	// s.mock.ExpectExec(`INSERT INTO "person"`).
// 	// 	WithArgs(id, name).
// 	// 	WillReturnResult(sqlmock.NewResult(1, 1))
//
// 	mocket.Catcher.Logging = true
//
// 	commonReply := []map[string]interface{}{{"id": id, "name": name}}
// 	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"person\"").WithID(int64(10)).WithReply(commonReply)
//
// 	// err := s.repo.Create(id, name)
//
// 	// require.NoError(s.T(), err)
// }
//
// func (s *Suite) Test_insertRecord() {
// 	var mockedId int64
// 	mockedId = 64
//
// 	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO　FOO").WithID(mockedId)
// 	returnedId, err := InsertRecord(s.DB)
// 	require.NoError(s.T(), err)
//
// 	// if int64(returnedId) != mockedId {
// 	// 	t.Fatalf("Last insert id not returned. Expected: [%v] , Got: [%v]", mockedId, returnedId)
// 	// require.NoError(s.T(), fmt.Errorf("a"))
// 	// }
// 	//
// 	// require.NoError(s.T(), err)
// }
