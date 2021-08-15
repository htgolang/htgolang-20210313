package responses

type JSONResponse struct {
	Ok     bool
	Reason string
	Result interface{}
}

func NewJsonSuccess() JSONResponse {
	return JSONResponse{Ok: true}
}

func NewJsonFailure(reason string) JSONResponse {
	return JSONResponse{Ok: false, Reason: reason}
}

func NewJsonResponse(result interface{}) JSONResponse {
	return JSONResponse{Ok: true, Result: result}
}
