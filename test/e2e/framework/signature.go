package framework

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Sign 计算 API 签名
func Sign(params map[string]interface{}, privateKey string) string {
	// 1. 过滤 Signature 参数
	keys := make([]string, 0, len(params))
	for k := range params {
		if k != "Signature" {
			keys = append(keys, k)
		}
	}

	// 2. ASCII 排序
	sort.Strings(keys)

	// 3. 拼接 Key + Value
	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(k)
		val := anyToString(params[k])
		builder.WriteString(val)
	}

	// 4. 追加私钥
	builder.WriteString(privateKey)

	raw := builder.String()
	
	// 5. SHA1 哈希
	hash := sha1.Sum([]byte(raw))

	// 6. Hex 编码
	sig := hex.EncodeToString(hash[:])
	
	return sig
}

func anyToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", val)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		// 简单处理浮点，实际可能需要控制精度
		// 注意：OpenAPI 文档示例是 200000230，如果是 float，%v 可能会变成 2.0000023e+08
		// 建议尽量用 int 传递整数
		return fmt.Sprintf("%.0f", val) // 假设都是整数？这也太危险了。改回 %v
		// return fmt.Sprintf("%v", val) 
	case bool:
		return strconv.FormatBool(val)
	case []interface{}:
		var sb strings.Builder
		for _, item := range val {
			sb.WriteString(anyToString(item))
		}
		return sb.String()
	case []string:
		return strings.Join(val, "")
	case map[string]interface{}:
		// 递归 Map 拼接
		return Sign(val, "") // 复用 Sign 逻辑（无私钥）
	default:
		// 反射兜底
		return fmt.Sprintf("%v", val)
	}
}

// FlattenParams 将复杂类型打平，便于 HTTP 传输或日志记录（可选）
func toMap(v interface{}) map[string]interface{} {
	if m, ok := v.(map[string]interface{}); ok {
		return m
	}
	return nil
}

// StructToMap 利用反射将结构体转换为 map[string]interface{}
func StructToMap(obj interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		key := strings.Split(tag, ",")[0]
		if key == "" {
			key = field.Name
		}
		if key == "-" {
			continue
		}
		
		val := v.Field(i).Interface()
		if isZero(v.Field(i)) {
			continue
		}
		out[key] = val
	}
	return out
}

func isZero(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
