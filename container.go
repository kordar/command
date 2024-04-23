package command

import log "github.com/kordar/gologger"

var container = map[string]FuncCli{}

func RegisterCliFunc(cipher FuncCli) {
	container[cipher.Name()] = cipher
}

func GetCliFunc(name string) FuncCli {
	return container[name]
}

func ExecuteCliFunc(name string) {
	cliFunc := GetCliFunc(name)
	if cliFunc == nil {
		log.Errorf("待执行的程序[%s]不存在", name)
		return
	}
	args := cliFunc.GetArgs()
	cliFunc.Execute(args...)
}
