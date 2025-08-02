package tenants

type ContactInput struct {
	Name  string `json:"name" example:"Renil"`
	Email string `json:"email" example:"renil@gmail.com" validate:"isEmail"`
	Phone string `json:"phone" example:"012345678"`
}

type TenantCreateInput struct {
	Name        string        `json:"name" example:"Netze" validate:"required"`
	Description string        `json:"description" example:"Netze Home Inc" validate:"required"`
	Contact     *ContactInput `json:"contact" validate:"nested"`
}
