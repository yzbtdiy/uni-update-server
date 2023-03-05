package models

// 数据库结构体
type DbDeviceInfoTable struct {
	Id         int    `gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
	DeviceId   string `gorm:"column:device_id;type:TEXT NOT NULL UNIQUE;"`
	DeviceMode string `gorm:"column:device_mode;type:TEXT NOT NULL;"`
	Activate   int    `gorm:"column:activate;type:INTEGER NOT NULL DEFAULT 0"`
	// EditionNumber int    `gorm:"column:edition_number;type:INTEGER NOT NULL;"`
}

func (DbDeviceInfoTable) TableName() string {
	return "device_info"
}

type DbVersionInfoTable struct {
	Id int `json:"id" gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
	*VersionInfo
}

type VersionInfo struct {
	Describe       string `json:"describe" gorm:"column:describe;type:TEXT NOT NULL;"`
	EditionUrl     string `json:"edition_url" gorm:"column:edition_url;type:TEXT NOT NULL;"`
	EditionForce   int    `json:"edition_force" gorm:"column:edition_force;type:INTEGER NOT NULL;"`
	PackageType    int    `json:"package_type" gorm:"column:package_type;type:INTEGER NOT NULL;"`
	EditionIssue   int    `json:"edition_issue" gorm:"column:edition_issue;type:INTEGER NOT NULL;"`
	EditionNumber  int    `json:"edition_number" gorm:"column:edition_number;type:INTEGER NOT NULL;"`
	EditionName    string `json:"edition_name" gorm:"column:edition_name;type:TEXT NOT NULL;"`
	EditionSilence int    `json:"edition_silence" gorm:"column:edition_silence;type:INTEGER NOT NULL;"`
}

func (DbVersionInfoTable) TableName() string {
	return "version_info"
}

type DbUserTable struct {
	Id       int    `gorm:"column:id;type:INTEGER NOT NULL"`
	Username string `gorm:"column:username;type:TEXT NOT NULL UNIQUE;"`
	Password string `gorm:"column:password;type:TEXT NOT NULL;"`
}

func (DbUserTable) TableName() string {
	return "admin_user"
}

type DbAppListTable struct {
	Id int `json:"id" gorm:"column:id;type:INTEGER NOT NULL;primaryKey;autoIncrement;"`
	*AppInfo
}

type AppInfo struct {
	Name string `json:"name" gorm:"column:name;type:TEXT NOT NULL"`
	Url  string `json:"url" gorm:"column:url;type:TEXT NOT NULL"`
}

func (DbAppListTable) TableName() string {
	return "app_list"
}
