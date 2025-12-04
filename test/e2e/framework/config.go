package framework

import (
	"os"
	"time"
)

// Config 包含全局测试配置
type Config struct {
	BaseURL    string
	PublicKey  string
	PrivateKey string

	// 默认上下文，若环境变量未设置则为空，测试运行时需自动探测或跳过
	Region    string
	CompanyID string
	ProjectID string

	Timeout time.Duration
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	cfg := &Config{
		BaseURL:    os.Getenv("USTACK_API_BASE_URL"),
		PublicKey:  os.Getenv("USTACK_API_PUBLIC_KEY"),
		PrivateKey: os.Getenv("USTACK_API_PRIVATE_KEY"),
		Region:     os.Getenv("USTACK_API_REGION"),
		CompanyID:  os.Getenv("USTACK_API_COMPANY_ID"),
		ProjectID:  os.Getenv("USTACK_API_PROJECT_ID"),
		Timeout:    30 * time.Second,
	}

	if t := os.Getenv("USTACK_API_TIMEOUT"); t != "" {
		if d, err := time.ParseDuration(t); err == nil {
			cfg.Timeout = d
		}
	}

	return cfg
}

// IsReady 检查最小运行环境是否满足
func (c *Config) IsReady() bool {
	return c.BaseURL != "" && c.PublicKey != "" && c.PrivateKey != ""
}
