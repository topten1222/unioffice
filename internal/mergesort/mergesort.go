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

package mergesort ;func MergeSort (array []float64 )[]float64 {if len (array )<=1{_d :=make ([]float64 ,len (array ));copy (_d ,array );return _d ;};_a :=len (array )/2;_dd :=MergeSort (array [:_a ]);_fc :=MergeSort (array [_a :]);_fg :=make ([]float64 ,len (array ));_e :=0;_b :=0;_eb :=0;for _b < len (_dd )&&_eb < len (_fc ){if _dd [_b ]<=_fc [_eb ]{_fg [_e ]=_dd [_b ];_b ++;}else {_fg [_e ]=_fc [_eb ];_eb ++;};_e ++;};for _b < len (_dd ){_fg [_e ]=_dd [_b ];_b ++;_e ++;};for _eb < len (_fc ){_fg [_e ]=_fc [_eb ];_eb ++;_e ++;};return _fg ;};