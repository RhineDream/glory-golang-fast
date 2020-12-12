package systemdao

import (
	"github.com/astaxie/beego/orm"
	"glory-golang/models/system"
	"glory-golang/utils"
	"strconv"
	"strings"
)

type BaseDao struct {
}

func GetPageData(pageNo int64, pageSize int64, sql string, obj []system.SysUser) (int64, string) {
	o := orm.NewOrm()

	begin := strings.Index(sql, "from")
	subSql := utils.SubString(sql, begin, len(sql))

	var count int64
	countSql := "select count(1) " + subSql
	o.Raw(countSql).QueryRow(&count)

	pageStart := (pageNo - 1) * pageSize
	pageEnd := pageSize
	resultSql := sql + " limit " + strconv.FormatInt(pageStart, 10) + "," + strconv.FormatInt(pageEnd, 10)
	return count, resultSql
}
