module github.com/baker-yuan/go-mall/application/portal

go 1.20

require (
	github.com/baker-yuan/go-mall/common v0.0.0-00010101000000-000000000000
	github.com/baker-yuan/go-mall/proto v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/ilyakaznacheev/cleanenv v1.2.5
	github.com/rs/cors v1.10.1
	golang.org/x/net v0.17.0
	google.golang.org/grpc v1.59.0
	gorm.io/gorm v1.25.5
)

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.2 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/baker-yuan/go-mall/common => ../../common
	github.com/baker-yuan/go-mall/proto => ../../proto
)
