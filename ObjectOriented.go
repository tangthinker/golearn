package main

import "fmt"

//类成员变量
type User struct {
	username string
	password string
	age      int
	sex      string
	phone    string
}

//构造函数
func newUser(username string, password string, age int, sex string, phone string) User {
	return User{username, password, age, sex, phone}
}

//类方法
func (user User) printUser() {
	fmt.Println("I am ", user.username, ", my password is ", user.password,
		", And my phone number is ", user.phone, " I am a ", user.sex, "!")
}

func main() {
	user := newUser("shanliao", "333", 21, "boy", "10086")
	user.printUser()
}
