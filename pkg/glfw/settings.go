package glfw

import "github.com/go-gl/glfw/v3.3/glfw"

type Settings struct {
	Major   int
	Minor   int
	Core    bool
	Forward bool
}

func (s Settings) SupportsCore() bool {
	return s.Major > 3 || (s.Major == 3 && s.Minor >= 2)
}

func (s Settings) HasVersion() bool {
	return s.Major != 0 || s.Minor != 0
}

func (s Settings) apply() {
	if s.HasVersion() {
		glfw.WindowHint(glfw.ContextVersionMajor, s.Major)
		glfw.WindowHint(glfw.ContextVersionMinor, s.Minor)
	}
	if s.Core && s.SupportsCore() {
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		if s.Forward {
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		}
	}
}
