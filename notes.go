package main

import (
	"slices"
)

// "github.com/gin-gonic/gin"

// the below struct is used for one note.

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// the below note is used to store all the notes created through the Note struct.
type NotesManager struct {
	notes  []Note // storest every single note in []notes.
	nextID int    // nextID is used to provide the id for new note.
}

// the below function is used to create notes.
func (m *NotesManager) CreateNotes(title, content string) Note {
	// the variable "NewNote" stores the the new note and appends to the NotesManager notes[]
	NewNote := Note{
		ID:      m.nextID,
		Title:   title,
		Content: content,
	}

	// the below line appends the NewNote to m.notes.
	m.notes = append(m.notes, NewNote)
	// increments the ID for next new note.
	m.nextID++

	// returns the NewNote to see which note was created.
	return NewNote
}

// the below function lists all notes.
func (m *NotesManager) ListNotes() []Note {
	return m.notes
}

// The below function is used to get a perticular note for the in-memory.
func (m *NotesManager) GetNote(id int) (Note, bool) {
	// loops till the the length of the slice "m.notes".
	for i := 0; i < len(m.notes); i++ {
		//the len is used here cuz we need to search all the notes inside the slice.
		if m.notes[i].ID == id { // it compares the ID in the index "i" with id.
			return m.notes[i], true
		}
	}
	return Note{}, false // returns an empty Note struct if it couldnt find the note with id.
}

func (m *NotesManager) UpdateNotes(id int, title, content string) bool {
	for i := 0; i < len(m.notes); i++ {
		//the len is used here cuz we need to search all the notes inside the slice.
		if m.notes[i].ID == id { // it compares the ID in the index "i" with id.
			m.notes[i].Title = title // updates the title and content to the new one.
			m.notes[i].Content = content
			return true
		}
	}

	return false

}

// same if idnetifing the note but uses slices.Delete to delete the note.
func (m *NotesManager) DeleteNote(id int) bool {
	for i := 0; i < len(m.notes); i++ {
		if m.notes[i].ID == id {
			m.notes = slices.Delete(m.notes, i, i+1)
			return true
		}
	}
	return false
}
