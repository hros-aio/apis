package tenants

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/psql/common/tenant"
	"github.com/hros-aio/apis/libs/saga"
	"github.com/hros-aio/apis/libs/saga/events"
	"github.com/hros-aio/apis/libs/saga/messages"
	"github.com/tinh-tinh/tinhtinh/v2/core"
	"github.com/tinh-tinh/tinhtinh/v2/middleware/logger"
)

type TenantService struct {
	tenantRepo     *tenant.Repository
	eventPublisher *saga.EventPulisher
	logger         *logger.Logger
}

func NewService(module core.Module) core.Provider {
	tenantRepo := module.Ref(tenant.REPOSITORY).(*tenant.Repository)
	eventPublisher := module.Ref(saga.EVENT_PUBLISHER).(*saga.EventPulisher)
	logger := logger.InjectLog(module)

	return module.NewProvider(&TenantService{
		tenantRepo:     tenantRepo,
		eventPublisher: eventPublisher,
		logger:         logger,
	})
}

func (s *TenantService) Create(ctx middleware.ContextInfo, input *TenantCreateInput) (*tenant.TenantModel, error) {
	model := input.Dto()
	tenantId, err := GenerateRandomString(6)
	if err != nil {
		s.logger.Errorf("Failed to generate random: %s", err.Error())
		return nil, err
	}
	model.TenantId = tenantId
	data, err := s.tenantRepo.Create(model)
	if err != nil {
		s.logger.Errorf("Failed to create tenant: %s", err.Error())
		return nil, err
	}
	go s.eventPublisher.Publish(events.TenantCreated, messages.TenantCreatedPayload{
		Id:        data.ID.String(),
		Name:      data.Contact.ContactName,
		CreatedAt: data.CreatedAt,
		TenantId:  data.TenantId,
		Domain:    data.Domain,
		Contact:   messages.ContactPerson(data.Contact),
	})
	return data, nil
}

func (s *TenantService) List(ctx middleware.ContextInfo) ([]tenant.TenantModel, error) {
	return s.tenantRepo.List()
}

func GenerateRandomString(n int) (string, error) {
	// The character set to use for generating the random string.
	// This version only uses lowercase letters and numbers.
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	sb.Grow(n)

	// Get the length of the character set.
	charsetLen := big.NewInt(int64(len(letters)))

	for i := 0; i < n; i++ {
		// Generate a random index for the character set.
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		// Append the character at the random index to the string builder.
		sb.WriteByte(letters[randomIndex.Int64()])
	}

	return sb.String(), nil
}
