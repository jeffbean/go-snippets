package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseZNode(t *testing.T) {
	tests := []struct {
		msg     string
		uri     string
		want    *ZNode
		wantErr error
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
			msg:     "missing slash",
			uri:     "zk://foo",
			wantErr: errMissingPath,
		},
		{
			msg:     "empty uri",
			uri:     "",
			wantErr: errEmptyURI,
		},
		{
			msg:     "invalid scheme",
			uri:     "http://foo/bar",
			wantErr: errInvalidScheme,
		},
		{
			msg:     "ensemble empty",
			uri:     "zk:///foo",
			wantErr: errEnsembleEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			got, err := ParseZNode(tt.uri)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}

			require.NoError(t, err, "did not expect error from parsing znode")
			assert.Equal(t, tt.want, got)
		})
	}
}
