package entity

type Property struct {
	ID       string  `json:"id" db:"id"`
	Location string  `json:"location" db:"location"`
	Lng      float64 `json:"lng" db:"lng"`
	Lat      float64 `json:"lat" db:"lat"`
	Price    float64 `json:"price" db:"price"`
}

type PropertyFilter struct {
	ID       string
	Location string
	MinLng   *ComparableFilter[float64]
	MinLat   *ComparableFilter[float64]
	MaxLat   *ComparableFilter[float64]
	Price    *ComparableFilter[float64]
	MaxLng   *ComparableFilter[float64]
}
