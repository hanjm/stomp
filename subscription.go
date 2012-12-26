package stomp

import (
	"github.com/jjeffery/stomp/message"
	"strconv"
)

// The Subscription type represents a client subscription to
// a destination. The subscription is created by calling Client.Subscribe.
//
// Once a client has subscribed, it can receive messages from the C channel.
type Subscription struct {
	//	C chan 
	id          int
	destination string
	client      *Client
	ackMode     AckMode
}

// BUG(jpj): If the client does not read messages from the Subscription.C channel quickly
// enough, the client will stop reading messages from the server.

// Identification for this subscription. Unique among
// all subscriptions for the same Client.
func (s *Subscription) Id() int {
	return s.id
}

// Destination for which the subscription applies.
func (s *Subscription) Destination() string {
	return s.destination
}

// The Ack mode for the subscription: auto, client or client-individual.
func (s *Subscription) AckMode() AckMode {
	return s.ackMode
}

// Unsubscribes and closes the channel C.
func (s *Subscription) Unsubscribe() error {
	_ = message.NewFrame(message.UNSUBSCRIBE, message.Id, strconv.Itoa(s.id))
	panic("not implemented")
}

// Read a message from the subscription
func (s *Subscription) Read() (*Message, error) {
	panic("not implemented")
}

// A Message is a message that is received from the server.
type Message struct {
	Destination  string        // Destination the message was sent to.
	ContentType  string        // MIME content
	Client       *Client       // Associated client
	Subscription *Subscription // Associated subscription
	Headers      Headers       // Optional headers
	Body         []byte        // Content of message
}

// The Headers interface represents a collection of headers, each having 
// a key  and a value. There may be more than one header in the collection 
// with the same key, in which case the first header's value is used.
type Headers interface {
	// Contains returns the value associated with the specified key, 
	// and whether it was found or not.
	Contains(key string) (string, bool)

	// Remove all headers with the specified key.
	Remove(key string)

	// Append the header to the end of the collection.
	Append(key, value string)

	// Set the value of the header. Replaces any existing header 
	// with the same key, or append if no header has the same key.
	Set(key, value string)

	// GetAt returns the header at the specified index.
	GetAt(index int) (key, value string)

	// Len returns the number of headers in the collection.
	Len() int
}
