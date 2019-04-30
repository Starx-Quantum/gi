// Code generated by "stringer -type=LightColors"; DO NOT EDIT.

package gi3d

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

const _LightColors_name = "DirectSunCarbonArcHalogenTungsten100WTungsten40WCandleOvercastFluorWarmFluorStdFluorCoolFluorFullFluorGrowMercuryVaporSodiumVaporMetalHalideLightColorsN"

var _LightColors_index = [...]uint8{0, 9, 18, 25, 37, 48, 54, 62, 71, 79, 88, 97, 106, 118, 129, 140, 152}

func (i LightColors) String() string {
	if i < 0 || i >= LightColors(len(_LightColors_index)-1) {
		return "LightColors(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LightColors_name[_LightColors_index[i]:_LightColors_index[i+1]]
}

func (i *LightColors) FromString(s string) error {
	for j := 0; j < len(_LightColors_index)-1; j++ {
		if s == _LightColors_name[_LightColors_index[j]:_LightColors_index[j+1]] {
			*i = LightColors(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: LightColors")
}
