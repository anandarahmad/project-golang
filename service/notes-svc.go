package service

import (
	"log"
	"notes-service/dto"
	"notes-service/entity"
	"notes-service/repository"

	"github.com/mashingan/smapping"
)

type NotesSvc interface {
	Create(notes dto.Create) entity.Note
	Update(notes dto.Update) entity.Note
	GetAll() []entity.Note
	CaribyID(id uint) entity.Note
	Delete(id uint)
}

type notesSvc struct {
	notesRepo repository.NotesRepo
}

func NewNotesSvc(notesRepo repository.NotesRepo) NotesSvc {
	return &notesSvc{
		notesRepo: notesRepo,
	}
}

func (service *notesSvc) Create(notes dto.Create) entity.Note {
	notesData := entity.Note{}
	err := smapping.FillStruct(&notesData, smapping.MapFields(&notes))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.notesRepo.Create(notesData)
	return res
}

func (service *notesSvc) Update(notes dto.Update) entity.Note {
	notesData := entity.Note{}
	err := smapping.FillStruct(&notesData, smapping.MapFields(&notes))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.notesRepo.Update(notesData)
	return res
}

func (service *notesSvc) GetAll() []entity.Note {
	return service.notesRepo.GetAll()
}

func (service *notesSvc) CaribyID(id uint) entity.Note {
	return service.notesRepo.CaribyID(id)
}

func (service *notesSvc) Delete(id uint) {
	service.notesRepo.Delete(id)
}
