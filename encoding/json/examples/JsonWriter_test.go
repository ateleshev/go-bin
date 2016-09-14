package json_test

import "testing"
import "github.com/ateleshev/go-bin/bytes"

// import "github.com/ateleshev/go-bin/strings"
import "github.com/ateleshev/go-bin/encoding/json"
import sjson "encoding/json"

/**
 * STDOUT example:
 *
 *  import "os"
 *  import "bufio"
 *
 *  out := bufio.NewWriter(os.Stdout)
 *  jout := json.NewJsonWriter(out)
 *  obj.JsonWriteTo(jout)
 *  if jout.HasErrors() {
 *    log.Println(jout.LastError())
 *  }
 *  jout.Release()
 */

var (
	JsonAttrTestObjectUrl             = "Url"
	JsonAttrTestObjectTitle           = "Title"
	JsonAttrTestObjectDescription     = "Description"
	JsonAttrTestObjectNumber          = "Number"
	JsonAttrTestObjectPrice           = "Price"
	JsonAttrTestObjectDiscount        = "Discount"
	JsonAttrTestObjectCurrency        = "Currency"
	JsonAttrTestObjectMonthsOfTheYear = "MonthsOfTheYear"
	JsonAttrTestObjectNumbersName     = "NumbersName"
	JsonAttrTestObjectAges            = "Ages"
	JsonAttrTestObjectSubObject       = "SubObject"
	JsonAttrTestObjectIsDone          = "IsDone"
	JsonAttrTestObjectSubObject2      = "SubObject2"

	JsonAttrTestSubObjectName       = "Name"
	JsonAttrTestSubObjectIntValue   = "IntValue"
	JsonAttrTestSubObjectInt8Value  = "Int8Value"
	JsonAttrTestSubObjectInt16Value = "Int16Value"
	JsonAttrTestSubObjectInt32Value = "Int32Value"
	JsonAttrTestSubObjectIsDone     = "IsDone"
)

type TestSubObject struct {
	Name       string
	IntValue   int
	Int8Value  int8
	Int16Value int16
	Int32Value int32
	IsDone     bool
}

func (this *TestSubObject) JsonWriteTo(jw *json.JsonWriter) *json.JsonWriter {
	jw.ObjOpen()

	// --[ Slower ]--
	//	jw.ObjElm(&JsonAttrTestSubObjectName, &this.Name, true)
	//	jw.ObjElm(&JsonAttrTestSubObjectIntValue, &this.IntValue, false)
	//	jw.ObjElm(&JsonAttrTestSubObjectInt8Value, &this.Int8Value, false)
	//	jw.ObjElm(&JsonAttrTestSubObjectInt16Value, &this.Int16Value, false)
	//	jw.ObjElm(&JsonAttrTestSubObjectInt32Value, &this.Int32Value, false)
	//	jw.ObjElm(&JsonAttrTestSubObjectIsDone, &this.IsDone, false)

	jw.StringValue(JsonAttrTestSubObjectName).Sep().StringValue(this.Name)
	jw.Next().StringValue(JsonAttrTestSubObjectIntValue).Sep().IntValue(this.IntValue)
	jw.Next().StringValue(JsonAttrTestSubObjectInt8Value).Sep().Int8Value(this.Int8Value)
	jw.Next().StringValue(JsonAttrTestSubObjectInt16Value).Sep().Int16Value(this.Int16Value)
	jw.Next().StringValue(JsonAttrTestSubObjectInt32Value).Sep().Int32Value(this.Int32Value)
	jw.Next().StringValue(JsonAttrTestSubObjectIsDone).Sep().BoolValue(this.IsDone)

	jw.ObjClose()

	return jw
}

type TestObject struct {
	Url             string
	Title           string
	Description     string
	Number          int
	Price           float64
	Discount        float32
	Currency        string
	MonthsOfTheYear map[string]int `json:"-"`
	NumbersName     map[int]string `json:"-"`
	Ages            []int8
	SubObject       *TestSubObject
	IsDone          bool
	SubObject2      *TestSubObject
}

var obj = &TestObject{
	Url:         "http://golang.org/dl/",
	Title:       "GOlang",
	Description: "Test: [ \", /, \\, \b, \f, \n, \r, \t ]. Тест Української Мови ... [ OK ]",
	Number:      1239874560,
	Price:       202.65,
	Discount:    -7.33,
	Currency:    "EUR",
	MonthsOfTheYear: map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	},
	NumbersName: map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	},
	Ages: []int8{
		25,
		30,
		35,
		40,
		45,
	},
	SubObject: &TestSubObject{
		Name:       "Some TestSubObject",
		IntValue:   123456798,
		Int8Value:  120,
		Int16Value: 1024,
		Int32Value: 600000,
		IsDone:     true,
	},
	SubObject2: nil,
}

func (this *TestObject) JsonWriteTo(jw *json.JsonWriter) *json.JsonWriter {
	jw.ObjOpen()

	// --[ Slower ]--
	//	jw.ObjElm(&JsonAttrTestObjectUrl, &this.Url, true)
	//	jw.ObjElm(&JsonAttrTestObjectTitle, &this.Title, false)
	//	jw.ObjElm(&JsonAttrTestObjectDescription, &this.Description, false)
	//	jw.ObjElm(&JsonAttrTestObjectNumber, &this.Number, false)
	//	jw.ObjElm(&JsonAttrTestObjectPrice, &this.Price, false)
	//	jw.ObjElm(&JsonAttrTestObjectDiscount, &this.Discount, false)
	//	jw.ObjElm(&JsonAttrTestObjectCurrency, &this.Currency, false)

	jw.StringValue(JsonAttrTestObjectUrl).Sep().StringValue(this.Url)
	jw.Next().StringValue(JsonAttrTestObjectTitle).Sep().StringValue(this.Title)
	jw.Next().StringValue(JsonAttrTestObjectDescription).Sep().StringValue(this.Description)
	jw.Next().StringValue(JsonAttrTestObjectNumber).Sep().IntValue(this.Number)
	jw.Next().StringValue(JsonAttrTestObjectPrice).Sep().Float64Value(this.Price)
	jw.Next().StringValue(JsonAttrTestObjectDiscount).Sep().Float32Value(this.Discount)
	jw.Next().StringValue(JsonAttrTestObjectCurrency).Sep().StringValue(this.Currency)

	//  >>> Does not support in standard library <<<

	// --[ Slower ]--
	//  var i int
	//	// -- MonthsOfTheYear
	//	jw.ObjElmName(&JsonAttrTestObjectMonthsOfTheYear, false)
	//	jw.ObjOpen()
	//	i = 0
	//	for v, k := range this.MonthsOfTheYear {
	//		jw.ObjElm(&k, &v, i == 0)
	//		i++
	//	}
	//	jw.ObjClose()
	//	// -- /MonthsOfTheYear
	//
	//	// -- NumbersName
	//	jw.ObjElmName(&JsonAttrTestObjectNumbersName, false)
	//	jw.ObjOpen()
	//	i = 0
	//	for v, k := range this.MonthsOfTheYear {
	//		jw.ObjElm(&k, &v, i == 0)
	//		i++
	//	}
	//	jw.ObjClose()
	//	// -- /NumbersName

	// -- Ages

	// --[ Slower ]--
	//	jw.ObjElmName(&JsonAttrTestObjectAges, false)
	//	jw.ArrOpen()
	//	for i, v := range this.Ages {
	//		jw.ArrElm(&v, i == 0)
	//	}
	//	jw.ArrClose()

	jw.Next().StringValue(JsonAttrTestObjectAges).Sep().ArrOpen()
	for i, v := range this.Ages {
		if i > 0 {
			jw.Next()
		}
		jw.Int8Value(v)
	}
	jw.ArrClose()

	// -- /Ages

	// --[ Slower ]--
	//	jw.ObjElm(&JsonAttrTestObjectSubObject, this.SubObject, false)
	//	jw.ObjElm(&JsonAttrTestObjectIsDone, &this.IsDone, false)
	//	jw.ObjElm(&JsonAttrTestObjectSubObject2, this.SubObject2, false)

	jw.Next().StringValue(JsonAttrTestObjectSubObject).Sep()
	if this.SubObject == nil {
		jw.NullValue()
	} else {
		this.SubObject.JsonWriteTo(jw)
	}
	jw.Next().StringValue(JsonAttrTestObjectIsDone).Sep().BoolValue(this.IsDone)

	jw.Next().StringValue(JsonAttrTestObjectSubObject2).Sep()
	if this.SubObject2 == nil {
		jw.NullValue()
	} else {
		this.SubObject2.JsonWriteTo(jw)
	}

	jw.ObjClose()

	return jw
}

var buf = bytes.NewBuffer(3 * 1024)
var jw = json.NewJsonWriter(buf)

/**
 * ==[ Tests ]==
 *
 * go test -v -run=JsonWriter_
 */

func TestJsonWriter_Object(t *testing.T) { // {{{
	if obj.JsonWriteTo(jw).HasErrors() {
		t.Fatal(jw.LastError())
	}
	// t.Log(buf.Bytes())
	t.Log(buf.String())
	jw.Release()
} // }}}

func TestStandart_Object(t *testing.T) { // {{{
	b, err := sjson.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(b)
	t.Log(string(b))
} // }}}

/**
 * ==[ Benchmarks ]==
 *
 * go test -v -run=^$ -benchmem -bench=JsonWriter_
 *
 * go test -v -run=^$ -benchmem -bench=JsonWriter_ -memprofile=mem.out -benchtime=3s | tee mem.profile
 * go test -v -run=^$ -benchmem -bench=JsonWriter_ -cpuprofile=cpu.out -benchtime=3s | tee cpu.profile
 */

func BenchmarkJsonWriter_Object(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		obj.JsonWriteTo(jw)
		if jw.HasErrors() {
			b.Fatal(jw.LastError())
		}
		buf.Reset()
	}
} // }}}

func BenchmarkStandart_Object(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		if bs, err := sjson.Marshal(obj); err != nil {
			b.Fatal(err)
		} else {
			_ = bs
		}
	}
} // }}}
