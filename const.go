package tclientlib

import (
	"regexp"
)

var CodeTOASCII = map[byte]string{
	IAC:            "IAC",
	WILL:           "WILL",
	WONT:           "WONT",
	DO:             "DO",
	DONT:           "DONT",
	SE:             "SE",
	SB:             "SB",
	BINARY:         "BINARY",
	ECHO:           "ECHO",
	RCP:            "RCP",
	SGA:            "SGA",
	NAMS:           "NAMS",
	STATUS:         "STATUS ",
	TM:             "TM",
	RCTE:           "RCTE",
	NAOL:           "NAOL",
	NAOP:           "NAOP",
	NAOCRD:         "NAOCRD",
	NAOHTS:         "NAOHTS",
	NAOHTD:         "NAOHTD",
	NAOFFD:         "NAOFFD",
	NAOVTS:         "NAOVTS",
	NAOVTD:         "NAOVTD",
	NAOLFD:         "NAOLFD",
	XASCII:         "XASCII",
	LOGOUT:         "LOGOUT",
	BM:             "BM",
	DET:            "DET",
	SUPDUP:         "SUPDUP",
	SUPDUPOUTPUT:   "SUPDUPOUTPUT",
	SNDLOC:         "SNDLOC",
	TTYPE:          "TTYPE",
	EOR:            "EOR",
	TUID:           "TUID",
	OUTMRK:         "OUTMRK",
	TTYLOC:         "TTYLOC",
	VT3270REGIME:   "VT3270REGIME",
	X3PAD:          "X3PAD",
	NAWS:           "NAWS",
	TSPEED:         "TSPEED",
	LFLOW:          "LFLOW",
	LINEMODE:       "LINEMODE",
	XDISPLOC:       "XDISPLOC",
	OLD_ENVIRON:    "OLD_ENVIRON",
	AUTHENTICATION: "AUTHENTICATION",
	ENCRYPT:        "ENCRYPT",
	NEW_ENVIRON:    "NEW_ENVIRON",
}

const (
	loginRegs          = "(?i)login:?\\s*$|username:?\\s*$|name:?\\s*$|用户名:?\\s*$|账\\s*号:?\\s*$|user:?\\s*$"
	passwordRegs       = "(?i)Password:?\\s*$|ssword:?\\s*$|passwd:?\\s*$|密\\s*码:?\\s*$"
	FailedRegs         = "(?i)incorrect|failed|失败|错误"
	DefaultSuccessRegs = "(?i)Last\\s*login|success|成功|#|>|\\$"
)

var (
	usernamePattern, _            = regexp.Compile(loginRegs)
	passwordPattern, _            = regexp.Compile(passwordRegs)
	defaultLoginFailedPattern, _  = regexp.Compile(FailedRegs)
	defaultLoginSuccessPattern, _ = regexp.Compile(DefaultSuccessRegs)
)

type AuthStatus int

const (
	AuthSuccess AuthStatus = iota
	AuthPartial
	AuthFailed
)
