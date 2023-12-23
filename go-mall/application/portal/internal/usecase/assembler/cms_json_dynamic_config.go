package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// JsonDynamicConfigEntityToModel entity转pb
func JsonDynamicConfigEntityToModel(jsonDynamicConfig *entity.JsonDynamicConfig) *pb.JsonDynamicConfig {
	return &pb.JsonDynamicConfig{}
}
