package controller

type MyCode int64

const (
	CodeSuccess                   MyCode = 200
	CodeFailure                   MyCode = 400
	CodePermissionError           MyCode = 401
	CodePermissionNoRedirectError MyCode = 402

	CodeInvalidParams   MyCode = 1001
	CodeUserExist       MyCode = 1002
	CodeUserNotExist    MyCode = 1003
	CodeInvalidPassword MyCode = 1004
	CodeServerBusy      MyCode = 1005

	CodeInvalidToken      MyCode = 1006
	CodeInvalidAuthFormat MyCode = 1007
	CodeNotLogin          MyCode = 1008

	CodeDictKeyExist     MyCode = 3000
	CodeDictValueExist   MyCode = 3001
	CodeDictTypeNotExist MyCode = 3002
	CodeUploadError      MyCode = 3003
	CodeUploadTypeError  MyCode = 3004

	CodeRoleInsertError MyCode = 3100
)

var msgFlags = map[MyCode]string{
	CodeSuccess:         "success",
	CodeFailure:         "failure",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",

	CodeDictKeyExist:              "字段key已存在，请勿重复添加",
	CodeDictValueExist:            "值已存在，请勿重复添加",
	CodeDictTypeNotExist:          "添加失败，只能添加字典库存在的字段key",
	CodeUploadError:               "上传错误",
	CodeUploadTypeError:           "文件类型不合法",
	CodePermissionError:           "用户没有接口权限",
	CodePermissionNoRedirectError: "用户没有接口权限",
	CodeRoleInsertError:           "添加失败-角色已存在",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
