package repository

import (
	"database/sql"

	"github.com/pkg/errors"
	"loong.me/gopher/internal/errcode"
	"loong.me/gopher/model"
)

// Hello 测试
func Hello() ([]model.Config, error) {
	var result []model.Config
	/* Access Database */
	err := sql.ErrNoRows
	if err != nil {
		return nil, errors.Wrap(errcode.ErrNotFound, "找不到记录")
	}
	return result, nil
}
