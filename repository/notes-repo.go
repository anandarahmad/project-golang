package repository

import (
	"notes-service/entity"

	"gorm.io/gorm"
)

type NotesRepo interface {
	Create(entity.Note) entity.Note
	Update(entity.Note) entity.Note
	GetAll() []entity.Note
	CaribyID(id uint) entity.Note
	Delete(id uint)
}

type dbConnect struct {
	conncetion *gorm.DB
}

func NewNotesRepo(db *gorm.DB) NotesRepo {
	return &dbConnect{
		conncetion: db,
	}
}

func (db *dbConnect) Create(entity entity.Note) entity.Note {
	db.conncetion.Save(&entity)
	return entity
}

func (db *dbConnect) Update(entity entity.Note) entity.Note {
	db.conncetion.Save(&entity)
	db.conncetion.Find(&entity)
	return entity
}

func (db *dbConnect) GetAll() []entity.Note {
	var notes []entity.Note
	db.conncetion.Order("created_at desc").Find(&notes)
	return notes
}

func (db *dbConnect) CaribyID(id uint) entity.Note {
	var notes entity.Note
	db.conncetion.Where("id = ?", id).Find(&notes)
	return notes
}

func (db *dbConnect) Delete(id uint) {
	var notes entity.Note
	db.conncetion.Where("id =?", id).Delete(&notes)
}
