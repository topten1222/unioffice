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

// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore ;import (_gb "github.com/unidoc/unioffice/common/tempstorage";_gg "io/ioutil";_b "os";_a "strings";);

// TempFile creates a new temp directory by calling ioutil TempDir
func (_ba diskStorage )TempDir (pattern string )(string ,error ){return _gg .TempDir ("",pattern )};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_e :=diskStorage {};_gb .SetAsStorage (&_e )};

// RemoveAll removes all files in the directory
func (_d diskStorage )RemoveAll (dir string )error {if _a .HasPrefix (dir ,_b .TempDir ()){return _b .RemoveAll (dir );};return nil ;};

// Add is not applicable in the diskstore implementation
func (_ee diskStorage )Add (path string )error {return nil };type diskStorage struct{};

// Open opens file from disk according to a path
func (_bd diskStorage )Open (path string )(_gb .File ,error ){return _b .OpenFile (path ,_b .O_RDWR ,0644)};

// TempFile creates a new temp file by calling ioutil TempFile
func (_ec diskStorage )TempFile (dir ,pattern string )(_gb .File ,error ){return _gg .TempFile (dir ,pattern );};