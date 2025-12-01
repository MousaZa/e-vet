package interfaces

type Subject interface {
	Register(observer Observer)
	DeRegister(observer Observer)
	NotifyAll()
}
