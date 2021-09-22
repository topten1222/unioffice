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

package elements ;import (_gg "encoding/xml";_a "fmt";_b "github.com/unidoc/unioffice";);func (_ggb *SimpleLiteral )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {for {_deg ,_gff :=d .Token ();if _gff !=nil {return _a .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_gff );};if _fc ,_ebag :=_deg .(_gg .EndElement );_ebag &&_fc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_bb *Any )ValidateWithPath (path string )error {if _ff :=_bb .SimpleLiteral .ValidateWithPath (path );_ff !=nil {return _ff ;};return nil ;};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_bec *ElementContainer )ValidateWithPath (path string )error {for _bc ,_gf :=range _bec .Choice {if _add :=_gf .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_bc ));_add !=nil {return _add ;};};return nil ;};func (_af *ElementContainer )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_fd :for {_cd ,_ef :=d .Token ();if _ef !=nil {return _ef ;};switch _fb :=_cd .(type ){case _gg .StartElement :switch _fb .Name {case _gg .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_eb :=NewElementsGroupChoice ();if _eba :=d .DecodeElement (&_eb .Any ,&_fb );_eba !=nil {return _eba ;};_af .Choice =append (_af .Choice ,_eb );default:_b .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_fb .Name );if _ggf :=d .Skip ();_ggf !=nil {return _ggf ;};};case _gg .EndElement :break _fd ;case _gg .CharData :};};return nil ;};

// Validate validates the SimpleLiteral and its children
func (_ded *SimpleLiteral )Validate ()error {return _ded .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};func (_ggg *ElementsGroup )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_gfed :for {_aa ,_aae :=d .Token ();if _aae !=nil {return _aae ;};switch _ca :=_aa .(type ){case _gg .StartElement :switch _ca .Name {case _gg .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_ggfc :=NewElementsGroupChoice ();if _ab :=d .DecodeElement (&_ggfc .Any ,&_ca );_ab !=nil {return _ab ;};_ggg .Choice =append (_ggg .Choice ,_ggfc );default:_b .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ca .Name );if _bfb :=d .Skip ();_bfb !=nil {return _bfb ;};};case _gg .EndElement :break _gfed ;case _gg .CharData :};};return nil ;};type ElementsGroupChoice struct{Any []*Any ;};

// Validate validates the ElementContainer and its children
func (_ae *ElementContainer )Validate ()error {return _ae .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_gbc *SimpleLiteral )ValidateWithPath (path string )error {return nil };func NewAny ()*Any {_ba :=&Any {};_ba .SimpleLiteral =*NewSimpleLiteral ();return _ba };type SimpleLiteral struct{};func NewSimpleLiteral ()*SimpleLiteral {_ffgb :=&SimpleLiteral {};return _ffgb };

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_de *ElementsGroupChoice )ValidateWithPath (path string )error {for _agg ,_ecg :=range _de .Any {if _ceb :=_ecg .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_agg ));_ceb !=nil {return _ceb ;};};return nil ;};

// Validate validates the ElementsGroup and its children
func (_gb *ElementsGroup )Validate ()error {return _gb .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};func NewElementsGroupChoice ()*ElementsGroupChoice {_bda :=&ElementsGroupChoice {};return _bda };func NewElementsGroup ()*ElementsGroup {_dc :=&ElementsGroup {};return _dc };func (_bg *ElementsGroup )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {if _bg .Choice !=nil {for _ ,_gfe :=range _bg .Choice {_gfe .MarshalXML (e ,_gg .StartElement {});};};return nil ;};func NewElementContainer ()*ElementContainer {_gca :=&ElementContainer {};return _gca };

// Validate validates the ElementsGroupChoice and its children
func (_fda *ElementsGroupChoice )Validate ()error {return _fda .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};func (_be *ElementContainer )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _be .Choice !=nil {for _ ,_bd :=range _be .Choice {_bd .MarshalXML (e ,_gg .StartElement {});};};e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};

// Validate validates the Any and its children
func (_fg *Any )Validate ()error {return _fg .ValidateWithPath ("\u0041\u006e\u0079")};type Any struct{SimpleLiteral };func (_db *ElementsGroupChoice )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_ce :for {_aef ,_afe :=d .Token ();if _afe !=nil {return _afe ;};switch _ecf :=_aef .(type ){case _gg .StartElement :switch _ecf .Name {case _gg .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fff :=NewAny ();if _ac :=d .DecodeElement (_fff ,&_ecf );_ac !=nil {return _ac ;};_db .Any =append (_db .Any ,_fff );default:_b .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_ecf .Name );if _bga :=d .Skip ();_bga !=nil {return _bga ;};};case _gg .EndElement :break _ce ;case _gg .CharData :};};return nil ;};func (_ffg *ElementsGroupChoice )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {if _ffg .Any !=nil {_ec :=_gg .StartElement {Name :_gg .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_afd :=range _ffg .Any {e .EncodeElement (_afd ,_ec );};};return nil ;};type ElementContainer struct{Choice []*ElementsGroupChoice ;};func (_dg *SimpleLiteral )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {e .EncodeToken (start );e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};func (_d *Any )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_d .SimpleLiteral =*NewSimpleLiteral ();for {_ad ,_gc :=d .Token ();if _gc !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_gc );};if _e ,_f :=_ad .(_gg .EndElement );_f &&_e .Name ==start .Name {break ;};};return nil ;};func (_c *Any )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {return _c .SimpleLiteral .MarshalXML (e ,start );};

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_cc *ElementsGroup )ValidateWithPath (path string )error {for _bdg ,_efd :=range _cc .Choice {if _bbg :=_efd .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_bdg ));_bbg !=nil {return _bbg ;};};return nil ;};type ElementsGroup struct{Choice []*ElementsGroupChoice ;};func init (){_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );};