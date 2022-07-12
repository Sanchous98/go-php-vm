package compiler

import (
	"github.com/Sanchous98/go-php-vm"
	"github.com/VKCOM/php-parser/pkg/ast"
)

type Compiler struct {
	ctx vm.Context
}

func NewCompiler(ctx vm.Context) *Compiler {
	return &Compiler{ctx: ctx}
}

func (c *Compiler) Root(n *ast.Root) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Nullable(n *ast.Nullable) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Parameter(n *ast.Parameter) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Identifier(n *ast.Identifier) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Argument(n *ast.Argument) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) MatchArm(n *ast.MatchArm) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Union(n *ast.Union) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Intersection(n *ast.Intersection) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) Attribute(n *ast.Attribute) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) AttributeGroup(n *ast.AttributeGroup) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtBreak(n *ast.StmtBreak) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtCase(n *ast.StmtCase) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtCatch(n *ast.StmtCatch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtEnum(n *ast.StmtEnum) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) EnumCase(n *ast.EnumCase) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtClass(n *ast.StmtClass) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtClassConstList(n *ast.StmtClassConstList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtClassMethod(n *ast.StmtClassMethod) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtConstList(n *ast.StmtConstList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtConstant(n *ast.StmtConstant) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtContinue(n *ast.StmtContinue) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtDeclare(n *ast.StmtDeclare) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtDefault(n *ast.StmtDefault) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtDo(n *ast.StmtDo) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtEcho(n *ast.StmtEcho) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtElse(n *ast.StmtElse) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtElseIf(n *ast.StmtElseIf) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtExpression(n *ast.StmtExpression) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtFinally(n *ast.StmtFinally) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtFor(n *ast.StmtFor) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtForeach(n *ast.StmtForeach) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtFunction(n *ast.StmtFunction) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtGlobal(n *ast.StmtGlobal) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtGoto(n *ast.StmtGoto) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtHaltCompiler(n *ast.StmtHaltCompiler) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtIf(n *ast.StmtIf) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtInlineHtml(n *ast.StmtInlineHtml) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtInterface(n *ast.StmtInterface) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtLabel(n *ast.StmtLabel) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtNamespace(n *ast.StmtNamespace) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtNop(n *ast.StmtNop) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtProperty(n *ast.StmtProperty) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtPropertyList(n *ast.StmtPropertyList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtReturn(n *ast.StmtReturn) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtStatic(n *ast.StmtStatic) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtStaticVar(n *ast.StmtStaticVar) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtStmtList(n *ast.StmtStmtList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtSwitch(n *ast.StmtSwitch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtThrow(n *ast.StmtThrow) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtTrait(n *ast.StmtTrait) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtTraitUse(n *ast.StmtTraitUse) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtTraitUseAlias(n *ast.StmtTraitUseAlias) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtTraitUsePrecedence(n *ast.StmtTraitUsePrecedence) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtTry(n *ast.StmtTry) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtUnset(n *ast.StmtUnset) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtUse(n *ast.StmtUseList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtGroupUse(n *ast.StmtGroupUseList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtUseDeclaration(n *ast.StmtUse) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) StmtWhile(n *ast.StmtWhile) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprArray(n *ast.ExprArray) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprArrayDimFetch(n *ast.ExprArrayDimFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprArrayItem(n *ast.ExprArrayItem) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprArrowFunction(n *ast.ExprArrowFunction) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBrackets(n *ast.ExprBrackets) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBitwiseNot(n *ast.ExprBitwiseNot) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBooleanNot(n *ast.ExprBooleanNot) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprClassConstFetch(n *ast.ExprClassConstFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprClone(n *ast.ExprClone) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprClosure(n *ast.ExprClosure) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprClosureUse(n *ast.ExprClosureUse) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprConstFetch(n *ast.ExprConstFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprEmpty(n *ast.ExprEmpty) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprErrorSuppress(n *ast.ExprErrorSuppress) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprEval(n *ast.ExprEval) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprExit(n *ast.ExprExit) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprFunctionCall(n *ast.ExprFunctionCall) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprInclude(n *ast.ExprInclude) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprIncludeOnce(n *ast.ExprIncludeOnce) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprInstanceOf(n *ast.ExprInstanceOf) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprIsset(n *ast.ExprIsset) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprList(n *ast.ExprList) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprMethodCall(n *ast.ExprMethodCall) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprNullsafeMethodCall(n *ast.ExprNullsafeMethodCall) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprMatch(n *ast.ExprMatch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprNew(n *ast.ExprNew) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPostDec(n *ast.ExprPostDec) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPostInc(n *ast.ExprPostInc) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPreDec(n *ast.ExprPreDec) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPreInc(n *ast.ExprPreInc) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPrint(n *ast.ExprPrint) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprPropertyFetch(n *ast.ExprPropertyFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprNullsafePropertyFetch(n *ast.ExprNullsafePropertyFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprRequire(n *ast.ExprRequire) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprRequireOnce(n *ast.ExprRequireOnce) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprShellExec(n *ast.ExprShellExec) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprStaticCall(n *ast.ExprStaticCall) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprStaticPropertyFetch(n *ast.ExprStaticPropertyFetch) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprTernary(n *ast.ExprTernary) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprThrow(n *ast.ExprThrow) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprUnaryMinus(n *ast.ExprUnaryMinus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprUnaryPlus(n *ast.ExprUnaryPlus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprVariable(n *ast.ExprVariable) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprYield(n *ast.ExprYield) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprYieldFrom(n *ast.ExprYieldFrom) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssign(n *ast.ExprAssign) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignReference(n *ast.ExprAssignReference) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignBitwiseAnd(n *ast.ExprAssignBitwiseAnd) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignBitwiseOr(n *ast.ExprAssignBitwiseOr) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignBitwiseXor(n *ast.ExprAssignBitwiseXor) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignCoalesce(n *ast.ExprAssignCoalesce) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignConcat(n *ast.ExprAssignConcat) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignDiv(n *ast.ExprAssignDiv) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignMinus(n *ast.ExprAssignMinus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignMod(n *ast.ExprAssignMod) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignMul(n *ast.ExprAssignMul) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignPlus(n *ast.ExprAssignPlus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignPow(n *ast.ExprAssignPow) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignShiftLeft(n *ast.ExprAssignShiftLeft) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprAssignShiftRight(n *ast.ExprAssignShiftRight) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryBitwiseAnd(n *ast.ExprBinaryBitwiseAnd) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryBitwiseOr(n *ast.ExprBinaryBitwiseOr) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryBitwiseXor(n *ast.ExprBinaryBitwiseXor) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryBooleanAnd(n *ast.ExprBinaryBooleanAnd) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryBooleanOr(n *ast.ExprBinaryBooleanOr) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryCoalesce(n *ast.ExprBinaryCoalesce) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryConcat(n *ast.ExprBinaryConcat) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryDiv(n *ast.ExprBinaryDiv) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryEqual(n *ast.ExprBinaryEqual) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryGreater(n *ast.ExprBinaryGreater) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryGreaterOrEqual(n *ast.ExprBinaryGreaterOrEqual) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryIdentical(n *ast.ExprBinaryIdentical) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryLogicalAnd(n *ast.ExprBinaryLogicalAnd) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryLogicalOr(n *ast.ExprBinaryLogicalOr) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryLogicalXor(n *ast.ExprBinaryLogicalXor) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryMinus(n *ast.ExprBinaryMinus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryMod(n *ast.ExprBinaryMod) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryMul(n *ast.ExprBinaryMul) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryNotEqual(n *ast.ExprBinaryNotEqual) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryNotIdentical(n *ast.ExprBinaryNotIdentical) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryPlus(n *ast.ExprBinaryPlus) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryPow(n *ast.ExprBinaryPow) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryShiftLeft(n *ast.ExprBinaryShiftLeft) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinaryShiftRight(n *ast.ExprBinaryShiftRight) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinarySmaller(n *ast.ExprBinarySmaller) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinarySmallerOrEqual(n *ast.ExprBinarySmallerOrEqual) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprBinarySpaceship(n *ast.ExprBinarySpaceship) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastArray(n *ast.ExprCastArray) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastBool(n *ast.ExprCastBool) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastDouble(n *ast.ExprCastDouble) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastInt(n *ast.ExprCastInt) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastObject(n *ast.ExprCastObject) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastString(n *ast.ExprCastString) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ExprCastUnset(n *ast.ExprCastUnset) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarDnumber(n *ast.ScalarDnumber) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarEncapsed(n *ast.ScalarEncapsed) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarEncapsedStringPart(n *ast.ScalarEncapsedStringPart) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarEncapsedStringVar(n *ast.ScalarEncapsedStringVar) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarEncapsedStringBrackets(n *ast.ScalarEncapsedStringBrackets) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarHeredoc(n *ast.ScalarHeredoc) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarLnumber(n *ast.ScalarLnumber) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarMagicConstant(n *ast.ScalarMagicConstant) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) ScalarString(n *ast.ScalarString) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) NameName(n *ast.Name) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) NameFullyQualified(n *ast.NameFullyQualified) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) NameRelative(n *ast.NameRelative) {
	//TODO implement me
	panic("implement me")
}

func (c *Compiler) NameNamePart(n *ast.NamePart) {
	//TODO implement me
	panic("implement me")
}

func Compile(root *ast.Root) (ctx *vm.GlobalContext) {
	ctx = new(vm.GlobalContext)
	NewCompiler(ctx).Root(root)

	return ctx
}
