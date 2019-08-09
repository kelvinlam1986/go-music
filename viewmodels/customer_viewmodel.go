package viewmodels

type CustomerGetByIdVm struct {
	Id uint `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	LoggedIn bool `json:"loggedIn"`
}
