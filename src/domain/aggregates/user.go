package aggregates

type User struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty" gorm:"primary_key"`
	Name     string `json:"name,omitempty" bson:"name" gorm:"not null"`
	Lastname string `json:"lastname,omitempty" bson:"lastname" gorm:"not null"`
	Age      int32  `son:"age,omitempty" bson:"age" gorm:"not null"`
	Email    string `json:"email,omitempty" bson:"email" gorm:"not null"`
}
