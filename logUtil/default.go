package logUtil

import (
	"elkFormatter/constant"
	"elkFormatter/util"
)

func DefaultLog(row LogStruct) string {
	if row.ErrorMessage != "" {
		return constant.Red + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Red + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Red + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Red + row.Topic + " : " + row.Path + "\n" +
			constant.Red + LogErr(row) + constant.Reset
	} else {
		return "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			"Module : " + row.Module + " | " + row.Type + "\n" +
			row.ServiceName + " | " + row.EntryModule + "\n" +
			row.Topic + " : " + row.Path + "\n" +
			LogErr(row)
	}
}
