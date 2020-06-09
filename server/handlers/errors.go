package handlers

// Errors in string format useful to handle in Serve function
const (
	EMissingData            = "missing data in body"
	EUnreadableBodyFmt      = "can't read body: %s"
	EUnexpectedBodyFmt      = "can't unmarshal %s from body: %s"
	ERequest                = "Error returned on request to api: %s"
	EUnexpectedStatus       = "unexpected status returned from api: %s"
	EUnexpectedEmptyMessage = "unexpected empty message returned"
)
