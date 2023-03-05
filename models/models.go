package models

// 泛型 interface
type YamlParse interface {
	*Config
}

type PostJson interface {
	*PostUser | *PostDevice | *PostModDevInfo | *PostAppList | *DbVersionInfoTable | *DbAppListTable
}
