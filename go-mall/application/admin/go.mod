module github.com/baker-yuan/go-mall/application/admin

go 1.20

require (
	github.com/baker-yuan/go-mall/common v0.0.0-incompatible
	github.com/baker-yuan/go-mall/proto v0.0.0-incompatible
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/ilyakaznacheev/cleanenv v1.2.6
	github.com/rs/cors v1.10.1
	github.com/swaggo/swag v1.7.6
	golang.org/x/net v0.19.0
	google.golang.org/grpc v1.59.0
	gorm.io/gorm v1.25.5
)

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/baker-yuan/go-mall/common => ../../common
	github.com/baker-yuan/go-mall/proto => ../../proto
)
