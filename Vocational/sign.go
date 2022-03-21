package vocational

import (
	"Gozhijiao/request"
	"encoding/json"
	"errors"
	"time"

	"github.com/imroc/req/v3"
)

var (
	header   = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	UserInfo info
)

type User interface {
	Login(u, p string) *info
}

type info struct {
	UserInfo  LoginInfo
	today     today
	Classroom Classroom
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
type Classroom struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	IsEvaluation   int    `json:"isEvaluation"`
	FaceEvaluation int    `json:"faceEvaluation"`
	DataList       []struct {
		ID                string      `json:"Id"`
		Title             string      `json:"Title"`
		DateCreated       string      `json:"DateCreated"`
		CreatorID         string      `json:"CreatorId"`
		DataType          string      `json:"DataType"`
		State             int         `json:"State"`
		SignType          int         `json:"SignType"`
		Gesture           string      `json:"Gesture"`
		AskType           int         `json:"AskType"`
		ViewAnswer        int         `json:"ViewAnswer"`
		ResourceURL       interface{} `json:"resourceUrl"`
		CellType          int         `json:"cellType"`
		CategoryName      interface{} `json:"categoryName"`
		ModuleID          interface{} `json:"moduleId"`
		CellSort          int         `json:"cellSort"`
		HkOrExamType      int         `json:"hkOrExamType"`
		PaperType         int         `json:"paperType"`
		TermTimeID        interface{} `json:"termTimeId"`
		IsForbid          int         `json:"isForbid"`
		FixedPublishTime  interface{} `json:"fixedPublishTime"`
		ExamStuID         interface{} `json:"examStuId"`
		ExamWays          int         `json:"examWays"`
		IsAuthenticate    int         `json:"isAuthenticate"`
		IsAnswerOrPreview int         `json:"isAnswerOrPreview"`
		IsPreview         int         `json:"isPreview"`
		StuStartDate      interface{} `json:"StuStartDate"`
		StuEndDate        interface{} `json:"StuEndDate"`
	} `json:"dataList"`
}

//登录信息结构体
type LoginInfo struct {
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
// func Getdate(c UserInfo, date string) Classdetail {
// 	var (
// 		url    = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
// 		data   = map[string]string{"stuId": c.UserID, "faceDate": date, "newToken": c.NewToken}
// 		result *Classdetail
// 	)
// 	_, err := req.R().SetFormData(data).SetHeaders(header).SetResult(result).Post(url)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if result.Code == 1 {
// 		panic(errors.New("获取失败"))
// 	}

// 	result.UserID = c.UserID
// 	result.NewToken = c.NewToken
// 	return *result
// }

// 获取课堂
func (c *info) NewGetStuFaceActivityList() {
	var (
		url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/newGetStuFaceActivityList"
	)
	// $data=array("activityId"=>$i['Id'],"stuId"=>$stuId,"classState"=>$i['state'],"openClassId"=>$i['openClassId'],"newToken"=>$newtoken);
	for _, v := range c.today.DataList {

		data := map[string]string{
			"activityId":  v.ID,
			"stuId":       c.UserInfo.UserID,
			"classState":  "2",
			"openClassId": v.OpenClassID,
			"newToken":    c.UserInfo.NewToken,
		}
		resp, err := req.SetHeaders(header).SetFormData(data).Post(url)
		if err != nil {
			panic(err)
		}
		if !resp.IsSuccess() {
			panic("访问失败")
		}
		resp.Unmarshal(&UserInfo.Classroom)
		c.IsJoinActivities(v.ID, v.OpenClassID)
	}
}

func (c *info) IsJoinActivities(kcid, OpenClassID string) {
	var (
		url = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/isJoinActivities"
	)
	//fmt.Println(a)
	for _, n := range c.Classroom.DataList {
		if n.DataType == "签到" && n.State != 3 {
			// fmt.Println(n.ID, "open:"+n.OpenClassID, userid, newToken, kcid)
			// $attendData = array("activityId"=>$i['Id'],"openClassId"=>$i['openClassId'],"stuId"=>$stuId,"typeId"=>$n['Id'],"type"=>"1","newToken"=>$newtoken);
			data := map[string]string{
				"activityId":  kcid,
				"newToken":    c.UserInfo.NewToken,
				"stuId":       c.UserInfo.UserID,
				"openClassId": OpenClassID,
				"typeId":      n.ID,
				"type":        "1",
			}
			cs := request.Post(url, data, header)
			// fmt.Println(string(cs))
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
					"stuId":       c.UserInfo.UserID,
					"openClassId": OpenClassID,
					"sourceType":  "3",
					"checkInCode": n.Gesture,
					"activityId":  kcid,
					"newToken":    c.UserInfo.NewToken,
				}
				_ = request.Post("https://zjyapp.icve.com.cn/newmobileapi/faceteach/saveStuSignNew", data, header)
				// 这里发送一个信息通知
				// fmt.Println(s)
				request.Notice("签到成功", time.Now().Format("2006-01-02")+"课堂签到成功")

			}
		}
	}
}

// 获取今日课堂
func (c *info) GetToday() {
	var (
		url  string = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
		data        = map[string]string{"stuId": c.UserInfo.UserID, "faceDate": time.Now().Format("2006-01-02"), "newToken": c.UserInfo.NewToken}
	)
	resp, err := req.R().SetFormData(data).SetHeaders(header).Post(url)

	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		panic(errors.New("http请求失败"))

	}
	resp.Unmarshal(&UserInfo.today)
	c.NewGetStuFaceActivityList()
}

// 获取所有课堂
// func GetAll(c *info) Classdetail {
// 	var (
// 		url    string = "https://zjyapp.icve.com.cn/newmobileapi/faceteach/getStuFaceTeachList"
// 		data          = map[string]string{"stuId": c.UserInfo.UserID, "newToken": c.UserInfo.NewToken}
// 		result        = Classdetail{}
// 	)
// 	s := request.Post(url, data, header)

// 	json.Unmarshal(s, &result)
// 	if result.Code != 1 {
// 		panic(errors.New("获取今日课堂失败"))
// 	}
// 	var ret = Classdetail{}
// 	json.Unmarshal(s, &ret)

// 	return ret
// }

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
