package logUtil

import (
	"elkFormatter/constant"
	"elkFormatter/util"
)

func EventLog(row LogStruct) string {
	if row.ErrorMessage != "" {
		return constant.Red + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Red + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Red + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Red + row.Path + " : " + row.ToPath + "\n" +
			constant.Red + LogErr(row) + constant.Reset
	} else {
		return constant.Yellow + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Yellow + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Yellow + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Yellow + row.Path + " : " + row.ToPath + "\n" +
			constant.Yellow + LogErr(row)
	}
}
