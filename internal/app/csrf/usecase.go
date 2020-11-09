package csrf


type Usecase interface {
	CreateToken (sesID string, timeStamp int64 ) (string,error)
    CheckToken (sesID string, token string  ) (bool, error)
}
