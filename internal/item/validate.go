package item

import (
	"midterm-api/internal/constant"
	"strings"

	"github.com/go-faster/errors"
)

type Validate struct {

}

func NewValidate() Validate {
	return Validate{}
}

func (validator Validate) UpdateItem(status constant.ItemStatus) error {
	if status == constant.ItemApprovedStatus || status == constant.ItemRejectedStatus {
		return errors.Errorf("Cannot update item when status is %s", strings.ToLower(string(status)))
	}
	return nil
}