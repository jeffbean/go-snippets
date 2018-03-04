package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseZNode(t *testing.T) {
	tests := []struct {
		msg           string
		uri           string
		want          *ZNode
		wantErr       error
		wantErrString string
	}{
		{
			msg:  "valid node",
			uri:  "zk://foo/bar/baz",
			want: &ZNode{Ensemble: "foo", Path: "/bar/baz"},
		},
		{
			msg:  "root node",
			uri:  "zk://foo/",
			want: &ZNode{Ensemble: "foo", Path: "/"},
		},
		{
			msg:           "missing slash",
			uri:           "zk://foo",
			wantErr:       errMissingPath,
			wantErrString: "invalid ZNode URI: missing znode path from uri",
		},
		{
			msg:           "empty uri",
			uri:           "",
			wantErr:       errEmptyURI,
			wantErrString: "invalid ZNode URI: empty uri",
		},
		{
			msg:           "invalid scheme",
			uri:           "http://foo/bar",
			wantErr:       errInvalidScheme,
			wantErrString: "invalid ZNode URI: scheme must be zk:",
		},
		{
			msg:           "ensemble empty",
			uri:           "zk:///foo",
			wantErr:       errEnsembleEmpty,
			wantErrString: "invalid ZNode URI: ensemble name is empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			got, err := ParseZNode(tt.uri)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				assert.EqualError(t, err, tt.wantErrString)
				return
			}

			require.NoError(t, err, "did not expect error from parsing znode")
			assert.Equal(t, tt.want, got)
		})
	}
}
