echo "生成 rpc 代码"

OUT=./

# admin

protoc \
-I ${OUT} \
-I model/*.proto \
--go_out=":./mall/"  \
--go-grpc_out=":./mall/"  \
--grpc-gateway_out=":./mall/" \
--validate_out="lang=go:./mall/" \
admin/*.proto model/*.proto


# model

#protoc \
#-I ${OUT} \
#--go_out=":./model/"  \
#--go-grpc_out=":./model/"  \
#--grpc-gateway_out=":./model/" \
#--validate_out="lang=go:./model/" \
#model/*.proto

#protoc \
#-I ${OUT} \
#-I validate/ \
#--go_out=":./model/" \
#--validate_out="lang=go:./model/" \
#model/pms_product_category.proto
#
#protoc \
#--go_out=${OUT} \
#--go-grpc_out=${OUT} \
#--go-grpc_opt=require_unimplemented_servers=false \
#model/pms_product_category.proto



