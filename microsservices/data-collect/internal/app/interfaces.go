package app

type Collector interface {
	Collect() error
}
