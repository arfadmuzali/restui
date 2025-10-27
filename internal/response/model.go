package response

type ResponseModel struct {
	ResponseWidth  int
	ResponseHeight int
	result         []byte
}

func New() ResponseModel {
	return ResponseModel{}
}
