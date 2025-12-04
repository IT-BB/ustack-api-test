package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/onsi/ginkgo/v2"
)

// Client 封装 API 调用（基于原生 net/http）
type Client struct {
	HttpClient *http.Client
	Config     *Config
}

// NewClient 创建新的 API 客户端
func NewClient(cfg *Config) *Client {
	return &Client{
		HttpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		Config: cfg,
	}
}

// BaseRequest 包含所有请求的公共参数
type BaseRequest struct {
	Action    string `json:"Action"`
	Region    string `json:"Region,omitempty"`
	CompanyID string `json:"CompanyID,omitempty"`
	ProjectID string `json:"ProjectID,omitempty"`
	PublicKey string `json:"PublicKey"`
	Signature string `json:"Signature"`
}

// BaseResponse 包含所有响应的公共字段
type BaseResponse struct {
	Action  string `json:"Action"`
	RetCode int    `json:"RetCode"`
	Message string `json:"Message"`
}

// Do 发送 API 请求
func (c *Client) Do(ctx context.Context, action string, payload map[string]interface{}, result interface{}) error {
	// 1. 注入公共参数
	if _, ok := payload["Action"]; !ok {
		payload["Action"] = action
	}
	if _, ok := payload["Region"]; !ok && c.Config.Region != "" {
		payload["Region"] = c.Config.Region
	}
	if _, ok := payload["CompanyID"]; !ok && c.Config.CompanyID != "" {
		payload["CompanyID"] = c.Config.CompanyID
	}
	if _, ok := payload["ProjectID"]; !ok && c.Config.ProjectID != "" {
		payload["ProjectID"] = c.Config.ProjectID
	}
	payload["PublicKey"] = c.Config.PublicKey

	// 2. 扁平化处理 Array 参数 (Key: ["a","b"] -> Key.0: "a", Key.1: "b")
	// 注意：这是为了适配后端对 Array 参数的特殊签名和 Payload 格式要求
	flatPayload := flattenPayload(payload)

	// 3. 计算签名 (基于扁平化后的参数)
	flatPayload["Signature"] = Sign(flatPayload, c.Config.PrivateKey)

	// 4. 序列化 Request Body (使用扁平化后的 Payload)
	reqBody, err := json.Marshal(flatPayload)
	if err != nil {
		return fmt.Errorf("marshal payload failed: %w", err)
	}

	// 5. 构建 Request
	url := c.Config.BaseURL
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 6. 发送请求
	start := time.Now()
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		ginkgo.GinkgoWriter.Printf("[API FAIL] POST %s Action=%s Err=%v\n", url, action, err)
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	duration := time.Since(start)

	// 7. 读取 Response Body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body failed: %w", err)
	}

	// HTTP Error
	if resp.StatusCode >= 400 {
		ginkgo.GinkgoWriter.Printf("[API HTTP FAIL] Action=%s Status=%d Duration=%v\nPAYLOAD: %s\nBODY: %s\n",
			action, resp.StatusCode, duration, string(reqBody), string(respBody))
		return fmt.Errorf("http error: %d %s", resp.StatusCode, string(respBody))
	}

	// 8. 反序列化
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			ginkgo.GinkgoWriter.Printf("[API JSON FAIL] Action=%s Body=%s Err=%v\n", action, string(respBody), err)
			return fmt.Errorf("unmarshal response failed: %w, body: %s", err, string(respBody))
		}

		// 检查业务 RetCode
		var tmp map[string]interface{}
		if err := json.Unmarshal(respBody, &tmp); err == nil {
			if rc, ok := tmp["RetCode"]; ok {
				code := 0
				switch v := rc.(type) {
				case float64:
					code = int(v)
				case int:
					code = v
				}
				if code != 0 {
					msg, _ := tmp["Message"].(string)
					ginkgo.GinkgoWriter.Printf("[API RET FAIL] Action=%s RetCode=%d Message=%s\n",
						action, code, msg)
				}
			}
		}
	}

	return nil
}

// flattenPayload 递归将 Map 中的 Slice/Array 展开为 Key.Index 格式
// 同时处理 Base64 编码等特殊逻辑（如果有）
// 注意：这里只处理一层 Array，因为 API 文档似乎没有嵌套 Array
func flattenPayload(input map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(input))
	for k, v := range input {
		// 检查是否为 Slice/Array
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			// 如果是 []byte，通常是 JSON 里的 String (Base64)，不需要展开
			if val.Type().Elem().Kind() == reflect.Uint8 {
				out[k] = v
				continue
			}
			
			// 展开 Array
			for i := 0; i < val.Len(); i++ {
				key := fmt.Sprintf("%s.%d", k, i)
				out[key] = val.Index(i).Interface()
			}
		} else {
			out[k] = v
		}
	}
	return out
}

// GenericResponse 用于无需解析详细字段的请求
type GenericResponse struct {
	BaseResponse
}
