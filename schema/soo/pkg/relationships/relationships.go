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

package relationships ;import (_ee "encoding/xml";_b "fmt";_a "github.com/unidoc/unioffice";);

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_eeb *CT_Relationships )ValidateWithPath (path string )error {for _bfg ,_ed :=range _eeb .Relationship {if _bdd :=_ed .ValidateWithPath (_b .Sprintf ("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d",path ,_bfg ));_bdd !=nil {return _bdd ;};};return nil ;};func (_abb ST_TargetMode )MarshalXMLAttr (name _ee .Name )(_ee .Attr ,error ){_agg :=_ee .Attr {};_agg .Name =name ;switch _abb {case ST_TargetModeUnset :_agg .Value ="";case ST_TargetModeExternal :_agg .Value ="\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case ST_TargetModeInternal :_agg .Value ="\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return _agg ,nil ;};func (_dc *Relationship )MarshalXML (e *_ee .Encoder ,start _ee .StartElement )error {return _dc .CT_Relationship .MarshalXML (e ,start );};func (_fe *ST_TargetMode )UnmarshalXML (d *_ee .Decoder ,start _ee .StartElement )error {_gcg ,_ddeg :=d .Token ();if _ddeg !=nil {return _ddeg ;};if _dad ,_gee :=_gcg .(_ee .EndElement );_gee &&_dad .Name ==start .Name {*_fe =1;return nil ;};if _dda ,_baf :=_gcg .(_ee .CharData );!_baf {return _b .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gcg );}else {switch string (_dda ){case "":*_fe =0;case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_fe =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_fe =2;};};_gcg ,_ddeg =d .Token ();if _ddeg !=nil {return _ddeg ;};if _cb ,_cd :=_gcg .(_ee .EndElement );_cd &&_cb .Name ==start .Name {return nil ;};return _b .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gcg );};func (_cga *CT_Relationships )UnmarshalXML (d *_ee .Decoder ,start _ee .StartElement )error {_gd :for {_gbc ,_ec :=d .Token ();if _ec !=nil {return _ec ;};switch _cfg :=_gbc .(type ){case _ee .StartElement :switch _cfg .Name {case _ee .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_ag :=NewRelationship ();if _bcg :=d .DecodeElement (_ag ,&_cfg );_bcg !=nil {return _bcg ;};_cga .Relationship =append (_cga .Relationship ,_ag );default:_a .Log ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v",_cfg .Name );if _dde :=d .Skip ();_dde !=nil {return _dde ;};};case _ee .EndElement :break _gd ;case _ee .CharData :};};return nil ;};func (_bab ST_TargetMode )Validate ()error {return _bab .ValidateWithPath ("")};const (ST_TargetModeUnset ST_TargetMode =0;ST_TargetModeExternal ST_TargetMode =1;ST_TargetModeInternal ST_TargetMode =2;);

// Validate validates the CT_Relationship and its children
func (_eag *CT_Relationship )Validate ()error {return _eag .ValidateWithPath ("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};

// Validate validates the Relationship and its children
func (_abd *Relationship )Validate ()error {return _abd .ValidateWithPath ("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};func NewCT_Relationship ()*CT_Relationship {_f :=&CT_Relationship {};return _f };func (_dcg *Relationship )UnmarshalXML (d *_ee .Decoder ,start _ee .StartElement )error {_dcg .CT_Relationship =*NewCT_Relationship ();for _ ,_ggb :=range start .Attr {if _ggb .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_dcg .TargetModeAttr .UnmarshalXMLAttr (_ggb );continue ;};if _ggb .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_agb ,_gdg :=_ggb .Value ,error (nil );if _gdg !=nil {return _gdg ;};_dcg .TargetAttr =_agb ;continue ;};if _ggb .Name .Local =="\u0054\u0079\u0070\u0065"{_acf ,_dcc :=_ggb .Value ,error (nil );if _dcc !=nil {return _dcc ;};_dcg .TypeAttr =_acf ;continue ;};if _ggb .Name .Local =="\u0049\u0064"{_bce ,_abe :=_ggb .Value ,error (nil );if _abe !=nil {return _abe ;};_dcg .IdAttr =_bce ;continue ;};};for {_bcc ,_ega :=d .Token ();if _ega !=nil {return _b .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073",_ega );};if _df ,_geb :=_bcc .(_ee .EndElement );_geb &&_df .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Relationships and its children
func (_dg *CT_Relationships )Validate ()error {return _dg .ValidateWithPath ("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};type Relationship struct{CT_Relationship };

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_fgb *Relationship )ValidateWithPath (path string )error {if _bbe :=_fgb .CT_Relationship .ValidateWithPath (path );_bbe !=nil {return _bbe ;};return nil ;};func (_bfa ST_TargetMode )String ()string {switch _bfa {case 0:return "";case 1:return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case 2:return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return "";};

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_bc *CT_Relationship )ValidateWithPath (path string )error {if _fgg :=_bc .TargetModeAttr .ValidateWithPath (path +"\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072");_fgg !=nil {return _fgg ;};return nil ;};func NewRelationships ()*Relationships {_acg :=&Relationships {};_acg .CT_Relationships =*NewCT_Relationships ();return _acg ;};func (_gc ST_TargetMode )MarshalXML (e *_ee .Encoder ,start _ee .StartElement )error {return e .EncodeElement (_gc .String (),start );};

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_abg *Relationships )ValidateWithPath (path string )error {if _eae :=_abg .CT_Relationships .ValidateWithPath (path );_eae !=nil {return _eae ;};return nil ;};func (_eac *CT_Relationships )MarshalXML (e *_ee .Encoder ,start _ee .StartElement )error {e .EncodeToken (start );if _eac .Relationship !=nil {_bd :=_ee .StartElement {Name :_ee .Name {Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}};for _ ,_gg :=range _eac .Relationship {e .EncodeElement (_gg ,_bd );};};e .EncodeToken (_ee .EndElement {Name :start .Name });return nil ;};type ST_TargetMode byte ;func (_dag ST_TargetMode )ValidateWithPath (path string )error {switch _dag {case 0,1,2:default:return _b .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_dag ));};return nil ;};func (_ga *CT_Relationship )UnmarshalXML (d *_ee .Decoder ,start _ee .StartElement )error {for _ ,_be :=range start .Attr {if _be .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_ga .TargetModeAttr .UnmarshalXMLAttr (_be );continue ;};if _be .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_bb ,_ac :=_be .Value ,error (nil );if _ac !=nil {return _ac ;};_ga .TargetAttr =_bb ;continue ;};if _be .Name .Local =="\u0054\u0079\u0070\u0065"{_cg ,_ea :=_be .Value ,error (nil );if _ea !=nil {return _ea ;};_ga .TypeAttr =_cg ;continue ;};if _be .Name .Local =="\u0049\u0064"{_fg ,_cf :=_be .Value ,error (nil );if _cf !=nil {return _cf ;};_ga .IdAttr =_fg ;continue ;};};for {_d ,_bf :=d .Token ();if _bf !=nil {return _b .Errorf ("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073",_bf );};if _eea ,_bfe :=_d .(_ee .CharData );_bfe {_ga .Content =string (_eea );};if _ce ,_ab :=_d .(_ee .EndElement );_ab &&_ce .Name ==start .Name {break ;};};return nil ;};type CT_Relationships struct{Relationship []*Relationship ;};func (_fb *Relationships )MarshalXML (e *_ee .Encoder ,start _ee .StartElement )error {start .Attr =append (start .Attr ,_ee .Attr {Name :_ee .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"});start .Attr =append (start .Attr ,_ee .Attr {Name :_ee .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073";return _fb .CT_Relationships .MarshalXML (e ,start );};func (_ba *ST_TargetMode )UnmarshalXMLAttr (attr _ee .Attr )error {switch attr .Value {case "":*_ba =0;case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_ba =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_ba =2;};return nil ;};func (_g *CT_Relationship )MarshalXML (e *_ee .Encoder ,start _ee .StartElement )error {if _g .TargetModeAttr !=ST_TargetModeUnset {_gb ,_eg :=_g .TargetModeAttr .MarshalXMLAttr (_ee .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"});if _eg !=nil {return _eg ;};start .Attr =append (start .Attr ,_gb );};start .Attr =append (start .Attr ,_ee .Attr {Name :_ee .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074"},Value :_b .Sprintf ("\u0025\u0076",_g .TargetAttr )});start .Attr =append (start .Attr ,_ee .Attr {Name :_ee .Name {Local :"\u0054\u0079\u0070\u0065"},Value :_b .Sprintf ("\u0025\u0076",_g .TypeAttr )});start .Attr =append (start .Attr ,_ee .Attr {Name :_ee .Name {Local :"\u0049\u0064"},Value :_b .Sprintf ("\u0025\u0076",_g .IdAttr )});e .EncodeElement (_g .Content ,start );e .EncodeToken (_ee .EndElement {Name :start .Name });return nil ;};func (_fgba *Relationships )UnmarshalXML (d *_ee .Decoder ,start _ee .StartElement )error {_fgba .CT_Relationships =*NewCT_Relationships ();_eb :for {_ad ,_edf :=d .Token ();if _edf !=nil {return _edf ;};switch _dcf :=_ad .(type ){case _ee .StartElement :switch _dcf .Name {case _ee .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_ef :=NewRelationship ();if _ebg :=d .DecodeElement (_ef ,&_dcf );_ebg !=nil {return _ebg ;};_fgba .Relationship =append (_fgba .Relationship ,_ef );default:_a .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076",_dcf .Name );if _da :=d .Skip ();_da !=nil {return _da ;};};case _ee .EndElement :break _eb ;case _ee .CharData :};};return nil ;};func NewRelationship ()*Relationship {_bbc :=&Relationship {};_bbc .CT_Relationship =*NewCT_Relationship ();return _bbc ;};

// Validate validates the Relationships and its children
func (_cff *Relationships )Validate ()error {return _cff .ValidateWithPath ("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};type CT_Relationship struct{TargetModeAttr ST_TargetMode ;TargetAttr string ;TypeAttr string ;IdAttr string ;Content string ;};type Relationships struct{CT_Relationships };func NewCT_Relationships ()*CT_Relationships {_ge :=&CT_Relationships {};return _ge };func init (){_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewCT_Relationships );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewCT_Relationship );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewRelationships );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewRelationship );};