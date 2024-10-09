package endpoint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSetMeta(t *testing.T) {
	getUserPermissions := Set{}.GetPermissionsDefine().Get("User").Get("GetUsers")
	require.Len(t, getUserPermissions, 1)
	require.True(t, getUserPermissions[0].EnableAuth)
	require.Equal(t, getUserPermissions[0].Role, []string{"admin", "viewer"})
	require.Equal(t, getUserPermissions[0].Description, "Get user list")
}
