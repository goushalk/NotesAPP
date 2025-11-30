package main

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // gin default router

	router.LoadHTMLGlob("templates/*.html") // loads html file in patteren applied globaly

	router.GET("/", func(ctx *gin.Context) { // gin.Context is handler resposible to handle all the http request.
		ctx.HTML(200, "index.html", gin.H{ // loads the html from the templates with a status code of "200", gin.H is used to handel the JSON.
			"title": "Notes App", // updates the html dynamicly.
		})
	})

	// the the below function is about loading a html file for dashboard UI.

	router.GET("/dashboard", func(ctx *gin.Context) {
		ctx.HTML(200, "dashboard.html", gin.H{})
	})

	// this manager variable has a type of NotesManager form notes.go.
	var manager NotesManager

	// this below is POST function it is used to create a note.
	router.POST("/CreateNote", func(ctx *gin.Context) {
		var UserInput Note                                     // this variable is used to store the JSON inpu to Note struct from notes.go .
		if err := ctx.ShouldBindJSON(&UserInput); err != nil { // this is resposible for converting JSON value to the value go needs.
			ctx.JSON(400, gin.H{ // if the err variable is equal to nil then it returns a error in JSON format.
				"error": err.Error(),
			})
			return
		}
		noteCreate := manager.CreateNotes(UserInput.Title, UserInput.Content) // this line is responsible for calling the function to create the notes.
		ctx.JSON(201, noteCreate)                                             // it converts the noteCreate data into JSON format and sends it back to the client with an HTTP 201 Created status code.

	})

	router.GET("/notes", func(ctx *gin.Context) {
		allNotes := manager.ListNotes()
		ctx.JSON(200, allNotes)
	})

	router.DELETE("/notes/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		noteID, _ := strconv.Atoi(id)
		deleteNotes := manager.DeleteNote(noteID)

		if deleteNotes == true {
			ctx.JSON(200, gin.H{
				"deleted": true,
			})
		} else {
			ctx.JSON(404, gin.H{
				"error": "page not found",
			})
		}
	})

	if err := router.Run(); err != nil {
		log.Error(err)
	}

}
