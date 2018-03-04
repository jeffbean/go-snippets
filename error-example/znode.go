package example

import "strings"

// this should be "zk" but for the demo this makes it a little bit easier to show
const _zkScheme = "zk:"

// ZNode represents a znode in zookeeper
type ZNode struct {
	Ensemble string
	Path     string
}

// ParseZNode expects fully quialified uris for a znode
// eg. "zk://zk-cluster.example.com/foo/bar"
func ParseZNode(uri string) (*ZNode, error) {
	// TODO: hacked up parsing to demo error types
	if uri == "" {
		return nil, errEmptyURI
	}

	parts := strings.Split(uri, _zkScheme)
	if len(parts) < 2 {
		return nil, errInvalidScheme
	}
	// large assumptions
	rest := parts[1]

	if strings.HasPrefix(rest, "///") {
		return nil, errEnsembleEmpty
	}
	// more large assumptions here
	hostParts := strings.SplitN(rest[2:], "/", 2)
	if len(hostParts) < 2 {
		return nil, errMissingPath
	}

	// more assumptions for demo on appending slash
	return &ZNode{Ensemble: hostParts[0], Path: "/" + hostParts[1]}, nil
}
