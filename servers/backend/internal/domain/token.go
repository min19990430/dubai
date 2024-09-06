package domain

type TokenClaims struct {
	UserUUID    string `json:"user_uuid"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	NickName    string `json:"nick_name"`
	AuthorityID string `json:"authority_id"`
}
