package routes

import (
	"fmt"
	"gotask/blocks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

// 通用响应消息格式 HTTP协议里携带的消息正文部分
type RespData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 自定义一套错误编码 对应上面的Code
const (
	RESP_OK         = "0"
	RESP_PARAMERR   = "1"
	RESP_LOGINERR   = "2"
	RESP_BLOCKERR   = "3"
	RESP_UNKNOWNERR = "4"
)

// 记录错误信息的键值对
var respMap map[string]string = map[string]string{
	RESP_OK:         "正常",
	RESP_PARAMERR:   "参数异常",
	RESP_LOGINERR:   "登陆失败",
	RESP_BLOCKERR:   "区块链操作失败",
	RESP_UNKNOWNERR: "未知异常",
}

// 通过错误码拿到错误信息
func getRespMsg(code string) string {
	return respMap[code]
}

// 返回数据
func respJsonMsg(c *gin.Context, r *RespData) {
	r.Msg = getRespMsg(r.Code)
	c.JSON(http.StatusOK, r)
}

type UserInfo struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Money struct {
	To    string `json:"To"`
	Value int64  `json:"Value"`
}

// 注册功能函数
func Register(c *gin.Context) {
	// 3.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 1.解析请求消息
	var ui UserInfo
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		resp.Code = RESP_PARAMERR // 参数错误
		fmt.Println("Failed to shouldbindjson", err)
		return
	}
	fmt.Printf("Register:%+v\n", ui)
	// 2.业务规则处理
	err = blocks.Register(ui.UserName, ui.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR // 区块链操作异常
		fmt.Println("failed to Register!!", err)
		return
	}
}

// 登录功能函数
func Login(c *gin.Context) {
	// 3.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 1.解析请求消息
	var ui UserInfo
	err := c.ShouldBindJSON(&ui)
	if err != nil {
		resp.Code = RESP_PARAMERR // 参数错误
		fmt.Println("Failed to shouldbindjson", err)
		return
	}
	fmt.Printf("Login:%+v\n", ui)
	// 2.业务规则处理
	ok, err := blocks.Login(ui.UserName, ui.PassWord)
	if err != nil {
		resp.Code = RESP_BLOCKERR // 区块链操作异常
		fmt.Println("failed to Register!!", err)
		return
	}
	if !ok {
		fmt.Println("User not exists or passwd err!")
		resp.Code = RESP_LOGINERR
		return
	}
	// 4.设置session
	store := ginsession.FromContext(c)
	store.Set("username", ui.UserName)
	store.Set("password", ui.PassWord)
	err = store.Save()
	if err != nil {
		// c.AbortWithError(500, err)
		resp.Code = RESP_UNKNOWNERR
		return
	}
}

// 任务发布功能函数
func Issue(c *gin.Context) {
	// 1.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 2.解析请求消息
	var task blocks.TaskInfox
	err := c.ShouldBindJSON(&task)
	if err != nil {
		resp.Code = RESP_PARAMERR // 参数错误
		fmt.Println("Failed to shouldbindjson", err)
		return
	}
	fmt.Printf("Issue:%+v\n", task)
	// 3.读取session信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		// c.AbortWithStatus(404)
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("Get session username err!")
		return
	}
	password, ok := store.Get("password")
	if !ok {
		// c.AbortWithStatus(404)
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("Get session password err!")
		return
	}
	// 4.业务规则处理
	err = blocks.Issue(username.(string), password.(string), task.Desc, task.Bonus)
	if err != nil {
		resp.Code = RESP_BLOCKERR // 区块链操作异常
		fmt.Println("failed to blocks.Issue!", err)
		return
	}
}

// 挖矿功能函数
func Mint(c *gin.Context) {
	// 3.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 1.解析请求消息
	var mo Money
	err := c.ShouldBindJSON(&mo)
	if err != nil {
		resp.Code = RESP_PARAMERR // 参数错误
		fmt.Println("Failed to shouldbindjson", err)
		return
	}
	fmt.Printf("Mint:%+v\n", mo)
	// 2.业务规则处理
	err = blocks.Mint(mo.To, mo.Value)
	if err != nil {
		resp.Code = RESP_BLOCKERR // 区块链操作异常
		fmt.Println("failed to Mint!!", err)
		return
	}
}

// 任务修改规则
func Update(c *gin.Context) {
	// 1.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 2.解析请求消息
	var task blocks.TaskInfox
	err := c.ShouldBindJSON(&task)
	if err != nil {
		resp.Code = RESP_PARAMERR // 参数错误
		fmt.Println("Failed to shouldbindjson", err)
		return
	}
	fmt.Printf("Issue:%+v\n", task)
	// 3.读取session信息
	store := ginsession.FromContext(c)
	username, ok := store.Get("username")
	if !ok {
		// c.AbortWithStatus(404)
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("Get session username err!")
		return
	}
	password, ok := store.Get("password")
	if !ok {
		// c.AbortWithStatus(404)
		resp.Code = RESP_UNKNOWNERR
		fmt.Println("Get session password err!")
		return
	}
	// 4.业务规则处理
	taskID, err := strconv.Atoi(task.TaskID)
	if err != nil {
		fmt.Println("Failed to strconv.Atoi!", err)
	}
	err = blocks.Update(username.(string), password.(string), task.Comment, task.Status, int64(taskID))
	if err != nil {
		resp.Code = RESP_BLOCKERR // 区块链操作异常
		fmt.Println("Failed to blocks.Update!", err)
		return
	}
}

// 任务查询,列表请求
func Tasklist(c *gin.Context) {
	// 1.组织响应消息(因为有defer所以提前处理)
	var resp RespData   // 定义响应消息结构
	resp.Code = RESP_OK // 初始化为OK

	defer respJsonMsg(c, &resp)
	// 2.解析请求消息
	strpage := c.Query("page")
	ipage, _ := strconv.Atoi(strpage)

	fmt.Printf("Tasklist:ipage = %d\n", ipage)
	// 3.业务规则处理
	tasks, err := blocks.Tasklist()
	if err != nil {
		resp.Code = RESP_BLOCKERR
		fmt.Println("Failed to blocks.Tasklist!", err)
		return
	}
	// 4.重新设置响应消息中的任务列表数据
	// 分页
	begin := (ipage - 1) * 10
	end := ipage * 10
	if end > len(tasks) {
		end = len(tasks)
	}
	data := struct {
		Total int         `json:"total"`
		Data  interface{} `json:"data"`
	}{
		Total: len(tasks),
		Data:  tasks[begin:end],
	}
	resp.Data = data
}
