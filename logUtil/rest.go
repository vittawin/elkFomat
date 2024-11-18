package logUtil

import (
	"elkFormatter/constant"
	"elkFormatter/util"
)

func RestLog(row LogStruct) string {
	if row.OriginalPath == "" {
		row.OriginalPath = row.EventName
	}

	//errorMessage, ok := row.ErrorMessage.([]ErrorStruct)
	if row.ErrorMessage != "" {
		//fmt.Println("Error : ", row.ErrorMessage)
		return constant.Red + "Time : " + util.FormatTimeToString(row.Timestamp) + "\n" +
			constant.Red + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Red + restFormat(row) +
			constant.Red + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Red + row.Method + " : " + restPath(row) + "\n" +
			constant.Red + LogErr(row) + constant.Reset
	} else {
		return constant.Blue + "Module : " + row.Module + " | " + row.Type + "\n" +
			constant.Blue + restFormat(row) +
			constant.Blue + row.ServiceName + " | " + row.EntryModule + "\n" +
			constant.Blue + row.Method + " : " + restPath(row) + "\n" +
			constant.Blue + LogErr(row) + constant.Reset
	}
}

func restFormat(row LogStruct) string {
	if row.Module == "client" {
		return "ToPath : " + row.ToPath + "\n"
	}
	return ""
}

func restPath(row LogStruct) string {
	if row.OriginalPath != "" {
		return row.OriginalPath
	}
	return row.Path
}
