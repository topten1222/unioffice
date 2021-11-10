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

package convert ;import (_fg "github.com/unidoc/unioffice/common/logger";_dg "github.com/unidoc/unioffice/common/tempstorage";_bc "github.com/unidoc/unioffice/internal/convertutils";_be "github.com/unidoc/unioffice/measurement";_fdg "github.com/unidoc/unioffice/schema/soo/dml";_ec "github.com/unidoc/unioffice/schema/soo/dml/chart";_dd "github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes";_fa "github.com/unidoc/unioffice/schema/soo/sml";_b "github.com/unidoc/unioffice/spreadsheet";_c "github.com/unidoc/unioffice/spreadsheet/reference";_d "github.com/unidoc/unipdf/v3/creator";_fd "github.com/unidoc/unipdf/v3/model";_g "image";_f "strconv";);func (_eba *convertContext )makeAnchors (){_fb ,_acc :=_eba ._afgb .GetDrawing ();if _fb !=nil {for _ ,_ggf :=range _fb .EG_Anchor {_fbb :=&anchor {};if _da :=_ggf .TwoCellAnchor ;_da !=nil {_ccf ,_gfa :=_da .From ,_da .To ;if _ccf ==nil ||_gfa ==nil {return ;};_fbb ._bdbg =int (_ccf .Row );_fbb ._dcd =_bc .FromSTCoordinate (_ccf .RowOff );_fbb ._dfd =int (_ccf .Col );_fbb ._ddbg =_bc .FromSTCoordinate (_ccf .ColOff );_fbb ._agc =int (_gfa .Row );_fbb ._dbbe =_bc .FromSTCoordinate (_gfa .RowOff );_fbb ._fdce =int (_gfa .Col );_fbb ._bcgc =_bc .FromSTCoordinate (_gfa .ColOff );if _dff :=_da .Choice ;_dff !=nil {if _ddc :=_dff .Pic ;_ddc !=nil {if _eeg :=_ddc .BlipFill ;_eeg !=nil {if _db :=_eeg .Blip ;_db !=nil {if _gb :=_db .EmbedAttr ;_gb !=nil {for _ ,_daf :=range _acc .X ().Relationship {if _daf .IdAttr ==*_gb {for _ ,_ggb :=range _eba ._aadd .Images {if _ggb .Target ()==_daf .TargetAttr {_cd ,_ebc :=_dg .Open (_ggb .Path ());if _ebc !=nil {_fg .Log .Debug ("\u004fp\u0065\u006e\u0020\u0069m\u0061\u0067\u0065\u0020\u0066i\u006ce\u0020e\u0072\u0072\u006f\u0072\u003a\u0020\u0025s",_ebc );continue ;};_cg ,_ ,_ebc :=_g .Decode (_cd );if _ebc !=nil {_fg .Log .Debug ("\u0044\u0065\u0063\u006fde\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020%\u0073",_ebc );continue ;};_fbb ._bfee =_cg ;};};};};};};};}else if _fc :=_dff .GraphicFrame ;_fc !=nil {if _fgb :=_fc .Graphic ;_fgb !=nil {if _bfc :=_fgb .GraphicData ;_bfc !=nil {for _ ,_ba :=range _bfc .Any {if _bd ,_fae :=_ba .(*_ec .Chart );_fae {for _ ,_ff :=range _acc .X ().Relationship {if _ff .IdAttr ==_bd .IdAttr {_eac :=_eba ._aadd .GetChartByTargetId (_ff .TargetAttr );if _eac !=nil {_fbb ._cdb =_eac ;};};};};};};};};};};if _fbb ._bfee !=nil ||_fbb ._cdb !=nil {_eba ._acga =append (_eba ._acga ,_fbb );};};};};type rowspan struct{_fed float64 ;_dbgc int ;_ccfg int ;};type convertContext struct{_ebag *_d .Creator ;_aadd *_b .Workbook ;_cdca *_fdg .Theme ;_afgb *_b .Sheet ;_ecf *_b .StyleSheet ;_fcb int ;_gdad int ;_effd []*pagespan ;_dag *page ;_gde []*colInfo ;_dcg []*rowInfo ;_beeee []*rowspan ;_dgga float64 ;_ebab float64 ;_dade float64 ;_abc float64 ;_efa []*mergedCell ;_acga []*anchor ;};type border struct{_egg float64 ;_cdee _d .Color ;};var _feac =map[uint32 ]_d .PageSize {1:_d .PageSize {8.5*_be .Inch ,11*_be .Inch },2:_d .PageSize {8.5*_be .Inch ,11*_be .Inch },3:_d .PageSize {11*_be .Inch ,17*_be .Inch },4:_d .PageSize {17*_be .Inch ,11*_be .Inch },5:_d .PageSize {8.5*_be .Inch ,14*_be .Inch },6:_d .PageSize {5.5*_be .Inch ,8.5*_be .Inch },7:_d .PageSize {7.5*_be .Inch ,10*_be .Inch },8:_d .PageSize {_daabg (297),_daabg (420)},9:_d .PageSize {_daabg (210),_daabg (297)},10:_d .PageSize {_daabg (210),_daabg (297)},11:_d .PageSize {_daabg (148),_daabg (210)},70:_d .PageSize {_daabg (105),_daabg (148)},12:_d .PageSize {_daabg (250),_daabg (354)},13:_d .PageSize {_daabg (182),_daabg (257)},14:_d .PageSize {8.5*_be .Inch ,13*_be .Inch },20:_d .PageSize {4.125*_be .Inch ,9.5*_be .Inch },27:_d .PageSize {_daabg (110),_daabg (220)},28:_d .PageSize {_daabg (162),_daabg (229)},34:_d .PageSize {_daabg (250),_daabg (176)},29:_d .PageSize {_daabg (324),_daabg (458)},30:_d .PageSize {_daabg (229),_daabg (324)},31:_d .PageSize {_daabg (114),_daabg (162)},37:_d .PageSize {3.88*_be .Inch ,7.5*_be .Inch },43:_d .PageSize {_daabg (100),_daabg (148)},69:_d .PageSize {_daabg (200),_daabg (148)}};

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont (name string ,style FontStyle ,font *_fd .PdfFont ){_bc .RegisterFont (name ,style ,font );};func (_egb *convertContext )getSymbolsFromString (_ffa string ,_ccd *style )[]*symbol {_ega :=[]*symbol {};_ded :=_egb .makeTextStyleFromCellStyle (_ccd );for _ ,_ecag :=range _ffa {_ega =append (_ega ,&symbol {_ebde :string (_ecag ),_bafa :_ded });};return _ega ;};var _bfb =[]string {"\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0030\u0030\u0030\u0030\u0030\u0030","\u0066\u0066\u0066\u0066\u0066\u0066","\u0066\u0066\u0030\u0030\u0030\u0030","\u0030\u0030\u0066\u0066\u0030\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0030\u0030","\u0030\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0038\u0030\u0030\u0030","\u0038\u0030\u0030\u0030\u0038\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0063\u0030\u0063\u0030\u0063\u0030","\u0038\u0030\u0038\u0030\u0038\u0030","\u0039\u0039\u0039\u0039\u0066\u0066","\u0039\u0039\u0033\u0033\u0036\u0036","\u0066\u0066\u0066\u0066\u0063\u0063","\u0063\u0063\u0066\u0066\u0066\u0066","\u0036\u0036\u0030\u0030\u0036\u0036","\u0066\u0066\u0038\u0030\u0038\u0030","\u0030\u0030\u0036\u0036\u0063\u0063","\u0063\u0063\u0063\u0063\u0066\u0066","\u0030\u0030\u0030\u0030\u0038\u0030","\u0066\u0066\u0030\u0030\u0066\u0066","\u0066\u0066\u0066\u0066\u0030\u0030","\u0030\u0030\u0066\u0066\u0066\u0066","\u0038\u0030\u0030\u0030\u0038\u0030","\u0038\u0030\u0030\u0030\u0030\u0030","\u0030\u0030\u0038\u0030\u0038\u0030","\u0030\u0030\u0030\u0030\u0066\u0066","\u0030\u0030\u0063\u0063\u0066\u0066","\u0063\u0063\u0066\u0066\u0066\u0066","\u0063\u0063\u0066\u0066\u0063\u0063","\u0066\u0066\u0066\u0066\u0039\u0039","\u0039\u0039\u0063\u0063\u0066\u0066","\u0066\u0066\u0039\u0039\u0063\u0063","\u0063\u0063\u0039\u0039\u0066\u0066","\u0066\u0066\u0063\u0063\u0039\u0039","\u0033\u0033\u0036\u0036\u0066\u0066","\u0033\u0033\u0063\u0063\u0063\u0063","\u0039\u0039\u0063\u0063\u0030\u0030","\u0066\u0066\u0063\u0063\u0030\u0030","\u0066\u0066\u0039\u0039\u0030\u0030","\u0066\u0066\u0036\u0036\u0030\u0030","\u0036\u0036\u0036\u0036\u0039\u0039","\u0039\u0036\u0039\u0036\u0039\u0036","\u0030\u0030\u0033\u0033\u0036\u0036","\u0033\u0033\u0039\u0039\u0036\u0036","\u0030\u0030\u0033\u0033\u0030\u0030","\u0033\u0033\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0030\u0030","\u0039\u0039\u0033\u0033\u0036\u0036","\u0033\u0033\u0033\u0033\u0039\u0039","\u0033\u0033\u0033\u0033\u0033\u0033"};type cell struct{_fffa _fa .ST_CellType ;_ebbg int ;_fafd float64 ;_fcbe []*line ;_addce float64 ;_fecdf float64 ;_gadb float64 ;_fbd float64 ;_cbaa float64 ;_bfca *_d .TextStyle ;_adf *border ;_aead *border ;_agg *border ;_cbc *border ;_ebd bool ;_dcc bool ;};const (FontStyle_Regular FontStyle =0;FontStyle_Bold FontStyle =1;FontStyle_Italic FontStyle =2;FontStyle_BoldItalic FontStyle =3;);

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle =_bc .FontStyle ;type mergedCell struct{_ada uint32 ;_aggb uint32 ;_fada uint32 ;_dae uint32 ;_fdcc float64 ;_fbea float64 ;};const _gg =0.25;type colInfo struct{_agd float64 ;_fbad float64 ;_bdgc *style ;};func (_beee *convertContext )getContentFromCell (_cgea _b .Cell ,_aba *style ,_fff float64 ,_dac bool )([]*line ,_fa .ST_CellType ){_fdec :=_cgea .X ();var _aecc []*symbol ;switch _fdec .TAttr {case _fa .ST_CellTypeS :_gcb :=_fdec .V ;if _gcb !=nil {_dgf ,_bca :=_f .Atoi (*_gcb );if _bca ==nil {_aaf :=_beee ._aadd .SharedStrings .X ().Si [_dgf ];if _aaf .T !=nil {_aecc =_beee .getSymbolsFromString (*_aaf .T ,_aba );}else if _aaf .R !=nil {_aecc =_beee .getSymbolsFromR (_aaf .R ,_aba );};};};case _fa .ST_CellTypeB :_aeg :=_fdec .V ;if _aeg !=nil {if *_aeg =="\u0030"{_aecc =_beee .getSymbolsFromString ("\u0046\u0041\u004cS\u0045",_aba );}else {_aecc =_beee .getSymbolsFromString ("\u0054\u0052\u0055\u0045",_aba );};};default:_aecc =_beee .getSymbolsFromString (_cgea .GetFormattedValue (),_aba );};_def :=0.0;_egdf :=0.0;var _afd []*line ;var _beag bool ;if _aba !=nil {if _aba ._bggg !=nil {if *_aba ._bggg {_beag =true ;};};if _aba ._bgd !=nil {if *_aba ._bgd {_beag =true ;};};};if _dac {_afd =[]*line {};_ecedc :=_fff -2*_a ;_bcdb :=[]*symbol {};for _ ,_dge :=range _aecc {_bce (_dge );if _def +_dge ._edce >=_ecedc {_aae :=_deg (_bcdb );if _beag {_aae /=_gc ;};_afd =append (_afd ,&line {_cfbc :_egdf ,_eeb :_bcdb ,_eaec :_aae });_bcdb =[]*symbol {_dge };_def =_dge ._edce ;_egdf +=_aae ;}else {_dge ._fgfa =_def ;_def +=_dge ._edce ;_bcdb =append (_bcdb ,_dge );};};_eab :=_deg (_bcdb );if _beag {_eab /=_gc ;};if len (_bcdb )> 0{_afd =append (_afd ,&line {_cfbc :_egdf ,_eeb :_bcdb ,_eaec :_eab });};}else {for _ ,_ccc :=range _aecc {_bce (_ccc );_ccc ._fgfa =_def ;_def +=_ccc ._edce ;};if len (_aecc )> 0{_afd =[]*line {&line {_eeb :_aecc ,_eaec :_deg (_aecc )}};};};_bbfc :=_fdec .TAttr ;if _bbfc ==_fa .ST_CellTypeUnset {_bbfc =_fa .ST_CellTypeN ;};return _afd ,_bbfc ;};const _gf =1.5;func _deg (_aebc []*symbol )float64 {_addc :=0.0;for _ ,_gdbg :=range _aebc {if _gdbg ._aac > _addc {_addc =_gdbg ._aac ;};};return _addc ;};var _bg =_daabg (1);func (_bfe *convertContext )drawPage (_bdcb *page ){_bba :=_bfe ._dgga ;_aaga :=_bfe ._ebab ;for _ ,_eae :=range _bdcb ._gfad {_feaf :=_bfe ._dcg [_eae ._aega ];for _ ,_aeb :=range _eae ._dbag {_cdc :=_aeb ._fbd < _aeb ._fafd ;_cbgf :=_aeb ._cbaa > _aeb ._fafd +_aeb ._addce ;var _efg ,_bfg bool ;for _ ,_cca :=range _aeb ._fcbe {for _ ,_egd :=range _cca ._eeb {if _cdc &&!_efg {_efg =_egd ._fgfa < 0;};if _cbgf &&!_bfg {_bfg =_aeb ._addce < _egd ._fgfa +_egd ._edce ;};if _aeb ._fafd +_egd ._fgfa >=_aeb ._fbd &&_aeb ._fafd +_egd ._fgfa +_egd ._edce <=_aeb ._cbaa {_cde :=_bfe ._ebag .NewStyledParagraph ();_ged :=_aaga +_aeb ._fafd +_egd ._fgfa ;_fdca :=_bba +_feaf ._feab +_cca ._cfbc -_egd ._aac -_daabg (0.5);_cde .SetPos (_ged ,_fdca );var _ggad *_d .TextChunk ;if _egd ._dbdc !=""{_ggad =_cde .AddExternalLink (_egd ._ebde ,_egd ._dbdc );}else {_ggad =_cde .Append (_egd ._ebde );};if _egd ._bafa !=nil {_ggad .Style =*_egd ._bafa ;};_bfe ._ebag .Draw (_cde );};};};var _ecad ,_ggda ,_aab ,_beda ,_dbe ,_eaea float64 ;var _fdd ,_ebbf ,_dc ,_fcd _d .Color ;if _gafa :=_aeb ._adf ;_gafa !=nil {_ecad =_gafa ._egg ;_fdd =_gafa ._cdee ;};if _afg :=_aeb ._aead ;_afg !=nil {_ggda =_afg ._egg ;_ebbf =_afg ._cdee ;};if _bfd :=_aeb ._agg ;_bfd !=nil {_aab =_bfd ._egg ;_dbe =_aab /2;_dc =_bfd ._cdee ;};if _gffb :=_aeb ._cbc ;_gffb !=nil {_beda =_gffb ._egg ;_eaea =_beda /2;_fcd =_gffb ._cdee ;};var _gcag float64 ;if _eae ._aega > 1{_gcag =_bfe ._dcg [_eae ._aega -1]._fcbc ;};_bfgg :=_bba +_feaf ._feab -0.5*(_gcag -_ecad );_gdb :=_bba +_feaf ._feab +_feaf ._beg +0.5*(_feaf ._fcbc +_ggda );_ecef :=_aaga +_aeb ._fafd ;_abdd :=_ecef +_aeb ._fecdf ;_bc .DrawLine (_bfe ._ebag ,_ecef ,_bfgg ,_abdd ,_bfgg ,_ecad ,_fdd );_bc .DrawLine (_bfe ._ebag ,_ecef ,_gdb ,_abdd ,_gdb ,_ggda ,_ebbf );if !_efg {_bc .DrawLine (_bfe ._ebag ,_ecef -_dbe ,_bfgg ,_ecef -_dbe ,_gdb ,_aab ,_dc );};if !_bfg {_bc .DrawLine (_bfe ._ebag ,_abdd -_eaea ,_bfgg ,_abdd -_eaea ,_gdb ,_beda ,_fcd );};};};for _ ,_cggg :=range _bdcb ._acgae {if _cggg !=nil {_bfe ._ebag .Draw (_cggg );};};};func (_ddd *convertContext )addRowToPage (_eced []*cell ,_fdcf int ){_faf :=0.0;_bbea :=_ddd ._abc ;for _ ,_cgec :=range _eced {if len (_cgec ._fcbe )!=0{_cgec ._fbd =_faf ;_faf =_cgec ._fafd +_cgec ._addce ;};};for _gfdb :=len (_eced )-1;_gfdb >=0;_gfdb --{_eea :=_eced [_gfdb ];if len (_eea ._fcbe )!=0{_eea ._cbaa =_bbea ;_bbea =_eea ._fafd ;};};_ddd ._dag ._gfad =append (_ddd ._dag ._gfad ,&pageRow {_aega :_fdcf ,_dbag :_eced });};func (_edbf *convertContext )getImage (_fbef _g .Image ,_aef ,_eaeaa ,_edaf ,_bage ,_fadc ,_ceg float64 ,_ffad _bc .ImgPart )*_d .Image {_bage +=_edbf ._dgga ;_edaf +=_edbf ._ebab ;_dbed ,_ggbfe :=_bc .GetImage (_edbf ._ebag ,_fbef ,_aef ,_eaeaa ,_edaf ,_bage ,_fadc ,_ceg ,_ffad );if _ggbfe !=nil {_fg .Log .Debug ("\u0043\u0061\u006eno\u0074\u0020\u0067\u0065\u0074\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_ggbfe );return nil ;};return _dbed ;};func (_fcfg *convertContext )alignSymbolsVertically (_ccca *cell ,_fffe _fa .ST_VerticalAlignment ){var _gfac float64 ;switch _fffe {case _fa .ST_VerticalAlignmentTop :_gfac =_cc ;if _ccca ._ebd {_gfac -=_gf ;}else if _ccca ._dcc {_gfac +=4*_gf ;};for _ ,_ege :=range _ccca ._fcbe {_gfac +=_ege ._eaec ;_ege ._cfbc =_gfac ;_gfac +=_bg ;};case _fa .ST_VerticalAlignmentCenter :_ffb :=0.0;for _ ,_egba :=range _ccca ._fcbe {_ffb +=_egba ._eaec +_daabg (1);};_gfac =0.5*(_ccca ._gadb -_ffb );if _ccca ._ebd {_gfac -=2*_gf ;}else if _ccca ._dcc {_gfac +=2*_gf ;};for _ ,_dbbd :=range _ccca ._fcbe {_gfac +=_dbbd ._eaec +0.5*_bg ;_dbbd ._cfbc =_gfac ;_gfac +=0.5*_bg ;};default:_gfac =_ccca ._gadb -_cc ;if _ccca ._ebd {_gfac -=4*_gf ;}else if _ccca ._dcc {_gfac +=_gf ;};for _abf :=len (_ccca ._fcbe )-1;_abf >=0;_abf --{_ccca ._fcbe [_abf ]._cfbc =_gfac ;_gfac -=_ccca ._fcbe [_abf ]._eaec ;_gfac -=_bg ;};};};type anchor struct{_bfee _g .Image ;_cdb *_ec .ChartSpace ;_bdbg int ;_dcd int64 ;_dfd int ;_ddbg int64 ;_agc int ;_dbbe int64 ;_fdce int ;_bcgc int64 ;};var _ecc =_daabg (0.0625);func (_bcf *convertContext )alignSymbolsHorizontally (_gba *cell ,_edd _fa .ST_HorizontalAlignment ){if _edd ==_fa .ST_HorizontalAlignmentUnset {switch _gba ._fffa {case _fa .ST_CellTypeB :_edd =_fa .ST_HorizontalAlignmentCenter ;case _fa .ST_CellTypeN :_edd =_fa .ST_HorizontalAlignmentRight ;default:_edd =_fa .ST_HorizontalAlignmentLeft ;};};var _ddcc float64 ;for _ ,_degb :=range _gba ._fcbe {switch _edd {case _fa .ST_HorizontalAlignmentLeft :_ddcc =_a ;case _fa .ST_HorizontalAlignmentRight :_dfaa :=_ccfd (_degb ._eeb );_ddcc =_gba ._addce -_a -_dfaa ;case _fa .ST_HorizontalAlignmentCenter :_bdb :=_ccfd (_degb ._eeb );_ddcc =(_gba ._addce -_bdb )/2;};for _ ,_bcde :=range _degb ._eeb {_bcde ._fgfa +=_ddcc ;};};};type pageRow struct{_aega int ;_dbag []*cell ;};func (_gea *convertContext )makeMergedCells (){_aff :=[]*mergedCell {};for _ ,_abb :=range _gea ._afgb .MergedCells (){_gd ,_cab ,_add :=_c .ParseRangeReference (_abb .Reference ());if _add !=nil {_fg .Log .Debug ("\u0065\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c: \u0025\u0073",_add );continue ;};_aa :=mergedCell {_ada :_gd .RowIdx ,_aggb :_gd .ColumnIdx ,_fada :_cab .RowIdx ,_dae :_cab .ColumnIdx };for _bde :=_aa ._aggb -1;_bde < _aa ._dae ;_bde ++{_aa ._fdcc +=_gea ._gde [_bde ]._fbad ;};for _fdff :=_aa ._ada -1;_fdff < _aa ._fada ;_fdff ++{_aa ._fbea +=_gea ._dcg [_fdff ]._beg ;};_aff =append (_aff ,&_aa );};_gea ._efa =_aff ;};func _daabg (_abeb float64 )float64 {return _abeb *_be .Millimeter };func _ccfd (_cbd []*symbol )float64 {_fbe :=0.0;for _ ,_fdgf :=range _cbd {_fbe +=_fdgf ._edce ;};return _fbe ;};func (_fba *convertContext )makeCells (){_cbg :=_fba ._afgb ;_edb :=_cbg .Rows ();_fgcf :=0;for _ ,_edf :=range _fba ._dcg {_edf ._fecd =[]*cell {};_ffd :=0.0;_ggdc :=_edf ._ggc ;if _edf ._feb {_ffc :=_edb [_fgcf ];_fgcf ++;_bdf :=_edf ._beg ;for _ ,_cf :=range _ffc .Cells (){_fe ,_gcf :=_c .ParseCellReference (_cf .Reference ());if _gcf !=nil {_fg .Log .Debug ("\u0043\u0061\u006e\u006eo\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0061\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u003a \u0025\u0073",_gcf );continue ;};_gcc :=_fba ._gde [_fe .ColumnIdx ];_gca :=_gcc ._fbad ;_gfd :=_gca ;_bag :=_gcc ._bdgc ;var _fec ,_fbc ,_ced ,_gad bool ;for _ ,_faee :=range _fba ._efa {if _fe .RowIdx >=_faee ._ada &&_fe .RowIdx <=_faee ._fada &&_fe .ColumnIdx >=_faee ._aggb &&_fe .ColumnIdx <=_faee ._dae {if _fe .ColumnIdx ==_faee ._aggb &&_fe .RowIdx ==_faee ._ada {_gca =_faee ._fdcc ;_bdf =_faee ._fbea ;};_fec =_fe .RowIdx !=_faee ._ada ;_fbc =_fe .RowIdx !=_faee ._fada ;_ced =_fe .ColumnIdx !=_faee ._aggb ;_gad =_fe .ColumnIdx !=_faee ._dae ;};};_cdf :=_fba .getStyleFromCell (_cf ,_ggdc ,_bag );var _ddfb ,_cdg ,_afb bool ;var _fea ,_bed ,_cad ,_ffcb *border ;var _edcb _fa .ST_VerticalAlignment ;var _fac _fa .ST_HorizontalAlignment ;if _cdf !=nil {if !_fec {_fea =_cdf ._dea ;};if !_fbc {_bed =_cdf ._ddfa ;};if !_ced {_cad =_cdf ._bedg ;};if !_gad {_ffcb =_cdf ._bbg ;};if _bed !=nil &&_bed ._egg > _ffd {_ffd =_bed ._egg ;};_edcb =_cdf ._gdef ;_fac =_cdf ._bafd ;if _cdf ._bggg !=nil {_ddfb =*_cdf ._bggg ;};if _cdf ._bgd !=nil {_cdg =*_cdf ._bgd ;};_afb =_cdf ._gadc ;};_cff ,_cfe :=_fba .getContentFromCell (_cf ,_cdf ,_gca ,_afb );_fbac :=&cell {_fffa :_cfe ,_addce :_gca ,_fecdf :_gfd ,_gadb :_bdf ,_fcbe :_cff ,_adf :_fea ,_aead :_bed ,_agg :_cad ,_cbc :_ffcb ,_ebd :_ddfb ,_dcc :_cdg };_fba .alignSymbolsHorizontally (_fbac ,_fac );_fba .alignSymbolsVertically (_fbac ,_edcb );_edf ._fecd =append (_edf ._fecd ,_fbac );};};_edf ._fcbc =_ffd ;};};func (_ebb *convertContext )makeRowspans (){var _acea float64 ;_aad :=0;for _fead ,_aada :=range _ebb ._dcg {_eed :=_aada ._beg +_aada ._fcbc ;if _acea +_eed <=_ebb ._dade {_aada ._feab =_acea ;_acea +=_eed ;}else {_ebb ._beeee =append (_ebb ._beeee ,&rowspan {_fed :_acea ,_dbgc :_aad ,_ccfg :_fead });_aad =_fead ;_aada ._feab =0;_acea =_eed ;};};_ebb ._beeee =append (_ebb ._beeee ,&rowspan {_fed :_acea ,_dbgc :_aad ,_ccfg :len (_ebb ._dcg )});};func (_bdc *convertContext )fillPages (){for _facf ,_aea :=range _bdc ._beeee {_egc :=_bdc ._dcg [_aea ._dbgc :_aea ._ccfg ];for _bea ,_aec :=range _egc {_gce :=0;_ebf :=0.0;_cef :=[]*cell {};if _aec ._feb {for _ ,_cbe :=range _aec ._fecd {_afbc :=_bdc ._effd [_gce ];_bdc ._dag =_afbc ._dbd [_facf ];_bdc ._dag ._bada =true ;_fagc :=_cbe ._fecdf ;if _ebf +_fagc > _afbc ._gbac {_bdc .addRowToPage (_cef ,_bea );_cef =[]*cell {_cbe };_ebf =_fagc ;_gce ++;}else {_cbe ._fafd =_ebf ;_cef =append (_cef ,_cbe );_ebf +=_fagc ;};};if len (_cef )> 0{_dbg :=_bdc ._effd [_gce ];_bdc ._dag =_dbg ._dbd [_facf ];_bdc ._dag ._bada =true ;_bdc .addRowToPage (_cef ,_bea );};};};};};type pagespan struct{_gbac float64 ;_dbd []*page ;_fga int ;_afdd int ;};type rowInfo struct{_feab float64 ;_feb bool ;_beg float64 ;_ggc *style ;_fecd []*cell ;_fcbc float64 ;};var _ef =3.025/_daabg (1);func (_cgg *convertContext )makeRows (){_bad :=[]*rowInfo {};_ece :=_cgg ._afgb .Rows ();_gaf :=0;for _ ,_bac :=range _ece {_gaf ++;_ggd :=int (_bac .RowNumber ());if _ggd > _gaf {for _bbf :=_gaf ;_bbf < _ggd ;_bbf ++{_bad =append (_bad ,&rowInfo {_beg :16/_ef });};_gaf =_ggd ;};var _eec float64 ;if _bac .X ().HtAttr ==nil {_eec =16;}else {_eec =*_bac .X ().HtAttr ;};_bad =append (_bad ,&rowInfo {_beg :_eec /_ef ,_feb :true ,_ggc :_cgg .getStyle (_bac .X ().SAttr )});};for _fad :=len (_bad );_fad < _cgg ._fcb ;_fad ++{_bad =append (_bad ,&rowInfo {_beg :16/_ef });};_cgg ._dcg =_bad ;};func (_gbb *convertContext )getColorStringFromSmlColor (_agcc *_fa .CT_Color )*string {var _eaf string ;if _agcc .RgbAttr !=nil {_eaf =*_agcc .RgbAttr ;}else if _agcc .IndexedAttr !=nil &&*_agcc .IndexedAttr < 64{_eaf =_bfb [*_agcc .IndexedAttr ];}else if _agcc .ThemeAttr !=nil {_feafe :=*_agcc .ThemeAttr ;_eaf =_gbb .getColorFromTheme (_feafe );};if _eaf ==""{return nil ;};if len (_eaf )> 6{_eaf =_eaf [(len (_eaf )-6):];};if _agcc .TintAttr !=nil {_bgdg :=*_agcc .TintAttr ;_eaf =_bc .AdjustColorByTint (_eaf ,_bgdg );};_eaf ="\u0023"+_eaf ;return &_eaf ;};func (_cfc *convertContext )imageFromAnchor (_dfe *anchor ,_ggdg ,_bdg float64 )_g .Image {if _dfe ._bfee !=nil {return _dfe ._bfee ;};if _dfe ._cdb !=nil {_fda ,_ccb :=_bc .MakeImageFromChartSpace (_dfe ._cdb ,_ggdg ,_bdg ,_cfc ._cdca );if _ccb !=nil {_fg .Log .Debug ("C\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067e\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u0068\u0061\u0072tS\u0070\u0061\u0063e\u003a \u0025\u0073",_ccb );return nil ;};return _fda ;};return nil ;};func _gbc (_eaa ,_gbe *style ){if _gbe ==nil {return ;};if _eaa ==nil {_eaa =_gbe ;return ;};if _eaa ._affg ==nil {_eaa ._affg =_gbe ._affg ;};if _eaa ._geac ==nil {_eaa ._geac =_gbe ._geac ;};if _eaa ._bgb ==nil {_eaa ._bgb =_gbe ._bgb ;};if _eaa ._aagg ==nil {_eaa ._aagg =_gbe ._aagg ;};if _eaa ._ddfe ==nil {_eaa ._ddfe =_gbe ._ddfe ;};if _eaa ._cba ==nil {_eaa ._cba =_gbe ._cba ;};if _eaa ._bggg ==nil {_eaa ._bggg =_gbe ._bggg ;};if _eaa ._bgd ==nil {_eaa ._bgd =_gbe ._bgd ;};if _eaa ._dea ==nil {_eaa ._dea =_gbe ._dea ;};if _eaa ._ddfa ==nil {_eaa ._ddfa =_gbe ._ddfa ;};if _eaa ._bedg ==nil {_eaa ._bedg =_gbe ._bedg ;};if _eaa ._bbg ==nil {_eaa ._bbg =_gbe ._bbg ;};if _eaa ._gdef ==_fa .ST_VerticalAlignmentUnset {_eaa ._gdef =_gbe ._gdef ;};if _eaa ._bafd ==_fa .ST_HorizontalAlignmentUnset {_eaa ._bafd =_gbe ._bafd ;};};const _cc =2;

// ConvertToPdf converts a sheet to a PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf (s *_b .Sheet )*_d .Creator {_ab :=s .X ();if _ab ==nil {return nil ;};var _bf _d .PageSize ;var _bee bool ;if _ga :=_ab .PageSetup ;_ga !=nil {if _bcg :=_ga .PaperSizeAttr ;_bcg !=nil {_bf =_feac [*_bcg ];};_bee =_ga .OrientationAttr ==_fa .ST_OrientationLandscape ;};if (_bf ==_d .PageSize {}){_bf =_feac [1];};if _bee {_bf [0],_bf [1]=_bf [1],_bf [0];};_af :=_d .New ();_af .SetPageSize (_bf );var _fde ,_ea ,_eb ,_ee float64 ;if _ac :=_ab .PageMargins ;_ac !=nil {_eb =_ac .LeftAttr ;_ee =_ac .RightAttr ;_fde =_ac .TopAttr ;_ea =_ac .BottomAttr ;};if _eb < 0.25{_eb =0.25;};if _ee < 0.25{_ee =0.25;};_fde *=_be .Inch ;_ea *=_be .Inch ;_eb *=_be .Inch ;_ee *=_be .Inch ;_af .SetPageMargins (_eb ,_ee ,_fde ,_ea );_ge :=s .Workbook ();var _bb *_fdg .Theme ;if len (_ge .Themes ())> 0{_bb =_ge .Themes ()[0];};_df :=&convertContext {_ebag :_af ,_afgb :s ,_aadd :s .Workbook (),_cdca :_bb ,_ecf :&s .Workbook ().StyleSheet ,_dgga :_fde ,_ebab :_eb ,_dade :_bf [1]-_ea -_fde ,_abc :_bf [0]-_ee -_eb };_df .makeAnchors ();_df .determineMaxIndexes ();if _df ._fcb ==0&&_df ._gdad ==0{_af .NewPage ();return _af ;};_df .makeCols ();_df .makeRows ();_df .makeMergedCells ();_df .makeCells ();_df .makePagespans ();_df .makeRowspans ();_df .makePages ();_df .fillPages ();_df .distributeAnchors ();_df .drawSheet ();return _af ;};func (_fag *convertContext )makePagespans (){_fag ._effd =[]*pagespan {};_bgg :=0.0;_efb :=0;for _bae ,_gbg :=range _fag ._gde {_baf :=_gbg ._fbad ;if _bgg +_baf <=_fag ._abc {_gbg ._agd =_bgg ;_bgg +=_baf ;}else {_gbg ._agd =0;_fag ._effd =append (_fag ._effd ,&pagespan {_gbac :_bgg ,_fga :_efb ,_afdd :_bae });_bgg =_baf ;_efb =_bae ;};};_fag ._effd =append (_fag ._effd ,&pagespan {_gbac :_bgg ,_fga :_efb ,_afdd :len (_fag ._gde )});};type symbol struct{_ebde string ;_fgfa float64 ;_aac float64 ;_edce float64 ;_bafa *_d .TextStyle ;_dbdc string ;};func (_aaec *convertContext )getStyleFromCell (_aebd _b .Cell ,_bceg ,_acde *style )*style {_acgaea :=_aebd .X ();_gfda :=_aaec .getStyle (_acgaea .SAttr );_gbc (_gfda ,_bceg );_gbc (_gfda ,_acde );return _gfda ;};func (_bga *convertContext )makePages (){for _ ,_cae :=range _bga ._effd {for _ ,_ecd :=range _bga ._beeee {_cae ._dbd =append (_cae ._dbd ,&page {_gfad :[]*pageRow {},_fbgf :_cae ,_cbf :_ecd });};};};type style struct{_geac *string ;_bgb *float64 ;_affg *string ;_aagg *bool ;_ddfe *bool ;_cba *bool ;_bggg *bool ;_bgd *bool ;_dea *border ;_ddfa *border ;_bedg *border ;_bbg *border ;_gadc bool ;_gdef _fa .ST_VerticalAlignment ;_bafd _fa .ST_HorizontalAlignment ;};func (_dgfg *convertContext )makeTextStyleFromCellStyle (_dfda *style )*_d .TextStyle {_aafa :=_dgfg ._ebag .NewTextStyle ();if _dfda ==nil {_aafa .FontSize =_bc .DefaultFontSize ;_aafa .Font =_bc .AssignStdFontByName (_aafa ,_bc .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][FontStyle_Regular ]);return &_aafa ;};if _fecb (_dfda ._cba ){_aafa .Underline =true ;_aafa .UnderlineStyle =_d .TextDecorationLineStyle {Offset :0.5,Thickness :_daabg (1/32)};};var _cdd FontStyle ;if _fecb (_dfda ._aagg )&&_fecb (_dfda ._ddfe ){_cdd =FontStyle_BoldItalic ;}else if _fecb (_dfda ._aagg ){_cdd =FontStyle_Bold ;}else if _fecb (_dfda ._ddfe ){_cdd =FontStyle_Italic ;}else {_cdd =FontStyle_Regular ;};_dbda :="\u0064e\u0066\u0061\u0075\u006c\u0074";if _dfda ._affg !=nil {_dbda =*_dfda ._affg ;};if _dce ,_ecada :=_bc .StdFontsMap [_dbda ];_ecada {_aafa .Font =_bc .AssignStdFontByName (_aafa ,_dce [_cdd ]);}else if _fdba :=_bc .GetRegisteredFont (_dbda ,_cdd );_fdba !=nil {_aafa .Font =_fdba ;}else {_fg .Log .Debug ("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e",_dbda ,_cdd );_aafa .Font =_bc .AssignStdFontByName (_aafa ,_bc .StdFontsMap ["\u0064e\u0066\u0061\u0075\u006c\u0074"][_cdd ]);};if _dfda ._bgb !=nil {_aafa .FontSize =*_dfda ._bgb ;};if _dfda ._geac !=nil {_aafa .Color =_d .ColorRGBFromHex (*_dfda ._geac );};if _dfda ._bggg !=nil &&*_dfda ._bggg {_aafa .FontSize *=_gc ;}else if _dfda ._bgd !=nil &&*_dfda ._bgd {_aafa .FontSize *=_gc ;};return &_aafa ;};func (_ddf *convertContext )determineMaxIndexes (){var _dga ,_bab int ;_dga =int (_ddf ._afgb .MaxColumnIdx ());_dfa :=_ddf ._afgb .Rows ();if len (_dfa )> 0{_bab =int (_dfa [len (_dfa )-1].RowNumber ());};for _ ,_dba :=range _ddf ._acga {if _dba ._agc >=_bab {_bab =_dba ._agc +1;};if _dba ._fdce >=_dga {_dga =_dba ._fdce +1;};};_ddf ._fcb =_bab ;_ddf ._gdad =_dga ;};func (_dda *convertContext )combineCellStyleWithRPrElt (_feeg *style ,_gfb *_fa .CT_RPrElt )*style {_cdec :=*_feeg ;_cga :=_dda .getStyleFromRPrElt (_gfb );if _cga ==nil {return &_cdec ;};if _cga ._geac !=nil {_cdec ._geac =_cga ._geac ;};if _cga ._bgb !=nil {_cdec ._bgb =_cga ._bgb ;};if _cga ._affg !=nil {_cdec ._affg =_cga ._affg ;};if _cga ._aagg !=nil {_cdec ._aagg =_cga ._aagg ;};if _cga ._ddfe !=nil {_cdec ._ddfe =_cga ._ddfe ;};if _cga ._cba !=nil {_cdec ._cba =_cga ._cba ;};if _cga ._bggg !=nil {_cdec ._bggg =_cga ._bggg ;};if _cga ._bgd !=nil {_cdec ._bgd =_cga ._bgd ;};return &_cdec ;};func _bce (_gacc *symbol ){_fgd :=_d .New ();_fcg :=_fgd .NewStyledParagraph ();_fcg .SetMargins (0,0,0,0);_egdb :=_fcg .Append (_gacc ._ebde );if _gacc ._bafa !=nil {_egdb .Style =*_gacc ._bafa ;};_gacc ._aac =_fcg .Height ();if _gacc ._edce ==0{_gacc ._edce =_fcg .Width ();};};func (_cdcc *convertContext )getColorFromTheme (_addcb uint32 )string {_cda :=_cdcc ._aadd .Themes ();if len (_cda )!=0{_daaba :=_cda [0];if _egdc :=_daaba .ThemeElements ;_egdc !=nil {if _abddb :=_egdc .ClrScheme ;_abddb !=nil {switch _addcb {case 0:return _bc .GetColorStringFromDmlColor (_abddb .Lt1 );case 1:return _bc .GetColorStringFromDmlColor (_abddb .Dk1 );case 2:return _bc .GetColorStringFromDmlColor (_abddb .Lt2 );case 3:return _bc .GetColorStringFromDmlColor (_abddb .Dk2 );case 4:return _bc .GetColorStringFromDmlColor (_abddb .Accent1 );case 5:return _bc .GetColorStringFromDmlColor (_abddb .Accent2 );case 6:return _bc .GetColorStringFromDmlColor (_abddb .Accent3 );case 7:return _bc .GetColorStringFromDmlColor (_abddb .Accent4 );case 8:return _bc .GetColorStringFromDmlColor (_abddb .Accent5 );case 9:return _bc .GetColorStringFromDmlColor (_abddb .Accent6 );};};};};return "";};

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles.
func RegisterFontsFromDirectory (dirName string )error {return _bc .RegisterFontsFromDirectory (dirName )};type page struct{_gfad []*pageRow ;_bada bool ;_acgae []*_d .Image ;_fbgf *pagespan ;_cbf *rowspan ;};func _fecb (_cgcg *bool )bool {return _cgcg !=nil &&*_cgcg };func (_daga *convertContext )getBorder (_bfbb *_fa .CT_BorderPr )*border {_cedg :=&border {};switch _bfbb .StyleAttr {case _fa .ST_BorderStyleThin :_cedg ._egg =_ecc ;case _fa .ST_BorderStyleMedium :_cedg ._egg =_ecc *2;case _fa .ST_BorderStyleThick :_cedg ._egg =_ecc *4;};if _cedg ._egg ==0.0{return nil ;};if _aegd :=_bfbb .Color ;_aegd !=nil {_beeeg :=_daga .getColorStringFromSmlColor (_aegd );if _beeeg !=nil {_cedg ._cdee =_d .ColorRGBFromHex (*_beeeg );}else {_cedg ._cdee =_d .ColorBlack ;};};return _cedg ;};func (_ccg *convertContext )getStyleFromRPrElt (_ceff *_fa .CT_RPrElt )*style {if _ceff ==nil {return nil ;};_acgd :=&style {};_acgd ._affg =&_ceff .RFont .ValAttr ;if _bfaf :=_ceff .B ;_bfaf !=nil {_ccfb :=_bfaf .ValAttr ==nil ||*_bfaf .ValAttr ;_acgd ._aagg =&_ccfb ;};if _eag :=_ceff .I ;_eag !=nil {_dbba :=_eag .ValAttr ==nil ||*_eag .ValAttr ;_acgd ._ddfe =&_dbba ;};if _bfgb :=_ceff .U ;_bfgb !=nil {_dfc :=_bfgb .ValAttr ==_fa .ST_UnderlineValuesSingle ||_bfgb .ValAttr ==_fa .ST_UnderlineValuesUnset ;_acgd ._cba =&_dfc ;};if _abg :=_ceff .VertAlign ;_abg !=nil {_cefb :=_abg .ValAttr ==_dd .ST_VerticalAlignRunSuperscript ;_acgd ._bggg =&_cefb ;_dfcc :=_abg .ValAttr ==_dd .ST_VerticalAlignRunSubscript ;_acgd ._bgd =&_dfcc ;};if _fgdg :=_ceff .Sz ;_fgdg !=nil {_gcd :=_fgdg .ValAttr /12*_bc .DefaultFontSize ;_acgd ._bgb =&_gcd ;};if _beeb :=_ceff .Color ;_beeb !=nil {_acgd ._geac =_ccg .getColorStringFromSmlColor (_beeb );};return _acgd ;};func (_aaa *convertContext )drawSheet (){for _ggbe ,_baff :=range _aaa ._effd {_gffg :=len (_baff ._dbd );if _ggbe ==len (_aaa ._effd )-1{for _ebff :=len (_baff ._dbd )-1;_ebff >=0;_ebff --{if !_baff ._dbd [_ebff ]._bada {_gffg =_ebff ;};};};_cag :=_baff ._dbd [:_gffg ];for _ ,_gade :=range _cag {_aaa ._ebag .NewPage ();_aaa .drawPage (_gade );};};};type line struct{_cfbc float64 ;_eeb []*symbol ;_eaec float64 ;};func (_dfeb *convertContext )getStyle (_cadg *uint32 )*style {_dgb :=&style {};_ffab :=false ;if _cadg !=nil {_gcga :=_dfeb ._ecf .GetCellStyle (*_cadg );_fbba :=_gcga .GetFont ();for _ ,_bbb :=range _fbba .Name {if _bbb !=nil {_dgb ._affg =&_bbb .ValAttr ;_ffab =true ;break ;};};for _ ,_cagd :=range _fbba .B {if _cagd !=nil {_bgge :=_cagd .ValAttr ==nil ||*_cagd .ValAttr ;_dgb ._aagg =&_bgge ;_ffab =true ;break ;};};for _ ,_ggaed :=range _fbba .I {if _ggaed !=nil {_cbb :=_ggaed .ValAttr ==nil ||*_ggaed .ValAttr ;_dgb ._ddfe =&_cbb ;_ffab =true ;break ;};};for _ ,_abe :=range _fbba .U {if _abe !=nil {_cbdc :=_abe .ValAttr ==_fa .ST_UnderlineValuesSingle ||_abe .ValAttr ==_fa .ST_UnderlineValuesUnset ;_dgb ._cba =&_cbdc ;_ffab =true ;break ;};};for _ ,_adb :=range _fbba .Sz {if _adb !=nil {_beacc :=_adb .ValAttr /12*_bc .DefaultFontSize ;_dgb ._bgb =&_beacc ;_ffab =true ;break ;};};for _ ,_geaab :=range _fbba .VertAlign {if _geaab !=nil {_babg :=_geaab .ValAttr ==_dd .ST_VerticalAlignRunSuperscript ;_dgb ._bggg =&_babg ;_aecb :=_geaab .ValAttr ==_dd .ST_VerticalAlignRunSubscript ;_dgb ._bgd =&_aecb ;_ffab =true ;break ;};};for _ ,_ebaga :=range _fbba .Color {if _ebaga !=nil {_dgb ._geac =_dfeb .getColorStringFromSmlColor (_ebaga );_ffab =true ;break ;};};_fcfd :=_gcga .GetBorder ();if _fcfd .Top !=nil {_dgb ._dea =_dfeb .getBorder (_fcfd .Top );_ffab =true ;};if _fcfd .Bottom !=nil {_dgb ._ddfa =_dfeb .getBorder (_fcfd .Bottom );_ffab =true ;};if _fcfd .Left !=nil {_dgb ._bedg =_dfeb .getBorder (_fcfd .Left );_ffab =true ;};if _fcfd .Right !=nil {_dgb ._bbg =_dfeb .getBorder (_fcfd .Right );_ffab =true ;};if _gcga .Wrapped (){_dgb ._gadc =true ;_ffab =true ;};if _ebbe :=_gcga .GetVerticalAlignment ();_ebbe !=_fa .ST_VerticalAlignmentUnset {_dgb ._gdef =_ebbe ;_ffab =true ;};if _bbbg :=_gcga .GetHorizontalAlignment ();_bbbg !=_fa .ST_HorizontalAlignmentUnset {_dgb ._bafd =_bbbg ;_ffab =true ;};};if _ffab {return _dgb ;};return nil ;};const _a =3;func (_efe *convertContext )makeCols (){_gff :=_efe ._afgb ;_ce :=_gff .X ();_daa :=[]*colInfo {};_ed :=[]colWidthRange {};if _ddb :=_ce .Cols ;len (_ddb )> 0{for _ ,_gga :=range _ddb [0].Col {_ag :=65.0;if _fgc :=_gga .WidthAttr ;_fgc !=nil {if *_fgc > 0.83{*_fgc -=0.83;};if *_fgc <=1{_ag =*_fgc *11;}else {_ag =5+*_fgc *6;};};_edc :=int (_gga .MinAttr -1);_eg :=int (_gga .MaxAttr -1);_ed =append (_ed ,colWidthRange {_decg :_edc ,_fge :_eg ,_eead :_ag ,_caee :_efe .getStyle (_gga .StyleAttr )});};};_ca :=0;for _cb :=0;_cb <=_efe ._gdad ;_cb ++{var _fgbb float64 ;var _afe *style ;if _ca >=len (_ed ){_fgbb =65;}else {_fdb :=_ed [_ca ];if _cb >=_fdb ._decg &&_cb <=_fdb ._fge {_fgbb =_fdb ._eead ;_afe =_fdb ._caee ;if _cb ==_fdb ._fge {_ca ++;};}else {_fgbb =65;};};_daa =append (_daa ,&colInfo {_fbad :_fgbb ,_bdgc :_afe });};_efe ._gde =_daa ;};func (_fbbb *convertContext )distributeAnchors (){for _ ,_gef :=range _fbbb ._acga {_baba ,_dad :=_gef ._bdbg ,_gef ._dcd ;_fgbc ,_addb :=_gef ._dfd ,_gef ._ddbg ;_de ,_efd :=_gef ._agc ,_gef ._dbbe ;_fef ,_dbb :=_gef ._fdce ,_gef ._bcgc ;var _accf ,_cdff ,_fgcg ,_cgc *page ;for _ ,_dgae :=range _fbbb ._effd {for _ ,_cfb :=range _dgae ._dbd {if _baba >=_cfb ._cbf ._dbgc &&_baba < _cfb ._cbf ._ccfg {if _fgbc >=_cfb ._fbgf ._fga &&_fgbc < _cfb ._fbgf ._afdd {_cfb ._bada =true ;_accf =_cfb ;};if _fef >=_cfb ._fbgf ._fga &&_fef < _cfb ._fbgf ._afdd {_cfb ._bada =true ;_cdff =_cfb ;};};if _de >=_cfb ._cbf ._dbgc &&_de < _cfb ._cbf ._ccfg {if _fgbc >=_cfb ._fbgf ._fga &&_fgbc < _cfb ._fbgf ._afdd {_cfb ._bada =true ;_cgc =_cfb ;};if _fef >=_cfb ._fbgf ._fga &&_fef < _cfb ._fbgf ._afdd {_cfb ._bada =true ;_fgcg =_cfb ;};};};};_geg :=_accf !=_cdff ;_bcdc :=_accf !=_cgc ;if _geg &&_bcdc {_cee :=_fbbb ._gde [_fgbc ]._agd +_be .FromEMU (_addb );_dec :=_accf ._fbgf ._gbac ;_gefb :=_fbbb ._gde [_fef ]._agd +_be .FromEMU (_dbb );_fee :=_fbbb ._dcg [_baba ]._feab +_be .FromEMU (_dad );_fcf :=float64 (_accf ._cbf ._fed );_gcec :=_fbbb ._dcg [_de ]._feab +_be .FromEMU (_efd );_cadb :=_gefb +_dec -_cee ;_cge :=_gcec +_fcf -_fee ;_cgd :=_fbbb .imageFromAnchor (_gef ,_cadb ,_cge );_accf ._acgae =append (_accf ._acgae ,_fbbb .getImage (_cgd ,_cge ,_cadb ,_cee ,_fee ,_dec -_cee ,_fcf -_fee ,_bc .ImgPart_lt ));_cdff ._acgae =append (_cdff ._acgae ,_fbbb .getImage (_cgd ,_cge ,_cadb ,0,_fee ,_dec -_cee ,_fcf -_fee ,_bc .ImgPart_rt ));_cgc ._acgae =append (_cgc ._acgae ,_fbbb .getImage (_cgd ,_cge ,_cadb ,_cee ,0,_dec -_cee ,_fcf -_fee ,_bc .ImgPart_lb ));_fgcg ._acgae =append (_fgcg ._acgae ,_fbbb .getImage (_cgd ,_cge ,_cadb ,0,0,_dec -_cee ,_fcf -_fee ,_bc .ImgPart_rb ));}else if _geg {_fdgd :=_fbbb ._dcg [_baba ]._feab +_be .FromEMU (_dad );_efbb :=_fbbb ._dcg [_de ]._feab +_be .FromEMU (_efd );_gcg :=_fbbb ._gde [_fgbc ]._agd +_be .FromEMU (_addb );_fbg :=_accf ._fbgf ._gbac ;_dbc :=_fbbb ._gde [_fef ]._agd +_be .FromEMU (_dbb );_baee :=_dbc +_fbg -_gcg ;_fgf :=_efbb -_fdgd ;_acd :=_fbbb .imageFromAnchor (_gef ,_baee ,_fgf );_accf ._acgae =append (_accf ._acgae ,_fbbb .getImage (_acd ,_fgf ,_baee ,_gcg ,_fdgd ,_fbg -_gcg ,0,_bc .ImgPart_l ));_cdff ._acgae =append (_cdff ._acgae ,_fbbb .getImage (_acd ,_fgf ,_baee ,0,_fdgd ,_fbg -_gcg ,0,_bc .ImgPart_r ));}else if _bcdc {_gdc :=_fbbb ._gde [_fgbc ]._agd +_be .FromEMU (_addb );_fbae :=_fbbb ._gde [_fef ]._agd +_be .FromEMU (_dbb );_abd :=_fbbb ._dcg [_baba ]._feab +_be .FromEMU (_dad );_eca :=float64 (_accf ._cbf ._fed );_gacg :=_fbbb ._dcg [_de ]._feab +_be .FromEMU (_efd );_caf :=_fbae -_gdc ;_acdg :=_gacg +_eca -_abd ;_abbd :=_fbbb .imageFromAnchor (_gef ,_caf ,_acdg );_accf ._acgae =append (_accf ._acgae ,_fbbb .getImage (_abbd ,_acdg ,_caf ,_gdc ,_abd ,0,_eca -_abd ,_bc .ImgPart_t ));_cgc ._acgae =append (_cgc ._acgae ,_fbbb .getImage (_abbd ,_acdg ,_caf ,_gdc ,0,0,_eca -_abd ,_bc .ImgPart_b ));}else {_geaa :=_fbbb ._gde [_fgbc ]._agd +_be .FromEMU (_addb );_bafc :=_fbbb ._gde [_fef ]._agd +_be .FromEMU (_dbb );_fdc :=_fbbb ._dcg [_baba ]._feab +_be .FromEMU (_dad );_eee :=_fbbb ._dcg [_de ]._feab +_be .FromEMU (_efd );_bge :=_bafc -_geaa ;_aag :=_eee -_fdc ;_bfa :=_fbbb .imageFromAnchor (_gef ,_bge ,_aag );_accf ._acgae =append (_accf ._acgae ,_fbbb .getImage (_bfa ,_aag ,_bge ,_geaa ,_fdc ,0,0,_bc .ImgPart_whole ));};};};const _gc =0.64;func (_dada *convertContext )getSymbolsFromR (_dega []*_fa .CT_RElt ,_cgbb *style )[]*symbol {_eff :=[]*symbol {};for _ ,_daabe :=range _dega {_cgcb :=_dada .combineCellStyleWithRPrElt (_cgbb ,_daabe .RPr );for _ ,_caa :=range _daabe .T {_eff =append (_eff ,&symbol {_ebde :string (_caa ),_bafa :_dada .makeTextStyleFromCellStyle (_cgcb )});};};return _eff ;};type colWidthRange struct{_decg int ;_fge int ;_eead float64 ;_caee *style ;};