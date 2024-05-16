package sql

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-faker/faker/v4"
	"github.com/kmcclive/goapipattern"
	"github.com/kmcclive/goapipattern/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ManufacturerServiceSuite struct {
	suite.Suite
	assert  *assert.Assertions
	sqlmock sqlmock.Sqlmock
	service goapipattern.ManufacturerService
}

func TestManufacturerServiceSuite(t *testing.T) {
	suite.Run(t, new(ManufacturerServiceSuite))
}

func (s *ManufacturerServiceSuite) SetupTest() {
	t := s.T()

	db, sql, err := mock.DB()
	require.NoError(t, err)

	s.assert = assert.New(t)
	s.sqlmock = sql
	s.service = NewManufacturerService(db)
}

func (s *ManufacturerServiceSuite) TestFetch_QueriesForID() {
	id := mock.ID()
	s.sqlmock.ExpectQuery("^SELECT (.+) FROM `manufacturers` WHERE `manufacturers`.`id` = ?").
		WithArgs(id, 1).
		WillReturnRows(s.newRows())

	s.service.FetchByID(id)

	s.assert.NoError(s.sqlmock.ExpectationsWereMet())
}

func (s *ManufacturerServiceSuite) TestFetch_WithRow_ReturnsManufacturer() {
	expected := new(goapipattern.Manufacturer)
	faker.FakeData(expected)
	expected.DeletedAt = gorm.DeletedAt{}
	rows := s.newRows()
	s.addRow(rows, expected.ID, expected.CreatedAt, expected.UpdatedAt, nil, expected.Name)
	s.sqlmock.ExpectQuery("").WillReturnRows(rows)

	actual, err := s.service.FetchByID(expected.ID)

	s.assert.NoError(err)
	s.assert.EqualValues(expected, actual)
}

func (s *ManufacturerServiceSuite) TestFetch_WithoutRow_ReturnsErrNotFound() {
	s.sqlmock.ExpectQuery("").WillReturnRows(s.newRows())

	actual, err := s.service.FetchByID(mock.ID())

	s.assert.Nil(actual)
	s.assert.ErrorIs(err, goapipattern.ErrNotFound)
}

func (s *ManufacturerServiceSuite) TestFetch_WithError_ReturnsError() {
	expectedErr := mock.Error()
	s.sqlmock.ExpectQuery("").WillReturnError(expectedErr)

	actual, err := s.service.FetchByID(mock.ID())

	s.assert.Nil(actual)
	s.assert.ErrorIs(err, expectedErr)
}

func (s *ManufacturerServiceSuite) newRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"})
}

func (s *ManufacturerServiceSuite) addRow(
	rows *sqlmock.Rows,
	id uint,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
	name string,
) {
	rows.AddRow(id, createdAt, updatedAt, deletedAt, name)
}
