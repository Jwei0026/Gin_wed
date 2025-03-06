package main

import (
	"2024-2025/brands"
	"2024-2025/products"
	"2024-2025/tool"
	"2024-2025/user"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	//1.open server ：所有路由都在这里
	//为了方便调试，使用gin自带中间件Logger（） Recover（）
	r := gin.Default()

	//静态资源加载：static访问
	r.StaticFS("/static", http.Dir("./static")) //第一个参数是自己定义的一个路径

	//加载模板文件
	r.LoadHTMLGlob("./tpl/*")

	//定义一个高于路由的变量用于存储用户信息
	var users user.Users
	//中间件，验证用户是否登录（如果用户登录，则存到某个变量，未来在需要时进行验证和调用）
	r.Use(func(ctx *gin.Context) {
		//验证用户是否登录,看session中是否有user信息
		val := tool.Getsessions(ctx, "user") //使用val装得到的val，此时的是序列化的字符串
		users, _ = val.(user.Users)          //将val进行强制类型转换,如果这里用user的话，会导致命名冲突
		fmt.Println(0, users, users.Username)

		ctx.Next()
	})

	//1.首页路由，输入网址即可，不需要对应路径（直接根目录路由）
	r.GET("/", func(ctx *gin.Context) {
		//获取首页相关数据

		products := products.GetIndexProudcts()
		// fmt.Println(products)

		//最后加载模板，展示页面
		ctx.HTML(200, "index.html", gin.H{
			"products": products, //如果某个模板参数出错，会导致后续的模板都无法输出
			"user":     users,
		})
		/*
			问题
			1.index.html（原index-JS.html）gin捕获到问题
			1.1有一个html注释只有一半 24
			1.2 167嵌入超链接有误 href为缺，或多余可去掉
			2.格式出现，但样式出问题  -->静态资源（js，css）不对

		*/

	})

	//2.特价路由
	r.GET("/tejia", func(ctx *gin.Context) {
		//获取总记录数

		//获取特价商品的相关信息
		newproducts := products.GetSpecialProducts()

		// fmt.Println(products)

		//加载模板，并结合数据
		ctx.HTML(200, "tejia.html", gin.H{
			"SpecialProducts": newproducts, //使用模板技术传数据传的是SpecialProducts，不是products
			"user":            users,
		})
	})

	//3.男士商品路由
	r.GET("/boys", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		current, _ := strconv.Atoi(page)
		//获取新品商品的相关信息
		newProducts := products.GetNewProducts()
		//获取男士商品的相关信息
		manproducts := products.GetGenderProducts("男", current, tool.GetPageCount())
		//获取品牌相关信息
		brands := brands.GetBrands()
		//验证
		// fmt.Println(brands)
		// fmt.Println(manproducts)
		// fmt.Println(newProducts)

		//加载模板，并结合数据
		ctx.HTML(200, "boys.html", gin.H{
			"ManProducts": manproducts,
			"brands":      brands,
			"newProducts": newProducts,
			"user":        users,
			"router":      "boys",
			"pageInfo":    tool.GetPage(products.GetGenderCounts("男"), current),
		})
	})

	//4.女士商品路由
	r.GET("/girls", func(ctx *gin.Context) {
		page := ctx.DefaultQuery("page", "1")
		current, _ := strconv.Atoi(page)

		//获取新品商品的相关信息
		newProducts := products.GetNewProducts()
		//获取女士商品的相关信息
		Girlproducts := products.GetGenderProducts("女", current, tool.GetPageCount())
		//获取品牌相关信息
		brands := brands.GetBrands()
		//验证
		// fmt.Println(brands)
		// fmt.Println(Girlproducts)
		// fmt.Println(newProducts)

		//加载模板，并结合数据
		ctx.HTML(200, "girl.html", gin.H{
			"GirlProducts": Girlproducts,
			"brands":       brands,
			"newProducts":  newProducts,
			"user":         users,
			"router":       "girls",
			"pageInfo":     tool.GetPage(products.GetGenderCounts("女"), current),
		})
	})

	//5.登录功能
	//5.1登录入口
	r.GET("/login", func(ctx *gin.Context) {
		referer := ctx.Request.Header.Get("referer")
		ctx.HTML(200, "login.html", gin.H{
			"referer": referer,
		})
	})

	//5.2登录验证
	r.POST("/checkLogin", func(ctx *gin.Context) {
		//接受数据：用户名与密码 解析表单参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		referer := ctx.PostForm("referer") //非用户输出，不需要太多验证

		//数据加工（去除两边多余的可能出现的空格）
		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)
		// fmt.Println(username, password)

		//验证数据:合法（是否满足规则）性与合理（是否匹配数据）性验证
		//简单验证,输入不为空
		if username == "" {
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   "用户名不能为空！", //待修改
				"target": "/login",
			})
			return
		}

		if password == "" { //如果没有输入密码则直接进入这里，所以数据库查不到
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   "密码不能为空！",
				"target": "/login",
			})
			return
		}
		//合理性验证
		// fmt.Println("这是主函数输出账户：", username, "密码：", password)
		user, err := user.CheckLogin(username, password)

		if err != nil {
			//输入的内容存在错误
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   err.Error(),
				"target": "/login",
			})

			return
		}

		fmt.Println(user)
		err = tool.Setsessions(ctx, "user", user)
		if err != nil {
			//session存的失败
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   err.Error(),
				"target": "/login",
			})

			return
		}

		//登录成功，记录用户登录状态（将用户信息存到session中）

		//跳转到首页（或者是用户上一次进入登录页的页面，通过redict ）
		//获取当前请求页面的上一个路由，通过http协议中的referer协议获取
		fmt.Println(referer)
		user1 := tool.Getsessions(ctx, "user")
		fmt.Println(2, user1)
		ctx.HTML(200, "redirect.html", gin.H{
			"tips":   "登录成功！",
			"target": referer,
		})
	})

	//详情路由
	r.GET("/detail/:id", func(ctx *gin.Context) {
		//接收数据，此处接受到的是字符串数据
		id := ctx.Param("id")
		//验证数据 此处一般对传入的数据进行合法合理性验证
		//数据操作
		product := products.GetDetail(id)

		//如果查找模板数据不存在，则进行相应处理
		if product.Product_id == 0 {
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   "找不到商品信息！",
				"target": "/",
			})
			return
		}

		products := products.GetBrandHostProducts(product.Brand_id)
		// fmt.Println(product)
		// fmt.Println(products)
		//返回数据（模板展示数据）
		ctx.HTML(200, "xiangxi.html", gin.H{
			"product":  product,
			"products": products,
			"user":     users,
		})
	})

	//注册功能
	//1.显示表单
	r.GET("/register", func(ctx *gin.Context) {
		//3.加载模板
		ctx.HTML(200, "register.html", gin.H{
			"captcha": tool.GetCaptcha(),
		})
	})

	//2.收集数据
	r.POST("/checkRegister", func(ctx *gin.Context) {
		//接收数据
		var UserInfo user.UserInfoRegister
		if err := ctx.ShouldBindWith(&UserInfo, binding.Form); err != nil {
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   err.Error(),
				"target": "/register",
			})
			return
		}

		fmt.Println(UserInfo)
		//(处理)验证（合法性和合理性）
		//数据保存到数据库
		if err := user.CheckRegister(UserInfo); err != nil {
			ctx.HTML(200, "redirect.html", gin.H{
				"tips":   err.Error(),
				"target": "/register",
			})
			return
		}
		//响应结果
		ctx.HTML(200, "redirect.html", gin.H{
			"tips":   "注册成功，请先登录！",
			"target": "/login",
		})
		//启动服务
	})

	//独立获取验证码
	r.GET("/getCaptcha", func(ctx *gin.Context) {
		ctx.String(200, tool.GetCaptcha())
	})

	//6.智能客服
	r.GET("/chat", func(ctx *gin.Context) {
		// fmt.Println("进入chat")
		//加载模板
		ctx.HTML(200, "chat.html", gin.H{})
		// products.GetDetailone("12", "case_material") //测试成功
		// fmt.Println("chat结束")
	})

	r.GET("/chatsearch", func(ctx *gin.Context) {
		//接收用户传来的数据，获取目标数据：商品id，用户输入内容
		//对输入内容进行关键词抓取
		//通过字符串匹配匹配出关键属性
		//根据关键属性查找对应应该返回的信息
		//返回的内容包括用户查找的属性，以及属性的信息
		// 绑定JSON请求体到结构体
		useinput := ctx.Query("useinput")
		productid := ctx.Query("productid")

		// fmt.Println(useinput)
		// fmt.Println(productid)
		// fmt.Println(useinput)
		// fmt.Printf("over")

		/*进行查询数据处理begin*/
		chineseKeyword, _ := tool.ContainsStringInSlice(useinput)           //捕获字符串中的关键次
		englishKeyword, _ := tool.TranslateChineseToEnglish(chineseKeyword) //将关键词转化为对应属性
		attributemsg := products.GetDetailone(productid, englishKeyword)    //查找对应属性信息
		fmt.Println(chineseKeyword, attributemsg)
		/*进行查询数据处理over*/

		ctx.JSON(200, gin.H{
			"Chinesekeyword": chineseKeyword,
			"Attributemsg":   attributemsg,
		})

	})

	r.Run(":80")

}

// 模板技术即使被注释掉也依旧会生效的
