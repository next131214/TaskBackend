package database

import (
	"log"
	"strconv"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID     string `gorm:"column:id;primary_key"`
	Text   string `gorm:"column:text"`
	Done   bool   `gorm:"column:done"`
	UserID string `gorm:"column:user_id"`
}

func (u *Todo) TableName() string {
	return "todo"
}

type TodoDao interface {
	InsertOne(u *Todo) error
	UpdateOne(u *Todo) error
	DeleteOne(u *Todo) error
	FindAll() ([]*Todo, error)
	FindByUserID(userID string) ([]*Todo, error)
	FindOne(id string) (*Todo, error)
}

type todoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return &todoDao{db: db}
}
func (d *todoDao) DeleteOne(u *Todo) error {
	var sql string = "DELETE FROM todo WHERE id = '" + u.ID + "'"
	log.Printf(sql)
	log.Printf("%s\n", u)
	res := d.db.Exec(sql)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
func (d *todoDao) UpdateOne(u *Todo) error {
	var sql string = "UPDATE todo SET text='" + u.Text + "',done = " + strconv.FormatBool(u.Done) + " WHERE id = '" + u.ID + "'"
	log.Printf(sql)
	log.Printf("%s\n", u)
	res := d.db.Exec(sql)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
func (d *todoDao) InsertOne(u *Todo) error {
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
func (d *todoDao) FindAll() ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (d *todoDao) FindOne(id string) (*Todo, error) {
	var todos []*Todo
	res := d.db.Where("id = ?", id).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(todos) < 1 {
		return nil, nil
	}
	return todos[0], nil
}

func (d *todoDao) FindByUserID(userID string) ([]*Todo, error) {
	var todos []*Todo
	res := d.db.Where("user_id = ?", userID).Find(&todos)
	if err := res.Error; err != nil {
		return nil, err
	}
	return todos, nil
}
