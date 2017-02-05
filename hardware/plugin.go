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

const (
	errFormat = "%s had invalid type: expected %s, got %s"
)

func OpenPlugin(path string) (hardware dcpu.Hardware, err error) {
	p, err := plugin.Open(path)
	if err != nil {
		return
	}

	var errs *multierror.Error

	execute, err := p.Lookup("Execute")
	errs = multierror.Append(errs, err)

	getID, err := p.Lookup("GetID")
	errs = multierror.Append(errs, err)

	getManufacturerID, err := p.Lookup("GetManufacturerID")
	errs = multierror.Append(errs, err)

	getVersion, err := p.Lookup("GetVersion")
	errs = multierror.Append(errs, err)

	handleHardwareInterrupt, err := p.Lookup("HandleHardwareInterrupt")
	errs = multierror.Append(errs, err)

	d := &device{}

	if execute, ok := execute.(func(*dcpu.DCPU)); ok {
		d.execute = execute
	} else {
		errs = multierror.Append(errs, fmt.Errorf(errFormat, "Execute", reflect.TypeOf(d.execute), reflect.TypeOf(execute)))
	}

	if getID, ok := getID.(func() uint32); ok {
		d.getID = getID
	} else {
		errs = multierror.Append(errs, fmt.Errorf(errFormat, "GetID", reflect.TypeOf(d.getID), reflect.TypeOf(execute)))
	}

	if getManufacturerID, ok := getManufacturerID.(func() uint32); ok {
		d.getManufacturerID = getManufacturerID
	} else {
		errs = multierror.Append(errs, fmt.Errorf(errFormat,
			"GetManufacturerID", reflect.TypeOf(d.getManufacturerID), reflect.TypeOf(getManufacturerID)))
	}

	if getVersion, ok := getVersion.(func() uint16); ok {
		d.getVersion = getVersion
	} else {
		errs = multierror.Append(errs, fmt.Errorf(
			errFormat, "GetVersion", reflect.TypeOf(d.getVersion), reflect.TypeOf(getVersion)))
	}

	if handleHardwareInterrupt, ok := handleHardwareInterrupt.(func(*dcpu.DCPU)); ok {
		d.handleHardwareInterrupt = handleHardwareInterrupt
	} else {
		errs = multierror.Append(errs, fmt.Errorf(errFormat,
			"HandleHardwareInterrupt", reflect.TypeOf(d.handleHardwareInterrupt), reflect.TypeOf(handleHardwareInterrupt)))
	}

	hardware = d
	err = errs.ErrorOrNil()
	return
}
