package orm

import (
	"code_tpl_go/util/errs"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func init() {
	var err error
	dbUrl := "xxxx:my_sercret@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetData() (chat *Chat, err error) {
	chatID := 1
	err = DB.Take(&chat, chatID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("get chat error, not exist, err: %v", err)
		return nil, errs.New(-1, "会话不存在")
	} else if err != nil {
		fmt.Printf("get chat error, err:%v", err)
		return nil, errs.New(-2, "数据库会话查询失败")
	}
	return
}

func GetDataList() (err error) {
	//拉取用户的全部会话
	userID := "111"
	var chats = make([]Chat, 0)
	err = DB.Find(&chats, "user_id = ?", userID).Error
	if err != nil {
		fmt.Printf("delete chat error,  err:%v", err)
		err = errs.New(-1, "数据库操作失败")
		return
	}
	return
}

func GetDataList1() (err error) {
	var chats = make([]Chat, 0)
	err = DB.Where(&Chat{UserID: "111", Count: 1}).Find(&chats).Error
	if err != nil {
		fmt.Printf("cannot get chat: %v, err: %v", "111", err)
		err = errs.New(-1, "上轮对话查询失败")
		return
	}
	return
}

func GetDataList2() (err error) {
	var chats = make([]Chat, 0)
	err = DB.Order("created_at desc").Find(&chats, "user_id = ?", "111").Error
	if err != nil {
		return
	}
	return
}

func DeleteChat(chatID int, userID string) (err error) {
	var (
		LockingClause = clause.Locking{Strength: "UPDATE"}
	)
	var chat Chat
	err = DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Clauses(LockingClause).Take(&chat, chatID).Error
		if err != nil {
			return err
		}

		if chat.UserID != userID {
			return errs.New(-1, "无操作权限")
		}

		return tx.Delete(&chat).Error
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("delete chat error, not exist,userID:%v, chatID:%v, err:%v", userID, chatID, err)
		err = errs.New(-2, "会话不存在")
		return
	} else if err != nil {
		fmt.Printf("delete chat error, userID:%v, chatID:%v, err:%v", userID, chatID, err)
		err = errs.New(-3, "数据库操作失败")
		return
	}
	return
}

func Modify(chatID int) (err error) {
	var chat Chat
	var (
		LockingClause = clause.Locking{Strength: "UPDATE"}
	)
	userID := "11"
	err = DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Clauses(LockingClause).Take(&chat, chatID).Error
		if err != nil {
			return err
		}

		if chat.UserID != userID {
			return errs.New(-1, "无操作权限")
		}

		chat.Name = "name"

		return tx.Save(&chat).Error
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errs.New(-1, "会话不存在")
		return
	}
	return
}
