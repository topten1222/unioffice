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

package lockedCanvas ;import (_e "encoding/xml";_ac "fmt";_b "github.com/unidoc/unioffice";_g "github.com/unidoc/unioffice/schema/soo/dml";);func (_bd *LockedCanvas )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"});start .Attr =append (start .Attr ,_e .Attr {Name :_e .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073";return _bd .CT_GvmlGroupShape .MarshalXML (e ,start );};func NewLockedCanvas ()*LockedCanvas {_c :=&LockedCanvas {};_c .CT_GvmlGroupShape =*_g .NewCT_GvmlGroupShape ();return _c ;};

// Validate validates the LockedCanvas and its children
func (_bb *LockedCanvas )Validate ()error {return _bb .ValidateWithPath ("\u004c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073");};type LockedCanvas struct{_g .CT_GvmlGroupShape };func (_gg *LockedCanvas )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gg .CT_GvmlGroupShape =*_g .NewCT_GvmlGroupShape ();for {_d ,_bdc :=d .Token ();if _bdc !=nil {return _ac .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u004c\u006f\u0063\u006b\u0065d\u0043\u0061\u006e\u0076\u0061\u0073\u003a\u0020\u0025\u0073",_bdc );};if _bdcb ,_df :=_d .(_e .EndElement );_df &&_bdcb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the LockedCanvas and its children, prefixing error messages with path
func (_f *LockedCanvas )ValidateWithPath (path string )error {if _ca :=_f .CT_GvmlGroupShape .ValidateWithPath (path );_ca !=nil {return _ca ;};return nil ;};func init (){_b .RegisterConstructor ("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073","\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073",NewLockedCanvas );};