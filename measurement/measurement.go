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

package measurement ;

// FromEMU converts int64 EMU units to float64 distance units.
func FromEMU (emu int64 )float64 {return float64 (emu )/914400*Inch };

// Convert ST_Angle to degree.
func DegreeFromSTAngle (stAngle int32 )float64 {return -float64 (stAngle )/_c };

// ToEMU converts float64 distance units to int64 EMU.
func ToEMU (m float64 )int64 {return int64 (914400.0/Inch *m )};const _c =60000;

// Distance represents a distance and is automatically converted
// to the units needed internally in the various ECMA 376 formats.
type Distance float64 ;const (Zero Distance =0;Point =1;Pixel72 =1.0/72.0*Inch ;Pixel96 =1.0/96.0*Inch ;HalfPoint =1.0/2.0*Point ;Character =7*Point ;Millimeter =2.83465*Point ;Centimeter =10*Millimeter ;Inch =72*Point ;Foot =12*Inch ;Twips =1.0/20.0*Point ;
EMU =1.0/914400.0*Inch ;HundredthPoint =1/100.0;Dxa =Twips ;);