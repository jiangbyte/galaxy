package utils

import "regexp"

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// ValidatePhone 验证手机号格式
func ValidatePhone(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	matched, _ := regexp.MatchString(pattern, phone)
	return matched
}

// ValidateUsername 验证用户名格式
func ValidateUsername(username string) bool {
	// 用户名只能包含字母、数字、下划线，长度3-20位
	pattern := `^[a-zA-Z0-9_]{3,20}$`
	matched, _ := regexp.MatchString(pattern, username)
	return matched
}
