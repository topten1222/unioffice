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

package content_types ;import (_g "encoding/xml";_fb "fmt";_a "github.com/unidoc/unioffice";_c "github.com/unidoc/unioffice/common/logger";_f "regexp";);const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_dfg *Default )ValidateWithPath (path string )error {if _cad :=_dfg .CT_Default .ValidateWithPath (path );_cad !=nil {return _cad ;};return nil ;};

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_acd *Override )ValidateWithPath (path string )error {if _ebc :=_acd .CT_Override .ValidateWithPath (path );_ebc !=nil {return _ebc ;};return nil ;};

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_ggd *Types )ValidateWithPath (path string )error {if _cgc :=_ggd .CT_Types .ValidateWithPath (path );_cgc !=nil {return _cgc ;};return nil ;};

// Validate validates the Default and its children
func (_feg *Default )Validate ()error {return _feg .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};

// Validate validates the Types and its children
func (_edab *Types )Validate ()error {return _edab .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};type CT_Types struct{Default []*Default ;Override []*Override ;};func NewOverride ()*Override {_aeb :=&Override {};_aeb .CT_Override =*NewCT_Override ();return _aeb };type Override struct{CT_Override };type Default struct{CT_Default };func (_fa *CT_Override )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_fb .Sprintf ("\u0025\u0076",_fa .ContentTypeAttr )});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_fb .Sprintf ("\u0025\u0076",_fa .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};type CT_Default struct{ExtensionAttr string ;ContentTypeAttr string ;};type Types struct{CT_Types };func NewCT_Types ()*CT_Types {_cec :=&CT_Types {};return _cec };

// Validate validates the Override and its children
func (_fae *Override )Validate ()error {return _fae .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_cga *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_cga .ExtensionAttr ){return _fb .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_cga .ExtensionAttr );};if !ST_ContentTypePatternRe .MatchString (_cga .ContentTypeAttr ){return _fb .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_cga .ContentTypeAttr );};return nil ;};func (_e *CT_Default )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_e .ExtensionAttr ="\u0078\u006d\u006c";_e .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_d :=range start .Attr {if _d .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_gg ,_fd :=_d .Value ,error (nil );if _fd !=nil {return _fd ;};_e .ExtensionAttr =_gg ;continue ;};if _d .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_cg ,_cf :=_d .Value ,error (nil );if _cf !=nil {return _cf ;};_e .ContentTypeAttr =_cg ;continue ;};};for {_ba ,_fe :=d .Token ();if _fe !=nil {return _fb .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_fe );};if _fc ,_bc :=_ba .(_g .EndElement );_bc &&_fc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_gag *CT_Types )ValidateWithPath (path string )error {for _fba ,_ag :=range _gag .Default {if _cb :=_ag .ValidateWithPath (_fb .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_fba ));_cb !=nil {return _cb ;};};for _ec ,_dbb :=range _gag .Override {if _cdd :=_dbb .ValidateWithPath (_fb .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_ec ));_cdd !=nil {return _cdd ;};};return nil ;};

// Validate validates the CT_Override and its children
func (_aaa *CT_Override )Validate ()error {return _aaa .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};type CT_Override struct{ContentTypeAttr string ;PartNameAttr string ;};func (_df *Default )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {return _df .CT_Default .MarshalXML (e ,start );};func (_ef *CT_Types )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_aec :for {_ga ,_fca :=d .Token ();if _fca !=nil {return _fca ;};switch _gbb :=_ga .(type ){case _g .StartElement :switch _gbb .Name {case _g .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_ffgb :=NewDefault ();if _ede :=d .DecodeElement (_ffgb ,&_gbb );_ede !=nil {return _ede ;};_ef .Default =append (_ef .Default ,_ffgb );case _g .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_ee :=NewOverride ();if _ca :=d .DecodeElement (_ee ,&_gbb );_ca !=nil {return _ca ;};_ef .Override =append (_ef .Override ,_ee );default:_c .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_gbb .Name );if _dbg :=d .Skip ();_dbg !=nil {return _dbg ;};};case _g .EndElement :break _aec ;case _g .CharData :};};return nil ;};func (_af *CT_Override )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_af .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_ac :=range start .Attr {if _ac .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_fbc ,_eb :=_ac .Value ,error (nil );if _eb !=nil {return _eb ;};_af .ContentTypeAttr =_fbc ;continue ;};if _ac .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_ed ,_aa :=_ac .Value ,error (nil );if _aa !=nil {return _aa ;};_af .PartNameAttr =_ed ;continue ;};};for {_cfe ,_fcd :=d .Token ();if _fcd !=nil {return _fb .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_fcd );};if _cc ,_dd :=_cfe .(_g .EndElement );_dd &&_cc .Name ==start .Name {break ;};};return nil ;};const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";func NewCT_Override ()*CT_Override {_ab :=&CT_Override {};_ab .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _ab ;};func (_eda *Override )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {return _eda .CT_Override .MarshalXML (e ,start );};

// Validate validates the CT_Default and its children
func (_fg *CT_Default )Validate ()error {return _fg .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};func (_ad *Types )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_ad .CT_Types =*NewCT_Types ();_cbb :for {_ffa ,_aaae :=d .Token ();if _aaae !=nil {return _aaae ;};switch _agd :=_ffa .(type ){case _g .StartElement :switch _agd .Name {case _g .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_ge :=NewDefault ();if _aab :=d .DecodeElement (_ge ,&_agd );_aab !=nil {return _aab ;};_ad .Default =append (_ad .Default ,_ge );case _g .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_gbc :=NewOverride ();if _bab :=d .DecodeElement (_gbc ,&_agd );_bab !=nil {return _bab ;};_ad .Override =append (_ad .Override ,_gbc );default:_c .Log .Debug ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_agd .Name );if _fce :=d .Skip ();_fce !=nil {return _fce ;};};case _g .EndElement :break _cbb ;case _g .CharData :};};return nil ;};func NewTypes ()*Types {_bb :=&Types {};_bb .CT_Types =*NewCT_Types ();return _bb };func (_gb *CT_Types )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {e .EncodeToken (start );if _gb .Default !=nil {_ffg :=_g .StartElement {Name :_g .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_ae :=range _gb .Default {e .EncodeElement (_ae ,_ffg );};};if _gb .Override !=nil {_db :=_g .StartElement {Name :_g .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_gbf :=range _gb .Override {e .EncodeElement (_gbf ,_db );};};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func NewCT_Default ()*CT_Default {_ce :=&CT_Default {};_ce .ExtensionAttr ="\u0078\u006d\u006c";_ce .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _ce ;};func (_bac *Types )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0054\u0079\u0070e\u0073";return _bac .CT_Types .MarshalXML (e ,start );};var ST_ExtensionPatternRe =_f .MustCompile (ST_ExtensionPattern );var ST_ContentTypePatternRe =_f .MustCompile (ST_ContentTypePattern );

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_ccb *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_ccb .ContentTypeAttr ){return _fb .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_ccb .ContentTypeAttr );};return nil ;};func (_da *Default )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_da .CT_Default =*NewCT_Default ();for _ ,_fcg :=range start .Attr {if _fcg .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_ddd ,_fac :=_fcg .Value ,error (nil );if _fac !=nil {return _fac ;};_da .ExtensionAttr =_ddd ;continue ;};if _fcg .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_dg ,_dag :=_fcg .Value ,error (nil );if _dag !=nil {return _dag ;};_da .ContentTypeAttr =_dg ;continue ;};};for {_aef ,_aff :=d .Token ();if _aff !=nil {return _fb .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_aff );};if _abe ,_aca :=_aef .(_g .EndElement );_aca &&_abe .Name ==start .Name {break ;};};return nil ;};func (_ff *CT_Default )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_fb .Sprintf ("\u0025\u0076",_ff .ExtensionAttr )});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_fb .Sprintf ("\u0025\u0076",_ff .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func (_dcb *Override )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_dcb .CT_Override =*NewCT_Override ();for _ ,_dca :=range start .Attr {if _dca .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_gf ,_cfc :=_dca .Value ,error (nil );if _cfc !=nil {return _cfc ;};_dcb .ContentTypeAttr =_gf ;continue ;};if _dca .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_cef ,_ffc :=_dca .Value ,error (nil );if _ffc !=nil {return _ffc ;};_dcb .PartNameAttr =_cef ;continue ;};};for {_gc ,_gac :=d .Token ();if _gac !=nil {return _fb .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_gac );};if _fda ,_fbf :=_gc .(_g .EndElement );_fbf &&_fda .Name ==start .Name {break ;};};return nil ;};func NewDefault ()*Default {_dc :=&Default {};_dc .CT_Default =*NewCT_Default ();return _dc };

// Validate validates the CT_Types and its children
func (_baa *CT_Types )Validate ()error {return _baa .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};func init (){_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );};