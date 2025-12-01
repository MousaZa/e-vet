package interfaces

type Observer interface {
	Update(string)
	GetID() string
}
