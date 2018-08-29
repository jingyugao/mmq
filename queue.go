package mmq

type Queue interface {
	Put(string)
	Consume() string
}
