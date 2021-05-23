package structs

type Login struct {
	Username string `json:"name" bson:"name"`
}

type User struct {
	ObjectId string `json:"objectId"`
	Age      int    `json:"age"`
	Name     string `json:"name"`
}
