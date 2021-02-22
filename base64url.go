package toolkit

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Base64URLEncode encode source data into base64url code
func Base64URLEncode(src []byte) string {
	res := base64.StdEncoding.EncodeToString(src)
	res = strings.Replace(res, "+", "-", -1)
	res = strings.Replace(res, "/", "_", -1)
	res = strings.Replace(res, "=", "", -1)
	return res
}

// Base64URLDecode decode base64url code into original data bytes
func Base64URLDecode(code string) (data []byte, err error) {
	code = strings.Replace(code, "-", "+", -1)
	code = strings.Replace(code, "_", "/", -1)
	switch len(code) % 4 {
	case 1:
		return data, errors.New("invalid base64url code length")
	case 2:
		code += "=="
	case 3:
		code += "="
	}
	return base64.StdEncoding.DecodeString(code)
}
