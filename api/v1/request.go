package v1

type UserRegisterReq struct {
	UserName     string `json:"username"`
	PassWord     string `json:"password"`
	RegisterTime string `json:"registered"`
	TellPhone    string `json:"telephone"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Captcha      string `json:"captcha"`
	Gender       string `json:"gender"`
}
type UserLoginUPReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type UserLoginEPReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type FriendAddReq struct {
	OwnerName string `json:"ownername"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}
type CreateGroupReq struct {
	OwnerUserName string `json:"ownerusername"`
	GroupName     string `json:"groupname"`
}
