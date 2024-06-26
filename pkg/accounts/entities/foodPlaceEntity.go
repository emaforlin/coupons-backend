package entities

type (
	InsertFoodLocalDto struct {
		FoodPlaceId  uint32 `gorm:"primaryKey;autoIncrement" json:"food_place_id"`
		UserId       uint32 `gorm:"not null;unique" json:"user_id"`
		BusinessName string `gorm:"not null" json:"business_name"`
		Location     string `gorm:"not null;unique" json:"location"`
	}

	GetFoodPlaceDto struct {
		UserId       uint32
		BusinessName string
		Location     string
	}

	FoodPlace struct {
		FoodPlaceId  uint32 `json:"food_place_id"`
		UserId       uint32 `json:"user_id"`
		BusinessName string `json:"business_name"`
		Location     string `json:"location"`
	}
)
