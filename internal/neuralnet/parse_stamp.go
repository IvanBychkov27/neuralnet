package neuralnet

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type stampData struct {
	TCPWindowSize   string
	IPTTL           string
	IPFlags         string
	TCPFlags        string
	TCPHeaderLength string
	TCPOptions      string
	MSS             string
}

func parseStamp(stamp string) stampData {
	d := stampData{}
	s := strings.Split(stamp, ";")
	if len(s) != 6 {
		return d
	}

	d.TCPWindowSize = s[0]
	d.IPTTL = s[1]
	d.IPFlags = getIPFlags(s[2])
	d.TCPFlags = s[3]
	d.TCPHeaderLength = s[4]
	d.TCPOptions = convertHexInDec(md5Data(s[5]))
	d.MSS = getMSS(s[5])

	return d
}

func getIPFlags(f string) string {
	switch f {
	case "DF":
		return "9"
	case "MF":
		return "5"
	}
	return "0"
}

func getMSS(dTCPOptions string) string {
	ms := strings.Split(dTCPOptions, ",")
	if len(ms) < 1 {
		return ""
	}
	mss := ms[0]
	if mss == "" {
		return ""
	}
	return mss[1:]
}

func md5Data(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func convertHexInDec(hex string) string {
	res := 0

	for i := 0; i < 22; i += 11 {
		d, err := convertInt(hex[i:i+11], 16, 10)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		r, _ := strconv.Atoi(d)
		res += r
	}

	return strconv.Itoa(res)
}

// convertInt конвертирует значение из одной системы счисления в другую, которая указана в toBase
// https://golangify.com/binary-to-decimal-octal-and-hexadecimal
func convertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
