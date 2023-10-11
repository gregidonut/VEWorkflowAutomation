package appInterface

type AppInterface interface {
	CpUploadFileProgressPercentage(int64) int64
	Debug(string)
	Info(string)
	Warning(string)
}
