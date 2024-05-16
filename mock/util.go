package mock

import (
	"errors"
	"math/rand"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DB() (*gorm.DB, sqlmock.Sqlmock, error) {
	conn, sql, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	return db, sql, nil
}

func Error() error {
	return errors.New(faker.Sentence())
}

func ID() uint {
	return uint(rand.Uint32())
}
