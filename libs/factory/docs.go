package factory

import (
	"strings"

	"github.com/tinh-tinh/swagger/v2"
)

func DefaultSwagger(prefix string) *swagger.SpecBuilder {
	name := formatNamePrefix(prefix)
	return swagger.NewSpecBuilder().
		SetTitle(name + " Doc").
		SetDescription(name + " API documentation").
		SetVersion("v1.0.0").
		AddSecurity(&swagger.SecuritySchemeObject{
			Type:         "http",
			Scheme:       "Bearer",
			BearerFormat: "JWT",
			Name:         "bearerAuth",
		})
}

func formatNamePrefix(input string) string {
	// Split the input string by hyphen
	parts := strings.Split(input, "-")

	var formattedParts []string
	for _, part := range parts {
		// Trim any whitespace from the part
		trimmedPart := strings.TrimSpace(part)
		if trimmedPart == "" {
			continue // Skip empty parts if any
		}

		// Convert "api" to "API" specifically
		if strings.ToLower(trimmedPart) == "api" {
			formattedParts = append(formattedParts, strings.ToUpper(trimmedPart))
		} else {
			// Capitalize the first letter and make the rest lowercase
			// This handles cases like "admin" -> "Admin"
			formattedParts = append(formattedParts, strings.ToUpper(trimmedPart[:1])+strings.ToLower(trimmedPart[1:]))
		}
	}

	// Join the formatted parts with a space
	return strings.Join(formattedParts, " ")
}
