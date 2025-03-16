package storage

type Post struct {
	ID          int
	AuthorID    int
	Title       string
	Content     string
	Created     int64
	Author_name string // Для монго

}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	GetAllPosts() ([]Post, error) // Получение всех публикаций
	AddPost(Post) error           // Создание новой публикации
	UpdatePost(Post) error        // Обновление публикации
	DeletePost(Post) error        // Удаление публикации по ID
}
