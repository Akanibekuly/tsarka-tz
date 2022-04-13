package entities

type UserSt struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserUpdateSt struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
