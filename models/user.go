package models

type User struct {
	Name     string `from:"name"`
	Birthday string `form:"birthday"`
	Address  string `form:"address"`
	Nick     string `form:"nick"`
	Password string `form:"password"`
}
