package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
	"log"
	"strconv"

	"gopl.io/ch7/eval"
)

func main() {
	s, err := input("Input expression:")
	if err != nil {
		log.Fatal(err)
	}

	expr, vars, err := parseAndCheck(s)
	if err != nil {
		log.Fatal(err)
	}

	env, err := buildEnv(vars)
	if err != nil {
		log.Fatal(err)
	}

	res := expr.Eval(env)
	fmt.Printf("Result: %g", res)
}

func input(prompt string) (string, error){
	fmt.Println(prompt)

	r := bufio.NewReader(os.Stdin)
	t, err := r.ReadString('\n')
	if err == io.EOF{
		return "", fmt.Errorf("meet EOF")
	}
	if err != nil{
		return "", fmt.Errorf("input expr has failed: %v", err)
	}

	return t, nil
}

func parseAndCheck(s string) (eval.Expr, map[eval.Var]bool, error) {
	if s == "" {
		return nil, nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, nil, err
	}

	return expr, vars, nil
}

func buildEnv(vars map[eval.Var]bool) (eval.Env, error) {
	env := eval.Env{}

	for v := range vars {
		k, err := input(string(v) + ":")
		if err != nil {
			return nil, fmt.Errorf("input var %s: %v", v, err)
		}

		k = k[:len(k) - 1]	// Trim last character which is equal to the delimiter '\n'.
		env[v], err = strconv.ParseFloat(k, 64)
		if err != nil {
			return nil, fmt.Errorf("parse %s as float: %v", v, err)
		}
	}

	return env, nil
}
