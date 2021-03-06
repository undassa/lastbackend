//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2020] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package service

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lastbackend/lastbackend/internal/server/server/middleware"
	h "github.com/lastbackend/lastbackend/internal/util/http"
	"github.com/lastbackend/lastbackend/tools/logger"
	"net/http"
)

const (
	logPrefix  = "api:handler:service"
	BufferSize = 512
)

// Handler represent the http handler for service
type Handler struct {
	Config Config
}

type Config struct {
	SecretToken string
}

// NewServiceHandler will initialize the service resources endpoint
func NewServiceHandler(r *mux.Router, mw middleware.Middleware, cfg Config) {

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	log.Infof("%s:> init service routes", logPrefix)

	handler := &Handler{
		Config: cfg,
	}

	r.Handle("/namespace/{namespace}/service", h.Handle(mw.Authenticate(handler.ServiceCreateH))).Methods(http.MethodPost)
	r.Handle("/namespace/{namespace}/service", h.Handle(mw.Authenticate(handler.ServiceListH))).Methods(http.MethodGet)
	r.Handle("/namespace/{namespace}/service/{service}", h.Handle(mw.Authenticate(handler.ServiceInfoH))).Methods(http.MethodGet)
	r.Handle("/namespace/{namespace}/service/{service}", h.Handle(mw.Authenticate(handler.ServiceUpdateH))).Methods(http.MethodPut)
	r.Handle("/namespace/{namespace}/service/{service}", h.Handle(mw.Authenticate(handler.ServiceRemoveH))).Methods(http.MethodDelete)
	r.Handle("/namespace/{namespace}/service/{service}/logs", h.Handle(mw.Authenticate(handler.ServiceLogsH))).Methods(http.MethodGet)
}

func (handler Handler) ServiceListH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /namespace/{namespace}/service service serviceList
	//
	// Shows a list of services
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	// responses:
	//   '200':
	//     description: Applications list response
	//     schema:
	//       "$ref": "#/definitions/views_service_list"
	//   '404':
	//     description: Namespace not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	//nid := util.Vars(r)["namespace"]
	//
	//log.Debugf("%s:list:> list services in %s", logPrefix, nid)
	//
	//var (
	//	stg = envs.Get().GetStorage()
	//	sm  = model.NewServiceModel(r.Context(), stg)
	//	nsm = model.NewNamespaceModel(r.Context(), stg)
	//)
	//
	//ns, err := nsm.Get(nid)
	//if err != nil {
	//	log.Errorf("%s:list:> get namespace", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if ns == nil {
	//	err := errors.New("namespace not found")
	//	log.Errorf("%s:list:> get namespace", logPrefix, err.Error())
	//	errors.New("namespace").NotFound().Http(w)
	//	return
	//}
	//
	//items, err := sm.List(ns.Meta.Name)
	//if err != nil {
	//	log.Errorf("%s:list:> get service list in namespace `%s` err: %s", logPrefix, ns.Meta.Name, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//
	//response, err := v1.View().Service().NewList(items).ToJson()
	//if err != nil {
	//	log.Errorf("%s:list:> convert struct to json err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}

	response := []byte{}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Errorf("%s:list:> write response err: %s", logPrefix, err.Error())
		return
	}
}

func (handler Handler) ServiceInfoH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /namespace/{namespace}/service/{service} service serviceInfo
	//
	// Shows an info about service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	//   - name: service
	//     in: path
	//     description: service id
	//     required: true
	//     type: string
	// responses:
	//   '200':
	//     description: Applications list response
	//     schema:
	//       "$ref": "#/definitions/views_service"
	//   '404':
	//     description: Namespace not found / Service not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	//sid := util.Vars(r)["service"]
	//nid := util.Vars(r)["namespace"]
	//
	//log.Debugf("%s:info:> get service `%s` in namespace `%s`", logPrefix, sid, nid)
	//
	//ns, e := namespace.FetchFromRequest(r.Context(), nid)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//svc, e := service.Fetch(r.Context(), ns.Meta.Name, sid)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//response, err := v1.View().Service().NewWithDeployment(svc).ToJson()
	//if err != nil {
	//	log.Errorf("%s:info:> convert struct to json err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}

	response := []byte{}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Errorf("%s:get write response err: %s", logPrefix, err.Error())
		return
	}
}

func (handler Handler) ServiceCreateH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation POST /namespace/{namespace}/service service serviceCreate
	//
	// Create new service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	//   - name: body
	//     in: body
	//     required: true
	//     schema:
	//       "$ref": "#/definitions/request_service_create"
	// responses:
	//   '200':
	//     description: Applications was successfully created
	//     schema:
	//       "$ref": "#/definitions/views_service"
	//   '400':
	//     description: Name is already in use
	//   '404':
	//     description: Namespace not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	//nid := util.Vars(r)["namespace"]
	//
	//log.Debugf("%s:create:> create service in namespace `%s`", logPrefix, nid)
	//
	//var (
	//	opts = v1.Request().Service().Manifest()
	//)
	//
	//// request body struct
	//if err := opts.DecodeAndValidate(r.Body); err != nil {
	//	log.Errorf("%s:create:> validation incoming data err: %s", logPrefix, err.Err())
	//	err.Http(w)
	//	return
	//}
	//
	//ns, e := namespace.FetchFromRequest(r.Context(), nid)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//svc, e := service.Create(r.Context(), ns, opts)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//response, err := v1.View().Service().NewWithDeployment(svc).ToJson()
	//if err != nil {
	//	log.Errorf("%s:update:> convert struct to json err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}

	response := []byte{}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Errorf("%s:update:> write response err: %s", logPrefix, err.Error())
		return
	}
}

func (handler Handler) ServiceUpdateH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation PUT /namespace/{namespace}/service/{service} service serviceUpdate
	//
	// Update service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	//   - name: service
	//     in: path
	//     description: service id
	//     required: true
	//     type: string
	//   - name: body
	//     in: body
	//     required: true
	//     schema:
	//       "$ref": "#/definitions/request_service_update"
	// responses:
	//   '200':
	//     description: Applications was successfully updated
	//     schema:
	//       "$ref": "#/definitions/views_service"
	//   '404':
	//     description: Namespace not found / Service not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	//nid := util.Vars(r)["namespace"]
	//sid := util.Vars(r)["service"]
	//
	//redeploy := util.QueryBool(r, "redeploy")
	//
	//log.Debugf("%s:update:> update service `%s` in namespace `%s`", logPrefix, sid, nid)
	//
	//// request body struct
	//opts := v1.Request().Service().Manifest()
	//if e := opts.DecodeAndValidate(r.Body); e != nil {
	//	log.Errorf("%s:update:> validation incoming data err: %s", logPrefix, e.Err())
	//	e.Http(w)
	//	return
	//}
	//
	//ns, e := namespace.FetchFromRequest(r.Context(), nid)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//svc, e := service.Fetch(r.Context(), ns.Meta.Name, sid)
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//svc, e = service.Update(r.Context(), ns, svc, opts, &request.ServiceUpdateOptions{Redeploy: redeploy})
	//if e != nil {
	//	e.Http(w)
	//	return
	//}
	//
	//response, err := v1.View().Service().NewWithDeployment(svc).ToJson()
	//if err != nil {
	//	log.Errorf("%s:update:> convert struct to json err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}

	response := []byte{}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		log.Errorf("%s:update:> write response err: %s", logPrefix, err.Error())
		return
	}
}

func (handler Handler) ServiceRemoveH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation DELETE /namespace/{namespace}/service/{service} service serviceRemove
	//
	// Remove service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	//   - name: service
	//     in: path
	//     description: service id
	//     required: true
	//     type: string
	// responses:
	//   '200':
	//     description: Applications was successfully removed
	//   '404':
	//     description: Namespace not found / Service not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	//nid := util.Vars(r)["namespace"]
	//sid := util.Vars(r)["service"]
	//
	//log.Debugf("%s:remove:> remove service `%s` from app `%s`", logPrefix, sid, nid)
	//
	//var (
	//	stg = envs.Get().GetStorage()
	//	nsm = model.NewNamespaceModel(r.Context(), stg)
	//	sm  = model.NewServiceModel(r.Context(), stg)
	//	rm  = model.NewRouteModel(r.Context(), stg)
	//)
	//
	//ns, err := nsm.Get(nid)
	//if err != nil {
	//	log.Errorf("%s:remove:> get namespace", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if ns == nil {
	//	err := errors.New("namespace not found")
	//	log.Errorf("%s:remove:> get namespace", logPrefix, err.Error())
	//	errors.New("namespace").NotFound().Http(w)
	//	return
	//}
	//
	//svc, err := sm.Get(ns.Meta.Name, sid)
	//if err != nil {
	//	log.Errorf("%s:remove:> get service by name `%s` in namespace `%s` err: %s", logPrefix, sid, ns.Meta.Name, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if svc == nil {
	//	log.Warnf("%s:remove:> service name `%s` in namespace `%s` not found", logPrefix, sid, ns.Meta.Name)
	//	errors.New("service").NotFound().Http(w)
	//	return
	//}
	//
	//rl, err := rm.ListByNamespace(nid)
	//if err != nil {
	//	log.Errorf("%s:remove:> get routes list in namespace `%s` err: %s", logPrefix, ns.Meta.Name, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//
	//// check routes attached to routes
	//for _, r := range rl.Items {
	//	for _, rule := range r.Spec.Rules {
	//		if rule.Service == svc.Meta.Name {
	//			log.Errorf("%s:remove:> service used in route `%s`", logPrefix, r.Meta.Name)
	//			errors.HTTP.BadRequest(w, errors.New(r.Meta.Name).Service().RouteBinded(r.Meta.Name).Error())
	//			return
	//		}
	//	}
	//}
	//
	//if _, err := sm.Destroy(svc); err != nil {
	//	log.Errorf("%s:remove:> remove service err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte{}); err != nil {
		log.Errorf("%s:remove:> write response err: %s", logPrefix, err.Error())
		return
	}
}

func (handler Handler) ServiceLogsH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /namespace/{namespace}/service/{service}/logs service serviceLogs
	//
	// Shows logs of the service
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	//   - name: namespace
	//     in: path
	//     description: namespace id
	//     required: true
	//     type: string
	//   - name: service
	//     in: path
	//     description: service id
	//     required: true
	//     type: string
	//   - name: deployment
	//     in: query
	//     description: deployment id
	//     required: true
	//     type: string
	//   - name: pod
	//     in: query
	//     description: pod id
	//     required: true
	//     type: string
	//   - name: container
	//     in: query
	//     description: container id
	//     required: true
	//     type: string
	// responses:
	//   '200':
	//     description: Applications logs received
	//   '404':
	//     description: Namespace not found / Service not found
	//   '500':
	//     description: Internal server error

	ctx := logger.NewContext(context.Background(), nil)
	log := logger.WithContext(ctx)

	log.Infof("%s:> get service logs", logPrefix)

	//nid := util.Vars(r)["namespace"]
	//sid := util.Vars(r)["service"]
	//tail := util.QueryInt(r, "tail")
	//follow := util.QueryBool(r, "follow")
	//
	//log.Debugf("%s:logs:> get logs for service `%s` in namespace `%s`", logPrefix, sid, nid)
	//
	//var (
	//	nsm = model.NewNamespaceModel(r.Context(), envs.Get().GetStorage())
	//	sm  = model.NewServiceModel(r.Context(), envs.Get().GetStorage())
	//	em  = model.NewExporterModel(r.Context(), envs.Get().GetStorage())
	//)
	//
	//ns, err := nsm.Get(nid)
	//if err != nil {
	//	log.Errorf("%s:logs:> get namespace", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if ns == nil {
	//	err := errors.New("namespace not found")
	//	log.Errorf("%s:logs:> get namespace", logPrefix, err.Error())
	//	errors.New("namespace").NotFound().Http(w)
	//	return
	//}
	//
	//svc, err := sm.Get(ns.Meta.Name, sid)
	//if err != nil {
	//	log.Errorf("%s:logs:> get service by name `%s` err: %s", logPrefix, sid, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if svc == nil {
	//	log.Warnf("%s:logs:> service name `%s` in namespace `%s` not found", logPrefix, sid, ns.Meta.Name)
	//	errors.New("service").NotFound().Http(w)
	//	return
	//}
	//
	//el, err := em.List()
	//if err != nil {
	//	log.Errorf("%s:logs:> get exporters", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//if len(el.Items) == 0 {
	//	log.Errorf("%s:logs:> exporters not found", logPrefix)
	//	errors.HTTP.NotFound(w)
	//	return
	//}
	//
	//exp := new(types.Exporter)
	//
	//for _, e := range el.Items {
	//	if e.Status.Ready {
	//		exp = e
	//		break
	//	}
	//}
	//
	//if exp == nil {
	//	log.Errorf("%s:logs:> active exporters not found", logPrefix)
	//	errors.HTTP.NotFound(w)
	//	return
	//}
	//
	//cx, cancel := context.WithCancel(context.Background())
	//
	//flw := "true"
	//if !follow {
	//	flw = "false"
	//}
	//
	//url := fmt.Sprintf("http://%s:%d/logs?kind=%s&selflink=%s&lines=%d&follow=%s",
	//	exp.Status.Http.IP, exp.Status.Http.Port, types.KindService, svc.SelfLink().String(), tail, flw)
	//
	//req, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	log.Errorf("%s:logs:> create http client err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//
	//req.WithContext(cx)
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", handler.Config.SecretToken))
	//
	//res, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	log.Errorf("%s:logs:> get pod logs err: %s", logPrefix, err.Error())
	//	errors.HTTP.InternalServerError(w)
	//	return
	//}
	//
	//defer cancel()
	//var buffer = make([]byte, BufferSize)
	//
	//for {
	//
	//	select {
	//	case <-r.Context().Done():
	//		return
	//	default:
	//
	//		n, err := res.Body.Read(buffer)
	//		if err != nil {
	//
	//			if err == context.Canceled {
	//				log.Debug("Stream is canceled")
	//				return
	//			}
	//
	//			log.Errorf("Error read bytes from stream %s", err)
	//			return
	//		}
	//
	//		_, err = func(p []byte) (n int, err error) {
	//
	//			n, err = w.Write(p)
	//			if err != nil {
	//				log.Errorf("Error write bytes to stream %s", err)
	//				return n, err
	//			}
	//
	//			if f, ok := w.(http.Flusher); ok {
	//				f.Flush()
	//			}
	//
	//			return n, nil
	//		}(buffer[0:n])
	//
	//		if err != nil {
	//			log.Errorf("Error written to stream %s", err)
	//			return
	//		}
	//
	//		for i := 0; i < n; i++ {
	//			buffer[i] = 0
	//		}
	//	}
	//}

}
