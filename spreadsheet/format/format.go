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

// Package format provides support for parsing and evaluating
// spreadsheetml/Excel number formats.
//
// Internally spreadsheets store numbers and dates values as a text
// representation of a floating point number (e.g. 1.2345).  This number is then
// displayed in Excel or another spreadsheet viewer differently depending on the
// number fornat of the cell style applied to the cell.
//
// As an example, the same value of 1.2345 can be displayed as:
// - "1" with format "0"
// - "1.2" with format "0.0"
// - "1.23" with format "0.00"
// - "1.235" with format "0.000"
// - "123%" with format "0%"
// - "1 23/100" with fornat "0 0/100"
// - "1.23E+00" with format "0.00E+00"
// - "29:37:41s" with format `[h]:mm:ss"s"`
package format ;import (_ad "bytes";_e "fmt";_aa "github.com/unidoc/unioffice/common/logger";_f "io";_gf "math";_g "strconv";_eb "strings";_c "time";);const _gde int =-1;const _ea ="\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0074\u0065\u0072a\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0069\u0067\u0069\u0074\u0046\u006d\u0074\u0054y\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0043o\u006d\u006d\u0061\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0065\u0063\u0069\u006da\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065Pe\u0072\u0063e\u006e\u0074\u0046\u006d\u0074\u0054\u0079\u0070e\u0044\u006f\u006c\u006c\u0061\u0072\u0046\u006d\u0074Ty\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0054\u0068\u006f\u0075\u0073\u0061n\u0064\u0073\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0055n\u0064\u0065\u0072\u0073c\u006f\u0072\u0065\u0046\u006d\u0074T\u0079\u0070\u0065\u0044\u0061\u0074\u0065\u0046\u006d\u0074\u0054y\u0070e\u0054\u0069\u006d\u0065\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0046\u0072\u0061\u0063t\u0069\u006f\u006e\u0046\u006dt\u0054\u0079\u0070\u0065\u0054e\u0078\u0074";func _fgb (_gd []byte )[]byte {for _ae :=0;_ae < len (_gd )/2;_ae ++{_bb :=len (_gd )-1-_ae ;_gd [_ae ],_gd [_bb ]=_gd [_bb ],_gd [_ae ];};return _gd ;};const _aagc int =34;func _cf (_ccgc float64 )string {_ebb :=_g .FormatFloat (_ccgc ,'E',-1,64);_cag :=_g .FormatFloat (_ccgc ,'E',5,64);if len (_ebb )< len (_cag ){return _g .FormatFloat (_ccgc ,'E',2,64);};return _cag ;};

// Token is a format token in the Excel format string.
type Token struct{Type FmtType ;Literal byte ;DateTime string ;};const _d =1e11;func Parse (s string )[]Format {_bcbg :=Lexer {};_bcbg .Lex (_eb .NewReader (s ));_bcbg ._cad =append (_bcbg ._cad ,_bcbg ._eeb );return _bcbg ._cad ;};func _bge (_eaad int64 ,_gdc Format )[]byte {if !_gdc .IsExponential ||len (_gdc .Exponent )==0{return nil ;};_ac :=_g .AppendInt (nil ,_ecb (_eaad ),10);_fdg :=make ([]byte ,0,len (_ac )+2);_fdg =append (_fdg ,'E');if _eaad >=0{_fdg =append (_fdg ,'+');}else {_fdg =append (_fdg ,'-');_eaad *=-1;};_bfa :=0;_fbf :for _db :=len (_gdc .Exponent )-1;_db >=0;_db --{_cee :=len (_ac )-1-_bfa ;_aae :=_gdc .Exponent [_db ];switch _aae .Type {case FmtTypeDigit :if _cee >=0{_fdg =append (_fdg ,_ac [_cee ]);_bfa ++;}else {_fdg =append (_fdg ,'0');};case FmtTypeDigitOpt :if _cee >=0{_fdg =append (_fdg ,_ac [_cee ]);_bfa ++;}else {for _cgd :=_db ;_cgd >=0;_cgd --{_aeg :=_gdc .Exponent [_cgd ];if _aeg .Type ==FmtTypeLiteral {_fdg =append (_fdg ,_aeg .Literal );};};break _fbf ;};case FmtTypeLiteral :_fdg =append (_fdg ,_aae .Literal );default:_aa .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0065\u0078p\u0020\u0025\u0076",_aae );};};if _bfa < len (_ac ){_fdg =append (_fdg ,_ac [len (_ac )-_bfa -1:_bfa -1]...);};_fgb (_fdg [2:]);return _fdg ;};const (FmtTypeLiteral FmtType =iota ;FmtTypeDigit ;FmtTypeDigitOpt ;FmtTypeComma ;FmtTypeDecimal ;FmtTypePercent ;FmtTypeDollar ;FmtTypeDigitOptThousands ;FmtTypeUnderscore ;FmtTypeDate ;FmtTypeTime ;FmtTypeFraction ;FmtTypeText ;);type Lexer struct{_eeb Format ;_cad []Format ;};const _bdb int =0;

// Format is a parsed number format.
type Format struct{Whole []Token ;Fractional []Token ;Exponent []Token ;IsExponential bool ;_fga bool ;_bg bool ;_cg bool ;_adc bool ;_cc bool ;_gb bool ;_af int64 ;_ba int ;};

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number (v float64 ,f string )string {if f ==""||f =="\u0047e\u006e\u0065\u0072\u0061\u006c"||f =="\u0040"{return NumberGeneric (v );};_gfd :=Parse (f );if len (_gfd )==1{return _ee (v ,_gfd [0],false );}else if len (_gfd )> 1&&v < 0{return _ee (v ,_gfd [1],true );}else if len (_gfd )> 2&&v ==0{return _ee (v ,_gfd [2],false );};return _ee (v ,_gfd [0],false );};const _def int =0;func _ge (_bbg ,_fdf float64 ,_fdd Format )[]byte {if len (_fdd .Fractional )==0{return nil ;};_ddc :=_g .AppendFloat (nil ,_bbg ,'f',-1,64);if len (_ddc )> 2{_ddc =_ddc [2:];}else {_ddc =nil ;};_fdb :=make ([]byte ,0,len (_ddc ));_fdb =append (_fdb ,'.');_feg :=0;_dea :for _fb :=0;_fb < len (_fdd .Fractional );_fb ++{_ega :=_fb ;_aag :=_fdd .Fractional [_fb ];switch _aag .Type {case FmtTypeDigit :if _ega < len (_ddc ){_fdb =append (_fdb ,_ddc [_ega ]);_feg ++;}else {_fdb =append (_fdb ,'0');};case FmtTypeDigitOpt :if _ega >=0{_fdb =append (_fdb ,_ddc [_ega ]);_feg ++;}else {break _dea ;};case FmtTypeLiteral :_fdb =append (_fdb ,_aag .Literal );default:_aa .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020f\u0072\u0061\u0063\u0074\u0069\u006f\u006ea\u006c\u0020\u0025\u0076",_aag );};};return _fdb ;};const _bdff int =34;func _dbc (_cce _c .Time ,_ffa float64 ,_ddb string )[]byte {_gc :=[]byte {};_bff :=0;for _ceea :=0;_ceea < len (_ddb );_ceea ++{var _ggc string ;if _ddb [_ceea ]==':'{_ggc =string (_ddb [_bff :_ceea ]);_bff =_ceea +1;}else if _ceea ==len (_ddb )-1{_ggc =string (_ddb [_bff :_ceea +1]);}else {continue ;};switch _ggc {case "\u0064":_gc =_cce .AppendFormat (_gc ,"\u0032");case "\u0068":_gc =_cce .AppendFormat (_gc ,"\u0033");case "\u0068\u0068":_gc =_cce .AppendFormat (_gc ,"\u0031\u0035");case "\u006d":_gc =_cce .AppendFormat (_gc ,"\u0034");case "\u006d\u006d":_gc =_cce .AppendFormat (_gc ,"\u0030\u0034");case "\u0073":_gc =_cce .Round (_c .Second ).AppendFormat (_gc ,"\u0035");case "\u0073\u002e\u0030":_gc =_cce .Round (_c .Second /10).AppendFormat (_gc ,"\u0035\u002e\u0030");case "\u0073\u002e\u0030\u0030":_gc =_cce .Round (_c .Second /100).AppendFormat (_gc ,"\u0035\u002e\u0030\u0030");case "\u0073\u002e\u00300\u0030":_gc =_cce .Round (_c .Second /1000).AppendFormat (_gc ,"\u0035\u002e\u00300\u0030");case "\u0073\u0073":_gc =_cce .Round (_c .Second ).AppendFormat (_gc ,"\u0030\u0035");case "\u0073\u0073\u002e\u0030":_gc =_cce .Round (_c .Second /10).AppendFormat (_gc ,"\u0030\u0035\u002e\u0030");case "\u0073\u0073\u002e0\u0030":_gc =_cce .Round (_c .Second /100).AppendFormat (_gc ,"\u0030\u0035\u002e0\u0030");case "\u0073\u0073\u002e\u0030\u0030\u0030":_gc =_cce .Round (_c .Second /1000).AppendFormat (_gc ,"\u0030\u0035\u002e\u0030\u0030\u0030");case "\u0041\u004d\u002fP\u004d":_gc =_cce .AppendFormat (_gc ,"\u0050\u004d");case "\u005b\u0068\u005d":_gc =_g .AppendInt (_gc ,int64 (_ffa *24),10);case "\u005b\u006d\u005d":_gc =_g .AppendInt (_gc ,int64 (_ffa *24*60),10);case "\u005b\u0073\u005d":_gc =_g .AppendInt (_gc ,int64 (_ffa *24*60*60),10);case "":default:_aa .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0069\u006d\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_ggc );};if _ddb [_ceea ]==':'{_gc =append (_gc ,':');};};return _gc ;};func (_ecbg *Lexer )nextFmt (){_ecbg ._cad =append (_ecbg ._cad ,_ecbg ._eeb );_ecbg ._eeb =Format {}};const _febe int =-1;var _fg =[...]uint8 {0,14,26,41,53,67,81,94,118,135,146,157,172,183};func _ecb (_eaa int64 )int64 {if _eaa < 0{return -_eaa ;};return _eaa ;};func _befed (_adba _c .Time )_c .Time {_adba =_adba .UTC ();return _c .Date (_adba .Year (),_adba .Month (),_adba .Day (),_adba .Hour (),_adba .Minute (),_adba .Second (),_adba .Nanosecond (),_c .Local );};func _ee (_ef float64 ,_ab Format ,_fae bool )string {if _ab ._cg {return NumberGeneric (_ef );};_afd :=make ([]byte ,0,20);_da :=_gf .Signbit (_ef );_eg :=_gf .Abs (_ef );_aec :=int64 (0);_cga :=int64 (0);if _ab .IsExponential {for _eg >=10{_cga ++;_eg /=10;};for _eg < 1{_cga --;_eg *=10;};}else if _ab ._bg {_eg *=100;}else if _ab ._fga {if _ab ._af ==0{_be :=_gf .Pow (10,float64 (_ab ._ba ));_fe ,_ec :=1.0,1.0;_ =_fe ;for _ecg :=1.0;_ecg < _be ;_ecg ++{_ ,_bc :=_gf .Modf (_eg *float64 (_ecg ));if _bc < _ec {_ec =_bc ;_fe =_ecg ;if _bc ==0{break ;};};};_ab ._af =int64 (_fe );};_aec =int64 (_eg *float64 (_ab ._af )+0.5);if len (_ab .Whole )> 0&&_aec > _ab ._af {_aec =int64 (_eg *float64 (_ab ._af ))%_ab ._af ;_eg -=float64 (_aec )/float64 (_ab ._af );}else {_eg -=float64 (_aec )/float64 (_ab ._af );if _gf .Abs (_eg )< 1{_edg :=true ;for _ ,_bcb :=range _ab .Whole {if _bcb .Type ==FmtTypeDigitOpt {continue ;};if _bcb .Type ==FmtTypeLiteral &&_bcb .Literal ==' '{continue ;};_edg =false ;};if _edg {_ab .Whole =nil ;};};};};_fed :=1;for _ ,_eeg :=range _ab .Fractional {if _eeg .Type ==FmtTypeDigit ||_eeg .Type ==FmtTypeDigitOpt {_fed ++;};};_eg +=5*_gf .Pow10 (-_fed );_dag ,_fad :=_gf .Modf (_eg );_afd =append (_afd ,_gdg (_dag ,_ef ,_ab )...);_afd =append (_afd ,_ge (_fad ,_ef ,_ab )...);_afd =append (_afd ,_bge (_cga ,_ab )...);if _ab ._fga {_afd =_g .AppendInt (_afd ,_aec ,10);_afd =append (_afd ,'/');_afd =_g .AppendInt (_afd ,_ab ._af ,10);};if !_fae &&_da {return "\u002d"+string (_afd );};return string (_afd );};const _ceee int =0;const _aegc int =34;

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value (v string ,f string )string {if IsNumber (v ){_cgc ,_ :=_g .ParseFloat (v ,64);return Number (_cgc ,f );};return String (v ,f );};const _b =1e-10;func (_ag FmtType )String ()string {if _ag >=FmtType (len (_fg )-1){return _e .Sprintf ("F\u006d\u0074\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029",_ag );};return _ea [_fg [_ag ]:_fg [_ag +1]];};func (_ebca *Lexer )Lex (r _f .Reader ){_dfdd ,_egd ,_edaf :=0,0,0;_eabd :=-1;_abg ,_bgb ,_debb :=0,0,0;_ =_bgb ;_ =_debb ;_fce :=1;_ =_fce ;_fff :=make ([]byte ,4096);_bfff :=false ;for !_bfff {_acc :=0;if _abg > 0{_acc =_egd -_abg ;};_egd =0;_cegf ,_ccef :=r .Read (_fff [_acc :]);if _cegf ==0||_ccef !=nil {_bfff =true ;};_edaf =_cegf +_acc ;if _edaf < len (_fff ){_eabd =_edaf ;};{_dfdd =_aagc ;_abg =0;_bgb =0;_debb =0;};{if _egd ==_edaf {goto _fcc ;};switch _dfdd {case 34:goto _cgde ;case 35:goto _faa ;case 0:goto _abbbf ;case 36:goto _cbc ;case 37:goto _cdfe ;case 1:goto _cgcfc ;case 2:goto _bgaa ;case 38:goto _egdb ;case 3:goto _eag ;case 4:goto _fcg ;case 39:goto _egf ;case 5:goto _bde ;case 6:goto _fgbd ;case 7:goto _dgf ;case 8:goto _abce ;case 40:goto _ecc ;case 9:goto _eage ;case 41:goto _aac ;case 10:goto _abcb ;case 42:goto _ceb ;case 11:goto _dee ;case 43:goto _fbe ;case 44:goto _bbe ;case 45:goto _cbcc ;case 12:goto _abfe ;case 46:goto _bea ;case 13:goto _daaf ;case 14:goto _afge ;case 15:goto _geg ;case 16:goto _bcbd ;case 47:goto _bfg ;case 17:goto _cbcb ;case 48:goto _fcd ;case 18:goto _beg ;case 19:goto _fdfa ;case 20:goto _egce ;case 49:goto _beed ;case 50:goto _ccca ;case 21:goto _ceff ;case 22:goto _bfcg ;case 23:goto _cfd ;case 24:goto _aabg ;case 25:goto _aaad ;case 51:goto _daf ;case 26:goto _gceb ;case 52:goto _aef ;case 53:goto _bab ;case 54:goto _acd ;case 55:goto _gcgg ;case 56:goto _gda ;case 57:goto _aaca ;case 27:goto _fbag ;case 28:goto _fdga ;case 29:goto _eedc ;case 30:goto _aaac ;case 31:goto _dagc ;case 58:goto _debf ;case 32:goto _gba ;case 59:goto _gef ;case 33:goto _bebf ;case 60:goto _dcfd ;case 61:goto _ddf ;case 62:goto _eebe ;};goto _gfda ;_baa :switch _debb {case 2:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeDigit ,nil );};case 3:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeDigitOpt ,nil );};case 5:{_egd =(_bgb )-1;};case 8:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypePercent ,nil );};case 13:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeFraction ,_fff [_abg :_bgb ]);};case 14:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeDate ,_fff [_abg :_bgb ]);};case 15:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeTime ,_fff [_abg :_bgb ]);};case 16:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeTime ,_fff [_abg :_bgb ]);};case 18:{_egd =(_bgb )-1;};case 20:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg :_bgb ]);};case 21:{_egd =(_bgb )-1;_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg +1:_bgb -1]);};};goto _bbd ;_dgb :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeFraction ,_fff [_abg :_bgb ]);};goto _bbd ;_fbg :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeDigitOpt ,nil );};goto _bbd ;_daa :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeDigitOptThousands ,nil );};goto _bbd ;_aba :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypePercent ,nil );};goto _bbd ;_cef :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeDate ,_fff [_abg :_bgb ]);};goto _bbd ;_bgg :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeDigit ,nil );};goto _bbd ;_cdee :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeTime ,_fff [_abg :_bgb ]);};goto _bbd ;_addb :_egd =(_bgb )-1;{_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg :_bgb ]);};goto _bbd ;_ggcb :_bgb =_egd +1;{_ebca ._eeb ._cg =true ;};goto _bbd ;_cgcf :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg :_bgb ]);};goto _bbd ;_dba :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeDollar ,nil );};goto _bbd ;_cega :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeComma ,nil );};goto _bbd ;_eed :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeDecimal ,nil );};goto _bbd ;_deaa :_bgb =_egd +1;{_ebca .nextFmt ();};goto _bbd ;_gdga :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeText ,nil );};goto _bbd ;_ggg :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeUnderscore ,nil );};goto _bbd ;_abc :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg :_bgb ]);};goto _bbd ;_dbac :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg +1:_bgb -1]);};goto _bbd ;_egaa :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeDigitOpt ,nil );};goto _bbd ;_agg :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeFraction ,_fff [_abg :_bgb ]);};goto _bbd ;_edeg :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypePercent ,nil );};goto _bbd ;_gce :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeDate ,_fff [_abg :_bgb ]);};goto _bbd ;_cdg :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeDigit ,nil );};goto _bbd ;_gddg :_bgb =_egd ;_egd --;{_ebca ._eeb .AddToken (FmtTypeTime ,_fff [_abg :_bgb ]);};goto _bbd ;_gfa :_bgb =_egd ;_egd --;{};goto _bbd ;_aad :_bgb =_egd +1;{_ebca ._eeb .IsExponential =true ;};goto _bbd ;_adcb :_bgb =_egd +1;{_ebca ._eeb .AddToken (FmtTypeLiteral ,_fff [_abg +1:_bgb ]);};goto _bbd ;_bbd :_abg =0;if _egd ++;_egd ==_edaf {goto _dgfd ;};_cgde :_abg =_egd ;switch _fff [_egd ]{case 34:goto _dgee ;case 35:goto _cccg ;case 36:goto _dba ;case 37:goto _fbc ;case 44:goto _cega ;case 46:goto _eed ;case 47:goto _bbf ;case 48:goto _cff ;case 58:goto _cadb ;case 59:goto _deaa ;case 63:goto _bbc ;case 64:goto _gdga ;case 65:goto _cbcg ;case 69:goto _bgf ;case 71:goto _ffb ;case 91:goto _accb ;case 92:goto _ffgb ;case 95:goto _ggg ;case 100:goto _bbf ;case 104:goto _cadb ;case 109:goto _dcaf ;case 115:goto _bba ;case 121:goto _bdbg ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _ggd ;};goto _cgcf ;_dgee :_bgb =_egd +1;_debb =20;goto _fef ;_fef :if _egd ++;_egd ==_edaf {goto _cbe ;};_faa :if _fff [_egd ]==34{goto _gca ;};goto _gdf ;_gdf :if _egd ++;_egd ==_edaf {goto _cfb ;};_abbbf :if _fff [_egd ]==34{goto _gca ;};goto _gdf ;_gca :_bgb =_egd +1;_debb =21;goto _aab ;_aab :if _egd ++;_egd ==_edaf {goto _cdeef ;};_cbc :if _fff [_egd ]==34{goto _gdf ;};goto _dbac ;_cccg :_bgb =_egd +1;_debb =3;goto _bad ;_bad :if _egd ++;_egd ==_edaf {goto _gabf ;};_cdfe :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _agb ;case 44:goto _gcf ;case 47:goto _ecd ;case 48:goto _agb ;case 63:goto _agb ;};goto _egaa ;_agb :if _egd ++;_egd ==_edaf {goto _bcbe ;};_cgcfc :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _agb ;case 47:goto _ecd ;case 48:goto _agb ;case 63:goto _agb ;};goto _baa ;_ecd :if _egd ++;_egd ==_edaf {goto _cdac ;};_bgaa :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _efg ;case 48:goto _cgf ;case 63:goto _befe ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _gbb ;};goto _baa ;_befe :_bgb =_egd +1;goto _bgbg ;_bgbg :if _egd ++;_egd ==_edaf {goto _ffgf ;};_egdb :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _befe ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _befe ;case 63:goto _befe ;case 65:goto _ece ;};goto _agg ;_ece :if _egd ++;_egd ==_edaf {goto _gefe ;};_eag :switch _fff [_egd ]{case 47:goto _acf ;case 77:goto _cbdg ;};goto _dgb ;_acf :if _egd ++;_egd ==_edaf {goto _dde ;};_fcg :if _fff [_egd ]==80{goto _gggb ;};goto _dgb ;_gggb :_bgb =_egd +1;goto _fca ;_fca :if _egd ++;_egd ==_edaf {goto _dfce ;};_egf :if _fff [_egd ]==65{goto _ece ;};goto _agg ;_cbdg :if _egd ++;_egd ==_edaf {goto _dagf ;};_bde :if _fff [_egd ]==47{goto _bfd ;};goto _dgb ;_bfd :if _egd ++;_egd ==_edaf {goto _cec ;};_fgbd :if _fff [_egd ]==80{goto _efb ;};goto _dgb ;_efb :if _egd ++;_egd ==_edaf {goto _fea ;};_dgf :if _fff [_egd ]==77{goto _gggb ;};goto _dgb ;_efg :if _egd ++;_egd ==_edaf {goto _gea ;};_abce :switch _fff [_egd ]{case 35:goto _dgc ;case 37:goto _gaec ;case 63:goto _dgc ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _defd ;};goto _baa ;_dgc :_bgb =_egd +1;goto _fgf ;_fgf :if _egd ++;_egd ==_edaf {goto _deeb ;};_ecc :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _bgd ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _befe ;case 63:goto _befe ;case 65:goto _ece ;};goto _agg ;_bgd :if _egd ++;_egd ==_edaf {goto _bbae ;};_eage :switch _fff [_egd ]{case 35:goto _cfa ;case 44:goto _cfa ;case 46:goto _cfa ;case 48:goto _cfa ;case 63:goto _cfa ;};goto _dgb ;_cfa :_bgb =_egd +1;goto _aff ;_aff :if _egd ++;_egd ==_edaf {goto _fbad ;};_aac :switch _fff [_egd ]{case 35:goto _cfa ;case 44:goto _cfa ;case 46:goto _cfa ;case 48:goto _cfa ;case 63:goto _cfa ;case 65:goto _ece ;};goto _agg ;_gaec :if _egd ++;_egd ==_edaf {goto _bdaa ;};_abcb :if _fff [_egd ]==37{goto _gaec ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _defd ;};goto _baa ;_defd :_bgb =_egd +1;_debb =13;goto _ccgb ;_ccgb :if _egd ++;_egd ==_edaf {goto _fdfb ;};_ceb :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _bda ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _ebbe ;case 63:goto _befe ;case 65:goto _ece ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _defd ;};goto _agg ;_bda :if _egd ++;_egd ==_edaf {goto _fde ;};_dee :switch _fff [_egd ]{case 35:goto _cfa ;case 37:goto _gaec ;case 44:goto _cfa ;case 46:goto _cfa ;case 63:goto _cfa ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _defd ;};goto _dgb ;_ebbe :_bgb =_egd +1;goto _eae ;_eae :if _egd ++;_egd ==_edaf {goto _dab ;};_fbe :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _ebbe ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _ebbe ;case 63:goto _befe ;case 65:goto _ece ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _defd ;};goto _agg ;_cgf :_bgb =_egd +1;goto _egdc ;_egdc :if _egd ++;_egd ==_edaf {goto _fedd ;};_bbe :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _ebbe ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _cgf ;case 63:goto _befe ;case 65:goto _ece ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _gbb ;};goto _agg ;_gbb :_bgb =_egd +1;goto _ggga ;_ggga :if _egd ++;_egd ==_edaf {goto _gcd ;};_cbcc :switch _fff [_egd ]{case 35:goto _befe ;case 37:goto _defd ;case 44:goto _befe ;case 46:goto _befe ;case 48:goto _cgf ;case 63:goto _befe ;case 65:goto _ece ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _gbb ;};goto _agg ;_gcf :if _egd ++;_egd ==_edaf {goto _ddcb ;};_abfe :if _fff [_egd ]==35{goto _daa ;};goto _fbg ;_fbc :_bgb =_egd +1;_debb =8;goto _bec ;_bec :if _egd ++;_egd ==_edaf {goto _ebg ;};_bea :switch _fff [_egd ]{case 35:goto _aacc ;case 37:goto _gbbe ;case 48:goto _ffe ;case 63:goto _aacc ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _bcg ;};goto _edeg ;_aacc :if _egd ++;_egd ==_edaf {goto _efc ;};_daaf :switch _fff [_egd ]{case 35:goto _aacc ;case 47:goto _ecd ;case 48:goto _aacc ;case 63:goto _aacc ;};goto _aba ;_gbbe :if _egd ++;_egd ==_edaf {goto _dfb ;};_afge :if _fff [_egd ]==37{goto _gbbe ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _bcg ;};goto _baa ;_bcg :if _egd ++;_egd ==_edaf {goto _edafe ;};_geg :switch _fff [_egd ]{case 37:goto _gbbe ;case 47:goto _ecd ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _bcg ;};goto _baa ;_ffe :if _egd ++;_egd ==_edaf {goto _gcb ;};_bcbd :switch _fff [_egd ]{case 35:goto _aacc ;case 37:goto _gbbe ;case 47:goto _ecd ;case 48:goto _ffe ;case 63:goto _aacc ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _bcg ;};goto _aba ;_bbf :_bgb =_egd +1;goto _gbdf ;_gbdf :if _egd ++;_egd ==_edaf {goto _cbg ;};_bfg :switch _fff [_egd ]{case 47:goto _bbf ;case 100:goto _bbf ;case 109:goto _bbf ;case 121:goto _cfg ;};goto _gce ;_cfg :if _egd ++;_egd ==_edaf {goto _bgff ;};_cbcb :if _fff [_egd ]==121{goto _bbf ;};goto _cef ;_cff :_bgb =_egd +1;_debb =2;goto _ebcae ;_ebcae :if _egd ++;_egd ==_edaf {goto _bdbgf ;};_fcd :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _caa ;case 47:goto _ecd ;case 48:goto _dca ;case 63:goto _agb ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _deea ;};goto _cdg ;_caa :if _egd ++;_egd ==_edaf {goto _feaf ;};_beg :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _caa ;case 47:goto _ecd ;case 48:goto _caa ;case 63:goto _agb ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _bcg ;};goto _bgg ;_dca :if _egd ++;_egd ==_edaf {goto _beae ;};_fdfa :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _caa ;case 47:goto _ecd ;case 48:goto _dca ;case 63:goto _agb ;};if 49<=_fff [_egd ]&&_fff [_egd ]<=57{goto _deea ;};goto _bgg ;_deea :if _egd ++;_egd ==_edaf {goto _acff ;};_egce :switch _fff [_egd ]{case 37:goto _bcg ;case 47:goto _ecd ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _deea ;};goto _baa ;_ggd :_bgb =_egd +1;_debb =20;goto _degc ;_degc :if _egd ++;_egd ==_edaf {goto _cafd ;};_beed :switch _fff [_egd ]{case 37:goto _bcg ;case 47:goto _ecd ;};if 48<=_fff [_egd ]&&_fff [_egd ]<=57{goto _deea ;};goto _abc ;_cadb :_bgb =_egd +1;_debb =15;goto _cegfa ;_cegfa :if _egd ++;_egd ==_edaf {goto _ecda ;};_ccca :switch _fff [_egd ]{case 58:goto _cadb ;case 65:goto _ddd ;case 104:goto _cadb ;case 109:goto _cadb ;case 115:goto _bba ;};goto _gddg ;_ddd :if _egd ++;_egd ==_edaf {goto _bdee ;};_ceff :switch _fff [_egd ]{case 47:goto _dbag ;case 77:goto _cac ;};goto _baa ;_dbag :if _egd ++;_egd ==_edaf {goto _cebg ;};_bfcg :if _fff [_egd ]==80{goto _cadb ;};goto _baa ;_cac :if _egd ++;_egd ==_edaf {goto _cdb ;};_cfd :if _fff [_egd ]==47{goto _cfe ;};goto _baa ;_cfe :if _egd ++;_egd ==_edaf {goto _acfbf ;};_aabg :if _fff [_egd ]==80{goto _cdfg ;};goto _baa ;_cdfg :if _egd ++;_egd ==_edaf {goto _gaa ;};_aaad :if _fff [_egd ]==77{goto _cadb ;};goto _baa ;_bba :_bgb =_egd +1;_debb =15;goto _ead ;_ead :if _egd ++;_egd ==_edaf {goto _ggf ;};_daf :switch _fff [_egd ]{case 46:goto _fabe ;case 58:goto _cadb ;case 65:goto _ddd ;case 104:goto _cadb ;case 109:goto _cadb ;case 115:goto _bba ;};goto _gddg ;_fabe :if _egd ++;_egd ==_edaf {goto _fefb ;};_gceb :if _fff [_egd ]==48{goto _acfb ;};goto _cdee ;_acfb :_bgb =_egd +1;_debb =15;goto _cacb ;_cacb :if _egd ++;_egd ==_edaf {goto _cfgd ;};_aef :switch _fff [_egd ]{case 48:goto _ecbf ;case 58:goto _cadb ;case 65:goto _ddd ;case 104:goto _cadb ;case 109:goto _cadb ;case 115:goto _bba ;};goto _gddg ;_ecbf :_bgb =_egd +1;_debb =15;goto _gcgf ;_gcgf :if _egd ++;_egd ==_edaf {goto _ceae ;};_bab :switch _fff [_egd ]{case 48:goto _cadb ;case 58:goto _cadb ;case 65:goto _ddd ;case 104:goto _cadb ;case 109:goto _cadb ;case 115:goto _bba ;};goto _gddg ;_bbc :_bgb =_egd +1;_debb =5;goto _gdda ;_gdda :if _egd ++;_egd ==_edaf {goto _caab ;};_acd :switch _fff [_egd ]{case 35:goto _agb ;case 37:goto _agb ;case 47:goto _ecd ;case 48:goto _agb ;case 63:goto _agb ;};goto _gfa ;_cbcg :_bgb =_egd +1;_debb =20;goto _ebf ;_ebf :if _egd ++;_egd ==_edaf {goto _abac ;};_gcgg :switch _fff [_egd ]{case 47:goto _dbag ;case 77:goto _cac ;};goto _abc ;_bgf :if _egd ++;_egd ==_edaf {goto _fgfg ;};_gda :switch _fff [_egd ]{case 43:goto _aad ;case 45:goto _aad ;};goto _abc ;_ffb :_bgb =_egd +1;goto _dce ;_dce :if _egd ++;_egd ==_edaf {goto _aegg ;};_aaca :if _fff [_egd ]==101{goto _eaac ;};goto _abc ;_eaac :if _egd ++;_egd ==_edaf {goto _ddcc ;};_fbag :if _fff [_egd ]==110{goto _bae ;};goto _addb ;_bae :if _egd ++;_egd ==_edaf {goto _cgag ;};_fdga :if _fff [_egd ]==101{goto _cddd ;};goto _addb ;_cddd :if _egd ++;_egd ==_edaf {goto _fdgb ;};_eedc :if _fff [_egd ]==114{goto _ebbd ;};goto _addb ;_ebbd :if _egd ++;_egd ==_edaf {goto _gac ;};_aaac :if _fff [_egd ]==97{goto _cdag ;};goto _addb ;_cdag :if _egd ++;_egd ==_edaf {goto _cca ;};_dagc :if _fff [_egd ]==108{goto _ggcb ;};goto _addb ;_accb :_bgb =_egd +1;_debb =20;goto _edbd ;_edbd :if _egd ++;_egd ==_edaf {goto _ddeb ;};_debf :switch _fff [_egd ]{case 104:goto _aaec ;case 109:goto _aaec ;case 115:goto _aaec ;};goto _aacb ;_aacb :if _egd ++;_egd ==_edaf {goto _adb ;};_gba :if _fff [_egd ]==93{goto _dff ;};goto _aacb ;_dff :_bgb =_egd +1;_debb =18;goto _ged ;_gabg :_bgb =_egd +1;_debb =16;goto _ged ;_ged :if _egd ++;_egd ==_edaf {goto _eeaf ;};_gef :if _fff [_egd ]==93{goto _dff ;};goto _aacb ;_aaec :if _egd ++;_egd ==_edaf {goto _cacf ;};_bebf :if _fff [_egd ]==93{goto _gabg ;};goto _aacb ;_ffgb :if _egd ++;_egd ==_edaf {goto _bfb ;};_dcfd :goto _adcb ;_dcaf :_bgb =_egd +1;_debb =14;goto _faea ;_faea :if _egd ++;_egd ==_edaf {goto _cddc ;};_ddf :switch _fff [_egd ]{case 47:goto _bbf ;case 58:goto _cadb ;case 65:goto _ddd ;case 100:goto _bbf ;case 104:goto _cadb ;case 109:goto _dcaf ;case 115:goto _bba ;case 121:goto _cfg ;};goto _gce ;_bdbg :if _egd ++;_egd ==_edaf {goto _dfbf ;};_eebe :if _fff [_egd ]==121{goto _bbf ;};goto _abc ;_gfda :_dgfd :_dfdd =34;goto _fcc ;_cbe :_dfdd =35;goto _fcc ;_cfb :_dfdd =0;goto _fcc ;_cdeef :_dfdd =36;goto _fcc ;_gabf :_dfdd =37;goto _fcc ;_bcbe :_dfdd =1;goto _fcc ;_cdac :_dfdd =2;goto _fcc ;_ffgf :_dfdd =38;goto _fcc ;_gefe :_dfdd =3;goto _fcc ;_dde :_dfdd =4;goto _fcc ;_dfce :_dfdd =39;goto _fcc ;_dagf :_dfdd =5;goto _fcc ;_cec :_dfdd =6;goto _fcc ;_fea :_dfdd =7;goto _fcc ;_gea :_dfdd =8;goto _fcc ;_deeb :_dfdd =40;goto _fcc ;_bbae :_dfdd =9;goto _fcc ;_fbad :_dfdd =41;goto _fcc ;_bdaa :_dfdd =10;goto _fcc ;_fdfb :_dfdd =42;goto _fcc ;_fde :_dfdd =11;goto _fcc ;_dab :_dfdd =43;goto _fcc ;_fedd :_dfdd =44;goto _fcc ;_gcd :_dfdd =45;goto _fcc ;_ddcb :_dfdd =12;goto _fcc ;_ebg :_dfdd =46;goto _fcc ;_efc :_dfdd =13;goto _fcc ;_dfb :_dfdd =14;goto _fcc ;_edafe :_dfdd =15;goto _fcc ;_gcb :_dfdd =16;goto _fcc ;_cbg :_dfdd =47;goto _fcc ;_bgff :_dfdd =17;goto _fcc ;_bdbgf :_dfdd =48;goto _fcc ;_feaf :_dfdd =18;goto _fcc ;_beae :_dfdd =19;goto _fcc ;_acff :_dfdd =20;goto _fcc ;_cafd :_dfdd =49;goto _fcc ;_ecda :_dfdd =50;goto _fcc ;_bdee :_dfdd =21;goto _fcc ;_cebg :_dfdd =22;goto _fcc ;_cdb :_dfdd =23;goto _fcc ;_acfbf :_dfdd =24;goto _fcc ;_gaa :_dfdd =25;goto _fcc ;_ggf :_dfdd =51;goto _fcc ;_fefb :_dfdd =26;goto _fcc ;_cfgd :_dfdd =52;goto _fcc ;_ceae :_dfdd =53;goto _fcc ;_caab :_dfdd =54;goto _fcc ;_abac :_dfdd =55;goto _fcc ;_fgfg :_dfdd =56;goto _fcc ;_aegg :_dfdd =57;goto _fcc ;_ddcc :_dfdd =27;goto _fcc ;_cgag :_dfdd =28;goto _fcc ;_fdgb :_dfdd =29;goto _fcc ;_gac :_dfdd =30;goto _fcc ;_cca :_dfdd =31;goto _fcc ;_ddeb :_dfdd =58;goto _fcc ;_adb :_dfdd =32;goto _fcc ;_eeaf :_dfdd =59;goto _fcc ;_cacf :_dfdd =33;goto _fcc ;_bfb :_dfdd =60;goto _fcc ;_cddc :_dfdd =61;goto _fcc ;_dfbf :_dfdd =62;goto _fcc ;_fcc :{};if _egd ==_eabd {switch _dfdd {case 35:goto _abc ;case 0:goto _baa ;case 36:goto _dbac ;case 37:goto _egaa ;case 1:goto _baa ;case 2:goto _baa ;case 38:goto _agg ;case 3:goto _dgb ;case 4:goto _dgb ;case 39:goto _agg ;case 5:goto _dgb ;case 6:goto _dgb ;case 7:goto _dgb ;case 8:goto _baa ;case 40:goto _agg ;case 9:goto _dgb ;case 41:goto _agg ;case 10:goto _baa ;case 42:goto _agg ;case 11:goto _dgb ;case 43:goto _agg ;case 44:goto _agg ;case 45:goto _agg ;case 12:goto _fbg ;case 46:goto _edeg ;case 13:goto _aba ;case 14:goto _baa ;case 15:goto _baa ;case 16:goto _aba ;case 47:goto _gce ;case 17:goto _cef ;case 48:goto _cdg ;case 18:goto _bgg ;case 19:goto _bgg ;case 20:goto _baa ;case 49:goto _abc ;case 50:goto _gddg ;case 21:goto _baa ;case 22:goto _baa ;case 23:goto _baa ;case 24:goto _baa ;case 25:goto _baa ;case 51:goto _gddg ;case 26:goto _cdee ;case 52:goto _gddg ;case 53:goto _gddg ;case 54:goto _gfa ;case 55:goto _abc ;case 56:goto _abc ;case 57:goto _abc ;case 27:goto _addb ;case 28:goto _addb ;case 29:goto _addb ;case 30:goto _addb ;case 31:goto _addb ;case 58:goto _abc ;case 32:goto _baa ;case 59:goto _baa ;case 33:goto _addb ;case 60:goto _abc ;case 61:goto _gce ;case 62:goto _abc ;};};};if _abg > 0{copy (_fff [0:],_fff [_abg :]);};};_ =_eabd ;if _dfdd ==_febe {_aa .Log .Debug ("\u0066o\u0072m\u0061\u0074\u0020\u0070\u0061r\u0073\u0065 \u0065\u0072\u0072\u006f\u0072");};};func IsNumber (data string )(_dga bool ){_bce ,_dgea ,_fdc :=0,0,len (data );_gcg :=len (data );_cead ,_add ,_dcf :=0,0,0;_ =_add ;_ =_dcf ;_ =_cead ;{_bce =_def ;_cead =0;_add =0;_dcf =0;};{if _dgea ==_fdc {goto _gdd ;};switch _bce {case 0:goto _dad ;case 1:goto _bdg ;case 2:goto _efa ;case 3:goto _ccd ;case 4:goto _fab ;case 5:goto _ade ;case 6:goto _ebbf ;case 7:goto _cbd ;};goto _ffg ;_addd :_add =_dgea ;_dgea --;{_dga =false ;};goto _cdf ;_caf :_add =_dgea ;_dgea --;{_dga =_add ==len (data );};goto _cdf ;_dbe :_add =_dgea ;_dgea --;{_dga =_add ==len (data );};goto _cdf ;_ggcf :switch _dcf {case 2:{_dgea =(_add )-1;_dga =_add ==len (data );};case 3:{_dgea =(_add )-1;_dga =false ;};};goto _cdf ;_cdf :_cead =0;if _dgea ++;_dgea ==_fdc {goto _dfc ;};_dad :_cead =_dgea ;switch data [_dgea ]{case 43:goto _ebc ;case 45:goto _ebc ;};if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _bfc ;};goto _eda ;_eda :if _dgea ++;_dgea ==_fdc {goto _bdga ;};_bdg :goto _eda ;_ebc :if _dgea ++;_dgea ==_fdc {goto _gfe ;};_efa :if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _bfc ;};goto _eda ;_bfc :if _dgea ++;_dgea ==_fdc {goto _beb ;};_ccd :if data [_dgea ]==46{goto _eba ;};if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _bfc ;};goto _eda ;_eba :if _dgea ++;_dgea ==_fdc {goto _fec ;};_fab :if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _cage ;};goto _eda ;_cage :if _dgea ++;_dgea ==_fdc {goto _abbf ;};_ade :if data [_dgea ]==69{goto _eff ;};if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _cage ;};goto _eda ;_eff :if _dgea ++;_dgea ==_fdc {goto _cbdc ;};_ebbf :switch data [_dgea ]{case 43:goto _eaae ;case 45:goto _eaae ;};goto _eda ;_eaae :_add =_dgea +1;_dcf =3;goto _bee ;_eaf :_add =_dgea +1;_dcf =2;goto _bee ;_bee :if _dgea ++;_dgea ==_fdc {goto _fc ;};_cbd :if 48<=data [_dgea ]&&data [_dgea ]<=57{goto _eaf ;};goto _eda ;_ffg :_dfc :_bce =0;goto _gdd ;_bdga :_bce =1;goto _gdd ;_gfe :_bce =2;goto _gdd ;_beb :_bce =3;goto _gdd ;_fec :_bce =4;goto _gdd ;_abbf :_bce =5;goto _gdd ;_cbdc :_bce =6;goto _gdd ;_fc :_bce =7;goto _gdd ;_gdd :{};if _dgea ==_gcg {switch _bce {case 1:goto _addd ;case 2:goto _addd ;case 3:goto _caf ;case 4:goto _addd ;case 5:goto _dbe ;case 6:goto _addd ;case 7:goto _ggcf ;};};};if _bce ==_febe {return false ;};return ;};func _gdg (_eea ,_deb float64 ,_aed Format )[]byte {if len (_aed .Whole )==0{return nil ;};_dge :=_c .Date (1899,12,30,0,0,0,0,_c .UTC );_dd :=_dge .Add (_c .Duration (_deb *float64 (24*_c .Hour )));_dd =_befed (_dd );_ca :=_g .AppendFloat (nil ,_eea ,'f',-1,64);_deba :=make ([]byte ,0,len (_ca ));_ccc :=0;_bd :=1;_dc :for _gad :=len (_aed .Whole )-1;_gad >=0;_gad --{_aedg :=len (_ca )-1-_ccc ;_eabf :=_aed .Whole [_gad ];switch _eabf .Type {case FmtTypeDigit :if _aedg >=0{_deba =append (_deba ,_ca [_aedg ]);_ccc ++;_bd =_gad ;}else {_deba =append (_deba ,'0');};case FmtTypeDigitOpt :if _aedg >=0{_deba =append (_deba ,_ca [_aedg ]);_ccc ++;_bd =_gad ;}else {for _eee :=_gad ;_eee >=0;_eee --{_feb :=_aed .Whole [_eee ];if _feb .Type ==FmtTypeLiteral {_deba =append (_deba ,_feb .Literal );};};break _dc ;};case FmtTypeDollar :for _cde :=_ccc ;_cde < len (_ca );_cde ++{_deba =append (_deba ,_ca [len (_ca )-1-_cde ]);_ccc ++;};_deba =append (_deba ,'$');case FmtTypeComma :if !_aed ._adc {_deba =append (_deba ,',');};case FmtTypeLiteral :_deba =append (_deba ,_eabf .Literal );case FmtTypeDate :_deba =append (_deba ,_fgb (_afg (_dd ,_eabf .DateTime ))...);case FmtTypeTime :_deba =append (_deba ,_fgb (_dbc (_dd ,_deb ,_eabf .DateTime ))...);default:_aa .Log .Debug ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0077\u0068\u006f\u006c\u0065 \u0025\u0076",_eabf );};};_abf :=_fgb (_deba );if _ccc < len (_ca )&&(_ccc !=0||_aed ._gb ){_bcf :=len (_ca )-_ccc ;_aga :=make ([]byte ,len (_abf )+_bcf );copy (_aga ,_abf [0:_bd ]);copy (_aga [_bd :],_ca [0:]);copy (_aga [_bd +_bcf :],_abf [_bd :]);_abf =_aga ;};if _aed ._adc {_bf :=_ad .Buffer {};_ceg :=0;for _ccg :=len (_abf )-1;_ccg >=0;_ccg --{if !(_abf [_ccg ]>='0'&&_abf [_ccg ]<='9'){_ceg ++;}else {break ;};};for _df :=0;_df < len (_abf );_df ++{_bef :=(len (_abf )-_df -_ceg );if _bef %3==0&&_bef !=0&&_df !=0{_bf .WriteByte (',');};_bf .WriteByte (_abf [_df ]);};_abf =_bf .Bytes ();};return _abf ;};

// AddToken adds a format token to the format.
func (_gg *Format )AddToken (t FmtType ,l []byte ){if _gg ._cc {_gg ._cc =false ;return ;};switch t {case FmtTypeDecimal :_gg ._gb =true ;case FmtTypeUnderscore :_gg ._cc =true ;case FmtTypeText :_gg .Whole =append (_gg .Whole ,Token {Type :t });case FmtTypeDate ,FmtTypeTime :_gg .Whole =append (_gg .Whole ,Token {Type :t ,DateTime :string (l )});case FmtTypePercent :_gg ._bg =true ;t =FmtTypeLiteral ;l =[]byte {'%'};fallthrough;case FmtTypeDigitOpt :fallthrough;case FmtTypeLiteral ,FmtTypeDigit ,FmtTypeDollar ,FmtTypeComma :if l ==nil {l =[]byte {0};};for _ ,_fd :=range l {if _gg .IsExponential {_gg .Exponent =append (_gg .Exponent ,Token {Type :t ,Literal :_fd });}else if !_gg ._gb {_gg .Whole =append (_gg .Whole ,Token {Type :t ,Literal :_fd });}else {_gg .Fractional =append (_gg .Fractional ,Token {Type :t ,Literal :_fd });};};case FmtTypeDigitOptThousands :_gg ._adc =true ;case FmtTypeFraction :_de :=_eb .Split (string (l ),"\u002f");if len (_de )==2{_gg ._fga =true ;_gg ._af ,_ =_g .ParseInt (_de [1],10,64);for _ ,_fa :=range _de [1]{if _fa =='?'||_fa =='0'{_gg ._ba ++;};};};default:_aa .Log .Debug ("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0070\u0068\u0020\u0074y\u0070\u0065\u0020\u0069\u006e\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0076",t );};};func _fba (_egc []byte )[]byte {for _dcg :=len (_egc )-1;_dcg > 0;_dcg --{if _egc [_dcg ]=='9'+1{_egc [_dcg ]='0';if _egc [_dcg -1]=='.'{_dcg --;};_egc [_dcg -1]++;};};if _egc [0]=='9'+1{_egc [0]='0';copy (_egc [1:],_egc [0:]);_egc [0]='1';};return _egc ;};

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String (v string ,f string )string {_cd :=Parse (f );var _ce Format ;if len (_cd )==1{_ce =_cd [0];}else if len (_cd )==4{_ce =_cd [3];};_ff :=false ;for _ ,_gbd :=range _ce .Whole {if _gbd .Type ==FmtTypeText {_ff =true ;};};if !_ff {return v ;};_edee :=_ad .Buffer {};for _ ,_bga :=range _ce .Whole {switch _bga .Type {case FmtTypeLiteral :_edee .WriteByte (_bga .Literal );case FmtTypeText :_edee .WriteString (v );};};return _edee .String ();};func _gbg (_abb []byte )[]byte {_aaa :=len (_abb );_gae :=false ;_abbb :=false ;for _faeg :=len (_abb )-1;_faeg >=0;_faeg --{if _abb [_faeg ]=='0'&&!_abbb &&!_gae {_aaa =_faeg ;}else if _abb [_faeg ]=='.'{_gae =true ;}else {_abbb =true ;};};if _gae &&_abbb {if _abb [_aaa -1]=='.'{_aaa --;};return _abb [0:_aaa ];};return _abb ;};

// FmtType is the type of a format token.
//go:generate stringer -type=FmtType
type FmtType byte ;

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric (v float64 )string {if _gf .Abs (v )>=_d ||_gf .Abs (v )<=_b &&v !=0{return _cf (v );};_dda :=make ([]byte ,0,15);_dda =_g .AppendFloat (_dda ,v ,'f',-1,64);if len (_dda )> 11{_cb :=_dda [11]-'0';if _cb >=5&&_cb <=9{_dda [10]++;_dda =_dda [0:11];_dda =_fba (_dda );};_dda =_dda [0:11];}else if len (_dda )==11{if _dda [len (_dda )-1]=='9'{_dda [len (_dda )-1]++;_dda =_fba (_dda );};};_dda =_gbg (_dda );return string (_dda );};func _afg (_cgcd _c .Time ,_bdf string )[]byte {_aeb :=[]byte {};_deg :=0;for _cea :=0;_cea < len (_bdf );_cea ++{var _cgcdd string ;if _bdf [_cea ]=='/'{_cgcdd =string (_bdf [_deg :_cea ]);_deg =_cea +1;}else if _cea ==len (_bdf )-1{_cgcdd =string (_bdf [_deg :_cea +1]);}else {continue ;};switch _cgcdd {case "\u0079\u0079":_aeb =_cgcd .AppendFormat (_aeb ,"\u0030\u0036");case "\u0079\u0079\u0079\u0079":_aeb =_cgcd .AppendFormat (_aeb ,"\u0032\u0030\u0030\u0036");case "\u006d":_aeb =_cgcd .AppendFormat (_aeb ,"\u0031");case "\u006d\u006d":_aeb =_cgcd .AppendFormat (_aeb ,"\u0030\u0031");case "\u006d\u006d\u006d":_aeb =_cgcd .AppendFormat (_aeb ,"\u004a\u0061\u006e");case "\u006d\u006d\u006d\u006d":_aeb =_cgcd .AppendFormat (_aeb ,"\u004aa\u006e\u0075\u0061\u0072\u0079");case "\u006d\u006d\u006dm\u006d":switch _cgcd .Month (){case _c .January ,_c .July ,_c .June :_aeb =append (_aeb ,'J');case _c .February :_aeb =append (_aeb ,'M');case _c .March ,_c .May :_aeb =append (_aeb ,'M');case _c .April ,_c .August :_aeb =append (_aeb ,'A');case _c .September :_aeb =append (_aeb ,'S');case _c .October :_aeb =append (_aeb ,'O');case _c .November :_aeb =append (_aeb ,'N');case _c .December :_aeb =append (_aeb ,'D');};case "\u0064":_aeb =_cgcd .AppendFormat (_aeb ,"\u0032");case "\u0064\u0064":_aeb =_cgcd .AppendFormat (_aeb ,"\u0030\u0032");case "\u0064\u0064\u0064":_aeb =_cgcd .AppendFormat (_aeb ,"\u004d\u006f\u006e");case "\u0064\u0064\u0064\u0064":_aeb =_cgcd .AppendFormat (_aeb ,"\u004d\u006f\u006e\u0064\u0061\u0079");default:_aa .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_cgcdd );};if _bdf [_cea ]=='/'{_aeb =append (_aeb ,'/');};};return _aeb ;};