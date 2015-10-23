package query_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/private/protocol/query"
	"github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"
	"github.com/aws/aws-sdk-go/private/signer/v4"
	"github.com/aws/aws-sdk-go/private/util"
	"github.com/stretchr/testify/assert"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = awstesting.GenerateAssertions
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF
var _ = aws.String

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService1ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService1ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService1ProtocolTest client from just a session.
//     svc := inputservice1protocoltest.New(mySession)
//
//     // Create a InputService1ProtocolTest client with additional configuration
//     svc := inputservice1protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService1ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService1ProtocolTest {
	c := p.ClientConfig("inputservice1protocoltest", cfgs...)
	return newInputService1ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService1ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService1ProtocolTest {
	svc := &InputService1ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice1protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService1ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService1ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService1TestCaseOperation1 = "OperationName"

// InputService1TestCaseOperation1Request generates a request for the InputService1TestCaseOperation1 operation.
func (c *InputService1ProtocolTest) InputService1TestCaseOperation1Request(input *InputService1TestShapeInputService1TestCaseOperation1Input) (req *request.Request, output *InputService1TestShapeInputService1TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService1TestCaseOperation1,
	}

	if input == nil {
		input = &InputService1TestShapeInputService1TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService1TestShapeInputService1TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService1ProtocolTest) InputService1TestCaseOperation1(input *InputService1TestShapeInputService1TestCaseOperation1Input) (*InputService1TestShapeInputService1TestCaseOperation1Output, error) {
	req, out := c.InputService1TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService1TestShapeInputService1TestCaseOperation1Input struct {
	Bar *string `type:"string"`

	Foo *string `type:"string"`

	metadataInputService1TestShapeInputService1TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService1TestShapeInputService1TestCaseOperation1Output struct {
	metadataInputService1TestShapeInputService1TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService1TestShapeInputService1TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService2ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService2ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService2ProtocolTest client from just a session.
//     svc := inputservice2protocoltest.New(mySession)
//
//     // Create a InputService2ProtocolTest client with additional configuration
//     svc := inputservice2protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService2ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService2ProtocolTest {
	c := p.ClientConfig("inputservice2protocoltest", cfgs...)
	return newInputService2ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService2ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService2ProtocolTest {
	svc := &InputService2ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice2protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService2ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService2ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService2TestCaseOperation1 = "OperationName"

// InputService2TestCaseOperation1Request generates a request for the InputService2TestCaseOperation1 operation.
func (c *InputService2ProtocolTest) InputService2TestCaseOperation1Request(input *InputService2TestShapeInputService2TestCaseOperation1Input) (req *request.Request, output *InputService2TestShapeInputService2TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService2TestCaseOperation1,
	}

	if input == nil {
		input = &InputService2TestShapeInputService2TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService2TestShapeInputService2TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService2ProtocolTest) InputService2TestCaseOperation1(input *InputService2TestShapeInputService2TestCaseOperation1Input) (*InputService2TestShapeInputService2TestCaseOperation1Output, error) {
	req, out := c.InputService2TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService2TestShapeInputService2TestCaseOperation1Input struct {
	StructArg *InputService2TestShapeStructType `type:"structure"`

	metadataInputService2TestShapeInputService2TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService2TestShapeInputService2TestCaseOperation1Output struct {
	metadataInputService2TestShapeInputService2TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService2TestShapeInputService2TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService2TestShapeStructType struct {
	ScalarArg *string `type:"string"`

	metadataInputService2TestShapeStructType `json:"-" xml:"-"`
}

type metadataInputService2TestShapeStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService3ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService3ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService3ProtocolTest client from just a session.
//     svc := inputservice3protocoltest.New(mySession)
//
//     // Create a InputService3ProtocolTest client with additional configuration
//     svc := inputservice3protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService3ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService3ProtocolTest {
	c := p.ClientConfig("inputservice3protocoltest", cfgs...)
	return newInputService3ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService3ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService3ProtocolTest {
	svc := &InputService3ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice3protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService3ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService3ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService3TestCaseOperation1 = "OperationName"

// InputService3TestCaseOperation1Request generates a request for the InputService3TestCaseOperation1 operation.
func (c *InputService3ProtocolTest) InputService3TestCaseOperation1Request(input *InputService3TestShapeInputService3TestCaseOperation1Input) (req *request.Request, output *InputService3TestShapeInputService3TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService3TestCaseOperation1,
	}

	if input == nil {
		input = &InputService3TestShapeInputService3TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService3TestShapeInputService3TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService3ProtocolTest) InputService3TestCaseOperation1(input *InputService3TestShapeInputService3TestCaseOperation1Input) (*InputService3TestShapeInputService3TestCaseOperation1Output, error) {
	req, out := c.InputService3TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService3TestShapeInputService3TestCaseOperation1Input struct {
	ListArg []*string `type:"list"`

	metadataInputService3TestShapeInputService3TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService3TestShapeInputService3TestCaseOperation1Output struct {
	metadataInputService3TestShapeInputService3TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService3TestShapeInputService3TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService4ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService4ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService4ProtocolTest client from just a session.
//     svc := inputservice4protocoltest.New(mySession)
//
//     // Create a InputService4ProtocolTest client with additional configuration
//     svc := inputservice4protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService4ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService4ProtocolTest {
	c := p.ClientConfig("inputservice4protocoltest", cfgs...)
	return newInputService4ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService4ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService4ProtocolTest {
	svc := &InputService4ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice4protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService4ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService4ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService4TestCaseOperation1 = "OperationName"

// InputService4TestCaseOperation1Request generates a request for the InputService4TestCaseOperation1 operation.
func (c *InputService4ProtocolTest) InputService4TestCaseOperation1Request(input *InputService4TestShapeInputService4TestCaseOperation1Input) (req *request.Request, output *InputService4TestShapeInputService4TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService4TestCaseOperation1,
	}

	if input == nil {
		input = &InputService4TestShapeInputService4TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService4TestShapeInputService4TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService4ProtocolTest) InputService4TestCaseOperation1(input *InputService4TestShapeInputService4TestCaseOperation1Input) (*InputService4TestShapeInputService4TestCaseOperation1Output, error) {
	req, out := c.InputService4TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService4TestShapeInputService4TestCaseOperation1Input struct {
	ListArg []*string `type:"list" flattened:"true"`

	ScalarArg *string `type:"string"`

	metadataInputService4TestShapeInputService4TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService4TestShapeInputService4TestCaseOperation1Output struct {
	metadataInputService4TestShapeInputService4TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService4TestShapeInputService4TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService5ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService5ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService5ProtocolTest client from just a session.
//     svc := inputservice5protocoltest.New(mySession)
//
//     // Create a InputService5ProtocolTest client with additional configuration
//     svc := inputservice5protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService5ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService5ProtocolTest {
	c := p.ClientConfig("inputservice5protocoltest", cfgs...)
	return newInputService5ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService5ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService5ProtocolTest {
	svc := &InputService5ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice5protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService5ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService5ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService5TestCaseOperation1 = "OperationName"

// InputService5TestCaseOperation1Request generates a request for the InputService5TestCaseOperation1 operation.
func (c *InputService5ProtocolTest) InputService5TestCaseOperation1Request(input *InputService5TestShapeInputService5TestCaseOperation1Input) (req *request.Request, output *InputService5TestShapeInputService5TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService5TestCaseOperation1,
	}

	if input == nil {
		input = &InputService5TestShapeInputService5TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService5TestShapeInputService5TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService5ProtocolTest) InputService5TestCaseOperation1(input *InputService5TestShapeInputService5TestCaseOperation1Input) (*InputService5TestShapeInputService5TestCaseOperation1Output, error) {
	req, out := c.InputService5TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService5TestShapeInputService5TestCaseOperation1Input struct {
	ListArg []*string `locationNameList:"ListArgLocation" type:"list" flattened:"true"`

	ScalarArg *string `type:"string"`

	metadataInputService5TestShapeInputService5TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService5TestShapeInputService5TestCaseOperation1Output struct {
	metadataInputService5TestShapeInputService5TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService5TestShapeInputService5TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService6ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService6ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService6ProtocolTest client from just a session.
//     svc := inputservice6protocoltest.New(mySession)
//
//     // Create a InputService6ProtocolTest client with additional configuration
//     svc := inputservice6protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService6ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService6ProtocolTest {
	c := p.ClientConfig("inputservice6protocoltest", cfgs...)
	return newInputService6ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService6ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService6ProtocolTest {
	svc := &InputService6ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice6protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService6ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService6ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService6TestCaseOperation1 = "OperationName"

// InputService6TestCaseOperation1Request generates a request for the InputService6TestCaseOperation1 operation.
func (c *InputService6ProtocolTest) InputService6TestCaseOperation1Request(input *InputService6TestShapeInputService6TestCaseOperation1Input) (req *request.Request, output *InputService6TestShapeInputService6TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService6TestCaseOperation1,
	}

	if input == nil {
		input = &InputService6TestShapeInputService6TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService6TestShapeInputService6TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService6ProtocolTest) InputService6TestCaseOperation1(input *InputService6TestShapeInputService6TestCaseOperation1Input) (*InputService6TestShapeInputService6TestCaseOperation1Output, error) {
	req, out := c.InputService6TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService6TestShapeInputService6TestCaseOperation1Input struct {
	MapArg map[string]*string `type:"map"`

	metadataInputService6TestShapeInputService6TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService6TestShapeInputService6TestCaseOperation1Output struct {
	metadataInputService6TestShapeInputService6TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService6TestShapeInputService6TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService7ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService7ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService7ProtocolTest client from just a session.
//     svc := inputservice7protocoltest.New(mySession)
//
//     // Create a InputService7ProtocolTest client with additional configuration
//     svc := inputservice7protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService7ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService7ProtocolTest {
	c := p.ClientConfig("inputservice7protocoltest", cfgs...)
	return newInputService7ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService7ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService7ProtocolTest {
	svc := &InputService7ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice7protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService7ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService7ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService7TestCaseOperation1 = "OperationName"

// InputService7TestCaseOperation1Request generates a request for the InputService7TestCaseOperation1 operation.
func (c *InputService7ProtocolTest) InputService7TestCaseOperation1Request(input *InputService7TestShapeInputService7TestCaseOperation1Input) (req *request.Request, output *InputService7TestShapeInputService7TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService7TestCaseOperation1,
	}

	if input == nil {
		input = &InputService7TestShapeInputService7TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService7TestShapeInputService7TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService7ProtocolTest) InputService7TestCaseOperation1(input *InputService7TestShapeInputService7TestCaseOperation1Input) (*InputService7TestShapeInputService7TestCaseOperation1Output, error) {
	req, out := c.InputService7TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService7TestShapeInputService7TestCaseOperation1Input struct {
	MapArg map[string]*string `locationNameKey:"TheKey" locationNameValue:"TheValue" type:"map"`

	metadataInputService7TestShapeInputService7TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService7TestShapeInputService7TestCaseOperation1Output struct {
	metadataInputService7TestShapeInputService7TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService7TestShapeInputService7TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService8ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService8ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService8ProtocolTest client from just a session.
//     svc := inputservice8protocoltest.New(mySession)
//
//     // Create a InputService8ProtocolTest client with additional configuration
//     svc := inputservice8protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService8ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService8ProtocolTest {
	c := p.ClientConfig("inputservice8protocoltest", cfgs...)
	return newInputService8ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService8ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService8ProtocolTest {
	svc := &InputService8ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice8protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService8ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService8ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService8TestCaseOperation1 = "OperationName"

// InputService8TestCaseOperation1Request generates a request for the InputService8TestCaseOperation1 operation.
func (c *InputService8ProtocolTest) InputService8TestCaseOperation1Request(input *InputService8TestShapeInputService8TestCaseOperation1Input) (req *request.Request, output *InputService8TestShapeInputService8TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService8TestCaseOperation1,
	}

	if input == nil {
		input = &InputService8TestShapeInputService8TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService8TestShapeInputService8TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService8ProtocolTest) InputService8TestCaseOperation1(input *InputService8TestShapeInputService8TestCaseOperation1Input) (*InputService8TestShapeInputService8TestCaseOperation1Output, error) {
	req, out := c.InputService8TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService8TestShapeInputService8TestCaseOperation1Input struct {
	BlobArg []byte `type:"blob"`

	metadataInputService8TestShapeInputService8TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService8TestShapeInputService8TestCaseOperation1Output struct {
	metadataInputService8TestShapeInputService8TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService8TestShapeInputService8TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService9ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService9ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService9ProtocolTest client from just a session.
//     svc := inputservice9protocoltest.New(mySession)
//
//     // Create a InputService9ProtocolTest client with additional configuration
//     svc := inputservice9protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService9ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService9ProtocolTest {
	c := p.ClientConfig("inputservice9protocoltest", cfgs...)
	return newInputService9ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService9ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService9ProtocolTest {
	svc := &InputService9ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice9protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService9ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService9ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService9TestCaseOperation1 = "OperationName"

// InputService9TestCaseOperation1Request generates a request for the InputService9TestCaseOperation1 operation.
func (c *InputService9ProtocolTest) InputService9TestCaseOperation1Request(input *InputService9TestShapeInputService9TestCaseOperation1Input) (req *request.Request, output *InputService9TestShapeInputService9TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService9TestCaseOperation1,
	}

	if input == nil {
		input = &InputService9TestShapeInputService9TestCaseOperation1Input{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService9TestShapeInputService9TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService9ProtocolTest) InputService9TestCaseOperation1(input *InputService9TestShapeInputService9TestCaseOperation1Input) (*InputService9TestShapeInputService9TestCaseOperation1Output, error) {
	req, out := c.InputService9TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

type InputService9TestShapeInputService9TestCaseOperation1Input struct {
	TimeArg *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataInputService9TestShapeInputService9TestCaseOperation1Input `json:"-" xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation1Input struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService9TestShapeInputService9TestCaseOperation1Output struct {
	metadataInputService9TestShapeInputService9TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService9TestShapeInputService9TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type InputService10ProtocolTest struct {
	*client.Client
}

// New creates a new instance of the InputService10ProtocolTest client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a InputService10ProtocolTest client from just a session.
//     svc := inputservice10protocoltest.New(mySession)
//
//     // Create a InputService10ProtocolTest client with additional configuration
//     svc := inputservice10protocoltest.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func NewInputService10ProtocolTest(p client.ConfigProvider, cfgs ...*aws.Config) *InputService10ProtocolTest {
	c := p.ClientConfig("inputservice10protocoltest", cfgs...)
	return newInputService10ProtocolTestClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newInputService10ProtocolTestClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *InputService10ProtocolTest {
	svc := &InputService10ProtocolTest{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   "inputservice10protocoltest",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	return svc
}

// newRequest creates a new request for a InputService10ProtocolTest operation and runs any
// custom request initialization.
func (c *InputService10ProtocolTest) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opInputService10TestCaseOperation1 = "OperationName"

// InputService10TestCaseOperation1Request generates a request for the InputService10TestCaseOperation1 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation1Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation1Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation1,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation1Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation1(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation1Output, error) {
	req, out := c.InputService10TestCaseOperation1Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation2 = "OperationName"

// InputService10TestCaseOperation2Request generates a request for the InputService10TestCaseOperation2 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation2Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation2Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation2,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation2Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation2(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation2Output, error) {
	req, out := c.InputService10TestCaseOperation2Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation3 = "OperationName"

// InputService10TestCaseOperation3Request generates a request for the InputService10TestCaseOperation3 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation3Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation3Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation3,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation3Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation3(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation3Output, error) {
	req, out := c.InputService10TestCaseOperation3Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation4 = "OperationName"

// InputService10TestCaseOperation4Request generates a request for the InputService10TestCaseOperation4 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation4Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation4Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation4,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation4Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation4(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation4Output, error) {
	req, out := c.InputService10TestCaseOperation4Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation5 = "OperationName"

// InputService10TestCaseOperation5Request generates a request for the InputService10TestCaseOperation5 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation5Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation5Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation5,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation5Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation5(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation5Output, error) {
	req, out := c.InputService10TestCaseOperation5Request(input)
	err := req.Send()
	return out, err
}

const opInputService10TestCaseOperation6 = "OperationName"

// InputService10TestCaseOperation6Request generates a request for the InputService10TestCaseOperation6 operation.
func (c *InputService10ProtocolTest) InputService10TestCaseOperation6Request(input *InputService10TestShapeInputShape) (req *request.Request, output *InputService10TestShapeInputService10TestCaseOperation6Output) {
	op := &request.Operation{
		Name: opInputService10TestCaseOperation6,
	}

	if input == nil {
		input = &InputService10TestShapeInputShape{}
	}

	req = c.newRequest(op, input, output)
	output = &InputService10TestShapeInputService10TestCaseOperation6Output{}
	req.Data = output
	return
}

func (c *InputService10ProtocolTest) InputService10TestCaseOperation6(input *InputService10TestShapeInputShape) (*InputService10TestShapeInputService10TestCaseOperation6Output, error) {
	req, out := c.InputService10TestCaseOperation6Request(input)
	err := req.Send()
	return out, err
}

type InputService10TestShapeInputService10TestCaseOperation1Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation1Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation1Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation2Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation2Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation2Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation3Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation3Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation3Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation4Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation4Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation4Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation5Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation5Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation5Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputService10TestCaseOperation6Output struct {
	metadataInputService10TestShapeInputService10TestCaseOperation6Output `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputService10TestCaseOperation6Output struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeInputShape struct {
	RecursiveStruct *InputService10TestShapeRecursiveStructType `type:"structure"`

	metadataInputService10TestShapeInputShape `json:"-" xml:"-"`
}

type metadataInputService10TestShapeInputShape struct {
	SDKShapeTraits bool `type:"structure"`
}

type InputService10TestShapeRecursiveStructType struct {
	NoRecurse *string `type:"string"`

	RecursiveList []*InputService10TestShapeRecursiveStructType `type:"list"`

	RecursiveMap map[string]*InputService10TestShapeRecursiveStructType `type:"map"`

	RecursiveStruct *InputService10TestShapeRecursiveStructType `type:"structure"`

	metadataInputService10TestShapeRecursiveStructType `json:"-" xml:"-"`
}

type metadataInputService10TestShapeRecursiveStructType struct {
	SDKShapeTraits bool `type:"structure"`
}

//
// Tests begin here
//

func TestInputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService1ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService1TestShapeInputService1TestCaseOperation1Input{
		Bar: aws.String("val2"),
		Foo: aws.String("val1"),
	}
	req, _ := svc.InputService1TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&Bar=val2&Foo=val1&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService2ProtocolTestNestedStructureMembersCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService2ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService2TestShapeInputService2TestCaseOperation1Input{
		StructArg: &InputService2TestShapeStructType{
			ScalarArg: aws.String("foo"),
		},
	}
	req, _ := svc.InputService2TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&StructArg.ScalarArg=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService3ProtocolTestListTypesCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService3ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService3TestShapeInputService3TestCaseOperation1Input{
		ListArg: []*string{
			aws.String("foo"),
			aws.String("bar"),
			aws.String("baz"),
		},
	}
	req, _ := svc.InputService3TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&ListArg.member.1=foo&ListArg.member.2=bar&ListArg.member.3=baz&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService4ProtocolTestFlattenedListCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService4ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService4TestShapeInputService4TestCaseOperation1Input{
		ListArg: []*string{
			aws.String("a"),
			aws.String("b"),
			aws.String("c"),
		},
		ScalarArg: aws.String("foo"),
	}
	req, _ := svc.InputService4TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&ListArg.1=a&ListArg.2=b&ListArg.3=c&ScalarArg=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService5ProtocolTestFlattenedListWithLocationNameCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService5ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService5TestShapeInputService5TestCaseOperation1Input{
		ListArg: []*string{
			aws.String("a"),
			aws.String("b"),
			aws.String("c"),
		},
		ScalarArg: aws.String("foo"),
	}
	req, _ := svc.InputService5TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&ListArgLocation.1=a&ListArgLocation.2=b&ListArgLocation.3=c&ScalarArg=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService6ProtocolTestSerializeMapTypeCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService6ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService6TestShapeInputService6TestCaseOperation1Input{
		MapArg: map[string]*string{
			"key1": aws.String("val1"),
			"key2": aws.String("val2"),
		},
	}
	req, _ := svc.InputService6TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&MapArg.entry.1.key=key1&MapArg.entry.1.value=val1&MapArg.entry.2.key=key2&MapArg.entry.2.value=val2&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService7ProtocolTestSerializeMapTypeWithLocationNameCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService7ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService7TestShapeInputService7TestCaseOperation1Input{
		MapArg: map[string]*string{
			"key1": aws.String("val1"),
			"key2": aws.String("val2"),
		},
	}
	req, _ := svc.InputService7TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&MapArg.entry.1.TheKey=key1&MapArg.entry.1.TheValue=val1&MapArg.entry.2.TheKey=key2&MapArg.entry.2.TheValue=val2&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService8ProtocolTestBase64EncodedBlobsCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService8ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService8TestShapeInputService8TestCaseOperation1Input{
		BlobArg: []byte("foo"),
	}
	req, _ := svc.InputService8TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&BlobArg=Zm9v&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService9ProtocolTestTimestampValuesCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService9ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService9TestShapeInputService9TestCaseOperation1Input{
		TimeArg: aws.Time(time.Unix(1422172800, 0)),
	}
	req, _ := svc.InputService9TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&TimeArg=2015-01-25T08%3A00%3A00Z&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase1(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			NoRecurse: aws.String("foo"),
		},
	}
	req, _ := svc.InputService10TestCaseOperation1Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.NoRecurse=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase2(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			RecursiveStruct: &InputService10TestShapeRecursiveStructType{
				NoRecurse: aws.String("foo"),
			},
		},
	}
	req, _ := svc.InputService10TestCaseOperation2Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.RecursiveStruct.NoRecurse=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase3(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			RecursiveStruct: &InputService10TestShapeRecursiveStructType{
				RecursiveStruct: &InputService10TestShapeRecursiveStructType{
					RecursiveStruct: &InputService10TestShapeRecursiveStructType{
						NoRecurse: aws.String("foo"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService10TestCaseOperation3Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.RecursiveStruct.RecursiveStruct.RecursiveStruct.NoRecurse=foo&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase4(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			RecursiveList: []*InputService10TestShapeRecursiveStructType{
				{
					NoRecurse: aws.String("foo"),
				},
				{
					NoRecurse: aws.String("bar"),
				},
			},
		},
	}
	req, _ := svc.InputService10TestCaseOperation4Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.RecursiveList.member.1.NoRecurse=foo&RecursiveStruct.RecursiveList.member.2.NoRecurse=bar&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase5(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			RecursiveList: []*InputService10TestShapeRecursiveStructType{
				{
					NoRecurse: aws.String("foo"),
				},
				{
					RecursiveStruct: &InputService10TestShapeRecursiveStructType{
						NoRecurse: aws.String("bar"),
					},
				},
			},
		},
	}
	req, _ := svc.InputService10TestCaseOperation5Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.RecursiveList.member.1.NoRecurse=foo&RecursiveStruct.RecursiveList.member.2.RecursiveStruct.NoRecurse=bar&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}

func TestInputService10ProtocolTestRecursiveShapesCase6(t *testing.T) {
	sess := session.New()
	svc := NewInputService10ProtocolTest(sess, &aws.Config{Endpoint: aws.String("https://test")})

	input := &InputService10TestShapeInputShape{
		RecursiveStruct: &InputService10TestShapeRecursiveStructType{
			RecursiveMap: map[string]*InputService10TestShapeRecursiveStructType{
				"bar": {
					NoRecurse: aws.String("bar"),
				},
				"foo": {
					NoRecurse: aws.String("foo"),
				},
			},
		},
	}
	req, _ := svc.InputService10TestCaseOperation6Request(input)
	r := req.HTTPRequest

	// build request
	query.Build(req)
	assert.NoError(t, req.Error)

	// assert body
	assert.NotNil(t, r.Body)
	body, _ := ioutil.ReadAll(r.Body)
	awstesting.AssertQuery(t, `Action=OperationName&RecursiveStruct.RecursiveMap.entry.1.key=foo&RecursiveStruct.RecursiveMap.entry.1.value.NoRecurse=foo&RecursiveStruct.RecursiveMap.entry.2.key=bar&RecursiveStruct.RecursiveMap.entry.2.value.NoRecurse=bar&Version=2014-01-01`, util.Trim(string(body)))

	// assert URL
	awstesting.AssertURL(t, "https://test/", r.URL.String())

	// assert headers

}
