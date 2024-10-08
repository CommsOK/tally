// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package v1

import (
	"bytes"
	"fmt"
	"github.com/CommsOK/tally/v4/thirdparty/github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Different types of values that m3 emits. Each metric
// must contain one of these values
//
// Attributes:
//   - Count
//   - Gauge
//   - Timer
type MetricValue struct {
	Count *CountValue `thrift:"count,1" json:"count,omitempty"`
	Gauge *GaugeValue `thrift:"gauge,2" json:"gauge,omitempty"`
	Timer *TimerValue `thrift:"timer,3" json:"timer,omitempty"`
}

func NewMetricValue() *MetricValue {
	return &MetricValue{}
}

var MetricValue_Count_DEFAULT CountValue

func (p *MetricValue) GetCount() CountValue {
	if !p.IsSetCount() {
		return MetricValue_Count_DEFAULT
	}
	return *p.Count
}

var MetricValue_Gauge_DEFAULT GaugeValue

func (p *MetricValue) GetGauge() GaugeValue {
	if !p.IsSetGauge() {
		return MetricValue_Gauge_DEFAULT
	}
	return *p.Gauge
}

var MetricValue_Timer_DEFAULT TimerValue

func (p *MetricValue) GetTimer() TimerValue {
	if !p.IsSetTimer() {
		return MetricValue_Timer_DEFAULT
	}
	return *p.Timer
}
func (p *MetricValue) CountSetFieldsMetricValue() int {
	count := 0
	if p.IsSetCount() {
		count++
	}
	if p.IsSetGauge() {
		count++
	}
	if p.IsSetTimer() {
		count++
	}
	return count

}

func (p *MetricValue) IsSetCount() bool {
	return p.Count != nil
}

func (p *MetricValue) IsSetGauge() bool {
	return p.Gauge != nil
}

func (p *MetricValue) IsSetTimer() bool {
	return p.Timer != nil
}

func (p *MetricValue) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
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

func (p *MetricValue) readField1(iprot thrift.TProtocol) error {
	p.Count = &CountValue{}
	if err := p.Count.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Count), err)
	}
	return nil
}

func (p *MetricValue) readField2(iprot thrift.TProtocol) error {
	p.Gauge = &GaugeValue{}
	if err := p.Gauge.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Gauge), err)
	}
	return nil
}

func (p *MetricValue) readField3(iprot thrift.TProtocol) error {
	p.Timer = &TimerValue{}
	if err := p.Timer.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Timer), err)
	}
	return nil
}

func (p *MetricValue) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsMetricValue(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("MetricValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MetricValue) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetCount() {
		if err := oprot.WriteFieldBegin("count", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:count: ", p), err)
		}
		if err := p.Count.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Count), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:count: ", p), err)
		}
	}
	return err
}

func (p *MetricValue) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetGauge() {
		if err := oprot.WriteFieldBegin("gauge", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:gauge: ", p), err)
		}
		if err := p.Gauge.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Gauge), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:gauge: ", p), err)
		}
	}
	return err
}

func (p *MetricValue) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimer() {
		if err := oprot.WriteFieldBegin("timer", thrift.STRUCT, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timer: ", p), err)
		}
		if err := p.Timer.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Timer), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timer: ", p), err)
		}
	}
	return err
}

func (p *MetricValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MetricValue(%+v)", *p)
}

// Different types of count values
//
// Attributes:
//   - I64Value
type CountValue struct {
	I64Value *int64 `thrift:"i64Value,1" json:"i64Value,omitempty"`
}

func NewCountValue() *CountValue {
	return &CountValue{}
}

var CountValue_I64Value_DEFAULT int64

func (p *CountValue) GetI64Value() int64 {
	if !p.IsSetI64Value() {
		return CountValue_I64Value_DEFAULT
	}
	return *p.I64Value
}
func (p *CountValue) CountSetFieldsCountValue() int {
	count := 0
	if p.IsSetI64Value() {
		count++
	}
	return count

}

func (p *CountValue) IsSetI64Value() bool {
	return p.I64Value != nil
}

func (p *CountValue) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
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

func (p *CountValue) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.I64Value = &v
	}
	return nil
}

func (p *CountValue) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsCountValue(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("CountValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CountValue) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetI64Value() {
		if err := oprot.WriteFieldBegin("i64Value", thrift.I64, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:i64Value: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.I64Value)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.i64Value (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:i64Value: ", p), err)
		}
	}
	return err
}

func (p *CountValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CountValue(%+v)", *p)
}

// Different types of gauge values
//
// Attributes:
//   - I64Value
//   - DValue
type GaugeValue struct {
	I64Value *int64   `thrift:"i64Value,1" json:"i64Value,omitempty"`
	DValue   *float64 `thrift:"dValue,2" json:"dValue,omitempty"`
}

func NewGaugeValue() *GaugeValue {
	return &GaugeValue{}
}

var GaugeValue_I64Value_DEFAULT int64

func (p *GaugeValue) GetI64Value() int64 {
	if !p.IsSetI64Value() {
		return GaugeValue_I64Value_DEFAULT
	}
	return *p.I64Value
}

var GaugeValue_DValue_DEFAULT float64

func (p *GaugeValue) GetDValue() float64 {
	if !p.IsSetDValue() {
		return GaugeValue_DValue_DEFAULT
	}
	return *p.DValue
}
func (p *GaugeValue) CountSetFieldsGaugeValue() int {
	count := 0
	if p.IsSetI64Value() {
		count++
	}
	if p.IsSetDValue() {
		count++
	}
	return count

}

func (p *GaugeValue) IsSetI64Value() bool {
	return p.I64Value != nil
}

func (p *GaugeValue) IsSetDValue() bool {
	return p.DValue != nil
}

func (p *GaugeValue) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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

func (p *GaugeValue) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.I64Value = &v
	}
	return nil
}

func (p *GaugeValue) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DValue = &v
	}
	return nil
}

func (p *GaugeValue) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsGaugeValue(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("GaugeValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GaugeValue) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetI64Value() {
		if err := oprot.WriteFieldBegin("i64Value", thrift.I64, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:i64Value: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.I64Value)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.i64Value (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:i64Value: ", p), err)
		}
	}
	return err
}

func (p *GaugeValue) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDValue() {
		if err := oprot.WriteFieldBegin("dValue", thrift.DOUBLE, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:dValue: ", p), err)
		}
		if err := oprot.WriteDouble(float64(*p.DValue)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.dValue (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:dValue: ", p), err)
		}
	}
	return err
}

func (p *GaugeValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GaugeValue(%+v)", *p)
}

// Different types of timer values
//
// Attributes:
//   - I64Value
//   - DValue
type TimerValue struct {
	I64Value *int64   `thrift:"i64Value,1" json:"i64Value,omitempty"`
	DValue   *float64 `thrift:"dValue,2" json:"dValue,omitempty"`
}

func NewTimerValue() *TimerValue {
	return &TimerValue{}
}

var TimerValue_I64Value_DEFAULT int64

func (p *TimerValue) GetI64Value() int64 {
	if !p.IsSetI64Value() {
		return TimerValue_I64Value_DEFAULT
	}
	return *p.I64Value
}

var TimerValue_DValue_DEFAULT float64

func (p *TimerValue) GetDValue() float64 {
	if !p.IsSetDValue() {
		return TimerValue_DValue_DEFAULT
	}
	return *p.DValue
}
func (p *TimerValue) CountSetFieldsTimerValue() int {
	count := 0
	if p.IsSetI64Value() {
		count++
	}
	if p.IsSetDValue() {
		count++
	}
	return count

}

func (p *TimerValue) IsSetI64Value() bool {
	return p.I64Value != nil
}

func (p *TimerValue) IsSetDValue() bool {
	return p.DValue != nil
}

func (p *TimerValue) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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

func (p *TimerValue) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.I64Value = &v
	}
	return nil
}

func (p *TimerValue) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DValue = &v
	}
	return nil
}

func (p *TimerValue) Write(oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsTimerValue(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set).", p, c)
	}
	if err := oprot.WriteStructBegin("TimerValue"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TimerValue) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetI64Value() {
		if err := oprot.WriteFieldBegin("i64Value", thrift.I64, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:i64Value: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.I64Value)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.i64Value (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:i64Value: ", p), err)
		}
	}
	return err
}

func (p *TimerValue) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetDValue() {
		if err := oprot.WriteFieldBegin("dValue", thrift.DOUBLE, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:dValue: ", p), err)
		}
		if err := oprot.WriteDouble(float64(*p.DValue)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.dValue (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:dValue: ", p), err)
		}
	}
	return err
}

func (p *TimerValue) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TimerValue(%+v)", *p)
}

// Tags that can be applied to a metric
//
// Attributes:
//   - TagName
//   - TagValue
type MetricTag struct {
	TagName  string  `thrift:"tagName,1" json:"tagName"`
	TagValue *string `thrift:"tagValue,2" json:"tagValue,omitempty"`
}

func NewMetricTag() *MetricTag {
	return &MetricTag{}
}

func (p *MetricTag) GetTagName() string {
	return p.TagName
}

var MetricTag_TagValue_DEFAULT string

func (p *MetricTag) GetTagValue() string {
	if !p.IsSetTagValue() {
		return MetricTag_TagValue_DEFAULT
	}
	return *p.TagValue
}
func (p *MetricTag) IsSetTagValue() bool {
	return p.TagValue != nil
}

func (p *MetricTag) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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

func (p *MetricTag) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.TagName = v
	}
	return nil
}

func (p *MetricTag) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.TagValue = &v
	}
	return nil
}

func (p *MetricTag) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MetricTag"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MetricTag) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tagName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:tagName: ", p), err)
	}
	if err := oprot.WriteString(string(p.TagName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tagName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:tagName: ", p), err)
	}
	return err
}

func (p *MetricTag) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetTagValue() {
		if err := oprot.WriteFieldBegin("tagValue", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tagValue: ", p), err)
		}
		if err := oprot.WriteString(string(*p.TagValue)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.tagValue (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tagValue: ", p), err)
		}
	}
	return err
}

func (p *MetricTag) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MetricTag(%+v)", *p)
}

// The metric that is being emitted
//
// Attributes:
//   - Name
//   - MetricValue
//   - Timestamp
//   - Tags
type Metric struct {
	Name        string              `thrift:"name,1" json:"name"`
	MetricValue *MetricValue        `thrift:"metricValue,2" json:"metricValue,omitempty"`
	Timestamp   *int64              `thrift:"timestamp,3" json:"timestamp,omitempty"`
	Tags        map[*MetricTag]bool `thrift:"tags,4" json:"tags,omitempty"`
}

func NewMetric() *Metric {
	return &Metric{}
}

func (p *Metric) GetName() string {
	return p.Name
}

var Metric_MetricValue_DEFAULT *MetricValue

func (p *Metric) GetMetricValue() *MetricValue {
	if !p.IsSetMetricValue() {
		return Metric_MetricValue_DEFAULT
	}
	return p.MetricValue
}

var Metric_Timestamp_DEFAULT int64

func (p *Metric) GetTimestamp() int64 {
	if !p.IsSetTimestamp() {
		return Metric_Timestamp_DEFAULT
	}
	return *p.Timestamp
}

var Metric_Tags_DEFAULT map[*MetricTag]bool

func (p *Metric) GetTags() map[*MetricTag]bool {
	return p.Tags
}
func (p *Metric) IsSetMetricValue() bool {
	return p.MetricValue != nil
}

func (p *Metric) IsSetTimestamp() bool {
	return p.Timestamp != nil
}

func (p *Metric) IsSetTags() bool {
	return p.Tags != nil
}

func (p *Metric) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
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

func (p *Metric) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Metric) readField2(iprot thrift.TProtocol) error {
	p.MetricValue = &MetricValue{}
	if err := p.MetricValue.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.MetricValue), err)
	}
	return nil
}

func (p *Metric) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Timestamp = &v
	}
	return nil
}

func (p *Metric) readField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin()
	if err != nil {
		return thrift.PrependError("error reading set begin: ", err)
	}
	tSet := make(map[*MetricTag]bool, size)
	p.Tags = tSet
	for i := 0; i < size; i++ {
		_elem0 := &MetricTag{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Tags[_elem0] = true
	}
	if err := iprot.ReadSetEnd(); err != nil {
		return thrift.PrependError("error reading set end: ", err)
	}
	return nil
}

func (p *Metric) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Metric"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
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
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Metric) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *Metric) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetMetricValue() {
		if err := oprot.WriteFieldBegin("metricValue", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:metricValue: ", p), err)
		}
		if err := p.MetricValue.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.MetricValue), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:metricValue: ", p), err)
		}
	}
	return err
}

func (p *Metric) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamp() {
		if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:timestamp: ", p), err)
		}
		if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.timestamp (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:timestamp: ", p), err)
		}
	}
	return err
}

func (p *Metric) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetTags() {
		if err := oprot.WriteFieldBegin("tags", thrift.SET, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:tags: ", p), err)
		}
		if err := oprot.WriteSetBegin(thrift.STRUCT, len(p.Tags)); err != nil {
			return thrift.PrependError("error writing set begin: ", err)
		}
		for v, _ := range p.Tags {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteSetEnd(); err != nil {
			return thrift.PrependError("error writing set end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:tags: ", p), err)
		}
	}
	return err
}

func (p *Metric) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Metric(%+v)", *p)
}

// Structure that holds a group of metrics which share
// common properties like the cluster and service.
//
// Attributes:
//   - Metrics
//   - CommonTags
type MetricBatch struct {
	Metrics    []*Metric           `thrift:"metrics,1" json:"metrics"`
	CommonTags map[*MetricTag]bool `thrift:"commonTags,2" json:"commonTags,omitempty"`
}

func NewMetricBatch() *MetricBatch {
	return &MetricBatch{}
}

func (p *MetricBatch) GetMetrics() []*Metric {
	return p.Metrics
}

var MetricBatch_CommonTags_DEFAULT map[*MetricTag]bool

func (p *MetricBatch) GetCommonTags() map[*MetricTag]bool {
	return p.CommonTags
}
func (p *MetricBatch) IsSetCommonTags() bool {
	return p.CommonTags != nil
}

func (p *MetricBatch) Read(iprot thrift.TProtocol) error {
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
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
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

func (p *MetricBatch) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Metric, 0, size)
	p.Metrics = tSlice
	for i := 0; i < size; i++ {
		_elem1 := &Metric{}
		if err := _elem1.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem1), err)
		}
		p.Metrics = append(p.Metrics, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *MetricBatch) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin()
	if err != nil {
		return thrift.PrependError("error reading set begin: ", err)
	}
	tSet := make(map[*MetricTag]bool, size)
	p.CommonTags = tSet
	for i := 0; i < size; i++ {
		_elem2 := &MetricTag{}
		if err := _elem2.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
		}
		p.CommonTags[_elem2] = true
	}
	if err := iprot.ReadSetEnd(); err != nil {
		return thrift.PrependError("error reading set end: ", err)
	}
	return nil
}

func (p *MetricBatch) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("MetricBatch"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MetricBatch) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("metrics", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:metrics: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Metrics)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Metrics {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:metrics: ", p), err)
	}
	return err
}

func (p *MetricBatch) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetCommonTags() {
		if err := oprot.WriteFieldBegin("commonTags", thrift.SET, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:commonTags: ", p), err)
		}
		if err := oprot.WriteSetBegin(thrift.STRUCT, len(p.CommonTags)); err != nil {
			return thrift.PrependError("error writing set begin: ", err)
		}
		for v, _ := range p.CommonTags {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteSetEnd(); err != nil {
			return thrift.PrependError("error writing set end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:commonTags: ", p), err)
		}
	}
	return err
}

func (p *MetricBatch) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MetricBatch(%+v)", *p)
}
