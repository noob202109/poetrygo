package data

//Navbar 导航栏结构体
type Navbar struct {
	Epoch string `db:"epoch"` //朝代
}

//Poetry 诗词结构体
type Poetry struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`   // 诗词名
	Epoch   string `json:"epoch"`   // 朝代
	Author  string `json:"author"`  // 作者
	Content string `json:"content"` // 诗词内容
}

//Index 首页结构体
type Index struct {
	Navbar []Navbar //导航栏
	Page   Page     //分页
	Poetry []Poetry //诗词数组
}

//Other 其他页结构体
type Other struct {
	Navbar []Navbar //导航栏
	Page   Page     //分页
	Poetry []Poetry //数次数组
	Param  string   //参数
}

//Content 内容页结构体
type Content struct {
	Navbar  []Navbar //导航栏
	Poetry  Poetry   //诗词
	Poetry1 []string //诗词内容数组
}

//Page 分页结构体
type Page struct {
	Page     int64 //当前页
	PageSize int64 //每页显示的条数
	Count    int64 //总页数，通过计算得到
	Num      int64 //总记录数，通过查询数据库得到
	IsPrev   bool  //是否有上一页
	IsNext   bool  //是否有下一页
	Prev     int64 //上一页页码
	Next     int64 //下一页页码
}
