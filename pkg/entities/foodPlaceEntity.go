package entities

type (
	InsertFoodPlaceDto struct {
		UserId       uint32        `gorm:"not null;unique" json:"user_id"`
		BusinessName string        `gorm:"not null" json:"business_name"`
		Location     string        `gorm:"not null;unique" json:"location"`
		Tags         string        `gorm:"not null" json:"tags"`
		User         InsertUserDto `gorm:"foreignKey:UserId"`
	}

	GetFoodPlaceDto struct {
		UserId uint32
	}

	FoodPlace struct {
		UserId       uint32   `json:"user_id"`
		BusinessName string   `json:"business_name"`
		Location     string   `json:"location"`
		Tags         []string `json:"tags"`
	}
)

func (InsertFoodPlaceDto) TableName() string {
	return "food_places"
}
func (GetFoodPlaceDto) TableName() string {
	return "food_places"
}
