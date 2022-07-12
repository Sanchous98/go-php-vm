package vm

//go:generate stringer -type=Opcode -linecomment
type Opcode int
type Handler func(ctx Context, op1, op2 ZVal) (ZVal, error)

const (
	ZendNop                 Opcode = iota // ZEND_NOP
	ZendAdd                               // ZEND_ADD
	ZendSub                               // ZEND_SUB
	ZendMul                               // ZEND_MUL
	ZendDiv                               // ZEND_DIV
	ZendMod                               // ZEND_MOD
	ZendSl                                // ZEND_SL
	ZendSr                                // ZEND_SR
	ZendConcat                            // ZEND_CONCAT
	ZendBwOr                              // ZEND_BW_OR
	ZendBwAnd                             // ZEND_BW_AND
	ZendBwXor                             // ZEND_BW_XOR
	ZendPow                               // ZEND_POW
	ZendBwNot                             // ZEND_BW_NOT
	ZendBoolNot                           // ZEND_BOOL_NOT
	ZendBoolXor                           // ZEND_BOOL_XOR
	ZendIsIdentical                       // ZEND_IS_IDENTICAL
	ZendIsNotIdentical                    // ZEND_IS_NOT_IDENTICAL
	ZendIsEqual                           // ZEND_IS_EQUAL
	ZendIsNotEqual                        // ZEND_IS_NOT_EQUAL
	ZendIsSmaller                         // ZEND_IS_SMALLER
	ZendIsSmallerOrEqual                  // ZEND_IS_SMALLER_OR_EQUAL
	ZendAssign                            // ZEND_ASSIGN
	ZendAssignDim                         // ZEND_ASSIGN_DIM
	ZendAssignObj                         // ZEND_ASSIGN_OBJ
	ZendAssignStaticProp                  // ZEND_ASSIGN_STATIC_PROP
	ZendAssignOp                          // ZEND_ASSIGN_OP
	ZendAssignDimOp                       // ZEND_ASSIGN_DIM_OP
	ZendAssignObjOp                       // ZEND_ASSIGN_OBJ_OP
	ZendAssignStaticPropOp                // ZEND_ASSIGN_STATIC_PROP_OP
	ZendAssignRef                         // ZEND_ASSIGN_REF
	ZendQmAssign                          // ZEND_QM_ASSIGN
	ZendAssignObjRef                      // ZEND_ASSIGN_OBJ_REF
	ZendAssignStaticPropRef               // ZEND_ASSIGN_STATIC_PROP_REF
	ZendPreInc                            // ZEND_PRE_INC
	ZendPreDec                            // ZEND_PRE_DEC
	ZendPostInc                           // ZEND_POST_INC
	ZendPostDec                           // ZEND_POST_DEC
	ZendPreIncStaticProp                  // ZEND_PRE_INC_STATIC_PROP
	ZendPreDecStaticProp                  // ZEND_PRE_DEC_STATIC_PROP
	ZendPostIncStaticProp                 // ZEND_POST_INC_STATIC_PROP
	ZendPostDecStaticProp                 // ZEND_POST_DEC_STATIC_PROP
	ZendJmp                               // ZEND_JMP
	ZendJmpz                              // ZEND_JMPZ
	ZendJmpnz                             // ZEND_JMPNZ
	_
	ZendJmpzEx                 // ZEND_JMPZ_EX
	ZendJmpnzEx                // ZEND_JMPNZ_EX
	ZendCase                   // ZEND_CASE
	ZendCheckVar               // ZEND_CHECK_VAR
	ZendSendVarNoRefEx         // ZEND_SEND_VAR_NO_REF_EX
	ZendCast                   // ZEND_CAST
	ZendBool                   // ZEND_BOOL
	ZendFastConcat             // ZEND_FAST_CONCAT
	ZendRopeInit               // ZEND_ROPE_INIT
	ZendRopeAdd                // ZEND_ROPE_ADD
	ZendRopeEnd                // ZEND_ROPE_END
	ZendBeginSilence           // ZEND_BEGIN_SILENCE
	ZendEndSilence             // ZEND_END_SILENCE
	ZendInitFcallByName        // ZEND_INIT_FCALL_BY_NAME
	ZendDoFcall                // ZEND_DO_FCALL
	ZendInitFcall              // ZEND_INIT_FCALL
	ZendReturn                 // ZEND_RETURN
	ZendRecv                   // ZEND_RECV
	ZendRecvInit               // ZEND_RECV_INIT
	ZendSendVal                // ZEND_SEND_VAL
	ZendSendVarEx              // ZEND_SEND_VAR_EX
	ZendSendRef                // ZEND_SEND_REF
	ZendNew                    // ZEND_NEW
	ZendInitNsFcallByName      // ZEND_INIT_NS_FCALL_BY_NAME
	ZendFree                   // ZEND_FREE
	ZendInitArray              // ZEND_INIT_ARRAY
	ZendAddArrayElement        // ZEND_ADD_ARRAY_ELEMENT
	ZendIncludeOrEval          // ZEND_INCLUDE_OR_EVAL
	ZendUnsetVar               // ZEND_UNSET_VAR
	ZendUnsetDim               // ZEND_UNSET_DIM
	ZendUnsetObj               // ZEND_UNSET_OBJ
	ZendFeResetR               // ZEND_FE_RESET_R
	ZendFeFetchR               // ZEND_FE_FETCH_R
	ZendExit                   // ZEND_EXIT
	ZendFetchR                 // ZEND_FETCH_R
	ZendFetchDimR              // ZEND_FETCH_DIM_R
	ZendFetchObjR              // ZEND_FETCH_OBJ_R
	ZendFetchW                 // ZEND_FETCH_W
	ZendFetchDimW              // ZEND_FETCH_DIM_W
	ZendFetchObjW              // ZEND_FETCH_OBJ_W
	ZendFetchRw                // ZEND_FETCH_RW
	ZendFetchDimRw             // ZEND_FETCH_DIM_RW
	ZendFetchObjRw             // ZEND_FETCH_OBJ_RW
	ZendFetchIs                // ZEND_FETCH_IS
	ZendFetchDimIs             // ZEND_FETCH_DIM_IS
	ZendFetchObjIs             // ZEND_FETCH_OBJ_IS
	ZendFetchFuncArg           // ZEND_FETCH_FUNC_ARG
	ZendFetchDimFuncArg        // ZEND_FETCH_DIM_FUNC_ARG
	ZendFetchObjFuncArg        // ZEND_FETCH_OBJ_FUNC_ARG
	ZendFetchUnset             // ZEND_FETCH_UNSET
	ZendFetchDimUnset          // ZEND_FETCH_DIM_UNSET
	ZendFetchObjUnset          // ZEND_FETCH_OBJ_UNSET
	ZendFetchListR             // ZEND_FETCH_LIST_R
	ZendFetchConstant          // ZEND_FETCH_CONSTANT
	ZendCheckFuncArg           // ZEND_CHECK_FUNC_ARG
	ZendExtStmt                // ZEND_EXT_STMT
	ZendExtFcallBegin          // ZEND_EXT_FCALL_BEGIN
	ZendExtFcallEnd            // ZEND_EXT_FCALL_END
	ZendExtNop                 // ZEND_EXT_NOP
	ZendTicks                  // ZEND_TICKS
	ZendSendVarNoRef           // ZEND_SEND_VAR_NO_REF
	ZendCatch                  // ZEND_CATCH
	ZendThrow                  // ZEND_THROW
	ZendFetchClass             // ZEND_FETCH_CLASS
	ZendClone                  // ZEND_CLONE
	ZendReturnByRef            // ZEND_RETURN_BY_REF
	ZendInitMethodCall         // ZEND_INIT_METHOD_CALL
	ZendInitStaticMethodCall   // ZEND_INIT_STATIC_METHOD_CALL
	ZendIssetIsemptyVar        // ZEND_ISSET_ISEMPTY_VAR
	ZendIssetIsemptyDimObj     // ZEND_ISSET_ISEMPTY_DIM_OBJ
	ZendSendValEx              // ZEND_SEND_VAL_EX
	ZendSendVar                // ZEND_SEND_VAR
	ZendInitUserCall           // ZEND_INIT_USER_CALL
	ZendSendArray              // ZEND_SEND_ARRAY
	ZendSendUser               // ZEND_SEND_USER
	ZendStrlen                 // ZEND_STRLEN
	ZendDefined                // ZEND_DEFINED
	ZendTypeCheck              // ZEND_TYPE_CHECK
	ZendVerifyReturnType       // ZEND_VERIFY_RETURN_TYPE
	ZendFeResetRw              // ZEND_FE_RESET_RW
	ZendFeFetchRw              // ZEND_FE_FETCH_RW
	ZendFeFree                 // ZEND_FE_FREE
	ZendInitDynamicCall        // ZEND_INIT_DYNAMIC_CALL
	ZendDoIcall                // ZEND_DO_ICALL
	ZendDoUcall                // ZEND_DO_UCALL
	ZendDoFcallByName          // ZEND_DO_FCALL_BY_NAME
	ZendPreIncObj              // ZEND_PRE_INC_OBJ
	ZendPreDecObj              // ZEND_PRE_DEC_OBJ
	ZendPostIncObj             // ZEND_POST_INC_OBJ
	ZendPostDecObj             // ZEND_POST_DEC_OBJ
	ZendEcho                   // ZEND_ECHO
	ZendOpData                 // ZEND_OP_DATA
	ZendInstanceof             // ZEND_INSTANCEOF
	ZendGeneratorCreate        // ZEND_GENERATOR_CREATE
	ZendMakeRef                // ZEND_MAKE_REF
	ZendDeclareFunction        // ZEND_DECLARE_FUNCTION
	ZendDeclareLambdaFunction  // ZEND_DECLARE_LAMBDA_FUNCTION
	ZendDeclareConst           // ZEND_DECLARE_CONST
	ZendDeclareClass           // ZEND_DECLARE_CLASS
	ZendDeclareClassDelayed    // ZEND_DECLARE_CLASS_DELAYED
	ZendDeclareAnonClass       // ZEND_DECLARE_ANON_CLASS
	ZendAddArrayUnpack         // ZEND_ADD_ARRAY_UNPACK
	ZendIssetIsemptyPropObj    // ZEND_ISSET_ISEMPTY_PROP_OBJ
	ZendHandleException        // ZEND_HANDLE_EXCEPTION
	ZendUserOpcode             // ZEND_USER_OPCODE
	ZendAssertCheck            // ZEND_ASSERT_CHECK
	ZendJmpSet                 // ZEND_JMP_SET
	ZendUnsetCv                // ZEND_UNSET_CV
	ZendIssetIsemptyCv         // ZEND_ISSET_ISEMPTY_CV
	ZendFetchListW             // ZEND_FETCH_LIST_W
	ZendSeparate               // ZEND_SEPARATE
	ZendFetchClassName         // ZEND_FETCH_CLASS_NAME
	ZendCallTrampoline         // ZEND_CALL_TRAMPOLINE
	ZendDiscardException       // ZEND_DISCARD_EXCEPTION
	ZendYield                  // ZEND_YIELD
	ZendGeneratorReturn        // ZEND_GENERATOR_RETURN
	ZendFastCall               // ZEND_FAST_CALL
	ZendFastRet                // ZEND_FAST_RET
	ZendRecvVariadic           // ZEND_RECV_VARIADIC
	ZendSendUnpack             // ZEND_SEND_UNPACK
	ZendYieldFrom              // ZEND_YIELD_FROM
	ZendCopyTmp                // ZEND_COPY_TMP
	ZendBindGlobal             // ZEND_BIND_GLOBAL
	ZendCoalesce               // ZEND_COALESCE
	ZendSpaceship              // ZEND_SPACESHIP
	ZendFuncNumArgs            // ZEND_FUNC_NUM_ARGS
	ZendFuncGetArgs            // ZEND_FUNC_GET_ARGS
	ZendFetchStaticPropR       // ZEND_FETCH_STATIC_PROP_R
	ZendFetchStaticPropW       // ZEND_FETCH_STATIC_PROP_W
	ZendFetchStaticPropRw      // ZEND_FETCH_STATIC_PROP_RW
	ZendFetchStaticPropIs      // ZEND_FETCH_STATIC_PROP_IS
	ZendFetchStaticPropFuncArg // ZEND_FETCH_STATIC_PROP_FUNC_ARG
	ZendFetchStaticPropUnset   // ZEND_FETCH_STATIC_PROP_UNSET
	ZendUnsetStaticProp        // ZEND_UNSET_STATIC_PROP
	ZendIssetIsemptyStaticProp // ZEND_ISSET_ISEMPTY_STATIC_PROP
	ZendFetchClassConstant     // ZEND_FETCH_CLASS_CONSTANT
	ZendBindLexical            // ZEND_BIND_LEXICAL
	ZendBindStatic             // ZEND_BIND_STATIC
	ZendFetchThis              // ZEND_FETCH_THIS
	ZendSendFuncArg            // ZEND_SEND_FUNC_ARG
	ZendIssetIsemptyThis       // ZEND_ISSET_ISEMPTY_THIS
	ZendSwitchLong             // ZEND_SWITCH_LONG
	ZendSwitchString           // ZEND_SWITCH_STRING
	ZendInArray                // ZEND_IN_ARRAY
	ZendCount                  // ZEND_COUNT
	ZendGetClass               // ZEND_GET_CLASS
	ZendGetCalledClass         // ZEND_GET_CALLED_CLASS
	ZendGetType                // ZEND_GET_TYPE
	ZendArrayKeyExists         // ZEND_ARRAY_KEY_EXISTS
	ZendMatch                  // ZEND_MATCH
	ZendCaseStrict             // ZEND_CASE_STRICT
	ZendMatchError             // ZEND_MATCH_ERROR
	ZendJmpNull                // ZEND_JMP_NULL
	ZendCheckUndefArgs         // ZEND_CHECK_UNDEF_ARGS
	ZendFetchGlobals           // ZEND_FETCH_GLOBALS
	ZendVerifyNeverType        // ZEND_VERIFY_NEVER_TYPE
	ZendCallableConvert        // ZEND_CALLABLE_CONVERT

	ZendVMLastOpcode = ZendCallableConvert
)

func (i *Opcode) Handler() Handler {
	switch *i {
	case ZendAdd:
		return Add
	case ZendSub:
		return Sub
	case ZendMul:
		return Mul
	case ZendDiv:
		return Div
	case ZendMod:
		return Mod
	case ZendSl:
		return Sl
	case ZendSr:
		return Sr
	case ZendConcat:
		return Concat
	case ZendBwOr:
		return BitwiseOr
	case ZendBwAnd:
		return BitwiseAnd
	case ZendBwXor:
		return BitwiseXor
	case ZendPow:
		return Pow
	case ZendBwNot:
		return BitwiseNot
	case ZendBoolNot:
		return BoolNot
	case ZendBoolXor:
		return BoolXor
	case ZendIsIdentical:
		return IsIdentical
	case ZendIsNotIdentical:
		return IsNotIdentical
	case ZendIsEqual:
		return IsEqual
	case ZendIsNotEqual:
		return IsNotEqual
	case ZendIsSmaller:
		return IsSmaller
	case ZendIsSmallerOrEqual:
		return IsSmallerOrEqual
	case ZendAssign:
		return Assign
	case ZendAssignDim:
		return AssignDim
	case ZendAssignObj:
		return AssignObj
	case ZendAssignStaticProp:
		return AssignStaticProp
	case ZendAssignOp:
		return AssignOp
	case ZendAssignDimOp:
		return AssignDimOp
	case ZendAssignObjOp:
		return AssignObjOp
	case ZendAssignStaticPropOp:
		return AssignStaticPropOp
	case ZendAssignRef:
		return AssignRef
	case ZendQmAssign:
		return QmAssign
	case ZendAssignObjRef:
		return AssignObjRef
	case ZendAssignStaticPropRef:
		return AssignStaticPropRef
	case ZendPreInc:
		return PreInc
	case ZendPreDec:
		return PreDec
	case ZendPostInc:
		return PostInc
	case ZendPostDec:
		return PostDec
	case ZendPreIncStaticProp:
		return PreIncStaticProp
	case ZendPreDecStaticProp:
		return PreDecStaticProp
	case ZendPostIncStaticProp:
		return PostIncStaticProp
	case ZendPostDecStaticProp:
		return PostDecStaticProp
	case ZendJmp:
		return Jump
	case ZendJmpz:
		return JumpZ
	case ZendJmpnz:
		return JumpNZ
	case ZendJmpzEx:
		return JumpZEx
	case ZendJmpnzEx:
		return JumpNZEx
	case ZendCase:
		return Case
	case ZendCheckVar:
		return CheckVariable
	case ZendSendVarNoRefEx:
		return SendVariableNoReferenceEx
	case ZendCast:
		return Cast
	case ZendBool:
		return Bool
	case ZendFastConcat:
		return FastConcat
	case ZendRopeInit:
		return RopeInit
	case ZendRopeAdd:
		return RopeAdd
	case ZendRopeEnd:
		return RopeEnd
	case ZendBeginSilence:
		return BeginSilence
	case ZendEndSilence:
		return EndSilence
	case ZendInitFcallByName:
		return InitFcallByName
	case ZendDoFcall:
		return DoFcall
	case ZendInitFcall:
		return InitFcall
	case ZendReturn:
		return Return
	case ZendRecv:
		return Recv
	case ZendRecvInit:
		return RecvInit
	case ZendSendVal:
		return SendVal
	case ZendSendVarEx:
		return SendVarEx
	case ZendSendRef:
		return SendRef
	case ZendNew:
		return New
	case ZendInitNsFcallByName:
		return InitNsFcallByName
	case ZendFree:
		return Free
	case ZendInitArray:
		return InitArray
	case ZendAddArrayElement:
		return AddArrayElement
	case ZendIncludeOrEval:
		return IncludeOrEval
	case ZendUnsetVar:
		return UnsetVar
	case ZendUnsetDim:
		return UnsetDim
	case ZendUnsetObj:
		return UnsetObj
	case ZendFeResetR:
		return FeResetR
	case ZendFeFetchR:
		return FeFetchR
	case ZendExit:
		return Exit
	case ZendFetchR:
		return FetchR
	case ZendFetchDimR:
		return FetchDimR
	case ZendFetchObjR:
		return FetchObjR
	case ZendFetchW:
		return FetchW
	case ZendFetchDimW:
		return FetchDimW
	case ZendFetchObjW:
		return FetchObjW
	case ZendFetchRw:
		return FetchRw
	case ZendFetchDimRw:
		return FetchDimRw
	case ZendFetchObjRw:
		return FetchObjRw
	case ZendFetchIs:
		return FetchIs
	case ZendFetchDimIs:
		return FetchDimIs
	case ZendFetchObjIs:
		return FetchObjIs
	case ZendFetchFuncArg:
		return FetchFuncArg
	case ZendFetchDimFuncArg:
		return FetchDimFuncArg
	case ZendFetchObjFuncArg:
		return FetchObjFuncArg
	case ZendFetchUnset:
		return FetchUnset
	case ZendFetchDimUnset:
		return FetchDimUnset
	case ZendFetchObjUnset:
		return FetchObjUnset
	case ZendFetchListR:
		return FetchListR
	case ZendFetchConstant:
		return FetchConstant
	case ZendCheckFuncArg:
		return CheckFuncArg
	case ZendExtStmt:
		return ExtStmt
	case ZendExtFcallBegin:
		return ExtFcallBegin
	case ZendExtFcallEnd:
		return ExtFcallEnd
	case ZendExtNop:
		return ExtNop
	case ZendTicks:
		return Ticks
	case ZendSendVarNoRef:
		return SendVarNoRef
	case ZendCatch:
		return Catch
	case ZendThrow:
		return Throw
	case ZendFetchClass:
		return FetchClass
	case ZendClone:
		return Clone
	case ZendReturnByRef:
		return ReturnByRef
	case ZendInitMethodCall:
		return InitMethodCall
	case ZendInitStaticMethodCall:
		return InitStaticMethodCall
	case ZendIssetIsemptyVar:
		return IssetIsemptyVar
	case ZendIssetIsemptyDimObj:
		return IssetIsemptyDimObj
	case ZendSendValEx:
		return SendValEx
	case ZendSendVar:
		return SendVar
	case ZendInitUserCall:
		return InitUserCall
	case ZendSendArray:
		return SendArray
	case ZendSendUser:
		return SendUser
	case ZendStrlen:
		return Strlen
	case ZendDefined:
		return Defined
	case ZendTypeCheck:
		return TypeCheck
	case ZendVerifyReturnType:
		return VerifyReturnType
	case ZendFeResetRw:
		return FeResetRw
	case ZendFeFetchRw:
		return FeFetchRw
	case ZendFeFree:
		return FeFree
	case ZendInitDynamicCall:
		return InitDynamicCall
	case ZendDoIcall:
		return DoIcall
	case ZendDoUcall:
		return DoUcall
	case ZendDoFcallByName:
		return DoFcallByName
	case ZendPreIncObj:
		return PreIncObj
	case ZendPreDecObj:
		return PreDecObj
	case ZendPostIncObj:
		return PostIncObj
	case ZendPostDecObj:
		return PostDecObj
	case ZendEcho:
		return Echo
	case ZendOpData:
		return OpData
	case ZendInstanceof:
		return Instanceof
	case ZendGeneratorCreate:
		return GeneratorCreate
	case ZendMakeRef:
		return MakeRef
	case ZendDeclareFunction:
		return DeclareFunction
	case ZendDeclareLambdaFunction:
		return DeclareLambdaFunction
	case ZendDeclareConst:
		return DeclareConst
	case ZendDeclareClass:
		return DeclareClass
	case ZendDeclareClassDelayed:
		return DeclareClassDelayed
	case ZendDeclareAnonClass:
		return DeclareAnonClass
	case ZendAddArrayUnpack:
		return AddArrayUnpack
	case ZendIssetIsemptyPropObj:
		return IssetIsemptyPropObj
	case ZendHandleException:
		return HandleException
	case ZendUserOpcode:
		return UserOpcode
	case ZendAssertCheck:
		return AssertCheck
	case ZendJmpSet:
		return JmpSet
	case ZendUnsetCv:
		return UnsetCv
	case ZendIssetIsemptyCv:
		return IssetIsemptyCv
	case ZendFetchListW:
		return FetchListW
	case ZendSeparate:
		return Separate
	case ZendFetchClassName:
		return FetchClassName
	case ZendCallTrampoline:
		return CallTrampoline
	case ZendDiscardException:
		return DiscardException
	case ZendYield:
		return Yield
	case ZendGeneratorReturn:
		return GeneratorReturn
	case ZendFastCall:
		return FastCall
	case ZendFastRet:
		return FastRet
	case ZendRecvVariadic:
		return RecvVariadic
	case ZendSendUnpack:
		return SendUnpack
	case ZendYieldFrom:
		return YieldFrom
	case ZendCopyTmp:
		return CopyTmp
	case ZendBindGlobal:
		return BindGlobal
	case ZendCoalesce:
		return Coalesce
	case ZendSpaceship:
		return Spaceship
	case ZendFuncNumArgs:
		return FuncNumArgs
	case ZendFuncGetArgs:
		return FuncGetArgs
	case ZendFetchStaticPropR:
		return FetchStaticPropR
	case ZendFetchStaticPropW:
		return FetchStaticPropW
	case ZendFetchStaticPropRw:
		return FetchStaticPropRw
	case ZendFetchStaticPropIs:
		return FetchStaticPropIs
	case ZendFetchStaticPropFuncArg:
		return FetchStaticPropFuncArg
	case ZendFetchStaticPropUnset:
		return FetchStaticPropUnset
	case ZendUnsetStaticProp:
		return UnsetStaticProp
	case ZendIssetIsemptyStaticProp:
		return IssetIsemptyStaticProp
	case ZendFetchClassConstant:
		return FetchClassConstant
	case ZendBindLexical:
		return BindLexical
	case ZendBindStatic:
		return BindStatic
	case ZendFetchThis:
		return FetchThis
	case ZendSendFuncArg:
		return SendFuncArg
	case ZendIssetIsemptyThis:
		return IssetIsemptyThis
	case ZendSwitchLong:
		return SwitchLong
	case ZendSwitchString:
		return SwitchString
	case ZendInArray:
		return InArray
	case ZendCount:
		return Count
	case ZendGetClass:
		return GetClass
	case ZendGetCalledClass:
		return GetCalledClass
	case ZendGetType:
		return GetType
	case ZendArrayKeyExists:
		return ArrayKeyExists
	case ZendMatch:
		return Match
	case ZendCaseStrict:
		return CaseStrict
	case ZendMatchError:
		return MatchError
	case ZendJmpNull:
		return JmpNull
	case ZendCheckUndefArgs:
		return CheckUndefArgs
	case ZendFetchGlobals:
		return FetchGlobals
	case ZendVerifyNeverType:
		return VerifyNeverType
	case ZendCallableConvert:
		return CallableConvert
	}

	return nil
}
