package products

import (
	"2024-2025/database"
	"fmt"
)

//所有与products表相关的操作，都在这封装
/*
	1. 定义表结构（应该跟需要的数据有关）
	2.封装相关的操作函数
*/

type IndexProducts struct {
	//首页需要显示的商品数据：商品ID、品牌名字、系列名字、商品名字、商品价格、商品图片、商品类别、是否特价
	Product_id int `gorm:"primarykey"`

	Style            string  `gorm:"size:10"`
	Price            float64 `gorm:""`
	Is_special_offer int     `gorm:""`
	Special_price    float64 `gorm:""`
	Image            string  `gorm:"size:100"`
	Gender           string  `gorm:"size:10"`
	Ordered_num      int     `gorm:""`

	//品牌名称
	Chinese_name string `gorm:"size:100"`
	English_name string `gorm:"size:100"`
	Grade        string `gorm:"size:2"`

	//系列信息
	Series_name string `gorm:"cloumn:series_name;size:100"`
}

type SpecialProducts struct {
	//首页需要显示的商品数据：商品ID、品牌名字、系列名字、商品名字、商品价格、商品图片、商品类别、是否特价
	Product_id int `gorm:"primarykey"`

	Style            string  `gorm:"size:10"`
	Price            float64 `gorm:""`
	Is_special_offer int     `gorm:""`
	Special_price    float64 `gorm:""`
	Image            string  `gorm:"size:100"`
	Ordered_num      int     `gorm:""`
	Gender           string  `gorm:"size:10"`

	//品牌名称
	Chinese_name string `gorm:"size:100"`
	English_name string `gorm:"size:100"`

	//系列信息
	Series_name string `gorm:"cloumn:series_name;size:100"`
}

type ProductDetail struct {
	//商品详细信息
	Product_id       int     `gorm:"primary"`
	Series_id        int     `gorm:""`
	Case_material    string  `gorm:"size:10"`
	Case_back        string  `gorm:"size:10"`
	Strap_color      string  `gorm:"size:10"`
	Functions        string  `gorm:"size:100"`
	Style            string  `gorm:"size:10"`
	Size             string  `gorm:"size:20"`
	Watch_glass      string  `gorm:"size:20"`
	Watch_buckle     string  `gorm:"size:10"`
	Launch_year      string  `gorm:"size:8"`
	Movement         string  `gorm:"size:10"`
	Thickness        float64 `gorm:""`
	Dial             string  `gorm:"size:10"`
	Water_resistance string  `gorm:"size:50"`
	Price            float64 `gorm:""`
	Is_new           bool    `gorm:""`
	Is_special_offer bool    `gorm:""`
	Special_price    float64 `gorm:""`
	Image            string  `gorm:"size:100"`
	Gender           string  `gorm:"size:10"`
	Ordered_num      int     `gorm:""`

	//商品品牌信息
	English_name string `gorm:"size:100"`
	Chinese_name string `gorm:"size:100"`
	Grade        string `gorm:"size:2"`
	Brand_id     int    `gorm:""`

	//商品系列名字
	Series_name string `gorm:"size:100"`
}

// 获取首页所需数据（连表查询）
func GetIndexProudcts() []IndexProducts {
	//定义切片保存结果
	var products []IndexProducts

	database.Gdb.Raw("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.ordered_num,p.style,s.series_name,b.chinese_name,b.english_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id;").Scan(&products)

	// println(products)
	return products
}

// 获取所有特价商品
func GetSpecialProducts() []SpecialProducts {
	var products []SpecialProducts

	database.Gdb.Raw("select p.product_id,p.style,p.price,p.is_special_offer,p.special_price,p.image,p.ordered_num,p.gender,b.chinese_name,b.english_name,s.series_name from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id;").Scan(&products)

	return products
}

// 根据性别获取商品
func GetGenderProducts(gender string, page int, pagecount int) []IndexProducts {
	//将条件组装，至sql执行
	var genderProducts []IndexProducts

	//分页显示数据的逻辑。limit起始位置，限制数量
	start := (page - 1) * pagecount
	database.Gdb.Raw(fmt.Sprintf("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.gender ='%s' order by p.price desc limit %d,%d", gender, start, pagecount)).Scan(&genderProducts)
	return genderProducts
}

// 获取商品数据，限定四条
func GetNewProducts() []IndexProducts {
	var newProducts []IndexProducts
	database.Gdb.Raw("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.is_new = 1 limit 3;").Scan(&newProducts)
	return newProducts
}

// 获取商品详情数据
func GetDetail(id string) ProductDetail {
	var product ProductDetail
	database.Gdb.Raw("select p.*,s.series_name,b.chinese_name,b.english_name,b.grade,b.brand_id from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.product_id =" + id).Scan(&product)
	// fmt.Println(product)
	return product
}

// 获取商品详情数据的某一属性
func GetDetailone(id string, attribute string) string {
	var attributemsg string
	database.Gdb.Raw(fmt.Sprintf("select %s from products  where product_id ='%s'", attribute, id)).Scan(&attributemsg)
	fmt.Println(attributemsg)
	return attributemsg
}

// 获取品牌热销商品数据
func GetBrandHostProducts(brand_id int) []IndexProducts {
	var products []IndexProducts
	fmt.Println(brand_id)
	database.Gdb.Raw(fmt.Sprintf("select p.product_id,p.style,p.price,p.is_special_offer,p.special_price,p.image,p.ordered_num,s.series_name,b.chinese_name,b.english_name from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where b.brand_id = %d order by p.ordered_num desc limit 4", brand_id)).Scan(&products)
	// fmt.Println(products)
	return products
}

// type count struct {
// 	count int `gorm:""`
// }

// 获取总记录数
func GetGenderCounts(gender string) int {
	var count int //记录总记录数

	database.Gdb.Raw("select count(*) count from products where gender = '" + gender + "'").First(&count)
	fmt.Println(count)

	return count
}

// 获取商品数据，限定五条
func GetNewProductsfive() []IndexProducts {
	var newProducts []IndexProducts
	database.Gdb.Raw("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.is_new = 1 limit 4;").Scan(&newProducts)
	return newProducts
}

// 根据性别获取数据，限定前五条
func GetGenderProductsfive(gender string, page int, pagecount int) []IndexProducts {
	//将条件组装，至sql执行
	var genderProducts []IndexProducts

	//分页显示数据的逻辑。limit起始位置，限制数量
	start := (page - 1) * pagecount
	database.Gdb.Raw(fmt.Sprintf("select p.product_id,p.image,p.price,p.is_special_offer,p.special_price,p.gender,p.style,p.ordered_num,s.series_name,b.chinese_name,b.grade from products p left join series s on p.series_id = s.series_id left join brands b on s.brand_id = b.brand_id where p.gender ='%s' order by p.price desc limit %d,%d", gender, start, pagecount)).Scan(&genderProducts)
	return genderProducts
}
