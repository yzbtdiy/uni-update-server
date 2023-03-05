package models

// post 请求数据结构体
type PostUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostDevice struct {
	// EditionType   string `json:"edition_type"`
	// VersionType   string `json:"version_type"`
	// EditionNumber int `json:"edition_number"`
	DeviceId   string `json:"device_id"`
	DeviceMode string `json:"device_mode"`
}

type PostModDevInfo struct {
	DeviceId string `json:"device_id"`
	Activate int    `json:"activate"`
}

type PostAppList struct {
	DeviceId string `json:"device_id"`
}

// 响应数据结构体
type ResLogin struct {
	Token string `json:"token"`
}

type ResData struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
