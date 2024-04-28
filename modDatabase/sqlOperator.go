package modDatabase

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/tursodatabase/libsql-client-go/libsql"
)

type CSqlDbOpt struct {
	dbConnector *driver.Connector
	database    *sql.DB
}

func newSqlDbOpt() *CSqlDbOpt {
	return &CSqlDbOpt{}
}

func (pInst *CSqlDbOpt) connect(dbUrl, token string) error {
	connector, err := libsql.NewConnector(dbUrl, libsql.WithAuthToken(token))
	if err != nil {
		return err
	}
	pInst.dbConnector = &connector

	db := sql.OpenDB(connector)
	if db == nil {
		return errors.New("db open error.")
	}
	pInst.database = db

	return nil
}

func (instSelf *CSqlDbOpt) execsql(sql string) error {
	_, err := instSelf.database.Exec(sql)
	return err
}

func (instSelf *CSqlDbOpt) query(sql string) (*sql.Rows, error) {
	rows, err := instSelf.database.Query(sql)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (pInst *CSqlDbOpt) GetDatabase() *sql.DB {
	return pInst.database
}
