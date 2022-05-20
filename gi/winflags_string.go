// Code generated by "stringer -type=WinFlags"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WinFlagHasGeomPrefs-24]
	_ = x[WinFlagUpdating-25]
	_ = x[WinFlagIsClosing-26]
	_ = x[WinFlagIsResizing-27]
	_ = x[WinFlagGotPaint-28]
	_ = x[WinFlagGotFocus-29]
	_ = x[WinFlagSentShow-30]
	_ = x[WinFlagGoLoop-31]
	_ = x[WinFlagStopEventLoop-32]
	_ = x[WinFlagDoFullRender-33]
	_ = x[WinFlagPublishFullReRender-34]
	_ = x[WinFlagFocusActive-35]
	_ = x[WinFlagsN-36]
}

const _WinFlags_name = "WinFlagHasGeomPrefsWinFlagUpdatingWinFlagIsClosingWinFlagIsResizingWinFlagGotPaintWinFlagGotFocusWinFlagSentShowWinFlagGoLoopWinFlagStopEventLoopWinFlagDoFullRenderWinFlagPublishFullReRenderWinFlagFocusActiveWinFlagsN"

var _WinFlags_index = [...]uint8{0, 19, 34, 50, 67, 82, 97, 112, 125, 145, 164, 190, 208, 217}

func (i WinFlags) String() string {
	i -= 24
	if i < 0 || i >= WinFlags(len(_WinFlags_index)-1) {
		return "WinFlags(" + strconv.FormatInt(int64(i+24), 10) + ")"
	}
	return _WinFlags_name[_WinFlags_index[i]:_WinFlags_index[i+1]]
}

func StringToWinFlags(s string) (WinFlags, error) {
	for i := 0; i < len(_WinFlags_index)-1; i++ {
		if s == _WinFlags_name[_WinFlags_index[i]:_WinFlags_index[i+1]] {
			return WinFlags(i + 24), nil
		}
	}
	return 0, errors.New("String: " + s + " is not a valid option for type: WinFlags")
}
