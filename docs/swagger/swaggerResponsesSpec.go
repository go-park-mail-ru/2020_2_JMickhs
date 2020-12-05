package swagger

//unique data already exists
//swagger:response conflict
type Conflict struct {
}

//wrong csrf token
//swagger:response Forbidden
type Forbidden struct {
}

//cannot parse data or undefined query or path parameters
//swagger:response badrequest
type Badrequest struct {
}

//This data does not exist
//swagger:response gone
type Gone struct {
}

//user unauthorizied
//swagger:response unauthorizied
type Unauthorizied struct {
}

// Unsupported Media Type
//swagger:response unsupport
type Unsupport struct {
}

// wrong credentials
//swagger:response badCredentials
type Credentials struct {
}

// two times rate one hotel or want to get not your wishlists
// swagger:response locked
type Locked struct {
}

// wrong email
// swagger:response WrongEmail
type WrongEmail struct {
}
