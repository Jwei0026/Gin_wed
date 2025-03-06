package brands

import "2024-2025/database"

type Brands struct {
	Brand_id     int    `gorm:""`
	English_name string `gorm:"size:100"`
	Chinese_name string `gorm:"size:100"`
	Grade        string `gorm:"size:2"`
}

func GetBrands() []Brands {
	//定义变量
	var brands []Brands
	//执行SQL
	database.Gdb.Raw("select brand_id,english_name,chinese_name,grade from brands;").Find(&brands)
	//返回数据
	return brands
}
