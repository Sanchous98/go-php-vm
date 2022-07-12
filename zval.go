package vm

import (
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

var variablesPool = sync.Pool{New: func() any { return new(ZVariable) }}

type ZInt int

func (z ZInt) ToInt() ZInt                { return z }
func (z ZInt) ToDouble() ZDouble          { return ZDouble(z) }
func (z ZInt) ToBool() ZBool              { return z != 0 }
func (z ZInt) ToString() ZString          { return ZString(strconv.Itoa(int(z))) }
func (z ZInt) ToArray() ZArray            { return ZArray{List: []ZVal{z}} }
func (z ZInt) IsIdentical(val ZVal) ZBool { return z == val }
func (z ZInt) String() string             { return "int" }

func (z ZInt) Compare(val ZVal) ZInt {
	val = UnwrapVariable(val)

	switch val := val.(type) {
	case ZString, ZArray:
		return val.Compare(z)
	case ZDouble:
		if z.ToDouble() < val {
			return -1
		} else if z.ToDouble() > val {
			return 1
		} else {
			return 0
		}
	default:
		if z < val.ToInt() {
			return -1
		} else if z > val.ToInt() {
			return 1
		} else {
			return 0
		}
	}
}

func (z ZInt) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(ZInt)
	return ZBool(ok)
}

func (z ZInt) Size() int { return 8 }

type ZDouble float64

func (z ZDouble) ToInt() ZInt                { return ZInt(z) }
func (z ZDouble) ToDouble() ZDouble          { return z }
func (z ZDouble) ToBool() ZBool              { return z != 0 }
func (z ZDouble) ToString() ZString          { return ZString(fmt.Sprintf("%g", z)) }
func (z ZDouble) ToArray() ZArray            { return ZArray{List: []ZVal{z}} }
func (z ZDouble) IsIdentical(val ZVal) ZBool { return z == val }
func (z ZDouble) String() string             { return "float" }

func (z ZDouble) Compare(val ZVal) ZInt {
	switch val.(type) {
	case ZArray, ZString:
		return val.Compare(z)
	}

	if z > val.ToDouble() {
		return 1
	} else if z < val.ToDouble() {
		return -1
	} else {
		return 0
	}
}

func (z ZDouble) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(ZDouble)
	return ZBool(ok)
}

func (z ZDouble) Size() int { return 8 }

type ZBool bool

func (z ZBool) ToInt() ZInt {
	if z {
		return 1
	}

	return 0
}

func (z ZBool) ToDouble() ZDouble {
	if z {
		return 1
	}

	return 0
}

func (z ZBool) ToBool() ZBool              { return z }
func (z ZBool) ToString() ZString          { return z.ToInt().ToString() }
func (z ZBool) ToArray() ZArray            { return ZArray{List: []ZVal{z}} }
func (z ZBool) IsIdentical(val ZVal) ZBool { return val == z }
func (z ZBool) String() string             { return "bool" }

func (z ZBool) Compare(val ZVal) ZInt {
	if val.ToBool() == z {
		return 0
	} else if z {
		return 1
	} else {
		return -1
	}
}

func (z ZBool) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(ZBool)
	return ZBool(ok)
}

func (z ZBool) Size() int { return 1 }

type ZString string

func (z ZString) ToInt() ZInt {
	var integer []byte

loop:
	for _, letter := range z {
		switch letter {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			integer = append(integer, byte(letter))
		default:
			break loop
		}
	}

	res, _ := strconv.Atoi(*(*string)(unsafe.Pointer(&integer)))

	return ZInt(res)
}

func (z ZString) ToDouble() ZDouble {
	var float []byte

loop:
	for _, letter := range z {
		switch letter {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E', '-':
			float = append(float, byte(letter))
		default:
			break loop
		}
	}

	res, _ := strconv.ParseFloat(*(*string)(unsafe.Pointer(&float)), 64)

	return ZDouble(res)
}

func (z ZString) ToBool() ZBool              { return len(z) != 0 }
func (z ZString) ToString() ZString          { return z }
func (z ZString) ToArray() ZArray            { return ZArray{List: []ZVal{z}} }
func (z ZString) IsIdentical(val ZVal) ZBool { return z == val }
func (z ZString) String() string             { return "string" }

func (z ZString) Compare(val ZVal) ZInt {
	switch val.(type) {
	case ZDouble, ZInt, ZBool:
		if num, err := strconv.ParseFloat(string(z), 64); err != nil {
			return val.Compare(ZDouble(num))
		}
	}

	v := val.ToString()

	if z < v {
		return -1
	} else if z > v {
		return 1
	} else {
		return 0
	}
}

func (z ZString) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(ZString)

	return ZBool(ok)
}

func (z ZString) Size() int { return int(unsafe.Sizeof(z)) }

type ZArray struct {
	Hash map[ZVal]ZVal
	List []ZVal
}

func (z ZArray) ToInt() ZInt {
	if z.Hash == nil {
		if len(z.Hash) > 0 {
			return 1
		}

		return 0
	}

	if len(z.List) > 0 {
		return 1
	}

	return 0
}

func (z ZArray) ToDouble() ZDouble {
	if z.Hash == nil {
		if len(z.Hash) > 0 {
			return 1
		}

		return 0
	}

	if len(z.List) > 0 {
		return 1
	}

	return 0
}

func (z ZArray) ToBool() ZBool {
	if z.Len() > 0 {
		return true
	}

	return false
}

func (z ZArray) ToString() ZString { return "Array" }
func (z ZArray) ToArray() ZArray   { return z }
func (z ZArray) String() string    { return "array" }

func (z ZArray) IsIdentical(val ZVal) ZBool {
	_, ok := val.(*ZArray)

	return ZBool(ok) && z.Compare(val) == 0
}

func (z ZArray) Compare(val ZVal) ZInt {
	switch val := val.(type) {
	case *ZArray:
		if z.Len() > val.Len() {
			return 1
		} else if z.Len() < val.Len() {
			return -1
		}

		var result ZInt

		z.Iterate(func(key, value ZVal) bool {
			v := val.Get(key)

			if v == Null {
				result = 1
				return false
			}

			cmp := value.Compare(v)

			if cmp != 0 {
				result = cmp
				return false
			}

			return true
		})

		return result
	default:
		return 1
	}
}

func (z ZArray) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(*ZArray)
	return ZBool(ok)
}

func (z ZArray) ToHash() {
	z.Hash = make(map[ZVal]ZVal, len(z.List))

	for index, value := range z.List {
		z.Hash[ZInt(index)] = value
	}
}

func (z ZArray) IsList() ZBool {
	return z.List != nil
}

func (z ZArray) IsHash() ZBool {
	return z.Hash != nil
}

func (z ZArray) IsEmpty() ZBool {
	return z.Hash == nil && len(z.List) == 0 || z.List == nil && len(z.Hash) == 0
}

func (z ZArray) Len() ZInt {
	switch {
	case bool(z.IsList()):
		return ZInt(len(z.List))
	case bool(z.IsHash()):
		return ZInt(len(z.Hash))
	default:
		return 0
	}
}

func (z ZArray) Iterate(callback func(key, value ZVal) bool) {
	if z.Len() == 0 {
		return
	}

	if z.IsHash() {
		for key, value := range z.Hash {
			if !callback(key, value) {
				break
			}
		}

		return
	}

	for index, value := range z.List {
		if !callback(ZInt(index), value) {
			break
		}
	}
}

func (z ZArray) Get(key ZVal) ZVal {
	if z.Len() == 0 {
		return ZNull{}
	}

	if z.IsList() {
		if key, ok := key.(ZInt); ok && key > 0 && int(key) < len(z.List) {
			return z.List[key]
		}

		return ZNull{}
	}

	if value, ok := z.Hash[key]; ok {
		return value
	}

	return ZNull{}
}

func (z ZArray) Size() int { return int(unsafe.Sizeof(z)) }

type ZNull struct{}

func (z ZNull) ToInt() ZInt                { return 0 }
func (z ZNull) ToDouble() ZDouble          { return 0 }
func (z ZNull) ToBool() ZBool              { return false }
func (z ZNull) ToString() ZString          { return "" }
func (z ZNull) ToArray() ZArray            { return ZArray{} }
func (z ZNull) IsIdentical(val ZVal) ZBool { return z.IsOfSameType(val) }
func (z ZNull) Compare(val ZVal) ZInt      { return z.ToBool().Compare(val) }
func (z ZNull) String() string             { return "null" }

func (z ZNull) IsOfSameType(val ZVal) ZBool {
	_, ok := val.(ZNull)
	return ZBool(ok)
}

func (z ZNull) Size() int { return 0 }

type ZVal interface {
	// ToInt Casts value to integer
	ToInt() ZInt
	// ToDouble Casts value to float64
	ToDouble() ZDouble
	// ToBool Casts value to boolean
	ToBool() ZBool
	// ToString Casts value to string
	ToString() ZString
	// ToArray Casts value to array
	ToArray() ZArray
	// IsIdentical checks if values are of same type and have the same value
	IsIdentical(ZVal) ZBool
	// Compare returns 0 if values are equal and 1 or -1 if one of values is greater or less
	Compare(ZVal) ZInt
	// IsOfSameType checks if values are of same type
	IsOfSameType(ZVal) ZBool
	// String return string representation of type
	String() string
	// Size return size of underlying value
	Size() int
}

var Null ZNull

type ZVariable struct {
	ZVal
	Name string
}

func AllocVariable(name string, value ZVal) *ZVariable {
	variable := variablesPool.Get().(*ZVariable)
	variable.Name = name
	variable.ZVal = value

	return variable
}

func DeallocVariable(variable *ZVariable) {
	variablesPool.Put(variable)
}
