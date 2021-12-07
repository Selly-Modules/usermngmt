package internal

const (
	// Invalid

	ErrorInvalidName        = "tên không hợp lệ"
	ErrorInvalidPhoneNumber = "số điện thoại không hợp lệ"
	ErrorInvalidEmail       = "email không hợp lệ"
	ErrorInvalidPassword    = "mật khẩu không hợp lệ"
	ErrorInvalidStatus      = "trạng thái không hợp lệ"
	ErrorInvalidRole        = "vai trò không hợp lệ"
	ErrorInvalidOldPassword = "mật khẩu cũ không hợp lệ"
	ErrorInvalidNewPassword = "mật khẩu mới không hợp lệ"
	ErrorInvalidPermission  = "quyền không hợp lệ"
	ErrorInvalidUser        = "người dùng không hợp lệ"
	ErrorInvalidAvatar      = "ảnh đại diện không hợp lệ"

	// NotFound

	ErrorNotFoundPermission = "quyền không tồn tại"
	ErrorNotFoundRole       = "vai trò không tồn tại"
	ErrorNotFoundUser       = "người dùng không tồn tại"

	// AlreadyExisted

	ErrorAlreadyExistedPhoneNumber = "số điện thoại đã tồn tại"
	ErrorAlreadyExistedEmail       = "email đã tồn tại"

	// Incorrect

	ErrorIncorrectPassword = "mật khẩu không chính xác"
)
