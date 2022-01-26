package date_distance

type DateDistance struct {
	From string `json:"from" query:"from"`
	To   string `json:"to" query:"to"`
}
