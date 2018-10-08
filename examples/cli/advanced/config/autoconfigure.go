package config

import (
	"github.com/hidevopsio/hiboot/examples/cli/advanced/model"
	"github.com/hidevopsio/hiboot/pkg/app"
)

type configuration struct {
	app.Configuration
}

func init() {
	app.AutoConfiguration(newConfiguration)
}

func newConfiguration() *configuration {
	return new(configuration)
}

func (c *configuration) Foo() *model.Foo {
	return &model.Foo{Name: "foo"}
}

func (c *configuration) FooBar() *model.Foo {
	return &model.Foo{Name: "foobar"}
}
