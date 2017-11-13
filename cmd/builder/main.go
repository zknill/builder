package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zknill/builder"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{
		Name:        "builder",
		Usage:       "a java builder class code generator",
		Description: "generate code for a java builder class",
		Version:     "0.0.1",

		HelpName: "",
		Action:   requireFlags(build),
		Flags:    flags,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func build(c *cli.Context) error {
	var vars []builder.Variable

	vars = append(vars, parseSingleVars(c)...)
	vars = append(vars, parseListVars(c)...)

	className := c.String("class")
	return builder.Generate(os.Stdout, className, vars)
}

func parseSingleVars(c *cli.Context) []builder.Variable {
	var vars []builder.Variable

	for _, val := range c.StringSlice("string") {
		vars = append(vars, builder.Var(builder.String, val))
	}
	for _, val := range c.StringSlice("int") {
		vars = append(vars, builder.Var(builder.Int, val))
	}
	for _, val := range c.StringSlice("integer") {
		vars = append(vars, builder.Var(builder.Integer, val))
	}
	for _, val := range c.StringSlice("boolean") {
		vars = append(vars, builder.Var(builder.Bool, val))
	}
	for _, val := range c.StringSlice("Boolean") {
		vars = append(vars, builder.Var(builder.Bool, val))
	}
	for _, val := range c.StringSlice("custom") {
		t := builder.Custom(strings.Title(val))
		v := strings.ToLower(val)
		vars = append(vars, builder.Var(t, v))
	}

	return vars
}

func parseListVars(c *cli.Context) []builder.Variable {
	var vars []builder.Variable

	for _, val := range c.StringSlice("string-list") {
		vars = append(vars, builder.List(builder.String, val))
	}
	for _, val := range c.StringSlice("integer-list") {
		vars = append(vars, builder.List(builder.Integer, val))
	}
	for _, val := range c.StringSlice("int-array") {
		vars = append(vars, builder.Array(builder.Int, val))
	}
	for _, val := range c.StringSlice("boolean-array") {
		vars = append(vars, builder.Array(builder.Bool, val))
	}
	for _, val := range c.StringSlice("Boolean-list") {
		vars = append(vars, builder.List(builder.Boolean, val))
	}
	for _, val := range c.StringSlice("custom-list") {
		t := builder.Custom(strings.Title(val))
		v := strings.ToLower(val)
		if !strings.HasSuffix(v, "s") {
			v += "s"
		}
		vars = append(vars, builder.List(t, v))
	}

	return vars
}

func requireFlags(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}

		return f(c)
	}
}

var flags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    "string",
		Aliases: []string{"s"},
		Usage:   "string variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "int",
		Aliases: []string{"i"},
		Usage:   "int variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "integer",
		Aliases: []string{"I"},
		Usage:   "Integer variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "boolean",
		Aliases: []string{"b"},
		Usage:   "boolean variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "Boolean",
		Aliases: []string{"B"},
		Usage:   "Boolean variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:  "custom",
		Usage: "define a variable with a custom `type`, the var name will be lowercase custom name",
	},
	&cli.StringSliceFlag{
		Name:    "string-list",
		Aliases: []string{"sl"},
		Usage:   "string list variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "integer-list",
		Aliases: []string{"IL"},
		Usage:   "Integer list variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "int-array",
		Aliases: []string{"ia"},
		Usage:   "int array variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "boolean-array",
		Aliases: []string{"ba"},
		Usage:   "boolean array variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:    "Boolean-list",
		Aliases: []string{"BL"},
		Usage:   "Boolean list variable by `name`",
	},
	&cli.StringSliceFlag{
		Name:  "custom-list",
		Usage: "variable with a custom `type`, the var name will be lowercase custom name",
	},
	&cli.StringFlag{
		Name:    "class",
		Aliases: []string{"c"},
		Usage:   "class to generate builder for",
		Value:   "CLAZZ",
	},
}
