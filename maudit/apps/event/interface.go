package event

type Service interface {
	SaveEvent() error
}
