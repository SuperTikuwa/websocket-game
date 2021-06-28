package utility

type Message struct {
	ID  int      `json:"id,omitempty"`
	Pos Position `json:"position,omitempty"`
}

type Position struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}
