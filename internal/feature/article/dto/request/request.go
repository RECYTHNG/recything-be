package request

type ArticleRequest struct {
	Title       string   `from:"title"`
	Description string   `from:"description"`
	Category_id []string `form:"category_id"`
	Thumbnail   string   `from:"thumbnail"`
	Categories  []ArticleTrashCategoryRequest
}

type ArticleTrashCategoryRequest struct {
	Category string
}
