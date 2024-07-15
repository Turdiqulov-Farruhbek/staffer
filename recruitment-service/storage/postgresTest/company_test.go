package postgrestest

import (
	"context"
	"testing"

	gen "recruitment/genprotos" // Adjust import path as per your project structure
	"recruitment/storage/postgres"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *pgx.Conn {
	// Connect to a test database (replace with your test database connection parameters)
	connString := "postgres://postgres:root@localhost:5432/armish_db?sslmode=disable"
	db, err := pgx.Connect(context.Background(), connString)
	require.NoError(t, err)

	// Create necessary tables for testing
	_, err = db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS companies (
			company_id UUID PRIMARY KEY,
			name TEXT,
			address TEXT,
			industry TEXT,
			website TEXT,
			deleted_at BIGINT DEFAULT 0
		)
	`)
	require.NoError(t, err)

	return db
}

func teardownTestDB(t *testing.T, db *pgx.Conn) {
	// Drop tables after testing
	_, err := db.Exec(context.Background(), `DROP TABLE IF EXISTS companies`)
	require.NoError(t, err)
}

func TestCompanyCRUDOperations(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	storage := postgres.NewCompany(db)

	// Test case 1: Create a new company
	createReq := &gen.CreateCompany{
		Name:     "Test Company",
		Address:  "123 Test St",
		Industry: "Testing",
		Website:  "http://testcompany.com",
	}

	createResp, err := storage.Create(context.Background(), createReq)
	require.NoError(t, err)
	assert.NotNil(t, createResp)


	// Test case 3: Update company details
	updateReq := &gen.CompanyUpdate{
		Id:       uuid.NewString(),
		Name:     "Updated Test Company",
		Address:  "456 Updated St",
		Industry: "Updated Testing",
		Website:  "http://updatedtestcompany.com",
	}

	updateResp, err := storage.Update(context.Background(), updateReq)
	require.NoError(t, err)
	assert.NotNil(t, updateResp)

	// Verify updated details
	getByIDRespAfterUpdate, err := storage.GetByID(context.Background(), &gen.Byid{})
	require.NoError(t, err)
	assert.NotNil(t, getByIDRespAfterUpdate)

	// Verify company is marked as deleted
	getByIDRespAfterDelete, err := storage.GetByID(context.Background(), &gen.Byid{})
	require.Error(t, err) // Expecting an error since company should be deleted
	assert.Nil(t, getByIDRespAfterDelete)
}
