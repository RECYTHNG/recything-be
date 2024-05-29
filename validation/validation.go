package validation

import (
	"errors"
	"strconv"
	"strings"

	"github.com/sawalreverr/recything/pkg"
)

func ValidateParamsPagination(page, limit string) (int, int, error) {
	var limitInt int
	var pageInt int
	var err error
	if limit == "" {
		limitInt = 10
	}
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, errors.New("limit harus berupa angka")
		}

		if limitInt > 10 {
			return 0, 0, errors.New("limit tidak boleh lebih dari 10")
		}
	}

	if page == "" {
		pageInt = 1
	}
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, errors.New("page harus berupa angka")
		}
	}

	pageInt, limitInt = ValidateCountLimitAndPage(pageInt, limitInt)
	return pageInt, limitInt, nil

}

func ValidateCountLimitAndPage(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	maxLimit := 10
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}

func CheckEqualData(data string, validData []string) (string, error) {
	inputData := strings.ToLower(data)

	isValidData := false
	for _, data := range validData {
		if inputData == strings.ToLower(data) {
			isValidData = true
			break
		}
	}

	if !isValidData {
		return "", errors.New(pkg.ERROR_INVALID_INPUT)
	}

	return inputData, nil
}
