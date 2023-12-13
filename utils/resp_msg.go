package utils

func ErrorRequired(param string) string {
	return param + " cannot empty. Please input " + param
}

func ErrorFailedReadRequest() string {
	return "Failed to read request data"
}

func ErrorFailedExecData(action, param string) string {
	return "Failed to " + action + " " + param + " data"
}

func SuccessExecData(action, param string) string {
	return "Success to " + action + " " + param + " data"
}
