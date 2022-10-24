package entities

// swagger:model Auth
type Auth struct {
	// ID
	// Required: true
	UserId string `json:"userId"`
	// 패스워드
	// Required: true
	Password string `json:"password"`
}
