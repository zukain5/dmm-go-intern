package object

type Relationship struct {
	ID         int64 `json:"id"`
	Following  bool  `json:"following"`
	FollowedBy bool  `json:"followed_by"`
}
