package gitter

type User struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	DisplayName     string `json:"displayName"`
	Url             string `json:"url"`
	AvatarUrlSmall  string `json:"avatarUrlSmall"`
	AvatarUrlMedium string `json:"avatarUrlMedium"`
}
