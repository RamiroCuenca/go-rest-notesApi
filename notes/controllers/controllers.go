package controllers

import (
	"database/sql"

	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/database/connection"
	"github.com/RamiroCuenca/go-rest-notesApi/notes/models"
)

type Database struct {
	*connection.PostgreClient
}

func (db *Database) NotesCreate(note *models.Note) (err error) {
	// 1° Prepare the query that will be executed
	q := `INSERT INTO notes (owner_name, title, datails)
	VALUES ($1, $2, $3) RETURNING id`

	// Be prepared if details are empty
	stringNull := sql.NullString{}

	if note.Details == "" {
		stringNull.Valid = false
	} else {
		stringNull.String = note.Details
	}

	// A time ago... i used to open the database here, but at least on this
	// particular project we open it on the main file so it is not necessary
	// to be opened here

	// 2° Open a transaction
	// db.PostgreClient.DB.Begin() // Have it in case it fails the tx
	tx, err := db.Begin()
	if err != nil {
		logger.Log().Infof("There was an error beggining the tx. Error: %v", err)
		return
	}

	// 3° Prepare the query
	stmt, err := tx.Prepare(q)
	if err != nil {
		logger.Log().Infof("There was an error preparing the sql statement. Error: %v", err)
		tx.Rollback()
		return
	}
	defer stmt.Close()

	// 4° Execute the query. Send the parameters
	// We scan it and assign the returned id to the object
	// StringNull == note.details
	err = stmt.QueryRow(note.OwnerName, note.Title, stringNull).Scan(&note.ID)
	if err != nil {
		logger.Log().Infof("There was an error executing the sql statement. Error: %v", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	logger.Log().Info("Note created successfully :)")

	return nil
}

func (db *Database) NotesGetAll() ([]models.Note, error) {
	// 1° Prepare the query that will be executed
	q := `SELECT id, owner_name, title, details
		created_at, updated_at FROM notes`

	// Be prepared if details or updated at are empty
	// stringNull := sql.NullString{}

	// 2° Open a transaction
	tx, err := db.Begin()
	if err != nil {
		logger.Log().Infof("There was an error beggining the tx. Error: %v", err)
		return nil, err
	}

	// 3° Prepare the query
	stmt, err := tx.Prepare(q)
	if err != nil {
		logger.Log().Infof("There was an error preparing the sql statement. Error: %v", err)
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	// 4° Execute the query. Send the parameters
	// We scan it and assign the returned id to the object
	// StringNull == note.details
	// err = stmt.QueryRow(note.OwnerName, note.Title, stringNull).Scan(&note.ID)
	if err != nil {
		logger.Log().Infof("There was an error executing the sql statement. Error: %v", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	logger.Log().Info("Note created successfully :)")

	return nil, nil
}
