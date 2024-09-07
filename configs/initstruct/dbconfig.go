package initstruct

// DataBaseConfig json中的嵌套对应结构体的嵌套
type DataBaseConfig struct {
	Name         string
	Password     string
	Host         string
	Port         string
	User         string
	MaxIdleConns int
	MaxOpenConns int
}
