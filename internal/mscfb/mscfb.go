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

package mscfb ;import (_a "bytes";_ae "encoding/binary";_d "fmt";_ad "github.com/richardlehane/msoleps/types";_e "github.com/unidoc/unioffice/internal/mscfb/rw";_bd "io";_f "os";_cb "strconv";_be "time";_b "unicode";_g "unicode/utf16";);func (_cffb Error )Typ ()int {return _cffb ._cfbb };func (_dgc *File )WriteAt (p []byte ,off int64 )(_afbd int ,_gbc error ){_ebgc ,_eaa ,_gfg :=_dgc ._ceg ,_dgc ._aaf ,_dgc ._gb ;_ ,_gbc =_dgc .Seek (off ,0);if _gbc ==nil {_afbd ,_gbc =_dgc .Write (p );};_dgc ._ceg ,_dgc ._aaf ,_dgc ._gb =_ebgc ,_eaa ,_gfg ;return _afbd ,_gbc ;};func _ce (_defe *File ){if _defe ._de < 4||_defe ._de > 64{return ;};_fac :=int (_defe ._de /2-1);_defe .Initial =_defe ._fg [0];var _cg int ;if !_b .IsPrint (rune (_defe .Initial )){_cg =1;};_defe .Name =string (_g .Decode (_defe ._fg [_cg :_fac ]));};func (_babg *Reader )exportMiniStream ()(*_e .Writer ,*_e .Writer ,error ){_ebc ,_cce :=_e .NewWriter (),_e .NewWriter ();_cae :=uint32 (0);for _ ,_dcff :=range _babg .File {if _dcff .Size ==0{continue ;};_dcff .reset ();_dcff ._dc =_cae ;_eef :=int (_dcff .Size )/int (_ged );if int (_dcff .Size )%int (_ged )!=0{_eef ++;};for _cfgd :=1;_cfgd < _eef ;_cfgd ++{_cae ++;if _aeca :=_ae .Write (_ebc ,_ae .LittleEndian ,_cae );_aeca !=nil {return nil ,nil ,_aeca ;};};if _bdbf :=_ae .Write (_ebc ,_ae .LittleEndian ,_caf );_bdbf !=nil {return nil ,nil ,_bdbf ;};_cae ++;if _ ,_dcfg :=_bd .Copy (_cce ,_dcff );_dcfg !=nil {return nil ,nil ,_dcfg ;};if _bfc :=_cce .AlignLength (64);_bfc !=nil {return nil ,nil ,_bfc ;};};if _agg :=_ebc .FillWithByte (int (_babg ._aeb )-_ebc .Len (),_cea );_agg !=nil {return nil ,nil ,_agg ;};if _bdg :=_cce .AlignLength (int (_babg ._aeb ));_bdg !=nil {return nil ,nil ,_bdg ;};return _ebc ,_cce ,nil ;};func _dcb (_dfe []byte )*directoryEntryFields {_ge :=&directoryEntryFields {};for _ee :=range _ge ._fg {_ge ._fg [_ee ]=_ae .LittleEndian .Uint16 (_dfe [_ee *2:_ee *2+2]);};_ge ._de =_ae .LittleEndian .Uint16 (_dfe [64:66]);_ge ._fb =uint8 (_dfe [66]);_ge ._dd =uint8 (_dfe [67]);_ge ._df =_ae .LittleEndian .Uint32 (_dfe [68:72]);_ge ._bf =_ae .LittleEndian .Uint32 (_dfe [72:76]);_ge ._af =_ae .LittleEndian .Uint32 (_dfe [76:80]);_ge ._fa =_ad .MustGuid (_dfe [80:96]);copy (_ge ._fbg [:],_dfe [96:100]);_ge ._bde =_ad .MustFileTime (_dfe [100:108]);_ge ._da =_ad .MustFileTime (_dfe [108:116]);_ge ._dc =_ae .LittleEndian .Uint32 (_dfe [116:120]);copy (_ge ._dfg [:],_dfe [120:128]);return _ge ;};func (_daa *Reader )saveToFatLocs (_aagf uint32 ,_cfga interface{},_acdb bool )error {_bdda :=_a .NewBuffer ([]byte {});if _bbe :=_ae .Write (_bdda ,_ae .LittleEndian ,_cfga );_bbe !=nil {return _bbe ;};_cefc :=_daa .findFatLocsOffset (_acdb )+int64 (_aagf *4);_ ,_ded :=_daa ._fda .WriteAt (_bdda .Bytes (),_cefc );return _ded ;};func (_ag *File )FileInfo ()_f .FileInfo {return fileInfo {_ag }};func _bee (_aa uint16 ,_gf *File ){_ce (_gf );if _gf ._fb !=_dg {return ;};if _aa > 3{_gf .Size =int64 (_ae .LittleEndian .Uint64 (_gf ._dfg [:]));}else {_gf .Size =int64 (_ae .LittleEndian .Uint32 (_gf ._dfg [:4]));};};func (_ec fileInfo )Sys ()interface{}{return nil };type Error struct{_cfbb int ;_cafbb string ;_becdd int64 ;};func (_bcef *Reader )findNextFreeSector (_cbg bool )(uint32 ,error ){_dfa :=_bcef .findFatLocsOffset (_cbg );_cef :=uint32 (0);_eccb :=_bcef ._aeb /4;for {_bbgb ,_deaa :=_bcef .readAt (_dfa ,4);if _deaa !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_deaa .Error ()+"\u0029",_dfa };};_eddg :=_ae .LittleEndian .Uint32 (_bbgb );if _eddg ==_baeb {break ;};if _cef >=_eccb {return 0,Error {ErrRead ,"\u0065\u006e\u0064\u0020of\u0020\u006d\u0069\u006e\u0069\u0046\u0061\u0074\u0020\u0072\u0065\u0061\u0063\u0068e\u0064",_dfa };};_cef ++;_dfa +=4;};return _cef ,nil ;};func (_ceag *Reader )setHeader ()error {_ege ,_dcg :=_ceag .readAt (0,_feb );if _dcg !=nil {return _dcg ;};_ceag ._dcga =&header {headerFields :_eega (_ege )};if _ceag ._dcga ._gaa !=_ccbd {return Error {ErrFormat ,"\u0062\u0061\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065",int64 (_ceag ._dcga ._gaa )};};if _ceag ._dcga ._bffa ==0x0009||_ceag ._dcga ._bffa ==0x000c{_ceag ._aeb =uint32 (1<<_ceag ._dcga ._bffa );}else {return Error {ErrFormat ,"\u0069\u006c\u006c\u0065ga\u006c\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065",int64 (_ceag ._dcga ._bffa )};};if _ceag ._dcga ._dbaf > 0{_dfb :=(_ceag ._aeb /4)-1;if int (_ceag ._dcga ._dbaf *_dfb +109)< 0{return Error {ErrFormat ,"\u0044I\u0046A\u0054\u0020\u0069\u006e\u0074 \u006f\u0076e\u0072\u0066\u006c\u006f\u0077",int64 (_ceag ._dcga ._dbaf )};};if _ceag ._dcga ._dbaf *_dfb +109> _ceag ._dcga ._agcf +_dfb {return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u0044\u0049\u0046\u0041\u0054\u0073 \u0065\u0078\u0063\u0065\u0065\u0064\u0073 \u0046\u0041\u0054\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0073",int64 (_ceag ._dcga ._dbaf )};};};if _ceag ._dcga ._fegc > 0{if int (_ceag ._aeb /4*_ceag ._dcga ._fegc )< 0{return Error {ErrFormat ,"m\u0069\u006e\u0069\u0020FA\u0054 \u0069\u006e\u0074\u0020\u006fv\u0065\u0072\u0066\u006c\u006f\u0077",int64 (_ceag ._dcga ._fegc )};};if _ceag ._dcga ._fegc > _ceag ._dcga ._agcf *(_ceag ._aeb /_ged ){return Error {ErrFormat ,"\u006e\u0075\u006d\u0020\u006d\u0069n\u0069\u0020\u0046\u0041\u0054\u0073\u0020\u0065\u0078\u0063\u0065\u0065\u0064s\u0020\u0046\u0041\u0054\u0020\u0073\u0065c\u0074\u006f\u0072\u0073",int64 (_ceag ._dcga ._agcf )};};};return nil ;};func (_ga *File )ID ()string {return _ga ._fa .String ()};func _fae (_ccb ,_ebef uint32 )int64 {return int64 ((_ebef +1)*_ccb )};type header struct{*headerFields ;_bba []uint32 ;_dgb []uint32 ;_bgdc []uint32 ;};const (_aee uint8 =0x0;_ef uint8 =0x1;);func New (ra _bd .ReaderAt )(*Reader ,error ){_bfd :=&Reader {_aed :ra };if _ ,_dcca :=ra .(slicer );_dcca {_bfd ._gcag =true ;}else {_bfd ._geg =make ([]byte ,_feb );};if _bafc :=_bfd .setHeader ();_bafc !=nil {return nil ,_bafc ;};if !_bfd ._gcag &&int (_bfd ._aeb )> len (_bfd ._geg ){_bfd ._geg =make ([]byte ,_bfd ._aeb );};if _dggb :=_bfd .setDifats ();_dggb !=nil {return nil ,_dggb ;};if _afc :=_bfd .setDirEntries ();_afc !=nil {return nil ,_afc ;};if _dgdb :=_bfd .setMiniStream ();_dgdb !=nil {return nil ,_dgdb ;};if _dgae :=_bfd .traverse ();_dgae !=nil {return nil ,_dgae ;};return _bfd ,nil ;};func (_dca fileInfo )Name ()string {return _dca .File .Name };const (_eddb uint32 =0xFFFFFFFA;_gba uint32 =0xFFFFFFFC;_cba uint32 =0xFFFFFFFD;_caf uint32 =0xFFFFFFFE;_baeb uint32 =0xFFFFFFFF;_cea byte =0xFF;_efg uint32 =0xFFFFFFFA;_dea uint32 =0xFFFFFFFF;);const (ErrFormat =iota ;ErrRead ;ErrSeek ;ErrWrite ;ErrTraverse ;);func (_ba fileInfo )Mode ()_f .FileMode {return _ba .File .mode ()};func (_ddb *File )Write (b []byte )(int ,error ){if _ddb .Size < 1||_ddb ._ceg >=_ddb .Size {return 0,_bd .EOF ;};if _eaf :=_ddb .ensureWriterAt ();_eaf !=nil {return 0,_eaf ;};_acgb :=len (b );if int64 (_acgb )> _ddb .Size -_ddb ._ceg {_acgb =int (_ddb .Size -_ddb ._ceg );};_gag ,_cd :=_ddb .stream (_acgb );if _cd !=nil {return 0,_cd ;};var _ffg ,_bcb int ;for _ ,_cc :=range _gag {_dgd :=_ffg +int (_cc [1]);if _dgd < _ffg ||_dgd > _acgb {return 0,Error {ErrWrite ,"\u0062\u0061d\u0020\u0077\u0072i\u0074\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_dgd )};};_cdf ,_dec :=_ddb ._db ._fda .WriteAt (b [_ffg :_dgd ],_cc [0]);_bcb =_bcb +_cdf ;if _dec !=nil {_ddb ._ceg +=int64 (_bcb );return _bcb ,Error {ErrWrite ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0077\u0072i\u0074\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_dec .Error ()+"\u0029",int64 (_ffg )};};_ffg =_dgd ;};_ddb ._ceg +=int64 (_bcb );if _bcb !=_acgb {_cd =Error {ErrWrite ,"\u0062\u0079t\u0065\u0073\u0020\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0077\u0072\u0069\u0074\u0065\u0020\u0073\u0069\u007a\u0065",int64 (_bcb )};}else if _bcb < len (b ){_cd =_bd .EOF ;};return _bcb ,_cd ;};type fileInfo struct{*File };func (_cgg *File )SetEntryContent (b []byte )error {if _bce :=_cgg .ensureWriterAt ();_bce !=nil {return _bce ;};_cgg .reset ();if _bcd :=_cgg .changeSize (int64 (len (b )));_bcd !=nil {return _bcd ;};_ ,_aea :=_cgg .Write (b );return _aea ;};type headerFields struct{_gaa uint64 ;_ [16]byte ;_gad uint16 ;_eccg uint16 ;_ [2]byte ;_bffa uint16 ;_ [2]byte ;_ [6]byte ;_ecf uint32 ;_agcf uint32 ;_becd uint32 ;_ [4]byte ;_ [4]byte ;_geda uint32 ;_fegc uint32 ;_adfa uint32 ;_dbaf uint32 ;_ddbf [109]uint32 ;};func _bff (_beec [][2]int64 )[][2]int64 {_acb :=len (_beec );for _gfgg ,_fbf :=0,0;_gfgg < _acb &&_fbf +1< len (_beec );_gfgg ++{if _beec [_fbf ][0]+_beec [_fbf ][1]==_beec [_fbf +1][0]{_beec [_fbf ][1]=_beec [_fbf ][1]+_beec [_fbf +1][1];for _agd :=range _beec [_fbf +1:len (_beec )-1]{_beec [_fbf +1+_agd ]=_beec [_agd +_fbf +2];};_beec =_beec [:len (_beec )-1];}else {_fbf +=1;};};return _beec ;};func (_afaa *Reader )exportFAT (_efe *_e .Writer ,_bbfb []uint32 )error {if _afaa ._dcga ._agcf ==0{return nil ;};_dcf :=_a .NewBuffer ([]byte {});if _ccbg :=_ae .Write (_dcf ,_ae .LittleEndian ,_cba );_ccbg !=nil {return _ccbg ;};for _abb :=0;_abb < len (_bbfb )-1;_abb ++{for _dddg :=_bbfb [_abb ];_dddg < _bbfb [_abb +1]-1;_dddg ++{if _abcf :=_ae .Write (_dcf ,_ae .LittleEndian ,_dddg );_abcf !=nil {return _abcf ;};};if _gcf :=_ae .Write (_dcf ,_ae .LittleEndian ,_caf );_gcf !=nil {return _gcf ;};};_dceg :=512;for _ ,_bdefc :=range _dcf .Bytes (){if _cfg :=_efe .WriteByteAt (_bdefc ,_dceg );_cfg !=nil {return _cfg ;};_dceg ++;};return nil ;};func (_bbc *Reader )findNext (_cda uint32 ,_bbcf bool )(uint32 ,error ){_fgcdd :=_bbc ._aeb /4;_beae :=int (_cda /_fgcdd );var _ccba uint32 ;if _bbcf {if _beae < 0||_beae >=len (_bbc ._dcga ._dgb ){return 0,Error {ErrRead ,"\u006d\u0069\u006e\u0069\u0073e\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073 \u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_beae )};};_ccba =_bbc ._dcga ._dgb [_beae ];}else {if _beae < 0||_beae >=len (_bbc ._dcga ._bba ){return 0,Error {ErrRead ,"\u0046\u0041\u0054\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073\u0020\u006f\u0075t\u0073i\u0064\u0065\u0020\u0044\u0049\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_beae )};};_ccba =_bbc ._dcga ._bba [_beae ];};_fag :=_cda %_fgcdd ;_fbd :=_fae (_bbc ._aeb ,_ccba )+int64 (_fag *4);_dbab ,_ffe :=_bbc .readAt (_fbd ,4);if _ffe !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_ffe .Error ()+"\u0029",_fbd };};_dfgg :=_ae .LittleEndian .Uint32 (_dbab );return _dfgg ,nil ;};func (_ada *Reader )GetEntry (entryName string )(*File ,error ){for _aba ,_aadd :=_ada .Next ();_aadd ==nil ;_aba ,_aadd =_ada .Next (){if _aba .Name ==entryName {return _aba ,nil ;};};return nil ,Error {ErrTraverse ,"\u004e\u006f\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072 \u0067\u0069\u0076\u0065\u006e\u0020\u006e\u0061\u006d\u0065\u002e",0};};const (_ff uint8 =0x0;_bb uint8 =0x1;_dg uint8 =0x2;_gd uint8 =0x5;);type directoryEntryFields struct{_fg [32]uint16 ;_de uint16 ;_fb uint8 ;_dd uint8 ;_df uint32 ;_bf uint32 ;_af uint32 ;_fa _ad .Guid ;_fbg [4]byte ;_bde _ad .FileTime ;_da _ad .FileTime ;_dc uint32 ;_dfg [8]byte ;};func (_bafe *Reader )Next ()(*File ,error ){_bafe ._fdf ++;if _bafe ._fdf >=len (_bafe .File ){return nil ,_bd .EOF ;};return _bafe .File [_bafe ._fdf ],nil ;};func (_ecea *File )stream (_bdf int )([][2]int64 ,error ){var _bdce bool ;var _dbe int ;var _bad int64 ;if _ecea .Size < _fgb {_bdce =true ;_dbe =_bdf /int (_ged )+2;_bad =int64 (_ged );}else {_dbe =_bdf /int (_ecea ._db ._aeb )+2;_bad =int64 (_ecea ._db ._aeb );};_cca :=make ([][2]int64 ,0,_dbe );var _afad ,_bec int ;if _ecea ._aaf > 0{_dbeg ,_ecc :=_ecea ._db .getOffset (_ecea ._gb ,_bdce );if _ecc !=nil {return nil ,_ecc ;};if _bad -_ecea ._aaf >=int64 (_bdf ){_cca =append (_cca ,[2]int64 {_dbeg +_ecea ._aaf ,int64 (_bdf )});}else {_cca =append (_cca ,[2]int64 {_dbeg +_ecea ._aaf ,_bad -_ecea ._aaf });};if _bad -_ecea ._aaf <=int64 (_bdf ){_ecea ._gb ,_ecc =_ecea ._db .findNext (_ecea ._gb ,_bdce );if _ecc !=nil {return nil ,_ecc ;};_bec +=int (_bad -_ecea ._aaf );_ecea ._aaf =0;}else {_ecea ._aaf +=int64 (_bdf );};if _cca [0][1]==int64 (_bdf ){return _cca ,nil ;};if _ecea ._gb ==_caf {return nil ,Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_ecea ._gb )};};_afad ++;};for {if _afad >=cap (_cca ){return nil ,Error {ErrRead ,"\u0069\u006e\u0064\u0065x\u0020\u006f\u0076\u0065\u0072\u0072\u0075\u006e\u0073\u0020s\u0065c\u0074\u006f\u0072\u0020\u006c\u0065\u006eg\u0074\u0068",int64 (_afad )};};_fcb ,_gfa :=_ecea ._db .getOffset (_ecea ._gb ,_bdce );if _gfa !=nil {return nil ,_gfa ;};if _bdf -_bec < int (_bad ){_cca =append (_cca ,[2]int64 {_fcb ,int64 (_bdf -_bec )});_ecea ._aaf =int64 (_bdf -_bec );return _bff (_cca ),nil ;}else {_cca =append (_cca ,[2]int64 {_fcb ,_bad });_bec +=int (_bad );_ecea ._gb ,_gfa =_ecea ._db .findNext (_ecea ._gb ,_bdce );if _gfa !=nil {return nil ,_gfa ;};if _bec ==_bdf {return _bff (_cca ),nil ;};};_afad ++;};};func (_afa *Reader )setDirEntries ()error {_faa :=20;if _afa ._dcga ._ecf > 0{_faa =int (_afa ._dcga ._ecf );};_deb :=make ([]*File ,0,_faa );_bdd :=make (map[uint32 ]bool );_fgc :=int (_afa ._aeb /_bcee );_eg :=_afa ._dcga ._becd ;for _eg !=_caf {_def ,_ebg :=_afa .readAt (_fae (_afa ._aeb ,_eg ),int (_afa ._aeb ));if _ebg !=nil {return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020r\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_ebg .Error ()+"\u0029",_fae (_afa ._aeb ,_eg )};};for _gdg :=0;_gdg < _fgc ;_gdg ++{_fc :=&File {_db :_afa };_fc .directoryEntryFields =_dcb (_def [_gdg *int (_bcee ):]);_bee (_afa ._dcga ._eccg ,_fc );_fc ._gb =_fc ._dc ;_deb =append (_deb ,_fc );};_eeb ,_ebg :=_afa .findNext (_eg ,false );if _ebg !=nil {return Error {ErrRead ,"d\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0065r\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006eg \u0073\u0065\u0063t\u006fr\u0020\u0028"+_ebg .Error ()+"\u0029",int64 (_eeb )};};if _eeb <=_eg {if _eeb ==_eg ||_bdd [_eeb ]{return Error {ErrRead ,"\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020s\u0065\u0063\u0074\u006f\u0072\u0020\u0063\u0079\u0063\u006c\u0065",int64 (_eeb )};};_bdd [_eeb ]=true ;};_eg =_eeb ;};_afa ._egae =_deb ;return nil ;};func (_dgcf *Reader )GetHeader ()*header {return _dgcf ._dcga };func (_fbdg *Reader )ID ()string {return _fbdg .File [0].ID ()};func (_bcc *Reader )Modified ()_be .Time {return _bcc .File [0].Modified ()};func (_aagg *Reader )exportDifats (_dbaa *_e .Writer )error {if _aagg ._dcga ._dbaf ==0{return nil ;};return nil ;};func (_dcef *Reader )Created ()_be .Time {return _dcef .File [0].Created ()};func (_fbdd Error )Error ()string {return "\u006ds\u0063\u0066\u0062\u003a\u0020"+_fbdd ._cafbb +"\u003b\u0020"+_cb .FormatInt (_fbdd ._becdd ,10);};func (_ecec *Reader )getOffset (_bgbd uint32 ,_bdef bool )(int64 ,error ){if _bdef {_ccad :=_ecec ._aeb /64;_ccg :=int (_bgbd /_ccad );if _ccg >=len (_ecec ._dcga ._bgdc ){return 0,Error {ErrRead ,"\u006di\u006e\u0069s\u0065\u0063\u0074o\u0072\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0069\u0073\u0020\u006f\u0075t\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0073\u0065c\u0074\u006f\u0072\u0020\u0072\u0061\u006e\u0067\u0065",int64 (_ccg )};};_eag :=_bgbd %_ccad ;return int64 ((_ecec ._dcga ._bgdc [_ccg ]+1)*_ecec ._aeb +_eag *64),nil ;};return _fae (_ecec ._aeb ,_bgbd ),nil ;};type Reader struct{_gcag bool ;_aeb uint32 ;_geg []byte ;_dcga *header ;File []*File ;_egae []*File ;_fdf int ;_aed _bd .ReaderAt ;_fda _bd .WriterAt ;};func (_fd *File )changeSize (_ccc int64 )error {if _ccc ==_fd .Size {return nil ;};var _bbg *File ;for _ ,_afb :=range _fd ._db ._egae {if _afb .Name ==_fd .Name {_bbg =_afb ;break ;};};if _bbg ==nil {return _d .Errorf ("\u004e\u006f\u0020\u0064\u0069\u0072e\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0072\u0020\u0061\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_fd .Name );};_fgg :=_a .NewBuffer ([]byte {});if _abg :=_ae .Write (_fgg ,_ae .LittleEndian ,_ccc );_abg !=nil {return _abg ;};for _gdab ,_bdca :=range _fgg .Bytes (){_bbg ._dfg [_gdab ]=_bdca ;};var _bbf int64 ;var _gg bool ;if _fd .Size < _fgb {_gg =true ;_bbf =int64 (_ged );}else {_bbf =int64 (_fd ._db ._aeb );};_ddf :=int ((_fd .Size -1)/_bbf )+1;_dgg :=int ((_ccc -1)/_bbf )+1;if _ddf < _dgg {_aec ,_fe :=_fd .findLast (_gg );if _fe !=nil {return _fe ;};_cabe ,_fe :=_fd ._db .findNextFreeSector (_gg );if _fe !=nil {return _fe ;};for _cbe :=_dgg -_ddf ;_cbe > 0;_cbe --{if _dag :=_fd ._db .saveToFatLocs (_aec ,_cabe ,_gg );_dag !=nil {return _dag ;};if _cbe > 1{_aec =_cabe ;_cabe ++;}else if _cga :=_fd ._db .saveToFatLocs (_cabe ,_caf ,_gg );_cga !=nil {return _cga ;};};}else if _ddf > _dgg {_gcd :=_fd ._dc ;var _ade error ;for _ddgf :=0;_ddgf < _dgg -1;_ddgf ++{_gcd ,_ade =_fd ._db .findNext (_gcd ,_gg );if _ade !=nil {return _ade ;};};if _cff :=_fd ._db .saveToFatLocs (_gcd ,_caf ,_gg );_cff !=nil {return _cff ;};};_fd .Size =_ccc ;return nil ;};func (_dcbe *Reader )Debug ()map[string ][]uint32 {_eca :=map[string ][]uint32 {"s\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065":[]uint32 {_dcbe ._aeb },"\u006d\u0069\u006e\u0069\u0020\u0066\u0061\u0074\u0020\u006c\u006f\u0063\u0073":_dcbe ._dcga ._dgb ,"\u006d\u0069n\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006c\u006f\u0063\u0073":_dcbe ._dcga ._bgdc ,"\u0064\u0069r\u0065\u0063\u0074o\u0072\u0079\u0020\u0073\u0065\u0063\u0074\u006f\u0072":[]uint32 {_dcbe ._dcga ._becd },"\u006d\u0069\u006e\u0069 s\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069z\u0065":[]uint32 {_dcbe .File [0]._dc ,_ae .LittleEndian .Uint32 (_dcbe .File [0]._dfg [:])}};for _cdb ,_cbfc :=_dcbe .Next ();_cbfc ==nil ;_cdb ,_cbfc =_dcbe .Next (){_eca [_cdb .Name +" \u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069\u007a\u0065"]=[]uint32 {_cdb ._dc ,_ae .LittleEndian .Uint32 (_cdb ._dfg [:])};};return _eca ;};func (_adb *Reader )setDifats ()error {_adb ._dcga ._bba =_adb ._dcga ._ddbf [:];if _adb ._dcga ._dbaf ==0{return nil ;};_ddc :=(_adb ._aeb /4)-1;_fbfd :=make ([]uint32 ,109,_adb ._dcga ._dbaf *_ddc +109);copy (_fbfd ,_adb ._dcga ._bba );_adb ._dcga ._bba =_fbfd ;_cafb :=_adb ._dcga ._adfa ;for _eaag :=0;_eaag < int (_adb ._dcga ._dbaf );_eaag ++{_cfcb ,_cad :=_adb .readAt (_fae (_adb ._aeb ,_cafb ),int (_adb ._aeb ));if _cad !=nil {return Error {ErrFormat ,"e\u0072r\u006f\u0072\u0020\u0073\u0065\u0074\u0074\u0069n\u0067\u0020\u0044\u0049FA\u0054\u0028"+_cad .Error ()+"\u0029",int64 (_cafb )};};for _baf :=0;_baf < int (_ddc );_baf ++{_adb ._dcga ._bba =append (_adb ._dcga ._bba ,_ae .LittleEndian .Uint32 (_cfcb [_baf *4:_baf *4+4]));};_cafb =_ae .LittleEndian .Uint32 (_cfcb [len (_cfcb )-4:]);};return nil ;};func (_aebc *Reader )Export ()([]byte ,error ){_eebd :=_e .NewWriter ();if _dda :=_aebc .exportHeader (_eebd );_dda !=nil {return nil ,_dda ;};if _dfag :=_eebd .FillWithByte (512,_cea );_dfag !=nil {return nil ,_dfag ;};_fgf :=[]uint32 {};if _gdc :=_aebc .exportDifats (_eebd );_gdc !=nil {return nil ,_gdc ;};_gge ,_bgde ,_ggf :=_aebc .exportMiniStream ();if _ggf !=nil {return nil ,_ggf ;};_fgf =append (_fgf ,uint32 (_eebd .Len ())/_aebc ._aeb );if _fcfb :=_aebc .exportDirEntries (_eebd );_fcfb !=nil {return nil ,_fcfb ;};_fgf =append (_fgf ,uint32 (_eebd .Len ())/_aebc ._aeb );if _ ,_febb :=_gge .WriteTo (_eebd );_febb !=nil {return nil ,_febb ;};_fgf =append (_fgf ,uint32 (_eebd .Len ())/_aebc ._aeb );if _ ,_dad :=_bgde .WriteTo (_eebd );_dad !=nil {return nil ,_dad ;};_fgf =append (_fgf ,uint32 (_eebd .Len ())/_aebc ._aeb );if _bcba :=_aebc .exportFAT (_eebd ,_fgf );_bcba !=nil {return nil ,_bcba ;};return _eebd .Bytes (),nil ;};func (_ed fileInfo )ModTime ()_be .Time {return _ed .Modified ()};type slicer interface{Slice (_aca int64 ,_ggba int )([]byte ,error );};func (_eae *Reader )setMiniStream ()error {if _eae ._egae [0]._dc ==_caf ||_eae ._dcga ._geda ==_caf ||_eae ._dcga ._fegc ==0{return nil ;};_cfad :=int (_eae ._dcga ._fegc );_eae ._dcga ._dgb =make ([]uint32 ,_cfad );_eae ._dcga ._dgb [0]=_eae ._dcga ._geda ;for _acf :=1;_acf < _cfad ;_acf ++{_cbd ,_cfe :=_eae .findNext (_eae ._dcga ._dgb [_acf -1],false );if _cfe !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_cfe .Error ()+"\u0029",int64 (_eae ._dcga ._dgb [_acf -1])};};_eae ._dcga ._dgb [_acf ]=_cbd ;};_cfad =int (_eae ._aeb /4*_eae ._dcga ._fegc );_eae ._dcga ._bgdc =make ([]uint32 ,0,_cfad );_edb :=_eae ._egae [0]._dc ;var _dggad error ;for _edb !=_caf {_eae ._dcga ._bgdc =append (_eae ._dcga ._bgdc ,_edb );_edb ,_dggad =_eae .findNext (_edb ,false );if _dggad !=nil {return Error {ErrFormat ,"s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028"+_dggad .Error ()+"\u0029",int64 (_edb )};};};return nil ;};func (_dce *File )ReadAt (p []byte ,off int64 )(_cac int ,_fgcf error ){_fdc ,_aag ,_feg :=_dce ._ceg ,_dce ._aaf ,_dce ._gb ;_ ,_fgcf =_dce .Seek (off ,0);if _fgcf ==nil {_cac ,_fgcf =_dce .Read (p );};_dce ._ceg ,_dce ._aaf ,_dce ._gb =_fdc ,_aag ,_feg ;return _cac ,_fgcf ;};func (_cbf fileInfo )IsDir ()bool {return _cbf .mode ().IsDir ()};func (_gfb *Reader )traverse ()error {_gfb .File =make ([]*File ,0,len (_gfb ._egae ));var (_cf func (int ,[]string );_bda error ;_fgd int ;);_cf =func (_ca int ,_bdc []string ){_fgd ++;if _fgd > len (_gfb ._egae ){_bda =Error {ErrTraverse ,"\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0063o\u0075\u006e\u0074\u0065\u0072\u0020\u006f\u0076\u0065\u0072f\u006c\u006f\u0077",int64 (_ca )};return ;};if _ca < 0||_ca >=len (_gfb ._egae ){_bda =Error {ErrTraverse ,"\u0069\u006c\u006ceg\u0061\u006c\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0069\u006e\u0064\u0065\u0078",int64 (_ca )};return ;};_gc :=_gfb ._egae [_ca ];if _gc ._df !=_dea {_cf (int (_gc ._df ),_bdc );};_gfb .File =append (_gfb .File ,_gc );_gc .Path =_bdc ;if _gc ._af !=_dea {if _ca > 0{_cf (int (_gc ._af ),append (_bdc ,_gc .Name ));}else {_cf (int (_gc ._af ),_bdc );};};if _gc ._bf !=_dea {_cf (int (_gc ._bf ),_bdc );};return ;};_cf (0,[]string {});return _bda ;};const _feb int =8+16+10+6+12+8+16+109*4;func _eega (_acga []byte )*headerFields {_bgf :=&headerFields {};_bgf ._gaa =_ae .LittleEndian .Uint64 (_acga [:8]);_bgf ._gad =_ae .LittleEndian .Uint16 (_acga [24:26]);_bgf ._eccg =_ae .LittleEndian .Uint16 (_acga [26:28]);_bgf ._bffa =_ae .LittleEndian .Uint16 (_acga [30:32]);_bgf ._ecf =_ae .LittleEndian .Uint32 (_acga [40:44]);_bgf ._agcf =_ae .LittleEndian .Uint32 (_acga [44:48]);_bgf ._becd =_ae .LittleEndian .Uint32 (_acga [48:52]);_bgf ._geda =_ae .LittleEndian .Uint32 (_acga [60:64]);_bgf ._fegc =_ae .LittleEndian .Uint32 (_acga [64:68]);_bgf ._adfa =_ae .LittleEndian .Uint32 (_acga [68:72]);_bgf ._dbaf =_ae .LittleEndian .Uint32 (_acga [72:76]);var _gcc int ;for _gaad :=76;_gaad < 512;_gaad =_gaad +4{_bgf ._ddbf [_gcc ]=_ae .LittleEndian .Uint32 (_acga [_gaad :_gaad +4]);_gcc ++;};return _bgf ;};func (_ac *File )Modified ()_be .Time {return _ac ._da .Time ()};const (_ccbd uint64 =0xE11AB1A1E011CFD0;_ged uint32 =64;_fgb int64 =4096;_bcee uint32 =128;);func (_fgcd fileInfo )Size ()int64 {if _fgcd ._fb !=_dg {return 0;};return _fgcd .File .Size ;};func (_eeg *Reader )exportDirEntries (_fad *_e .Writer )error {if int64 (_fad .Len ())!=_fae (_eeg ._aeb ,_eeg ._dcga ._becd ){return Error {ErrWrite ,_d .Sprintf ("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020l\u0065\u006e\u0067t\u0068:\u0020\u0025\u0076",_fad .Len ()),0};};for _ ,_dfc :=range _eeg ._egae {_abf ,_gfgga :=_dgdd (_dfc .directoryEntryFields );if _gfgga !=nil {return _gfgga ;};if _ ,_dfec :=_bd .Copy (_fad ,_abf );_dfec !=nil {return _dfec ;};};return nil ;};func (_cbb *Reader )Read (b []byte )(_acd int ,_ddgg error ){if _cbb ._fdf >=len (_cbb .File ){return 0,_bd .EOF ;};return _cbb .File [_cbb ._fdf ].Read (b );};func (_ddgc *File )seek (_gfgb int64 )error {var _cfa bool ;var _faaf int64 ;if _ddgc .Size < _fgb {_cfa =true ;_faaf =64;}else {_faaf =int64 (_ddgc ._db ._aeb );};var _fbc int64 ;var _dcc error ;if _ddgc ._aaf > 0{if _faaf -_ddgc ._aaf <=_gfgb {_ddgc ._gb ,_dcc =_ddgc ._db .findNext (_ddgc ._gb ,_cfa );if _dcc !=nil {return _dcc ;};_fbc +=_faaf -_ddgc ._aaf ;_ddgc ._aaf =0;if _fbc ==_gfgb {return nil ;};}else {_ddgc ._aaf +=_gfgb ;return nil ;};if _ddgc ._gb ==_caf {return Error {ErrRead ,"\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e",int64 (_ddgc ._gb )};};};for {if _gfgb -_fbc < _faaf {_ddgc ._aaf =_gfgb -_fbc ;return nil ;}else {_fbc +=_faaf ;_ddgc ._gb ,_dcc =_ddgc ._db .findNext (_ddgc ._gb ,_cfa );if _dcc !=nil {return _dcc ;};if _fbc ==_gfgb {return nil ;};};};};func (_cbc *File )ensureWriterAt ()error {if _cbc ._db ._fda ==nil {_ebb ,_acg :=_cbc ._db ._aed .(_bd .WriterAt );if !_acg {return Error {ErrWrite ,"\u006d\u0073\u0063\u0066\u0062\u002e\u004ee\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067\u0069\u0076\u0065n\u0020R\u0065\u0061\u0064\u0065\u0072\u0041t\u0020\u0063\u006f\u006e\u0076\u0065\u0072t\u0069\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0020\u0069\u006f\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u0041\u0074\u0020\u0069n\u0020\u006f\u0072\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0077\u0072\u0069t\u0065",0};};_cbc ._db ._fda =_ebb ;};return nil ;};func (_eaff *File )findLast (_aad bool )(uint32 ,error ){_ega :=_eaff ._dc ;for {_bae ,_aeg :=_eaff ._db .findNext (_ega ,_aad );if _aeg !=nil {return 0,Error {ErrRead ,"\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028"+_aeg .Error ()+"\u0029",0};};if _bae ==_caf {break ;};_ega =_bae ;};return _ega ,nil ;};func (_cfeb *Reader )findFatLocsOffset (_faf bool )int64 {var _bgdb uint32 ;if _faf {_bgdb =_cfeb ._dcga ._dgb [0];}else {_bgdb =_cfeb ._dcga ._bba [0];};return _fae (_cfeb ._aeb ,_bgdb );};type File struct{Name string ;Initial uint16 ;Path []string ;Size int64 ;_ceg int64 ;_gb uint32 ;_aaf int64 ;*directoryEntryFields ;_db *Reader ;};func (_egea *Reader )exportHeader (_cfea *_e .Writer )error {if _fga :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._gaa );_fga !=nil {return _fga ;};if _cgb :=_cfea .Skip (16);_cgb !=nil {return _cgb ;};if _gega :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._gad );_gega !=nil {return _gega ;};if _fge :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._eccg );_fge !=nil {return _fge ;};if _gea :=_ae .Write (_cfea ,_ae .LittleEndian ,uint16 (0xfffe));_gea !=nil {return _gea ;};if _aeag :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._bffa );_aeag !=nil {return _aeag ;};if _ffgf :=_ae .Write (_cfea ,_ae .LittleEndian ,uint16 (0x0006));_ffgf !=nil {return _ffgf ;};if _aef :=_cfea .Skip (6);_aef !=nil {return _aef ;};if _ebec :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._ecf );_ebec !=nil {return _ebec ;};if _edf :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._agcf );_edf !=nil {return _edf ;};if _cag :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._becd );_cag !=nil {return _cag ;};if _adg :=_cfea .Skip (4);_adg !=nil {return _adg ;};if _ggec :=_ae .Write (_cfea ,_ae .LittleEndian ,uint32 (0x00001000));_ggec !=nil {return _ggec ;};if _gee :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._geda );_gee !=nil {return _gee ;};if _ddd :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._fegc );_ddd !=nil {return _ddd ;};if _efd :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._adfa );_efd !=nil {return _efd ;};if _gcef :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._dbaf );_gcef !=nil {return _gcef ;};for _cde :=0;_cde < 109;_cde ++{if _gdca :=_ae .Write (_cfea ,_ae .LittleEndian ,&_egea ._dcga ._ddbf [_cde ]);_gdca !=nil {return _gdca ;};};return nil ;};func (_gca *File )Created ()_be .Time {return _gca ._bde .Time ()};const _eb int =64+4*4+16+4+8*2+4+8;func (_agc *File )reset (){_agc ._ceg =0;_agc ._aaf =0;_agc ._gb =_agc ._dc };func (_eee *File )mode ()_f .FileMode {if _eee ._fb !=_dg {return _f .ModeDir |0777;};return 0666;};func (_cab *File )Read (b []byte )(int ,error ){if _cab .Size < 1||_cab ._ceg >=_cab .Size {return 0,_bd .EOF ;};_bc :=len (b );if int64 (_bc )> _cab .Size -_cab ._ceg {_bc =int (_cab .Size -_cab ._ceg );};_ab ,_bbb :=_cab .stream (_bc );if _bbb !=nil {return 0,_bbb ;};var _afg ,_gda int ;for _ ,_dbg :=range _ab {_egd :=_afg +int (_dbg [1]);if _egd < _afg ||_egd > _bc {return 0,Error {ErrRead ,"\u0062a\u0064 \u0072\u0065\u0061\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068",int64 (_egd )};};_aff ,_cgd :=_cab ._db ._aed .ReadAt (b [_afg :_egd ],_dbg [0]);_gda =_gda +_aff ;if _cgd !=nil {_cab ._ceg +=int64 (_gda );return _gda ,Error {ErrRead ,"\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0072\u0065a\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028"+_cgd .Error ()+"\u0029",int64 (_afg )};};_afg =_egd ;};_cab ._ceg +=int64 (_gda );if _gda !=_bc {_bbb =Error {ErrRead ,"\u0062\u0079\u0074e\u0073\u0020\u0072\u0065\u0061\u0064\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020r\u0065\u0061\u0064\u0020\u0073\u0069\u007a\u0065",int64 (_gda )};}else if _gda < len (b ){_bbb =_bd .EOF ;};return _gda ,_bbb ;};func _dgdd (_bab *directoryEntryFields )(*_a .Buffer ,error ){_cfc :=_a .NewBuffer ([]byte {});for _ ,_dga :=range _bab ._fg {if _bdb :=_ae .Write (_cfc ,_ae .LittleEndian ,&_dga );_bdb !=nil {return nil ,_bdb ;};};if _ede :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._de );_ede !=nil {return nil ,_ede ;};if _dba :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._fb );_dba !=nil {return nil ,_dba ;};if _bg :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._dd );_bg !=nil {return nil ,_bg ;};if _bbfa :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._df );_bbfa !=nil {return nil ,_bbfa ;};if _adf :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._bf );_adf !=nil {return nil ,_adf ;};if _gff :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._af );_gff !=nil {return nil ,_gff ;};if _cfb :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._fa .DataA );_cfb !=nil {return nil ,_cfb ;};if _bgd :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._fa .DataB );_bgd !=nil {return nil ,_bgd ;};if _gcaa :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._fa .DataC );_gcaa !=nil {return nil ,_gcaa ;};if _ ,_bfed :=_cfc .Write (_bab ._fa .DataD [:]);_bfed !=nil {return nil ,_bfed ;};if _ ,_bea :=_cfc .Write (_bab ._fbg [:]);_bea !=nil {return nil ,_bea ;};if _dbae :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._bde .Low );_dbae !=nil {return nil ,_dbae ;};if _gce :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._bde .High );_gce !=nil {return nil ,_gce ;};if _edd :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._da .Low );_edd !=nil {return nil ,_edd ;};if _daf :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._da .High );_daf !=nil {return nil ,_daf ;};if _egac :=_ae .Write (_cfc ,_ae .LittleEndian ,&_bab ._dc );_egac !=nil {return nil ,_egac ;};if _ ,_efa :=_cfc .Write (_bab ._dfg [:]);_efa !=nil {return nil ,_efa ;};return _cfc ,nil ;};func (_ggb *Reader )readAt (_gbb int64 ,_eegd int )([]byte ,error ){if _ggb ._gcag {_geb ,_faea :=_ggb ._aed .(slicer ).Slice (_gbb ,_eegd );if _faea !=nil {return nil ,Error {ErrRead ,"\u0073\u006c\u0069\u0063er\u0020\u0072\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028"+_faea .Error ()+"\u0029",_gbb };};return _geb ,nil ;};if _eegd > len (_ggb ._geg ){return nil ,Error {ErrRead ,"\u0072\u0065ad\u0020\u006c\u0065n\u0067\u0074\u0068\u0020gre\u0061te\u0072\u0020\u0074\u0068\u0061\u006e\u0020re\u0061\u0064\u0020\u0062\u0075\u0066\u0066e\u0072",int64 (_eegd )};};if _ ,_cfcf :=_ggb ._aed .ReadAt (_ggb ._geg [:_eegd ],_gbb );_cfcf !=nil {return nil ,Error {ErrRead ,_cfcf .Error (),_gbb };};return _ggb ._geg [:_eegd ],nil ;};func (_beg *File )Seek (offset int64 ,whence int )(int64 ,error ){var _bega int64 ;switch whence {default:return 0,Error {ErrSeek ,"\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077h\u0065\u006e\u0063\u0065",int64 (whence )};case 0:_bega =offset ;case 1:_bega =_beg ._ceg +offset ;case 2:_bega =_beg .Size -offset ;};switch {case _bega < 0:return _beg ._ceg ,Error {ErrSeek ,"\u0063\u0061\u006e'\u0074\u0020\u0073\u0065e\u006b\u0020\u0062\u0065\u0066\u006f\u0072e\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0046\u0069\u006c\u0065",_bega };case _bega >=_beg .Size :return _beg ._ceg ,Error {ErrSeek ,"c\u0061\u006e\u0027\u0074\u0020\u0073e\u0065\u006b\u0020\u0070\u0061\u0073\u0074\u0020\u0046i\u006c\u0065\u0020l\u0065n\u0067\u0074\u0068",_bega };case _bega ==_beg ._ceg :return _bega ,nil ;case _bega > _beg ._ceg :_gab :=_beg ._ceg ;_beg ._ceg =_bega ;return _beg ._ceg ,_beg .seek (_bega -_gab );};if _beg ._aaf >=_beg ._ceg -_bega {_beg ._aaf =_beg ._aaf -(_beg ._ceg -_bega );_beg ._ceg =_bega ;return _beg ._ceg ,nil ;};_beg ._aaf =0;_beg ._gb =_beg ._dc ;_beg ._ceg =_bega ;return _beg ._ceg ,_beg .seek (_bega );};