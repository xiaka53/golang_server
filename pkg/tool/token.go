package tool

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// token生成
func GenerateToken(id, length int) (string, error) {
	// 获取当前时间的 Unix 时间戳（秒级）
	currentTime := time.Now().Unix()

	// 转换时间戳为字节序列
	timeBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		timeBytes[7-i] = byte(currentTime >> (8 * i))
	}

	// 转换用户ID为字节序列
	userIDBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		userIDBytes[7-i] = byte(id >> (8 * i))
	}

	// 计算随机部分的字节数
	randomByteLength := (length - 24) / 2 // 时间戳和用户ID各占用8个字节，一个十六进制字符占用2个字节

	// 生成随机字节序列
	randomBytes := make([]byte, randomByteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 合并时间戳、用户ID和随机字节序列
	tokenBytes := append(timeBytes, userIDBytes...)
	tokenBytes = append(tokenBytes, randomBytes...)

	// 将字节序列转换为十六进制字符串
	token := hex.EncodeToString(tokenBytes)

	return token, nil
}
