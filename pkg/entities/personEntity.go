package entities

type (
	InsertPersonDto struct {
		FirstName string        `gorm:"not null" json:"first_name"`
		LastName  string        `gorm:"not null" json:"last_name"`
		UserId    uint32        `gorm:"not null;unique" json:"user_id"`
		User      InsertUserDto `gorm:"foreignKey:UserId"`
	}
	GetPersonDto struct {
		UserId    uint32
		FirstName string
		LastName  string
	}
	Person struct {
		UserId    uint32 `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)

func (InsertPersonDto) TableName() string {
	return "customers"
}

func (GetPersonDto) TableName() string {
	return "customers"
}
