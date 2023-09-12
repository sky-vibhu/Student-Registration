package services

import "api/models"

type StudentService interface {
	CreateStudent(*models.Student) error
	GetStudent(*string) (*models.Student, error)
	GetStudentbyId (*string)(*models.Student, error)
}

type SubjectService interface {
	CreateSubject(*models.Subject) error
	GetAllSubjects()([]*models.Subject, error)
}