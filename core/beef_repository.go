package core

type BeefRepository interface {
	GetData() ([]byte, error)
}
