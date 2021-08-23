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

package zippkg ;import (_d "archive/zip";_a "bytes";_ee "encoding/xml";_dg "fmt";_db "github.com/unidoc/unioffice";_be "github.com/unidoc/unioffice/algo";_bc "github.com/unidoc/unioffice/common/tempstorage";_ff "github.com/unidoc/unioffice/schema/soo/pkg/relationships";_ea "io";_f "path";_b "sort";_df "strings";_e "time";);var _fcf =[]byte {'/','>'};var _gcf =[]byte {'\r','\n'};

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_bec map[string ]Target ;_aa map[*_ff .Relationships ]string ;_cf []Target ;_ag OnNewRelationshipFunc ;_cg map[string ]struct{};_bf map[string ]int ;};

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_d .Writer ,zipPath string ,data []byte )error {_aaa ,_bcd :=z .Create (zipPath );if _bcd !=nil {return _dg .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_bcd );};_ ,_bcd =_ea .Copy (_aaa ,_a .NewReader (data ));return _bcd ;};const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _ea .Writer ;};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};func MarshalXMLByType (z *_d .Writer ,dt _db .DocType ,typ string ,v interface{})error {_bed :=_db .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_bed ,v );};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_ga *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_ga ._ag =fn };func (_ffc SelfClosingWriter )Write (b []byte )(int ,error ){_egc :=0;_def :=0;for _ebf :=0;_ebf < len (b )-2;_ebf ++{if b [_ebf ]=='>'&&b [_ebf +1]=='<'&&b [_ebf +2]=='/'{_agd :=[]byte {};_daf :=_ebf ;for _dc :=_ebf ;_dc >=0;_dc --{if b [_dc ]==' '{_daf =_dc ;}else if b [_dc ]=='<'{_agd =b [_dc +1:_daf ];break ;};};_gaf :=[]byte {};for _bcb :=_ebf +3;_bcb < len (b );_bcb ++{if b [_bcb ]=='>'{_gaf =b [_ebf +3:_bcb ];break ;};};if !_a .Equal (_agd ,_gaf ){continue ;};_aadd ,_dgg :=_ffc .W .Write (b [_egc :_ebf ]);if _dgg !=nil {return _def +_aadd ,_dgg ;};_def +=_aadd ;_ ,_dgg =_ffc .W .Write (_fcf );if _dgg !=nil {return _def ,_dgg ;};_def +=3;for _defe :=_ebf +2;_defe < len (b )&&b [_defe ]!='>';_defe ++{_def ++;_egc =_defe +2;_ebf =_egc ;};};};_fea ,_faa :=_ffc .W .Write (b [_egc :]);return _fea +_def ,_faa ;};

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_d .File ,path string )(string ,error ){_caf ,_ddd :=_bc .TempFile (path ,"\u007a\u007a");if _ddd !=nil {return "",_ddd ;};defer _caf .Close ();_egg ,_ddd :=f .Open ();if _ddd !=nil {return "",_ddd ;};defer _egg .Close ();_ ,_ddd =_ea .Copy (_caf ,_egg );if _ddd !=nil {return "",_ddd ;};return _caf .Name (),nil ;};

// Decode loops decoding targets registered with AddTarget and calling th
func (_dgd *DecodeMap )Decode (files []*_d .File )error {_ge :=1;for _ge > 0{for len (_dgd ._cf )> 0{_cc :=_dgd ._cf [0];_dgd ._cf =_dgd ._cf [1:];_ef :=_cc .Ifc .(*_ff .Relationships );for _ ,_fc :=range _ef .Relationship {_fff ,_ :=_dgd ._aa [_ef ];_dgd ._ag (_dgd ,_fff +_fc .TargetAttr ,_fc .TypeAttr ,files ,_fc ,_cc );};};for _aege ,_dff :=range files {if _dff ==nil {continue ;};if _eg ,_dfe :=_dgd ._bec [_dff .Name ];_dfe {delete (_dgd ._bec ,_dff .Name );if _bb :=Decode (_dff ,_eg .Ifc );_bb !=nil {return _bb ;};files [_aege ]=nil ;if _dda ,_dgf :=_eg .Ifc .(*_ff .Relationships );_dgf {_dgd ._cf =append (_dgd ._cf ,_eg );_gd ,_ :=_f .Split (_f .Clean (_dff .Name +"\u002f\u002e\u002e\u002f"));_dgd ._aa [_dda ]=_gd ;_ge ++;};};};_ge --;};return nil ;};func (_af *DecodeMap )IndexFor (path string )int {return _af ._bf [path ]};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_d .Writer ,zipPath ,storagePath string )error {_bfa ,_fb :=z .Create (zipPath );if _fb !=nil {return _dg .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_fb );};_gad ,_fb :=_bc .Open (storagePath );if _fb !=nil {return _dg .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_fb );};defer _gad .Close ();_ ,_fb =_ea .Copy (_bfa ,_gad );return _fb ;};

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_d .File ,dest interface{})error {_fdc ,_eea :=f .Open ();if _eea !=nil {return _dg .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_eea );};defer _fdc .Close ();_gf :=_ee .NewDecoder (_fdc );if _afc :=_gf .Decode (dest );_afc !=nil {return _dg .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_afc );};if _cfa ,_gcb :=dest .(*_ff .Relationships );_gcb {for _afce ,_aeb :=range _cfa .Relationship {switch _aeb .TypeAttr {case _db .OfficeDocumentTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .OfficeDocumentType ;case _db .StylesTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .StylesType ;case _db .ThemeTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ThemeType ;case _db .ControlTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ControlType ;case _db .SettingsTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .SettingsType ;case _db .ImageTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ImageType ;case _db .CommentsTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .CommentsType ;case _db .ThumbnailTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ThumbnailType ;case _db .DrawingTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .DrawingType ;case _db .ChartTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ChartType ;case _db .ExtendedPropertiesTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .ExtendedPropertiesType ;case _db .CustomXMLTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .CustomXMLType ;case _db .WorksheetTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .WorksheetType ;case _db .SharedStringsTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .SharedStringsType ;case _db .TableTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .TableType ;case _db .HeaderTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .HeaderType ;case _db .FooterTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .FooterType ;case _db .NumberingTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .NumberingType ;case _db .FontTableTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .FontTableType ;case _db .WebSettingsTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .WebSettingsType ;case _db .FootNotesTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .FootNotesType ;case _db .EndNotesTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .EndNotesType ;case _db .SlideTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .SlideType ;case _db .VMLDrawingTypeStrict :_cfa .Relationship [_afce ].TypeAttr =_db .VMLDrawingType ;};};_b .Slice (_cfa .Relationship ,func (_fe ,_da int )bool {_fa :=_cfa .Relationship [_fe ];_bee :=_cfa .Relationship [_da ];return _be .NaturalLess (_fa .IdAttr ,_bee .IdAttr );});};return nil ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_d .Writer ,filename string ,v interface{})error {_aee :=&_d .FileHeader {};_aee .Method =_d .Deflate ;_aee .Name =filename ;_aee .SetModTime (_e .Now ());_gfd ,_cbe :=z .CreateHeader (_aee );if _cbe !=nil {return _dg .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_cbe );};_ ,_cbe =_gfd .Write ([]byte (XMLHeader ));if _cbe !=nil {return _dg .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_cbe );};if _cbe =_ee .NewEncoder (SelfClosingWriter {_gfd }).Encode (v );_cbe !=nil {return _dg .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_cbe );};_ ,_cbe =_gfd .Write (_gcf );return _cbe ;};

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ec *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _ec ._bec ==nil {_ec ._bec =make (map[string ]Target );_ec ._aa =make (map[*_ff .Relationships ]string );_ec ._cg =make (map[string ]struct{});_ec ._bf =make (map[string ]int );};_cfb :=_f .Clean (filePath );if _ ,_gb :=_ec ._cg [_cfb ];_gb {return false ;};_ec ._cg [_cfb ]=struct{}{};_ec ._bec [_cfb ]=Target {Path :filePath ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };return true ;};func (_aeg *DecodeMap )RecordIndex (path string ,idx int ){_aeg ._bf [path ]=idx };func MarshalXMLByTypeIndex (z *_d .Writer ,dt _db .DocType ,typ string ,idx int ,v interface{})error {_aad :=_db .AbsoluteFilename (dt ,typ ,idx );return MarshalXML (z ,_aad ,v );};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_ca :=_df .Split (path ,"\u002f");_gac :=_df .Join (_ca [0:len (_ca )-1],"\u002f");_gc :=_ca [len (_ca )-1];_gac +="\u002f_\u0072\u0065\u006c\u0073\u002f";_gc +="\u002e\u0072\u0065l\u0073";return _gac +_gc ;};

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_c *DecodeMap ,_ac ,_ae string ,_de []*_d .File ,_dd *_ff .Relationship ,_gg Target )error ;