package tracks

type TrackDto struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Position  *int   `json:"position,omitempty"`
}
