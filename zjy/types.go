package zjy

var (
	header  = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	sendKey string
)

// 这个可以和下面的合并
type Msg struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	IsAttend int    `json:"isAttend"`
}
type res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//  上面两个结构体 下次开发的时候可以合并一下
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
	KID         string // 这个就是留着
	DataType    string `json:"DataType"`
	SignType    int    `json:"SignType"`
	Gesture     string `json:"Gesture"`
	OpenClassID string
	State       int `json:"State"`
}

// 配置文件
// User  登陆账号
// Pass  登陆密码
// 通知的Key
type config struct {
	User, Pass, SendKey string
}

func NewConfig(user, pass, key string) *config {
	return &config{
		User:    user,
		Pass:    pass,
		SendKey: key,
	}
}

type Study interface {
}

// zyj.run(config)
