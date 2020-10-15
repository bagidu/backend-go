//go:generate mockery  --inpackage --name Service

package auth

// Service interface
type Service interface {
	Login() error
}
