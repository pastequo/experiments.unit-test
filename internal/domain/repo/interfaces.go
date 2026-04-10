package repo

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces_mock.go -package=mocks

type Actor interface {
	Act() error
}

type Doer interface {
	Do() error
}
