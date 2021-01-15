package context

import (
	"fmt"

	"github.com/loopicorp/loopgo/internal/utils"
)

type Context struct {
	name   string
	parent *Context
}

func NewContext(parent interface{}, ctxName ...string) *Context {
	c := new(Context)
	c.setParentName(parent, ctxName...)

	return c
}

func (c *Context) GetDebugNamespace() string {
	return c.name
}

func (c *Context) getParent() *Context {
	return c.parent
}

func (c *Context) setParentName(parent interface{}, ctxName ...string) {
	if len(ctxName) > 1 {
		panic("cannot take anything aside a single name")
	}
	if p, ok := parent.(string); ok {
		c.name = p
		c.parent = nil
	} else {
		c.parent = parent.(*Context)
	}
	if len(ctxName) > 0 {
		c.name = ctxName[0]
	} else {
		c.name = c.generateName()
	}
}

func (c *Context) generateName() string {
	name := "loopgo:context"
	if len(c.name) > 0 {
		name = c.name
	}
	return fmt.Sprintf("%s:%s", name, utils.GenerateUniqueID())
}
