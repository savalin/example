package internal

type Edge struct {
	ID     int64   `json:"id"`
	From   int64   `json:"from"`
	To     int64   `json:"to"`
	Weight float64 `json:"weight"`
}
