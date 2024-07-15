package postgres

import (
	"context"
	"strconv"

	gen "recruitment/genprotos"

	"github.com/jackc/pgx/v5"
)

type JobApp struct {
	db *pgx.Conn
}

func NewJobApp(db *pgx.Conn) *JobApp {
	return &JobApp{db: db}
}

// CreateJobApplication inserts a new job application into the database.
func (s *JobApp) CreateJobApplication(ctx context.Context, req *gen.CreateJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	query := `INSERT INTO jobapplications (applicant_name, position, status) VALUES ($1, $2, $3) RETURNING id`
	row := s.db.QueryRow(ctx, query, req.ApplicantName, req.Position)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &gen.JobApplicationResponse{JobApplication: &gen.JobApplication{
		Id:            id,
		ApplicantName: req.ApplicantName,
		Position:      req.Position,
		Steps:         req.Steps,
	}}, nil
}

// CreateJobStep inserts a new job step into the database.
func (s *JobApp) CreateJobStep(ctx context.Context, req *gen.CreateJobStepRequest) (*gen.JobStepResponse, error) {
	query := `INSERT INTO jobsteps (application_id, description, status) VALUES ($1, $2, $3) RETURNING id`
	row := s.db.QueryRow(ctx, query, req.ApplicationId, req.Description, req.Status)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &gen.JobStepResponse{JobStep: &gen.JobStep{
		Id:            id,
		ApplicationId: req.ApplicationId,
		Description:   req.Description,
		Status:        req.Status,
	}}, nil
}

// UpdateJobStep updates an existing job step in the database.
func (s *JobApp) UpdateJobStep(ctx context.Context, req *gen.UpdateJobStepRequest) (*gen.JobStepResponse, error) {
	query := `UPDATE jobsteps SET description = $1, status = $2 WHERE id = $3`
	_, err := s.db.Exec(ctx, query, req.Description, req.Status, req.StepId)
	if err != nil {
		return nil, err
	}

	return s.GetJobStepById(ctx, &gen.GetJobStepByIdRequest{ApplicationId: req.StepId})
}

// UpdateJobApplication updates an existing job application in the database.
func (s *JobApp) UpdateJobApplication(ctx context.Context, req *gen.UpdateJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	query := `UPDATE jobapplications SET status = $1 WHERE id = $2`
	_, err := s.db.Exec(ctx, query, req.Status, req.Id)
	if err != nil {
		return nil, err
	}

	return s.GetJobApplicationById(ctx, &gen.GetJobApplicationByIdRequest{Id: req.Id})
}

// GetJobApplicationById retrieves a job application by ID.
func (s *JobApp) GetJobApplicationById(ctx context.Context, req *gen.GetJobApplicationByIdRequest) (*gen.JobApplicationResponse, error) {
	query := `SELECT id, applicant_name, position, status FROM jobapplications WHERE id = $1`
	row := s.db.QueryRow(ctx, query, req.Id)

	var jobApplication gen.JobApplication
	err := row.Scan(&jobApplication.Id, &jobApplication.ApplicantName, &jobApplication.Position, &jobApplication.Status)
	if err != nil {
		return nil, err
	}

	return &gen.JobApplicationResponse{JobApplication: &jobApplication}, nil
}

// GetJobStepById retrieves a job step by ID.
func (s *JobApp) GetJobStepById(ctx context.Context, req *gen.GetJobStepByIdRequest) (*gen.JobStepResponse, error) {
	query := `SELECT id, application_id, description, status FROM jobsteps WHERE id = $1`
	row := s.db.QueryRow(ctx, query, req.ApplicationId)

	var jobStep gen.JobStep
	err := row.Scan(&jobStep.Id, &jobStep.ApplicationId, &jobStep.Description, &jobStep.Status)
	if err != nil {
		return nil, err
	}

	return &gen.JobStepResponse{JobStep: &jobStep}, nil
}

// GetAllJobApplications retrieves all job applications with optional filters.
func (s *JobApp) GetAllJobApplications(ctx context.Context, req *gen.GetJobApplicationsRequest) (*gen.JobApplicationsResponse, error) {
	query := `SELECT id, applicant_name, position, status FROM jobapplications WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

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

	var jobApplications []*gen.JobApplication
	for rows.Next() {
		var jobApplication gen.JobApplication
		err := rows.Scan(&jobApplication.Id, &jobApplication.ApplicantName, &jobApplication.Position, &jobApplication.Status)
		if err != nil {
			return nil, err
		}
		jobApplications = append(jobApplications, &jobApplication)
	}

	return &gen.JobApplicationsResponse{
		JobApplications: jobApplications,
	}, nil
}

// DeleteJobApplication deletes a job application by ID.
func (s *JobApp) DeleteJobApplication(ctx context.Context, req *gen.DeleteJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	query := `DELETE FROM jobapplications WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.JobApplicationResponse{}, nil
}

// DeleteJobStep deletes a job step by ID.
func (s *JobApp) DeleteJobStep(ctx context.Context, req *gen.DeleteJobStepRequest) (*gen.JobStepResponse, error) {
	query := `DELETE FROM jobsteps WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.StepId)
	if err != nil {
		return nil, err
	}

	return &gen.JobStepResponse{}, nil
}
