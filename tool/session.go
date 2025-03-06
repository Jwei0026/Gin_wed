package tool

import (
	"2024-2025/user"
	"encoding/gob"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

//此工具包专门处理session数据

//session启动与配置

var store *sessions.CookieStore

func init() {
	//session的初始化
	store = sessions.NewCookieStore([]byte("recommend:32bytes-secrect")) //这里至少用32个字符加密

	//session的配置
	store.Options = &sessions.Options{
		// 主要配置生命周期还有cookie相关的内容
		Path:     "/",
		MaxAge:   0,  //浏览器关闭时自动失效
		Domain:   "", // 是否跨域：同一个高级别域名（taobao.com == www.taobao.com + m.taobao.com）
		Secure:   false,
		HttpOnly: true,
		// SameSite：cookie是否跨站：不同的站点（www.taobao.com + www.jd.com）
		SameSite: http.SameSiteLaxMode,
	}

	gob.Register(user.Users{}) //使用复杂结构之前得告诉session
}

func Setsessions(ctx *gin.Context, key string, value interface{}) error {
	//得先获取到cookie
	session, err := store.Get(ctx.Request, "SHOPID") //shopID为一个特定的cookie会话ID
	if err != nil {
		return err
	}

	//对session进行赋值
	session.Values[key] = value

	//保存session到文件
	return session.Save(ctx.Request, ctx.Writer)
}

func Getsessions(ctx *gin.Context, key string) interface{} {
	session, err := store.Get(ctx.Request, "SHOPID")
	if err != nil {
		return err
	}

	//返回session内容
	return session.Values[key]
}
