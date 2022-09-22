// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// ForgeFedFilesAddedPropertyIterator represents a single value for the
// "filesAdded" property.
type ForgeFedFilesAddedPropertyIterator interface {
	// Get returns the value of this property. When IsXMLSchemaString returns
	// false, Get will return any arbitrary value.
	Get() string
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// IsXMLSchemaString returns true if this property is set and not an IRI.
	IsXMLSchemaString() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ForgeFedFilesAddedPropertyIterator) bool
	// Name returns the name of this property: "ForgeFedFilesAdded".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ForgeFedFilesAddedPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ForgeFedFilesAddedPropertyIterator
	// Set sets the value of this property. Calling IsXMLSchemaString
	// afterwards will return true.
	Set(v string)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}

// Specifies a filename, as a relative path, relative to the top of the tree of
// files in the Repository, of a file that got added in this Commit, and
// didn’t exist in the previous version of the tree.
type ForgeFedFilesAddedProperty interface {
	// AppendIRI appends an IRI value to the back of a list of the property
	// "filesAdded"
	AppendIRI(v *url.URL)
	// AppendXMLSchemaString appends a string value to the back of a list of
	// the property "filesAdded". Invalidates iterators that are
	// traversing using Prev.
	AppendXMLSchemaString(v string)
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ForgeFedFilesAddedPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ForgeFedFilesAddedPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ForgeFedFilesAddedPropertyIterator
	// Insert inserts an IRI value at the specified index for a property
	// "filesAdded". Existing elements at that index and higher are
	// shifted back once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// InsertXMLSchemaString inserts a string value at the specified index for
	// a property "filesAdded". Existing elements at that index and higher
	// are shifted back once. Invalidates all iterators.
	InsertXMLSchemaString(idx int, v string)
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "filesAdded"
	// property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ForgeFedFilesAddedProperty) bool
	// Name returns the name of this property ("filesAdded") with any alias.
	Name() string
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "filesAdded".
	PrependIRI(v *url.URL)
	// PrependXMLSchemaString prepends a string value to the front of a list
	// of the property "filesAdded". Invalidates all iterators.
	PrependXMLSchemaString(v string)
	// Remove deletes an element at the specified index from a list of the
	// property "filesAdded", regardless of its type. Panics if the index
	// is out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets a string value to be at the specified index for the property
	// "filesAdded". Panics if the index is out of bounds. Invalidates all
	// iterators.
	Set(idx int, v string)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "filesAdded". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "filesAdded"
	// property.
	Swap(i, j int)
}
