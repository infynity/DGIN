package model

import "errors"

func Txn() error{
	var tx = GetDB().Begin()
	//defer tx.Commit()

	var ask Ask
	err := tx.Where("`id` = ?", 1).First(&ask).Error
	if err != nil {
		return err
	}

	tx.Model(&ask).Create(&Ask{Content: "qqqqqq"})
	return errors.New("123")


}