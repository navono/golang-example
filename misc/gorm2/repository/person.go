package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	"golang-example/misc/gorm2/model"
)

type Repository interface {
	Get(id uuid.UUID) (*model.Person, error)
	Create(id uuid.UUID, name string) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) Create(id uuid.UUID, name string) error {
	person := &model.Person{
		ID:   id,
		Name: name,
	}

	return p.DB.Create(person).Error
}

func (p *repo) Get(id uuid.UUID) (*model.Person, error) {
	person := new(model.Person)

	err := p.DB.Where("id = ?", id).Find(person).Error

	return person, err
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// func InsertRecord(db *sql.DB) int64 {
// 	res, err := db.Exec(`INSERT INTO foo VALUES("bar", ?))`, "value")
// 	if err != nil {
// 		return 0
// 	}
// 	id, _ := res.LastInsertId()
// 	return id
// }

type (
	Foo struct {
		ID   uint `gorm:"primary_key"`
		Name string
	}
)

func InsertRecord(db *gorm.DB) (uint, error) {
	f := Foo{
		Name: "foo",
	}

	db.Save(&f)

	var ff Foo
	db.Find(&Foo{
		Name: "foo",
	}).Find(&ff)

	return ff.ID, nil
}
