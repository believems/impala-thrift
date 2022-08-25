// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package metricdefs

import (
  "bytes"
  "context"
  "fmt"
  "github.com/apache/thrift/lib/go/thrift"
  "github.com/believems/impala-thrift/metrics"
  "reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = metrics.GoUnusedProtection__

// Attributes:
//  - Key
//  - Kind
//  - Units
//  - Contexts
//  - Label
//  - Description
type TMetricDef struct {
	Key         *string              `thrift:"key,1" db:"key" json:"key,omitempty"`
	Kind        *metrics.TMetricKind `thrift:"kind,2" db:"kind" json:"kind,omitempty"`
	Units       *metrics.TUnit       `thrift:"units,3" db:"units" json:"units,omitempty"`
	Contexts    []string             `thrift:"contexts,4" db:"contexts" json:"contexts,omitempty"`
	Label       *string              `thrift:"label,5" db:"label" json:"label,omitempty"`
	Description *string              `thrift:"description,6" db:"description" json:"description,omitempty"`
}

func NewTMetricDef() *TMetricDef {
	return &TMetricDef{}
}

var TMetricDef_Key_DEFAULT string

func (p *TMetricDef) GetKey() string {
	if !p.IsSetKey() {
		return TMetricDef_Key_DEFAULT
	}
	return *p.Key
}

var TMetricDef_Kind_DEFAULT metrics.TMetricKind

func (p *TMetricDef) GetKind() metrics.TMetricKind {
	if !p.IsSetKind() {
		return TMetricDef_Kind_DEFAULT
	}
	return *p.Kind
}

var TMetricDef_Units_DEFAULT metrics.TUnit

func (p *TMetricDef) GetUnits() metrics.TUnit {
	if !p.IsSetUnits() {
		return TMetricDef_Units_DEFAULT
	}
	return *p.Units
}

var TMetricDef_Contexts_DEFAULT []string

func (p *TMetricDef) GetContexts() []string {
	return p.Contexts
}

var TMetricDef_Label_DEFAULT string

func (p *TMetricDef) GetLabel() string {
	if !p.IsSetLabel() {
		return TMetricDef_Label_DEFAULT
	}
	return *p.Label
}

var TMetricDef_Description_DEFAULT string

func (p *TMetricDef) GetDescription() string {
	if !p.IsSetDescription() {
		return TMetricDef_Description_DEFAULT
	}
	return *p.Description
}
func (p *TMetricDef) IsSetKey() bool {
	return p.Key != nil
}

func (p *TMetricDef) IsSetKind() bool {
	return p.Kind != nil
}

func (p *TMetricDef) IsSetUnits() bool {
	return p.Units != nil
}

func (p *TMetricDef) IsSetContexts() bool {
	return p.Contexts != nil
}

func (p *TMetricDef) IsSetLabel() bool {
	return p.Label != nil
}

func (p *TMetricDef) IsSetDescription() bool {
	return p.Description != nil
}

func (p *TMetricDef) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
        }
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TMetricDef) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = &v
	}
	return nil
}

func (p *TMetricDef) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := metrics.TMetricKind(v)
		p.Kind = &temp
	}
	return nil
}

func (p *TMetricDef) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := metrics.TUnit(v)
		p.Units = &temp
	}
	return nil
}

func (p *TMetricDef) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Contexts = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Contexts = append(p.Contexts, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TMetricDef) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Label = &v
	}
	return nil
}

func (p *TMetricDef) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Description = &v
	}
	return nil
}

func (p *TMetricDef) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TMetricDef"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TMetricDef) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetKey() {
		if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Key)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetKind() {
		if err := oprot.WriteFieldBegin("kind", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:kind: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Kind)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.kind (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:kind: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetUnits() {
		if err := oprot.WriteFieldBegin("units", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:units: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Units)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.units (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:units: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetContexts() {
		if err := oprot.WriteFieldBegin("contexts", thrift.LIST, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:contexts: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Contexts)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Contexts {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:contexts: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetLabel() {
		if err := oprot.WriteFieldBegin("label", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:label: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Label)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.label (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:label: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetDescription() {
		if err := oprot.WriteFieldBegin("description", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:description: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Description)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.description (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:description: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TMetricDef(%+v)", *p)
}
