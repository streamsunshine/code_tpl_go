package logic

import (
	"errors"
	"fmt"
	"strings"
)

type CommonAggregateRs struct{
	ItemName string `json:"item_name"`
	Total    int    `json:"total"`
}

func AggregateData(startDate int, endDate int) ([]*CommonAggregateRs, error) {
	dbTableName := "dbName" + "." + "tableName"

	itemNameFormat := ""
	itemNameFormat = "%Y%m%d"    // "%Y%u" ,"%Y%m"

	args := []interface{}{startDate, endDate}
	whereStr := "date >= ? and date <= ? and book_id = ?"

	itemName := fmt.Sprintf("FROM_UNIXTIME(unix_timestamp(CONVERT(date,char)),'%s') as item_name", itemNameFormat)
	itemStrList := []string{itemName, "sum(num) total"}
	itemStr := strings.Join(itemStrList, ",")

	//确保按照 date 聚合
	sql := fmt.Sprintf("select  %s from %s where %s group by item_name;", itemStr, dbTableName, whereStr)

	list := make([]*CommonAggregateRs, 0)
	if rows, err := DataSQLObj.Raw(sql, args...).Rows(); err != nil {
		return nil, errors.New("1005")
	} else {
		for rows.Next() {
			rs := &CommonAggregateRs{}
			err = rows.Scan(&rs.ItemName, &rs.Total)
				err = rows.Scan(&rs.ItemName, &rs.Total)
			if err != nil {
				return nil, errors.New("1005")
			}
			list = append(list, rs)
		}
	}
	return list, nil
}
