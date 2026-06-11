// Package ast re-exports the internal AST API for external consumers of this
// fork. See package github.com/microsoft/typescript-go/shim for the governing
// policy.
package ast

import (
	"github.com/microsoft/typescript-go/internal/ast"
)

type (
	Node          = ast.Node
	NodeList      = ast.NodeList
	SourceFile    = ast.SourceFile
	Symbol        = ast.Symbol
	Diagnostic    = ast.Diagnostic
	Kind          = ast.Kind
	ModifierFlags = ast.ModifierFlags
	SymbolFlags   = ast.SymbolFlags
)

// Node kinds the schema walker matches on. Expand as walker code demands.
const (
	KindNumericLiteral              = ast.KindNumericLiteral
	KindStringLiteral               = ast.KindStringLiteral
	KindIdentifier                  = ast.KindIdentifier
	KindDecorator                   = ast.KindDecorator
	KindPropertyDeclaration         = ast.KindPropertyDeclaration
	KindMethodDeclaration           = ast.KindMethodDeclaration
	KindTypeReference               = ast.KindTypeReference
	KindArrayType                   = ast.KindArrayType
	KindUnionType                   = ast.KindUnionType
	KindObjectLiteralExpression     = ast.KindObjectLiteralExpression
	KindPropertyAccessExpression    = ast.KindPropertyAccessExpression
	KindCallExpression              = ast.KindCallExpression
	KindClassDeclaration            = ast.KindClassDeclaration
	KindInterfaceDeclaration        = ast.KindInterfaceDeclaration
	KindTypeAliasDeclaration        = ast.KindTypeAliasDeclaration
	KindEnumDeclaration             = ast.KindEnumDeclaration
	KindHeritageClause              = ast.KindHeritageClause
	KindSourceFile                  = ast.KindSourceFile
	KindExtendsKeyword              = ast.KindExtendsKeyword
	KindImplementsKeyword           = ast.KindImplementsKeyword
	KindImportDeclaration           = ast.KindImportDeclaration
	KindNamedImports                = ast.KindNamedImports
	KindNamespaceImport             = ast.KindNamespaceImport
	KindImportSpecifier             = ast.KindImportSpecifier
	KindArrayLiteralExpression      = ast.KindArrayLiteralExpression
	KindTrueKeyword                 = ast.KindTrueKeyword
	KindFalseKeyword                = ast.KindFalseKeyword
	KindNullKeyword                 = ast.KindNullKeyword
	KindPrefixUnaryExpression       = ast.KindPrefixUnaryExpression
	KindMinusToken                  = ast.KindMinusToken
	KindTypeLiteral                 = ast.KindTypeLiteral
	KindPropertySignature           = ast.KindPropertySignature
	KindLiteralType                 = ast.KindLiteralType
	KindParameter                   = ast.KindParameter
	KindEnumMember                  = ast.KindEnumMember
	KindExpressionWithTypeArguments = ast.KindExpressionWithTypeArguments
	KindPropertyAssignment          = ast.KindPropertyAssignment
	KindExportAssignment            = ast.KindExportAssignment
	KindTypeOperator                = ast.KindTypeOperator
	KindParenthesizedType           = ast.KindParenthesizedType
	KindQualifiedName               = ast.KindQualifiedName
)

// Modifier flags the schema walker matches on.
const (
	ModifierFlagsNone     = ast.ModifierFlagsNone
	ModifierFlagsExport   = ast.ModifierFlagsExport
	ModifierFlagsAbstract = ast.ModifierFlagsAbstract
)

var (
	GetSourceFileOfNode  = ast.GetSourceFileOfNode
	HasSyntacticModifier = ast.HasSyntacticModifier
)
