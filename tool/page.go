package tool

import (
	"fmt"
	"math"
)

//完成分页功能的需求

// 先定义装数据的结构
type Page struct {
	TotalCont  int
	TotalPages int
	Current    int
	Prev       int
	Next       int
	Pages      []int
}

// 定义一个全局变量，记录每页显示的数量
var pageCount = 5

// 增加一个包方法，允许修改每页显示的数据
func SetPageCount(pc int) {
	pageCount = pc
}

func GetPageCount() int {
	return pageCount
}

func GetPage(totalCount int, current int) Page {
	//得到页码的对象
	page := Page{
		TotalCont:  totalCount,
		TotalPages: int(math.Ceil(float64(totalCount) / float64(pageCount))),
		Current:    current,
		Prev:       current - 1,
		Next:       current + 1,
	}

	//动态生成所有页码
	page = *getSinglePages(&page)
	fmt.Println(page)

	//计算上一页和下一页
	var prev int
	var next int

	if current <= 1 {
		prev = 1
	} else {
		prev = current - 1
	}

	if next >= page.TotalPages {
		next = page.TotalPages
	} else {
		next = current + 1
	}

	// 将上一页和下一页放到目标对象
	page.Prev = prev
	page.Next = next

	return page
}

func getSinglePages(page *Page) *Page {
	//定义起始、结束位置
	var start int = 1
	var end int = page.TotalPages

	//逻辑处理，解决起始位置的问题
	//如果当前页面大于7，起始位置=当前页-5
	if page.Current >= 6 {
		start = page.Current - 5

		//当前页面落在后五页，保证有十页
		if page.Current+4 >= page.TotalPages {
			start = page.TotalPages - 4 - 5
		}

		//前提是end不超过总页数
		if page.Current+4 <= page.TotalPages {
			end = page.Current + 4
		}
	}

	//当现实页码小于6页事故。end保留10页逻辑
	if page.Current <= 6 {
		end = 10
	}

	//3.安全处理，如果总页数达不到10时
	if page.TotalPages <= 10 {
		start = 1
		end = page.TotalPages
	}

	//4.安全处理，用户手动输入超过的页数
	if page.Current <= 1 {
		//少于第一页
		page.Current = 1
	}

	if page.Current > page.TotalCont {
		//超过最后一页
		page.Current = page.TotalPages
	}

	for i := start; i <= end; i++ {
		page.Pages = append(page.Pages, i)
	}
	return page
}
