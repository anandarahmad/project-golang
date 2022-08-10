package main

import (
	"html/template"
	"notes-service/config"
	"notes-service/controller"
	"notes-service/repository"
	"notes-service/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB             = config.SetupDB()
	notesRepo repository.NotesRepo = repository.NewNotesRepo(db)
	notesSvc  service.NotesSvc     = service.NewNotesSvc(notesRepo)
	notesCtr  controller.NotesCtr  = controller.NewNotesCtr(notesSvc)
	pageCtr   controller.PageCtr   = controller.NewPageCtr()
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to Load env")
	}

	defer config.CloseDB(db)
	r := gin.Default()
	notesRoute := r.Group("api/notes")
	{
		notesRoute.POST("create", notesCtr.Create)
		notesRoute.POST("update", notesCtr.Update)
		notesRoute.GET("all", notesCtr.GetAll)
		notesRoute.GET("detail/:id", notesCtr.CaribyID)
		notesRoute.GET("done/:id", notesCtr.Done)
		notesRoute.GET("undone/:id", notesCtr.UnDone)
		notesRoute.DELETE("delete/:id", notesCtr.Delete)
	}

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"add": func(x int, y int) int {
			return x + y
		},
	}

	r.SetFuncMap(funcMap)
	r.StaticFile("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*.html")
	pageRoute := r.Group("")
	{
		pageRoute.GET("", pageCtr.Index)
	}

	r.Run(":80")
}
