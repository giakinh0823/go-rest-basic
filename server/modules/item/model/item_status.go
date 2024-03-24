package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	DOING ItemStatus = iota
	DONE
	DELETED
)

var ITEM_STATUS = [3]string{"DOING", "DONE", "DELETED"}

func (item *ItemStatus) String() string {
	return ITEM_STATUS[*item]
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	val, err := parseStrToItemStatus(str)

	if err != nil {
		return err
	}

	*item = val

	return nil
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to  scan data from sql: %s", value))
	}

	val, err := parseStrToItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to  scan data from sql: %s", value))
	}

	*item = val

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

func parseStrToItemStatus(s string) (ItemStatus, error) {
	for i := range ITEM_STATUS {
		if ITEM_STATUS[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New(fmt.Sprintf("Invalid  status string"))
}

