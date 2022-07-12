package vm

import (
	"fmt"
	"math"
)

func Add(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	if op1.IsOfSameType(op2) {
		switch op1.(type) {
		case ZInt:
			return op1.(ZInt) + op2.(ZInt), nil
		case ZDouble:
			return op1.(ZDouble) + op2.(ZDouble), nil
		case ZBool:
			if op1.(ZBool) && op2.(ZBool) {
				return ZInt(2), nil
			} else if !(op1.(ZBool) || op2.(ZBool)) {
				return ZInt(0), nil
			} else {
				return ZInt(1), nil
			}
		case ZArray:
			return arrayAdd(op1.(ZArray), op2.(ZArray)), nil
		}
	}

	switch op1.(type) {
	case ZInt:
		switch op2.(type) {
		case ZBool:
			return op1.(ZInt) + op2.ToInt(), nil
		case ZDouble:
			return op1.ToDouble() + op2.(ZDouble), nil
		}
	case ZDouble:
		return op1.(ZDouble) + op2.ToDouble(), nil
	case ZBool:
		switch op2.(type) {
		case ZDouble:
			return op1.ToDouble() + op2.(ZDouble), nil
		case ZInt:
			return op1.ToInt() + op2.(ZInt), nil
		}
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "+")
}

// Adding arrays algorithm is to complex, so it's moved to the separate function
func arrayAdd(op1, op2 ZArray) (result ZArray) {
	if op1.Compare(op2) == 0 || op2.IsEmpty() {
		return op1
	}

	if op1.IsEmpty() {
		return op2
	}

	if op1.IsList() && op2.IsList() {
		result.List = make([]ZVal, NumMax(op1.Len(), op2.Len()))
		copy(result.List, op2.List)

		for index, value := range op1.List {
			result.List[index] = value
		}

		return
	}

	result.Hash = make(map[ZVal]ZVal, NumMax(op1.Len(), op2.Len()))

	op2.Iterate(func(key, value ZVal) bool {
		result.Hash[key] = value
		return true
	})

	op1.Iterate(func(key, value ZVal) bool {
		result.Hash[key] = value
		return true
	})

	return
}

func Sub(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZInt:
		switch op2.(type) {
		case ZInt, ZBool:
			return op1.(ZInt) - op2.ToInt(), nil
		case ZDouble:
			return op1.ToDouble() - op2.(ZDouble), nil
		}
	case ZDouble:
		return op1.(ZDouble) - op2.ToDouble(), nil
	case ZBool:
		switch op2.(type) {
		case ZDouble:
			return op1.ToDouble() - op2.(ZDouble), nil
		case ZInt, ZBool:
			return op1.ToInt() - op2.ToInt(), nil
		}
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "-")
}

func Mul(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZInt:
		switch op2.(type) {
		case ZInt, ZBool:
			return op1.(ZInt) * op2.ToInt(), nil
		case ZDouble:
			return op1.ToDouble() * op2.(ZDouble), nil
		}
	case ZDouble:
		return op1.(ZDouble) * op2.ToDouble(), nil
	case ZBool:
		switch op2.(type) {
		case ZDouble:
			return op1.ToDouble() * op2.(ZDouble), nil
		case ZInt, ZBool:
			return op1.ToInt() * op2.ToInt(), nil
		}
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "-")
}

func Div(_ Context, op1, op2 ZVal) (ZVal, error) {
	if op2.ToDouble() == 0 {
		return nil, fmt.Errorf("Division by zero")
	}

	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZInt:
		switch op2.(type) {
		case ZInt, ZBool:
			return op1.(ZInt) / op2.ToInt(), nil
		case ZDouble:
			return op1.ToDouble() / op2.(ZDouble), nil
		}
	case ZDouble:
		return op1.(ZDouble) / op2.ToDouble(), nil
	case ZBool:
		switch op2.(type) {
		case ZDouble:
			return op1.ToDouble() / op2.(ZDouble), nil
		case ZInt, ZBool:
			return op1.ToInt() / op2.ToInt(), nil
		}
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "/")
}

func Mod(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() % op2.ToInt(), nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "%")
}

func Sl(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() << op2.ToInt(), nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "<<")
}

func Sr(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() >> op2.ToInt(), nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, ">>")
}

func Concat(_ Context, op1, op2 ZVal) (ZVal, error) {
	return op1.ToString() + op2.ToString(), nil
}

func BitwiseOr(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() | op2.ToInt(), nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "|")
}

func BitwiseAnd(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() & op2.ToInt(), nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "&")
}

func BitwiseXor(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return op1.ToInt() ^ op2.ToInt(), nil
	case ZString:
		// Special case for strings
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "^")
}

func Pow(_ Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		result := ZDouble(math.Pow(float64(op1.ToDouble()), float64(op2.ToDouble())))

		if _, ok := op1.(ZDouble); !ok {
			if _, ok := op2.(ZDouble); !ok {
				return result.ToInt(), nil
			}
		}

		return result, nil
	}

	return ZInt(0), UnsupportedOperandType(op1, op2, "**")
}

func BitwiseNot(_ Context, op1, _ ZVal) (ZVal, error) {
	op1 = UnwrapVariable(op1)

	switch op1.(type) {
	case ZDouble, ZBool, ZInt:
		return ^op1.ToInt(), nil
	case ZString:
		// Special case for strings
	}

	return ZInt(0), UnsupportedOperandType(op1, nil, "~")
}

func BoolNot(_ Context, op1, _ ZVal) (ZVal, error) {
	return !op1.ToBool(), nil
}

func BoolXor(_ Context, op1, op2 ZVal) (ZVal, error) {
	return ZBool(op1.ToBool() != op2.ToBool()), nil
}

func IsIdentical(_ Context, op1, op2 ZVal) (ZVal, error) {
	return op1.IsIdentical(op2), nil
}

func IsNotIdentical(_ Context, op1, op2 ZVal) (ZVal, error) {
	return !op1.IsIdentical(op2), nil
}

func IsEqual(_ Context, op1, op2 ZVal) (ZVal, error) {
	return ZBool(op1.Compare(op2) == 0), nil
}

func IsNotEqual(_ Context, op1, op2 ZVal) (ZVal, error) {
	return ZBool(op1.Compare(op2) != 0), nil
}

func IsSmaller(_ Context, op1, op2 ZVal) (ZVal, error) {
	return ZBool(op1.Compare(op2) < 0), nil
}

func IsSmallerOrEqual(_ Context, op1, op2 ZVal) (ZVal, error) {
	return ZBool(op1.Compare(op2) <= 0), nil
}

func Assign(ctx Context, op1, op2 ZVal) (ZVal, error) {
	if v1, ok := op1.(*ZVariable); ok {
		if v2, ok := op2.(*ZVariable); ok {
			v1.ZVal = v2.ZVal
		} else {
			v1.ZVal = op2
		}

		return nil, nil
	}

	return nil, fmt.Errorf("unexpected token \"=\" ")
}

func AssignDim(ctx Context, op1, op2 ZVal) (ZVal, error) {
	op1, op2 = UnwrapVariable(op1), UnwrapVariable(op2)
	var v1 ZArray
	var ok bool

	if v1, ok = op1.(ZArray); !ok {
		return nil, fmt.Errorf("")
	}

	return AllocVariable("$1", v1.Get(UnwrapVariable(op2))), nil
}

func AssignObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignOp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignDimOp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignObjOp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignStaticPropOp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func QmAssign(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignObjRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssignStaticPropRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreInc(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreDec(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostInc(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostDec(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreIncStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreDecStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostIncStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostDecStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Jump(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JumpZ(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JumpNZ(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JumpZEx(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JumpNZEx(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Case(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CheckVariable(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendVariableNoReferenceEx(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Cast(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Bool(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FastConcat(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func RopeInit(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func RopeAdd(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func RopeEnd(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func BeginSilence(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func EndSilence(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitFcallByName(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DoFcall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitFcall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Return(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Recv(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func RecvInit(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendVal(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendVarEx(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func New(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitNsFcallByName(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Free(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitArray(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AddArrayElement(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IncludeOrEval(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UnsetVar(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UnsetDim(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UnsetObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FeResetR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FeFetchR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Exit(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchW(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimW(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjW(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchIs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimIs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjIs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchUnset(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchDimUnset(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchObjUnset(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchListR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchConstant(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CheckFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ExtStmt(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ExtFcallBegin(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ExtFcallEnd(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ExtNop(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Ticks(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendVarNoRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Catch(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Throw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchClass(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Clone(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ReturnByRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitMethodCall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitStaticMethodCall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyVar(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyDimObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendValEx(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendVar(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitUserCall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendArray(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendUser(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Strlen(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Defined(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func TypeCheck(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func VerifyReturnType(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FeResetRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FeFetchRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FeFree(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InitDynamicCall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DoIcall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DoUcall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DoFcallByName(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreIncObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PreDecObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostIncObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func PostDecObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Echo(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func OpData(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Instanceof(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func GeneratorCreate(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func MakeRef(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareFunction(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareLambdaFunction(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareConst(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareClass(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareClassDelayed(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DeclareAnonClass(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AddArrayUnpack(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyPropObj(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func HandleException(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UserOpcode(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func AssertCheck(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JmpSet(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UnsetCv(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyCv(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchListW(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Separate(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchClassName(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CallTrampoline(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func DiscardException(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Yield(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func GeneratorReturn(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FastCall(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FastRet(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func RecvVariadic(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendUnpack(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func YieldFrom(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CopyTmp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func BindGlobal(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Coalesce(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Spaceship(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FuncNumArgs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FuncGetArgs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropR(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropW(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropRw(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropIs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchStaticPropUnset(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func UnsetStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyStaticProp(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchClassConstant(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func BindLexical(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func BindStatic(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchThis(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SendFuncArg(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func IssetIsemptyThis(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SwitchLong(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func SwitchString(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func InArray(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Count(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func GetClass(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func GetCalledClass(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func GetType(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func ArrayKeyExists(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func Match(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CaseStrict(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func MatchError(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func JmpNull(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CheckUndefArgs(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func FetchGlobals(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func VerifyNeverType(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}

func CallableConvert(ctx Context, op1, op2 ZVal) (ZVal, error) {
	return nil, nil
}
