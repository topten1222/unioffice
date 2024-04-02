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

package terms ;import (_c "encoding/xml";_cc "fmt";_e "github.com/unidoc/unioffice";_g "github.com/unidoc/unioffice/common/logger";_d "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_bda *ISO639_2 )ValidateWithPath (path string )error {return nil };func (_gfg *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_edb :for {_gbe ,_cd :=d .Token ();if _cd !=nil {return _cd ;};switch _gab :=_gbe .(type ){case _c .StartElement :switch _gab .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fd :=_d .NewAny ();
if _ecc :=d .DecodeElement (_fd ,&_gab );_ecc !=nil {return _ecc ;};_gfg .Any =append (_gfg .Any ,_fd );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_gab .Name );
if _aec :=d .Skip ();_aec !=nil {return _aec ;};};case _c .EndElement :break _edb ;case _c .CharData :};};return nil ;};

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_ggab *RFC1766 )ValidateWithPath (path string )error {return nil };func NewDDC ()*DDC {_ga :=&DDC {};return _ga };type ElementsAndRefinementsGroupChoice struct{Any []*_d .Any ;};type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;
};

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_dbb *ISO3166 )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_daf *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _cb ,_ebe :=range _daf .Any {if _gfc :=_ebe .ValidateWithPath (_cc .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_cb ));_gfc !=nil {return _gfc ;
};};return nil ;};func (_ed *DDC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func NewPeriod ()*Period {_ced :=&Period {};
return _ced };func (_aae *RFC1766 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_agcg ,_eea :=d .Token ();if _eea !=nil {return _cc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_eea );
};if _dcf ,_abaf :=_agcg .(_c .EndElement );_abaf &&_dcf .Name ==start .Name {break ;};};return nil ;};func NewIMT ()*IMT {_fg :=&IMT {};return _fg };func (_dga *ElementsAndRefinementsGroup )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_dgde :for {_cf ,_ace :=d .Token ();
if _ace !=nil {return _ace ;};switch _bbff :=_cf .(type ){case _c .StartElement :switch _bbff .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_afa :=NewElementsAndRefinementsGroupChoice ();
if _acf :=d .DecodeElement (&_afa .Any ,&_bbff );_acf !=nil {return _acf ;};_dga .Choice =append (_dga .Choice ,_afa );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_bbff .Name );
if _acg :=d .Skip ();_acg !=nil {return _acg ;};};case _c .EndElement :break _dgde ;case _c .CharData :};};return nil ;};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_aba :=&ElementsAndRefinementsGroup {};return _aba ;};func (_dfb *LCC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_abd *Point )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};type Period struct{};func (_dgd *ElementOrRefinementContainer )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";
e .EncodeToken (start );if _dgd .Choice !=nil {for _ ,_dcg :=range _dgd .Choice {_dcg .MarshalXML (e ,_c .StartElement {});};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_bed *Period )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_cedc *Period )ValidateWithPath (path string )error {return nil };

// Validate validates the ISO3166 and its children
func (_fe *ISO3166 )Validate ()error {return _fe .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};func NewDCMIType ()*DCMIType {_eg :=&DCMIType {};return _eg };

// Validate validates the ElementOrRefinementContainer and its children
func (_bc *ElementOrRefinementContainer )Validate ()error {return _bc .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};func (_bb *Box )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_df ,_cg :=d .Token ();
if _cg !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_cg );};if _dfg ,_eb :=_df .(_c .EndElement );_eb &&_dfg .Name ==start .Name {break ;};};return nil ;};func NewUDC ()*UDC {_cag :=&UDC {};
return _cag };

// Validate validates the TGN and its children
func (_fbd *TGN )Validate ()error {return _fbd .ValidateWithPath ("\u0054\u0047\u004e")};func (_agc *ElementsAndRefinementsGroupChoice )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _agc .Any !=nil {_dfgd :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_eeg :=range _agc .Any {e .EncodeElement (_eeg ,_dfgd );};};return nil ;};type MESH struct{};func (_eda *UDC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};func (_gg *DCMIType )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_dgf ,_gd :=d .Token ();if _gd !=nil {return _cc .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_gd );
};if _dgfe ,_dc :=_dgf .(_c .EndElement );_dc &&_dgfe .Name ==start .Name {break ;};};return nil ;};func (_ac *ElementOrRefinementContainer )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_fc :for {_dca ,_ae :=d .Token ();if _ae !=nil {return _ae ;
};switch _bf :=_dca .(type ){case _c .StartElement :switch _bf .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_gf :=NewElementsAndRefinementsGroupChoice ();
if _ggg :=d .DecodeElement (&_gf .Any ,&_bf );_ggg !=nil {return _ggg ;};_ac .Choice =append (_ac .Choice ,_gf );default:_g .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_bf .Name );
if _afc :=d .Skip ();_afc !=nil {return _afc ;};};case _c .EndElement :break _fc ;case _c .CharData :};};return nil ;};func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_afb :=&ElementsAndRefinementsGroupChoice {};return _afb ;
};func (_fff *RFC3066 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_ba *DCMIType )ValidateWithPath (path string )error {return nil };

// Validate validates the LCSH and its children
func (_aaf *LCSH )Validate ()error {return _aaf .ValidateWithPath ("\u004c\u0043\u0053\u0048")};func NewLCSH ()*LCSH {_bcb :=&LCSH {};return _bcb };func NewRFC3066 ()*RFC3066 {_eec :=&RFC3066 {};return _eec };

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_bcf *LCC )ValidateWithPath (path string )error {return nil };func (_bae *ISO639_2 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_eba ,_ggd :=d .Token ();if _ggd !=nil {return _cc .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_ggd );
};if _bbfd ,_ce :=_eba .(_c .EndElement );_ce &&_bbfd .Name ==start .Name {break ;};};return nil ;};type Box struct{};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_bcd *ElementsAndRefinementsGroup )Validate ()error {return _bcd .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};func NewPoint ()*Point {_gfca :=&Point {};
return _gfca };

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_bdf *UDC )ValidateWithPath (path string )error {return nil };func (_acd *RFC3066 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_cac ,_dfga :=d .Token ();if _dfga !=nil {return _cc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_dfga );
};if _gdd ,_cga :=_cac .(_c .EndElement );_cga &&_gdd .Name ==start .Name {break ;};};return nil ;};

// Validate validates the Box and its children
func (_af *Box )Validate ()error {return _af .ValidateWithPath ("\u0042\u006f\u0078")};type URI struct{};type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};func (_fae *W3CDTF )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_cfcg ,_bbb :=d .Token ();
if _bbb !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_bbb );};if _gfee ,_gad :=_cfcg .(_c .EndElement );_gad &&_gfee .Name ==start .Name {break ;};};return nil ;};

// Validate validates the UDC and its children
func (_gcad *UDC )Validate ()error {return _gcad .ValidateWithPath ("\u0055\u0044\u0043")};func (_gea *MESH )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_gda ,_bacf :=d .Token ();if _bacf !=nil {return _cc .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_bacf );
};if _bbc ,_afbb :=_gda .(_c .EndElement );_afbb &&_bbc .Name ==start .Name {break ;};};return nil ;};type W3CDTF struct{};func NewRFC1766 ()*RFC1766 {_aff :=&RFC1766 {};return _aff };type RFC3066 struct{};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_gcg *ElementsAndRefinementsGroupChoice )Validate ()error {return _gcg .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");
};func NewLCC ()*LCC {_gga :=&LCC {};return _gga };func NewISO3166 ()*ISO3166 {_bd :=&ISO3166 {};return _bd };

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_gdef *MESH )ValidateWithPath (path string )error {return nil };type Point struct{};func NewW3CDTF ()*W3CDTF {_eacc :=&W3CDTF {};return _eacc };

// Validate validates the MESH and its children
func (_dde *MESH )Validate ()error {return _dde .ValidateWithPath ("\u004d\u0045\u0053\u0048")};func NewMESH ()*MESH {_cea :=&MESH {};return _cea };type IMT struct{};func (_bg *LCSH )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_eccg *Point )ValidateWithPath (path string )error {return nil };func (_gfb *TGN )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};type DCMIType struct{};func (_fb *LCSH )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fdg ,_efg :=d .Token ();if _efg !=nil {return _cc .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_efg );
};if _cfcb ,_bace :=_fdg .(_c .EndElement );_bace &&_cfcb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the DCMIType and its children
func (_dgc *DCMIType )Validate ()error {return _dgc .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};func (_cdf *MESH )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";e .EncodeToken (start );
e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_eaa *LCC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_aea ,_gcd :=d .Token ();if _gcd !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_gcd );
};if _cfa ,_cdg :=_aea .(_c .EndElement );_cdg &&_cfa .Name ==start .Name {break ;};};return nil ;};

// Validate validates the DDC and its children
func (_ec *DDC )Validate ()error {return _ec .ValidateWithPath ("\u0044\u0044\u0043")};

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_aab *IMT )ValidateWithPath (path string )error {return nil };

// Validate validates the Point and its children
func (_afda *Point )Validate ()error {return _afda .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};

// Validate validates the URI and its children
func (_bad *URI )Validate ()error {return _bad .ValidateWithPath ("\u0055\u0052\u0049")};func (_gca *ISO639_2 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";e .EncodeToken (start );
e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the ISO639_2 and its children
func (_bdb *ISO639_2 )Validate ()error {return _bdb .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};func NewBox ()*Box {_be :=&Box {};return _be };func (_ged *Point )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_ggb ,_eff :=d .Token ();
if _eff !=nil {return _cc .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_eff );};if _bde ,_gcaa :=_ggb .(_c .EndElement );_gcaa &&_bde .Name ==start .Name {break ;};};return nil ;};func (_bfg *ISO3166 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_dgff ,_ef :=d .Token ();
if _ef !=nil {return _cc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_ef );};if _ge ,_bac :=_dgff .(_c .EndElement );_bac &&_ge .Name ==start .Name {break ;};};return nil ;};func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_gac :=&ElementOrRefinementContainer {};
return _gac ;};func (_a *Box )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_bdc *UDC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_de ,_dgb :=d .Token ();
if _dgb !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_dgb );};if _gfe ,_bgf :=_de .(_c .EndElement );_bgf &&_gfe .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_ece *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _bag ,_ag :=range _ece .Choice {if _ggf :=_ag .ValidateWithPath (_cc .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_bag ));
_ggf !=nil {return _ggf ;};};return nil ;};

// Validate validates the RFC3066 and its children
func (_aecg *RFC3066 )Validate ()error {return _aecg .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func (_dg *DCMIType )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_def *W3CDTF )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";e .EncodeToken (start );
e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the RFC1766 and its children
func (_gacd *RFC1766 )Validate ()error {return _gacd .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};func (_efb *RFC1766 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the Period and its children
func (_ccc *Period )Validate ()error {return _ccc .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};type RFC1766 struct{};type UDC struct{};

// Validate validates the IMT and its children
func (_afbf *IMT )Validate ()error {return _afbf .ValidateWithPath ("\u0049\u004d\u0054")};

// Validate validates the W3CDTF and its children
func (_agb *W3CDTF )Validate ()error {return _agb .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};func (_f *DDC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_ab ,_ad :=d .Token ();if _ad !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_ad );
};if _aa ,_cce :=_ab .(_c .EndElement );_cce &&_aa .Name ==start .Name {break ;};};return nil ;};type ISO639_2 struct{};type LCC struct{};func (_efe *TGN )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_adg ,_geg :=d .Token ();if _geg !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_geg );
};if _cgag ,_ebd :=_adg .(_c .EndElement );_ebd &&_cgag .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_dae *URI )ValidateWithPath (path string )error {return nil };func (_db *ElementsAndRefinementsGroup )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _db .Choice !=nil {for _ ,_ee :=range _db .Choice {_ee .MarshalXML (e ,_c .StartElement {});
};};return nil ;};

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_gc *Box )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_afd *LCSH )ValidateWithPath (path string )error {return nil };func (_eac *Period )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_gef ,_ca :=d .Token ();if _ca !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_ca );
};if _fec ,_eef :=_gef .(_c .EndElement );_eef &&_fec .Name ==start .Name {break ;};};return nil ;};

// Validate validates the LCC and its children
func (_dbf *LCC )Validate ()error {return _dbf .ValidateWithPath ("\u004c\u0043\u0043")};func NewURI ()*URI {_fac :=&URI {};return _fac };func NewTGN ()*TGN {_gcgb :=&TGN {};return _gcgb };type TGN struct{};func (_dd *IMT )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";
e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type DDC struct{};func (_ea *IMT )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_dffb ,_gde :=d .Token ();if _gde !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_gde );
};if _fa ,_dda :=_dffb .(_c .EndElement );_dda &&_fa .Name ==start .Name {break ;};};return nil ;};func NewISO639_2 ()*ISO639_2 {_aca :=&ISO639_2 {};return _aca };

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_da *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _cca ,_ff :=range _da .Choice {if _dcag :=_ff .ValidateWithPath (_cc .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_cca ));
_dcag !=nil {return _dcag ;};};return nil ;};

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_bab *TGN )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_eeb *RFC3066 )ValidateWithPath (path string )error {return nil };func (_cdd *URI )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fgf ,_faf :=d .Token ();if _faf !=nil {return _cc .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_faf );
};if _cfcd ,_bfd :=_fgf .(_c .EndElement );_bfd &&_cfcd .Name ==start .Name {break ;};};return nil ;};func (_cfc *ISO3166 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type ISO3166 struct{};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_aad *W3CDTF )ValidateWithPath (path string )error {return nil };type LCSH struct{};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_bbf *DDC )ValidateWithPath (path string )error {return nil };func (_cccf *URI )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};func init (){_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );
_e .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );
};