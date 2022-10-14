package utils
import (
	"time"
)
type SampleResponse struct{
Sum float64  `json:"sum"`	
Avg float64  `json:"avg"`
Max float64  `json:"max"`
Min float64  `json:"min"`
Count int    `json:"count"`
}
type SampleInput struct{
Amount float64  `json:"amount"`
Timestamp time.Time `json:"timestamp"`
}

// Author model (Struct)
type Author struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}

// Course model (Struct)
type Course struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Price  string  `json:"price"`
    Link   string  `json:"link"`
    Author *Author `json:"author"`
}

