package common

type User struct {
	UserId        int    `json:"id"`
	UserId2       int    `json:"user_id" gorm:"-"`
	Username      string `json:"name";gorm:"column:username"`
	Password      string `json:"password";gorm:"column:password"`
	FollowCount   int    `json:"follow_count";gorm:"column:follow_count"`
	FollowerCount int    `json:"follower_count";gorm:"column:follower_count"`
}

type Video struct {
	VideoId       int    `json:"id";gorm:"column:video_id"`
	Title         string `json:"title"`
	AuthorId      int    `json:"author_id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" gorm:"-"`
	User          User   `json:"author" gorm:"foreignKey:VideoId;references:user_id"`
}

type VideoFollow struct {
	Id      int `json:"id"`
	VideoId int `json:"video_id"`
	UserId  int `json:"user_id"`
}

type VideoComment struct {
	Id          int    `json:"id"`
	VideoId     int    `json:"video_id"`
	UserId      int    `json:"user_id"`
	CreateDate  string `json:"create_date"`
	CommentText string `json:"content"`
	User        User   `json:"user" gorm:"foreignKey:Id;references:user_id"`
}

func (Video) TableName() string {
	return "video"
}

func (VideoFollow) TableName() string {
	return "video_follow"
}

func (VideoComment) TableName() string {
	return "video_comment"
}
