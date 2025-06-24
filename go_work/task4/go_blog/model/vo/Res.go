package vo

import (
	"net/http"
)

const (
	CodeSuccess      = 200 // 成功
	CodeParamError   = 201 // 成功
	CodeBadRequest   = 400 // 请求错误
	CodeUnauthorized = 401 // 未授权
	CodeForbidden    = 403 // 禁止访问
	CodeNotFound     = 404 // 资源不存在
	CodeServerError  = 500 // 服务器错误
)

// 统一响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 消息
	Data    T      `json:"data"`    // 数据(泛型)
}

// 成功响应
func Success[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

// 失败响应
func Error(code int, message string) *Response[any] {
	return &Response[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 带分页的响应
type PageResponse[T any] struct {
	Response[T]
	Page     int `json:"page"`      // 当前页码
	PageSize int `json:"page_size"` // 每页数量
	Total    int `json:"total"`     // 总数
}

// 成功分页响应
func PageSuccess[T any](data T, page, pageSize, total int) *PageResponse[T] {
	return &PageResponse[T]{
		Response: Response[T]{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data,
		},
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}
}
