package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/RamiroCuenca/go-rest-notesApi/common/handler"
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/database/connection"
	"github.com/RamiroCuenca/go-rest-notesApi/notes/models"
)

// This handler creates a new note
func NotesCreate(w http.ResponseWriter, r *http.Request) {
	// 1° Decode the json received on a Note object
	n := models.Note{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		logger.Log().Infof("Error decoding request: %v", err)
		handler.SendError(w, http.StatusBadRequest)
		return
	}

	// 2° Create the sql statement and prepare null fields
	q := `INSERT INTO notes (owner_name, title, details)
	 	VALUES ($1, $2, $3) RETURNING id`

	// A time ago... i used to open the database here, but at least on this
	// particular project we open it on the main file so it is not necessary
	// to be opened here

	// 3° Start a transaction
	db := connection.NewPostgresClient()
	tx, err := db.Begin()
	if err != nil {
		logger.Log().Infof("Error starting transaction: %v", err)
		handler.SendError(w, 500) // Internal Server Error
		return
	}

	// 4° Prepare the transaction
	stmt, err := tx.Prepare(q)
	if err != nil {
		logger.Log().Infof("Error preparing transaction: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}
	defer stmt.Close()

	// 5° Execute the query and assign the returned id to the Note object
	// We will use QueryRow because the exec method returns two methods that are
	// not compatible with psql!
	err = stmt.QueryRow(
		n.OwnerName,
		n.Title,
		stringToNull(n.Details), // Send a nill if it's null
	).Scan(&n.ID)
	if err != nil {
		logger.Log().Infof("Error executing query: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}

	// 6° As there are no errors, commit the transaction
	tx.Commit()
	logger.Log().Infof("Note created successfully! :)")

	// 7° Encode the Note into a JSON object
	json, _ := json.Marshal(n)

	// 8° Send the response
	handler.SendResponse(w, http.StatusCreated, json)

}

// This function manages the null string values
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}

	if null.String != "" {
		null.Valid = true
	}

	return null
}

// func (db *Database) NotesGetAll() ([]models.Note, error) {
// 	// 1° Prepare the query that will be executed
// 	q := `SELECT id, owner_name, title, details
// 		created_at, updated_at FROM notes`

// 	// Be prepared if details or updated at are empty
// 	// detailsNull := sql.NullString{}

// 	// 2° Open a transaction
// 	tx, err := db.Begin()
// 	if err != nil {
// 		logger.Log().Infof("There was an error beggining the tx. Error: %v", err)
// 		return nil, err
// 	}

// 	// 3° Prepare the query
// 	stmt, err := tx.Prepare(q)
// 	if err != nil {
// 		logger.Log().Infof("There was an error preparing the sql statement. Error: %v", err)
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	// 4° Execute the query. Send the parameters
// 	// We scan it and assign the returned id to the object
// 	// StringNull == note.details
// 	// err = stmt.QueryRow(note.OwnerName, note.Title, stringNull).Scan(&note.ID)
// 	if err != nil {
// 		logger.Log().Infof("There was an error executing the sql statement. Error: %v", err)
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	tx.Commit()
// 	logger.Log().Info("Note created successfully :)")

// 	return nil, nil
// }
