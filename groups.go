package groups

type Callback func() error

type Group interface {
	Go(Callback)
	Wait() error
}
