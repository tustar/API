package middlewares

import (
	"strings"
	"sort"
	"fmt"
	"crypto/x509"
	"crypto/rsa"
	"encoding/base64"
	"crypto/sha1"
	"crypto"
	"os"
	"bufio"
	"io"
	"ushare/logger"
)

func BuildSignContent(params map[string]interface{}) string {
	// STEP1, 过滤sign参数不验签
	result := ParamsFilter(params)
	//
	return CreateLinkString(result)
}

func ParamsFilter(params map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range params {
		if strings.EqualFold(k, "sign") {
			continue
		}
		fmt.Sprintf("%v = %v\n", k, v)
		result[k] = v
	}
	return result
}

func CreateLinkString(params map[string]interface{}) string {
	// STEP2, 对key进行升序排序
	sortedKeys := make([]string, 0)
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	//STEP3, 对key=value的键值对用&连接起来，略过空值
	var signString string
	for i, k := range sortedKeys {
		v := fmt.Sprintf("%v", params[k])
		if v != "" {
			if i > 0 {
				signString = signString + "&"
			}
			signString = signString + k + "=" + v
		}

	}
	return signString
}

func CheckSign(params map[string]interface{}, signData string) error {
	originalData := BuildSignContent(params)
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		logger.E("Base64 decode signData error: ", err)
		return err
	}

	key, err := GetPublicKey("./src/ushare/middlewares/public.pub")
	if err != nil {
		logger.E("Get public key error: ", err)
		return err
	}

	public, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		logger.E("Base64 decode public key error: ", err)
		return err
	}

	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		logger.E("x509.ParsePKIXPublicKey public error: ", err)
		return err
	}

	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}

func GetPublicKey(path string) (string, error) {

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := bufio.NewReader(f)

	var publicKey = ""
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.Contains(line, "-----") {
			continue
		}
		publicKey = strings.Join([]string{publicKey, line}, "")
		if err != nil {
			if err == io.EOF {
				return publicKey, nil
			}
			return "", err
		}
	}

	return publicKey, nil
}
