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

package endian ;import (_f "encoding/binary";_a "unsafe";);func IsLittle ()bool {return !_aa };func init (){const _b =int (_a .Sizeof (0));_ad :=1;_e :=(*[_b ]byte )(_a .Pointer (&_ad ));if _e [0]==0{_aa =true ;ByteOrder =_f .BigEndian ;}else {ByteOrder =_f .LittleEndian ;
};};var (ByteOrder _f .ByteOrder ;_aa bool ;);func IsBig ()bool {return _aa };