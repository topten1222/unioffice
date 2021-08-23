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
package diskstore ;import (_a "github.com/unidoc/unioffice/common/tempstorage";_b "io/ioutil";_ba "os";_e "strings";);

// Add is not applicable in the diskstore implementation
func (_de diskStorage )Add (path string )error {return nil };

// RemoveAll removes all files in the directory
func (_f diskStorage )RemoveAll (dir string )error {if _e .HasPrefix (dir ,_ba .TempDir ()){return _ba .RemoveAll (dir );};return nil ;};type diskStorage struct{};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_bg diskStorage )TempDir (pattern string )(string ,error ){return _b .TempDir ("",pattern )};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_d :=diskStorage {};_a .SetAsStorage (&_d )};

// TempFile creates a new temp file by calling ioutil TempFile
func (_eg diskStorage )TempFile (dir ,pattern string )(_a .File ,error ){return _b .TempFile (dir ,pattern );};

// Open opens file from disk according to a path
func (_bf diskStorage )Open (path string )(_a .File ,error ){return _ba .Open (path )};