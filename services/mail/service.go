package mail

// Service interface
type Service interface {
	Send(*Mail) error
}
