// Package schema provides access to the Bridge API.
//
// See http://github.com/coreos-inc/bridge
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/schema/v1"
//   ...
//   schemaService, err := schema.New(oauthHttpClient)
package schema

import (
	"bytes"
	"github.com/coreos-inc/bridge/Godeps/_workspace/src/code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

const apiId = "bridge:v1"
const apiName = "schema"
const apiVersion = "v1"
const basePath = "http://localhost:9000/api/bridge/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Controllers = NewControllersService(s)
	s.Pods = NewPodsService(s)
	s.Services = NewServicesService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Controllers *ControllersService

	Pods *PodsService

	Services *ServicesService

	Users *UsersService
}

func NewControllersService(s *Service) *ControllersService {
	rs := &ControllersService{s: s}
	return rs
}

type ControllersService struct {
	s *Service
}

func NewPodsService(s *Service) *PodsService {
	rs := &PodsService{s: s}
	return rs
}

type PodsService struct {
	s *Service
}

func NewServicesService(s *Service) *ServicesService {
	rs := &ServicesService{s: s}
	return rs
}

type ServicesService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type Controller struct {
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// DesiredState: The desired configuration of the replicationController
	DesiredState *ControllerDesiredState `json:"desiredState,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Labels *ControllerLabels `json:"labels,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type ControllerDesiredState struct {
	// PodTemplate: Template from which to create new pods, as necessary.
	// Identical to pod schema.
	PodTemplate *ControllerDesiredStatePodTemplate `json:"podTemplate,omitempty"`

	// ReplicaSelector: Required labels used to identify pods in the set
	ReplicaSelector *ControllerDesiredStateReplicaSelector `json:"replicaSelector,omitempty"`

	// Replicas: Number of pods desired in the set
	Replicas float64 `json:"replicas,omitempty"`
}

type ControllerDesiredStatePodTemplate struct {
}

type ControllerDesiredStateReplicaSelector struct {
}

type ControllerLabels struct {
}

type ControllerList struct {
	Items []*Controller `json:"items,omitempty"`
}

type Pod struct {
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// CurrentState: The current configuration and status of the pod. Fields
	// in common with desiredState have the same meaning.
	CurrentState *PodCurrentState `json:"currentState,omitempty"`

	// DesiredState: The desired configuration of the pod
	DesiredState *PodDesiredState `json:"desiredState,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Labels *PodLabels `json:"labels,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type PodCurrentState struct {
	Host string `json:"host,omitempty"`

	HostIP string `json:"hostIP,omitempty"`

	Info *PodCurrentStateInfo `json:"info,omitempty"`

	Manifest *PodCurrentStateManifest `json:"manifest,omitempty"`

	Status string `json:"status,omitempty"`
}

type PodCurrentStateInfo struct {
}

type PodCurrentStateManifest struct {
}

type PodDesiredState struct {
	Host string `json:"host,omitempty"`

	HostIP string `json:"hostIP,omitempty"`

	Info *PodDesiredStateInfo `json:"info,omitempty"`

	// Manifest: Manifest describing group of [Docker
	// containers](http://docker.io); compatible with format used by [Google
	// Cloud Platform's container-vm
	// images](https://developers.google.com/compute/docs/containers)
	Manifest *PodDesiredStateManifest `json:"manifest,omitempty"`

	Status string `json:"status,omitempty"`
}

type PodDesiredStateInfo struct {
}

type PodDesiredStateManifest struct {
}

type PodLabels struct {
}

type PodList struct {
	Items []*Pod `json:"items,omitempty"`
}

type Service1 struct {
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Labels *ServiceLabels `json:"labels,omitempty"`

	Name string `json:"name,omitempty"`

	Port float64 `json:"port,omitempty"`

	Selector *ServiceSelector `json:"selector,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type ServiceLabels struct {
}

type ServiceSelector struct {
}

type ServiceList struct {
	Items []*Service1 `json:"items,omitempty"`
}

type User struct {
	FirstName string `json:"firstName,omitempty"`

	Id string `json:"id,omitempty"`

	LastName string `json:"lastName,omitempty"`
}

type UserPage struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Users []*User `json:"users,omitempty"`
}

// method id "bridge.controllers.get":

type ControllersGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Retrieve a Controllers.
func (r *ControllersService) Get(id string) *ControllersGetCall {
	c := &ControllersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Id sets the optional parameter "id":
func (c *ControllersGetCall) Id(id string) *ControllersGetCall {
	c.opt_["id"] = id
	return c
}

func (c *ControllersGetCall) Do() (*Controller, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "controllers/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{id}", url.QueryEscape(c.id), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Controller)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Controllers.",
	//   "httpMethod": "GET",
	//   "id": "bridge.controllers.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "type": "string"
	//     }
	//   },
	//   "path": "controllers/{id}",
	//   "response": {
	//     "$ref": "controller"
	//   }
	// }

}

// method id "bridge.controllers.list":

type ControllersListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve a list of Controllers.
func (r *ControllersService) List() *ControllersListCall {
	c := &ControllersListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *ControllersListCall) Do() (*ControllerList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "controllers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(ControllerList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a list of Controllers.",
	//   "httpMethod": "GET",
	//   "id": "bridge.controllers.list",
	//   "path": "controllers",
	//   "response": {
	//     "$ref": "controllerList"
	//   }
	// }

}

// method id "bridge.pods.get":

type PodsGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Retrieve a Pod.
func (r *PodsService) Get(id string) *PodsGetCall {
	c := &PodsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Id sets the optional parameter "id":
func (c *PodsGetCall) Id(id string) *PodsGetCall {
	c.opt_["id"] = id
	return c
}

func (c *PodsGetCall) Do() (*Pod, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "pods/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{id}", url.QueryEscape(c.id), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Pod)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Pod.",
	//   "httpMethod": "GET",
	//   "id": "bridge.pods.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "type": "string"
	//     }
	//   },
	//   "path": "pods/{id}",
	//   "response": {
	//     "$ref": "pod"
	//   }
	// }

}

// method id "bridge.pods.list":

type PodsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve a list of Pods.
func (r *PodsService) List() *PodsListCall {
	c := &PodsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *PodsListCall) Do() (*PodList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "pods")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(PodList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a list of Pods.",
	//   "httpMethod": "GET",
	//   "id": "bridge.pods.list",
	//   "path": "pods",
	//   "response": {
	//     "$ref": "podList"
	//   }
	// }

}

// method id "bridge.services.get":

type ServicesGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Retrieve a Service.
func (r *ServicesService) Get(id string) *ServicesGetCall {
	c := &ServicesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Id sets the optional parameter "id":
func (c *ServicesGetCall) Id(id string) *ServicesGetCall {
	c.opt_["id"] = id
	return c
}

func (c *ServicesGetCall) Do() (*Service1, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "services/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{id}", url.QueryEscape(c.id), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(Service1)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a Service.",
	//   "httpMethod": "GET",
	//   "id": "bridge.services.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "type": "string"
	//     }
	//   },
	//   "path": "services/{id}",
	//   "response": {
	//     "$ref": "service"
	//   }
	// }

}

// method id "bridge.services.list":

type ServicesListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve a list of Services.
func (r *ServicesService) List() *ServicesListCall {
	c := &ServicesListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

func (c *ServicesListCall) Do() (*ServiceList, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	urls := googleapi.ResolveRelative(c.s.BasePath, "services")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(ServiceList)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a list of Services.",
	//   "httpMethod": "GET",
	//   "id": "bridge.services.list",
	//   "path": "services",
	//   "response": {
	//     "$ref": "serviceList"
	//   }
	// }

}

// method id "bridge.user.destroy":

type UsersDestroyCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Destroy: Destroy a User.
func (r *UsersService) Destroy(id string) *UsersDestroyCall {
	c := &UsersDestroyCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Id sets the optional parameter "id":
func (c *UsersDestroyCall) Id(id string) *UsersDestroyCall {
	c.opt_["id"] = id
	return c
}

func (c *UsersDestroyCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{id}", url.QueryEscape(c.id), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Destroy a User.",
	//   "httpMethod": "DELETE",
	//   "id": "bridge.user.destroy",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}"
	// }

}

// method id "bridge.users.get":

type UsersGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Retrieve a User.
func (r *UsersService) Get(id string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Id sets the optional parameter "id":
func (c *UsersGetCall) Id(id string) *UsersGetCall {
	c.opt_["id"] = id
	return c
}

func (c *UsersGetCall) Do() (*User, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["id"]; ok {
		params.Set("id", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.URL.Path = strings.Replace(req.URL.Path, "{id}", url.QueryEscape(c.id), 1)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(User)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a User.",
	//   "httpMethod": "GET",
	//   "id": "bridge.users.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "location": "path",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{id}",
	//   "response": {
	//     "$ref": "user"
	//   }
	// }

}

// method id "bridge.users.list":

type UsersListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieve a page of Users.
func (r *UsersService) List() *UsersListCall {
	c := &UsersListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// NextPageToken sets the optional parameter "nextPageToken":
func (c *UsersListCall) NextPageToken(nextPageToken string) *UsersListCall {
	c.opt_["nextPageToken"] = nextPageToken
	return c
}

func (c *UsersListCall) Do() (*UserPage, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["nextPageToken"]; ok {
		params.Set("nextPageToken", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := new(UserPage)
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve a page of Users.",
	//   "httpMethod": "GET",
	//   "id": "bridge.users.list",
	//   "parameters": {
	//     "nextPageToken": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users",
	//   "response": {
	//     "$ref": "userPage"
	//   }
	// }

}