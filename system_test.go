package cas

import (
	"fmt"
	"testing"
)

func TestSystem(t *testing.T) {
	fmt.Println("Testing system")

	es := NewEvalState()

	CasAssertSame(t, es, "f", "Head[f[x]]")
	CasAssertSame(t, es, "Symbol", "Head[x]")
	CasAssertSame(t, es, "List", "Head[{x}]")
	CasAssertSame(t, es, "Plus", "Head[a + b]")
	CasAssertSame(t, es, "Integer", "Head[1]")
	CasAssertSame(t, es, "Real", "Head[1.]")
	CasAssertSame(t, es, "Rational", "Head[2/7]")
	//CasAssertSame(t, es, "Rational", "Head[1/7]")
	CasAssertSame(t, es, "String", "Head[\"1\"]")
	CasAssertSame(t, es, "Plus", "Head[Head[(a + b)[x]]]")

	// Test speed of CommutativeIsMatchQ
	CasAssertSame(t, es, "Null", "Plus[foo, -foo, rest___] := bar + rest")
	CasAssertSame(t, es, "bar + 1 + a + b + c + d + e", "Plus[foo,1,-foo,a,b,c,d,e]")
}
