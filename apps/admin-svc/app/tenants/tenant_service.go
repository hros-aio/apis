package tenants

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/hros-aio/apis/libs/factory/middleware"
	"github.com/hros-aio/apis/libs/sql/common/tenant"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

type TenantService struct {
	tenantRepo *tenant.Repository
}

func NewService(module core.Module) core.Provider {
	tenantRepo := module.Ref(tenant.REPOSITORY).(*tenant.Repository)

	return module.NewProvider(&TenantService{
		tenantRepo: tenantRepo,
	})
}

func (s *TenantService) Create(ctx middleware.ContextInfo, input *TenantCreateInput) (*tenant.TenantModel, error) {
	model := &tenant.TenantDB{
		Name:        input.Name,
		Description: input.Description,
		Contact: tenant.ContactPersonDb{
			ContactName:  input.Contact.Name,
			ContactEmail: input.Contact.Email,
			ContactPhone: input.Contact.Phone,
		},
	}
	tenantId, err := GenerateRandomString(6)
	if err != nil {
		return nil, err
	}
	model.TenantId = tenantId
	data, err := s.tenantRepo.Create(model)
	if err != nil {
		return nil, err
	}
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
