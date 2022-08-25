// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package logging

import(
	"bytes"
	"context"
	"reflect"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type TLogLevel int64
const (
  TLogLevel_VLOG_3 TLogLevel = 0
  TLogLevel_VLOG_2 TLogLevel = 1
  TLogLevel_VLOG TLogLevel = 2
  TLogLevel_INFO TLogLevel = 3
  TLogLevel_WARN TLogLevel = 4
  TLogLevel_ERROR TLogLevel = 5
  TLogLevel_FATAL TLogLevel = 6
)

func (p TLogLevel) String() string {
  switch p {
  case TLogLevel_VLOG_3: return "VLOG_3"
  case TLogLevel_VLOG_2: return "VLOG_2"
  case TLogLevel_VLOG: return "VLOG"
  case TLogLevel_INFO: return "INFO"
  case TLogLevel_WARN: return "WARN"
  case TLogLevel_ERROR: return "ERROR"
  case TLogLevel_FATAL: return "FATAL"
  }
  return "<UNSET>"
}

func TLogLevelFromString(s string) (TLogLevel, error) {
  switch s {
  case "VLOG_3": return TLogLevel_VLOG_3, nil 
  case "VLOG_2": return TLogLevel_VLOG_2, nil 
  case "VLOG": return TLogLevel_VLOG, nil 
  case "INFO": return TLogLevel_INFO, nil 
  case "WARN": return TLogLevel_WARN, nil 
  case "ERROR": return TLogLevel_ERROR, nil 
  case "FATAL": return TLogLevel_FATAL, nil 
  }
  return TLogLevel(0), fmt.Errorf("not a valid TLogLevel string")
}


func TLogLevelPtr(v TLogLevel) *TLogLevel { return &v }

func (p TLogLevel) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *TLogLevel) UnmarshalText(text []byte) error {
q, err := TLogLevelFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *TLogLevel) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = TLogLevel(v)
return nil
}

func (p * TLogLevel) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - ClassName
//  - LogLevel
type TSetJavaLogLevelParams struct {
  ClassName string `thrift:"class_name,1,required" db:"class_name" json:"class_name"`
  LogLevel string `thrift:"log_level,2,required" db:"log_level" json:"log_level"`
}

func NewTSetJavaLogLevelParams() *TSetJavaLogLevelParams {
  return &TSetJavaLogLevelParams{}
}


func (p *TSetJavaLogLevelParams) GetClassName() string {
  return p.ClassName
}

func (p *TSetJavaLogLevelParams) GetLogLevel() string {
  return p.LogLevel
}
func (p *TSetJavaLogLevelParams) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetClassName bool = false;
  var issetLogLevel bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
        issetClassName = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
        issetLogLevel = true
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
  if !issetClassName{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ClassName is not set"));
  }
  if !issetLogLevel{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field LogLevel is not set"));
  }
  return nil
}

func (p *TSetJavaLogLevelParams)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ClassName = v
}
  return nil
}

func (p *TSetJavaLogLevelParams)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.LogLevel = v
}
  return nil
}

func (p *TSetJavaLogLevelParams) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TSetJavaLogLevelParams"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSetJavaLogLevelParams) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("class_name", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:class_name: ", p), err) }
  if err := oprot.WriteString(string(p.ClassName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.class_name (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:class_name: ", p), err) }
  return err
}

func (p *TSetJavaLogLevelParams) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("log_level", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:log_level: ", p), err) }
  if err := oprot.WriteString(string(p.LogLevel)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.log_level (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:log_level: ", p), err) }
  return err
}

func (p *TSetJavaLogLevelParams) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSetJavaLogLevelParams(%+v)", *p)
}

// Attributes:
//  - LogLevels
type TGetJavaLogLevelsResult_ struct {
  LogLevels []string `thrift:"log_levels,1,required" db:"log_levels" json:"log_levels"`
}

func NewTGetJavaLogLevelsResult_() *TGetJavaLogLevelsResult_ {
  return &TGetJavaLogLevelsResult_{}
}


func (p *TGetJavaLogLevelsResult_) GetLogLevels() []string {
  return p.LogLevels
}
func (p *TGetJavaLogLevelsResult_) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetLogLevels bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
        issetLogLevels = true
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
  if !issetLogLevels{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field LogLevels is not set"));
  }
  return nil
}

func (p *TGetJavaLogLevelsResult_)  ReadField1(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.LogLevels =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.LogLevels = append(p.LogLevels, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TGetJavaLogLevelsResult_) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TGetJavaLogLevelsResult"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TGetJavaLogLevelsResult_) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("log_levels", thrift.LIST, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:log_levels: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRING, len(p.LogLevels)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.LogLevels {
    if err := oprot.WriteString(string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:log_levels: ", p), err) }
  return err
}

func (p *TGetJavaLogLevelsResult_) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TGetJavaLogLevelsResult_(%+v)", *p)
}

