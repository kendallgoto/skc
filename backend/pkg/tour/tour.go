package tour

type TourItem struct {
	Title     string
	ShortName string
	Body      string
	Code      string
}

var TourItems = []TourItem{}

func init() {
	// populate touritems in order ...
	TourItems = append(TourItems, buildBasics()...)
}
