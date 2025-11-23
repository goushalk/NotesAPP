package main

import (
// "github.com/gin-gonic/gin"
)

type Note struct {
	ID      int
	title   string
	content string
}

type NotesManager struct {
	notes  []Note
	nextID int
}

func (m *NotesManager) Create(title, content string) Note {}

func (m *NotesManager) List() []Note {}

func (m *NotesManager) Get(id int) (Note, bool) {}

func (m *NotesManager) Update(id int, title, content string) bool {}

func (m *NotesManager) Delete(id int) bool {}
