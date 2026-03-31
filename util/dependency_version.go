package util

import (
	"runtime/debug"
)

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("no build info found")
	}

	return info.Main.Version
}

func DependencyVersion(dependency string) string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		panic("no build info found")
	}

	for _, dep := range info.Deps {
		if dep.Path == dependency {
			return dep.Version
		}
	}

	return "?"
}
