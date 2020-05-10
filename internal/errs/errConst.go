package errs

type ErrCode struct{ Code, Desc string }

var (
	// get config
	E10_01 = ErrCode{"10.01", "Unable to get config"}
)
