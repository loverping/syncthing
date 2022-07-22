package runtimeos

import "runtime"

const (
	AIX       = "aix"
	Android   = "android"
	Darwin    = "darwin"
	Dragonfly = "dragonfly"
	FreeBSD   = "freebsd"
	Illumos   = "illumos"
	IOS       = "ios"
	JS        = "js"
	Linux     = "linux"
	NetBSD    = "netbsd"
	OpenBSD   = "openbsd"
	Plan9     = "plan9"
	Solaris   = "solaris"
	Windows   = "windows"
)

const (
	IsAIX       = runtime.GOOS == AIX
	IsAndroid   = runtime.GOOS == Android
	IsDarwin    = runtime.GOOS == Darwin
	IsDragonfly = runtime.GOOS == Dragonfly
	IsFreeBSD   = runtime.GOOS == FreeBSD
	IsIllumos   = runtime.GOOS == Illumos
	IsIOS       = runtime.GOOS == IOS
	IsJS        = runtime.GOOS == JS
	IsLinux     = runtime.GOOS == Linux
	IsNetBSD    = runtime.GOOS == NetBSD
	IsOpenBSD   = runtime.GOOS == OpenBSD
	IsPlan9     = runtime.GOOS == Plan9
	IsSolaris   = runtime.GOOS == Solaris
	IsWindows   = runtime.GOOS == Windows
)
