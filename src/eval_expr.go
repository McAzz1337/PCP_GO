package main

import "fmt"

type Expr interface {
	eval() int
}

type ConstantExpr struct {
	value int
}

func (c ConstantExpr) eval() int {
	return c.value
}

type PlusExpr struct {
	left  Expr
	right Expr
}

func (p PlusExpr) eval() int {
	return p.left.eval() + p.right.eval()
}

type TimesExpr struct {
	left  Expr
	right Expr
}

func (t TimesExpr) eval() int {
	return t.left.eval() * t.right.eval()
}

type NegExpr struct {
	inner Expr
}

func (n NegExpr) eval() int {
	return -n.inner.eval()
}

func eval_expression() {
	expr := PlusExpr{ConstantExpr{1}, TimesExpr{ConstantExpr{2}, ConstantExpr{3}}}
	fmt.Println("Result of evaluating the expression:", expr.eval())
}

