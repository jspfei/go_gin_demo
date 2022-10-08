package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) Sendcode(phone string) bool {

	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	fmt.Println(code)
	//2.调用阿里云sdk 完成发送
	// 接受返回结果，并判断发送状态

	//将数据保存到数据库中
	smsCode := model.SmsCode{Phone: phone, Code: code, BizId: "1", CreateTime: time.Now().Unix()}

	//插入数据库服务
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.InsertCode(smsCode)
	return result > 0

	return false
}
