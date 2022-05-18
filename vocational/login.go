package vocational

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

//  现在访问过多会曝出 504 的错误

// Login 登录
func Login(u, p string) *UserInfo {
	var (
		emit             = fmt.Sprint(time.Now().Unix(), "000")
		device           = "Xiaomi Redmi K20 Pro"
		deviceApiVersion = "10"
		appVersion       = getAppVersion()
		Userinfo         UserInfo
		url              = "https://zjyapp.icve.com.cn/newMobileAPI/MobileLogin/newSignIn"
		data             = map[string]string{"clientId": "d902c875d5f34c0f93362139f5af0c4c", "sourceType": "2", "userPwd": p, "userName": u, "appVersion": appVersion, "equipmentAppVersion": appVersion, "equipmentApiVersion": deviceApiVersion, "equipmentModel": device}
	)
	header["emit"] = emit
	header["device"] = getDeviceEncryption(device, deviceApiVersion, appVersion, emit)
	resp, err := req.R().SetHeaders(header).SetFormData(data).Post(url)
	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		panic(resp.Error())
	}
	err = resp.Unmarshal(&Userinfo)
	if err != nil {
		panic(err)
	}
	if Userinfo.Code != 1 {
		panic(Userinfo.Msg)
	}
	return &Userinfo
}
func (i *UserInfo) NewGetStuFaceActivityList() {
	url := "https://zjyapp.icve.com.cn/newmobileapi/faceteach/newGetStuFaceActivityList"
	var c Classroom
	data := map[string]string{"stuId": i.UserID, "newToken": i.NewToken, "classState": "2"}
	for _, v := range i.DataList {
		data["activityId"], data["openClassId"] = v.ID, v.OpenClassID
		resp, err := req.SetHeaders(header).SetFormData(data).Post(url)
		if err != nil {
			panic(err)
		}
		if !resp.IsSuccess() {
			panic(resp.Error())
		}

		err = resp.Unmarshal(&c)
		if err != nil {
			panic(err)
		}
		if c.Code != 1 {
			panic(c.Msg)
		}
		for _, n := range c.DataList {
			if n.DataType == "签到" && n.State != 3 {
				n.KID, n.OpenClassID = v.ID, v.OpenClassID
				i.SingIn = append(i.SingIn, n)
			}
		}
	}
}
func (i *UserInfo) IsJoinActivities() {
	var url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/IsJoinActivities"
	var data = map[string]string{"newToken": i.NewToken, "stuId": i.UserID, "typeId": "1"}
	var msg Msg
	for _, v := range i.SingIn {
		data["activityId"], data["openClassId"], data["typeId"] = v.KID, v.OpenClassID, v.ID
		resp, err := req.SetHeaders(header).SetFormData(data).Post(url)
		if err != nil {
			panic(err)
		}
		if !resp.IsSuccess() {
			panic(resp.Error())
		}
		err = resp.Unmarshal(&msg)
		if err != nil {
			panic(err)
		}
		if msg.IsAttend != 1 {
			medata := map[string]string{"signId": v.ID, "stuId": i.UserID, "openClassId": v.OpenClassID, "sourceType": "3", "checkInCode": v.Gesture, "activityId": v.KID, "newToken": i.NewToken}
			resp, err = req.SetHeaders(header).SetFormData(medata).Post("https://zjyapp.icve.com.cn/newmobileapi/faceteach/saveStuSignNew")
			if err != nil {
				panic(err)
			}
			if !resp.IsSuccess() {
				panic(resp.Error())
			}
			m := res{}
			_ = resp.Unmarshal(&m)
			if m.Msg == "签到成功！" {
				fmt.Println("签到成功")
				_, err = req.Get(fmt.Sprintf("https://sctapi.ftqq.com/%s.send?title=%s&desp=%s", sendKey, "签到成功", time.Now().Format("2006-01-02 15:04-05")+"签到成功"))
				if err != nil {
					panic(err)
				}
			}
		}

	}
}
func (i *UserInfo) GetToday() {
	var (
		url  = "https://zjyapp.icve.com.cn/newMobileAPI/FaceTeach/getStuFaceTeachList"
		data = map[string]string{"stuId": i.UserID, "faceDate": time.Now().Format("2006-01-02"), "newToken": i.NewToken}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)
	if err != nil {
		panic(err)
	}
	if resp.IsSuccess() {
		err = resp.Unmarshal(&i)
		if err != nil {
			return
		}
		if i.Code != 1 {
			panic(i.Msg)
		}

	}
	return
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

// GetDate 获取指定日期课程
func (i *UserInfo) GetDate(date string) {

	var (
		url  = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data = map[string]string{"stuId": i.UserID, "faceDate": date, "newToken": i.NewToken}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)
	if err != nil {
		panic(err)
	}

	if !resp.IsSuccess() {
		panic(resp.Error())
	}
	err = resp.Unmarshal(&i)
	if err != nil {
		panic(err)
	}
	if i.Code != 1 {
		panic(i.Msg)
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
