package movie

type Category struct {
	Name string `json:"Name" gorm:"unique"`
	ID   int    `gorm:"primaryKey, AUTO_INCREMENT, unique" json:"id"`
}
