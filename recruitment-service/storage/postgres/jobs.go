package postgres

import (
	"context"
	"strconv"
	"time"

	gen "recruitment/genprotos"

	"github.com/jackc/pgx/v5"
)

type Job struct {
	db *pgx.Conn
}

func NewJob(db *pgx.Conn) *Job {
	return &Job{db: db}
}

// Create inserts a new job application into the database.
func (s *Job) Create(ctx context.Context, req *gen.JobApplicationCreate) (*gen.Void, error) {
	query := `INSERT INTO job_applications (candidate_id, vacancy_id, resume_id, status) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(ctx, query, req.CandidateId, req.VacancyId, req.ResumeId, req.Status)
	if err != nil {
		return nil, err
	}
	return &gen.Void{}, nil
}

// Get retrieves a job application by ID.
func (s *Job) GetByID(ctx context.Context, req *gen.Byid) (*gen.JobApplitcationGetRes, error) {
	query := `SELECT id, candidate_id, vacancy_id, resume_id, status FROM job_applications WHERE id = $1 AND deleted_at = 0`
	row := s.db.QueryRow(ctx, query, req.Id)

	var jobApplication gen.JobApplitcationGetRes
	err := row.Scan(&jobApplication.Id, &jobApplication.CandidateId, &jobApplication.VacancyId, &jobApplication.ResumeId, &jobApplication.Status)
	if err != nil {
		return nil, err
	}

	return &jobApplication, nil
}

// Update modifies an existing job application in the database.
func (s *Job) Update(ctx context.Context, req *gen.JobApplicationUpdate) (*gen.Void, error) {
	// First, retrieve the existing job application details
	var existing gen.JobApplitcationGetRes
	query := `SELECT resume_id, status FROM job_applications WHERE id = $1 AND deleted_at = 0`
	err := s.db.QueryRow(ctx, query, req.Id).Scan(&existing.ResumeId, &existing.Status)
	if err != nil {
		return nil, err
	}

	// Use existing values if the new values are empty
	if req.ResumeId == "" {
		req.ResumeId = existing.ResumeId
	}
	if req.Status == "" {
		req.Status = existing.Status
	}

	// Prepare the SQL statement for updating a job application
	updateQuery := `UPDATE job_applications SET resume_id = $1, status = $2 WHERE id = $3 AND deleted_at = 0`
	_, err = s.db.Exec(ctx, updateQuery, req.ResumeId, req.Status, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.Void{}, nil
}

// Delete marks a job application as deleted by setting the deleted_at timestamp.
func (s *Job) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	query := `UPDATE job_applications SET deleted_at = $1 WHERE id = $2 AND deleted_at = 0`
	deletedAt := time.Now().Unix()
	_, err := s.db.Exec(ctx, query, deletedAt, req.Id)
	if err != nil {
		return nil, err
	}
	return &gen.Void{}, nil
}

// GetAll retrieves all job applications from the database with optional filtering.
func (s *Job) GetAll(ctx context.Context, req *gen.JobApplicationGetAll) (*gen.JobApplicationGetAllRes, error) {
	query := `SELECT id, candidate_id, vacancy_id, resume_id, status FROM job_applications WHERE deleted_at = 0`
	args := []interface{}{}
	argIndex := 1

	// Add filters to the query if they are provided in the request
	if req.CandidateId != "" {
		query += ` AND candidate_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.CandidateId)
		argIndex++
	}
	if req.VacancyId != "" {
		query += ` AND vacancy_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.VacancyId)
		argIndex++
	}
	if req.ResumeId != "" {
		query += ` AND resume_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.ResumeId)
		argIndex++
	}
	if req.Status != "" {
		query += ` AND status = $` + strconv.Itoa(argIndex)
		args = append(args, req.Status)
		argIndex++
	}

	if req.Filter != nil {
		if req.Filter.Limit > 0 {
			query += ` LIMIT $` + strconv.Itoa(argIndex)
			args = append(args, req.Filter.Limit)
			argIndex++
		}
		if req.Filter.Offset > 0 {
			query += ` OFFSET $` + strconv.Itoa(argIndex)
			args = append(args, req.Filter.Offset)
			argIndex++
		}
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobApplications []*gen.JobApplitcationGetRes
	for rows.Next() {
		var jobApplication gen.JobApplitcationGetRes
		err := rows.Scan(&jobApplication.Id, &jobApplication.CandidateId, &jobApplication.VacancyId, &jobApplication.ResumeId, &jobApplication.Status)
		if err != nil {
			return nil, err
		}
		jobApplications = append(jobApplications, &jobApplication)
	}

	return &gen.JobApplicationGetAllRes{
		JobApplications: jobApplications,
		Success:         true,
	}, nil
}
