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

package testutils ;import (_g "crypto/md5";_fg "encoding/hex";_e "errors";_cf "fmt";_bd "image";_ge "image/png";_b "io";_ga "os";_gf "os/exec";_ca "path/filepath";_c "strings";_d "testing";);func HashFile (file string )(string ,error ){_da ,_cab :=_ga .Open (file );if _cab !=nil {return "",_cab ;};defer _da .Close ();_fe :=_g .New ();if _ ,_cab =_b .Copy (_fe ,_da );_cab !=nil {return "",_cab ;};return _fg .EncodeToString (_fe .Sum (nil )),nil ;};func ReadPNG (file string )(_bd .Image ,error ){_fd ,_ag :=_ga .Open (file );if _ag !=nil {return nil ,_ag ;};defer _fd .Close ();return _ge .Decode (_fd );};var (ErrRenderNotSupported =_e .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m"););func CopyFile (src ,dst string )error {_bf ,_a :=_ga .Open (src );if _a !=nil {return _a ;};defer _bf .Close ();_bb ,_a :=_ga .Create (dst );if _a !=nil {return _a ;};defer _bb .Close ();_ ,_a =_b .Copy (_bb ,_bf );return _a ;};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_agd ,_fga :=HashFile (file1 );if _fga !=nil {return false ,_fga ;};_dd ,_fga :=HashFile (file2 );if _fga !=nil {return false ,_fga ;};if _agd ==_dd {return true ,nil ;};_fcde ,_fga :=ReadPNG (file1 );if _fga !=nil {return false ,_fga ;};_fa ,_fga :=ReadPNG (file2 );if _fga !=nil {return false ,_fga ;};if _fcde .Bounds ()!=_fa .Bounds (){return false ,nil ;};return CompareImages (_fcde ,_fa );};func CompareImages (img1 ,img2 _bd .Image )(bool ,error ){_fc :=img1 .Bounds ();_fb :=0;for _bg :=0;_bg < _fc .Size ().X ;_bg ++{for _aa :=0;_aa < _fc .Size ().Y ;_aa ++{_bc ,_bce ,_bbd ,_ :=img1 .At (_bg ,_aa ).RGBA ();_gb ,_fdd ,_bgg ,_ :=img2 .At (_bg ,_aa ).RGBA ();if _bc !=_gb ||_bce !=_fdd ||_bbd !=_bgg {_fb ++;};};};_fcd :=float64 (_fb )/float64 (_fc .Dx ()*_fc .Dy ());if _fcd > 0.0001{_cf .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_fcd ,_fb );return false ,nil ;};return true ,nil ;};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_cd :=_gf .LookPath ("\u0067\u0073");_cd !=nil {return ErrRenderNotSupported ;};return _gf .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_cf .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();};func RunRenderTest (t *_d .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_dg :=_c .TrimSuffix (_ca .Base (pdfPath ),_ca .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_aef *_d .T ){_gff :=_ca .Join (outputDir ,_dg );_bdb :=_gff +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _agc :=RenderPDFToPNGs (pdfPath ,0,_bdb );_agc !=nil {_aef .Skip (_agc );};for _bbg :=1;true ;_bbg ++{_afd :=_cf .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_gff ,_bbg );_dab :=_ca .Join (baselineRenderPath ,_cf .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_dg ,_bbg ));if _ ,_fba :=_ga .Stat (_afd );_fba !=nil {break ;};_aef .Logf ("\u0025\u0073",_dab );if saveBaseline {_aef .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_afd ,_dab );_ce :=CopyFile (_afd ,_dab );if _ce !=nil {_aef .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_dab ,_ce );};continue ;};_aef .Run (_cf .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_bbg ),func (_be *_d .T ){_be .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_afd ,_dab );_ef ,_dde :=ComparePNGFiles (_afd ,_dab );if _ga .IsNotExist (_dde ){_be .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_ef {_be .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");};});};});};