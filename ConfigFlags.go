package imgui

const (
	// Master keyboard navigation enable flag. NewFrame() will automatically fill io.NavInputs[] based on io.KeysDown[].
	ConfigFlagNavEnableKeyboard = 1 << iota
	// Master gamepad navigation enable flag. This is mostly to instruct your imgui back-end to fill io.NavInputs[]. Back-end also needs to set ImGuiBackendFlags_HasGamepad.
	ConfigFlagNavEnableGamepad
	// Instruct navigation to move the mouse cursor. May be useful on TV/console systems where moving a virtual mouse is awkward. Will update io.MousePos and set io.WantSetMousePos=true. If enabled you MUST honor io.WantSetMousePos requests in your binding, otherwise ImGui will react as if the mouse is jumping around back and forth.
	ConfigFlagNavEnableSetMousePos
	// Instruct navigation to not set the io.WantCaptureKeyboard flag when io.NavActive is set.
	ConfigFlagNavNoCaptureKeyboard
	// Instruct imgui to clear mouse position/buttons in NewFrame(). This allows ignoring the mouse information set by the back-end.
	ConfigFlagNoMouse
	// Instruct back-end to not alter mouse cursor shape and visibility. Use if the back-end cursor changes are interfering with yours and you don't want to use SetMouseCursor() to change mouse cursor. You may want to honor requests from imgui by reading GetMouseCursor() yourself instead.
	ConfigFlagNoMouseCursorChange

	// User storage (to allow your back-end/engine to communicate to code that may be shared between multiple projects. Those flags are not used by core ImGui)
	// Application is SRGB-aware.
	ConfigFlagIsSRGB = 1 << 20
	// Application is using a touch screen instead of a mouse.
	ConfigFlagIsTouchScreen = 1 << 21
)
