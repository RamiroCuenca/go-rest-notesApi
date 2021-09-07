package connection

import (
	"database/sql"
	"fmt"

	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	_ "github.com/lib/pq" // Dont forget to import it, it provides the drivers for postgres
)

// Postgre db
type PostgreClient struct {
	*sql.DB
}

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "notes-app-db"
)

// Starts the connection to our postgres database
func NewPostgresClient() *PostgreClient {

	// Log db credentials
	fmt.Println("Psql is using the following config:")
	fmt.Printf("host: %15s\n", host)
	fmt.Printf("port: %15s\n", port)
	fmt.Printf("user: %15s\n", user)
	fmt.Printf("password: %11s\n", password)
	fmt.Printf("database: %11s\n", dbname)

	// dsn = data source name
	// The idea is that the parameters are variable according to the environment
	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"

	// Open method does not create the connection, it simply check if the arguments work properly
	// That's why we MUST check with Ping() method if it's working!
	db, err := sql.Open("postgres", dsn)
	// Close DB after program exits
	defer db.Close()

	if err != nil {
		// If we can not connect to the database, log the error and close the app with panic
		logger.Log().Errorf("Error opening the database. Reason: %v", err)
		panic(err)
	}

	// Check if the connection with the database is stable and alive
	err = db.Ping()

	if err != nil {
		logger.Log().Errorf("Error with the connection with the database. Reason: %v", err)
	}

	// fmt.Printf("\nConnected to %s succesfully!\n", dbname)
	logger.Log().Infof("Connected to %s succesfully!", dbname)

	// As there are no errors, return the database connection
	return &PostgreClient{db}
}
