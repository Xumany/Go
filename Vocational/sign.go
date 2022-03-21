package vocational

import (
	"Gozhijiao/request"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

var (
	header   = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	UserInfo UserInfo
)

type info struct {
	UserInfo UserInfo
	today    today
}

type today struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DataList []struct {
		ID           string `json:"Id"`
		CourseOpenID string `json:"courseOpenId"`
		CourseName   string `json:"courseName"`
		ClassName    string `json:"className"`
		Title        string `json:"Title"`
		OpenClassID  string `json:"openClassId"`
		DateCreated  string `json:"dateCreated"`
		TeachDate    string `json:"teachDate"`
		ClassSection string `json:"classSection"`
		Address      string `json:"Address"`
		State        int    `json:"state"`
	} `json:"dataList"`
}
type UserInfo struct {
	Code                int    `json:"code"`
	UserType            int    `json:"userType"`
	Token               string `json:"token"`
	UserName            string `json:"userName"`
	SecondUserName      string `json:"secondUserName"`
	UserID              string `json:"userId"`
	NewToken            string `json:"newToken"`
	DisplayName         string `json:"displayName"`
	EmployeeNumber      string `json:"employeeNumber"`
	URL                 string `json:"url"`
	SchoolName          string `json:"schoolName"`
	SchoolID            string `json:"schoolId"`
	IsValid             int    `json:"isValid"`
	IsNeedMergeUserName int    `json:"isNeedMergeUserName"`
	IsZjyUser           int    `json:"isZjyUser"`
	IsGameUser          int    `json:"isGameUser"`
	IsNeedUpdatePwd     int    `json:"isNeedUpdatePwd"`
	PwdMsg              string `json:"pwdMsg"`
}

type Classdetail struct {
	UserID   string
	NewToken string
	Code     int        `json:"code"`
	Datalist []DataList `json:"dataList"`
}
type ActivityList struct {
	DataList []DataList `json:"dataList"`
}
type DataList struct {
	ID           string `json:"Id"`
	CourseOpenID string `json:"courseOpenId"`
	OpenClassID  string `json:"openClassId"`
	State        int    `json:"state"`
	Gesture      string `json:"Gesture"`
	DataType     string `json:"DataType"`
}

// 获取指定日期课堂
func Getdate(c UserInfo, date string) Classdetail {
	var (
		url    = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data   = map[string]string{"stuId": c.UserID, "faceDate": date, "newToken": c.NewToken}
		result *Classdetail
	)
	_, err := req.R().SetFormData(data).SetHeaders(header).SetResult(result).Post(url)
	if err != nil {
		panic(err)
	}

	if result.Code == 1 {
		panic(errors.New("获取失败"))
	}

	result.UserID = c.UserID
	result.NewToken = c.NewToken
	return *result
}

// 获取课堂
func NewGetStuFaceActivityList(c Classdetail) {
	var (
		url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/newGetStuFaceActivityList"
	)
	// $data=array("activityId"=>$i['Id'],"stuId"=>$stuId,"classState"=>$i['state'],"openClassId"=>$i['openClassId'],"newToken"=>$newtoken);
	for _, v := range c.Datalist {

		data := map[string]string{
			"activityId":  v.ID,
			"stuId":       c.UserID,
			"classState":  "2",
			"openClassId": v.OpenClassID,
			"newToken":    c.NewToken,
		}
		d := request.Post(url, data, header)
		var tmp = ActivityList{}
		json.Unmarshal(d, &tmp)
		tmp.IsJoinActivities(c.UserID, c.NewToken, v.ID, v.OpenClassID)
	}
}

func (a ActivityList) IsJoinActivities(userid, newToken, kcid, OpenClassID string) {
	var (
		url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/isJoinActivities"
	)
	//fmt.Println(a)
	for _, n := range a.DataList {
		if n.DataType == "签到" && n.State != 3 {
			fmt.Println(n.ID, "open:"+n.OpenClassID, userid, newToken, kcid)
			// $attendData = array("activityId"=>$i['Id'],"openClassId"=>$i['openClassId'],"stuId"=>$stuId,"typeId"=>$n['Id'],"type"=>"1","newToken"=>$newtoken);
			data := map[string]string{
				"activityId":  kcid,
				"newToken":    newToken,
				"stuId":       userid,
				"openClassId": OpenClassID,
				"typeId":      n.ID,
				"type":        "1",
			}
			cs := request.Post(url, data, header)
			fmt.Println(string(cs))
			tamp := struct {
				Code     int    `json:"code"`
				Msg      string `json:"msg"`
				IsAttend int    `json:"isAttend"`
			}{}
			json.Unmarshal(cs, &tamp)
			// 判断是否结束
			// 3 结束
			// 1 开启中
			if tamp.IsAttend != 1 {
				// 如果这个是开启的就进行签到一次
				data = map[string]string{
					"signId":      n.ID,
					"stuId":       userid,
					"openClassId": OpenClassID,
					"sourceType":  "3",
					"checkInCode": n.Gesture,
					"activityId":  kcid,
					"newToken":    newToken,
				}
				s := request.Post("https://zjyapp.icve.com.cn/newmobileapi/faceteach/saveStuSignNew", data, header)
				// 这里发送一个信息通知
				fmt.Println(s)
				request.Notice("签到成功", time.Now().Format("2006-01-02")+"课堂签到成功")

			}
		}
	}
}

// 获取今日课堂
func GetToday(c UserInfo) Classdetail {

	req.DevMode()
	var (
		url    string = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data          = map[string]string{"stuId": c.UserID, "faceDate": time.Now().Format("2006-01-02"), "newToken": c.NewToken}
		result        = Classdetail{}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)

	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		panic(errors.New("http请求失败"))

	}
	resp.Unmarshal(&result)
	if result.Code != 1 {
		panic(errors.New("获取失败"))
	}

	result.UserID = c.UserID
	result.NewToken = c.NewToken

	return result
}

// 获取所有课堂
func GetAll(c UserInfo) Classdetail {
	var (
		url    string = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data          = map[string]string{"stuId": c.UserID, "newToken": c.NewToken}
		result        = Classdetail{}
	)
	s := request.Post(url, data, header)

	json.Unmarshal(s, &result)
	if result.Code != 1 {
		panic(errors.New("获取今日课堂失败"))
	}
	var ret = Classdetail{}
	json.Unmarshal(s, &ret)
	ret.UserID = c.UserID
	ret.NewToken = c.NewToken
	return ret
}

func RepairSign() {
	//var url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/changeSignType"
	//	c := struct {
	// 	OpenClassId    string
	// 	ID             string
	// 	SignId         string
	// 	StuId          string
	// 	SignResultType string
	// 	SourceType     int
	// 	schoolId       string
	// }{}
	//request.PostJson(url, c, "application/json")
}
