package tool

import (
	"fmt"
	"math/rand/v2"
)

func GetCaptcha() string {
	//定义一个空的字符串
	var captcha string

	//随机生成一个长度为5的验证码
	for i := 0; i < 5; i++ {

		myrand := rand.IntN(3) // 生成0-3的随机谁[0,3)
		var ascii int          //ascii码的序号
		if myrand == 0 {
			//数字
			ascii = rand.IntN(9) + 49
		} else if myrand == 1 {
			//大写
			ascii = rand.IntN(26) + 65
		} else {
			//小写
			ascii = rand.IntN(26) + 97
		}
		fmt.Println(ascii)
		//生产目标码
		captcha += fmt.Sprintf("%c", ascii)
	}

	return captcha
}
