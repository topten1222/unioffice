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
package diskstore ;import (_eg "github.com/topten1222/unioffice/common/tempstorage";_c "io/ioutil";_a "os";_e "strings";);

// Open opens file from disk according to a path
func (_cb diskStorage )Open (path string )(_eg .File ,error ){return _a .OpenFile (path ,_a .O_RDWR ,0644)};type diskStorage struct{};

// TempFile creates a new temp file by calling ioutil TempFile
func (_d diskStorage )TempFile (dir ,pattern string )(_eg .File ,error ){return _c .TempFile (dir ,pattern );};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_fd :=diskStorage {};_eg .SetAsStorage (&_fd )};

// RemoveAll removes all files in the directory
func (_fe diskStorage )RemoveAll (dir string )error {if _e .HasPrefix (dir ,_a .TempDir ()){return _a .RemoveAll (dir );};return nil ;};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_fb diskStorage )TempDir (pattern string )(string ,error ){return _c .TempDir ("",pattern )};

// Add is not applicable in the diskstore implementation
func (_ed diskStorage )Add (path string )error {return nil };