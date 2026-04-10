package observability

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces_mock.go -package=mocks

type Metrics interface {
	Increment(name string)
}
