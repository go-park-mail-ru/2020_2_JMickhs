//go:generate mockgen -source usecase.go -destination mocks/csrf_usecase_mock.go -package csrf_mock
package csrf


type Usecase interface {
	CreateToken (sesID string, timeStamp int64 ) (string,error)
    CheckToken (sesID string, token string  ) (bool, error)
}
