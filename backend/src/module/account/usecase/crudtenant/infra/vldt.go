package infra

type InputData struct {
	AuthClientID uint   `json:"auth_client_id" validate:"required"`
	Uid          string `json:"uid" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Avatar       string `json:"avatar"`
	AvatarStr    string `json:"avatar_str"`
}
