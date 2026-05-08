package layouts

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultHead_ContainsAndroidAppConfig(t *testing.T) {
	var out bytes.Buffer
	err := DefaultHead().Render(context.Background(), &out)
	require.NoError(t, err)

	head := out.String()
	require.Contains(t, head, `rel="manifest"`)
	require.Contains(t, head, `.webmanifest"`)
	require.Contains(t, head, `name="theme-color" content="#111827"`)
	require.Contains(t, head, `name="mobile-web-app-capable" content="yes"`)
	require.True(t, strings.Contains(head, `rel="icon"`))
}
