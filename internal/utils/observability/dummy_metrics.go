package observability

type DummyMetrics struct{}

func (d DummyMetrics) Increment(_ string) {
}
