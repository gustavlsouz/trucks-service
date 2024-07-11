package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/lib/pq"
)

var singleton *persistence

type Persistence interface {
	Start(string) error
	Database() *sql.DB
	ScanStruct(rows *sql.Rows, destiny interface{}) error
}

type persistence struct {
	database *sql.DB
}

func GetPersistenceInstance() Persistence {
	if singleton != nil {
		return singleton
	}
	singleton = &persistence{}
	return singleton
}

func (persistence *persistence) connect() error {
	if persistence.database != nil {
		return nil
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	sslmode := os.Getenv("DB_SSLMODE")
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	database, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Panicf("error to open connection with database %v", err)
		return err
	}

	if err := database.Ping(); err != nil {
		log.Panic(err)
	}

	persistence.database = database
	return err
}

func (persistence *persistence) Start(migrationPath string) error {
	persistence.connect()

	migrator := newMigrator(persistence)

	migrator.Migrate(migrationPath)

	return nil
}

func (persistence *persistence) Database() *sql.DB {
	return persistence.database
}

func (persistence *persistence) ScanStruct(rows *sql.Rows, destiny interface{}) error {
	value := reflect.ValueOf(destiny)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("destiny must be a pointer to a struct")
	}

	value = value.Elem()
	valueType := value.Type()

	// Obtem as colunas do resultado da consulta
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("error to list columns: %w", err)
	}

	if len(columns) > valueType.NumField() {
		return errors.New("number of columns exceeds number of struct fields")
	}

	scanArgs := make([]interface{}, len(columns))

	for i := 0; i < len(columns); i++ {
		scanArgs[i] = value.Field(i).Addr().Interface()
	}

	if err := rows.Scan(scanArgs...); err != nil {
		return err
	}

	return nil
}
