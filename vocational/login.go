package vocational

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

// 登录
func Login(u, p string) *info {
	var (
		emit             = fmt.Sprint(time.Now().Unix(), "000")
		device           = "Xiaomi Redmi K20 Pro"
		deviceApiVersion = "10"
		appVersion       = getAppVersion()
		Userinfo         info
		url              = "https://zjyapp.icve.com.cn/newMobileAPI/MobileLogin/newSignIn"
		data             = map[string]string{"clientId": "d902c875d5f34c0f93362139f5af0c4c", "sourceType": "2", "userPwd": p, "userName": u, "appVersion": appVersion, "equipmentAppVersion": appVersion, "equipmentApiVersion": deviceApiVersion, "equipmentModel": device}
	)
	header["emit"] = emit
	header["device"] = getDeviceEncryption(device, deviceApiVersion, appVersion, emit)

	// 这个函数里面要判断他是否登录成功
	resp, err := req.R().SetHeaders(header).SetFormData(data).Post(url)
	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		panic("访问错误")
	}
	err = resp.Unmarshal(&Userinfo.UserInfo)
	if err != nil {
		panic(err)
	}
	return &Userinfo
}

func (i *info) NewGetStuFaceActivityList() {
	url := "https://zjyapp.icve.com.cn/newmobileapi/faceteach/newGetStuFaceActivityList"
	data := map[string]string{"stuId": i.UserInfo.UserID, "newToken": i.UserInfo.NewToken, "classState": "2"}
	for _, v := range i.Today.DataList {
		data["activityId"], data["openClassId"] = v.ID, v.OpenClassID
		resp, err := req.SetHeaders(header).SetFormData(data).Post(url)
		if err != nil {
			panic(err)
		}
		if !resp.IsSuccess() {
			panic("访问失败")
		}
	}
}
func (i *info) IsJoinActivities(kid, OpenClassID string) {
	_, _ = kid, OpenClassID
}
func (i *info) GetToday() {
	var (
		url  = "https://zjyapp.icve.com.cn/newMobileAPI/FaceTeach/getStuFaceTeachList"
		data = map[string]string{"stuId": i.UserInfo.UserID, "faceDate": time.Now().Format("2006-01-02"), "newToken": i.UserInfo.NewToken}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)

	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		panic(errors.New("http请求失败"))

	}
	err = resp.Unmarshal(&i.Today)
	if err != nil {
		panic(err)
	}
}

// 获取最新的APP版本
func getAppVersion() string {
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

// 获取指定日期课程
func (i *info) GetDate(date string) {
	fmt.Sprintln(i)
	var (
		url  = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data = map[string]string{"stuId": i.UserInfo.UserID, "faceDate": date, "newToken": i.UserInfo.NewToken}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)
	if err != nil {
		panic(err)
	}

	if !resp.IsSuccess() {
		panic(errors.New("获取失败"))
	}
	err = resp.Unmarshal(&i.Today)
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
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
