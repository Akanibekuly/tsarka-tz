package interfaces

type Cacher interface {
	Add(n int) error
	Sub(n int) error
	Val() (int, error)
}
