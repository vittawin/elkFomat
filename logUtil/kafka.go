package logUtil

import (
	"elkFormatter/constant"
	"elkFormatter/util"
)

func KafkaLog(row LogStruct) string {
	if row.ErrorMessage != "" {
		return constant.Red + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Red + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Red + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Red + row.Topic + " : " + row.Path + "\n" +
			constant.Red + LogErr(row) + constant.Reset
	} else {
		return constant.Green + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Green + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Green + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Green + row.Topic + " : " + row.Path + "\n" +
			constant.Green + LogErr(row)
	}
}
