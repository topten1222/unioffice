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

package elements ;import (_e "encoding/xml";_be "fmt";_g "github.com/unidoc/unioffice";_c "github.com/unidoc/unioffice/common/logger";);func (_ed *ElementContainer )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_dc :for {_ec ,_cd :=d .Token ();
if _cd !=nil {return _cd ;};switch _fd :=_ec .(type ){case _e .StartElement :switch _fd .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_bb :=NewElementsGroupChoice ();
if _ba :=d .DecodeElement (&_bb .Any ,&_fd );_ba !=nil {return _ba ;};_ed .Choice =append (_ed .Choice ,_bb );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_fd .Name );
if _dcb :=d .Skip ();_dcb !=nil {return _dcb ;};};case _e .EndElement :break _dc ;case _e .CharData :};};return nil ;};func NewElementContainer ()*ElementContainer {_ac :=&ElementContainer {};return _ac };func (_fcf *ElementsGroupChoice )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_fbb :for {_abf ,_ccf :=d .Token ();
if _ccf !=nil {return _ccf ;};switch _ecg :=_abf .(type ){case _e .StartElement :switch _ecg .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fcg :=NewAny ();
if _gea :=d .DecodeElement (_fcg ,&_ecg );_gea !=nil {return _gea ;};_fcf .Any =append (_fcf .Any ,_fcg );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_ecg .Name );
if _gfb :=d .Skip ();_gfb !=nil {return _gfb ;};};case _e .EndElement :break _fbb ;case _e .CharData :};};return nil ;};func NewSimpleLiteral ()*SimpleLiteral {_dg :=&SimpleLiteral {};return _dg };func (_ae *ElementsGroupChoice )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _ae .Any !=nil {_gcc :=_e .StartElement {Name :_e .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_cca :=range _ae .Any {e .EncodeElement (_cca ,_gcc );};};return nil ;};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_fdb *ElementContainer )ValidateWithPath (path string )error {for _ab ,_ee :=range _fdb .Choice {if _cad :=_ee .ValidateWithPath (_be .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_ab ));_cad !=nil {return _cad ;
};};return nil ;};func NewElementsGroupChoice ()*ElementsGroupChoice {_bc :=&ElementsGroupChoice {};return _bc };

// Validate validates the ElementContainer and its children
func (_bgb *ElementContainer )Validate ()error {return _bgb .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};func NewElementsGroup ()*ElementsGroup {_df :=&ElementsGroup {};return _df };func (_f *Any )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return _f .SimpleLiteral .MarshalXML (e ,start );
};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_fcd *SimpleLiteral )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_cab *ElementsGroup )ValidateWithPath (path string )error {for _acf ,_cc :=range _cab .Choice {if _af :=_cc .ValidateWithPath (_be .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_acf ));_af !=nil {return _af ;
};};return nil ;};type ElementsGroupChoice struct{Any []*Any ;};

// Validate validates the ElementsGroupChoice and its children
func (_ccd *ElementsGroupChoice )Validate ()error {return _ccd .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};type SimpleLiteral struct{};

// Validate validates the SimpleLiteral and its children
func (_gd *SimpleLiteral )Validate ()error {return _gd .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_efa *Any )ValidateWithPath (path string )error {if _d :=_efa .SimpleLiteral .ValidateWithPath (path );_d !=nil {return _d ;};return nil ;};type Any struct{SimpleLiteral };func (_gc *ElementsGroup )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _gc .Choice !=nil {for _ ,_abg :=range _gc .Choice {_abg .MarshalXML (e ,_e .StartElement {});
};};return nil ;};

// Validate validates the Any and its children
func (_a *Any )Validate ()error {return _a .ValidateWithPath ("\u0041\u006e\u0079")};func (_dea *ElementsGroup )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_gf :for {_bgg ,_gg :=d .Token ();if _gg !=nil {return _gg ;};switch _ge :=_bgg .(type ){case _e .StartElement :switch _ge .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fa :=NewElementsGroupChoice ();
if _bbb :=d .DecodeElement (&_fa .Any ,&_ge );_bbb !=nil {return _bbb ;};_dea .Choice =append (_dea .Choice ,_fa );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ge .Name );
if _fb :=d .Skip ();_fb !=nil {return _fb ;};};case _e .EndElement :break _gf ;case _e .CharData :};};return nil ;};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};func (_bg *Any )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_bg .SimpleLiteral =*NewSimpleLiteral ();
for {_ea ,_eg :=d .Token ();if _eg !=nil {return _be .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_eg );};if _bd ,_gb :=_ea .(_e .EndElement );_gb &&_bd .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_afe *ElementsGroupChoice )ValidateWithPath (path string )error {for _bbbg ,_eb :=range _afe .Any {if _ddf :=_eb .ValidateWithPath (_be .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_bbbg ));_ddf !=nil {return _ddf ;
};};return nil ;};

// Validate validates the ElementsGroup and its children
func (_fc *ElementsGroup )Validate ()error {return _fc .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};type ElementContainer struct{Choice []*ElementsGroupChoice ;};func (_efg *SimpleLiteral )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for {_fcga ,_cf :=d .Token ();
if _cf !=nil {return _be .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_cf );};if _bcc ,_fdg :=_fcga .(_e .EndElement );_fdg &&_bcc .Name ==start .Name {break ;};};return nil ;
};func (_ca *ElementContainer )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _ca .Choice !=nil {for _ ,_de :=range _ca .Choice {_de .MarshalXML (e ,_e .StartElement {});
};};e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};func NewAny ()*Any {_ef :=&Any {};_ef .SimpleLiteral =*NewSimpleLiteral ();return _ef };func (_ga *SimpleLiteral )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );
e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};func init (){_g .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );
_g .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );
_g .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_g .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );
};