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

package core_properties ;import (_d "encoding/xml";_a "fmt";_ad "github.com/unidoc/unioffice";_b "github.com/unidoc/unioffice/common/logger";_f "time";);

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_ebd *CoreProperties )ValidateWithPath (path string )error {if _cce :=_ebd .CT_CoreProperties .ValidateWithPath (path );_cce !=nil {return _cce ;};return nil ;};func (_cc *CT_Keywords )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _cc .LangAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_a .Sprintf ("\u0025\u0076",*_cc .LangAttr )});};e .EncodeToken (start );if _cc .Value !=nil {_fa :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}};for _ ,_fcf :=range _cc .Value {e .EncodeElement (_fcf ,_fa );};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func NewCT_Keywords ()*CT_Keywords {_ecd :=&CT_Keywords {};return _ecd };func (_abc *CT_Keyword )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_eaa :=range start .Attr {if _eaa .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_eaa .Name .Local =="\u006c\u0061\u006e\u0067"{_fbc ,_gbe :=_eaa .Value ,error (nil );if _gbe !=nil {return _gbe ;};_abc .LangAttr =&_fbc ;continue ;};};for {_baa ,_gba :=d .Token ();if _gba !=nil {return _a .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073",_gba );};if _bb ,_cd :=_baa .(_d .CharData );_cd {_abc .Content =string (_bb );};if _fee ,_cge :=_baa .(_d .EndElement );_cge &&_fee .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_bf *CT_CoreProperties )ValidateWithPath (path string )error {if _bf .Keywords !=nil {if _efc :=_bf .Keywords .ValidateWithPath (path +"\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s");_efc !=nil {return _efc ;};};return nil ;};func (_ba *CT_CoreProperties )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );if _ba .Category !=nil {_fb :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}};_ad .AddPreserveSpaceAttr (&_fb ,*_ba .Category );e .EncodeElement (_ba .Category ,_fb );};if _ba .ContentStatus !=nil {_df :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}};_ad .AddPreserveSpaceAttr (&_df ,*_ba .ContentStatus );e .EncodeElement (_ba .ContentStatus ,_df );};if _ba .Created !=nil {_c :=_d .StartElement {Name :_d .Name {Local :"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}};e .EncodeElement (_ba .Created ,_c );};if _ba .Creator !=nil {_bg :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};e .EncodeElement (_ba .Creator ,_bg );};if _ba .Description !=nil {_e :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};e .EncodeElement (_ba .Description ,_e );};if _ba .Identifier !=nil {_dg :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}};e .EncodeElement (_ba .Identifier ,_dg );};if _ba .Keywords !=nil {_cb :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}};e .EncodeElement (_ba .Keywords ,_cb );};if _ba .Language !=nil {_fd :=_d .StartElement {Name :_d .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};e .EncodeElement (_ba .Language ,_fd );};if _ba .LastModifiedBy !=nil {_eb :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}};_ad .AddPreserveSpaceAttr (&_eb ,*_ba .LastModifiedBy );e .EncodeElement (_ba .LastModifiedBy ,_eb );};if _ba .LastPrinted !=nil {_ce :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}};e .EncodeElement (_ba .LastPrinted ,_ce );};if _ba .Modified !=nil {_eg :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}};e .EncodeElement (_ba .Modified ,_eg );};if _ba .Revision !=nil {_ceb :=_d .StartElement {Name :_d .Name {Local :"c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}};_ad .AddPreserveSpaceAttr (&_ceb ,*_ba .Revision );e .EncodeElement (_ba .Revision ,_ceb );};if _ba .Subject !=nil {_aa :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}};e .EncodeElement (_ba .Subject ,_aa );};if _ba .Title !=nil {_adg :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};e .EncodeElement (_ba .Title ,_adg );};if _ba .Version !=nil {_gc :=_d .StartElement {Name :_d .Name {Local :"\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}};_ad .AddPreserveSpaceAttr (&_gc ,*_ba .Version );e .EncodeElement (_ba .Version ,_gc );};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// Validate validates the CT_CoreProperties and its children
func (_da *CT_CoreProperties )Validate ()error {return _da .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};

// Validate validates the CoreProperties and its children
func (_fge *CoreProperties )Validate ()error {return _fge .ValidateWithPath ("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};func (_fca *CoreProperties )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073";return _fca .CT_CoreProperties .MarshalXML (e ,start );};

// Validate validates the CT_Keywords and its children
func (_bfa *CT_Keywords )Validate ()error {return _bfa .ValidateWithPath ("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073");};func (_dgf *CT_Keyword )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _dgf .LangAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_a .Sprintf ("\u0025\u0076",*_dgf .LangAttr )});};e .EncodeElement (_dgf .Content ,start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};type CoreProperties struct{CT_CoreProperties };func NewCT_CoreProperties ()*CT_CoreProperties {_ge :=&CT_CoreProperties {};return _ge };func NewCoreProperties ()*CoreProperties {_gbg :=&CoreProperties {};_gbg .CT_CoreProperties =*NewCT_CoreProperties ();return _gbg ;};type CT_CoreProperties struct{Category *string ;ContentStatus *string ;Created *_ad .XSDAny ;Creator *_ad .XSDAny ;Description *_ad .XSDAny ;Identifier *_ad .XSDAny ;Keywords *CT_Keywords ;Language *_ad .XSDAny ;LastModifiedBy *string ;LastPrinted *_f .Time ;Modified *_ad .XSDAny ;Revision *string ;Subject *_ad .XSDAny ;Title *_ad .XSDAny ;Version *string ;};

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_ggc *CT_Keywords )ValidateWithPath (path string )error {for _fae ,_ff :=range _ggc .Value {if _bge :=_ff .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d",path ,_fae ));_bge !=nil {return _bge ;};};return nil ;};func (_be *CT_Keywords )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_bgd :=range start .Attr {if _bgd .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_bgd .Name .Local =="\u006c\u0061\u006e\u0067"{_geg ,_gef :=_bgd .Value ,error (nil );if _gef !=nil {return _gef ;};_be .LangAttr =&_geg ;continue ;};};_efe :for {_cebe ,_cbc :=d .Token ();if _cbc !=nil {return _cbc ;};switch _baab :=_cebe .(type ){case _d .StartElement :switch _baab .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076\u0061\u006cu\u0065"}:_ag :=NewCT_Keyword ();if _gcgc :=d .DecodeElement (_ag ,&_baab );_gcgc !=nil {return _gcgc ;};_be .Value =append (_be .Value ,_ag );default:_b .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076",_baab .Name );if _ca :=d .Skip ();_ca !=nil {return _ca ;};};case _d .EndElement :break _efe ;case _d .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_gbaf *CT_Keyword )ValidateWithPath (path string )error {return nil };type CT_Keyword struct{LangAttr *string ;Content string ;};

// Validate validates the CT_Keyword and its children
func (_eae *CT_Keyword )Validate ()error {return _eae .ValidateWithPath ("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064");};type CT_Keywords struct{LangAttr *string ;Value []*CT_Keyword ;};func (_ac *CoreProperties )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_ac .CT_CoreProperties =*NewCT_CoreProperties ();_dee :for {_gegc ,_ee :=d .Token ();if _ee !=nil {return _ee ;};switch _af :=_gegc .(type ){case _d .StartElement :switch _af .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_ac .Category =new (string );if _cef :=d .DecodeElement (_ac .Category ,&_af );_cef !=nil {return _cef ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_ac .ContentStatus =new (string );if _ddc :=d .DecodeElement (_ac .ContentStatus ,&_af );_ddc !=nil {return _ddc ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_ac .Created =new (_ad .XSDAny );if _aag :=d .DecodeElement (_ac .Created ,&_af );_aag !=nil {return _aag ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_ac .Creator =new (_ad .XSDAny );if _egg :=d .DecodeElement (_ac .Creator ,&_af );_egg !=nil {return _egg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_ac .Description =new (_ad .XSDAny );if _eaec :=d .DecodeElement (_ac .Description ,&_af );_eaec !=nil {return _eaec ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_ac .Identifier =new (_ad .XSDAny );if _bff :=d .DecodeElement (_ac .Identifier ,&_af );_bff !=nil {return _bff ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_ac .Keywords =NewCT_Keywords ();if _ed :=d .DecodeElement (_ac .Keywords ,&_af );_ed !=nil {return _ed ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_ac .Language =new (_ad .XSDAny );if _bab :=d .DecodeElement (_ac .Language ,&_af );_bab !=nil {return _bab ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_ac .LastModifiedBy =new (string );if _acg :=d .DecodeElement (_ac .LastModifiedBy ,&_af );_acg !=nil {return _acg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_ac .LastPrinted =new (_f .Time );if _fdca :=d .DecodeElement (_ac .LastPrinted ,&_af );_fdca !=nil {return _fdca ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_ac .Modified =new (_ad .XSDAny );if _gbed :=d .DecodeElement (_ac .Modified ,&_af );_gbed !=nil {return _gbed ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_ac .Revision =new (string );if _bdg :=d .DecodeElement (_ac .Revision ,&_af );_bdg !=nil {return _bdg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_ac .Subject =new (_ad .XSDAny );if _bfc :=d .DecodeElement (_ac .Subject ,&_af );_bfc !=nil {return _bfc ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_ac .Title =new (_ad .XSDAny );if _cee :=d .DecodeElement (_ac .Title ,&_af );_cee !=nil {return _cee ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_ac .Version =new (string );if _aab :=d .DecodeElement (_ac .Version ,&_af );_aab !=nil {return _aab ;};default:_b .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076",_af .Name );if _gf :=d .Skip ();_gf !=nil {return _gf ;};};case _d .EndElement :break _dee ;case _d .CharData :};};return nil ;};func (_de *CT_CoreProperties )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fdg :for {_ae ,_ec :=d .Token ();if _ec !=nil {return _ec ;};switch _ea :=_ae .(type ){case _d .StartElement :switch _ea .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_de .Category =new (string );if _dec :=d .DecodeElement (_de .Category ,&_ea );_dec !=nil {return _dec ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_de .ContentStatus =new (string );if _cg :=d .DecodeElement (_de .ContentStatus ,&_ea );_cg !=nil {return _cg ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_de .Created =new (_ad .XSDAny );if _gcg :=d .DecodeElement (_de .Created ,&_ea );_gcg !=nil {return _gcg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_de .Creator =new (_ad .XSDAny );if _bc :=d .DecodeElement (_de .Creator ,&_ea );_bc !=nil {return _bc ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_de .Description =new (_ad .XSDAny );if _ef :=d .DecodeElement (_de .Description ,&_ea );_ef !=nil {return _ef ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_de .Identifier =new (_ad .XSDAny );if _eca :=d .DecodeElement (_de .Identifier ,&_ea );_eca !=nil {return _eca ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_de .Keywords =NewCT_Keywords ();if _fg :=d .DecodeElement (_de .Keywords ,&_ea );_fg !=nil {return _fg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_de .Language =new (_ad .XSDAny );if _bd :=d .DecodeElement (_de .Language ,&_ea );_bd !=nil {return _bd ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_de .LastModifiedBy =new (string );if _ab :=d .DecodeElement (_de .LastModifiedBy ,&_ea );_ab !=nil {return _ab ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_de .LastPrinted =new (_f .Time );if _gb :=d .DecodeElement (_de .LastPrinted ,&_ea );_gb !=nil {return _gb ;};case _d .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_de .Modified =new (_ad .XSDAny );if _ege :=d .DecodeElement (_de .Modified ,&_ea );_ege !=nil {return _ege ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_de .Revision =new (string );if _fe :=d .DecodeElement (_de .Revision ,&_ea );_fe !=nil {return _fe ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_de .Subject =new (_ad .XSDAny );if _gg :=d .DecodeElement (_de .Subject ,&_ea );_gg !=nil {return _gg ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_de .Title =new (_ad .XSDAny );if _cga :=d .DecodeElement (_de .Title ,&_ea );_cga !=nil {return _cga ;};case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_de .Version =new (string );if _abb :=d .DecodeElement (_de .Version ,&_ea );_abb !=nil {return _abb ;};default:_b .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076",_ea .Name );if _fdc :=d .Skip ();_fdc !=nil {return _fdc ;};};case _d .EndElement :break _fdg ;case _d .CharData :};};return nil ;};func NewCT_Keyword ()*CT_Keyword {_dd :=&CT_Keyword {};return _dd };func init (){_ad .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCT_CoreProperties );_ad .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",NewCT_Keywords );_ad .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064",NewCT_Keyword );_ad .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCoreProperties );};