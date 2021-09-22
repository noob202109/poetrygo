package model

import (
	"PoetryWeb/data"
	"PoetryWeb/utils/utils"
	"fmt"
	"strconv"
	"strings"
)

//Getnavbar 获取导航栏
func Getnavbar() []data.Navbar {
	//sql语句
	sqlStr := "select DISTINCT epoch from poetry"
	//声明model 下的结构体
	var navbar []data.Navbar
	//执行sql
	err := utils.Db.Select(&navbar, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}

	return navbar
}

//Getpageindexs 获取带分页的主页诗词信息
func Getpageindexs(page string) (data.Page, []data.Poetry) {
	//将传过来的页码转换为int
	ipage, _ := strconv.ParseInt(page, 10, 64)
	//获取总记录数
	//sql语句
	sqlStr := "SELECT COUNT(*) FROM poetry"
	//总记录数 定义一个变量保存总记录数
	var Num int64
	//执行sql
	utils.Db.Get(&Num, sqlStr)
	//设置每页显示的记录数
	var pagesize int64 = 10
	//总页数 定义一个变量保存总页数
	var count int64
	//计算总页数
	if Num%pagesize == 0 {
		//如果总页数除每页条数可以除净 那直接除就可以了
		count = Num / pagesize
	} else {
		//如果总页数除每页条数可以除不净 那就需要在加一页
		count = Num/pagesize + 1
	}

	//判断是否有上一页
	IsPrev := ipage > 1
	//计算上一页页码
	var Prev int64
	if IsPrev {
		Prev = ipage - 1
	} else {
		Prev = 1
	}
	//判断是否有下一页
	IsNext := ipage < count
	//计算下一页页码
	var Next int64
	if IsNext {
		Next = ipage + 1
	} else {
		Next = count
	}

	//获取当前页中的诗词
	sqlStr2 := "select * from poetry limit ?,?"
	//声明model 下的结构体
	var poetry []data.Poetry
	//执行sql
	utils.Db.Select(&poetry, sqlStr2, (ipage-1)*pagesize, pagesize)
	res := data.Page{
		Page:     ipage,
		PageSize: pagesize,
		Count:    count,
		Num:      Num,
		IsPrev:   IsPrev,
		IsNext:   IsNext,
		Prev:     Prev,
		Next:     Next,
	}

	return res, poetry
}

//Getpagesearchs 获取带分页的搜索诗词信息
func Getpagesearchs(page, keyword string) (data.Page, []data.Poetry) {
	//将传过来的页码转换为int
	ipage, _ := strconv.ParseInt(page, 10, 64)
	//获取总记录数
	//sql语句(sqlite没有CONCAT()函数)
	var sqlStr string
	if utils.Conf.Database == "mysql" {
		sqlStr = "SELECT COUNT(*) FROM poetry WHERE CONCAT(title,epoch,author,content) LIKE ?" //mysql
	} else {
		sqlStr = "SELECT COUNT(*) FROM poetry WHERE \"title\"||\"epoch\"||\"author\"||\"content\" LIKE ?" //sqlite
	}

	//总记录数 定义一个变量保存总记录数
	var Num int64
	//执行sql
	utils.Db.Get(&Num, sqlStr, "%"+keyword+"%")

	//设置每页显示的记录数
	var pagesize int64 = 10
	//总页数 定义一个变量保存总页数
	var count int64
	//计算总页数
	if Num%pagesize == 0 {
		//如果总页数除每页条数可以除净 那直接除就可以了
		count = Num / pagesize
	} else {
		//如果总页数除每页条数可以除不净 那就需要在加一页
		count = Num/pagesize + 1
	}

	//判断是否有上一页
	IsPrev := ipage > 1
	//计算上一页页码
	var Prev int64
	if IsPrev {
		Prev = ipage - 1
	} else {
		Prev = 1
	}
	//判断是否有下一页
	IsNext := ipage < count
	//计算下一页页码
	var Next int64
	if IsNext {
		Next = ipage + 1
	} else {
		Next = count
	}

	//获取当前页中的诗词(sqlite没有CONCAT()函数)
	var sqlStr2 string
	if utils.Conf.Database == "mysql" {
		sqlStr2 = "SELECT * FROM `poetry` WHERE CONCAT(title,epoch,author,content) LIKE ? limit ?,?" //mysql
	} else {
		sqlStr2 = "SELECT * FROM `poetry` WHERE \"title\"||\"epoch\"||\"author\"||\"content\" LIKE ? limit ?,?" //sqlite
	}

	//声明model 下的结构体
	var poetry []data.Poetry
	//执行sql
	utils.Db.Select(&poetry, sqlStr2, "%"+keyword+"%", (ipage-1)*pagesize, pagesize)
	res := data.Page{
		Page:     ipage,
		PageSize: pagesize,
		Count:    count,
		Num:      Num,
		IsPrev:   IsPrev,
		IsNext:   IsNext,
		Prev:     Prev,
		Next:     Next,
	}

	return res, poetry
}

//Getpageepochs 获取带分页的朝代诗词信息
func Getpageepochs(page, epoch string) (data.Page, []data.Poetry) {
	//将传过来的页码转换为int
	ipage, _ := strconv.ParseInt(page, 10, 64)
	//获取总记录数
	//sql语句
	sqlStr := "SELECT COUNT(*) FROM poetry where epoch=?"
	//总记录数 定义一个变量保存总记录数
	var Num int64
	//执行sql
	utils.Db.Get(&Num, sqlStr, epoch)
	//设置每页显示的记录数
	var pagesize int64 = 10
	//总页数 定义一个变量保存总页数
	var count int64
	//计算总页数
	if Num%pagesize == 0 {
		//如果总页数除每页条数可以除净 那直接除就可以了
		count = Num / pagesize
	} else {
		//如果总页数除每页条数可以除不净 那就需要在加一页
		count = Num/pagesize + 1
	}

	//判断是否有上一页
	IsPrev := ipage > 1
	//计算上一页页码
	var Prev int64
	if IsPrev {
		Prev = ipage - 1
	} else {
		Prev = 1
	}
	//判断是否有下一页
	IsNext := ipage < count
	//计算下一页页码
	var Next int64
	if IsNext {
		Next = ipage + 1
	} else {
		Next = count
	}

	//获取当前页中的诗词
	sqlStr2 := "select * from poetry where epoch=? limit ?,?"
	//声明model 下的结构体
	var poetry []data.Poetry
	//执行sql
	utils.Db.Select(&poetry, sqlStr2, epoch, (ipage-1)*pagesize, pagesize)
	res := data.Page{
		Page:     ipage,
		PageSize: pagesize,
		Count:    count,
		Num:      Num,
		IsPrev:   IsPrev,
		IsNext:   IsNext,
		Prev:     Prev,
		Next:     Next,
	}

	return res, poetry
}

//Getpagepoets 获取带分页的作者诗词信息
func Getpagepoets(page, poet string) (data.Page, []data.Poetry) {
	//将传过来的页码转换为int
	ipage, _ := strconv.ParseInt(page, 10, 64)
	//获取总记录数
	//sql语句
	sqlStr := "SELECT COUNT(*) FROM poetry where author=?"
	//总记录数 定义一个变量保存总记录数
	var Num int64
	//执行sql
	utils.Db.Get(&Num, sqlStr, poet)
	//设置每页显示的记录数
	var pagesize int64 = 10
	//总页数 定义一个变量保存总页数
	var count int64
	//计算总页数
	if Num%pagesize == 0 {
		//如果总页数除每页条数可以除净 那直接除就可以了
		count = Num / pagesize
	} else {
		//如果总页数除每页条数可以除不净 那就需要在加一页
		count = Num/pagesize + 1
	}

	//判断是否有上一页
	IsPrev := ipage > 1
	//计算上一页页码
	var Prev int64
	if IsPrev {
		Prev = ipage - 1
	} else {
		Prev = 1
	}
	//判断是否有下一页
	IsNext := ipage < count
	//计算下一页页码
	var Next int64
	if IsNext {
		Next = ipage + 1
	} else {
		Next = count
	}

	//获取当前页中的诗词
	sqlStr2 := "select * from poetry where author=? limit ?,?"
	//声明model 下的结构体
	var poetry []data.Poetry
	//执行sql
	utils.Db.Select(&poetry, sqlStr2, poet, (ipage-1)*pagesize, pagesize)
	res := data.Page{
		Page:     ipage,
		PageSize: pagesize,
		Count:    count,
		Num:      Num,
		IsPrev:   IsPrev,
		IsNext:   IsNext,
		Prev:     Prev,
		Next:     Next,
	}

	return res, poetry
}

//Getcontent 获取诗词信息
func Getcontent(id string) (data.Poetry, []string) {
	//初始化诗词结构体
	var poetry data.Poetry
	//sql语句
	sqlStr := "SELECT * FROM poetry where id=?"
	//执行sql
	utils.Db.Get(&poetry, sqlStr, id)

	str := strings.Replace(poetry.Content, "。", "，", -1) //将全部句号换成逗号
	str1 := strings.Trim(str, "，")                       //将左右两边的逗号去掉
	content := strings.Split(str1, "，")                  //将诗词以逗号分割为数组

	return poetry, content
}
