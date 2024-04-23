package command

import log "github.com/kordar/gologger"

// VersionFunc 查询软件版本
type VersionFunc struct {
}

func (v VersionFunc) Name() string {
	// TODO implement me
	return "version"
}

func (v VersionFunc) Execute(args ...interface{}) {
	// TODO implement me
	log.Infof("当前软件版本：%s", "v0.0.1")
}

func (v VersionFunc) GetArgs() []interface{} {
	// TODO implement me
	return []interface{}{}
}
