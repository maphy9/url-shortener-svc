package data

type MasterQ interface {
	Mapping() MappingQ

	Transaction(fn func(db MasterQ) error) error
}