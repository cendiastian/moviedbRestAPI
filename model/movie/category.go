package movie

type Category struct {
	Name  string   `json:"name" gorm:"unique"`
	ID    int      `gorm:"primaryKey, AUTO_INCREMENT" json:"id"`
	Movie []*Movie `gorm:"many2many:movie_category;" json:"Genre"`
}
