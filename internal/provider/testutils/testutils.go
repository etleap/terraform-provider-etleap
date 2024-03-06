package testutils

import (
	"bytes"
	"text/template"
)

type TestConstants struct {
	MysqlConnectionId    string
	RedshiftConnectionId string
	PipelineId           string
}

var Constants = &TestConstants{
	MysqlConnectionId:    `A5p7oGO4`,
	RedshiftConnectionId: `hvgjRKgo`,
	PipelineId:           `qqkU5JdA`,
}

func RunTemplate[T any](template *template.Template, config T) string {
	var byteBuffer bytes.Buffer
	_ = template.Execute(&byteBuffer, config)
	return byteBuffer.String()
}
