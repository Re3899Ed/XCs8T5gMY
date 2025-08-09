// 代码生成时间: 2025-08-09 08:39:14
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "revel"
)

// PasswordTool 用于密码加密和解密的结构体
type PasswordTool struct {
    Key []byte
}

// NewPasswordTool 创建并初始化一个PasswordTool实例
func NewPasswordTool(key []byte) *PasswordTool {
    return &PasswordTool{Key: key}
}

// Encrypt 加密密码
func (p *PasswordTool) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(p.Key)
    if err != nil {
        return "", err
    }

    // 填充至块大小的倍数
    plaintextBytes := PKCS5Padding([]byte(plaintext), aes.BlockSize)

    // 加密
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)

    // 编码为base64字符串
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密密码
func (p *PasswordTool) Decrypt(ciphertext string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(p.Key)
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("ciphertext too short")
    }
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintextBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    // 移除填充
    plaintext := string(PKCS5UnPadding(plaintextBytes))
    return plaintext, nil
}

// PKCS5Padding 填充至块大小的倍数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

// PKCS5UnPadding 移除填充
func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func main() {
    revel.RunProd()
}

// Actions 是Revel框架的控制器
type Actions struct {
    *revel.Controller
}

//加密操作
func (c Actions) EncryptPassword() revel.Result {
    key := []byte("your-secret-key") // 密钥必须是AES块大小的倍数，例如16, 24, 32字节
    passwordTool := NewPasswordTool(key)
    plaintext := c.Params.Get("plaintext")
    encrypted, err := passwordTool.Encrypt(plaintext)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(revel.Result{
        "filename": "password_tool.go",
        "code": "// 加密结果
" + encrypted,
    })
}

//解密操作
func (c Actions) DecryptPassword() revel.Result {
    key := []byte("your-secret-key") // 密钥必须是AES块大小的倍数，例如16, 24, 32字节
    passwordTool := NewPasswordTool(key)
    ciphertext := c.Params.Get("ciphertext")
    decrypted, err := passwordTool.Decrypt(ciphertext)
    if err != nil {
        return c.RenderError(err)
    }
    return c.RenderJSON(revel.Result{
        "filename": "password_tool.go",
        "code": "// 解密结果
" + decrypted,
    })
}
