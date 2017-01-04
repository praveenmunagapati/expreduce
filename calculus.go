package cas

func GetCalculusDefinitions() (defs []Definition) {
	defs = append(defs, Definition{
		name: "D",
		rules: map[string]string{
			"D[x_,x_]": "1",
			"D[a_,x_]": "0",
			"D[a_+b__,x_]": "D[a,x]+D[Plus[b],x]",
			"D[a_ b__,x_]": "D[a,x] b+a D[Times[b],x]",
			// The times operator is needed here. Whitespace precedence is messed up
			"D[a_^(b_), x_]": "a^b*(D[b,x] Log[a]+D[a,x]/a*b)",
			"D[Log[a_], x_]": "D[a, x]/a",
			"D[Sin[a_], x_]": "D[a,x] Cos[a]",
			"D[Cos[a_], x_]": "-D[a,x] Sin[a]",
		},
	})
	defs = append(defs, Definition{
		name: "Integrate",
		rules: map[string]string{
			// Might need to be implemented in code. Try running Integrate[-10x, {x, 1, 5}]
			// with this
			//"Integrate[a_,{x_Symbol,start_Integer,end_Integer}]": "ReplaceAll[Integrate[a, x],x->end] - ReplaceAll[Integrate[a, x],x->start]",
			"Integrate[a_Integer,x_Symbol]": "a*x",
			"Integrate[a_Integer*b_,x_Symbol]": "a*Integrate[b,x]",
			// An outstanding bug is requiring me to write this as amatch and bmatch
			// instead of a and b, because doing the latter causes issues with
			// Integrate[a+b+c,x]
			"Integrate[amatch_+bmatch__,x_Symbol]": "Integrate[amatch,x]+Integrate[Plus[bmatch],x]",
			"Integrate[x_Symbol^n_Integer, x_Symbol]": "x^(n+1)/(n+1)",
			"Integrate[x_Symbol^n_Rational, x_Symbol]": "x^(n+1)/(n+1)",
		},
	})
	return
}
