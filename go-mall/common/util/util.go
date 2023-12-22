package util

import (
	"github.com/baker-yuan/go-mall/common/util/conversion"
	"github.com/baker-yuan/go-mall/common/util/ctx"
	"github.com/baker-yuan/go-mall/common/util/decimal"
	"github.com/baker-yuan/go-mall/common/util/img"
	"github.com/baker-yuan/go-mall/common/util/str"
)

// https://gitee.com/baker-yuan/gotool/blob/master/tool.go

var (
	StrUtils            str.StrUtils               // 字符串操作
	TypeConversionUtils conversion.ConversionUtils // 类型转换，用于string，int，int64，float等数据转换，免去err的接收，和设置默认值
	ImgUtils            img.ImgUtils               // 图片路径处理
	CtxUtils            ctx.CtxUtils               // 上下文处理
	DecimalUtils        decimal.DecimalUtils       // 浮点数计算
)
