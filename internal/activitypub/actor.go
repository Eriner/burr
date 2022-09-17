package activitypub

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/go-fed/activity/pub"
	"github.com/go-fed/activity/streams/vocab"
)

// var _ pub.FederatingActor = (*pub.FederatingActor)(nil)

var _ pub.CommonBehavior = &Service{}

type Service struct{}

// GetOutbox returns the OrderedCollection inbox of the actor for this
// context. It is up to the implementation to provide the correct
// collection for the kind of authorization given in the request.
//
// AuthenticateGetOutbox will be called prior to this.
//
// Always called, regardless whether the Federated Protocol or Social
// API is enabled.
func (s *Service) GetOutbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	panic("not implemented") // TODO: Implement
}

// PostInbox returns true if the request was handled as an ActivityPub
// POST to an actor's inbox. If false, the request was not an
// ActivityPub request and may still be handled by the caller in
// another way, such as serving a web page.
//
// If the error is nil, then the ResponseWriter's headers and response
// has already been written. If a non-nil error is returned, then no
// response has been written.
//
// If the Actor was constructed with the Federated Protocol enabled,
// side effects will occur.
//
// If the Federated Protocol is not enabled, writes the
// http.StatusMethodNotAllowed status code in the response. No side
// effects occur.
func (s *Service) PostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	panic("not implemented")
}

// PostOutbox returns true if the request was handled as an ActivityPub
// POST to an actor's outbox. If false, the request was not an
// ActivityPub request and may still be handled by the caller in another
// way, such as serving a web page.
//
// If the error is nil, then the ResponseWriter's headers and response
// has already been written. If a non-nil error is returned, then no
// response has been written.
//
// If the Actor was constructed with the Social Protocol enabled, side
// effects will occur.
//
// If the Social Protocol is not enabled, writes the
// http.StatusMethodNotAllowed status code in the response. No side
// effects occur.
//
// If the Social and Federated Protocol are both enabled, it will handle
// the side effects of receiving an ActivityStream Activity, and then
// federate the Activity to peers.
func (s *Service) PostOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
	panic("not implemented")
}

// GetOutbox returns true if the request was handled as an ActivityPub
// GET to an actor's outbox. If false, the request was not an
// ActivityPub request.
//
// If the error is nil, then the ResponseWriter's headers and response
// has already been written. If a non-nil error is returned, then no
// response has been written.
//
// If the request is an ActivityPub request, the Actor will defer to the
// application to determine the correct authorization of the request and
// the resulting OrderedCollection to respond with. The Actor handles
// serializing this OrderedCollection and responding with the correct
// headers and http.StatusOK.
//func (s *Service) GetOutbox(c context.Context, w http.ResponseWriter, r *http.Request) (bool, error) {
//	panic("not implemented")
//}

// Send a federated activity.
//
// The provided url must be the outbox of the sender. All processing of
// the activity occurs similarly to the C2S flow:
//   - If t is not an Activity, it is wrapped in a Create activity.
//   - A new ID is generated for the activity.
//   - The activity is added to the specified outbox.
//   - The activity is prepared and delivered to recipients.
//
// Note that this function will only behave as expected if the
// implementation has been constructed to support federation. This
// method will guaranteed work for non-custom Actors. For custom actors,
// care should be used to not call this method if only C2S is supported.
func (s *Service) Send(c context.Context, outbox *url.URL, t vocab.Type) (pub.Activity, error) {
	panic("not implemented") // TODO: Implement
}

// Hook callback after parsing the request body for a federated request
// to the Actor's inbox.
//
// Can be used to set contextual information based on the Activity
// received.
//
// Only called if the Federated Protocol is enabled.
//
// Warning: Neither authentication nor authorization has taken place at
// this time. Doing anything beyond setting contextual information is
// strongly discouraged.
//
// If an error is returned, it is passed back to the caller of
// PostInbox. In this case, the DelegateActor implementation must not
// write a response to the ResponseWriter as is expected that the caller
// to PostInbox will do so when handling the error.
func (s *Service) PostInboxRequestBodyHook(c context.Context, r *http.Request, activity pub.Activity) (context.Context, error) {
	panic("not implemented") // TODO: Implement
}

// AuthenticatePostInbox delegates the authentication of a POST to an
// inbox.
//
// If an error is returned, it is passed back to the caller of
// PostInbox. In this case, the implementation must not write a
// response to the ResponseWriter as is expected that the client will
// do so when handling the error. The 'authenticated' is ignored.
//
// If no error is returned, but authentication or authorization fails,
// then authenticated must be false and error nil. It is expected that
// the implementation handles writing to the ResponseWriter in this
// case.
//
// Finally, if the authentication and authorization succeeds, then
// authenticated must be true and error nil. The request will continue
// to be processed.
func (s *Service) AuthenticatePostInbox(c context.Context, w http.ResponseWriter, r *http.Request) (out context.Context, authenticated bool, err error) {
	panic("not implemented") // TODO: Implement
}

// Blocked should determine whether to permit a set of actors given by
// their ids are able to interact with this particular end user due to
// being blocked or other application-specific logic.
//
// If an error is returned, it is passed back to the caller of
// PostInbox.
//
// If no error is returned, but authentication or authorization fails,
// then blocked must be true and error nil. An http.StatusForbidden
// will be written in the wresponse.
//
// Finally, if the authentication and authorization succeeds, then
// blocked must be false and error nil. The request will continue
// to be processed.
func (s *Service) Blocked(c context.Context, actorIRIs []*url.URL) (blocked bool, err error) {
	panic("not implemented") // TODO: Implement
}

// FederatingCallbacks returns the application logic that handles
// ActivityStreams received from federating peers.
//
// Note that certain types of callbacks will be 'wrapped' with default
// behaviors supported natively by the library. Other callbacks
// compatible with streams.TypeResolver can be specified by 'other'.
//
// For example, setting the 'Create' field in the
// FederatingWrappedCallbacks lets an application dependency inject
// additional behaviors they want to take place, including the default
// behavior supplied by this library. This is guaranteed to be compliant
// with the ActivityPub Social protocol.
//
// To override the default behavior, instead supply the function in
// 'other', which does not guarantee the application will be compliant
// with the ActivityPub Social Protocol.
//
// Applications are not expected to handle every single ActivityStreams
// type and extension. The unhandled ones are passed to DefaultCallback.
func (s *Service) FederatingCallbacks(c context.Context) (wrapped pub.FederatingWrappedCallbacks, other []interface{}, err error) {
	panic("not implemented") // TODO: Implement
}

// DefaultCallback is called for types that go-fed can deserialize but
// are not handled by the application's callbacks returned in the
// Callbacks method.
//
// Applications are not expected to handle every single ActivityStreams
// type and extension, so the unhandled ones are passed to
// DefaultCallback.
func (s *Service) DefaultCallback(c context.Context, activity pub.Activity) error {
	panic("not implemented") // TODO: Implement
}

// MaxInboxForwardingRecursionDepth determines how deep to search within
// an activity to determine if inbox forwarding needs to occur.
//
// Zero or negative numbers indicate infinite recursion.
func (s *Service) MaxInboxForwardingRecursionDepth(c context.Context) int {
	panic("not implemented") // TODO: Implement
}

// MaxDeliveryRecursionDepth determines how deep to search within
// collections owned by peers when they are targeted to receive a
// delivery.
//
// Zero or negative numbers indicate infinite recursion.
func (s *Service) MaxDeliveryRecursionDepth(c context.Context) int {
	panic("not implemented") // TODO: Implement
}

// FilterForwarding allows the implementation to apply business logic
// such as blocks, spam filtering, and so on to a list of potential
// Collections and OrderedCollections of recipients when inbox
// forwarding has been triggered.
//
// The activity is provided as a reference for more intelligent
// logic to be used, but the implementation must not modify it.
func (s *Service) FilterForwarding(c context.Context, potentialRecipients []*url.URL, a pub.Activity) (filteredRecipients []*url.URL, err error) {
	panic("not implemented") // TODO: Implement
}

// GetInbox returns the OrderedCollection inbox of the actor for this
// context. It is up to the implementation to provide the correct
// collection for the kind of authorization given in the request.
//
// AuthenticateGetInbox will be called prior to this.
//
// Always called, regardless whether the Federated Protocol or Social
// API is enabled.
func (s *Service) GetInbox(c context.Context, r *http.Request) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	panic("not implemented") // TODO: Implement
}

func (*Service) AuthenticateGetInbox(ctx context.Context, w http.ResponseWriter, r *http.Request) (out context.Context, authed bool, err error) {
	panic("TODO")
}

func (*Service) AuthenticateGetOutbox(ctx context.Context, w http.ResponseWriter, r *http.Request) (out context.Context, authed bool, err error) {
	panic("TODO")
}

func (*Service) NewTransport(ctx context.Context, boxIRI *url.URL, goFedAgent string) (t pub.Transport, err error) {
	panic("TODO")
}

func (*Service) Now() time.Time {
	return time.Now()
}
