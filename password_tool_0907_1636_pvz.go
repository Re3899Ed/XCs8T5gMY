// 代码生成时间: 2025-09-07 16:36:06
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// PasswordTool 提供了密码加密解密的功能
type PasswordTool struct {
    Key []byte
}

// NewPasswordTool 初始化 PasswordTool 并设置密钥
func NewPasswordTool(key []byte) *PasswordTool {
    return &PasswordTool{Key: key}
}

// Encrypt 将明文密码加密
func (pt *PasswordTool) Encrypt(password string) (string, error) {
    block, err := aes.NewCipher(pt.Key)
    if err != nil {
        return "", err
    }

    plaintext := []byte(password)
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 将密文密码解密
func (pt *PasswordTool) Decrypt(encryptedPassword string) (string, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(encryptedPassword)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(pt.Key)
    if err != nil {
        return "", err
    }

    if len(ciphertext) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)
    return string(ciphertext), nil
}

func main() {
    key := []byte("32-byte-long-key----")
    pt := NewPasswordTool(key)

    password := "mysecretpassword"

    encryptedPassword, err := pt.Encrypt(password)
    if err != nil {
        fmt.Println("Error encrypting password: ", err)
        return
    }
    fmt.Printf("Encrypted password: %s
", encryptedPassword)

    decryptedPassword, err := pt.Decrypt(encryptedPassword)
    if err != nil {
        fmt.Println("Error decrypting password: ", err)
        return
    }
    fmt.Printf("Decrypted password: %s
", decryptedPassword)
}
