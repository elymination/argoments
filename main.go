package argoments

import (
	"fmt"
	"os"
	"strings"
)

type Args struct {
	argsCount     int
	args          []string
	paramlessArgs []string
	paramedArgs   map[string]string
	parsedIndices []int
}

func Init() *Args {
	a := new(Args)
	a.paramlessArgs = make([]string, 0)
	a.paramedArgs = make(map[string]string, 0)
	a.argsCount = len(os.Args)
	a.args = make([]string, 0)
	return a
}

func (a *Args) RegisterParamed(paramedArgs []string) {
	for _, arg := range paramedArgs {
		arg = trimDashed(arg)
		a.paramedArgs[arg] = ""
	}
}
func (a *Args) Parse() {
	for index, arg := range os.Args {
		a.args = append(a.args, arg)
		if !a.isParsed(index) {
			a.parseArg(index, arg)
		}
	}
}

func (a *Args) GetValue(arg string) (string, error) {
	if value, found := a.paramedArgs[arg]; found {
		return value, nil
	}
	return "", fmt.Errorf("could not find %s parameter value", arg)
}

func (a *Args) GetArgs() []string {
	return a.args
}

func (a *Args) GetParamlessArgs() []string {
	return a.paramlessArgs
}

func (a *Args) GetRegisteredParamedArgs() []string {
	registered := []string{}
	for arg := range a.paramedArgs {
		registered = append(registered, arg)
	}
	return registered
}

func (a *Args) GetUsedParamedArgs() []string {
	registeredAndUsed := []string{}
	for _, arg := range a.GetRegisteredParamedArgs() {
		if a.IsUsed(arg) {
			registeredAndUsed = append(registeredAndUsed, arg)
		}
	}
	return registeredAndUsed
}

func (a *Args) GetUnusedParamedArgs() []string {
	registeredAndUnused := []string{}
	for _, arg := range a.GetRegisteredParamedArgs() {
		if !a.IsUsed(arg) {
			registeredAndUnused = append(registeredAndUnused, arg)
		}
	}
	return registeredAndUnused
}

func (a *Args) IsUsed(arg string) bool {
	arg = trimDashed(arg)
	if a, found := a.paramedArgs[arg]; found {
		return a != ""
	}
	return false
}

func (a *Args) parseArg(index int, arg string) {
	isDashed := isDashed(arg)
	arg = trimDashed(arg)
	if isDashed && a.isRegistered(arg) {
		if index+1 < a.argsCount {
			a.addParam([]string{arg, os.Args[index+1]})
			a.parsedIndices = append(a.parsedIndices, index, index+1)
		}
	} else {
		a.addParamless(arg)
	}
}

func (a *Args) addParamless(arg string) {
	a.paramlessArgs = append(a.paramlessArgs, arg)
}

func (a *Args) addParam(args []string) {
	a.paramedArgs[args[0]] = args[1]
}

func (a *Args) isRegistered(arg string) bool {
	arg = trimDashed(arg)
	_, found := a.paramedArgs[arg]
	return found
}

func (a *Args) isParsed(index int) bool {
	lastParsed := len(a.parsedIndices) - 1
	if lastParsed >= 0 {
		return index <= a.parsedIndices[lastParsed]
	}
	return false
}

func trimDashed(s string) string {
	for isDashed(s) {
		s = strings.TrimPrefix(s, "-")
	}
	return s
}

func isDashed(s string) bool {
	return strings.HasPrefix(s, "-")
}
