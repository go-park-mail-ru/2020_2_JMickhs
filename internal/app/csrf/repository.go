package csrf

type Repository interface {
	Add(token string) error
	Check(token string) error
}
