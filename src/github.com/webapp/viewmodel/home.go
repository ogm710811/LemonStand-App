package viewmodel

// Home Base struct
type Home struct {
	Title  string
	Active string
}

// NewHome constructor
func NewHome() Home {
	return Home{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
}
