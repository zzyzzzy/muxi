package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// 解析 extra_info
	extraInfo := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	decoded, _ := base64.StdEncoding.DecodeString(extraInfo)
	decodedStr := string(decoded)

	// 提取 secret_key 和 error_code
	parts := strings.Split(decodedStr, ", ")
	secretKey := strings.TrimPrefix(parts[0], "secret_key:")
	errorCode := strings.TrimPrefix(parts[1], "error_code:")

	// AES-256 加密
	encrypted, err := AES256Encrypt(errorCode, secretKey)
	if err != nil {
		panic(err)
	}

	// 发送请求
	sendRequest(encrypted, secretKey)
}

func AES256Encrypt(plaintext, key string) (string, error) {
	// 处理密钥为32字节
	keyBytes := []byte(key)
	if len(keyBytes) < 32 {
		padded := make([]byte, 32)
		copy(padded, keyBytes)
		keyBytes = padded
	} else {
		keyBytes = keyBytes[:32]
	}

	// 创建加密块
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// PKCS7填充
	plaintextBytes := []byte(plaintext)
	padding := aes.BlockSize - (len(plaintextBytes) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintextBytes = append(plaintextBytes, padText...)

	// 生成IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	mode.CryptBlocks(ciphertext, plaintextBytes)

	// 组合IV和密文
	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func sendRequest(encryptedErrorCode, secretKey string) {
	url := "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate"

	requestData := map[string]string{
		"error_code": encryptedErrorCode,
		"secret_key": secretKey,
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	req.Header.Set("code", "11")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiMTHpmr7pgZPmmK_kvaDlrrPkuobmiJEiLCJpYXQiOjE3NjI5MTg4NjksIm5iZiI6MTc2MjkxODg2OX0.koIQNVwXpMRFESk_aK1PtYs9RKcBFUXyY7SApsgLPPM")
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应: %s\n", string(body))
}
