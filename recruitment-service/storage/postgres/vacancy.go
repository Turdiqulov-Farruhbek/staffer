package postgres

import (
	"context"
	"strconv"
	"time"

	gen "recruitment/genprotos"

	"github.com/jackc/pgx/v5"
)

type Vacancy struct {
	db *pgx.Conn
}

func NewVacancy(db *pgx.Conn) *Vacancy {
	return &Vacancy{db: db}
}

// Create inserts a new vacancy into the database.
func (s *Vacancy) Create(ctx context.Context, req *gen.VacancyCreate) (*gen.Void, error) {
	query := `INSERT INTO vacancies (title, description, department_id, position_id, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(ctx, query, req.Title, req.Description, req.DepartmentId, req.PositionId, req.Status)
	if err != nil {
		return nil, err
	}
	return &gen.Void{}, nil
}

// GetByID retrieves a vacancy by ID.
func (s *Vacancy) GetByID(ctx context.Context, req *gen.Byid) (*gen.VacancyGetResUpdateReq, error) {
	query := `SELECT id, title, description, department_id, position_id, status FROM vacancies WHERE id = $1 AND deleted_at = 0`
	row := s.db.QueryRow(ctx, query, req.Id)

	var vacancy gen.VacancyGetResUpdateReq
	err := row.Scan(&vacancy.Id, &vacancy.Title, &vacancy.Description, &vacancy.DepartmentId, &vacancy.PositionId, &vacancy.Status)
	if err != nil {
		return nil, err
	}

	return &vacancy, nil
}

// Update modifies an existing vacancy in the database.
func (s *Vacancy) Update(ctx context.Context, req *gen.VacancyGetResUpdateReq) (*gen.Void, error) {
	// First, retrieve the existing vacancy details
	var existing gen.VacancyGetResUpdateReq
	query := `SELECT title, description, department_id, position_id, status FROM vacancies WHERE id = $1 AND deleted_at = 0`
	err := s.db.QueryRow(ctx, query, req.Id).Scan(&existing.Title, &existing.Description, &existing.DepartmentId, &existing.PositionId, &existing.Status)
	if err != nil {
		return nil, err
	}

	// Use existing values if the new values are empty
	if req.Title == "" {
		req.Title = existing.Title
	}
	if req.Description == "" {
		req.Description = existing.Description
	}
	if req.DepartmentId == "" {
		req.DepartmentId = existing.DepartmentId
	}
	if req.PositionId == "" {
		req.PositionId = existing.PositionId
	}
	if req.Status == "" {
		req.Status = existing.Status
	}

	// Prepare the SQL statement for updating a vacancy
	updateQuery := `UPDATE vacancies SET title = $1, description = $2, department_id = $3, position_id = $4, status = $5 WHERE id = $6 AND deleted_at = 0`
	_, err = s.db.Exec(ctx, updateQuery, req.Title, req.Description, req.DepartmentId, req.PositionId, req.Status, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.Void{}, nil
}

// Delete marks a vacancy as deleted by setting the deleted_at timestamp.
func (s *Vacancy) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	query := `UPDATE vacancies SET deleted_at = $1 WHERE id = $2 AND deleted_at = 0`
	deletedAt := time.Now().Unix()
	_, err := s.db.Exec(ctx, query, deletedAt, req.Id)
	if err != nil {
		return nil, err
	}
	return &gen.Void{}, nil
}

// GetAll retrieves all vacancies from the database with optional filtering.
func (s *Vacancy) GetAll(ctx context.Context, req *gen.VacancyGetAllReq) (*gen.VacancyGetAllRes, error) {
	query := `SELECT id, title, description, department_id, position_id, status FROM vacancies WHERE deleted_at = 0`
	args := []interface{}{}
	argIndex := 1

	// Add filters to the query if they are provided in the request
	if req.Title != "" {
		query += ` AND title = $` + strconv.Itoa(argIndex)
		args = append(args, req.Title)
		argIndex++
	}
	if req.Description != "" {
		query += ` AND description = $` + strconv.Itoa(argIndex)
		args = append(args, req.Description)
		argIndex++
	}
	if req.DepartmentId != "" {
		query += ` AND department_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.DepartmentId)
		argIndex++
	}
	if req.PositionId != "" {
		query += ` AND position_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.PositionId)
		argIndex++
	}
	if req.Status != "" {
		query += ` AND status = $` + strconv.Itoa(argIndex)
		args = append(args, req.Status)
		argIndex++
	}

	if req.Pagination != nil {
		if req.Pagination.Limit > 0 {
			query += ` LIMIT $` + strconv.Itoa(argIndex)
			args = append(args, req.Pagination.Limit)
			argIndex++
		}
		if req.Pagination.Offset > 0 {
			query += ` OFFSET $` + strconv.Itoa(argIndex)
			args = append(args, req.Pagination.Offset)
			argIndex++
		}
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []*gen.VacancyGetResUpdateReq
	for rows.Next() {
		var vacancy gen.VacancyGetResUpdateReq
		err := rows.Scan(&vacancy.Id, &vacancy.Title, &vacancy.Description, &vacancy.DepartmentId, &vacancy.PositionId, &vacancy.Status)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, &vacancy)
	}

	return &gen.VacancyGetAllRes{
		Vacancies: vacancies,
		Success:   true,
	}, nil
}
