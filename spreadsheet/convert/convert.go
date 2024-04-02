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

package convert ;import (_dd "github.com/unidoc/unioffice/common/logger";_eg "github.com/unidoc/unioffice/common/tempstorage";_dc "github.com/unidoc/unioffice/internal/convertutils";_ed "github.com/unidoc/unioffice/measurement";_cf "github.com/unidoc/unioffice/schema/soo/dml";
_bb "github.com/unidoc/unioffice/schema/soo/dml/chart";_f "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_b "github.com/unidoc/unioffice/schema/soo/sml";_gf "github.com/unidoc/unioffice/spreadsheet";_c "github.com/unidoc/unioffice/spreadsheet/formula";
_e "github.com/unidoc/unioffice/spreadsheet/reference";_dg "github.com/unidoc/unipdf/v3/creator";_gg "github.com/unidoc/unipdf/v3/model";_da "image";_g "strconv";);type colWidthRange struct{_abgb int ;_cbbf int ;_gdg float64 ;_aegb *style ;};type border struct{_fdcc float64 ;
_fffdd _dg .Color ;};func _cbae (_fee *symbol ){_dgd :=_dg .New ();_aaf :=_dgd .NewStyledParagraph ();_aaf .SetMargins (0,0,0,0);_gcg :=_aaf .Append (_fee ._acgga );if _fee ._adb !=nil {_gcg .Style =*_fee ._adb ;};_fee ._aefce =_aaf .Height ();if _fee ._dgbc ==0{_fee ._dgbc =_aaf .Width ();
};};type rowspan struct{_feb float64 ;_egdd int ;_dee int ;};type symbol struct{_acgga string ;_bfcd float64 ;_aefce float64 ;_dgbc float64 ;_adb *_dg .TextStyle ;_ggac string ;};func (_bag *convertContext )getBorder (_fcfc *_b .CT_BorderPr )*border {_ffag :=&border {};
switch _fcfc .StyleAttr {case _b .ST_BorderStyleThin :_ffag ._fdcc =_af ;case _b .ST_BorderStyleMedium :_ffag ._fdcc =_af *2;case _b .ST_BorderStyleThick :_ffag ._fdcc =_af *4;};if _ffag ._fdcc ==0.0{return nil ;};if _eag :=_fcfc .Color ;_eag !=nil {_ebea :=_bag .getColorStringFromSmlColor (_eag );
if _ebea !=nil {_ffag ._fffdd =_dg .ColorRGBFromHex (*_ebea );}else {_ffag ._fffdd =_dg .ColorBlack ;};};return _ffag ;};func (_cab *convertContext )addRowToPage (_acfc []*cell ,_eedd int ){_ecfag :=0.0;_cdc :=_cab ._fgee ;for _ ,_dab :=range _acfc {if len (_dab ._ebg )!=0{_dab ._dea =_ecfag ;
_ecfag =_dab ._eegc +_dab ._caba ;};};for _ebb :=len (_acfc )-1;_ebb >=0;_ebb --{_cfa :=_acfc [_ebb ];if len (_cfa ._ebg )!=0{_cfa ._bcfe =_cdc ;_cdc =_cfa ._eegc ;};};_cab ._bdga ._eeec =append (_cab ._bdga ._eeec ,&pageRow {_ccbcc :_eedd ,_gbefe :_acfc });
};func (_aga *convertContext )getSymbolsFromR (_bbc []*_b .CT_RElt ,_aagde *style )[]*symbol {_gacf :=[]*symbol {};for _ ,_dacf :=range _bbc {_bdbd :=_aga .combineCellStyleWithRPrElt (_aagde ,_dacf .RPr );for _ ,_eac :=range _dacf .T {_gacf =append (_gacf ,&symbol {_acgga :string (_eac ),_adb :_aga .makeTextStyleFromCellStyle (_bdbd )});
};};return _gacf ;};var _af =_agaga (0.0625);func _cecd (_beca *bool )bool {return _beca !=nil &&*_beca };func (_aac *convertContext )makeRowspans (){var _fgbg float64 ;_ggb :=0;for _ddga ,_gfc :=range _aac ._fgf {_bgc :=_gfc ._bbbb +_gfc ._fdce ;if _fgbg +_bgc <=_aac ._gcaaa {_gfc ._gfde =_fgbg ;
_fgbg +=_bgc ;}else {_aac ._afad =append (_aac ._afad ,&rowspan {_feb :_fgbg ,_egdd :_ggb ,_dee :_ddga });_ggb =_ddga ;_gfc ._gfde =0;_fgbg =_bgc ;};};_aac ._afad =append (_aac ._afad ,&rowspan {_feb :_fgbg ,_egdd :_ggb ,_dee :len (_aac ._fgf )});};func (_afe *convertContext )drawSheet (_fdf bool ){for _bba ,_ega :=range _afe ._acgg {_ebd :=len (_ega ._gdcc );
if _bba ==len (_afe ._acgg )-1{for _abg :=len (_ega ._gdcc )-1;_abg >=0;_abg --{if !_ega ._gdcc [_abg ]._dbg {_ebd =_abg ;};};};_beg :=_ega ._gdcc [:_ebd ];for _ ,_ceg :=range _beg {_afe ._fcc .NewPage ();if _fdf {_afe ._fcc .RotateDeg (270);};_afe .drawPage (_ceg );
};};};const _ff =0.64;const _cc =1.5;func (_fbce *convertContext )makeTextStyleFromCellStyle (_ddeg *style )*_dg .TextStyle {_bdca :=_fbce ._fcc .NewTextStyle ();if _ddeg ==nil {_bdca .FontSize =_dc .DefaultFontSize ;_bdca .Font =_dc .AssignStdFontByName (_bdca ,_dc .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][FontStyle_Regular ]);
return &_bdca ;};if _cecd (_ddeg ._baee ){_bdca .Underline =true ;_bdca .UnderlineStyle =_dg .TextDecorationLineStyle {Offset :0.5,Thickness :_agaga (1/32)};};var _fgea FontStyle ;if _cecd (_ddeg ._dcdc )&&_cecd (_ddeg ._gfaec ){_fgea =FontStyle_BoldItalic ;
}else if _cecd (_ddeg ._dcdc ){_fgea =FontStyle_Bold ;}else if _cecd (_ddeg ._gfaec ){_fgea =FontStyle_Italic ;}else {_fgea =FontStyle_Regular ;};_gfgf :="\u0064e\u0066\u0061\u0075\u006c\u0074";if _ddeg ._gaa !=nil {_gfgf =*_ddeg ._gaa ;};if _eedg ,_cgcf :=_dc .StdFontsMap [_gfgf ];
_cgcf {_bdca .Font =_dc .AssignStdFontByName (_bdca ,_eedg [_fgea ]);}else if _cdfcb :=_dc .GetRegisteredFont (_gfgf ,_fgea );_cdfcb !=nil {_bdca .Font =_cdfcb ;}else {_dd .Log .Debug ("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e",_gfgf ,_fgea );
_bdca .Font =_dc .AssignStdFontByName (_bdca ,_dc .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][_fgea ]);};if _ddeg ._gef !=nil {_bdca .FontSize =*_ddeg ._gef ;};if _ddeg ._gddc !=nil {_bdca .Color =_dg .ColorRGBFromHex (*_ddeg ._gddc );};if _ddeg ._bgbde !=nil &&*_ddeg ._bgbde {_bdca .FontSize *=_ff ;
}else if _ddeg ._bcfd !=nil &&*_ddeg ._bcfd {_bdca .FontSize *=_ff ;};return &_bdca ;};const _ea =0.25;const _cd =3;const _ga =2;

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle =_dc .FontStyle ;func (_dced *convertContext )makeMergedCells (){_ae :=[]*mergedCell {};for _ ,_eebb :=range _dced ._agc .MergedCells (){_efg ,_gfa ,_ccg :=_e .ParseRangeReference (_eebb .Reference ());if _ccg !=nil {_dd .Log .Debug ("\u0065\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c: \u0025\u0073",_ccg );
continue ;};_gfg :=mergedCell {_gafb :_efg .RowIdx ,_dgcg :_efg .ColumnIdx ,_agca :_gfa .RowIdx ,_dcdcg :_gfa .ColumnIdx };for _fe :=_gfg ._dgcg -1;_fe < _gfg ._dcdcg ;_fe ++{_gfg ._aeec +=_dced ._gdcf [_fe ]._fcf ;};for _gee :=_gfg ._gafb -1;_gee < _gfg ._agca ;
_gee ++{_gfg ._acga +=_dced ._fgf [_gee ]._bbbb ;};_ae =append (_ae ,&_gfg );};_dced ._dcc =_ae ;};type pageRow struct{_ccbcc int ;_gbefe []*cell ;};func (_bfd *convertContext )makePagespans (){_bfd ._acgg =[]*pagespan {};_acf :=0.0;_bea :=0;for _df ,_ggfg :=range _bfd ._gdcf {_bcc :=_ggfg ._fcf ;
if _acf +_bcc <=_bfd ._fgee {_ggfg ._cgge =_acf ;_acf +=_bcc ;}else {_ggfg ._cgge =0;_bfd ._acgg =append (_bfd ._acgg ,&pagespan {_bbfe :_acf ,_edbb :_bea ,_bcgfg :_df });_acf =_bcc ;_bea =_df ;};};_bfd ._acgg =append (_bfd ._acgg ,&pagespan {_bbfe :_acf ,_edbb :_bea ,_bcgfg :len (_bfd ._gdcf )});
};func (_bdd *convertContext )getSymbolsFromString (_bfe string ,_fed *style )[]*symbol {_fcd :=[]*symbol {};_bbda :=_bdd .makeTextStyleFromCellStyle (_fed );for _ ,_ebbg :=range _bfe {_fcd =append (_fcd ,&symbol {_acgga :string (_ebbg ),_adb :_bbda });
};return _fcd ;};type mergedCell struct{_gafb uint32 ;_dgcg uint32 ;_agca uint32 ;_dcdcg uint32 ;_aeec float64 ;_acga float64 ;};func (_dgfb *convertContext )fillPages (){for _aeg ,_ebe :=range _dgfb ._afad {_ade :=_dgfb ._fgf [_ebe ._egdd :_ebe ._dee ];
for _cgg ,_ebaa :=range _ade {_fcb :=0;_acfg :=0.0;_dgce :=[]*cell {};if _ebaa ._beac {for _ ,_eefd :=range _ebaa ._feeb {_ecbd :=_dgfb ._acgg [_fcb ];_dgfb ._bdga =_ecbd ._gdcc [_aeg ];_dgfb ._bdga ._dbg =true ;_dfd :=_eefd ._fdag ;if _acfg +_dfd > _ecbd ._bbfe {_dgfb .addRowToPage (_dgce ,_cgg );
_dgce =[]*cell {_eefd };_acfg =_dfd ;_fcb ++;}else {_eefd ._eegc =_acfg ;_dgce =append (_dgce ,_eefd );_acfg +=_dfd ;};};if len (_dgce )> 0{_fdd :=_dgfb ._acgg [_fcb ];_dgfb ._bdga =_fdd ._gdcc [_aeg ];_dgfb ._bdga ._dbg =true ;_dgfb .addRowToPage (_dgce ,_cgg );
};};};};};func (_dcf *convertContext )getContentFromCell (_ffg *_gf .Sheet ,_gcaa _gf .Cell ,_cagd *style ,_bbe float64 ,_cad bool )([]*line ,_b .ST_CellType ){_cbbd :=_gcaa .X ();var _abf []*symbol ;switch _cbbd .TAttr {case _b .ST_CellTypeS :_bcgf :=_cbbd .V ;
if _bcgf !=nil {_gfae ,_fbc :=_g .Atoi (*_bcgf );if _fbc ==nil {_beef :=_dcf ._dfg .SharedStrings .X ().Si [_gfae ];if _beef .T !=nil {_abf =_dcf .getSymbolsFromString (*_beef .T ,_cagd );}else if _beef .R !=nil {_abf =_dcf .getSymbolsFromR (_beef .R ,_cagd );
};};};case _b .ST_CellTypeB :_abd :=_cbbd .V ;if _abd !=nil {if *_abd =="\u0030"{_abf =_dcf .getSymbolsFromString ("\u0046\u0041\u004cS\u0045",_cagd );}else {_abf =_dcf .getSymbolsFromString ("\u0054\u0052\u0055\u0045",_cagd );};};case _b .ST_CellTypeStr :if _cbbd .F !=nil {_aed :=_c .NewEvaluator ();
_eeff :=_ffg .FormulaContext ().Cell (_gcaa .Reference (),_aed );_abf =_dcf .getSymbolsFromString (_eeff .Value (),_cagd );};default:_abf =_dcf .getSymbolsFromString (_gcaa .GetFormattedValue (),_cagd );};_cfb :=0.0;_dag :=0.0;var _fba []*line ;var _dac bool ;
if _cagd !=nil {if _cagd ._bgbde !=nil {if *_cagd ._bgbde {_dac =true ;};};if _cagd ._bcfd !=nil {if *_cagd ._bcfd {_dac =true ;};};};if _cad {_fba =[]*line {};_bbac :=_bbe -2*_cd ;_aefc :=[]*symbol {};for _ ,_ccac :=range _abf {_cbae (_ccac );if _cfb +_ccac ._dgbc >=_bbac {_beec :=_eafd (_aefc );
if _dac {_beec /=_ff ;};_fba =append (_fba ,&line {_gced :_dag ,_ffgf :_aefc ,_cbe :_beec });_aefc =[]*symbol {_ccac };_cfb =_ccac ._dgbc ;_dag +=_beec ;}else {_ccac ._bfcd =_cfb ;_cfb +=_ccac ._dgbc ;_aefc =append (_aefc ,_ccac );};};_dbag :=_eafd (_aefc );
if _dac {_dbag /=_ff ;};if len (_aefc )> 0{_fba =append (_fba ,&line {_gced :_dag ,_ffgf :_aefc ,_cbe :_dbag });};}else {for _ ,_dcd :=range _abf {_cbae (_dcd );_dcd ._bfcd =_cfb ;_cfb +=_dcd ._dgbc ;};if len (_abf )> 0{_fba =[]*line {&line {_ffgf :_abf ,_cbe :_eafd (_abf )}};
};};_acfcg :=_cbbd .TAttr ;if _acfcg ==_b .ST_CellTypeUnset {_acfcg =_b .ST_CellTypeN ;};return _fba ,_acfcg ;};type style struct{_gddc *string ;_gef *float64 ;_gaa *string ;_dcdc *bool ;_gfaec *bool ;_baee *bool ;_bgbde *bool ;_bcfd *bool ;_dadd *border ;
_ggbf *border ;_efbdf *border ;_bad *border ;_eedb bool ;_bgdc _b .ST_VerticalAlignment ;_afc _b .ST_HorizontalAlignment ;};type page struct{_eeec []*pageRow ;_dbg bool ;_cge []*_dg .Image ;_cfgfa *pagespan ;_fce *rowspan ;};func (_baea *convertContext )combineCellStyleWithRPrElt (_bab *style ,_bcb *_b .CT_RPrElt )*style {_egbe :=*_bab ;
_fccf :=_baea .getStyleFromRPrElt (_bcb );if _fccf ==nil {return &_egbe ;};if _fccf ._gddc !=nil {_egbe ._gddc =_fccf ._gddc ;};if _fccf ._gef !=nil {_egbe ._gef =_fccf ._gef ;};if _fccf ._gaa !=nil {_egbe ._gaa =_fccf ._gaa ;};if _fccf ._dcdc !=nil {_egbe ._dcdc =_fccf ._dcdc ;
};if _fccf ._gfaec !=nil {_egbe ._gfaec =_fccf ._gfaec ;};if _fccf ._baee !=nil {_egbe ._baee =_fccf ._baee ;};if _fccf ._bgbde !=nil {_egbe ._bgbde =_fccf ._bgbde ;};if _fccf ._bcfd !=nil {_egbe ._bcfd =_fccf ._bcfd ;};return &_egbe ;};func (_gfd *convertContext )alignSymbolsVertically (_cgcg *cell ,_bgb _b .ST_VerticalAlignment ){var _bcac float64 ;
switch _bgb {case _b .ST_VerticalAlignmentTop :_bcac =_ga ;if _cgcg ._gae {_bcac -=_cc ;}else if _cgcg ._ddb {_bcac +=4*_cc ;};for _ ,_ccfg :=range _cgcg ._ebg {_bcac +=_ccfg ._cbe ;_ccfg ._gced =_bcac ;_bcac +=_cbf ;};case _b .ST_VerticalAlignmentCenter :_befb :=0.0;
for _ ,_cdg :=range _cgcg ._ebg {_befb +=_cdg ._cbe +_agaga (1);};_bcac =0.5*(_cgcg ._bcfb -_befb );if _cgcg ._gae {_bcac -=2*_cc ;}else if _cgcg ._ddb {_bcac +=2*_cc ;};for _ ,_cde :=range _cgcg ._ebg {_bcac +=_cde ._cbe +0.5*_cbf ;_cde ._gced =_bcac ;
_bcac +=0.5*_cbf ;};default:_bcac =_cgcg ._bcfb -_ga ;if _cgcg ._gae {_bcac -=4*_cc ;}else if _cgcg ._ddb {_bcac +=_cc ;};for _cdb :=len (_cgcg ._ebg )-1;_cdb >=0;_cdb --{_cgcg ._ebg [_cdb ]._gced =_bcac ;_bcac -=_cgcg ._ebg [_cdb ]._cbe ;_bcac -=_cbf ;
};};};func (_gdac *convertContext )drawPage (_effae *page ){_abb :=_gdac ._bbf ;_aadf :=_gdac ._ffbd ;for _ ,_gddb :=range _effae ._eeec {_eed :=_gdac ._fgf [_gddb ._ccbcc ];for _ ,_ccb :=range _gddb ._gbefe {_gaf :=_ccb ._dea < _ccb ._eegc ;_ccf :=_ccb ._bcfe > _ccb ._eegc +_ccb ._caba ;
var _efb ,_acc bool ;for _ ,_cedd :=range _ccb ._ebg {for _ ,_fbg :=range _cedd ._ffgf {if _gaf &&!_efb {_efb =_fbg ._bfcd < 0;};if _ccf &&!_acc {_acc =_ccb ._caba < _fbg ._bfcd +_fbg ._dgbc ;};if _ccb ._eegc +_fbg ._bfcd >=_ccb ._dea &&_ccb ._eegc +_fbg ._bfcd +_fbg ._dgbc <=_ccb ._bcfe {_feg :=_gdac ._fcc .NewStyledParagraph ();
_eabd :=_aadf +_ccb ._eegc +_fbg ._bfcd ;_dge :=_abb +_eed ._gfde +_cedd ._gced -_fbg ._aefce -_agaga (0.5);_feg .SetPos (_eabd ,_dge );var _gfga *_dg .TextChunk ;if _fbg ._ggac !=""{_gfga =_feg .AddExternalLink (_fbg ._acgga ,_fbg ._ggac );}else {_gfga =_feg .Append (_fbg ._acgga );
};if _fbg ._adb !=nil {_gfga .Style =*_fbg ._adb ;};_gdac ._fcc .Draw (_feg );};};};var _eede ,_ccbc ,_dgfba ,_cbg ,_bcg ,_eefc float64 ;var _cba ,_acaf ,_fec ,_ceeb _dg .Color ;if _cdff :=_ccb ._dgbd ;_cdff !=nil {_eede =_cdff ._fdcc ;_cba =_cdff ._fffdd ;
};if _fdbg :=_ccb ._dbe ;_fdbg !=nil {_ccbc =_fdbg ._fdcc ;_acaf =_fdbg ._fffdd ;};if _dae :=_ccb ._gefd ;_dae !=nil {_dgfba =_dae ._fdcc ;_bcg =_dgfba /2;_fec =_dae ._fffdd ;};if _bff :=_ccb ._afg ;_bff !=nil {_cbg =_bff ._fdcc ;_eefc =_cbg /2;_ceeb =_bff ._fffdd ;
};var _eeg float64 ;if _gddb ._ccbcc > 1{_eeg =_gdac ._fgf [_gddb ._ccbcc -1]._fdce ;};_gfcd :=_abb +_eed ._gfde -0.5*(_eeg -_eede );_cafa :=_abb +_eed ._gfde +_eed ._bbbb +0.5*(_eed ._fdce +_ccbc );_bfdf :=_aadf +_ccb ._eegc ;_efge :=_bfdf +_ccb ._fdag ;
_dc .DrawLine (_gdac ._fcc ,_bfdf ,_gfcd ,_efge ,_gfcd ,_eede ,_cba );_dc .DrawLine (_gdac ._fcc ,_bfdf ,_cafa ,_efge ,_cafa ,_ccbc ,_acaf );if !_efb {_dc .DrawLine (_gdac ._fcc ,_bfdf -_bcg ,_gfcd ,_bfdf -_bcg ,_cafa ,_dgfba ,_fec );};if !_acc {_dc .DrawLine (_gdac ._fcc ,_efge -_eefc ,_gfcd ,_efge -_eefc ,_cafa ,_cbg ,_ceeb );
};};};for _ ,_dcedb :=range _effae ._cge {if _dcedb !=nil {_gdac ._fcc .Draw (_dcedb );};};};var _bfc =[]string {"\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0030\u0030","\u0030\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0038\u0030\u0030\u0030","\u0038\u0030\u0030\u0030\u0038\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0063\u0030\u0063\u0030\u0063\u0030","\u0038\u0030\u0038\u0030\u0038\u0030","\u0039\u0039\u0039\u0039\u0066\u0066","\u0039\u0039\u0033\u0033\u0036\u0036","\u0066\u0066\u0066\u0066\u0063\u0063","\u0063\u0063\u0066\u0066\u0066\u0066","\u0036\u0036\u0030\u0030\u0036\u0036","\u0066\u0066\u0038\u0030\u0038\u0030","\u0030\u0030\u0036\u0036\u0063\u0063","\u0063\u0063\u0063\u0063\u0066\u0066","\u0030\u0030\u0030\u0030\u0038\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0030\u0030\u0063\u0063\u0066\u0066","\u0063\u0063\u0066\u0066\u0066\u0066","\u0063\u0063\u0066\u0066\u0063\u0063","\u0066\u0066\u0066\u0066\u0039\u0039","\u0039\u0039\u0063\u0063\u0066\u0066","\u0066\u0066\u0039\u0039\u0063\u0063","\u0063\u0063\u0039\u0039\u0066\u0066","\u0066\u0066\u0063\u0063\u0039\u0039","\u0033\u0033\u0036\u0036\u0066\u0066","\u0033\u0033\u0063\u0063\u0063\u0063","\u0039\u0039\u0063\u0063\u0030\u0030","\u0066\u0066\u0063\u0063\u0030\u0030","\u0066\u0066\u0039\u0039\u0030\u0030","\u0066\u0066\u0036\u0036\u0030\u0030","\u0036\u0036\u0036\u0036\u0039\u0039","\u0039\u0036\u0039\u0036\u0039\u0036","\u0030\u0030\u0033\u0033\u0036\u0036","\u0033\u0033\u0039\u0039\u0036\u0036","\u0030\u0030\u0033\u0033\u0030\u0030","\u0033\u0033\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0036\u0036","\u0033\u0033\u0033\u0033\u0039\u0039","\u0033\u0033\u0033\u0033\u0033\u0033"};
func (_bead *convertContext )makePages (){for _ ,_cfgde :=range _bead ._acgg {for _ ,_eba :=range _bead ._afad {_cfgde ._gdcc =append (_cfgde ._gdcc ,&page {_eeec :[]*pageRow {},_cfgfa :_cfgde ,_fce :_eba });};};};type cell struct{_gde _b .ST_CellType ;
_egdg int ;_eegc float64 ;_ebg []*line ;_caba float64 ;_fdag float64 ;_bcfb float64 ;_dea float64 ;_bcfe float64 ;_deeg *_dg .TextStyle ;_dgbd *border ;_dbe *border ;_gefd *border ;_afg *border ;_gae bool ;_ddb bool ;};type pagespan struct{_bbfe float64 ;
_gdcc []*page ;_edbb int ;_bcgfg int ;};type colInfo struct{_cgge float64 ;_fcf float64 ;_beaf *style ;};func (_cfgd *convertContext )makeCells (){_bf :=_cfgd ._agc ;_fffa :=_bf .Rows ();_fbe :=0;for _beea ,_eab :=range _cfgd ._fgf {if _beea < _cfgd ._bdac ||(_beea > _cfgd ._caeg &&_cfgd ._caeg > 0){continue ;
};_eab ._feeb =[]*cell {};_aag :=0.0;_ggdd :=_eab ._ccea ;if _eab ._beac {_gac :=_fffa [_fbe ];_fbe ++;_gacd :=_eab ._bbbb ;for _ ,_fc :=range _gac .Cells (){_dgf ,_effa :=_e .ParseCellReference (_fc .Reference ());if _effa !=nil {_dd .Log .Debug ("\u0043\u0061\u006e\u006eo\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0061\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u003a \u0025\u0073",_effa );
continue ;};if int (_dgf .ColumnIdx )< _cfgd ._fffd ||(int (_dgf .ColumnIdx )> _cfgd ._fbgb &&_cfgd ._fbgb > 0){continue ;};_dbf :=_cfgd ._gdcf [_dgf .ColumnIdx ];_agga :=_dbf ._fcf ;_gdc :=_agga ;_dec :=_dbf ._beaf ;var _cbc ,_afd ,_gab ,_bfa bool ;for _ ,_eb :=range _cfgd ._dcc {if _dgf .RowIdx >=_eb ._gafb &&_dgf .RowIdx <=_eb ._agca &&_dgf .ColumnIdx >=_eb ._dgcg &&_dgf .ColumnIdx <=_eb ._dcdcg {if _dgf .ColumnIdx ==_eb ._dgcg &&_dgf .RowIdx ==_eb ._gafb {_agga =_eb ._aeec ;
_gacd =_eb ._acga ;};_cbc =_dgf .RowIdx !=_eb ._gafb ;_afd =_dgf .RowIdx !=_eb ._agca ;_gab =_dgf .ColumnIdx !=_eb ._dgcg ;_bfa =_dgf .ColumnIdx !=_eb ._dcdcg ;};};_aeee :=_cfgd .getStyleFromCell (_fc ,_ggdd ,_dec );var _def ,_cgd ,_dca bool ;var _efff ,_deb ,_fgd ,_dadg *border ;
var _gbgb _b .ST_VerticalAlignment ;var _gbfd _b .ST_HorizontalAlignment ;if _aeee !=nil {if !_cbc {_efff =_aeee ._dadd ;};if !_afd {_deb =_aeee ._ggbf ;};if !_gab {_fgd =_aeee ._efbdf ;};if !_bfa {_dadg =_aeee ._bad ;};if _deb !=nil &&_deb ._fdcc > _aag {_aag =_deb ._fdcc ;
};_gbgb =_aeee ._bgdc ;_gbfd =_aeee ._afc ;if _aeee ._bgbde !=nil {_def =*_aeee ._bgbde ;};if _aeee ._bcfd !=nil {_cgd =*_aeee ._bcfd ;};_dca =_aeee ._eedb ;};_fdcb ,_eca :=_cfgd .getContentFromCell (_bf ,_fc ,_aeee ,_agga ,_dca );_bdb :=&cell {_gde :_eca ,_caba :_agga ,_fdag :_gdc ,_bcfb :_gacd ,_ebg :_fdcb ,_dgbd :_efff ,_dbe :_deb ,_gefd :_fgd ,_afg :_dadg ,_gae :_def ,_ddb :_cgd };
_cfgd .alignSymbolsHorizontally (_bdb ,_gbfd );_cfgd .alignSymbolsVertically (_bdb ,_gbgb );_eab ._feeb =append (_eab ._feeb ,_bdb );};};_eab ._fdce =_aag ;};};func (_effaf *convertContext )getStyleFromCell (_efa _gf .Cell ,_cbba ,_egf *style )*style {_ccef :=_efa .X ();
_egfb :=_effaf .getStyle (_ccef .SAttr );_aede (_egfb ,_cbba );_aede (_egfb ,_egf );return _egfb ;};type line struct{_gced float64 ;_ffgf []*symbol ;_cbe float64 ;};func _cdffa (_fda []*symbol )float64 {_fedg :=0.0;for _ ,_accb :=range _fda {_fedg +=_accb ._dgbc ;
};return _fedg ;};

// ConvertToPdf converts a sheet to a PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf (s *_gf .Sheet )*_dg .Creator {_dgg :=s .X ();if _dgg ==nil {return nil ;};var _edf _dg .PageSize ;_cfc :=true ;_ab :=false ;if _dgg .SheetPr !=nil &&_dgg .SheetPr .PageSetUpPr !=nil &&_dgg .SheetPr .PageSetUpPr .FitToPageAttr !=nil &&*_dgg .SheetPr .PageSetUpPr .FitToPageAttr {_ab =true ;
};if _ac :=_dgg .PageSetup ;_ac !=nil {_cfc =_ac .OrientationAttr ==_b .ST_OrientationLandscape ;if _fff :=_ac .PaperSizeAttr ;_fff !=nil {_edf =_bcdf [*_fff ];};};if (_edf ==_dg .PageSize {}){_edf =_dg .PageSizeA4 ;};if _cfc {_edf [0],_edf [1]=_edf [1],_edf [0];
};_cfg :=_dg .New ();_cfg .SetPageSize (_edf );var _ec ,_cbb ,_dce ,_edg float64 ;if _bg :=_dgg .PageMargins ;_bg !=nil {_dce =_bg .LeftAttr ;_edg =_bg .RightAttr ;_ec =_bg .TopAttr ;_cbb =_bg .BottomAttr ;};if _dce < _ea {_dce =_ea ;};if _edg < _ea {_edg =_ea ;
};if _ec < _be {_ec =_be ;};if _cbb < _be {_cbb =_be ;};_ec *=_ed .Inch ;_cbb *=_ed .Inch ;_dce *=_ed .Inch ;_edg *=_ed .Inch ;_cfg .SetPageMargins (_dce ,_edg ,_ec ,_cbb );_gc :=s .Workbook ();var _de *_cf .Theme ;if len (_gc .Themes ())> 0{_de =_gc .Themes ()[0];
};var _fg ,_egb ,_cfgf ,_dgc int ;for _ ,_caf :=range _gc .DefinedNames (){if _caf .Name ()=="\u005f\u0078l\u006e\u006d\u002eP\u0072\u0069\u006e\u0074\u005f\u0041\u0072\u0065\u0061"{_cg ,_ce ,_ag ,_ced :=_dc .ParseExcelRange (_caf .Content ());if _ced ==nil &&s .Name ()==_cg {_fg =int (_ce .ColumnIdx );
_egb =int (_ag .ColumnIdx );_cfgf =int (_ce .RowIdx );_dgc =int (_ag .RowIdx );};};};_ge :=&convertContext {_fcc :_cfg ,_agc :s ,_dfg :s .Workbook (),_eggae :_de ,_bdf :&s .Workbook ().StyleSheet ,_bbf :_ec ,_ffbd :_dce ,_gcaaa :_edf [1]-_cbb -_ec ,_fgee :_edf [0]-_edg -_dce ,_fffd :_fg ,_fbgb :_egb ,_bdac :_cfgf ,_caeg :_dgc ,_egbg :_ab };
_ge .makeAnchors ();_ge .determineMaxIndexes ();if _ge ._ecbde ==0&&_ge ._bgbd ==0{_cfg .NewPage ();return _cfg ;};_ge .makeCols ();_ge .makeRows ();_ge .makeMergedCells ();_ge .makeCells ();_ge .makePagespans ();_ge .makeRowspans ();_ge .makePages ();
_ge .fillPages ();_ge .distributeAnchors ();_ge .drawSheet (_cfc );return _cfg ;};var _bcdf =map[uint32 ]_dg .PageSize {1:_dg .PageSize {8.5*_ed .Inch ,11*_ed .Inch },2:_dg .PageSize {8.5*_ed .Inch ,11*_ed .Inch },3:_dg .PageSize {11*_ed .Inch ,17*_ed .Inch },4:_dg .PageSize {17*_ed .Inch ,11*_ed .Inch },5:_dg .PageSize {8.5*_ed .Inch ,14*_ed .Inch },6:_dg .PageSize {5.5*_ed .Inch ,8.5*_ed .Inch },7:_dg .PageSize {7.5*_ed .Inch ,10*_ed .Inch },8:_dg .PageSize {_agaga (297),_agaga (420)},9:_dg .PageSize {_agaga (210),_agaga (297)},10:_dg .PageSize {_agaga (210),_agaga (297)},11:_dg .PageSize {_agaga (148),_agaga (210)},70:_dg .PageSize {_agaga (105),_agaga (148)},12:_dg .PageSize {_agaga (250),_agaga (354)},13:_dg .PageSize {_agaga (182),_agaga (257)},14:_dg .PageSize {8.5*_ed .Inch ,13*_ed .Inch },20:_dg .PageSize {4.125*_ed .Inch ,9.5*_ed .Inch },27:_dg .PageSize {_agaga (110),_agaga (220)},28:_dg .PageSize {_agaga (162),_agaga (229)},34:_dg .PageSize {_agaga (250),_agaga (176)},29:_dg .PageSize {_agaga (324),_agaga (458)},30:_dg .PageSize {_agaga (229),_agaga (324)},31:_dg .PageSize {_agaga (114),_agaga (162)},37:_dg .PageSize {3.88*_ed .Inch ,7.5*_ed .Inch },43:_dg .PageSize {_agaga (100),_agaga (148)},69:_dg .PageSize {_agaga (200),_agaga (148)}};
const (FontStyle_Regular FontStyle =0;FontStyle_Bold FontStyle =1;FontStyle_Italic FontStyle =2;FontStyle_BoldItalic FontStyle =3;);func (_cgef *convertContext )getColorStringFromSmlColor (_fabd *_b .CT_Color )*string {var _cedc string ;if _fabd .RgbAttr !=nil {_cedc =*_fabd .RgbAttr ;
}else if _fabd .IndexedAttr !=nil &&*_fabd .IndexedAttr < 64{_cedc =_bfc [*_fabd .IndexedAttr ];}else if _fabd .ThemeAttr !=nil {_ddegg :=*_fabd .ThemeAttr ;_cedc =_cgef .getColorFromTheme (_ddegg );};if _cedc ==""{return nil ;};if len (_cedc )> 6{_cedc =_cedc [(len (_cedc )-6):];
};if _fabd .TintAttr !=nil {_ecgb :=*_fabd .TintAttr ;_cedc =_dc .AdjustColorByTint (_cedc ,_ecgb );};_cedc ="\u0023"+_cedc ;return &_cedc ;};func _eafd (_eega []*symbol )float64 {_bbag :=0.0;for _ ,_gdce :=range _eega {if _gdce ._aefce > _bbag {_bbag =_gdce ._aefce ;
};};return _bbag ;};func (_abca *convertContext )alignSymbolsHorizontally (_fcba *cell ,_fbb _b .ST_HorizontalAlignment ){if _fbb ==_b .ST_HorizontalAlignmentUnset {switch _fcba ._gde {case _b .ST_CellTypeB :_fbb =_b .ST_HorizontalAlignmentCenter ;case _b .ST_CellTypeN :_fbb =_b .ST_HorizontalAlignmentRight ;
default:_fbb =_b .ST_HorizontalAlignmentLeft ;};};var _beee float64 ;for _ ,_agf :=range _fcba ._ebg {switch _fbb {case _b .ST_HorizontalAlignmentLeft :_beee =_cd ;case _b .ST_HorizontalAlignmentRight :_fefg :=_cdffa (_agf ._ffgf );_beee =_fcba ._caba -_cd -_fefg ;
case _b .ST_HorizontalAlignmentCenter :_bcd :=_cdffa (_agf ._ffgf );_beee =(_fcba ._caba -_bcd )/2;};for _ ,_daeg :=range _agf ._ffgf {_daeg ._bfcd +=_beee ;};};};func (_eee *convertContext )makeRows (){_eeef :=[]*rowInfo {};_fae :=_eee ._agc .Rows ();
_eff :=0;_cec :=0.0;for _bef ,_aca :=range _fae {if _bef < _eee ._bdac ||(_bef > _eee ._caeg &&_eee ._caeg > 0){continue ;};_eff ++;_dga :=int (_aca .RowNumber ());if _dga > _eff {for _gbf :=_eff ;_gbf < _dga ;_gbf ++{_eeef =append (_eeef ,&rowInfo {_bbbb :_ca /_a });
_cec +=_ca /_a ;};_eff =_dga ;};var _fdb float64 ;if _aca .X ().HtAttr ==nil {_fdb =_ca ;}else {_fdb =*_aca .X ().HtAttr ;};_eeef =append (_eeef ,&rowInfo {_bbbb :_fdb /_a ,_beac :true ,_ccea :_eee .getStyle (_aca .X ().SAttr )});_cec +=_fdb /_a ;};for _eeeg :=len (_eeef );
_eeeg < _eee ._ecbde ;_eeeg ++{_eeef =append (_eeef ,&rowInfo {_bbbb :_ca /_a });_cec +=_ca /_a ;};if _eee ._egbg ||_cec > _eee ._gcaaa {for _ ,_gda :=range _eeef {_gda ._bbbb *=_eee ._cbgd ;};};_eee ._fgf =_eeef ;};const _ca =15.0;type anchor struct{_eebbe _da .Image ;
_eea *_bb .ChartSpace ;_acbf int ;_bdff int64 ;_edc int ;_ggc int64 ;_bde int ;_eefa int64 ;_edfa int ;_ffab int64 ;};func (_ege *convertContext )getStyle (_fad *uint32 )*style {_abe :=&style {};_bbgb :=false ;if _fad !=nil {_ddc :=_ege ._bdf .GetCellStyle (*_fad );
_ccgc :=_ddc .GetFont ();for _ ,_dcae :=range _ccgc .Name {if _dcae !=nil {_abe ._gaa =&_dcae .ValAttr ;_bbgb =true ;break ;};};for _ ,_beafg :=range _ccgc .B {if _beafg !=nil {_bfg :=_beafg .ValAttr ==nil ||*_beafg .ValAttr ;_abe ._dcdc =&_bfg ;_bbgb =true ;
break ;};};for _ ,_afbd :=range _ccgc .I {if _afbd !=nil {_adcd :=_afbd .ValAttr ==nil ||*_afbd .ValAttr ;_abe ._gfaec =&_adcd ;_bbgb =true ;break ;};};for _ ,_cga :=range _ccgc .U {if _cga !=nil {_bdba :=_cga .ValAttr ==_b .ST_UnderlineValuesSingle ||_cga .ValAttr ==_b .ST_UnderlineValuesUnset ;
_abe ._baee =&_bdba ;_bbgb =true ;break ;};};for _ ,_fgce :=range _ccgc .Sz {if _fgce !=nil {_fbge :=_fgce .ValAttr ;_abe ._gef =&_fbge ;_bbgb =true ;break ;};};for _ ,_cbd :=range _ccgc .VertAlign {if _cbd !=nil {_egbf :=_cbd .ValAttr ==_f .ST_VerticalAlignRunSuperscript ;
_abe ._bgbde =&_egbf ;_aeff :=_cbd .ValAttr ==_f .ST_VerticalAlignRunSubscript ;_abe ._bcfd =&_aeff ;_bbgb =true ;break ;};};for _ ,_bfdfa :=range _ccgc .Color {if _bfdfa !=nil {_abe ._gddc =_ege .getColorStringFromSmlColor (_bfdfa );_bbgb =true ;break ;
};};_bbfea :=_ddc .GetBorder ();if _bbfea .Top !=nil {_abe ._dadd =_ege .getBorder (_bbfea .Top );_bbgb =true ;};if _bbfea .Bottom !=nil {_abe ._ggbf =_ege .getBorder (_bbfea .Bottom );_bbgb =true ;};if _bbfea .Left !=nil {_abe ._efbdf =_ege .getBorder (_bbfea .Left );
_bbgb =true ;};if _bbfea .Right !=nil {_abe ._bad =_ege .getBorder (_bbfea .Right );_bbgb =true ;};if _ddc .Wrapped (){_abe ._eedb =true ;_bbgb =true ;};if _fcde :=_ddc .GetVerticalAlignment ();_fcde !=_b .ST_VerticalAlignmentUnset {_abe ._bgdc =_fcde ;
_bbgb =true ;};if _aaa :=_ddc .GetHorizontalAlignment ();_aaa !=_b .ST_HorizontalAlignmentUnset {_abe ._afc =_aaa ;_bbgb =true ;};};if _bbgb {return _abe ;};return nil ;};func (_ecb *convertContext )makeAnchors (){_fb ,_ecg :=_ecb ._agc .GetDrawing ();
if _fb !=nil {for _ ,_bd :=range _fb .EG_Anchor {_gbe :=&anchor {};if _gcf :=_bd .TwoCellAnchor ;_gcf !=nil {_gbef ,_cce :=_gcf .From ,_gcf .To ;if _gbef ==nil ||_cce ==nil {return ;};_gbe ._acbf =int (_gbef .Row );_gbe ._bdff =_dc .FromSTCoordinate (_gbef .RowOff );
_gbe ._edc =int (_gbef .Col );_gbe ._ggc =_dc .FromSTCoordinate (_gbef .ColOff );_gbe ._bde =int (_cce .Row );_gbe ._eefa =_dc .FromSTCoordinate (_cce .RowOff );_gbe ._edfa =int (_cce .Col );_gbe ._ffab =_dc .FromSTCoordinate (_cce .ColOff );if _bec :=_gcf .Choice ;
_bec !=nil {if _gcb :=_bec .Pic ;_gcb !=nil {if _gca :=_gcb .BlipFill ;_gca !=nil {if _gbg :=_gca .Blip ;_gbg !=nil {if _ggd :=_gbg .EmbedAttr ;_ggd !=nil {for _ ,_gfe :=range _ecg .X ().Relationship {if _gfe .IdAttr ==*_ggd {for _ ,_agg :=range _ecb ._dfg .Images {if _agg .Target ()==_gfe .TargetAttr {_bc ,_daf :=_eg .Open (_agg .Path ());
if _daf !=nil {_dd .Log .Debug ("\u004fp\u0065\u006e\u0020\u0069m\u0061\u0067\u0065\u0020\u0066i\u006ce\u0020e\u0072\u0072\u006f\u0072\u003a\u0020\u0025s",_daf );continue ;};_bgf ,_ ,_daf :=_da .Decode (_bc );if _daf !=nil {_dd .Log .Debug ("\u0044\u0065\u0063\u006fde\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020%\u0073",_daf );
continue ;};_gbe ._eebbe =_bgf ;};};};};};};};}else if _ggf :=_bec .GraphicFrame ;_ggf !=nil {if _bgg :=_ggf .Graphic ;_bgg !=nil {if _ba :=_bgg .GraphicData ;_ba !=nil {for _ ,_fdc :=range _ba .Any {if _efca ,_dda :=_fdc .(*_bb .Chart );_dda {for _ ,_fa :=range _ecg .X ().Relationship {if _fa .IdAttr ==_efca .IdAttr {_cgf :=_ecb ._dfg .GetChartByTargetId (_fa .TargetAttr );
if _cgf !=nil {_gbe ._eea =_cgf ;};};};};};};};};};};if _gbe ._eebbe !=nil ||_gbe ._eea !=nil {_ecb ._bggd =append (_ecb ._bggd ,_gbe );};};};};

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles.
func RegisterFontsFromDirectory (dirName string )error {return _dc .RegisterFontsFromDirectory (dirName )};var _a =3.025/_agaga (1);func (_ccd *convertContext )imageFromAnchor (_egga *anchor ,_ded ,_ebf float64 )_da .Image {if _egga ._eebbe !=nil {return _egga ._eebbe ;
};if _egga ._eea !=nil {_bdbc ,_fgcf :=_dc .MakeImageFromChartSpace (_egga ._eea ,_ded ,_ebf ,_ccd ._eggae ,_ccd ._dfg );if _fgcf !=nil {_dd .Log .Debug ("C\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067e\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u0068\u0061\u0072tS\u0070\u0061\u0063e\u003a \u0025\u0073",_fgcf );
return nil ;};return _bdbc ;};return nil ;};func (_gfeb *convertContext )determineMaxIndexes (){var _edgf ,_cdf int ;_edgf =int (_gfeb ._agc .MaxColumnIdx ());_fgb :=_gfeb ._agc .Rows ();if len (_fgb )> 0{_cdf =int (_fgb [len (_fgb )-1].RowNumber ());};
for _ ,_edfg :=range _gfeb ._bggd {if _edfg ._bde >=_cdf {_cdf =_edfg ._bde +1;};if _edfg ._edfa >=_edgf {_edgf =_edfg ._edfa +1;};};_gfeb ._ecbde =_cdf ;_gfeb ._bgbd =_edgf ;};func (_dde *convertContext )makeCols (){_egg :=_dde ._agc ;_eeb :=_egg .X ();
_ddg :=[]*colInfo {};_ad :=[]colWidthRange {};if _dgb :=_eeb .Cols ;len (_dgb )> 0{for _ ,_fgc :=range _dgb [0].Col {_egc :=_cb ;if _fge :=_fgc .WidthAttr ;_fge !=nil {if *_fge > 0.83{*_fge -=0.83;};if *_fge <=1{_egc =*_fge *11;}else {_egc =5+*_fge *6;
};};_bge :=int (_fgc .MinAttr -1);_bbb :=int (_fgc .MaxAttr -1);_ad =append (_ad ,colWidthRange {_abgb :_bge ,_cbbf :_bbb ,_gdg :_egc ,_aegb :_dde .getStyle (_fgc .StyleAttr )});};};_efd :=0;for _bee :=0;_bee <=_dde ._bgbd ;_bee ++{if _bee < _dde ._fffd ||(_bee > _dde ._fbgb &&_dde ._fbgb > 0){continue ;
};var _dcea float64 ;var _acb *style ;if _efd >=len (_ad ){_dcea =_cb ;}else {_fbf :=_ad [_efd ];if _bee >=_fbf ._abgb &&_bee <=_fbf ._cbbf {_dcea =_fbf ._gdg ;_acb =_fbf ._aegb ;if _bee ==_fbf ._cbbf {_efd ++;};}else {_dcea =_cb ;};};_ddg =append (_ddg ,&colInfo {_fcf :_dcea ,_beaf :_acb });
};_cee :=0.0;for _ ,_db :=range _ddg {_cee +=_db ._fcf ;};_dde ._cbgd =1.0;if _cee > _dde ._fgee {_dde ._cbgd =_dde ._fgee /_cee ;for _ ,_eef :=range _ddg {_eef ._fcf *=_dde ._cbgd ;};};_dde ._gdcf =_ddg ;};func (_fdcd *convertContext )getStyleFromRPrElt (_eebeb *_b .CT_RPrElt )*style {if _eebeb ==nil {return nil ;
};_caa :=&style {};_caa ._gaa =&_eebeb .RFont .ValAttr ;if _cggd :=_eebeb .B ;_cggd !=nil {_bgce :=_cggd .ValAttr ==nil ||*_cggd .ValAttr ;_caa ._dcdc =&_bgce ;};if _bdcae :=_eebeb .I ;_bdcae !=nil {_bcaa :=_bdcae .ValAttr ==nil ||*_bdcae .ValAttr ;_caa ._gfaec =&_bcaa ;
};if _edcc :=_eebeb .U ;_edcc !=nil {_dfbe :=_edcc .ValAttr ==_b .ST_UnderlineValuesSingle ||_edcc .ValAttr ==_b .ST_UnderlineValuesUnset ;_caa ._baee =&_dfbe ;};if _eafb :=_eebeb .VertAlign ;_eafb !=nil {_ffgd :=_eafb .ValAttr ==_f .ST_VerticalAlignRunSuperscript ;
_caa ._bgbde =&_ffgd ;_cda :=_eafb .ValAttr ==_f .ST_VerticalAlignRunSubscript ;_caa ._bcfd =&_cda ;};if _gfad :=_eebeb .Sz ;_gfad !=nil {_dff :=_gfad .ValAttr ;_caa ._gef =&_dff ;};if _gddbf :=_eebeb .Color ;_gddbf !=nil {_caa ._gddc =_fdcd .getColorStringFromSmlColor (_gddbf );
};return _caa ;};func _aede (_efab ,_agb *style ){if _agb ==nil {return ;};if _efab ==nil {_efab =_agb ;return ;};if _efab ._gaa ==nil {_efab ._gaa =_agb ._gaa ;};if _efab ._gddc ==nil {_efab ._gddc =_agb ._gddc ;};if _efab ._gef ==nil {_efab ._gef =_agb ._gef ;
};if _efab ._dcdc ==nil {_efab ._dcdc =_agb ._dcdc ;};if _efab ._gfaec ==nil {_efab ._gfaec =_agb ._gfaec ;};if _efab ._baee ==nil {_efab ._baee =_agb ._baee ;};if _efab ._bgbde ==nil {_efab ._bgbde =_agb ._bgbde ;};if _efab ._bcfd ==nil {_efab ._bcfd =_agb ._bcfd ;
};if _efab ._dadd ==nil {_efab ._dadd =_agb ._dadd ;};if _efab ._ggbf ==nil {_efab ._ggbf =_agb ._ggbf ;};if _efab ._efbdf ==nil {_efab ._efbdf =_agb ._efbdf ;};if _efab ._bad ==nil {_efab ._bad =_agb ._bad ;};if _efab ._bgdc ==_b .ST_VerticalAlignmentUnset {_efab ._bgdc =_agb ._bgdc ;
};if _efab ._afc ==_b .ST_HorizontalAlignmentUnset {_efab ._afc =_agb ._afc ;};};

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont (name string ,style FontStyle ,font *_gg .PdfFont ){_dc .RegisterFont (name ,style ,font );};func _agaga (_egcc float64 )float64 {return _egcc *_ed .Millimeter };type rowInfo struct{_gfde float64 ;_beac bool ;_bbbb float64 ;_ccea *style ;
_feeb []*cell ;_fdce float64 ;};func (_edb *convertContext )distributeAnchors (){for _ ,_ebc :=range _edb ._bggd {_cgc ,_aagd :=_ebc ._acbf ,_ebc ._bdff ;_fef ,_cafb :=_ebc ._edc ,_ebc ._ggc ;_eec ,_abc :=_ebc ._bde ,_ebc ._eefa ;_gce ,_adc :=_ebc ._edfa ,_ebc ._ffab ;
if _cgc < _edb ._bdac ||(_eec > _edb ._caeg &&_edb ._caeg > 0){continue ;};if _fef < _edb ._fffd ||(_gce > _edb ._fbgb &&_edb ._fbgb > 0){continue ;};var _gfac ,_gag ,_afb ,_fab *page ;for _ ,_afa :=range _edb ._acgg {for _ ,_edgd :=range _afa ._gdcc {if _cgc >=_edgd ._fce ._egdd &&_cgc < _edgd ._fce ._dee {if _fef >=_edgd ._cfgfa ._edbb &&_fef < _edgd ._cfgfa ._bcgfg {_edgd ._dbg =true ;
_gfac =_edgd ;};if _gce >=_edgd ._cfgfa ._edbb &&_gce < _edgd ._cfgfa ._bcgfg {_edgd ._dbg =true ;_gag =_edgd ;};};if _eec >=_edgd ._fce ._egdd &&_eec < _edgd ._fce ._dee {if _fef >=_edgd ._cfgfa ._edbb &&_fef < _edgd ._cfgfa ._bcgfg {_edgd ._dbg =true ;
_fab =_edgd ;};if _gce >=_edgd ._cfgfa ._edbb &&_gce < _edgd ._cfgfa ._bcgfg {_edgd ._dbg =true ;_afb =_edgd ;};};};};_gegf :=_gfac !=_gag ;_cafbf :=_gfac !=_fab ;if _gegf &&_cafbf {_bgca :=_edb ._gdcf [_fef ]._cgge +_ed .FromEMU (_cafb );_bgcd :=_gfac ._cfgfa ._bbfe ;
_bdc :=_edb ._gdcf [_gce ]._cgge +_ed .FromEMU (_adc );_ddee :=_edb ._fgf [_cgc ]._gfde +_ed .FromEMU (_aagd );_gdf :=float64 (_gfac ._fce ._feb );_dcb :=_edb ._fgf [_eec ]._gfde +_ed .FromEMU (_abc );_ggdf :=_bdc +_bgcd -_bgca ;_gdd :=_dcb +_gdf -_ddee ;
_cef :=_edb .imageFromAnchor (_ebc ,_ggdf ,_gdd );_gfac ._cge =append (_gfac ._cge ,_edb .getImage (_cef ,_gdd ,_ggdf ,_bgca ,_ddee ,_bgcd -_bgca ,_gdf -_ddee ,_dc .ImgPart_lt ));_gag ._cge =append (_gag ._cge ,_edb .getImage (_cef ,_gdd ,_ggdf ,0,_ddee ,_bgcd -_bgca ,_gdf -_ddee ,_dc .ImgPart_rt ));
_fab ._cge =append (_fab ._cge ,_edb .getImage (_cef ,_gdd ,_ggdf ,_bgca ,0,_bgcd -_bgca ,_gdf -_ddee ,_dc .ImgPart_lb ));_afb ._cge =append (_afb ._cge ,_edb .getImage (_cef ,_gdd ,_ggdf ,0,0,_bgcd -_bgca ,_gdf -_ddee ,_dc .ImgPart_rb ));}else if _gegf {_bbd :=_edb ._fgf [_cgc ]._gfde +_ed .FromEMU (_aagd );
_ffa :=_edb ._fgf [_eec ]._gfde +_ed .FromEMU (_abc );_gga :=_edb ._gdcf [_fef ]._cgge +_ed .FromEMU (_cafb );_age :=_gfac ._cfgfa ._bbfe ;_effc :=_edb ._gdcf [_gce ]._cgge +_ed .FromEMU (_adc );_bgab :=_effc +_age -_gga ;_dcef :=_ffa -_bbd ;_eaf :=_edb .imageFromAnchor (_ebc ,_bgab ,_dcef );
_gfac ._cge =append (_gfac ._cge ,_edb .getImage (_eaf ,_dcef ,_bgab ,_gga ,_bbd ,_age -_gga ,0,_dc .ImgPart_l ));_gag ._cge =append (_gag ._cge ,_edb .getImage (_eaf ,_dcef ,_bgab ,0,_bbd ,_age -_gga ,0,_dc .ImgPart_r ));}else if _cafbf {_daa :=_edb ._gdcf [_fef ]._cgge +_ed .FromEMU (_cafb );
_fefb :=_edb ._gdcf [_gce ]._cgge +_ed .FromEMU (_adc );_egd :=_edb ._fgf [_cgc ]._gfde +_ed .FromEMU (_aagd );_ecfa :=float64 (_gfac ._fce ._feb );_ebac :=_edb ._fgf [_eec ]._gfde +_ed .FromEMU (_abc );_adf :=_fefb -_daa ;_aad :=_ebac +_ecfa -_egd ;_fag :=_edb .imageFromAnchor (_ebc ,_adf ,_aad );
_gfac ._cge =append (_gfac ._cge ,_edb .getImage (_fag ,_aad ,_adf ,_daa ,_egd ,0,_ecfa -_egd ,_dc .ImgPart_t ));_fab ._cge =append (_fab ._cge ,_edb .getImage (_fag ,_aad ,_adf ,_daa ,0,0,_ecfa -_egd ,_dc .ImgPart_b ));}else {_cefa :=_edb ._gdcf [_fef ]._cgge +_ed .FromEMU (_cafb );
_gbeg :=_edb ._gdcf [_gce ]._cgge +_ed .FromEMU (_adc );_gbgc :=_edb ._fgf [_cgc ]._gfde +_ed .FromEMU (_aagd );_cgcc :=_edb ._fgf [_eec ]._gfde +_ed .FromEMU (_abc );_bgd :=_gbeg -_cefa ;_bae :=_cgcc -_gbgc ;_ebad :=_edb .imageFromAnchor (_ebc ,_bgd ,_bae );
_gfac ._cge =append (_gfac ._cge ,_edb .getImage (_ebad ,_bae ,_bgd ,_cefa ,_gbgc ,0,0,_dc .ImgPart_whole ));};};};func (_fde *convertContext )getColorFromTheme (_ged uint32 )string {_dgfa :=_fde ._dfg .Themes ();if len (_dgfa )!=0{_gcbe :=_dgfa [0];if _fdbc :=_gcbe .ThemeElements ;
_fdbc !=nil {if _becg :=_fdbc .ClrScheme ;_becg !=nil {switch _ged {case 0:return _dc .GetColorStringFromDmlColor (_becg .Lt1 );case 1:return _dc .GetColorStringFromDmlColor (_becg .Dk1 );case 2:return _dc .GetColorStringFromDmlColor (_becg .Lt2 );case 3:return _dc .GetColorStringFromDmlColor (_becg .Dk2 );
case 4:return _dc .GetColorStringFromDmlColor (_becg .Accent1 );case 5:return _dc .GetColorStringFromDmlColor (_becg .Accent2 );case 6:return _dc .GetColorStringFromDmlColor (_becg .Accent3 );case 7:return _dc .GetColorStringFromDmlColor (_becg .Accent4 );
case 8:return _dc .GetColorStringFromDmlColor (_becg .Accent5 );case 9:return _dc .GetColorStringFromDmlColor (_becg .Accent6 );};};};};return "";};var _cbf =_agaga (1);const _cb =64.0;const _be =0.0;type convertContext struct{_fcc *_dg .Creator ;_dfg *_gf .Workbook ;
_eggae *_cf .Theme ;_agc *_gf .Sheet ;_bdf *_gf .StyleSheet ;_ecbde int ;_bgbd int ;_acgg []*pagespan ;_bdga *page ;_gdcf []*colInfo ;_fgf []*rowInfo ;_afad []*rowspan ;_bbf float64 ;_ffbd float64 ;_gcaaa float64 ;_fgee float64 ;_dcc []*mergedCell ;_bggd []*anchor ;
_cbgd float64 ;_fffd int ;_fbgb int ;_bdac int ;_caeg int ;_egbg bool ;};func (_cbab *convertContext )getImage (_agfg _da .Image ,_egbc ,_cdfc ,_agag ,_fcbd ,_cabd ,_fbcb float64 ,_fece _dc .ImgPart )*_dg .Image {_fcbd +=_cbab ._bbf ;_agag +=_cbab ._ffbd ;
_cceg ,_bcf :=_dc .GetImage (_cbab ._fcc ,_agfg ,_egbc ,_cdfc ,_agag ,_fcbd ,_cabd ,_fbcb ,_fece );if _bcf !=nil {_dd .Log .Debug ("\u0043\u0061\u006eno\u0074\u0020\u0067\u0065\u0074\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_bcf );
return nil ;};return _cceg ;};