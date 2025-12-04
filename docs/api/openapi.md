# OpenAPI 签名生成

为了保证 API 请求的安全性，所有请求都必须包含签名。本文档描述了如何生成签名。

## 签名算法

签名生成步骤如下：

0.  **获取密钥**:
    *   用户可以通过调用 `DescribeUser` 接口获取公钥 (PublicKey) 和私钥 (PrivateKey)。

1.  **准备参数**:
    *   收集所有请求参数。
    *   确保包含 `PublicKey`。
    *   移除所有值为空的参数。
    *   排除 `Signature` 参数本身。

2.  **键排序**:
    *   按照 ASCII 码对参数键进行升序排序。

3.  **拼接字符串**:
    *   遍历排序后的键。
    *   对于每个键，将 **Key** 和对应的 **Value** 拼接到字符串缓冲区。
    *   **值转换规则**:
        *   **基本类型** (string, int, bool 等): 转换为字符串表示。
        *   **数组/切片**: 拼接所有元素的字符串表示。
        *   **Map**: 递归应用排序和拼接逻辑 (Key + Value)。

4.  **追加私钥**:
    *   在拼接好的字符串末尾追加用户的 `PrivateKey`。

5.  **计算哈希**:
    *   计算结果字符串的 **SHA1** 哈希值。

6.  **Hex 编码**:
    *   将二进制哈希值转换为十六进制字符串。这就是最终的签名。

## 示例

假设我们有以下请求参数：

*   `Action`: `DeleteVMInstance`
*   `Region`: `cong-arm`
*   `CompanyID`: `200000230`
*   `VMID`: `vm-uf8mjntt2tqndp`

以及密钥信息：

*   `PublicKey`: `nDVv-arKQuZzS326dors0c1RFCgampVsL1Ppygy4aKt6bJrRM1BxiYHV`
*   `PrivateKey`: `stvC_notwaEnD9klufFttH24ormYM_m6OQT8TxN3Jln2XB0kFx3QbXcTTiIfksO5`

**步骤 1: 准备参数**

将所有参数（包括 `PublicKey`）收集起来。

**步骤 2: 键排序**

对参数键进行 ASCII 升序排序：
1.  `Action`
2.  `CompanyID`
3.  `PublicKey`
4.  `Region`
5.  `VMID`

**步骤 3: 拼接字符串**

按照排序后的顺序，将 Key 和 Value 拼接：

`ActionDeleteVMInstanceCompanyID200000230PublicKeynDVv-arKQuZzS326dors0c1RFCgampVsL1Ppygy4aKt6bJrRM1BxiYHVRegioncong-armVMIDvm-uf8mjntt2tqndp`

**步骤 4: 追加私钥**

在末尾追加 `PrivateKey`：

`ActionDeleteVMInstanceCompanyID200000230PublicKeynDVv-arKQuZzS326dors0c1RFCgampVsL1Ppygy4aKt6bJrRM1BxiYHVRegioncong-armVMIDvm-uf8mjntt2tqndpstvC_notwaEnD9klufFttH24ormYM_m6OQT8TxN3Jln2XB0kFx3QbXcTTiIfksO5`

**步骤 5 & 6: 计算哈希并编码**

对上述字符串进行 SHA1 计算，并转换为 Hex 字符串，得到最终签名：

`8adc30f47a1cd4f0850ec3ac3709ed45fe7e3d01`

## 参考实现 (Go)

以下是 Go 语言的参考实现。

```go
package signature

import (
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

// Sign generates the signature for the given input parameters and private key.
func Sign(input map[string]interface{}, privateKey string) string {
	str := map2String(input) + privateKey
	hashed := sha1.Sum([]byte(str))
	return hex.EncodeToString(hashed[:])
}

// map2String converts a map to a sorted concatenated string.
func map2String(params map[string]interface{}) (str string) {
	for _, k := range extractSortedKeys(params) {
		str += k + any2String(params[k])
	}
	return
}

// extractSortedKeys extracts and sorts keys from the map.
func extractSortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		if k == "Signature" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// any2String converts any type to its string representation.
func any2String(v interface{}) string {
	switch v := v.(type) {
	case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64,
		*string, *bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
		return simple2String(v)
	case []interface{}:
		return slice2String(v)
	case map[string]interface{}:
		return map2String(v)
	default:
		return reflectArr2String(reflect.ValueOf(v))
	}
}

// simple2String converts simple types to string.
func simple2String(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case *string:
		return *v
	case *bool:
		return strconv.FormatBool(*v)
	case *int:
		return strconv.FormatInt(int64(*v), 10)
	case *int8:
		return strconv.FormatInt(int64(*v), 10)
	case *int16:
		return strconv.FormatInt(int64(*v), 10)
	case *int32:
		return strconv.FormatInt(int64(*v), 10)
	case *int64:
		return strconv.FormatInt(*v, 10)
	case *uint:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint8:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint16:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint32:
		return strconv.FormatUint(uint64(*v), 10)
	case *uint64:
		return strconv.FormatUint(*v, 10)
	case *float32:
		return strconv.FormatFloat(float64(*v), 'f', -1, 64)
	case *float64:
		return strconv.FormatFloat(*v, 'f', -1, 64)
	}
	return ""
}

// slice2String converts a slice to a concatenated string.
func slice2String(arr []interface{}) (str string) {
	for _, v := range arr {
		switch v := v.(type) {
		case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64,
			*string, *bool, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
			str += simple2String(v)
		case []interface{}:
			str += slice2String(v)
		case map[string]interface{}:
			str += map2String(v)
		default:
			str += reflectArr2String(reflect.ValueOf(v))
		}
	}
	return
}

// reflectArr2String converts array/slice to string using reflection.
func reflectArr2String(rv reflect.Value) (str string) {
	if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < rv.Len(); i++ {
		str += any2String(rv.Index(i).Interface())
	}
	return
}

// sign prepares the data and calls Sign.
func sign(publicKey, privateKey string, data url.Values) string {
	data.Set("PublicKey", publicKey)
	var paramsMap = make(map[string]interface{})
	for k := range data {
		if data.Get(k) == "" {
			continue
		}
		paramsMap[k] = data.Get(k)
	}
	return Sign(paramsMap, privateKey)
}
```