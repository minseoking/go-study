package result

type Result[T any] struct {
	ErrorMessage string `json:"errorMessage"`
	Result       T      `json:"result"`
	Success      bool   `json:"success"`
}
