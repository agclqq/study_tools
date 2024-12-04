package demo

type EntityA struct {
	Id     int    `json:"id" gorm:"primary_key;auto_increment`
	Name   string `json:"name omitempty"`
	Status int    `json:"status"`
}
