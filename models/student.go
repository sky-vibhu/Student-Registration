package models 


type Student struct {
	Name           string   			`json: "name"     bson: "student_name"`
	PhoneNumber    string       		`json: "number"   bson: "student_phone_number"`
	Gender         string				`json: "gender"   bson: "student_gender"`
	Subjects      []Subject	    		`json: "subjects" bson: "student_subjects"`
}
