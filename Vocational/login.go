package vocational

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

// 登录
func Login(u, p string) *info {
	var (
		emit             string = fmt.Sprint(time.Now().Unix(), "000")
		device           string = "Xiaomi Redmi K20 Pro"
		deviceApiVerison string = "10"
		appVersion       string = getAppVeriosn()
		url                     = "https://zjyapp.icve.com.cn/newMobileAPI/MobileLogin/newSignIn"
		data                    = map[string]string{"clientId": "d902c875d5f34c0f93362139f5af0c4c", "sourceType": "2", "userPwd": p, "userName": u, "appVersion": appVersion, "equipmentAppVersion": appVersion, "equipmentApiVersion": deviceApiVerison, "equipmentModel": device}
	)
	header["emit"] = emit
	header["device"] = getDeviceEncryption(device, deviceApiVerison, appVersion, emit)

	// 这个函数里面要判断他是否登录成功
	s, err := req.SetHeaders(header).SetFormData(data).SetResult(&UserInfo.UserInfoda).Post(url)
	if err != nil {
		panic(err)
	}
	return s
}

// 获取最新的APP版本
func getAppVeriosn() string {
	result := &struct {
		AppVersionInfo struct {
			VersionCode string `json:"VersionCode"`
		} `json:"appVersionInfo"`
	}{}
	resp, err := req.Get("https://zjy2.icve.com.cn/portal/AppVersion/getLatestVersionInfo")
	if err != nil {
		panic(err)
	}
	if resp.IsSuccess() {
		err = resp.Unmarshal(result)
		if err != nil {
			panic(err)
		}
	}

	return result.AppVersionInfo.VersionCode

}

// 计算设备信息md5
func getDeviceEncryption(args ...string) string {
	// var md5str1 string
	var tmp string
	for i, v := range args {
		if i == 0 {
			tmp = md5str(v)
			continue
		}
		tmp = tmp + args[i]
		tmp = md5str(tmp)
	}
	return tmp
}

// 封装MD5函数
func md5str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}
