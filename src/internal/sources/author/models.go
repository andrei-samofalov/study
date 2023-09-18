package sources

type Author struct {
	ID        int    `uri:"id" json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	IsDeleted bool   `json:"is_deleted" form:"is_deleted"`
}
