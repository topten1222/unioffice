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

package elements ;import (_e "encoding/xml";_g "fmt";_d "github.com/unidoc/unioffice";_c "github.com/unidoc/unioffice/common/logger";);func NewSimpleLiteral ()*SimpleLiteral {_deb :=&SimpleLiteral {};return _deb };func (_cda *ElementsGroup )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_cg :for {_ab ,_cdd :=d .Token ();if _cdd !=nil {return _cdd ;};switch _dfa :=_ab .(type ){case _e .StartElement :switch _dfa .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_dfb :=NewElementsGroupChoice ();if _ee :=d .DecodeElement (&_dfb .Any ,&_dfa );_ee !=nil {return _ee ;};_cda .Choice =append (_cda .Choice ,_dfb );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_dfa .Name );if _de :=d .Skip ();_de !=nil {return _de ;};};case _e .EndElement :break _cg ;case _e .CharData :};};return nil ;};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_df *ElementContainer )ValidateWithPath (path string )error {for _age ,_efa :=range _df .Choice {if _cc :=_efa .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_age ));_cc !=nil {return _cc ;};};return nil ;};

// Validate validates the ElementsGroupChoice and its children
func (_daf *ElementsGroupChoice )Validate ()error {return _daf .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};func (_eg *Any )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {return _eg .SimpleLiteral .MarshalXML (e ,start );};func (_adf *SimpleLiteral )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {e .EncodeToken (start );e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};

// Validate validates the ElementContainer and its children
func (_gca *ElementContainer )Validate ()error {return _gca .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_ded *ElementsGroup )ValidateWithPath (path string )error {for _ff ,_cgg :=range _ded .Choice {if _fcc :=_cgg .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_ff ));_fcc !=nil {return _fcc ;};};return nil ;};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};func NewElementContainer ()*ElementContainer {_ebd :=&ElementContainer {};return _ebd };func (_egc *SimpleLiteral )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {for {_ga ,_ebg :=d .Token ();if _ebg !=nil {return _g .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_ebg );};if _bf ,_fg :=_ga .(_e .EndElement );_fg &&_bf .Name ==start .Name {break ;};};return nil ;};func (_dbe *ElementsGroupChoice )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _dbe .Any !=nil {_af :=_e .StartElement {Name :_e .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_bb :=range _dbe .Any {e .EncodeElement (_bb ,_af );};};return nil ;};type Any struct{SimpleLiteral };func NewAny ()*Any {_eb :=&Any {};_eb .SimpleLiteral =*NewSimpleLiteral ();return _eb };func (_efb *ElementsGroup )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {if _efb .Choice !=nil {for _ ,_fd :=range _efb .Choice {_fd .MarshalXML (e ,_e .StartElement {});};};return nil ;};

// Validate validates the ElementsGroup and its children
func (_ac *ElementsGroup )Validate ()error {return _ac .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};func (_cd *ElementContainer )MarshalXML (e *_e .Encoder ,start _e .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _cd .Choice !=nil {for _ ,_da :=range _cd .Choice {_da .MarshalXML (e ,_e .StartElement {});};};e .EncodeToken (_e .EndElement {Name :start .Name });return nil ;};type ElementsGroupChoice struct{Any []*Any ;};func (_aa *Any )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_aa .SimpleLiteral =*NewSimpleLiteral ();for {_ag ,_f :=d .Token ();if _f !=nil {return _g .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_f );};if _db ,_ef :=_ag .(_e .EndElement );_ef &&_db .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_ca *Any )ValidateWithPath (path string )error {if _fe :=_ca .SimpleLiteral .ValidateWithPath (path );_fe !=nil {return _fe ;};return nil ;};

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_fcf *ElementsGroupChoice )ValidateWithPath (path string )error {for _eeb ,_ge :=range _fcf .Any {if _ad :=_ge .ValidateWithPath (_g .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_eeb ));_ad !=nil {return _ad ;};};return nil ;};func (_be *ElementsGroupChoice )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_dc :for {_efdg ,_gfb :=d .Token ();if _gfb !=nil {return _gfb ;};switch _ccg :=_efdg .(type ){case _e .StartElement :switch _ccg .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ddd :=NewAny ();if _acb :=d .DecodeElement (_ddd ,&_ccg );_acb !=nil {return _acb ;};_be .Any =append (_be .Any ,_ddd );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_ccg .Name );if _dad :=d .Skip ();_dad !=nil {return _dad ;};};case _e .EndElement :break _dc ;case _e .CharData :};};return nil ;};func NewElementsGroupChoice ()*ElementsGroupChoice {_gf :=&ElementsGroupChoice {};return _gf };

// Validate validates the SimpleLiteral and its children
func (_cab *SimpleLiteral )Validate ()error {return _cab .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};func (_fc *ElementContainer )UnmarshalXML (d *_e .Decoder ,start _e .StartElement )error {_fb :for {_gc ,_dg :=d .Token ();if _dg !=nil {return _dg ;};switch _dd :=_gc .(type ){case _e .StartElement :switch _dd .Name {case _e .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fcg :=NewElementsGroupChoice ();if _b :=d .DecodeElement (&_fcg .Any ,&_dd );_b !=nil {return _b ;};_fc .Choice =append (_fc .Choice ,_fcg );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_dd .Name );if _fef :=d .Skip ();_fef !=nil {return _fef ;};};case _e .EndElement :break _fb ;case _e .CharData :};};return nil ;};func NewElementsGroup ()*ElementsGroup {_fec :=&ElementsGroup {};return _fec };

// Validate validates the Any and its children
func (_efd *Any )Validate ()error {return _efd .ValidateWithPath ("\u0041\u006e\u0079")};type SimpleLiteral struct{};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_fbb *SimpleLiteral )ValidateWithPath (path string )error {return nil };type ElementContainer struct{Choice []*ElementsGroupChoice ;};func init (){_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_d .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );};