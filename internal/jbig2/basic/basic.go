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

package basic ;import _da "github.com/unidoc/unipdf/v3/internal/jbig2/errors";type Stack struct{Data []interface{};Aux *Stack ;};func (_gf IntSlice )Get (index int )(int ,error ){if index > len (_gf )-1{return 0,_da .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );};return _gf [index ],nil ;};func (_ac NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_ac )-1{return 0,_da .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};return _ac [i ],nil ;};func NewNumSlice (i int )*NumSlice {_ed :=NumSlice (make ([]float32 ,i ));return &_ed };func (_ee IntSlice )Size ()int {return len (_ee )};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_eeb NumSlice )GetIntSlice ()[]int {_cdd :=make ([]int ,len (_eeb ));for _gg ,_bb :=range _eeb {_cdd [_gg ]=int (_bb );};return _cdd ;};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_ad NumSlice )GetInt (i int )(int ,error ){const _fd ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_ad )-1{return 0,_da .Errorf (_fd ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_cb :=_ad [i ];return int (_cb +Sign (_cb )*0.5),nil ;};func (_dcb *Stack )Pop ()(_gd interface{},_gfa bool ){_gd ,_gfa =_dcb .peek ();if !_gfa {return nil ,_gfa ;};_dcb .Data =_dcb .Data [:_dcb .top ()];return _gd ,true ;};func (_cd *NumSlice )Add (v float32 ){*_cd =append (*_cd ,v )};func (_ged *Stack )Push (v interface{}){_ged .Data =append (_ged .Data ,v )};type IntsMap map[uint64 ][]int ;func (_gb IntsMap )Delete (key uint64 ){delete (_gb ,key )};func (_b *IntSlice )Copy ()*IntSlice {_dc :=IntSlice (make ([]int ,len (*_b )));copy (_dc ,*_b );return &_dc ;};func (_ab *Stack )top ()int {return len (_ab .Data )-1};func (_e IntsMap )Add (key uint64 ,value int ){_e [key ]=append (_e [key ],value )};func (_gbc *Stack )Len ()int {return len (_gbc .Data )};func Min (x ,y int )int {if x < y {return x ;};return y ;};func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_bd *Stack )peek ()(interface{},bool ){_cbg :=_bd .top ();if _cbg ==-1{return nil ,false ;};return _bd .Data [_cbg ],true ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_fc IntsMap )GetSlice (key uint64 )([]int ,bool ){_a ,_dad :=_fc [key ];if !_dad {return nil ,false ;};return _a ,true ;};type IntSlice []int ;func NewIntSlice (i int )*IntSlice {_dg :=IntSlice (make ([]int ,i ));return &_dg };func (_ff *IntSlice )Add (v int )error {if _ff ==nil {return _da .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");};*_ff =append (*_ff ,v );return nil ;};func (_fdd *Stack )Peek ()(_be interface{},_eee bool ){return _fdd .peek ()};func (_f IntsMap )Get (key uint64 )(int ,bool ){_g ,_c :=_f [key ];if !_c {return 0,false ;};if len (_g )==0{return 0,false ;};return _g [0],true ;};func (_ge *NumSlice )AddInt (v int ){*_ge =append (*_ge ,float32 (v ))};type NumSlice []float32 ;