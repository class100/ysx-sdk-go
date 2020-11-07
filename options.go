package ysx

import (
	"github.com/class100/core"
)

type (
	// Option 配置选项
	Option func(*options)

	options struct {
		options []core.Option

		// Endpoint 服务端点
		Endpoint string
	}
)

func defaultOptions() options {
	return options{}
}

// WithSecret 配置应用授权
func WithSecret(secret core.Secret) Option {
	return func(options *options) {
		options.options = append(options.options, core.WithSecret(secret))
	}
}

// WithAlgorithms 配置签名算法
func WithAlgorithms(algorithms ...core.SignatureAlgorithm) Option {
	return func(options *options) {
		options.options = append(options.options, core.WithAlgorithms(algorithms...))
	}
}

// WithEndpoint 配置服务端点
func WithEndpoint(endpoint string) Option {
	return func(options *options) {
		options.Endpoint = endpoint
	}
}
