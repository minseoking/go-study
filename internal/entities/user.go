package entities

// swagger:model User
type User struct {
	// ID
	// Required: true
	UserId string `json:"userId"`
	// 이름
	// Required: true
	Name string `json:"Name"`
}
