package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	store *NoteStore
}

type RequestNote struct {
	Text string `json:"text"`
}

type ResponseId struct {
	Id int `json:"id"`
}

func NewServer() *Server {
	store := NewNoteStore()
	return &Server{store: store}
}

func getJSON(w http.ResponseWriter, k interface{}) {
	js, err := json.Marshal(k)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

func (s *Server) noteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/note/" {
		switch r.Method {
		case http.MethodPost:
			s.createNoteHandler(w, r)
		case http.MethodGet:
			s.getAllNotesHandler(w, r)
		case http.MethodDelete:
			s.deleteAllNotesHandler(w, r)
		default:
			http.Error(w, fmt.Sprintf("except GET, DELETE and POST, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		if len(parts) < 2 {
			http.Error(w, "except /note/<id>", http.StatusBadRequest)
			return
		}
		for _, e := range parts {
			if e == "first" {
				s.getFirstNoteHandler(w, r)
				return
			} else if e == "last" {
				s.getLastNoteHandler(w, r)
				return
			}
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodDelete:
			s.deleteNoteHandler(w, r, int(id))
		case http.MethodGet:
			s.getNoteHandler(w, r, int(id))
		default:
			http.Error(w, fmt.Sprintf("except GET, DELETE and POST, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (s *Server) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var rn RequestNote
	if err := dec.Decode(&rn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := s.store.CreateNote(rn.Text)
	getJSON(w, ResponseId{Id: id})
}

func (s *Server) getNoteHandler(w http.ResponseWriter, r *http.Request, id int) {
	note, err := s.store.GetNote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	getJSON(w, note)
}

func (s *Server) getFirstNoteHandler(w http.ResponseWriter, r *http.Request) {
	note, err := s.store.GetFirstNote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	getJSON(w, note)
}

func (s *Server) getLastNoteHandler(w http.ResponseWriter, r *http.Request) {
	note, err := s.store.GetLastNote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	getJSON(w, note)
}

func (s *Server) getAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	allNotes := s.store.GetAllNotes()
	getJSON(w, allNotes)
}

func (s *Server) deleteNoteHandler(w http.ResponseWriter, r *http.Request, id int) {
	err := s.store.DeleteNote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (s *Server) deleteAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	s.store.DeleteAllNotes()
}
