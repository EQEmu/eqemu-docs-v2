package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v4"
	"os"
	"time"
)

type Instance struct {
	db *sql.DB
}

func (i *Instance) Db() *sql.DB {
	return i.db
}

func NewInstance() *Instance {
	return &Instance{db: newConnection()}
}

func newConnection() *sql.DB {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

type DbSchemaRowResult struct {
	Table           string      `yaml:",omitempty"`
	Column          string      `yaml:"column"`
	DataType        string      `yaml:"data_type"`
	ColumnKey       null.String `yaml:"column_key"`
	OrdinalPosition string      `yaml:"ordinal_position"`
	IsNullable      string      `yaml:"is_nullable"`
	ColumnType      string      `yaml:"column_type"`
	ColumnDefault   null.String `yaml:"column_default"`
}

func (i *Instance) GetSchema() (map[string][]DbSchemaRowResult, error) {
	rows, err := i.db.Query(
		fmt.Sprintf(
			`
		SELECT
		  TABLE_NAME,
		  COLUMN_NAME,
		  DATA_TYPE,
		  COLUMN_KEY,
		  ORDINAL_POSITION,
		  IS_NULLABLE,
		  COLUMN_TYPE,
		  COLUMN_DEFAULT
		FROM
		  INFORMATION_SCHEMA.COLUMNS
		WHERE
		  TABLE_SCHEMA = '%v';
	  `, os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return nil, err
	}

	tableColumnResponses := make(map[string][]DbSchemaRowResult, 0)

	defer rows.Close()
	for rows.Next() {
		var row DbSchemaRowResult
		err = rows.Scan(
			&row.Table,
			&row.Column,
			&row.DataType,
			&row.ColumnKey,
			&row.OrdinalPosition,
			&row.IsNullable,
			&row.ColumnType,
			&row.ColumnDefault,
		)
		if err != nil {
			return nil, err
		}

		tableColumnResponses[row.Table] = append(tableColumnResponses[row.Table], row)
	}

	return tableColumnResponses, nil
}

func (i *Instance) GetSchemaTableNames() ([]string, error) {
	rows, err := i.db.Query(
		fmt.Sprintf(
			`
		SELECT
		  TABLE_NAME
		FROM
		  INFORMATION_SCHEMA.COLUMNS
		WHERE
		  TABLE_SCHEMA = '%v' 
		  AND TABLE_NAME NOT LIKE 'spire_%%'
		  AND TABLE_NAME NOT LIKE '_backup_%%'
		GROUP BY TABLE_NAME;
	  `, os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return nil, err
	}

	tables := []string{}

	defer rows.Close()
	for rows.Next() {
		var row DbSchemaRowResult
		err = rows.Scan(
			&row.Table,
		)
		if err != nil {
			return nil, err
		}

		tables = append(tables, row.Table)
	}

	return tables, nil
}
