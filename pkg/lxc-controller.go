package pkg

import (
	"fmt"

	"gopkg.in/lxc/go-lxc.v2"
)

func newcontainer(name string) (*lxc.Container, error) {
	c, err := lxc.NewContainer(name, lxc.DefaultConfigPath())
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Create(obj interface{}) error {
	spec, ok := obj.(LxcSpec)
	if !ok {
		fmt.Println("spec struct unmatch")
	}
	fmt.Println(ok)
	c, err := newcontainer(spec.Name)
	if err != nil {
		return err
	}
	defer c.Release()

	fmt.Println(spec)

	createtemplate := lxc.TemplateOptions {
		Template:	"download",
		//Backend:	lxc.Directory,
		Distro:		spec.Distro,
		Release:	spec.Release,
		Arch:			spec.Arch,
		//Variant: "default",
		FlushCache: false,
		DisableGPGValidation: false,
	}

	fmt.Println("create")
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
	return nil
}
