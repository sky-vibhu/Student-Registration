package controllers

import (
	"net/http"

	"api/models"
	"api/services" // "example.com/student-apis/services"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService services.StudentService
}

type SubjectController struct {
	SubjectService services.SubjectService
}

func NewStudent(studentservice services.StudentService) StudentController {
	return StudentController{
		StudentService: studentservice,
	}
}

func NewSubject(subjectservice services.SubjectService) SubjectController {
	return SubjectController{
		SubjectService: subjectservice,
	}
}

func (uc *SubjectController) CreateSubject(ctx *gin.Context) {
	var subject models.Subject
	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.SubjectService.CreateSubject(&subject)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *StudentController) CreateStudent(ctx *gin.Context) {
	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.StudentService.CreateStudent(&student)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *StudentController) GetStudent(ctx *gin.Context) {

	ctx.JSON(200, "")
}

func (uc *SubjectController) GetAllSubjects(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *StudentController) GetStudentbyId(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *StudentController) RegisterStudentRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/student")
	userroute.POST("/create", uc.CreateStudent)
	userroute.GET("/get/:name", uc.GetStudent)
	userroute.GET("/get/:id", uc.GetStudentbyId)
}

func (uc *SubjectController) RegisterSubjectRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/subject")
	userroute.POST("/create", uc.CreateSubject)
	userroute.GET("/getall", uc.GetAllSubjects)
}
