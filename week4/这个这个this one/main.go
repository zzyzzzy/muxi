package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 用户结构体
type User struct {
	UserName string `json:"username"` //用户名
	UserId   string `json:"userid"`   //昵称，为什么不是nickname，因为写错了
	PassWord string `json:"password"` //密码
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// Session结构体
type Session struct {
	UserName string    //用户名，关联到哪个用户
	Expiry   time.Time //过期时间，超过这个时间就失效
}

// 相应结构体
type Response struct {
	Success bool        `json:"success"` //是否成功
	Message string      `json:"message"` //提示消息
	Data    interface{} `json:"data"`    //返回数据，可以是任意类型
}

//请求登录结构体
// type LoginResquest struct{
// 	UserId string `json:"userid"`
// 	PassWord string `json:"password"`
// }
// //
// type ResgisterResquest struct{
// 	UserName string  `json:"username"`
// 	UserId string  `json:"userid"`
// 	PassWord string  `json:"password"`
// 	Email string  `json:"email"`
// 	Age string  `json:"age"`
// }

// key是用户名，value是user对象
var users = make(map[string]User)

// key是sessionID，value是Session对象
var sessions = make(map[string]Session)

// 创建一个函数，为用户创建一个新的会话
func CreateSession(username string) string {
	//生成一个唯一的SessionID，用用户名和时间戳确保唯一性
	SessionID := fmt.Sprintf("%s-%d", username, time.Now().UnixNano())

	//设置会话过期时间为24小时后
	expiry := time.Now().Add(24 * time.Hour)

	sessions[SessionID] = Session{
		UserName: username,
		Expiry:   expiry,
	}
	return SessionID
}

// GetSession函数，根据sessionID获取会话信息
func GetSession(SessionID string) (Session, bool) {
	//查找是否存在session
	session, exists := sessions[SessionID]
	if !exists {
		return Session{}, false //没找到，返回空session和false
	}
	//检查session是否过期
	if time.Now().After(session.Expiry) {
		//过期了就删除这个session
		delete(sessions, SessionID)
		return Session{}, false
	}
	return session, true //找到了有效的session
}

// cooike管理
// SetSessionCooike函数，在浏览器中设置session cookie
func SetSessionCookie(w http.ResponseWriter, SessionID string) {
	//创建一个新的cookie
	cookie := &http.Cookie{
		Name:     "session_id", //cookie的名称
		Value:    SessionID,    //cookie的值(SessionID)
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	//设置cooki到相应中
	http.SetCookie(w, cookie)
}

// ClearSessionCoolie函数，清除浏览器的session cookie
func ClearSessionCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	}
	http.SetCookie(w, cookie)
}

// GetCurrentUser函数，从请求中获取当前登录的用户
func GetCurrentUser(r *http.Request) (User, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return User{}, false //没有cookie，用户位登录
	}
	//根据cookie中的SessionID获取session信息
	session, exists := GetSession(cookie.Value)
	if !exists {
		return User{}, false //session无效或过期
	}
	//根据session中的用户名获取用户信息
	user, exists := users[session.UserName]
	if !exists {
		return User{}, false //用户不存在
	}
	return user, true
}

// RespondJson函数，统一json，让所有API返回相同格式的相应
func RespondJson(w http.ResponseWriter, status int, data Response) {
	//设置响应头，告诉浏览器返回的是json格式
	w.Header().Set("Content-Type", "application/json")
	//设置http状态码
	w.WriteHeader(status)
	//将Response结构体转换为json并写入响应
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println("json err=", err)
	}
	//deepseek不处理错误，我已严厉批评了它，但是这样处理错误，好像也并没有什么用嘻嘻
}

// API处理
// RegisterHandle，处理注册
func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	//只接收Post请求
	if r.Method != "POST" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持post请求",
		})
		return
	}
	//解析请求体中的json数据
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		RespondJson(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "请求数据格式错误",
		})
		return
	}
	//检查用户名是否存在
	if _, exists := users[newUser.UserName]; exists {
		RespondJson(w, http.StatusConflict, Response{
			Success: false,
			Message: "用户已存在",
		})
		return
	}
	//将新用户储存到users的map中
	users[newUser.UserName] = newUser
	//返回成功相应
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "注册成功",
	})
}

// LoginHandle，处理用户登录
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持POST请求",
		})
		return
	}
	//解析登录请求数据
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password`
	}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		RespondJson(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "请求数据格式错误",
		})
		return
	}

	//检查用户是否存在
	user, exists := users[loginData.Username]
	if !exists {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "用户名或密码错误",
		})
		return
	}
	//检查密码是否正确（简单版本，直接比较字符串）
	if user.PassWord != loginData.Password {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "用户名或密码错误",
		})
		return
	}
	//创建session
	SessionID := CreateSession(loginData.Username)
	//设置cookie
	SetSessionCookie(w, SessionID)
	//返回成功相应和用户信息(不包含密码)
	userInfo := map[string]interface{}{
		"username": user.UserName,
		"email":    user.Email,
		"age":      user.Age,
	}
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "登录成功",
		Data:    userInfo,
	})
}

// LogoutHandler处理用户注销
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持post请求",
		})
		return
	}
	//获取当前session cookie
	cookie, err := r.Cookie("session_id")
	if err != nil {
		//从session里删除这个session
		delete(sessions, cookie.Value)

	}
	//清除浏览器的cookie
	ClearSessionCookie(w)
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "注销成功",
	})
}

// ChangePassword ,修改密码
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持post请求",
		})
		return

	}
	//检查登录
	user, exists := GetCurrentUser(r)
	if !exists {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "请先登录",
		})
		return
	}
	//解析修改密码请求
	var PasswordData struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	err := json.NewDecoder(r.Body).Decode(&PasswordData)
	if err != nil {
		RespondJson(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "请求数据结构错误",
		})

		return
	}
	//检查旧密码是否正确
	if user.PassWord != PasswordData.OldPassword {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "旧密码错误",
		})
		return
	}
	//更新密码
	user.PassWord = PasswordData.NewPassword
	users[user.UserName] = user
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "修改密码成功",
	})
}

// GetUserInfo 获取用户当前信息
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持Get请求",
		})
		return
	}
	//检查登录
	user, exists := GetCurrentUser(r)
	if !exists {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "请先登录",
		})
		return
	}
	//返回用户信息(不含密码)
	userInfo := map[string]interface{}{
		"username": user.UserName,
		"userid":   user.UserId,
		"email":    user.Email,
		"age":      user.Age,
	}
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "获取用户信息成功",
		Data:    userInfo,
	})
}

// updataUserInfo,更新用户信息
func UpdataUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		RespondJson(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "只支持post请求",
		})
		return
	}
	//检查用户是否登录
	user, exists := GetCurrentUser(r)
	if !exists {
		RespondJson(w, http.StatusUnauthorized, Response{
			Success: false,
			Message: "请先登录",
		})
		return
	}
	//解析要更新的数据
	var updataData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updataData)
	if err != nil {
		RespondJson(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "请求数据格式错误",
		})
		return
	}
	//更新用户信息
	//更新昵称，嘻嘻嘻写错了，请把id当成昵称
	if userid, ok := updataData["userid"].(string); ok {
		user.UserId = userid
	}
	//检查并更新邮箱
	if email, ok := updataData["email"].(string); ok {
		user.Email = email
	}
	//检查并更新年龄
	if age, ok := updataData["age"].(float64); ok {
		user.Age = int(age)
	}
	//保存更新后的用户信息
	users[user.UserName] = user
	//返回更新后的用户信息
	userInfo := map[string]interface{}{
		"usesrname": user.UserName,
		"userid":    user.UserId,
		"email":     user.Email,
		"age":       user.Age,
	}
	RespondJson(w, http.StatusOK, Response{
		Success: true,
		Message: "用户信息更新成功",
		Data:    userInfo,
	})
}

// 啊啊啊终于到主程序了我哭透了
func main() {
	//设置路由
	//用户注册
	http.HandleFunc("/register", RegisterHandle)
	//用户登录
	http.HandleFunc("/login", LoginHandle)
	//用户注销
	http.HandleFunc("/logout", LogoutHandler)
	//修改密码
	http.HandleFunc("/change-password", ChangePassword)
	//获取用户信息
	http.HandleFunc("/user/info", GetUserInfo)
	//更新用户信息
	http.HandleFunc("/user/updata", UpdataUserInfo)
	//打印启动信息
	fmt.Println("用户管理系统后端服务启动")
	fmt.Println("服务地址:http://localhost:8080")
	fmt.Println()
	fmt.Println("可用接口：")
	fmt.Println("POST/register         -用户注册")
	fmt.Println("POST /login           - 用户登录")
	fmt.Println("POST /logout          - 用户注销")
	fmt.Println("POST /change-password - 修改密码")
	fmt.Println("GET  /user/info       - 获取用户信息")
	fmt.Println("POST /user/update     - 更新用户信息")
	fmt.Println("")
	fmt.Println("使用Ctrl+C停止服务")
	//启动http服务器 ，8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败err=", err)

	}
}
