package entity

import (
	tcm "github.com/sawalreverr/recything/internal/trash_category/model"
)

func CategoryModelToCategoryCore(category tcm.TrashCategory) ArticleTrashCategoryCore {
	return ArticleTrashCategoryCore{
		// TrashCategoryID: category.ID,
		Category: category.TrashType,
	}
}

func ListCategoryModelToCategoryCore(category []tcm.TrashCategory) []ArticleTrashCategoryCore {
	coreCategory := []ArticleTrashCategoryCore{}
	for _, v := range category {
		category := CategoryModelToCategoryCore(v)
		coreCategory = append(coreCategory, category)
	}
	return coreCategory
}

func CategoryCoreToCategoryModel(category ArticleTrashCategoryCore) tcm.TrashCategory {
	return tcm.TrashCategory{
		// ID:        category.TrashCategoryID,
		TrashType: category.Category,
	}
}

func ListCategoryCoreToCategoryModel(category []ArticleTrashCategoryCore) []tcm.TrashCategory {
	coreCategorys := []tcm.TrashCategory{}
	for _, v := range category {
		categorys := CategoryCoreToCategoryModel(v)
		coreCategorys = append(coreCategorys, categorys)
	}
	return coreCategorys
}
