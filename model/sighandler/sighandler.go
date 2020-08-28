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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_dd "bytes";_ee "crypto";_ab "crypto/rand";_g "crypto/rsa";_ebe "crypto/x509";_ac "crypto/x509/pkix";_cb "encoding/asn1";_eb "errors";_f "fmt";_ga "github.com/unidoc/pkcs7";_b "github.com/unidoc/timestamp";_gf "github.com/unidoc/unipdf/v3/core";_eea "github.com/unidoc/unipdf/v3/model";_af "hash";_cc "io";_a "io/ioutil";_c "net/http";_e "time";);func (_eeab *docTimeStamp )getCertificate (_acce *_eea .PdfSignature )(*_ebe .Certificate ,error ){var _cbe []byte ;switch _ce :=_acce .Cert .(type ){case *_gf .PdfObjectString :_cbe =_ce .Bytes ();case *_gf .PdfObjectArray :if _ce .Len ()==0{return nil ,_eb .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_gca :=range _ce .Elements (){_cba ,_edd :=_gf .GetString (_gca );if !_edd {return nil ,_f .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_gca );};_cbe =append (_cbe ,_cba .Bytes ()...);};default:return nil ,_f .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_ce );};_gg ,_geb :=_ebe .ParseCertificates (_cbe );if _geb !=nil {return nil ,_geb ;};return _gg [0],nil ;};

// NewDigest creates a new digest.
func (_bd *adobePKCS7Detached )NewDigest (sig *_eea .PdfSignature )(_eea .Hasher ,error ){return _dd .NewBuffer (nil ),nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_beg *adobeX509RSASHA1 )IsApplicable (sig *_eea .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_bf *adobePKCS7Detached )IsApplicable (sig *_eea .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_g .PrivateKey ,certificate *_ebe .Certificate )(_eea .SignatureHandler ,error ){return &adobePKCS7Detached {_dg :certificate ,_ge :privateKey },nil ;};func (_ege *adobePKCS7Detached )getCertificate (_gff *_eea .PdfSignature )(*_ebe .Certificate ,error ){if _ege ._dg !=nil {return _ege ._dg ,nil ;};var _bb []byte ;switch _ff :=_gff .Cert .(type ){case *_gf .PdfObjectString :_bb =_ff .Bytes ();case *_gf .PdfObjectArray :if _ff .Len ()==0{return nil ,_eb .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_dge :=range _ff .Elements (){_dc ,_de :=_gf .GetString (_dge );if !_de {return nil ,_f .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_dge );};_bb =append (_bb ,_dc .Bytes ()...);};default:return nil ,_f .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_ff );};_ec ,_da :=_ebe .ParseCertificates (_bb );if _da !=nil {return nil ,_da ;};return _ec [0],nil ;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// The timestampServerURL parameter can be empty string for the signature validation.
// The hashAlgorithm parameter can be crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _ee .Hash )(_eea .SignatureHandler ,error ){return &docTimeStamp {_fgfb :timestampServerURL ,_deec :hashAlgorithm },nil ;};type adobeX509RSASHA1 struct{_acg *_g .PrivateKey ;_caa *_ebe .Certificate ;_fe SignFunc ;};

// Validate validates PdfSignature.
func (_be *adobeX509RSASHA1 )Validate (sig *_eea .PdfSignature ,digest _eea .Hasher )(_eea .SignatureValidationResult ,error ){_gde ,_dfc :=_be .getCertificate (sig );if _dfc !=nil {return _eea .SignatureValidationResult {},_dfc ;};_fgc :=sig .Contents .Bytes ();var _acc []byte ;if _ ,_add :=_cb .Unmarshal (_fgc ,&_acc );_add !=nil {return _eea .SignatureValidationResult {},_add ;};_dbc ,_daa :=digest .(_af .Hash );if !_daa {return _eea .SignatureValidationResult {},_eb .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_feg ,_ :=_bce (_gde .SignatureAlgorithm );if _bdc :=_g .VerifyPKCS1v15 (_gde .PublicKey .(*_g .PublicKey ),_feg ,_dbc .Sum (nil ),_acc );_bdc !=nil {return _eea .SignatureValidationResult {},_bdc ;};return _eea .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};type docTimeStamp struct{_fgfb string ;_deec _ee .Hash ;};

// Validate validates PdfSignature.
func (_deg *adobePKCS7Detached )Validate (sig *_eea .PdfSignature ,digest _eea .Hasher )(_eea .SignatureValidationResult ,error ){_db :=sig .Contents .Bytes ();_geg ,_dgf :=_ga .Parse (_db );if _dgf !=nil {return _eea .SignatureValidationResult {},_dgf ;};_dee :=digest .(*_dd .Buffer );_geg .Content =_dee .Bytes ();if _dgf =_geg .Verify ();_dgf !=nil {return _eea .SignatureValidationResult {},_dgf ;};return _eea .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// InitSignature initialises the PdfSignature.
func (_acec *adobeX509RSASHA1 )InitSignature (sig *_eea .PdfSignature )error {if _acec ._caa ==nil {return _eb .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _acec ._acg ==nil &&_acec ._fe ==nil {return _eb .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_ebb :=*_acec ;sig .Handler =&_ebb ;sig .Filter =_gf .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_gf .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");sig .Cert =_gf .MakeString (string (_ebb ._caa .Raw ));sig .Reference =nil ;_fa ,_fg :=_ebb .NewDigest (sig );if _fg !=nil {return _fg ;};_fa .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _ebb .Sign (sig ,_fa );};func (_ef *adobeX509RSASHA1 )getCertificate (_aa *_eea .PdfSignature )(*_ebe .Certificate ,error ){if _ef ._caa !=nil {return _ef ._caa ,nil ;};var _bbg []byte ;switch _fgf :=_aa .Cert .(type ){case *_gf .PdfObjectString :_bbg =_fgf .Bytes ();case *_gf .PdfObjectArray :if _fgf .Len ()==0{return nil ,_eb .New ("\u006e\u006f\u0020s\u0069\u0067\u006e\u0061t\u0075\u0072\u0065\u0020\u0063\u0065\u0072t\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0075\u006e\u0064");};for _ ,_cg :=range _fgf .Elements (){_fac ,_bbfa :=_gf .GetString (_cg );if !_bbfa {return nil ,_f .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062j\u0065\u0063\u0074\u0020\u0074\u0079p\u0065\u0020\u0069\u006e\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065 \u0063\u0065r\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u0063h\u0061\u0069\u006e\u003a\u0020\u0025\u0054",_cg );};_bbg =append (_bbg ,_fac .Bytes ()...);};default:return nil ,_f .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072e\u0020\u0063\u0065\u0072\u0074\u0069f\u0069\u0063\u0061\u0074\u0065\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0054",_fgf );};_ebbg ,_degf :=_ebe .ParseCertificates (_bbg );if _degf !=nil {return nil ,_degf ;};return _ebbg [0],nil ;};func _ecg (_aea _cb .ObjectIdentifier )(_ee .Hash ,error ){switch {case _aea .Equal (_ga .OIDDigestAlgorithmSHA1 ),_aea .Equal (_ga .OIDDigestAlgorithmECDSASHA1 ),_aea .Equal (_ga .OIDDigestAlgorithmDSA ),_aea .Equal (_ga .OIDDigestAlgorithmDSASHA1 ),_aea .Equal (_ga .OIDEncryptionAlgorithmRSA ):return _ee .SHA1 ,nil ;case _aea .Equal (_ga .OIDDigestAlgorithmSHA256 ),_aea .Equal (_ga .OIDDigestAlgorithmECDSASHA256 ):return _ee .SHA256 ,nil ;case _aea .Equal (_ga .OIDDigestAlgorithmSHA384 ),_aea .Equal (_ga .OIDDigestAlgorithmECDSASHA384 ):return _ee .SHA384 ,nil ;case _aea .Equal (_ga .OIDDigestAlgorithmSHA512 ),_aea .Equal (_ga .OIDDigestAlgorithmECDSASHA512 ):return _ee .SHA512 ,nil ;};return _ee .Hash (0),_ga .ErrUnsupportedAlgorithm ;};

// Validate validates PdfSignature.
func (_ba *docTimeStamp )Validate (sig *_eea .PdfSignature ,digest _eea .Hasher )(_eea .SignatureValidationResult ,error ){_bfe :=sig .Contents .Bytes ();_eec ,_edg :=_ga .Parse (_bfe );if _edg !=nil {return _eea .SignatureValidationResult {},_edg ;};if _edg =_eec .Verify ();_edg !=nil {return _eea .SignatureValidationResult {},_edg ;};var _fef timestampInfo ;_ ,_edg =_cb .Unmarshal (_eec .Content ,&_fef );if _edg !=nil {return _eea .SignatureValidationResult {},_edg ;};_dfb ,_edg :=_ecg (_fef .MessageImprint .HashAlgorithm .Algorithm );if _edg !=nil {return _eea .SignatureValidationResult {},_edg ;};_ag :=_dfb .New ();_fga :=digest .(*_dd .Buffer );_ag .Write (_fga .Bytes ());_gea :=_ag .Sum (nil );_edb :=_eea .SignatureValidationResult {IsSigned :true ,IsVerified :_dd .Equal (_gea ,_fef .MessageImprint .HashedMessage ),GeneralizedTime :_fef .GeneralizedTime };return _edb ,nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_abg *adobeX509RSASHA1 )Sign (sig *_eea .PdfSignature ,digest _eea .Hasher )error {var _gcb []byte ;var _gb error ;if _abg ._fe !=nil {_gcb ,_gb =_abg ._fe (sig ,digest );if _gb !=nil {return _gb ;};}else {_ede ,_cfc :=digest .(_af .Hash );if !_cfc {return _eb .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_ebbga ,_ :=_bce (_abg ._caa .SignatureAlgorithm );_gcb ,_gb =_g .SignPKCS1v15 (_ab .Reader ,_abg ._acg ,_ebbga ,_ede .Sum (nil ));if _gb !=nil {return _gb ;};};_gcb ,_gb =_cb .Marshal (_gcb );if _gb !=nil {return _gb ;};sig .Contents =_gf .MakeHexString (string (_gcb ));return nil ;};type timestampInfo struct{Version int ;Policy _cb .RawValue ;MessageImprint struct{HashAlgorithm _ac .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _cb .RawValue ;GeneralizedTime _e .Time ;};func _bce (_df _ebe .SignatureAlgorithm )(_ee .Hash ,bool ){return _ee .SHA1 ,true };

// Sign sets the Contents fields for the PdfSignature.
func (_edbc *docTimeStamp )Sign (sig *_eea .PdfSignature ,digest _eea .Hasher )error {_cbef :=digest .(*_dd .Buffer );_gbb :=_edbc ._deec .New ();if _ ,_fgca :=_cc .Copy (_gbb ,_cbef );_fgca !=nil {return _fgca ;};_gee :=_gbb .Sum (nil );_egg :=_b .Request {HashAlgorithm :_edbc ._deec ,HashedMessage :_gee ,Certificates :true ,Extensions :nil ,ExtraExtensions :nil };_cgd ,_dae :=_egg .Marshal ();if _dae !=nil {return _dae ;};_ada ,_dae :=_c .Post (_edbc ._fgfb ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_dd .NewBuffer (_cgd ));if _dae !=nil {return _dae ;};defer _ada .Body .Close ();_dcb ,_dae :=_a .ReadAll (_ada .Body );if _dae !=nil {return _dae ;};if _ada .StatusCode !=_c .StatusOK {return _f .Errorf ("\u0068\u0074\u0074\u0070\u0020\u0073\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064e\u0020n\u006f\u0074\u0020\u006f\u006b\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0064\u0029",_ada .StatusCode );};var _agg struct{Version _cb .RawValue ;Content _cb .RawValue ;};_ ,_dae =_cb .Unmarshal (_dcb ,&_agg );if _dae !=nil {return _dae ;};sig .Contents =_gf .MakeHexString (string (_agg .Content .FullBytes ));return nil ;};

// Sign sets the Contents fields.
func (_bbf *adobePKCS7Detached )Sign (sig *_eea .PdfSignature ,digest _eea .Hasher )error {if _bbf ._aff {_bbaf :=_bbf ._bc ;if _bbaf <=0{_bbaf =8192;};sig .Contents =_gf .MakeHexString (string (make ([]byte ,_bbaf )));return nil ;};_gd :=digest .(*_dd .Buffer );_gc ,_adg :=_ga .NewSignedData (_gd .Bytes ());if _adg !=nil {return _adg ;};if _cca :=_gc .AddSigner (_bbf ._dg ,_bbf ._ge ,_ga .SignerInfoConfig {});_cca !=nil {return _cca ;};_gc .Detach ();_eeg ,_adg :=_gc .Finish ();if _adg !=nil {return _adg ;};_ace :=make ([]byte ,8192);copy (_ace ,_eeg );sig .Contents =_gf .MakeHexString (string (_ace ));return nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler
// with a custom signing function. Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1Custom (certificate *_ebe .Certificate ,signFunc SignFunc )(_eea .SignatureHandler ,error ){return &adobeX509RSASHA1 {_caa :certificate ,_fe :signFunc },nil ;};

// InitSignature initialises the PdfSignature.
func (_ca *adobePKCS7Detached )InitSignature (sig *_eea .PdfSignature )error {if !_ca ._aff {if _ca ._dg ==nil {return _eb .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");};if _ca ._ge ==nil {return _eb .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_ad :=*_ca ;sig .Handler =&_ad ;sig .Filter =_gf .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_gf .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_eg ,_adf :=_ad .NewDigest (sig );if _adf !=nil {return _adf ;};_eg .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _ad .Sign (sig ,_eg );};

// NewDigest creates a new digest.
func (_acf *docTimeStamp )NewDigest (sig *_eea .PdfSignature )(_eea .Hasher ,error ){return _dd .NewBuffer (nil ),nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_bg *_eea .PdfSignature ,_bcf _eea .Hasher )([]byte ,error );type adobePKCS7Detached struct{_ge *_g .PrivateKey ;_dg *_ebe .Certificate ;_aff bool ;_bc int ;};

// NewDigest creates a new digest.
func (_fc *adobeX509RSASHA1 )NewDigest (sig *_eea .PdfSignature )(_eea .Hasher ,error ){_ffe ,_ed :=_fc .getCertificate (sig );if _ed !=nil {return nil ,_ed ;};_dab ,_ :=_bce (_ffe .SignatureAlgorithm );return _dab .New (),nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_cge *docTimeStamp )IsApplicable (sig *_eea .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite adbe.x509.rsa_sha1 signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_g .PrivateKey ,certificate *_ebe .Certificate )(_eea .SignatureHandler ,error ){return &adobeX509RSASHA1 {_caa :certificate ,_acg :privateKey },nil ;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_eea .SignatureHandler ,error ){return &adobePKCS7Detached {_aff :true ,_bc :signatureLen },nil ;};

// InitSignature initialises the PdfSignature.
func (_fcf *docTimeStamp )InitSignature (sig *_eea .PdfSignature )error {_fgb :=*_fcf ;sig .Handler =&_fgb ;sig .Filter =_gf .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_gf .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");sig .Reference =nil ;_gag ,_cag :=_fcf .NewDigest (sig );if _cag !=nil {return _cag ;};_gag .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));return _fgb .Sign (sig ,_gag );};