package entities

type (
	InsertFoodPlaceDto struct {
		BusinessName string        `gorm:"not null" json:"business_name"`
		Location     string        `gorm:"not null;unique" json:"location"`
		UserId       uint32        `gorm:"not null;unique" json:"user_id"`
		User         InsertUserDto `gorm:"foreignKey:UserId"`
	}

	GetFoodPlaceDto struct {
		UserId       uint32
		BusinessName string
		Location     string
	}

	FoodPlace struct {
		UserId       uint32   `json:"user_id"`
		BusinessName string   `json:"business_name"`
		Location     string   `json:"location"`
		Tags         []string `json:"tags"`
	}
)

func (f InsertFoodPlaceDto) TableName() string {
	return "food_places"
}
