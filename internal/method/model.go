package method

type ActiveState int

const (
	GET ActiveState = iota
	POST
	PUT
	PATCH
	DELETE
)

func (s ActiveState) String() string {
	switch s {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	default:
		return "Unknown"
	}
}

type MethodModel struct {
	ActiveState   ActiveState
	OverlayActive bool
	windowWidth   int
	windowHeight  int
}

func New() MethodModel {

	return MethodModel{OverlayActive: false, ActiveState: GET}
}
