package logUtil

func LogErr(row LogStruct) string {
	if row.ErrorMessage != "" {
		return "Error : " + row.ErrorMessage
	}
	return ""
}
