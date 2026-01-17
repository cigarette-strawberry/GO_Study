package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello 为指定的人返回问候语.
func Hello(name string) (string, error) {
	// 如果没有给出名字，返回一个带有消息的错误.
	if name == "" {
		return "", errors.New("empty name")
	}

	// 如果接收到名称，则返回一个嵌入名称的值
	// 在问候消息中.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos 返回一个map，该地图将每个已命名的人员
// 与问候消息相关联.
func Hellos(names []string) (map[string]string, error) {
	// 将名称与消息关联的map.
	messages := make(map[string]string)
	// 遍历接收到的名称切片，调用
	// Hello 函数为每个名字获取一条消息.
	// 使用 Go 空白标识符（下划线）来忽略
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// 在map中，将检索到的消息与
		// 名称相关联.
		messages[name] = message
	}
	return messages, nil
}

// init 为函数中使用的变量设置初始值.
/* func init() {
	// 获取当前时间的纳秒数作为种子
	seed := time.Now().UnixNano()
	// 创建一个新的随机数生成器
	r := rand.New(rand.NewSource(seed))

	// 使用新的随机数生成器生成随机数
	randomNumber := r.Intn(100)
	fmt.Println(randomNumber)
} */

// randomFormat 返回一组问候消息中的一个。返回的
// 消息是随机选择的.
func randomFormat() string {
	// 消息格式的切片.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 通过为格式切片指定随机索引
	// rand.Intn(n) 是 Go 语言 rand 包中的函数，它会返回一个介于 0（包括）和 n（不包括）之间的伪随机整数。
	// 来返回随机选择的消息格式.
	return formats[rand.Intn(len(formats))]
}
