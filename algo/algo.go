//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package algo ;import _g "strconv";func _f (_be byte )bool {return _be >='0'&&_be <='9'};func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_fc :=make ([]byte ,len (s )*cnt );_df :=[]byte (s );for _da :=0;_da < cnt ;_da ++{copy (_fc [_da :],_df );};return string (_fc );};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_e ,_d :=0,0;for _e < len (lhs )&&_d < len (rhs ){_c :=lhs [_e ];_dc :=rhs [_d ];_ed :=_f (_c );_gb :=_f (_dc );switch {case _ed &&!_gb :return true ;case !_ed &&_gb :return false ;case !_ed &&!_gb :if _c !=_dc {return _c < _dc ;};_e ++;_d ++;default:_ea :=_e +1;_eab :=_d +1;for _ea < len (lhs )&&_f (lhs [_ea ]){_ea ++;};for _eab < len (rhs )&&_f (rhs [_eab ]){_eab ++;};_cc ,_ :=_g .ParseUint (lhs [_e :_ea ],10,64);_ga ,_ :=_g .ParseUint (rhs [_e :_eab ],10,64);if _cc !=_ga {return _cc < _ga ;};_e =_ea ;_d =_eab ;};};return len (lhs )< len (rhs );};