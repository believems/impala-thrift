// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package metricdefs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/believems/impala-thrift/metrics"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

var _ = metrics.GoUnusedProtection__

// Attributes:
//   - Key
//   - Kind
//   - Units
//   - Contexts
//   - Label
//   - Description
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

func (p *TMetricDef) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TMetricDef) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = &v
	}
	return nil
}

func (p *TMetricDef) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := metrics.TMetricKind(v)
		p.Kind = &temp
	}
	return nil
}

func (p *TMetricDef) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := metrics.TUnit(v)
		p.Units = &temp
	}
	return nil
}

func (p *TMetricDef) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Contexts = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Contexts = append(p.Contexts, _elem0)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TMetricDef) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Label = &v
	}
	return nil
}

func (p *TMetricDef) ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Description = &v
	}
	return nil
}

func (p *TMetricDef) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "TMetricDef"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField6(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TMetricDef) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetKey() {
		if err := oprot.WriteFieldBegin(ctx, "key", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Key)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetKind() {
		if err := oprot.WriteFieldBegin(ctx, "kind", thrift.I32, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:kind: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.Kind)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.kind (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:kind: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetUnits() {
		if err := oprot.WriteFieldBegin(ctx, "units", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:units: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.Units)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.units (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:units: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetContexts() {
		if err := oprot.WriteFieldBegin(ctx, "contexts", thrift.LIST, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:contexts: ", p), err)
		}
		if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.Contexts)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Contexts {
			if err := oprot.WriteString(ctx, string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(ctx); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:contexts: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetLabel() {
		if err := oprot.WriteFieldBegin(ctx, "label", thrift.STRING, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:label: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Label)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.label (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:label: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetDescription() {
		if err := oprot.WriteFieldBegin(ctx, "description", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:description: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Description)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.description (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:description: ", p), err)
		}
	}
	return err
}

func (p *TMetricDef) Equals(other *TMetricDef) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Key != other.Key {
		if p.Key == nil || other.Key == nil {
			return false
		}
		if (*p.Key) != (*other.Key) {
			return false
		}
	}
	if p.Kind != other.Kind {
		if p.Kind == nil || other.Kind == nil {
			return false
		}
		if (*p.Kind) != (*other.Kind) {
			return false
		}
	}
	if p.Units != other.Units {
		if p.Units == nil || other.Units == nil {
			return false
		}
		if (*p.Units) != (*other.Units) {
			return false
		}
	}
	if len(p.Contexts) != len(other.Contexts) {
		return false
	}
	for i, _tgt := range p.Contexts {
		_src1 := other.Contexts[i]
		if _tgt != _src1 {
			return false
		}
	}
	if p.Label != other.Label {
		if p.Label == nil || other.Label == nil {
			return false
		}
		if (*p.Label) != (*other.Label) {
			return false
		}
	}
	if p.Description != other.Description {
		if p.Description == nil || other.Description == nil {
			return false
		}
		if (*p.Description) != (*other.Description) {
			return false
		}
	}
	return true
}

func (p *TMetricDef) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TMetricDef(%+v)", *p)
}
