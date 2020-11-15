package model

// Reason 原因模型
type GoodsModel struct {
	Id    int
	Stock int
	Price int
}

func (GoodsModel) TableName() string{
	return "goods"
}

