// Code generated by "stringer -type KicadLayer"; DO NOT EDIT

package gokicadlib

import "fmt"

const _KicadLayer_name = "Edge_CutsMarginEco2_UserEco1_UserCmts_UserDwgs_UserF_FabF_CrtYdF_AdhesF_SilkSF_PasteF_MaskF_CuIn1_CuIn2_CuIn3_CuIn4_CuIn5_CuIn6_CuIn7_CuIn8_CuIn9_CuIn10_CuIn11_CuIn12_CuIn13_CuIn14_CuIn15_CuIn16_CuIn17_CuIn18_CuIn19_CuIn20_CuIn21_CuIn22_CuIn23_CuIn24_CuIn25_CuIn26_CuIn27_CuIn28_CuIn29_CuIn30_CuB_CuB_MaskB_PasteB_SilkSB_AdhesB_CrtYdB_Fab"

var _KicadLayer_index = [...]uint16{0, 9, 15, 24, 33, 42, 51, 56, 63, 70, 77, 84, 90, 94, 100, 106, 112, 118, 124, 130, 136, 142, 148, 155, 162, 169, 176, 183, 190, 197, 204, 211, 218, 225, 232, 239, 246, 253, 260, 267, 274, 281, 288, 295, 299, 305, 312, 319, 326, 333, 338}

func (i KicadLayer) String() string {
	if i < 0 || i >= KicadLayer(len(_KicadLayer_index)-1) {
		return fmt.Sprintf("KicadLayer(%d)", i)
	}
	return _KicadLayer_name[_KicadLayer_index[i]:_KicadLayer_index[i+1]]
}

func (s KicadLayerSlice) String() []string {
	var ret []string
	for _, l := range s {
		ret = append(ret, l.String())
	}
	return ret
}
