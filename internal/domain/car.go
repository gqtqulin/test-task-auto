package domain

type Currency string

const (
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Car struct {
	CarId      int64    `json:"car_id" omitempty:"true"`
	Mark       string   `json:"mark" required:"true"`
	Model      string   `json:"model" required:"true"`
	OwnerCount int64    `json:"owner_count"`
	Price      int64    `json:"price"`
	Currency   Currency `json:"currency"`
	Options    []string `json:"options" required:"true"`
}
