package user

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserFriends struct {
	Friends []int `json:"friends"`
}

type FriendsRequest struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}