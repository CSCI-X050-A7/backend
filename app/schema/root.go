package schema

// Pointer returns a pointer to a variable holding the supplied T constant
func Pointer[T any](x T) *T {
	return &x
}

type Pagination struct {
	Ordering *string `form:"ordering,omitempty" json:"ordering,omitempty"`
	Offset   *int    `form:"offset,omitempty"   json:"offset,omitempty"`
	Limit    *int    `form:"limit,omitempty"    json:"limit,omitempty"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Msg         string `json:"msg"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

type ListResponse[T any] struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Data   []T `json:"data"`
}
