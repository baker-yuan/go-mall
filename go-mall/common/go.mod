module github.com/baker-yuan/go-mall/common

go 1.20

require (
    github.com/baker-yuan/go-mall/proto v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1
	github.com/rs/zerolog v1.31.0
	github.com/shopspring/decimal v1.3.1
	golang.org/x/crypto v0.14.0
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848
	google.golang.org/grpc v1.59.0
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
replace (
	github.com/baker-yuan/go-mall/proto => ../proto
)