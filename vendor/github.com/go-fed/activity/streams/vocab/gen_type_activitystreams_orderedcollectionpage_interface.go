// Code generated by astool. DO NOT EDIT.

package vocab

// Used to represent ordered subsets of items from an OrderedCollection. Refer to
// the Activity Streams 2.0 Core for a complete description of the
// OrderedCollectionPage object.
//
// Example 8 (https://www.w3.org/TR/activitystreams-vocabulary/#ex6c-jsonld):
//   {
//     "id": "http://example.org/foo?page=1",
//     "orderedItems": [
//       {
//         "name": "A Simple Note",
//         "type": "Note"
//       },
//       {
//         "name": "Another Simple Note",
//         "type": "Note"
//       }
//     ],
//     "partOf": "http://example.org/foo",
//     "summary": "Page 1 of Sally's notes",
//     "type": "OrderedCollectionPage"
//   }
type ActivityStreamsOrderedCollectionPage interface {
	// GetActivityStreamsAltitude returns the "altitude" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAltitude() ActivityStreamsAltitudeProperty
	// GetActivityStreamsAttachment returns the "attachment" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAttachment() ActivityStreamsAttachmentProperty
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() ActivityStreamsAttributedToProperty
	// GetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAudience() ActivityStreamsAudienceProperty
	// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBcc() ActivityStreamsBccProperty
	// GetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBto() ActivityStreamsBtoProperty
	// GetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsCc() ActivityStreamsCcProperty
	// GetActivityStreamsContent returns the "content" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContent() ActivityStreamsContentProperty
	// GetActivityStreamsContext returns the "context" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContext() ActivityStreamsContextProperty
	// GetActivityStreamsCurrent returns the "current" property if it exists,
	// and nil otherwise.
	GetActivityStreamsCurrent() ActivityStreamsCurrentProperty
	// GetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDuration() ActivityStreamsDurationProperty
	// GetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	GetActivityStreamsEndTime() ActivityStreamsEndTimeProperty
	// GetActivityStreamsFirst returns the "first" property if it exists, and
	// nil otherwise.
	GetActivityStreamsFirst() ActivityStreamsFirstProperty
	// GetActivityStreamsGenerator returns the "generator" property if it
	// exists, and nil otherwise.
	GetActivityStreamsGenerator() ActivityStreamsGeneratorProperty
	// GetActivityStreamsIcon returns the "icon" property if it exists, and
	// nil otherwise.
	GetActivityStreamsIcon() ActivityStreamsIconProperty
	// GetActivityStreamsImage returns the "image" property if it exists, and
	// nil otherwise.
	GetActivityStreamsImage() ActivityStreamsImageProperty
	// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it
	// exists, and nil otherwise.
	GetActivityStreamsInReplyTo() ActivityStreamsInReplyToProperty
	// GetActivityStreamsLast returns the "last" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLast() ActivityStreamsLastProperty
	// GetActivityStreamsLikes returns the "likes" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLikes() ActivityStreamsLikesProperty
	// GetActivityStreamsLocation returns the "location" property if it
	// exists, and nil otherwise.
	GetActivityStreamsLocation() ActivityStreamsLocationProperty
	// GetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	GetActivityStreamsMediaType() ActivityStreamsMediaTypeProperty
	// GetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	GetActivityStreamsName() ActivityStreamsNameProperty
	// GetActivityStreamsNext returns the "next" property if it exists, and
	// nil otherwise.
	GetActivityStreamsNext() ActivityStreamsNextProperty
	// GetActivityStreamsObject returns the "object" property if it exists,
	// and nil otherwise.
	GetActivityStreamsObject() ActivityStreamsObjectProperty
	// GetActivityStreamsOrderedItems returns the "orderedItems" property if
	// it exists, and nil otherwise.
	GetActivityStreamsOrderedItems() ActivityStreamsOrderedItemsProperty
	// GetActivityStreamsPartOf returns the "partOf" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPartOf() ActivityStreamsPartOfProperty
	// GetActivityStreamsPrev returns the "prev" property if it exists, and
	// nil otherwise.
	GetActivityStreamsPrev() ActivityStreamsPrevProperty
	// GetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPreview() ActivityStreamsPreviewProperty
	// GetActivityStreamsPublished returns the "published" property if it
	// exists, and nil otherwise.
	GetActivityStreamsPublished() ActivityStreamsPublishedProperty
	// GetActivityStreamsReplies returns the "replies" property if it exists,
	// and nil otherwise.
	GetActivityStreamsReplies() ActivityStreamsRepliesProperty
	// GetActivityStreamsShares returns the "shares" property if it exists,
	// and nil otherwise.
	GetActivityStreamsShares() ActivityStreamsSharesProperty
	// GetActivityStreamsSource returns the "source" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSource() ActivityStreamsSourceProperty
	// GetActivityStreamsStartIndex returns the "startIndex" property if it
	// exists, and nil otherwise.
	GetActivityStreamsStartIndex() ActivityStreamsStartIndexProperty
	// GetActivityStreamsStartTime returns the "startTime" property if it
	// exists, and nil otherwise.
	GetActivityStreamsStartTime() ActivityStreamsStartTimeProperty
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTag() ActivityStreamsTagProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() ActivityStreamsToProperty
	// GetActivityStreamsTotalItems returns the "totalItems" property if it
	// exists, and nil otherwise.
	GetActivityStreamsTotalItems() ActivityStreamsTotalItemsProperty
	// GetActivityStreamsUpdated returns the "updated" property if it exists,
	// and nil otherwise.
	GetActivityStreamsUpdated() ActivityStreamsUpdatedProperty
	// GetActivityStreamsUrl returns the "url" property if it exists, and nil
	// otherwise.
	GetActivityStreamsUrl() ActivityStreamsUrlProperty
	// GetForgeFedEarlyItems returns the "earlyItems" property if it exists,
	// and nil otherwise.
	GetForgeFedEarlyItems() ForgeFedEarlyItemsProperty
	// GetForgeFedTeam returns the "team" property if it exists, and nil
	// otherwise.
	GetForgeFedTeam() ForgeFedTeamProperty
	// GetForgeFedTicketsTrackedBy returns the "ticketsTrackedBy" property if
	// it exists, and nil otherwise.
	GetForgeFedTicketsTrackedBy() ForgeFedTicketsTrackedByProperty
	// GetForgeFedTracksTicketsFor returns the "tracksTicketsFor" property if
	// it exists, and nil otherwise.
	GetForgeFedTracksTicketsFor() ForgeFedTracksTicketsForProperty
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetJSONLDType returns the "type" property if it exists, and nil
	// otherwise.
	GetJSONLDType() JSONLDTypeProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the
	// OrderedCollectionPage type. Note that this should not be used by
	// app developers. It is only used to help determine which
	// implementation is LessThan the other. Developers who are creating a
	// different implementation of this type's interface can use this
	// method in their LessThan implementation, but routine ActivityPub
	// applications should not use this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the OrderedCollectionPage type extends from
	// the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this OrderedCollectionPage is lesser, with an
	// arbitrary but stable determination.
	LessThan(o ActivityStreamsOrderedCollectionPage) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsAltitude sets the "altitude" property.
	SetActivityStreamsAltitude(i ActivityStreamsAltitudeProperty)
	// SetActivityStreamsAttachment sets the "attachment" property.
	SetActivityStreamsAttachment(i ActivityStreamsAttachmentProperty)
	// SetActivityStreamsAttributedTo sets the "attributedTo" property.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsAudience sets the "audience" property.
	SetActivityStreamsAudience(i ActivityStreamsAudienceProperty)
	// SetActivityStreamsBcc sets the "bcc" property.
	SetActivityStreamsBcc(i ActivityStreamsBccProperty)
	// SetActivityStreamsBto sets the "bto" property.
	SetActivityStreamsBto(i ActivityStreamsBtoProperty)
	// SetActivityStreamsCc sets the "cc" property.
	SetActivityStreamsCc(i ActivityStreamsCcProperty)
	// SetActivityStreamsContent sets the "content" property.
	SetActivityStreamsContent(i ActivityStreamsContentProperty)
	// SetActivityStreamsContext sets the "context" property.
	SetActivityStreamsContext(i ActivityStreamsContextProperty)
	// SetActivityStreamsCurrent sets the "current" property.
	SetActivityStreamsCurrent(i ActivityStreamsCurrentProperty)
	// SetActivityStreamsDuration sets the "duration" property.
	SetActivityStreamsDuration(i ActivityStreamsDurationProperty)
	// SetActivityStreamsEndTime sets the "endTime" property.
	SetActivityStreamsEndTime(i ActivityStreamsEndTimeProperty)
	// SetActivityStreamsFirst sets the "first" property.
	SetActivityStreamsFirst(i ActivityStreamsFirstProperty)
	// SetActivityStreamsGenerator sets the "generator" property.
	SetActivityStreamsGenerator(i ActivityStreamsGeneratorProperty)
	// SetActivityStreamsIcon sets the "icon" property.
	SetActivityStreamsIcon(i ActivityStreamsIconProperty)
	// SetActivityStreamsImage sets the "image" property.
	SetActivityStreamsImage(i ActivityStreamsImageProperty)
	// SetActivityStreamsInReplyTo sets the "inReplyTo" property.
	SetActivityStreamsInReplyTo(i ActivityStreamsInReplyToProperty)
	// SetActivityStreamsLast sets the "last" property.
	SetActivityStreamsLast(i ActivityStreamsLastProperty)
	// SetActivityStreamsLikes sets the "likes" property.
	SetActivityStreamsLikes(i ActivityStreamsLikesProperty)
	// SetActivityStreamsLocation sets the "location" property.
	SetActivityStreamsLocation(i ActivityStreamsLocationProperty)
	// SetActivityStreamsMediaType sets the "mediaType" property.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName sets the "name" property.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsNext sets the "next" property.
	SetActivityStreamsNext(i ActivityStreamsNextProperty)
	// SetActivityStreamsObject sets the "object" property.
	SetActivityStreamsObject(i ActivityStreamsObjectProperty)
	// SetActivityStreamsOrderedItems sets the "orderedItems" property.
	SetActivityStreamsOrderedItems(i ActivityStreamsOrderedItemsProperty)
	// SetActivityStreamsPartOf sets the "partOf" property.
	SetActivityStreamsPartOf(i ActivityStreamsPartOfProperty)
	// SetActivityStreamsPrev sets the "prev" property.
	SetActivityStreamsPrev(i ActivityStreamsPrevProperty)
	// SetActivityStreamsPreview sets the "preview" property.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsPublished sets the "published" property.
	SetActivityStreamsPublished(i ActivityStreamsPublishedProperty)
	// SetActivityStreamsReplies sets the "replies" property.
	SetActivityStreamsReplies(i ActivityStreamsRepliesProperty)
	// SetActivityStreamsShares sets the "shares" property.
	SetActivityStreamsShares(i ActivityStreamsSharesProperty)
	// SetActivityStreamsSource sets the "source" property.
	SetActivityStreamsSource(i ActivityStreamsSourceProperty)
	// SetActivityStreamsStartIndex sets the "startIndex" property.
	SetActivityStreamsStartIndex(i ActivityStreamsStartIndexProperty)
	// SetActivityStreamsStartTime sets the "startTime" property.
	SetActivityStreamsStartTime(i ActivityStreamsStartTimeProperty)
	// SetActivityStreamsSummary sets the "summary" property.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsTag sets the "tag" property.
	SetActivityStreamsTag(i ActivityStreamsTagProperty)
	// SetActivityStreamsTo sets the "to" property.
	SetActivityStreamsTo(i ActivityStreamsToProperty)
	// SetActivityStreamsTotalItems sets the "totalItems" property.
	SetActivityStreamsTotalItems(i ActivityStreamsTotalItemsProperty)
	// SetActivityStreamsUpdated sets the "updated" property.
	SetActivityStreamsUpdated(i ActivityStreamsUpdatedProperty)
	// SetActivityStreamsUrl sets the "url" property.
	SetActivityStreamsUrl(i ActivityStreamsUrlProperty)
	// SetForgeFedEarlyItems sets the "earlyItems" property.
	SetForgeFedEarlyItems(i ForgeFedEarlyItemsProperty)
	// SetForgeFedTeam sets the "team" property.
	SetForgeFedTeam(i ForgeFedTeamProperty)
	// SetForgeFedTicketsTrackedBy sets the "ticketsTrackedBy" property.
	SetForgeFedTicketsTrackedBy(i ForgeFedTicketsTrackedByProperty)
	// SetForgeFedTracksTicketsFor sets the "tracksTicketsFor" property.
	SetForgeFedTracksTicketsFor(i ForgeFedTracksTicketsForProperty)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// SetJSONLDType sets the "type" property.
	SetJSONLDType(i JSONLDTypeProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
