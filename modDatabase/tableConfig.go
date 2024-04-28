package modDatabase

import (
	"database/sql"
	"fmt"
)

type cFieldDataConfig struct {
	Name string
	Text string
}

type CTableConfig struct {
}

var g_singleTableConfig *CTableConfig = &CTableConfig{}

func getSingleTableConfig() *CTableConfig {
	return g_singleTableConfig
}

const sql_create_testproject = `CREATE TABLE if not exists configServer (
	Name CHAR(50) PRIMARY KEY,
	Text1 TEXT,
 );`

func (pInst *CTableConfig) initialize() error {
	err := getSingleTurso().connect()
	if err != nil {
		return err
	}
	err = pInst.createTable()

	return err
}

func (pInst *CTableConfig) createTable() error {
	err := getSingleTurso().execsql(sql_create_testproject)
	if err != nil {
		return err
	}

	return nil
}

func (pInst *CTableConfig) Insert(name, text string) error {
	sql := fmt.Sprintf("INSERT INTO configServer(Name, Text1) VALUES('%s', '%s');", name, text)
	err := getSingleTurso().execsql(sql)

	return err
}
func (pInst *CTableConfig) Query() (*sql.Rows, error) {
	sql := "SELECT * FROM configServer"

	rows, err := getSingleTurso().query(sql)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
