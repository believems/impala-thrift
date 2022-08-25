// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package partitions

import (
  "bytes"
  "context"
  "database/sql/driver"
  "errors"
  "fmt"
  "github.com/apache/thrift/lib/go/thrift"
  "github.com/believems/impala-thrift/exprs"
  "reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = exprs.GoUnusedProtection__

type TPartitionType int64

const (
	TPartitionType_UNPARTITIONED     TPartitionType = 0
	TPartitionType_RANDOM            TPartitionType = 1
	TPartitionType_HASH_PARTITIONED  TPartitionType = 2
	TPartitionType_RANGE_PARTITIONED TPartitionType = 3
	TPartitionType_KUDU              TPartitionType = 4
)

func (p TPartitionType) String() string {
	switch p {
	case TPartitionType_UNPARTITIONED:
		return "UNPARTITIONED"
	case TPartitionType_RANDOM:
		return "RANDOM"
	case TPartitionType_HASH_PARTITIONED:
		return "HASH_PARTITIONED"
	case TPartitionType_RANGE_PARTITIONED:
		return "RANGE_PARTITIONED"
	case TPartitionType_KUDU:
		return "KUDU"
	}
	return "<UNSET>"
}

func TPartitionTypeFromString(s string) (TPartitionType, error) {
	switch s {
	case "UNPARTITIONED":
		return TPartitionType_UNPARTITIONED, nil
	case "RANDOM":
		return TPartitionType_RANDOM, nil
	case "HASH_PARTITIONED":
		return TPartitionType_HASH_PARTITIONED, nil
	case "RANGE_PARTITIONED":
		return TPartitionType_RANGE_PARTITIONED, nil
	case "KUDU":
		return TPartitionType_KUDU, nil
	}
	return TPartitionType(0), fmt.Errorf("not a valid TPartitionType string")
}

func TPartitionTypePtr(v TPartitionType) *TPartitionType { return &v }

func (p TPartitionType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TPartitionType) UnmarshalText(text []byte) error {
	q, err := TPartitionTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *TPartitionType) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = TPartitionType(v)
	return nil
}

func (p *TPartitionType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Type
//  - PartitionExprs
type TDataPartition struct {
	Type           TPartitionType `thrift:"type,1,required" db:"type" json:"type"`
	PartitionExprs []*exprs.TExpr `thrift:"partition_exprs,2" db:"partition_exprs" json:"partition_exprs,omitempty"`
}

func NewTDataPartition() *TDataPartition {
	return &TDataPartition{}
}

func (p *TDataPartition) GetType() TPartitionType {
	return p.Type
}

var TDataPartition_PartitionExprs_DEFAULT []*exprs.TExpr

func (p *TDataPartition) GetPartitionExprs() []*exprs.TExpr {
	return p.PartitionExprs
}
func (p *TDataPartition) IsSetPartitionExprs() bool {
	return p.PartitionExprs != nil
}

func (p *TDataPartition) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetType bool = false

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
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
				issetType = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField2(iprot); err != nil {
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
	if !issetType {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Type is not set"))
    }
	return nil
}

func (p *TDataPartition) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := TPartitionType(v)
		p.Type = temp
	}
	return nil
}

func (p *TDataPartition) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*exprs.TExpr, 0, size)
	p.PartitionExprs = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &exprs.TExpr{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.PartitionExprs = append(p.PartitionExprs, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TDataPartition) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TDataPartition"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
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

func (p *TDataPartition) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("type", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err)
	}
	return err
}

func (p *TDataPartition) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetPartitionExprs() {
		if err := oprot.WriteFieldBegin("partition_exprs", thrift.LIST, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:partition_exprs: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.PartitionExprs)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.PartitionExprs {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:partition_exprs: ", p), err)
		}
	}
	return err
}

func (p *TDataPartition) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TDataPartition(%+v)", *p)
}
