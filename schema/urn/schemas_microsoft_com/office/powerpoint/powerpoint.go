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

package powerpoint ;import (_c "encoding/xml";_b "fmt";_cd "github.com/topten1222/unioffice";);

// Validate validates the Textdata and its children
func (_dgd *Textdata )Validate ()error {return _dgd .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};func (_bd *CT_Empty )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fc ,_dd :=d .Token ();if _dd !=nil {return _b .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_dd );
};if _fa ,_a :=_fc .(_c .EndElement );_a &&_fa .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Empty and its children
func (_aa *CT_Empty )Validate ()error {return _aa .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};type Iscomment struct{CT_Empty };

// Validate validates the CT_Rel and its children
func (_fd *CT_Rel )Validate ()error {return _fd .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func NewIscomment ()*Iscomment {_gc :=&Iscomment {};_gc .CT_Empty =*NewCT_Empty ();return _gc };func (_be *Textdata )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _be .CT_Rel .MarshalXML (e ,start );};type CT_Empty struct{};type CT_Rel struct{IdAttr *string ;};

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_dg *CT_Empty )ValidateWithPath (path string )error {return nil };func NewCT_Rel ()*CT_Rel {_ba :=&CT_Rel {};return _ba };func (_e *CT_Empty )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};func NewCT_Empty ()*CT_Empty {_f :=&CT_Empty {};return _f };

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_bb *CT_Rel )ValidateWithPath (path string )error {return nil };type Textdata struct{CT_Rel };

// Validate validates the Iscomment and its children
func (_ac *Iscomment )Validate ()error {return _ac .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func (_dc *Iscomment )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_dc .CT_Empty =*NewCT_Empty ();for {_fce ,_cbd :=d .Token ();
if _cbd !=nil {return _b .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_cbd );};if _faf ,_bf :=_fce .(_c .EndElement );_bf &&_faf .Name ==start .Name {break ;};};return nil ;};func (_eae *Textdata )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_eae .CT_Rel =*NewCT_Rel ();
for _ ,_dcg :=range start .Attr {if _dcg .Name .Local =="\u0069\u0064"{_bad ,_eb :=_dcg .Value ,error (nil );if _eb !=nil {return _eb ;};_eae .IdAttr =&_bad ;continue ;};};for {_ace ,_baf :=d .Token ();if _baf !=nil {return _b .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_baf );
};if _bbg ,_cg :=_ace .(_c .EndElement );_cg &&_bbg .Name ==start .Name {break ;};};return nil ;};func (_fdd *Iscomment )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _fdd .CT_Empty .MarshalXML (e ,start );};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_bbe *Iscomment )ValidateWithPath (path string )error {if _fgf :=_bbe .CT_Empty .ValidateWithPath (path );_fgf !=nil {return _fgf ;};return nil ;};func NewTextdata ()*Textdata {_dcc :=&Textdata {};_dcc .CT_Rel =*NewCT_Rel ();return _dcc };func (_ea *CT_Rel )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _ea .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0069\u0064"},Value :_b .Sprintf ("\u0025\u0076",*_ea .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_ab *CT_Rel )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_cc :=range start .Attr {if _cc .Name .Local =="\u0069\u0064"{_gg ,_cb :=_cc .Value ,error (nil );
if _cb !=nil {return _cb ;};_ab .IdAttr =&_gg ;continue ;};};for {_fb ,_ag :=d .Token ();if _ag !=nil {return _b .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_ag );};if _fg ,_bag :=_fb .(_c .EndElement );
_bag &&_fg .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_ef *Textdata )ValidateWithPath (path string )error {if _gb :=_ef .CT_Rel .ValidateWithPath (path );_gb !=nil {return _gb ;};return nil ;};func init (){_cd .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_cd .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_cd .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_cd .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};