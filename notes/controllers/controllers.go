package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/RamiroCuenca/go-rest-notesApi/common/handler"
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/database/connection"
	"github.com/RamiroCuenca/go-rest-notesApi/notes/models"
	"github.com/lib/pq"
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

// This handler fetch all notes
func NotesGetAll(w http.ResponseWriter, r *http.Request) {
	// 1° Create the sql statement and prepare null fields
	q := `SELECT * FROM notes`

	// 2° Open DB connection and start a transaction
	db := connection.NewPostgresClient()

	tx, err := db.Begin()
	if err != nil {
		logger.Log().Infof("Error starting transaction: %v", err)
		handler.SendError(w, 500) // Internal Server Error
		return
	}

	// 3° Prepare the transaction
	stmt, err := tx.Prepare(q)
	if err != nil {
		logger.Log().Infof("Error preparing transaction: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}
	defer stmt.Close()

	// 4° Execute the query and assign the returned id to the Note object
	// We will use QueryRow because the exec method returns two methods that are
	// not compatible with psql!
	rows, err := stmt.Query() // Query() return Rows object
	if err != nil {
		logger.Log().Infof("Error executing query: %v", err)
		tx.Rollback()
		handler.SendError(w, 500) // Internal Server Error
		return
	}
	defer rows.Close() // Close Rows object

	// 5° Create an array that will hold the results
	// and go over the rows object and assign each value to
	// a Note object and then append it to the notesArr
	// Also look that we handle the possible null values!
	notesArr := []models.Note{}

	for rows.Next() {

		n := models.Note{}
		nullDetail := sql.NullString{}
		nullUpdateAt := pq.NullTime{}

		err := rows.Scan(
			&n.ID,
			&n.OwnerName,
			&n.Title,
			&nullDetail,
			&n.CreatedAt,
			&nullUpdateAt,
		)

		if err != nil {
			logger.Log().Infof("Error scaning the received rows: %v", err)
			tx.Rollback()
			handler.SendError(w, 500) // Internal Server Error
			return
		}

		n.Details = nullDetail.String
		n.UpdatedAt = nullUpdateAt.Time

		notesArr = append(notesArr, n)

	}

	// 6° As there are no errors, commit the transaction
	tx.Commit()
	logger.Log().Infof("Notes fetched successfully! :)")

	// 7° Encode the Note into a JSON object
	// There a many ways we can do this, but i choose to use the Marshal
	// method as it organices for us the Json an it looks preety.
	json, _ := json.Marshal(notesArr)

	// Otherway, we can do this:
	// _ = json.NewEncoder(w).Encode(notesArr)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// 8° Send the response
	handler.SendResponse(w, http.StatusOK, json)
}

// This function manages the null string values
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}

	if null.String != "" {
		null.Valid = true
	}

	return null
}
