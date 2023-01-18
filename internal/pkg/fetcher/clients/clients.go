//go:generate mockery --case underscore --name RegistratorServiceClient --with-expecter

package clients

import "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

// RegistratorServiceClient ...
type RegistratorServiceClient interface {
	registrator.RegistratorServiceClient
}
