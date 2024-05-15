package sql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/kmcclive/goapipattern"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ManufacturerServiceSuite struct {
	suite.Suite
	assert  *assert.Assertions
	db      *gorm.DB
	mock    sqlmock.Sqlmock
	service goapipattern.ManufacturerService
}

func (s *ManufacturerServiceSuite) SetupTest() {
	require := require.New(s.T())

	conn, mock, err := sqlmock.New()
	require.NoError(err)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(err)

	s.db = db
	s.mock = mock
	s.service = NewManufacturerService(db)
	s.assert = assert.New(s.T())
}

func (s *ManufacturerServiceSuite) TestFetch() {
	expected := new(goapipattern.Manufacturer)
	faker.FakeData(expected)
	expected.DeletedAt = gorm.DeletedAt{}
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
	rows.AddRow(expected.ID, expected.CreatedAt, expected.UpdatedAt, nil, expected.Name)
	s.mock.ExpectQuery("^SELECT (.+) FROM `manufacturers` WHERE `manufacturers`.`id` = ?").WillReturnRows(rows)

	actual, err := s.service.FetchByID(expected.ID)

	s.assert.NoError(err)
	s.assert.EqualValues(expected, actual)
}

func TestManufacturerServiceSuite(t *testing.T) {
	suite.Run(t, new(ManufacturerServiceSuite))
}
