package main

import (
	"fmt"

	"gopkg.in/lxc/go-lxc.v2"
)

func newcontainer(name string) (*lxc.Container, error) {
	c, err := lxc.NewContainer(name, lxc.DegaultConfigPath())
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Create(name string) error {
	c, err := newcontainer(name)
	if err != nil {
		return err
	}
	defer c.Release()

	createtemplate := lxc.TemplateOptions {
		Distro:		"debian",
		Release:	"buster",
		Arch:			"amd64",
	}

	err = c.Create(createtemplate)
	if err != nil {
		return err
	}
	return nil
}

func Delete(name string) error {
	c, err := newcontainer(name)
	if err != nil {
		return err
	}
	defer c.Release()

	err = c.Destroy()
	if err != nil {
		return err
	}
}
