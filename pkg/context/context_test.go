package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetParentNamespace(t *testing.T) {
	parentName := "root-ctx"
	rootCtx := NewContext(parentName)
	serverCtx := NewContext(rootCtx, "server-ctx")

	debugNamespace := serverCtx.getParent().GetDebugNamespace()

	assert.Contains(t, debugNamespace, parentName)
	assert.Equal(t, len(debugNamespace), len(parentName)+11)
}
