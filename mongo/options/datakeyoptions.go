// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package options

// DataKeyOptions represents all possible options used to create a new data key.
//
// See corresponding setter methods for documentation.
type DataKeyOptions struct {
	MasterKey   interface{}
	KeyAltNames []string
	KeyMaterial []byte
}

// DataKeyOptionsBuilder contains options to configure DataKey operations. Each
// option can be set through setter functions. See documentation for each setter
// function for an explanation of the option.
type DataKeyOptionsBuilder struct {
	Opts []func(*DataKeyOptions) error
}

// DataKey creates a new DataKeyOptions instance.
func DataKey() *DataKeyOptionsBuilder {
	return &DataKeyOptionsBuilder{}
}

// List returns a list of DataKey setter functions.
func (dk *DataKeyOptionsBuilder) List() []func(*DataKeyOptions) error {
	return dk.Opts
}

// SetMasterKey specifies a KMS-specific key used to encrypt the new data key.
//
// If being used with a local KMS provider, this option is not applicable and should not be specified.
//
// For the AWS, Azure, and GCP KMS providers, this option is required and must be a document. For each, the value of the
// "endpoint" or "keyVaultEndpoint" must be a host name with an optional port number (e.g. "foo.com" or "foo.com:443").
//
// When using AWS, the document must have the format:
//
//	{
//	  region: <string>,
//	  key: <string>,             // The Amazon Resource Name (ARN) to the AWS customer master key (CMK).
//	  endpoint: Optional<string> // An alternate host identifier to send KMS requests to.
//	}
//
// If unset, the "endpoint" defaults to "kms.<region>.amazonaws.com".
//
// When using Azure, the document must have the format:
//
//	{
//	  keyVaultEndpoint: <string>,  // A host identifier to send KMS requests to.
//	  keyName: <string>,
//	  keyVersion: Optional<string> // A specific version of the named key.
//	}
//
// If unset, "keyVersion" defaults to the key's primary version.
//
// When using GCP, the document must have the format:
//
//	{
//	  projectId: <string>,
//	  location: <string>,
//	  keyRing: <string>,
//	  keyName: <string>,
//	  keyVersion: Optional<string>, // A specific version of the named key.
//	  endpoint: Optional<string>    // An alternate host identifier to send KMS requests to.
//	}
//
// If unset, "keyVersion" defaults to the key's primary version and "endpoint" defaults to "cloudkms.googleapis.com".
func (dk *DataKeyOptionsBuilder) SetMasterKey(masterKey interface{}) *DataKeyOptionsBuilder {
	dk.Opts = append(dk.Opts, func(opts *DataKeyOptions) error {
		opts.MasterKey = masterKey

		return nil
	})

	return dk
}

// SetKeyAltNames specifies an optional list of string alternate names used to reference a key. If a key is created'
// with alternate names, encryption may refer to the key by a unique alternate name instead of by _id.
func (dk *DataKeyOptionsBuilder) SetKeyAltNames(keyAltNames []string) *DataKeyOptionsBuilder {
	dk.Opts = append(dk.Opts, func(opts *DataKeyOptions) error {
		opts.KeyAltNames = keyAltNames

		return nil
	})

	return dk
}

// SetKeyMaterial will set a custom keyMaterial to DataKeyOptions which can be used to encrypt data. If omitted,
// keyMaterial is generated form a cryptographically secure random source. "Key Material" is used interchangeably
// with "dataKey" and "Data Encryption Key" (DEK).
func (dk *DataKeyOptionsBuilder) SetKeyMaterial(keyMaterial []byte) *DataKeyOptionsBuilder {
	dk.Opts = append(dk.Opts, func(opts *DataKeyOptions) error {
		opts.KeyMaterial = keyMaterial

		return nil
	})

	return dk
}
