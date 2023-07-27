package db

import (
	"fmt"
	"gin_app/app/config"
	"log"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var sqliteDBs map[string]*gorm.DB = make(map[string]*gorm.DB)
var mysqlDBs map[string]*gorm.DB = make(map[string]*gorm.DB)

func setDBPool(dbName string, db *gorm.DB, connPool *config.ConnPool) {
	if db != nil && connPool != nil && connPool.Enable {
		sqlDB, err := db.DB()
		if err != nil {
			panic(fmt.Sprintf("meet error when set DB pool for DB %v, error: %v", dbName, err))
		}
		sqlDB.SetMaxIdleConns(connPool.MaxIdleConns)
		sqlDB.SetMaxOpenConns(connPool.MaxOpenConns)
		duration, err := time.ParseDuration(connPool.ConnMaxLifetime)
		if err != nil {
			panic(fmt.Sprintf("meet error when set ConnMaxLifetime for DB %v, error: %v", dbName, err))
		}
		sqlDB.SetConnMaxLifetime(duration)
	}
}

func openSqliteDBs(dbConfs *map[string]config.SqliteDB) {
	for dbName, dbconf := range *dbConfs {
		if !dbconf.Enable {
			continue
		}
		db, err := gorm.Open(sqlite.Open(dbconf.File), &gorm.Config{})
		if err != nil {
			log.Printf("meet error when open Sqlite DB %v, error: %v", dbName, err)
			continue
		}
		setDBPool(dbName, db, &dbconf.ConnPool)
		sqliteDBs[dbName] = db
		log.Printf("Sqlite DB %v opened", dbName)
	}
}

func openMysqlDBs(dbConfs *map[string]config.MysqlDB) {
	for dbName, dbconf := range *dbConfs {
		if !dbconf.Enable {
			continue
		}
		dsn := strings.Builder{}
		dsn.WriteString(dbconf.User)
		dsn.WriteString(":")
		dsn.WriteString(dbconf.Password)
		dsn.WriteString("@tcp(")
		dsn.WriteString(strings.TrimSpace(dbconf.Host))
		dsn.WriteString(":")
		dsn.WriteString(strings.TrimSpace(dbconf.Port))
		dsn.WriteString(")/")
		dsn.WriteString(strings.TrimSpace(dbconf.Database))
		dsn.WriteString("?")
		dsn.WriteString("&parseTime=True&loc=Local")
		charset := strings.TrimSpace(dbconf.Charset)
		if charset != "" {
			dsn.WriteString("&charset=")
			dsn.WriteString(charset)
		}

		db, err := gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})
		if err != nil {
			log.Printf("meet error when open Mysql DB %v, error: %v", dbName, err)
			continue
		}
		setDBPool(dbName, db, &dbconf.ConnPool)
		sqliteDBs[dbName] = db
		log.Printf("Mysql DB %v opened", dbName)
	}
}

func closeDBs(dbs *map[string]*gorm.DB) {
	closed := make([]string, 0)
	for dbName, db := range *dbs {
		closed = append(closed, dbName)
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("meet error when get sqlDB for %v, error: %v\n", dbName, err)
			continue
		}
		err = sqlDB.Close()
		if err != nil {
			log.Printf("close sqlDB %v, error: %v\n", dbName, err)
		}
		log.Printf("DB %v closed", dbName)
	}

	for _, dbName := range closed {
		delete(*dbs, dbName)
	}
}

func OpenDBConnection() {
	openSqliteDBs(&config.DB_CONFIG.DB.Sqlite)
	openMysqlDBs(&config.DB_CONFIG.DB.Mysql)
}

func CloseDBConnection() {
	closeDBs(&sqliteDBs)
	closeDBs(&mysqlDBs)
}

func getSqliteDB(dbName ...string) *gorm.DB {
	if len(dbName) > 0 {
		db, ok := sqliteDBs[dbName[0]]
		if ok {
			log.Println("get Sqlite DB", dbName)
			return db
		} else {
			log.Println("don't found Sqlite connection for:", dbName)
			return nil
		}
	} else {
		var db *gorm.DB = nil
		for dbName, v := range sqliteDBs {
			db = v
			log.Println("get Sqlite DB", dbName)
			break
		}
		return db
	}
}

func GetDB() *gorm.DB {
	return getSqliteDB()
}
