package shared

type ParamID struct {
	ID string `path:"id" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
}

type QueryCompany struct {
	CompanyID string `query:"companyId" example:"6cdad833-ba6d-49e3-889c-da23b764bb21"`
}
