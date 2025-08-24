package middleware

import (
	"slices"
	"strings"

	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const PAGINATE core.CtxKey = "Paginate"

type PaginateInput struct {
	Page  int `json:"page" query:"page" example:"1"`
	Limit int `json:"limit" query:"limit" example:"10"`
}

type Paginate struct {
	Skip  int
	Limit int
	Sort  map[string]string
}

func Pagination(ctx core.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	paginate := &Paginate{
		Limit: limit,
		Skip:  (page - 1) * limit,
	}

	sortStr := ctx.Query("sort")
	if sortStr != "" {
		sorts := strings.Split(sortStr, ",")
		sortMapper := make(map[string]string)
		for _, sort := range sorts {
			keyVal := strings.Split(sort, ":")
			if !isValidSortValue(keyVal[1]) {
				continue
			}
			sortMapper[keyVal[0]] = keyVal[1]
		}

		if len(sortMapper) > 0 {
			paginate.Sort = sortMapper
		}
		sortMapper = nil
	}

	ctx.Set(PAGINATE, paginate)
	return ctx.Next()
}

func isValidSortValue(key string) bool {
	validKey := []string{"ASC", "DESC", "1", "-1"}
	return slices.Index(validKey, key) > -1
}
