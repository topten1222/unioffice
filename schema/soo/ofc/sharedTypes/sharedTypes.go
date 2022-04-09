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

package sharedTypes ;import (_b "encoding/xml";_g "fmt";_f "regexp";);type ST_OnOff1 byte ;func (_agg *ST_YAlign )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_agg =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_agg =1;case "\u0074\u006f\u0070":*_agg =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_agg =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_agg =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_agg =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_agg =6;};return nil ;};func (_efdg ST_YAlign )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_fde :=_b .Attr {};_fde .Name =name ;switch _efdg {case ST_YAlignUnset :_fde .Value ="";case ST_YAlignInline :_fde .Value ="\u0069\u006e\u006c\u0069\u006e\u0065";case ST_YAlignTop :_fde .Value ="\u0074\u006f\u0070";case ST_YAlignCenter :_fde .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_YAlignBottom :_fde .Value ="\u0062\u006f\u0074\u0074\u006f\u006d";case ST_YAlignInside :_fde .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_YAlignOutside :_fde .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";};return _fde ,nil ;};const (ST_AlgTypeUnset ST_AlgType =0;ST_AlgTypeTypeAny ST_AlgType =1;ST_AlgTypeCustom ST_AlgType =2;);func (_ad *ST_CalendarType )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_ag ,_ba :=d .Token ();if _ba !=nil {return _ba ;};if _gd ,_ab :=_ag .(_b .EndElement );_ab &&_gd .Name ==start .Name {*_ad =1;return nil ;};if _gee ,_ggf :=_ag .(_b .CharData );!_ggf {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ag );}else {switch string (_gee ){case "":*_ad =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_ad =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_ad =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_ad =3;case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_ad =4;case "\u0068\u0069\u006ar\u0069":*_ad =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_ad =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_ad =7;case "\u006a\u0061\u0070a\u006e":*_ad =8;case "\u0074\u0068\u0061\u0069":*_ad =9;case "\u006b\u006f\u0072e\u0061":*_ad =10;case "\u0073\u0061\u006b\u0061":*_ad =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_ad =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_ad =13;case "\u006e\u006f\u006e\u0065":*_ad =14;};};_ag ,_ba =d .Token ();if _ba !=nil {return _ba ;};if _ef ,_fec :=_ag .(_b .EndElement );_fec &&_ef .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ag );};func (_fd ST_TrueFalse )ValidateWithPath (path string )error {switch _fd {case 0,1,2,3,4:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fd ));};return nil ;};type ST_AlgType byte ;func (_dgf *ST_OnOff1 )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_dgf =0;case "\u006f\u006e":*_dgf =1;case "\u006f\u0066\u0066":*_dgf =2;};return nil ;};func (_bbf ST_AlgClass )Validate ()error {return _bbf .ValidateWithPath ("")};func (_gge ST_AlgClass )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_efc :=_b .Attr {};_efc .Name =name ;switch _gge {case ST_AlgClassUnset :_efc .Value ="";case ST_AlgClassHash :_efc .Value ="\u0068\u0061\u0073\u0068";case ST_AlgClassCustom :_efc .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _efc ,nil ;};const (ST_ConformanceClassUnset ST_ConformanceClass =0;ST_ConformanceClassStrict ST_ConformanceClass =1;ST_ConformanceClassTransitional ST_ConformanceClass =2;);const ST_FixedPercentagePattern ="-\u003f\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039]\u003f\u0029\u0029\u0028\u005c\u002e\u005b\u0030\u002d\u0039][\u0030\u002d\u0039]\u003f)\u003f\u0025";func (_gb ST_CryptProv )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_gb .String (),start );};func (_ecb ST_VerticalAlignRun )ValidateWithPath (path string )error {switch _ecb {case 0,1,2,3:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ecb ));};return nil ;};func (_agd ST_CryptProv )ValidateWithPath (path string )error {switch _agd {case 0,1,2,3:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_agd ));};return nil ;};func (_caa ST_VerticalAlignRun )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_caa .String (),start );};func (_abe ST_XAlign )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_abe .String (),start );};

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct{ST_UnsignedDecimalNumber *uint64 ;ST_PositiveUniversalMeasure *string ;};const (ST_CalendarTypeUnset ST_CalendarType =0;ST_CalendarTypeGregorian ST_CalendarType =1;ST_CalendarTypeGregorianUs ST_CalendarType =2;ST_CalendarTypeGregorianMeFrench ST_CalendarType =3;ST_CalendarTypeGregorianArabic ST_CalendarType =4;ST_CalendarTypeHijri ST_CalendarType =5;ST_CalendarTypeHebrew ST_CalendarType =6;ST_CalendarTypeTaiwan ST_CalendarType =7;ST_CalendarTypeJapan ST_CalendarType =8;ST_CalendarTypeThai ST_CalendarType =9;ST_CalendarTypeKorea ST_CalendarType =10;ST_CalendarTypeSaka ST_CalendarType =11;ST_CalendarTypeGregorianXlitEnglish ST_CalendarType =12;ST_CalendarTypeGregorianXlitFrench ST_CalendarType =13;ST_CalendarTypeNone ST_CalendarType =14;);func (_dc ST_CalendarType )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_bde :=_b .Attr {};_bde .Name =name ;switch _dc {case ST_CalendarTypeUnset :_bde .Value ="";case ST_CalendarTypeGregorian :_bde .Value ="\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case ST_CalendarTypeGregorianUs :_bde .Value ="g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";case ST_CalendarTypeGregorianMeFrench :_bde .Value ="\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";case ST_CalendarTypeGregorianArabic :_bde .Value ="\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case ST_CalendarTypeHijri :_bde .Value ="\u0068\u0069\u006ar\u0069";case ST_CalendarTypeHebrew :_bde .Value ="\u0068\u0065\u0062\u0072\u0065\u0077";case ST_CalendarTypeTaiwan :_bde .Value ="\u0074\u0061\u0069\u0077\u0061\u006e";case ST_CalendarTypeJapan :_bde .Value ="\u006a\u0061\u0070a\u006e";case ST_CalendarTypeThai :_bde .Value ="\u0074\u0068\u0061\u0069";case ST_CalendarTypeKorea :_bde .Value ="\u006b\u006f\u0072e\u0061";case ST_CalendarTypeSaka :_bde .Value ="\u0073\u0061\u006b\u0061";case ST_CalendarTypeGregorianXlitEnglish :_bde .Value ="g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case ST_CalendarTypeGregorianXlitFrench :_bde .Value ="\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";case ST_CalendarTypeNone :_bde .Value ="\u006e\u006f\u006e\u0065";};return _bde ,nil ;};var ST_PositiveUniversalMeasurePatternRe =_f .MustCompile (ST_PositiveUniversalMeasurePattern );func (_fe ST_OnOff )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {e .EncodeToken (start );if _fe .Bool !=nil {e .EncodeToken (_b .CharData (_g .Sprintf ("\u0025\u0064",_dd (*_fe .Bool ))));};if _fe .ST_OnOff1 !=ST_OnOff1Unset {e .EncodeToken (_b .CharData (_fe .ST_OnOff1 .String ()));};return e .EncodeToken (_b .EndElement {Name :start .Name });};type ST_TrueFalse byte ;func (_dg *ST_CryptProv )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_fb ,_fgg :=d .Token ();if _fgg !=nil {return _fgg ;};if _gf ,_ggc :=_fb .(_b .EndElement );_ggc &&_gf .Name ==start .Name {*_dg =1;return nil ;};if _ca ,_cfe :=_fb .(_b .CharData );!_cfe {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_fb );}else {switch string (_ca ){case "":*_dg =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_dg =1;case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_dg =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_dg =3;};};_fb ,_fgg =d .Token ();if _fgg !=nil {return _fgg ;};if _ga ,_aae :=_fb .(_b .EndElement );_aae &&_ga .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_fb );};func (_fa ST_AlgType )ValidateWithPath (path string )error {switch _fa {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fa ));};return nil ;};func (_bb *ST_CalendarType )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_bb =0;case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":*_bb =1;case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":*_bb =2;case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":*_bb =3;case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":*_bb =4;case "\u0068\u0069\u006ar\u0069":*_bb =5;case "\u0068\u0065\u0062\u0072\u0065\u0077":*_bb =6;case "\u0074\u0061\u0069\u0077\u0061\u006e":*_bb =7;case "\u006a\u0061\u0070a\u006e":*_bb =8;case "\u0074\u0068\u0061\u0069":*_bb =9;case "\u006b\u006f\u0072e\u0061":*_bb =10;case "\u0073\u0061\u006b\u0061":*_bb =11;case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":*_bb =12;case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":*_bb =13;case "\u006e\u006f\u006e\u0065":*_bb =14;};return nil ;};func (_fbf *ST_AlgType )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_ddg ,_ecg :=d .Token ();if _ecg !=nil {return _ecg ;};if _gbg ,_ee :=_ddg .(_b .EndElement );_ee &&_gbg .Name ==start .Name {*_fbf =1;return nil ;};if _eab ,_acb :=_ddg .(_b .CharData );!_acb {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ddg );}else {switch string (_eab ){case "":*_fbf =0;case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_fbf =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_fbf =2;};};_ddg ,_ecg =d .Token ();if _ecg !=nil {return _ecg ;};if _gfd ,_bee :=_ddg .(_b .EndElement );_bee &&_gfd .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ddg );};func (_cba *ST_AlgClass )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_cba =0;case "\u0068\u0061\u0073\u0068":*_cba =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_cba =2;};return nil ;};var ST_FixedPercentagePatternRe =_f .MustCompile (ST_FixedPercentagePattern );func (_egd *ST_VerticalAlignRun )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_bdf ,_ade :=d .Token ();if _ade !=nil {return _ade ;};if _gbc ,_bcd :=_bdf .(_b .EndElement );_bcd &&_gbc .Name ==start .Name {*_egd =1;return nil ;};if _fgc ,_ebe :=_bdf .(_b .CharData );!_ebe {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_bdf );}else {switch string (_fgc ){case "":*_egd =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_egd =1;case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_egd =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_egd =3;};};_bdf ,_ade =d .Token ();if _ade !=nil {return _ade ;};if _acd ,_efd :=_bdf .(_b .EndElement );_efd &&_acd .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_bdf );};func (_ge *ST_TwipsMeasure )Validate ()error {return _ge .ValidateWithPath ("")};func (_efe ST_CalendarType )String ()string {switch _efe {case 0:return "";case 1:return "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n";case 2:return "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073";case 3:return "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068";case 4:return "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063";case 5:return "\u0068\u0069\u006ar\u0069";case 6:return "\u0068\u0065\u0062\u0072\u0065\u0077";case 7:return "\u0074\u0061\u0069\u0077\u0061\u006e";case 8:return "\u006a\u0061\u0070a\u006e";case 9:return "\u0074\u0068\u0061\u0069";case 10:return "\u006b\u006f\u0072e\u0061";case 11:return "\u0073\u0061\u006b\u0061";case 12:return "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068";case 13:return "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068";case 14:return "\u006e\u006f\u006e\u0065";};return "";};func (_ed ST_CalendarType )ValidateWithPath (path string )error {switch _ed {case 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ed ));};return nil ;};func (_edgd ST_ConformanceClass )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_edgd .String (),start );};func (_efef *ST_TrueFalseBlank )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_acge ,_dcab :=d .Token ();if _dcab !=nil {return _dcab ;};if _bfd ,_bc :=_acge .(_b .EndElement );_bc &&_bfd .Name ==start .Name {*_efef =1;return nil ;};if _cfg ,_ggd :=_acge .(_b .CharData );!_ggd {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_acge );}else {switch string (_cfg ){case "":*_efef =0;case "\u0074":*_efef =1;case "\u0066":*_efef =2;case "\u0074\u0072\u0075\u0065":*_efef =3;case "\u0066\u0061\u006cs\u0065":*_efef =4;case "\u0054\u0072\u0075\u0065":*_efef =6;case "\u0046\u0061\u006cs\u0065":*_efef =7;};};_acge ,_dcab =d .Token ();if _dcab !=nil {return _dcab ;};if _eabd ,_aag :=_acge .(_b .EndElement );_aag &&_eabd .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_acge );};func (_gff ST_AlgType )Validate ()error {return _gff .ValidateWithPath ("")};func (_gdc ST_TrueFalseBlank )ValidateWithPath (path string )error {switch _gdc {case 0,1,2,3,4,6,7:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gdc ));};return nil ;};func (_cg ST_AlgClass )ValidateWithPath (path string )error {switch _cg {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_cg ));};return nil ;};func (_dfe ST_OnOff1 )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_dfe .String (),start );};const (ST_CryptProvUnset ST_CryptProv =0;ST_CryptProvRsaAES ST_CryptProv =1;ST_CryptProvRsaFull ST_CryptProv =2;ST_CryptProvCustom ST_CryptProv =3;);func (_ddc ST_VerticalAlignRun )Validate ()error {return _ddc .ValidateWithPath ("")};func (_de *ST_CryptProv )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_de =0;case "\u0072\u0073\u0061\u0041\u0045\u0053":*_de =1;case "\u0072s\u0061\u0046\u0075\u006c\u006c":*_de =2;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_de =3;};return nil ;};const ST_PositiveUniversalMeasurePattern ="\u005b\u0030-9\u005d\u002b\u0028\\\u002e\u005b\u0030\u002d9]+\u0029?(\u006d\u006d\u007c\u0063\u006d\u007c\u0069n|\u0070\u0074\u007c\u0070\u0063\u007c\u0070i\u0029";type ST_CalendarType byte ;func (_eg ST_AlgClass )String ()string {switch _eg {case 0:return "";case 1:return "\u0068\u0061\u0073\u0068";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};func (_fecc *ST_YAlign )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_beeb ,_ebc :=d .Token ();if _ebc !=nil {return _ebc ;};if _edcd ,_eee :=_beeb .(_b .EndElement );_eee &&_edcd .Name ==start .Name {*_fecc =1;return nil ;};if _afd ,_cbad :=_beeb .(_b .CharData );!_cbad {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_beeb );}else {switch string (_afd ){case "":*_fecc =0;case "\u0069\u006e\u006c\u0069\u006e\u0065":*_fecc =1;case "\u0074\u006f\u0070":*_fecc =2;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_fecc =3;case "\u0062\u006f\u0074\u0074\u006f\u006d":*_fecc =4;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_fecc =5;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_fecc =6;};};_beeb ,_ebc =d .Token ();if _ebc !=nil {return _ebc ;};if _cd ,_dcg :=_beeb .(_b .EndElement );_dcg &&_cd .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_beeb );};const ST_PositivePercentagePattern ="\u005b0\u002d9\u005d\u002b\u0028\u005c\u002e[\u0030\u002d9\u005d\u002b\u0029\u003f\u0025";var ST_UniversalMeasurePatternRe =_f .MustCompile (ST_UniversalMeasurePattern );func (_edb ST_XAlign )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_dec :=_b .Attr {};_dec .Name =name ;switch _edb {case ST_XAlignUnset :_dec .Value ="";case ST_XAlignLeft :_dec .Value ="\u006c\u0065\u0066\u0074";case ST_XAlignCenter :_dec .Value ="\u0063\u0065\u006e\u0074\u0065\u0072";case ST_XAlignRight :_dec .Value ="\u0072\u0069\u0067h\u0074";case ST_XAlignInside :_dec .Value ="\u0069\u006e\u0073\u0069\u0064\u0065";case ST_XAlignOutside :_dec .Value ="\u006fu\u0074\u0073\u0069\u0064\u0065";};return _dec ,nil ;};const (ST_YAlignUnset ST_YAlign =0;ST_YAlignInline ST_YAlign =1;ST_YAlignTop ST_YAlign =2;ST_YAlignCenter ST_YAlign =3;ST_YAlignBottom ST_YAlign =4;ST_YAlignInside ST_YAlign =5;ST_YAlignOutside ST_YAlign =6;);func (_beg ST_TwipsMeasure )String ()string {if _beg .ST_UnsignedDecimalNumber !=nil {return _g .Sprintf ("\u0025\u0076",*_beg .ST_UnsignedDecimalNumber );};if _beg .ST_PositiveUniversalMeasure !=nil {return _g .Sprintf ("\u0025\u0076",*_beg .ST_PositiveUniversalMeasure );};return "";};func (_gc *ST_AlgClass )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_gef ,_acg :=d .Token ();if _acg !=nil {return _acg ;};if _ea ,_cc :=_gef .(_b .EndElement );_cc &&_ea .Name ==start .Name {*_gc =1;return nil ;};if _fg ,_db :=_gef .(_b .CharData );!_db {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gef );}else {switch string (_fg ){case "":*_gc =0;case "\u0068\u0061\u0073\u0068":*_gc =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_gc =2;};};_gef ,_acg =d .Token ();if _acg !=nil {return _acg ;};if _efb ,_ec :=_gef .(_b .EndElement );_ec &&_efb .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gef );};func (_cb *ST_TwipsMeasure )ValidateWithPath (path string )error {_ff :=[]string {};if _cb .ST_UnsignedDecimalNumber !=nil {_ff =append (_ff ,"\u0053T\u005f\u0055\u006e\u0073\u0069\u0067\u006e\u0065\u0064\u0044\u0065c\u0069\u006d\u0061\u006c\u004e\u0075\u006d\u0062\u0065\u0072");};if _cb .ST_PositiveUniversalMeasure !=nil {_ff =append (_ff ,"S\u0054\u005f\u0050\u006f\u0073\u0069t\u0069\u0076\u0065\u0055\u006e\u0069\u0076\u0065\u0072s\u0061\u006c\u004de\u0061s\u0075\u0072\u0065");};if len (_ff )> 1{return _g .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_ff );};return nil ;};func (_a *ST_OnOff )ValidateWithPath (path string )error {_be :=[]string {};if _a .Bool !=nil {_be =append (_be ,"\u0042\u006f\u006f\u006c");};if _a .ST_OnOff1 !=ST_OnOff1Unset {_be =append (_be ,"\u0053T\u005f\u004f\u006e\u004f\u0066\u00661");};if len (_be )> 1{return _g .Errorf ("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076",path ,_be );};return nil ;};func (_agda *ST_TrueFalse )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_dbe ,_eff :=d .Token ();if _eff !=nil {return _eff ;};if _eef ,_afb :=_dbe .(_b .EndElement );_afb &&_eef .Name ==start .Name {*_agda =1;return nil ;};if _aeae ,_afa :=_dbe .(_b .CharData );!_afa {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_dbe );}else {switch string (_aeae ){case "":*_agda =0;case "\u0074":*_agda =1;case "\u0066":*_agda =2;case "\u0074\u0072\u0075\u0065":*_agda =3;case "\u0066\u0061\u006cs\u0065":*_agda =4;};};_dbe ,_eff =d .Token ();if _eff !=nil {return _eff ;};if _bbfa ,_ccd :=_dbe .(_b .EndElement );_ccd &&_bbfa .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_dbe );};func (_fc ST_YAlign )String ()string {switch _fc {case 0:return "";case 1:return "\u0069\u006e\u006c\u0069\u006e\u0065";case 2:return "\u0074\u006f\u0070";case 3:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 4:return "\u0062\u006f\u0074\u0074\u006f\u006d";case 5:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 6:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";};func (_egc *ST_VerticalAlignRun )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_egc =0;case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":*_egc =1;case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":*_egc =2;case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":*_egc =3;};return nil ;};func (_bf *ST_TrueFalse )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_bf =0;case "\u0074":*_bf =1;case "\u0066":*_bf =2;case "\u0074\u0072\u0075\u0065":*_bf =3;case "\u0066\u0061\u006cs\u0065":*_bf =4;};return nil ;};const (ST_OnOff1Unset ST_OnOff1 =0;ST_OnOff1On ST_OnOff1 =1;ST_OnOff1Off ST_OnOff1 =2;);type ST_YAlign byte ;var ST_PositivePercentagePatternRe =_f .MustCompile (ST_PositivePercentagePattern );func (_eed ST_OnOff1 )Validate ()error {return _eed .ValidateWithPath ("")};func _dd (_da bool )uint8 {if _da {return 1;};return 0;};func (_dee ST_YAlign )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_dee .String (),start );};func (_dbee ST_TrueFalseBlank )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_fdg :=_b .Attr {};_fdg .Name =name ;switch _dbee {case ST_TrueFalseBlankUnset :_fdg .Value ="";case ST_TrueFalseBlankT :_fdg .Value ="\u0074";case ST_TrueFalseBlankF :_fdg .Value ="\u0066";case ST_TrueFalseBlankTrue :_fdg .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseBlankFalse :_fdg .Value ="\u0066\u0061\u006cs\u0065";case ST_TrueFalseBlankTrue_ :_fdg .Value ="\u0054\u0072\u0075\u0065";case ST_TrueFalseBlankFalse_ :_fdg .Value ="\u0046\u0061\u006cs\u0065";};return _fdg ,nil ;};var ST_PercentagePatternRe =_f .MustCompile (ST_PercentagePattern );const ST_PercentagePattern ="-\u003f[\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e[\u0030\u002d\u0039\u005d+)\u003f\u0025";

// ST_OnOff is a union type
type ST_OnOff struct{Bool *bool ;ST_OnOff1 ST_OnOff1 ;};func (_bdc ST_AlgClass )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_bdc .String (),start );};const (ST_TrueFalseUnset ST_TrueFalse =0;ST_TrueFalseT ST_TrueFalse =1;ST_TrueFalseF ST_TrueFalse =2;ST_TrueFalseTrue ST_TrueFalse =3;ST_TrueFalseFalse ST_TrueFalse =4;);func (_d ST_OnOff )String ()string {if _d .Bool !=nil {return _g .Sprintf ("\u0025\u0076",*_d .Bool );};if _d .ST_OnOff1 !=ST_OnOff1Unset {return _d .ST_OnOff1 .String ();};return "";};func (_bgd ST_TrueFalse )Validate ()error {return _bgd .ValidateWithPath ("")};func (_fad ST_TrueFalseBlank )String ()string {switch _fad {case 0:return "";case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";case 6:return "\u0054\u0072\u0075\u0065";case 7:return "\u0046\u0061\u006cs\u0065";};return "";};func (_bca *ST_ConformanceClass )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_bca =0;case "\u0073\u0074\u0072\u0069\u0063\u0074":*_bca =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_bca =2;};return nil ;};func (_gae ST_YAlign )ValidateWithPath (path string )error {switch _gae {case 0,1,2,3,4,5,6:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_gae ));};return nil ;};type ST_VerticalAlignRun byte ;func (_gcc ST_TrueFalse )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_ffd :=_b .Attr {};_ffd .Name =name ;switch _gcc {case ST_TrueFalseUnset :_ffd .Value ="";case ST_TrueFalseT :_ffd .Value ="\u0074";case ST_TrueFalseF :_ffd .Value ="\u0066";case ST_TrueFalseTrue :_ffd .Value ="\u0074\u0072\u0075\u0065";case ST_TrueFalseFalse :_ffd .Value ="\u0066\u0061\u006cs\u0065";};return _ffd ,nil ;};func ParseUnionST_OnOff (s string )(ST_OnOff ,error ){_bd :=ST_OnOff {};switch s {case "\u0074\u0072\u0075\u0065","\u0031","\u006f\u006e":_gg :=true ;_bd .Bool =&_gg ;default:_ae :=false ;_bd .Bool =&_ae ;};return _bd ,nil ;};func (_geb ST_ConformanceClass )ValidateWithPath (path string )error {switch _geb {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_geb ));};return nil ;};func (_agdd ST_AlgType )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_dfd :=_b .Attr {};_dfd .Name =name ;switch _agdd {case ST_AlgTypeUnset :_dfd .Value ="";case ST_AlgTypeTypeAny :_dfd .Value ="\u0074y\u0070\u0065\u0041\u006e\u0079";case ST_AlgTypeCustom :_dfd .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _dfd ,nil ;};func (_dff ST_ConformanceClass )Validate ()error {return _dff .ValidateWithPath ("")};func (_aea ST_AlgType )String ()string {switch _aea {case 0:return "";case 1:return "\u0074y\u0070\u0065\u0041\u006e\u0079";case 2:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};type ST_ConformanceClass byte ;const ST_GuidPattern ="\u005c\u007b\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b\u0038\u007d\u002d\u005b\u0030\u002d9\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030-\u0039\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b4\u007d\u002d\u005b\u0030\u002d\u0039A\u002d\u0046]\u007b\u00312\u007d\\\u007d";func (_ecgb ST_OnOff1 )ValidateWithPath (path string )error {switch _ecgb {case 0,1,2:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ecgb ));};return nil ;};type ST_CryptProv byte ;const (ST_AlgClassUnset ST_AlgClass =0;ST_AlgClassHash ST_AlgClass =1;ST_AlgClassCustom ST_AlgClass =2;);func (_fac ST_VerticalAlignRun )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_eb :=_b .Attr {};_eb .Name =name ;switch _fac {case ST_VerticalAlignRunUnset :_eb .Value ="";case ST_VerticalAlignRunBaseline :_eb .Value ="\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";case ST_VerticalAlignRunSuperscript :_eb .Value ="s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case ST_VerticalAlignRunSubscript :_eb .Value ="\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return _eb ,nil ;};func (_aaa ST_TrueFalseBlank )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_aaa .String (),start );};type ST_AlgClass byte ;func (_abd ST_XAlign )Validate ()error {return _abd .ValidateWithPath ("")};func (_ada *ST_XAlign )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_ada =0;case "\u006c\u0065\u0066\u0074":*_ada =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_ada =2;case "\u0072\u0069\u0067h\u0074":*_ada =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_ada =4;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_ada =5;};return nil ;};var ST_GuidPatternRe =_f .MustCompile (ST_GuidPattern );const (ST_XAlignUnset ST_XAlign =0;ST_XAlignLeft ST_XAlign =1;ST_XAlignCenter ST_XAlign =2;ST_XAlignRight ST_XAlign =3;ST_XAlignInside ST_XAlign =4;ST_XAlignOutside ST_XAlign =5;);func (_ega ST_ConformanceClass )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_ece :=_b .Attr {};_ece .Name =name ;switch _ega {case ST_ConformanceClassUnset :_ece .Value ="";case ST_ConformanceClassStrict :_ece .Value ="\u0073\u0074\u0072\u0069\u0063\u0074";case ST_ConformanceClassTransitional :_ece .Value ="\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";};return _ece ,nil ;};func (_bac ST_AlgType )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_bac .String (),start );};type ST_XAlign byte ;func (_fada ST_VerticalAlignRun )String ()string {switch _fada {case 0:return "";case 1:return "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065";case 2:return "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074";case 3:return "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t";};return "";};func (_dgd *ST_XAlign )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_fef ,_ce :=d .Token ();if _ce !=nil {return _ce ;};if _edc ,_eedf :=_fef .(_b .EndElement );_eedf &&_edc .Name ==start .Name {*_dgd =1;return nil ;};if _baf ,_bgc :=_fef .(_b .CharData );!_bgc {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_fef );}else {switch string (_baf ){case "":*_dgd =0;case "\u006c\u0065\u0066\u0074":*_dgd =1;case "\u0063\u0065\u006e\u0074\u0065\u0072":*_dgd =2;case "\u0072\u0069\u0067h\u0074":*_dgd =3;case "\u0069\u006e\u0073\u0069\u0064\u0065":*_dgd =4;case "\u006fu\u0074\u0073\u0069\u0064\u0065":*_dgd =5;};};_fef ,_ce =d .Token ();if _ce !=nil {return _ce ;};if _cgb ,_edg :=_fef .(_b .EndElement );_edg &&_cgb .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_fef );};func (_aa ST_CryptProv )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_cf :=_b .Attr {};_cf .Name =name ;switch _aa {case ST_CryptProvUnset :_cf .Value ="";case ST_CryptProvRsaAES :_cf .Value ="\u0072\u0073\u0061\u0041\u0045\u0053";case ST_CryptProvRsaFull :_cf .Value ="\u0072s\u0061\u0046\u0075\u006c\u006c";case ST_CryptProvCustom :_cf .Value ="\u0063\u0075\u0073\u0074\u006f\u006d";};return _cf ,nil ;};func (_gaf ST_OnOff1 )MarshalXMLAttr (name _b .Name )(_b .Attr ,error ){_ede :=_b .Attr {};_ede .Name =name ;switch _gaf {case ST_OnOff1Unset :_ede .Value ="";case ST_OnOff1On :_ede .Value ="\u006f\u006e";case ST_OnOff1Off :_ede .Value ="\u006f\u0066\u0066";};return _ede ,nil ;};func (_af ST_CryptProv )String ()string {switch _af {case 0:return "";case 1:return "\u0072\u0073\u0061\u0041\u0045\u0053";case 2:return "\u0072s\u0061\u0046\u0075\u006c\u006c";case 3:return "\u0063\u0075\u0073\u0074\u006f\u006d";};return "";};const ST_UniversalMeasurePattern ="\u002d\u003f\u005b\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u003f\u0028\u006d\u006d\u007c\u0063m\u007c\u0069\u006e\u007c\u0070t\u007c\u0070c\u007c\u0070\u0069\u0029";func (_c *ST_OnOff )Validate ()error {return _c .ValidateWithPath ("")};func (_ebf ST_YAlign )Validate ()error {return _ebf .ValidateWithPath ("")};type ST_TrueFalseBlank byte ;func (_ggec ST_XAlign )String ()string {switch _ggec {case 0:return "";case 1:return "\u006c\u0065\u0066\u0074";case 2:return "\u0063\u0065\u006e\u0074\u0065\u0072";case 3:return "\u0072\u0069\u0067h\u0074";case 4:return "\u0069\u006e\u0073\u0069\u0064\u0065";case 5:return "\u006fu\u0074\u0073\u0069\u0064\u0065";};return "";};func (_fbe ST_XAlign )ValidateWithPath (path string )error {switch _fbe {case 0,1,2,3,4,5:default:return _g .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_fbe ));};return nil ;};func (_dca ST_CalendarType )Validate ()error {return _dca .ValidateWithPath ("")};const (ST_VerticalAlignRunUnset ST_VerticalAlignRun =0;ST_VerticalAlignRunBaseline ST_VerticalAlignRun =1;ST_VerticalAlignRunSuperscript ST_VerticalAlignRun =2;ST_VerticalAlignRunSubscript ST_VerticalAlignRun =3;);func (_adg ST_OnOff1 )String ()string {switch _adg {case 0:return "";case 1:return "\u006f\u006e";case 2:return "\u006f\u0066\u0066";};return "";};func (_cce ST_TrueFalse )String ()string {switch _cce {case 0:return "";case 1:return "\u0074";case 2:return "\u0066";case 3:return "\u0074\u0072\u0075\u0065";case 4:return "\u0066\u0061\u006cs\u0065";};return "";};func (_ccc *ST_AlgType )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_ccc =0;case "\u0074y\u0070\u0065\u0041\u006e\u0079":*_ccc =1;case "\u0063\u0075\u0073\u0074\u006f\u006d":*_ccc =2;};return nil ;};func (_bed ST_TrueFalseBlank )Validate ()error {return _bed .ValidateWithPath ("")};func (_aec ST_TrueFalse )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_aec .String (),start );};const (ST_TrueFalseBlankUnset ST_TrueFalseBlank =0;ST_TrueFalseBlankT ST_TrueFalseBlank =1;ST_TrueFalseBlankF ST_TrueFalseBlank =2;ST_TrueFalseBlankTrue ST_TrueFalseBlank =3;ST_TrueFalseBlankFalse ST_TrueFalseBlank =4;ST_TrueFalseBlankTrue_ ST_TrueFalseBlank =6;ST_TrueFalseBlankFalse_ ST_TrueFalseBlank =7;);func (_efa *ST_TrueFalseBlank )UnmarshalXMLAttr (attr _b .Attr )error {switch attr .Value {case "":*_efa =0;case "\u0074":*_efa =1;case "\u0066":*_efa =2;case "\u0074\u0072\u0075\u0065":*_efa =3;case "\u0066\u0061\u006cs\u0065":*_efa =4;case "\u0054\u0072\u0075\u0065":*_efa =6;case "\u0046\u0061\u006cs\u0065":*_efa =7;};return nil ;};var ST_PositiveFixedPercentagePatternRe =_f .MustCompile (ST_PositiveFixedPercentagePattern );func (_efg ST_CryptProv )Validate ()error {return _efg .ValidateWithPath ("")};const ST_PositiveFixedPercentagePattern ="\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039\u005d\u003f\u0029\u0029\u0028\u005c\u002e[\u0030\u002d\u0039\u005d\u005b0\u002d\u0039]\u003f\u0029\u003f\u0025";func (_cdf *ST_ConformanceClass )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_gea ,_fga :=d .Token ();if _fga !=nil {return _fga ;};if _geea ,_cge :=_gea .(_b .EndElement );_cge &&_geea .Name ==start .Name {*_cdf =1;return nil ;};if _eea ,_cfd :=_gea .(_b .CharData );!_cfd {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gea );}else {switch string (_eea ){case "":*_cdf =0;case "\u0073\u0074\u0072\u0069\u0063\u0074":*_cdf =1;case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":*_cdf =2;};};_gea ,_fga =d .Token ();if _fga !=nil {return _fga ;};if _cged ,_acde :=_gea .(_b .EndElement );_acde &&_cged .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gea );};func (_bbe *ST_OnOff1 )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_aga ,_fba :=d .Token ();if _fba !=nil {return _fba ;};if _cga ,_gac :=_aga .(_b .EndElement );_gac &&_cga .Name ==start .Name {*_bbe =1;return nil ;};if _bag ,_bg :=_aga .(_b .CharData );!_bg {return _g .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_aga );}else {switch string (_bag ){case "":*_bbe =0;case "\u006f\u006e":*_bbe =1;case "\u006f\u0066\u0066":*_bbe =2;};};_aga ,_fba =d .Token ();if _fba !=nil {return _fba ;};if _bge ,_aac :=_aga .(_b .EndElement );_aac &&_bge .Name ==start .Name {return nil ;};return _g .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_aga );};func (_df ST_CalendarType )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {return e .EncodeElement (_df .String (),start );};func (_ac ST_TwipsMeasure )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {e .EncodeToken (start );if _ac .ST_UnsignedDecimalNumber !=nil {e .EncodeToken (_b .CharData (_g .Sprintf ("\u0025\u0064",*_ac .ST_UnsignedDecimalNumber )));};if _ac .ST_PositiveUniversalMeasure !=nil {e .EncodeToken (_b .CharData (*_ac .ST_PositiveUniversalMeasure ));};return e .EncodeToken (_b .EndElement {Name :start .Name });};func (_ace ST_ConformanceClass )String ()string {switch _ace {case 0:return "";case 1:return "\u0073\u0074\u0072\u0069\u0063\u0074";case 2:return "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c";};return "";};