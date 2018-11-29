package viewmodel

// Login Struct
type Login struct {
	Title    string
	Active   string
	Email    string
	Password string
}

// NewLogin constructor
func NewLogin() Login {
	result := Login{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
	return result
}
