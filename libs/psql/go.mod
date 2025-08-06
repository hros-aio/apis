module github.com/hros-aio/apis/libs/psql

go 1.24.1

require (
	github.com/google/uuid v1.6.0
	github.com/hros-aio/apis/libs/factory v0.0.1
	github.com/tinh-tinh/config/v2 v2.1.0
	github.com/tinh-tinh/sqlorm/v2 v2.3.1
	github.com/tinh-tinh/tinhtinh/v2 v2.3.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.30.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/hros-aio/apis/libs/factory => ../factory
