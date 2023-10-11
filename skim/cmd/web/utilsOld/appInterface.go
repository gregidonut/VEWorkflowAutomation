package utilsOld

type AppInterface interface {
	Debug(string, ...string)
	Info(string, ...string)
	Warning(string, ...string)
}
