package translator

import (
	"fmt"
	"strings"
)

type CFunctionSpec struct {
	Typedef   string
	Return    CType
	Params    []*CDecl
	Arrays    string
	VarArrays uint8
	Pointers  uint8
}

func (c *CFunctionSpec) AddArray(size uint64) {
	if size > 0 {
		c.Arrays = fmt.Sprintf("%s[%d]", c.Arrays, size)
		return
	}
	c.VarArrays++
}

func (c CFunctionSpec) String() string {
	var params []string
	for i, param := range c.Params {
		if len(param.Name) == 0 {
			params = append(params, fmt.Sprintf("arg%d", i))
			continue
		}
		params = append(params, param.Name)
	}
	paramList := strings.Join(params, ", ")
	if c.Return != nil {
		return fmt.Sprintf("%s (%s)", c.Return, paramList)
	}
	return fmt.Sprintf("void (%s)", paramList)
}

func (c *CFunctionSpec) SetPointers(n uint8) {
	c.Pointers = n
}

func (c *CFunctionSpec) Kind() CTypeKind {
	return FunctionKind
}

func (c *CFunctionSpec) IsComplete() bool {
	return true
}

func (c *CFunctionSpec) IsOpaque() bool {
	return len(c.Params) == 0
}

func (c CFunctionSpec) Copy() CType {
	return &c
}

func (c *CFunctionSpec) GetBase() string {
	return ""
}

func (c *CFunctionSpec) GetTag() string {
	return ""
}

func (c *CFunctionSpec) CGoName() string {
	return ""
}

func (c *CFunctionSpec) GetArrays() string {
	return c.Arrays
}

func (c *CFunctionSpec) GetVarArrays() uint8 {
	return c.VarArrays
}

func (c *CFunctionSpec) GetPointers() uint8 {
	return c.Pointers
}

func (c *CFunctionSpec) IsConst() bool {
	return false
}
