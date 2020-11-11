//go:generate mockgen -source repository.go -destination mocks/csrf_repository_mock.go -package csrf_mock
package csrf

type Repository interface {
	Add(token string) error
	Check(token string) error
}
