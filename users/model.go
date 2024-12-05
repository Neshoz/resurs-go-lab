package users

type MonetaryAmount struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type Link struct {
	Href string `json:"href"`
	Access string `json:"access"`
}

type ContactPreferences struct {
	ContactName string `json:"contactName"`
	PhoneNumber string `json:"phoneNumber"`
	Email string `json:"email"`
}

type UserRole struct {
	Name string `json:"name"`
	Rank int `json:"rank"`
	DisposalRight string `json:"DisposalRight"`
}

type Department struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TotalApprovedLimit MonetaryAmount `json:"totalApprovedLimit"`
	AvailableLimit MonetaryAmount `json:"availableLimit"`
	Visible bool `json:"visible"`
	Links struct {
		Details Link `json:"details"`
		Update Link `json:"update"`
	} `json:"_links"`
}

type User struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Title string `json:"title"`
	OrganisationName string `json:"organisationName"`
	ContactPreferences ContactPreferences `json:"contactPreferences"`
	UserRole UserRole`json:"userRole"`
	ViewRight string `json:"viewRight"`
	EditRight string `json:"editRight"`
	TotalLimit MonetaryAmount `json:"totalLimit"`
	ConsumedLimit MonetaryAmount `json:"consumedLimit"`
	AvailableLimit MonetaryAmount `json:"availableLimit"`
	Administrators []string `json:"administrators"`
	Departments []Department `json:"departments"`
	Links struct {
		UpdateMe Link `json:"updateMe"`
	} `json:"_links"`
}