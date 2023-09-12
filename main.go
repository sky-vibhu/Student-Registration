package main

import (
	"context"
	"fmt"
	"log"

	"api/controllers"
	"api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server            *gin.Engine
	studentservice    services.StudentService
	subjectservice    services.SubjectService
	studentcontroller controllers.StudentController
	subjectcontroller controllers.SubjectController
	ctx               context.Context
	studentcollection *mongo.Collection
	subjectcollection *mongo.Collection
	mongoclient       *mongo.Client
	err               error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongodb", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	studentcollection = mongoclient.Database("studentdb").Collection("students")
	studentservice = services.NewStudentService(studentcollection, ctx)
	studentcontroller = controllers.NewStudent(studentservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	studentcontroller.RegisterStudentRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
