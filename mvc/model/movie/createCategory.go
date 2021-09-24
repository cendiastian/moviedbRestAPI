package movie

type CreateCategory struct {
	Name string `json:"Name" gorm:"unique"`
}
