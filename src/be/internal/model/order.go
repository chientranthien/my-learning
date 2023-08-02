package model

type Order struct {
	Id int64
	Region string
	UserId int64
	Status int32
	ExtraInfo []byte
	Ctime int64
	Mtime int64
}
func (Order) TableName() string {
    return "order_tab"
}