package controller

import (
	"net/http"
	"notes-service/dto"
	"notes-service/entity"
	"notes-service/helper"
	"notes-service/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type NotesCtr interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Done(ctx *gin.Context)
	UnDone(ctx *gin.Context)
	CaribyID(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type notesCtr struct {
	notesSvc service.NotesSvc
}

func NewNotesCtr(notesSvc service.NotesSvc) NotesCtr {
	return &notesCtr{
		notesSvc: notesSvc,
	}
}

func (c *notesCtr) Create(ctx *gin.Context) {
	var dataNotes dto.Create
	errDTO := ctx.ShouldBind(&dataNotes)
	if errDTO != nil {
		response := helper.BuildErrResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	dataCreate := c.notesSvc.Create(dataNotes)
	response := helper.BuildResponse(true, "Data berhasil dibuat", dataCreate)
	ctx.JSON(http.StatusCreated, response)
}

func (c *notesCtr) Update(ctx *gin.Context) {
	var dataNotes dto.Update
	errDTO := ctx.ShouldBind(&dataNotes)
	if errDTO != nil {
		response := helper.BuildErrResponse("Gagal memproses permintaan", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	cekData := c.notesSvc.CaribyID(dataNotes.ID)
	if (cekData == entity.Note{}) {
		response := helper.BuildErrResponse("Gagal memproses permintaan", "Data tidak ditemukan", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	dataNotes.CreatedAt = cekData.CreatedAt
	dataNotes.IsDone = cekData.IsDone
	dataNotes.UpdatedAt = time.Now()

	dataUpdate := c.notesSvc.Update(dataNotes)
	response := helper.BuildResponse(true, "Data berhasil diubah", dataUpdate)
	ctx.JSON(http.StatusCreated, response)
}

func (c *notesCtr) GetAll(ctx *gin.Context) {
	getData := c.notesSvc.GetAll()
	response := helper.BuildResponse(true, "Data berhasil ditampilkan", getData)
	ctx.JSON(http.StatusOK, response)
}

func (c *notesCtr) CaribyID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrResponse("Parameter tidak ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	getData := c.notesSvc.CaribyID(uint(id))
	if (getData == entity.Note{}) {
		response := helper.BuildErrResponse("Gagal memproses permintaan", "Data tidak ditemukan", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse(true, "Data berhasil ditampilkan", getData)
	ctx.JSON(http.StatusOK, response)
}

func (c *notesCtr) Done(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrResponse("Parameter tidak ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	cekData := c.notesSvc.CaribyID(uint(id))
	if (cekData == entity.Note{}) {
		response := helper.BuildErrResponse("Gagal memproses permintaan", "Data tidak ditemukan", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var dataNotes dto.Update
	smapping.FillStruct(&dataNotes, smapping.MapFields(&cekData))
	dataNotes.IsDone = true
	dataNotes.UpdatedAt = time.Now()
	dataUpdate := c.notesSvc.Update(dataNotes)
	response := helper.BuildResponse(true, "Data berhasil diubah menjadi DONE", dataUpdate)
	ctx.JSON(http.StatusCreated, response)

}

func (c *notesCtr) UnDone(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrResponse("Parameter tidak ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	cekData := c.notesSvc.CaribyID(uint(id))
	if (cekData == entity.Note{}) {
		response := helper.BuildErrResponse("Gagal memproses permintaan", "Data tidak ditemukan", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var dataNotes dto.Update
	smapping.FillStruct(&dataNotes, smapping.MapFields(&cekData))
	dataNotes.IsDone = false
	dataNotes.UpdatedAt = time.Now()
	dataUpdate := c.notesSvc.Update(dataNotes)
	response := helper.BuildResponse(true, "Data berhasil diubah menjadi UNDONE", dataUpdate)
	ctx.JSON(http.StatusCreated, response)

}

func (c *notesCtr) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrResponse("Parameter tidak ditemukan", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	cekData := c.notesSvc.CaribyID(uint(id))
	if (cekData == entity.Note{}) {
		response := helper.BuildErrResponse("Gagal memproses permintaan", "Data tidak ditemukan", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.notesSvc.Delete(uint(id))

	response := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	ctx.JSON(http.StatusCreated, response)
}
