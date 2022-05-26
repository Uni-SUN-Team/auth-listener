package models

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Agreement bool   `json:"agreement"`
}

type RegisterStrapiResponse struct {
	User RegisterStrapiResponseUserDatail `json:"user"`
}

type RegisterStrapiResponseUserDatail struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Provider   string `json:"provider"`
	Confirmed  bool   `json:"confirmed"`
	Blocked    bool   `json:"blocked"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Telephone  string `json:"telephone"`
	Agreement  string `json:"agreement"`
	Middlename string `json:"middlename"`
}
