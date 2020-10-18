package tracks

type TrackDto struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Position  *uint  `json:"position,omitempty"`
}
