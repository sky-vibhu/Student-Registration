package services

import (
	"context"
	"errors"

	"api/models"
	// "github.com/charmbracelet/charm/ui/username"
	// "github.com/gobuffalo/flect/name"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type StudentServiceImpl struct {
	studentcollection  *mongo.Collection
	ctx                context.Context
}

type SubjectServiceImpl struct {
	subjectcollection  *mongo.Collection
	ctx                context.Context
}

func NewStudentService(studentcollection *mongo.Collection, ctx context.Context ) StudentService {
	return &StudentServiceImpl {
		studentcollection: studentcollection,
		ctx:   ctx,
	}
}

func NewSubjectService(subjectcollection *mongo.Collection, ctx context.Context) SubjectService {
	return &SubjectServiceImpl {
		subjectcollection: subjectcollection,
		ctx: ctx,
	}
}


func (u *SubjectServiceImpl) CreateSubject(subject *models.Subject) error {
	_, err := u.subjectcollection.InsertOne(u.ctx, subject)
	return err
}


func (u *StudentServiceImpl) CreateStudent(student *models.Student) error {
	_, err := u.studentcollection.InsertOne(u.ctx, student)
	return err
}


func (u *StudentServiceImpl) GetStudent(name *string) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.studentcollection.FindOne(u.ctx, query).Decode(&student)
	return student, err
}

func (u *StudentServiceImpl) GetStudentbyId(Id *string) (*models.Student, error) {
	var student *models.Student
	query := bson.D{bson.E{Key: "id", Value: Id}}
	err := u.studentcollection.FindOne(u.ctx, query).Decode(&student)
	return student, err
}


func (u *SubjectServiceImpl) GetAllSubjects() ([]*models.Subject, error) {
	var subjects []*models.Subject
	cursor, err := u.subjectcollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var subject models.Subject
		err := cursor.Decode(&subject)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, &subject)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	
	if len(subjects) == 0 {
		return nil, errors.New("Documents not found")
	}
	return subjects, nil
}

