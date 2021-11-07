package main

import "GeekTime/SecondWeek/dao"

// 我这里目录设置的中文的，可能不能执行，根据代码大体是这个意思
// 我认为不应该是否应该 Wrap 这个 error，准确说不能把这个当成error处理
// 原因再下面的代码

func main() {

	id := 1
	name, err := dao.GetUserNameById(id)
	if err != nil {
		// 打印具体的日志
	}
	if name == "" {
		// 对应的逻辑判断
	}

}
