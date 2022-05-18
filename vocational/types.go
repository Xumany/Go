package vocational

var (
	header  = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	sendKey = "SCT132670TKeOPbiky28mDLzKZyfniZ9uh"
)

type Msg struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	IsAttend int    `json:"isAttend"`
}
type res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type UserInfo struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	UserType    int    `json:"userType"`
	UserName    string `json:"userName"`
	UserID      string `json:"userId"`
	NewToken    string `json:"newToken"`
	DisplayName string `json:"displayName"`
	SchoolName  string `json:"schoolName"`
	SchoolID    string `json:"schoolId"`
	DataList    []struct {
		ID           string `json:"Id"`
		CourseOpenID string `json:"courseOpenId"`
		OpenClassID  string `json:"openClassId"`
		State        int    `json:"state"`
	} `json:"dataList"`
	SingIn []ClassRoomInfo
}

type Classroom struct {
	Code     int             `json:"code"`
	Msg      string          `json:"msg"`
	DataList []ClassRoomInfo `json:"dataList"`
}
type ClassRoomInfo struct {
	ID          string `json:"Id"`
	KID         string
	DataType    string `json:"DataType"`
	SignType    int    `json:"SignType"`
	Gesture     string `json:"Gesture"`
	OpenClassID string
	State       int `json:"State"`
}
type Config struct {
	User    string
	Pass    string
	SendKey string
}
