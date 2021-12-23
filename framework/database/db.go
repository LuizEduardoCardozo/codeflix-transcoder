package database

import (
	"encoder/domain"
	"log"

	"gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

type Database struct {
	Env             string
	DB              *gorm.DB
	DBOptions       gorm.Option
	DBOptionsTest   gorm.Option
	DBDSN           string
	DBDSNTest       string
	Debug           bool
	AutoMigrationDB bool
}

func (db *Database) Connect() (*gorm.DB, error) {
	var ConnectDsn gorm.Dialector
	var ConnectOptions gorm.Option

	switch db.Env {
	case "test":
		{
			ConnectDsn = sqlite.Open(db.DBDSNTest)
			ConnectOptions = db.DBOptionsTest
		}
	default:
		{
			ConnectDsn = sqlite.Open(db.DBDSN)
			ConnectOptions = db.DBOptions
		}
	}

	var err error
	db.DB, err = gorm.Open(ConnectDsn, ConnectOptions)
	if err != nil {
		log.Fatalf("error while openning database, in mode %s: %s\n", db.Env, err.Error())
		return nil, err
	}

	if db.AutoMigrationDB {
		if err := db.DB.AutoMigrate(&domain.Video{}, &domain.Job{}); err != nil {
			log.Fatalf("error while migrating database: %s\n", err.Error())
			return nil, err
		}
	}

	return db.DB, nil
}

func NewDB() *Database {
	return &Database{}
}

func NewTestDB() *gorm.DB {
	dbInstance := NewDB()

	dbInstance.Env = "test"
	dbInstance.DBDSNTest = "file::memory:?cache=shared"
	dbInstance.DBOptionsTest = &gorm.Config{}
	dbInstance.AutoMigrationDB = true
	dbInstance.Debug = true

	conn, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("test db error while connecting: %s\n", err.Error())
	}

	return conn
}
