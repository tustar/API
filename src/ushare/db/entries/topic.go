package entries

type Topic struct {
	BaseModel
	Title  string `json:"title" form:"title"`
	Desc   string `json:"description" form:"description"`
	Begin  int64  `json:"begin" form:"begin"`
	End    int64  `json:"end" form:"end"`
	Shared bool   `json:"shared" form:"shared"`
	UserID int64  `json:"user_id" form:"user_id"`
}
