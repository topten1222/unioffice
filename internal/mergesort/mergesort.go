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

package mergesort ;func MergeSort (array []float64 )[]float64 {if len (array )<=1{_c :=make ([]float64 ,len (array ));copy (_c ,array );return _c ;};_f :=len (array )/2;_a :=MergeSort (array [:_f ]);_g :=MergeSort (array [_f :]);_b :=make ([]float64 ,len (array ));_bg :=0;_e :=0;_df :=0;for _e < len (_a )&&_df < len (_g ){if _a [_e ]<=_g [_df ]{_b [_bg ]=_a [_e ];_e ++;}else {_b [_bg ]=_g [_df ];_df ++;};_bg ++;};for _e < len (_a ){_b [_bg ]=_a [_e ];_e ++;_bg ++;};for _df < len (_g ){_b [_bg ]=_g [_df ];_df ++;_bg ++;};return _b ;};