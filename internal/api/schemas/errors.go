package schemas

type ApiErr struct {
	Code    int
	Message string
}

func (apiErr *ApiErr) Error() string {
	return apiErr.Message
}
