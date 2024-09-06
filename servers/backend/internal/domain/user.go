package domain

type User struct {
	UUID        string `json:"uuid"`
	IsEnable    bool   `json:"is_enable"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	FullName    string `json:"full_name"`
	NickName    string `json:"nick_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type UserWithAuthority struct {
	User
	Authority Authority `json:"authority"`
}
