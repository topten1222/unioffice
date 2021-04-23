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

package core_properties ;import (_ce "encoding/xml";_e "fmt";_f "github.com/unidoc/unioffice";_eb "time";);type CT_CoreProperties struct{Category *string ;ContentStatus *string ;Created *_f .XSDAny ;Creator *_f .XSDAny ;Description *_f .XSDAny ;Identifier *_f .XSDAny ;Keywords *CT_Keywords ;Language *_f .XSDAny ;LastModifiedBy *string ;LastPrinted *_eb .Time ;Modified *_f .XSDAny ;Revision *string ;Subject *_f .XSDAny ;Title *_f .XSDAny ;Version *string ;};

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_cg *CT_Keyword )ValidateWithPath (path string )error {return nil };

// Validate validates the CT_Keywords and its children
func (_aba *CT_Keywords )Validate ()error {return _aba .ValidateWithPath ("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073");};func (_ad *CT_Keywords )UnmarshalXML (d *_ce .Decoder ,start _ce .StartElement )error {for _ ,_gca :=range start .Attr {if _gca .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_gca .Name .Local =="\u006c\u0061\u006e\u0067"{_fbd ,_efc :=_gca .Value ,error (nil );if _efc !=nil {return _efc ;};_ad .LangAttr =&_fbd ;continue ;};};_ebb :for {_de ,_cca :=d .Token ();if _cca !=nil {return _cca ;};switch _ggc :=_de .(type ){case _ce .StartElement :switch _ggc .Name {case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076\u0061\u006cu\u0065"}:_geg :=NewCT_Keyword ();if _ebc :=d .DecodeElement (_geg ,&_ggc );_ebc !=nil {return _ebc ;};_ad .Value =append (_ad .Value ,_geg );default:_f .Log ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076",_ggc .Name );if _gcc :=d .Skip ();_gcc !=nil {return _gcc ;};};case _ce .EndElement :break _ebb ;case _ce .CharData :};};return nil ;};func (_afc *CoreProperties )UnmarshalXML (d *_ce .Decoder ,start _ce .StartElement )error {_afc .CT_CoreProperties =*NewCT_CoreProperties ();_gce :for {_ecg ,_fga :=d .Token ();if _fga !=nil {return _fga ;};switch _cgg :=_ecg .(type ){case _ce .StartElement :switch _cgg .Name {case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_afc .Category =new (string );if _bcg :=d .DecodeElement (_afc .Category ,&_cgg );_bcg !=nil {return _bcg ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_afc .ContentStatus =new (string );if _cdf :=d .DecodeElement (_afc .ContentStatus ,&_cgg );_cdf !=nil {return _cdf ;};case _ce .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_afc .Created =new (_f .XSDAny );if _fae :=d .DecodeElement (_afc .Created ,&_cgg );_fae !=nil {return _fae ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_afc .Creator =new (_f .XSDAny );if _fea :=d .DecodeElement (_afc .Creator ,&_cgg );_fea !=nil {return _fea ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_afc .Description =new (_f .XSDAny );if _dde :=d .DecodeElement (_afc .Description ,&_cgg );_dde !=nil {return _dde ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_afc .Identifier =new (_f .XSDAny );if _dge :=d .DecodeElement (_afc .Identifier ,&_cgg );_dge !=nil {return _dge ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_afc .Keywords =NewCT_Keywords ();if _ca :=d .DecodeElement (_afc .Keywords ,&_cgg );_ca !=nil {return _ca ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_afc .Language =new (_f .XSDAny );if _face :=d .DecodeElement (_afc .Language ,&_cgg );_face !=nil {return _face ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_afc .LastModifiedBy =new (string );if _gee :=d .DecodeElement (_afc .LastModifiedBy ,&_cgg );_gee !=nil {return _gee ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_afc .LastPrinted =new (_eb .Time );if _fdg :=d .DecodeElement (_afc .LastPrinted ,&_cgg );_fdg !=nil {return _fdg ;};case _ce .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_afc .Modified =new (_f .XSDAny );if _edg :=d .DecodeElement (_afc .Modified ,&_cgg );_edg !=nil {return _edg ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_afc .Revision =new (string );if _ada :=d .DecodeElement (_afc .Revision ,&_cgg );_ada !=nil {return _ada ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_afc .Subject =new (_f .XSDAny );if _ag :=d .DecodeElement (_afc .Subject ,&_cgg );_ag !=nil {return _ag ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_afc .Title =new (_f .XSDAny );if _ee :=d .DecodeElement (_afc .Title ,&_cgg );_ee !=nil {return _ee ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_afc .Version =new (string );if _bbb :=d .DecodeElement (_afc .Version ,&_cgg );_bbb !=nil {return _bbb ;};default:_f .Log ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076",_cgg .Name );if _ceeb :=d .Skip ();_ceeb !=nil {return _ceeb ;};};case _ce .EndElement :break _gce ;case _ce .CharData :};};return nil ;};func NewCT_CoreProperties ()*CT_CoreProperties {_ed :=&CT_CoreProperties {};return _ed };func (_d *CT_CoreProperties )MarshalXML (e *_ce .Encoder ,start _ce .StartElement )error {e .EncodeToken (start );if _d .Category !=nil {_ea :=_ce .StartElement {Name :_ce .Name {Local :"c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}};_f .AddPreserveSpaceAttr (&_ea ,*_d .Category );e .EncodeElement (_d .Category ,_ea );};if _d .ContentStatus !=nil {_fe :=_ce .StartElement {Name :_ce .Name {Local :"\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}};_f .AddPreserveSpaceAttr (&_fe ,*_d .ContentStatus );e .EncodeElement (_d .ContentStatus ,_fe );};if _d .Created !=nil {_dg :=_ce .StartElement {Name :_ce .Name {Local :"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}};e .EncodeElement (_d .Created ,_dg );};if _d .Creator !=nil {_df :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};e .EncodeElement (_d .Creator ,_df );};if _d .Description !=nil {_eg :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};e .EncodeElement (_d .Description ,_eg );};if _d .Identifier !=nil {_fg :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}};e .EncodeElement (_d .Identifier ,_fg );};if _d .Keywords !=nil {_g :=_ce .StartElement {Name :_ce .Name {Local :"c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}};e .EncodeElement (_d .Keywords ,_g );};if _d .Language !=nil {_fb :=_ce .StartElement {Name :_ce .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};e .EncodeElement (_d .Language ,_fb );};if _d .LastModifiedBy !=nil {_ec :=_ce .StartElement {Name :_ce .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}};_f .AddPreserveSpaceAttr (&_ec ,*_d .LastModifiedBy );e .EncodeElement (_d .LastModifiedBy ,_ec );};if _d .LastPrinted !=nil {_db :=_ce .StartElement {Name :_ce .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}};e .EncodeElement (_d .LastPrinted ,_db );};if _d .Modified !=nil {_ef :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}};e .EncodeElement (_d .Modified ,_ef );};if _d .Revision !=nil {_fgc :=_ce .StartElement {Name :_ce .Name {Local :"c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}};_f .AddPreserveSpaceAttr (&_fgc ,*_d .Revision );e .EncodeElement (_d .Revision ,_fgc );};if _d .Subject !=nil {_b :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}};e .EncodeElement (_d .Subject ,_b );};if _d .Title !=nil {_da :=_ce .StartElement {Name :_ce .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};e .EncodeElement (_d .Title ,_da );};if _d .Version !=nil {_ga :=_ce .StartElement {Name :_ce .Name {Local :"\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}};_f .AddPreserveSpaceAttr (&_ga ,*_d .Version );e .EncodeElement (_d .Version ,_ga );};e .EncodeToken (_ce .EndElement {Name :start .Name });return nil ;};type CT_Keyword struct{LangAttr *string ;Content string ;};func (_dcg *CoreProperties )MarshalXML (e *_ce .Encoder ,start _ce .StartElement )error {start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"});start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073";return _dcg .CT_CoreProperties .MarshalXML (e ,start );};func NewCT_Keywords ()*CT_Keywords {_fac :=&CT_Keywords {};return _fac };func NewCT_Keyword ()*CT_Keyword {_cd :=&CT_Keyword {};return _cd };func NewCoreProperties ()*CoreProperties {_gdb :=&CoreProperties {};_gdb .CT_CoreProperties =*NewCT_CoreProperties ();return _gdb ;};

// Validate validates the CT_CoreProperties and its children
func (_gc *CT_CoreProperties )Validate ()error {return _gc .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};func (_edb *CT_Keyword )UnmarshalXML (d *_ce .Decoder ,start _ce .StartElement )error {for _ ,_cc :=range start .Attr {if _cc .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_cc .Name .Local =="\u006c\u0061\u006e\u0067"{_ff ,_bab :=_cc .Value ,error (nil );if _bab !=nil {return _bab ;};_edb .LangAttr =&_ff ;continue ;};};for {_gf ,_dgd :=d .Token ();if _dgd !=nil {return _e .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073",_dgd );};if _cdb ,_bf :=_gf .(_ce .CharData );_bf {_edb .Content =string (_cdb );};if _dfb ,_fdd :=_gf .(_ce .EndElement );_fdd &&_dfb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CoreProperties and its children
func (_bde *CoreProperties )Validate ()error {return _bde .ValidateWithPath ("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_dfg *CoreProperties )ValidateWithPath (path string )error {if _gda :=_dfg .CT_CoreProperties .ValidateWithPath (path );_gda !=nil {return _gda ;};return nil ;};func (_a *CT_CoreProperties )UnmarshalXML (d *_ce .Decoder ,start _ce .StartElement )error {_gg :for {_ba ,_bb :=d .Token ();if _bb !=nil {return _bb ;};switch _fc :=_ba .(type ){case _ce .StartElement :switch _fc .Name {case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_a .Category =new (string );if _fa :=d .DecodeElement (_a .Category ,&_fc );_fa !=nil {return _fa ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_a .ContentStatus =new (string );if _bd :=d .DecodeElement (_a .ContentStatus ,&_fc );_bd !=nil {return _bd ;};case _ce .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_a .Created =new (_f .XSDAny );if _cea :=d .DecodeElement (_a .Created ,&_fc );_cea !=nil {return _cea ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_a .Creator =new (_f .XSDAny );if _dc :=d .DecodeElement (_a .Creator ,&_fc );_dc !=nil {return _dc ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_a .Description =new (_f .XSDAny );if _eca :=d .DecodeElement (_a .Description ,&_fc );_eca !=nil {return _eca ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_a .Identifier =new (_f .XSDAny );if _bc :=d .DecodeElement (_a .Identifier ,&_fc );_bc !=nil {return _bc ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_a .Keywords =NewCT_Keywords ();if _cee :=d .DecodeElement (_a .Keywords ,&_fc );_cee !=nil {return _cee ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_a .Language =new (_f .XSDAny );if _ebf :=d .DecodeElement (_a .Language ,&_fc );_ebf !=nil {return _ebf ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_a .LastModifiedBy =new (string );if _cb :=d .DecodeElement (_a .LastModifiedBy ,&_fc );_cb !=nil {return _cb ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_a .LastPrinted =new (_eb .Time );if _ge :=d .DecodeElement (_a .LastPrinted ,&_fc );_ge !=nil {return _ge ;};case _ce .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_a .Modified =new (_f .XSDAny );if _ab :=d .DecodeElement (_a .Modified ,&_fc );_ab !=nil {return _ab ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_a .Revision =new (string );if _cf :=d .DecodeElement (_a .Revision ,&_fc );_cf !=nil {return _cf ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_a .Subject =new (_f .XSDAny );if _fgf :=d .DecodeElement (_a .Subject ,&_fc );_fgf !=nil {return _fgf ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_a .Title =new (_f .XSDAny );if _fd :=d .DecodeElement (_a .Title ,&_fc );_fd !=nil {return _fd ;};case _ce .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_a .Version =new (string );if _edf :=d .DecodeElement (_a .Version ,&_fc );_edf !=nil {return _edf ;};default:_f .Log ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076",_fc .Name );if _faa :=d .Skip ();_faa !=nil {return _faa ;};};case _ce .EndElement :break _gg ;case _ce .CharData :};};return nil ;};

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_bag *CT_CoreProperties )ValidateWithPath (path string )error {if _bag .Keywords !=nil {if _gb :=_bag .Keywords .ValidateWithPath (path +"\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s");_gb !=nil {return _gb ;};};return nil ;};func (_gd *CT_Keywords )MarshalXML (e *_ce .Encoder ,start _ce .StartElement )error {if _gd .LangAttr !=nil {start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_e .Sprintf ("\u0025\u0076",*_gd .LangAttr )});};e .EncodeToken (start );if _gd .Value !=nil {_af :=_ce .StartElement {Name :_ce .Name {Local :"\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}};for _ ,_fff :=range _gd .Value {e .EncodeElement (_fff ,_af );};};e .EncodeToken (_ce .EndElement {Name :start .Name });return nil ;};

// Validate validates the CT_Keyword and its children
func (_aa *CT_Keyword )Validate ()error {return _aa .ValidateWithPath ("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064");};type CT_Keywords struct{LangAttr *string ;Value []*CT_Keyword ;};

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_ecf *CT_Keywords )ValidateWithPath (path string )error {for _be ,_gbe :=range _ecf .Value {if _cbc :=_gbe .ValidateWithPath (_e .Sprintf ("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d",path ,_be ));_cbc !=nil {return _cbc ;};};return nil ;};type CoreProperties struct{CT_CoreProperties };func (_dgb *CT_Keyword )MarshalXML (e *_ce .Encoder ,start _ce .StartElement )error {if _dgb .LangAttr !=nil {start .Attr =append (start .Attr ,_ce .Attr {Name :_ce .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_e .Sprintf ("\u0025\u0076",*_dgb .LangAttr )});};e .EncodeElement (_dgb .Content ,start );e .EncodeToken (_ce .EndElement {Name :start .Name });return nil ;};func init (){_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCT_CoreProperties );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",NewCT_Keywords );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064",NewCT_Keyword );_f .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCoreProperties );};