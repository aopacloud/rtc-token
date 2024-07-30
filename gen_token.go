package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	// 这个是声网的加密算法，注意版本必须是2020年12月18日的，对应的commitID是53b7d2461db6，可以git checkout 53b7d2461db6
	dk5 "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/DynamicKey5"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// GenerateToken 生成token，服务器与客户端必须使用一样的版本，自研RTC与声网使用一样的加密算法
func GenerateToken(appId, appCert, cid, uid string) (string, error) {
	uidInt, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return "", fmt.Errorf("uid %s is not int", uid)
	}

	unixTs := time.Now().Unix()
	expiredTs := unixTs + 24*3600 // 24小时过期
	randomInt := uint32(random(1, 99999999))
	fmt.Printf("appId:%s, appCert:%s, cid:%s, unixTs:%d, randomInt:%d, uidInt:%d, expireTs:%d\n", appId, appCert, cid, unixTs, randomInt, uidInt, expiredTs)
	return dk5.GenerateMediaChannelKey(appId, appCert, cid, uint32(unixTs), randomInt, uint32(uidInt), uint32(expiredTs))
}

func main() {
	token, _ := GenerateToken("328579ea8dc94bc8ba45b25521f1e3c5", "aba841fb842d855255dc58eebe011f87", "123456789", "200300400")
	fmt.Printf("token: %s\n", token)
}
