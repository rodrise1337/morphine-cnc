package events

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Secret key for HMAC (должен совпадать с ключом на сервере)
var secretKey = []byte("supersecretkey")

// VerifyHMAC verifies the HMAC signature
func VerifyHMAC(message, signature string) bool {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(message))
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return expectedSignature == signature
}

// CheckLicense отправляет запрос на проверку лицензии и возвращает результат
func CheckLicense(licenseKey, product, clientIP string) (bool, error) {
	url := "http://54.38.62.45:8080/check"
	requestBody, _ := json.Marshal(map[string]string{
		"key":     licenseKey,
		"product": product,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Raw Response:", string(body))

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return false, fmt.Errorf("error parsing response: %w. Raw Response: %s", err, string(body))
	}

	// Извлечение сигнатуры и проверяемых данных
	signature, ok := response["signature"].(string)
	if !ok {
		return false, fmt.Errorf("signature missing in response")
	}

	expiry := response["expiry"].(string)
	products := []string{}
	for _, product := range []string{"cnc", "bss", "funnel"} {
		if response[product] == "true" {
			products = append(products, product)
		}
	}

	message := fmt.Sprintf("%s|%s|%s", expiry, strings.Join(products, ","), clientIP)
	if VerifyHMAC(message, signature) {
		return true, nil
	}

	return false, fmt.Errorf("signature verification failed")

}
