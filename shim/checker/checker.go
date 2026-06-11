// Package checker re-exports the internal type checker API for external
// consumers of this fork. See package github.com/microsoft/typescript-go/shim
// for the governing policy.
package checker

import (
	"github.com/microsoft/typescript-go/internal/checker"
)

type (
	Checker   = checker.Checker
	Type      = checker.Type
	TypeFlags = checker.TypeFlags
)

// Type flags the schema walker matches on. Expand as walker code demands.
const (
	TypeFlagsNone           = checker.TypeFlagsNone
	TypeFlagsAny            = checker.TypeFlagsAny
	TypeFlagsUnknown        = checker.TypeFlagsUnknown
	TypeFlagsUndefined      = checker.TypeFlagsUndefined
	TypeFlagsNull           = checker.TypeFlagsNull
	TypeFlagsVoid           = checker.TypeFlagsVoid
	TypeFlagsString         = checker.TypeFlagsString
	TypeFlagsNumber         = checker.TypeFlagsNumber
	TypeFlagsBoolean        = checker.TypeFlagsBoolean
	TypeFlagsStringLiteral  = checker.TypeFlagsStringLiteral
	TypeFlagsNumberLiteral  = checker.TypeFlagsNumberLiteral
	TypeFlagsBooleanLiteral = checker.TypeFlagsBooleanLiteral
	TypeFlagsEnumLiteral    = checker.TypeFlagsEnumLiteral
	TypeFlagsNever          = checker.TypeFlagsNever
	TypeFlagsTypeParameter  = checker.TypeFlagsTypeParameter
	TypeFlagsObject         = checker.TypeFlagsObject
	TypeFlagsUnion          = checker.TypeFlagsUnion
	TypeFlagsIntersection   = checker.TypeFlagsIntersection
	TypeFlagsBooleanLike    = checker.TypeFlagsBooleanLike
	TypeFlagsEnumLike       = checker.TypeFlagsEnumLike
	TypeFlagsLiteral        = checker.TypeFlagsLiteral
	TypeFlagsNullable       = checker.TypeFlagsNullable
	TypeFlagsStringLike     = checker.TypeFlagsStringLike
	TypeFlagsNumberLike     = checker.TypeFlagsNumberLike
)
