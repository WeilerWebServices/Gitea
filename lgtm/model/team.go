package model

// Team represents a team from the the remote API.
type Team struct {
	Login  string `json:"login"`
	Avatar string `json:"avatar"`
}

// Member represents a member from the the remote API.
type Member struct {
	Login string `json:"login"`
}
