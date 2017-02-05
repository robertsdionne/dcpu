package hardware

import (
	"fmt"
	"plugin"
	"reflect"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/robertsdionne/dcpu"
)

type device struct {
	execute                 func(*dcpu.DCPU)
	getID                   func() uint32
	getManufacturerID       func() uint32
	getVersion              func() uint16
	handleHardwareInterrupt func(*dcpu.DCPU)
}

func (d *device) Execute(dcpu *dcpu.DCPU) {
	d.execute(dcpu)
}

func (d *device) GetID() uint32 {
	return d.getID()
}

func (d *device) GetManufacturerID() uint32 {
	return d.getManufacturerID()
}

func (d *device) GetVersion() uint16 {
	return d.getVersion()
}

func (d *device) HandleHardwareInterrupt(dcpu *dcpu.DCPU) {
	d.handleHardwareInterrupt(dcpu)
}

func OpenPlugin(path string) (hardware dcpu.Hardware, err error) {
	p, err := plugin.Open(path)
	if err != nil {
		return
	}

	execute, e := p.Lookup("Execute")
	err = multierror.Append(err, e)

	getID, e := p.Lookup("GetID")
	err = multierror.Append(err, e)

	getManufacturerID, e := p.Lookup("GetManufacturerID")
	err = multierror.Append(err, e)

	getVersion, e := p.Lookup("GetVersion")
	err = multierror.Append(err, e)

	handleHardwareInterrupt, e := p.Lookup("HandleHardwareInterrupt")
	err = multierror.Append(err, e)

	d := &device{}

	if execute, ok := execute.(func(*dcpu.DCPU)); ok {
		d.execute = execute
	} else {
		err = multierror.Append(err, fmt.Errorf(
			"Execute had invalid type: expected %s, got %s", reflect.TypeOf(d.execute), reflect.TypeOf(execute)))
	}

	if getID, ok := getID.(func() uint32); ok {
		d.getID = getID
	} else {
		err = multierror.Append(
			err, fmt.Errorf("GetID had invalid type: expected %s, got %s", reflect.TypeOf(d.getID), reflect.TypeOf(execute)))
	}

	if getManufacturerID, ok := getManufacturerID.(func() uint32); ok {
		d.getManufacturerID = getManufacturerID
	} else {
		err = multierror.Append(err, fmt.Errorf("GetManufacturerID had invalid type: expected %s, got %s",
			reflect.TypeOf(d.getManufacturerID), reflect.TypeOf(getManufacturerID)))
	}

	if getVersion, ok := getVersion.(func() uint16); ok {
		d.getVersion = getVersion
	} else {
		err = multierror.Append(err, fmt.Errorf(
			"GetVersion had invalid type: expected %s, got %s", reflect.TypeOf(d.getVersion), reflect.TypeOf(getVersion)))
	}

	if handleHardwareInterrupt, ok := handleHardwareInterrupt.(func(*dcpu.DCPU)); ok {
		d.handleHardwareInterrupt = handleHardwareInterrupt
	} else {
		err = multierror.Append(err, fmt.Errorf("HandleHardwareInterrupt had invalid type: expected %s, got %s",
			reflect.TypeOf(d.handleHardwareInterrupt), reflect.TypeOf(handleHardwareInterrupt)))
	}

	hardware = d
	return
}
