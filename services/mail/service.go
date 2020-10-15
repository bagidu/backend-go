//go:generate mockery --inpackage --name Service

package mail

// Service interface
type Service interface {
	Send(*Mail) error
}
