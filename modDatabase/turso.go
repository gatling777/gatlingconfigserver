package modDatabase

import (
	"configServer/modUtility"
	"database/sql"
)

type CDBTurso struct {
	sqlOpt *CSqlDbOpt
}

var g_singleTruso *CDBTurso = &CDBTurso{sqlOpt: newSqlDbOpt()}

func getSingleTurso() *CDBTurso {
	return g_singleTruso
}

func (pInst *CDBTurso) connect() error {
	dbUrl := modUtility.G_DBUrl
	token := modUtility.G_DBToken
	return pInst.sqlOpt.connect(dbUrl, token)
}

func (pInst *CDBTurso) execsql(sql string) error {
	return pInst.sqlOpt.execsql(sql)
}

func (pInst *CDBTurso) query(sql string) (*sql.Rows, error) {
	return pInst.sqlOpt.query(sql)
}
