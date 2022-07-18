package utils

func GetError(err error) string {
	if err != nil {
		return ""
	}
	return err.Error()
}
