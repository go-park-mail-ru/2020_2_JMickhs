package swagger

//unique data already exists
//swagger:response conflict
type conflict struct {
}

//wrong csrf token
//swagger:response Forbidden
type forbidden struct {
}

//cannot parse data or undefined query or path parameters
//swagger:response badrequest
type badrequest struct {
}

//This data does not exist
//swagger:response gone
type gone struct {
}

//user unauthorizied
//swagger:response unauthorizied
type unauthorizied struct {
}

// Unsupported Media Type
//swagger:response unsupport
type unsupport struct {
}

// wrong credentials
//swagger:response badCredentials
type credentials struct {
}

// two times rate one hotel
// swagger:response locked
type locked struct {
}

// wrong email
// swagger:response WrongEmail
type wrongEmail struct {
}
