package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	InternalServerError = Err("internal_server_error")
	BadRequest          = Err("bad_request")
)
