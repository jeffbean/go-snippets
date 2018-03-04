package example

import "fmt"

type errZNodeInvalid struct {
	reason string
}

func (e errZNodeInvalid) Error() string {
	return "invalid ZNode URI: " + e.reason
}

var (
	errEmptyURI      = errZNodeInvalid{"empty uri"}
	errInvalidScheme = errZNodeInvalid{fmt.Sprintf("path scheme must be %v", _zkScheme)}
	errEnsembleEmpty = errZNodeInvalid{"ensemble name is empty"}
	errMissingPath   = errZNodeInvalid{"missing znode path from uri"}
)
