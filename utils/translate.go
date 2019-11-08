package utils

import "github.com/letanthang/echo_stackdriver/types"

type TransLateResult struct {
	HttpCode int
	Message  string
}

var Directory = map[string]TransLateResult{
	"BAD_REQUEST":        {400, "Thông tin không hợp lệ, vui lòng kiểm tra lại"},
	"NOT_FOUND":          {404, "Không tìm thấy thông tin, vui lòng kiểm tra lại"},
	"PHONE_EXISTS":       {409, "Số điện thoại đã tồn tại, vui lòng kiểm tra lại"},
	"EMAIL_EXISTS":       {409, "Email đã tồn tại, vui lòng kiểm tra lại"},
	"UNKNOWN":            {400, "Thông tin không hợp lệ, vui lòng kiểm tra lại"},
	"FULLNAME_INVALID":   {400, "Họ và Tên không hợp lệ, vui lòng kiểm tra lại."},
	"ADDRESS_INVALID":    {400, "Địa chỉ không hợp lệ, vui lòng kiểm tra lại."},
	"EMAIL_INVALID":      {400, "Định dạng email không hợp lệ, vui lòng kiểm tra lại"},
	"PHONE_INVALID":      {400, "Số điện thoại không hợp lệ, vui lòng kiểm tra lại"},
	"NOT_MATCH_OLD_PASS": {400, "Mật khẩu cũ không đúng, vui lòng kiểm tra lại."},
	"WRONG_PASS":         {400, "Mật khẩu không đúng, vui lòng kiểm tra lại."},
	"CONTACT_TYPE_WRONG": {400, "Loại liên hệ không tồn tại, vui lòng kiểm tra lại."},
	"SERVICE_NOT_FOUND":  {404, "Không tìm thấy thông tin dich vụ."},
	"PROFILE_NOT_FOUND":  {404, "Không tìm thấy thông tin tài khoản."},
	"ACCOUNT_NOT_FOUND":  {404, "Không tìm thấy thông tin tài khoản."},
	"ACCOUNT_EXISTS":     {409, "Tài khoản đã tồn tại."},
	"LOGIN_FAIL":         {400, "Mật khẩu không đúng, vui lòng kiểm tra lại."},
}

const (
	CODE_BAD_REQUEST = "BAD_REQUEST"
	CODE_NOT_FOUND   = "NOT_FOUND"
)

func Translate(code string) TransLateResult {
	rs := Directory[code]
	if rs.HttpCode == 0 {
		rs.HttpCode = 400
		rs.Message = "Có lỗi xảy ra, vui lòng kiểm tra lại"
	}
	return rs
}
func ResponseCode(code string) int {
	tr := Translate(code)
	return tr.HttpCode
}

func ResponseBody(code string) types.ErrorResponse {
	tr := Translate(code)

	return types.ErrorResponse{ErrorCode: code, Message: tr.Message}
}
