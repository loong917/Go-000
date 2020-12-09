package hello

import (
	"database/sql"

	"github.com/pkg/errors"

	"loong.me/gopher/model/hello"
)

// ErrNotFound 哨兵错误
var ErrNotFound = errors.New("NOTFOUND")

// Hello 测试
func Hello() ([]hello.TaskConfig, error) {
	var result []hello.TaskConfig
	/* Access Database */
	err := sql.ErrNoRows
	if err != nil {
		return nil, errors.Wrap(ErrNotFound, "找不到记录")
	}
	return result, nil
}
