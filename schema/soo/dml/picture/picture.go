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

package picture ;import (_f "encoding/xml";_bg "github.com/topten1222/unioffice";_d "github.com/topten1222/unioffice/common/logger";_e "github.com/topten1222/unioffice/schema/soo/dml";);

// Validate validates the CT_Picture and its children
func (_ae *CT_Picture )Validate ()error {return _ae .ValidateWithPath ("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");};

// Validate validates the CT_PictureNonVisual and its children
func (_fa *CT_PictureNonVisual )Validate ()error {return _fa .ValidateWithPath ("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c");};func (_cg *CT_PictureNonVisual )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_cg .CNvPr =_e .NewCT_NonVisualDrawingProps ();
_cg .CNvPicPr =_e .NewCT_NonVisualPictureProperties ();_cf :for {_ffb ,_fgfb :=d .Token ();if _fgfb !=nil {return _fgfb ;};switch _cge :=_ffb .(type ){case _f .StartElement :switch _cge .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0063\u004e\u0076P\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0063\u004e\u0076P\u0072"}:if _ad :=d .DecodeElement (_cg .CNvPr ,&_cge );
_ad !=nil {return _ad ;};case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:if _ea :=d .DecodeElement (_cg .CNvPicPr ,&_cge );
_ea !=nil {return _ea ;};default:_d .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076",_cge .Name );
if _afg :=d .Skip ();_afg !=nil {return _afg ;};};case _f .EndElement :break _cf ;case _f .CharData :};};return nil ;};func NewCT_Picture ()*CT_Picture {_a :=&CT_Picture {};_a .NvPicPr =NewCT_PictureNonVisual ();_a .BlipFill =_e .NewCT_BlipFillProperties ();
_a .SpPr =_e .NewCT_ShapeProperties ();return _a ;};func (_ff *CT_Picture )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_ff .NvPicPr =NewCT_PictureNonVisual ();_ff .BlipFill =_e .NewCT_BlipFillProperties ();_ff .SpPr =_e .NewCT_ShapeProperties ();
_be :for {_bc ,_eg :=d .Token ();if _eg !=nil {return _eg ;};switch _fg :=_bc .(type ){case _f .StartElement :switch _fg .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"}:if _af :=d .DecodeElement (_ff .NvPicPr ,&_fg );
_af !=nil {return _af ;};case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:if _bd :=d .DecodeElement (_ff .BlipFill ,&_fg );
_bd !=nil {return _bd ;};case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"}:if _ec :=d .DecodeElement (_ff .SpPr ,&_fg );
_ec !=nil {return _ec ;};default:_d .Log .Debug ("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076",_fg .Name );
if _fb :=d .Skip ();_fb !=nil {return _fb ;};};case _f .EndElement :break _be ;case _f .CharData :};};return nil ;};type CT_PictureNonVisual struct{CNvPr *_e .CT_NonVisualDrawingProps ;CNvPicPr *_e .CT_NonVisualPictureProperties ;};func NewPic ()*Pic {_dg :=&Pic {};
_dg .CT_Picture =*NewCT_Picture ();return _dg };type CT_Picture struct{NvPicPr *CT_PictureNonVisual ;BlipFill *_e .CT_BlipFillProperties ;SpPr *_e .CT_ShapeProperties ;};func (_fgf *CT_PictureNonVisual )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );
_ce :=_f .StartElement {Name :_f .Name {Local :"\u0070i\u0063\u003a\u0063\u004e\u0076\u0050r"}};e .EncodeElement (_fgf .CNvPr ,_ce );_fcb :=_f .StartElement {Name :_f .Name {Local :"\u0070\u0069\u0063:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}};
e .EncodeElement (_fgf .CNvPicPr ,_fcb );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};func (_cfe *Pic )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_cfe .CT_Picture =*NewCT_Picture ();_gg :for {_bf ,_ecg :=d .Token ();
if _ecg !=nil {return _ecg ;};switch _eb :=_bf .(type ){case _f .StartElement :switch _eb .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"}:if _bef :=d .DecodeElement (_cfe .NvPicPr ,&_eb );
_bef !=nil {return _bef ;};case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:if _fgb :=d .DecodeElement (_cfe .BlipFill ,&_eb );
_fgb !=nil {return _fgb ;};case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"},_f .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"}:if _bfd :=d .DecodeElement (_cfe .SpPr ,&_eb );
_bfd !=nil {return _bfd ;};default:_d .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u0050i\u0063\u0020\u0025\u0076",_eb .Name );
if _de :=d .Skip ();_de !=nil {return _de ;};};case _f .EndElement :break _gg ;case _f .CharData :};};return nil ;};func (_c *CT_Picture )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );_ag :=_f .StartElement {Name :_f .Name {Local :"p\u0069\u0063\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}};
e .EncodeElement (_c .NvPicPr ,_ag );_fc :=_f .StartElement {Name :_f .Name {Local :"\u0070\u0069\u0063:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}};e .EncodeElement (_c .BlipFill ,_fc );_g :=_f .StartElement {Name :_f .Name {Local :"\u0070\u0069\u0063\u003a\u0073\u0070\u0050\u0072"}};
e .EncodeElement (_c .SpPr ,_g );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};type Pic struct{CT_Picture };

// Validate validates the Pic and its children
func (_fbag *Pic )Validate ()error {return _fbag .ValidateWithPath ("\u0050\u0069\u0063")};

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (_bb *Pic )ValidateWithPath (path string )error {if _cb :=_bb .CT_Picture .ValidateWithPath (path );_cb !=nil {return _cb ;};return nil ;};

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_fba *CT_PictureNonVisual )ValidateWithPath (path string )error {if _cc :=_fba .CNvPr .ValidateWithPath (path +"\u002f\u0043\u004e\u0076\u0050\u0072");_cc !=nil {return _cc ;};if _bgg :=_fba .CNvPicPr .ValidateWithPath (path +"\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r");
_bgg !=nil {return _bgg ;};return nil ;};

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_ed *CT_Picture )ValidateWithPath (path string )error {if _ba :=_ed .NvPicPr .ValidateWithPath (path +"\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072");_ba !=nil {return _ba ;};if _edd :=_ed .BlipFill .ValidateWithPath (path +"\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl");
_edd !=nil {return _edd ;};if _fga :=_ed .SpPr .ValidateWithPath (path +"\u002f\u0053\u0070P\u0072");_fga !=nil {return _fga ;};return nil ;};func (_cgea *Pic )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0070\u0069c"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0070i\u0063\u003a\u0070\u0069\u0063";return _cgea .CT_Picture .MarshalXML (e ,start );};func NewCT_PictureNonVisual ()*CT_PictureNonVisual {_fd :=&CT_PictureNonVisual {};_fd .CNvPr =_e .NewCT_NonVisualDrawingProps ();_fd .CNvPicPr =_e .NewCT_NonVisualPictureProperties ();
return _fd ;};func init (){_bg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c",NewCT_PictureNonVisual );
_bg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065",NewCT_Picture );
_bg .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0070\u0069\u0063",NewPic );
};