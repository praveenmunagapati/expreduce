package expreduce

import "bytes"

type ToStringParams struct {
	form         string
	context      *String
	contextPath  *Expression
	previousHead string
}

func needsParens(thisHead string, previousHead string) bool {
	if previousHead == "<TOPLEVEL>" {
		return false
	}
	// TODO: look up actual precedence values using interface with the
	// interpreter.
	if previousHead == "System`Plus" && thisHead == "System`Times" {
		return false
	}
	if previousHead == "System`Rule" && thisHead == "System`Times" {
		return false
	}
	if previousHead == "System`RuleDelayed" && thisHead == "System`Times" {
		return false
	}
	if previousHead == "System`Rule" && thisHead == "System`Plus" {
		return false
	}
	if previousHead == "System`RuleDelayed" && thisHead == "System`Plus" {
		return false
	}
	return true
}

func ToStringInfix(parts []Ex, delim string, thisHead string, p ToStringParams) (bool, string) {
	if p.form != "InputForm" && p.form != "OutputForm" {
		return false, ""
	}
	if len(parts) < 2 {
		return false, ""
	}
	addParens := needsParens(thisHead, p.previousHead)
	var buffer bytes.Buffer
	if addParens {
		buffer.WriteString("(")
	}
	nextParams := ToStringParams{
		form:         p.form,
		context:      p.context,
		contextPath:  p.contextPath,
		previousHead: thisHead,
	}
	for i := 0; i < len(parts); i++ {
		buffer.WriteString(parts[i].StringForm(nextParams))
		if i != len(parts)-1 {
			buffer.WriteString(delim)
		}
	}
	if addParens {
		buffer.WriteString(")")
	}
	return true, buffer.String()
}

func (this *Expression) ToStringInfix(p ToStringParams) (bool, string) {
	if len(this.Parts) != 3 {
		return false, ""
	}
	expr, isExpr := this.Parts[1].(*Expression)
	delim, delimIsStr := this.Parts[2].(*String)
	if !isExpr || !delimIsStr {
		return false, ""
	}
	return ToStringInfix(expr.Parts[1:], delim.Val, "", p)
}

func ToStringInfixAdvanced(parts []Ex, delim string, thisHead string, surroundEachArg bool, start string, end string, params ToStringParams) (bool, string) {
	if params.form != "InputForm" && params.form != "OutputForm" {
		return false, ""
	}
	if len(parts) < 2 {
		return false, ""
	}
	var buffer bytes.Buffer
	if !surroundEachArg {
		buffer.WriteString(start)
	}
	nextParams := ToStringParams{
		form:         params.form,
		context:      params.context,
		contextPath:  params.contextPath,
		previousHead: thisHead,
	}
	for i := 0; i < len(parts); i++ {
		if surroundEachArg {
			buffer.WriteString("(")
			buffer.WriteString(parts[i].StringForm(nextParams))
			buffer.WriteString(")")
		} else {
			buffer.WriteString(parts[i].StringForm(nextParams))
		}
		if i != len(parts)-1 {
			buffer.WriteString(delim)
		}
	}
	if !surroundEachArg {
		buffer.WriteString(end)
	}
	return true, buffer.String()
}

func DefaultStringFormArgs() (*String, *Expression) {
	return NewString("Global`"), NewExpression([]Ex{
		NewSymbol("System`List"),
		NewString("System`"),
	})
}

func DefinitionComplexityStringFormArgs() (*String, *Expression) {
	// This was created because the "Private`" names in the blanks were
	// indicating greater complexity than they deserved.
	return NewString("Global`"), NewExpression([]Ex{
		NewSymbol("System`List"),
		NewString("System`"),
		NewString("Private`"),
	})
}

func ActualStringFormArgs(es *EvalState) (*String, *Expression) {
	return NewString(es.GetStringDef("System`$Context", "Global`")), es.GetListDef("System`$ContextPath")
}

func ActualStringFormArgsFull(form string, es *EvalState) ToStringParams {
	return ToStringParams{
		form:         form,
		context:      NewString(es.GetStringDef("System`$Context", "Global`")),
		contextPath:  es.GetListDef("System`$ContextPath"),
		previousHead: "<TOPLEVEL>",
	}

}
