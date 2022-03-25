package vocational

var (
	header = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
)

type info struct {
	UserInfo LoginInfo
	Today    Today
}
type Today struct {
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
