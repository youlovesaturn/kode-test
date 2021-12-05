package main

import (
	"fmt"
	"math"
)

type Note struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type NoteStore struct {
	notes  map[int]Note
	nextId int
}

func NewNoteStore() *NoteStore {
	ns := &NoteStore{}
	ns.notes = make(map[int]Note)
	ns.nextId = 0
	return ns
}

func (ns *NoteStore) CreateNote(text string) int {
	note := Note{
		Id:   ns.nextId,
		Text: text,
	}

	ns.notes[ns.nextId] = note
	ns.nextId++
	return note.Id
}

func (ns *NoteStore) GetNote(id int) (Note, error) {
	n, ok := ns.notes[id]
	if ok {
		return n, nil
	} else {
		return Note{}, fmt.Errorf("note with id = %d doesn't exists", id)
	}
}

func findLastNote(n map[int]Note) (max int) {
	max = math.MinInt
	for n := range n {
		if n > max {
			max = n
		}
	}
	return max
}

func findFirstNote(n map[int]Note) (min int) {
	min = math.MaxInt
	for n := range n {
		if n < min {
			min = n
		}
	}
	return min
}

func (ns *NoteStore) GetFirstNote() (Note, error) {
	n, ok := ns.notes[findFirstNote(ns.notes)]
	if ok {
		return n, nil
	} else {
		return Note{}, fmt.Errorf("notes don't exist yet")
	}
}

func (ns *NoteStore) GetLastNote() (Note, error) {
	n, ok := ns.notes[findLastNote(ns.notes)]
	if ok {
		return n, nil
	} else {
		return Note{}, fmt.Errorf("notes don't exist yet")
	}
}

func (ns *NoteStore) DeleteNote(id int) error {
	if _, ok := ns.notes[id]; !ok {
		return fmt.Errorf("note with id = %d doesn't exists", id)
	}

	delete(ns.notes, id)
	return nil
}

func (ns *NoteStore) DeleteAllNotes() error {
	ns.notes = make(map[int]Note)
	return nil
}

func (ns *NoteStore) GetAllNotes() []Note {
	allNotes := make([]Note, 0, len(ns.notes))
	for _, note := range ns.notes {
		allNotes = append(allNotes, note)
	}
	return allNotes
}
