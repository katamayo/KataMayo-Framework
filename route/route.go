package route

import (
	"log"
	"net/http"
	"time"

	"github.com/bayurstarcool/bayurGo/app/controllers"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (r *router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func NewRouter() *router {
	return &router{httprouter.New()}
}
func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

// We could also put *httprouter.Router in a field to not get access to the original methods (GET, POST, etc. in uppercase)
type router struct {
	*httprouter.Router
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
func RouteApp(appContext *controllers.AppContext) (r *router) {
	appC := appContext
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	router := NewRouter()
	router.NotFound = http.HandlerFunc(controllers.MyNotFound)
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	// router.Get("/admin", commonHandlers.Append(appC.AuthHandler).ThenFunc(appC.AdminHandler))
	router.Get("/about", commonHandlers.ThenFunc(controllers.AboutHandler))
	router.Get("/", commonHandlers.ThenFunc(appC.IndexHandler))
	router.Get("/login", commonHandlers.ThenFunc(appC.LoginHandler))
	router.Get("/dashboard", commonHandlers.ThenFunc(appC.DashboardHandler))
	router.Get("/anggota", commonHandlers.ThenFunc(appC.AnggotaHandler))
	router.Get("/listanggota", commonHandlers.ThenFunc(appC.ListAnggotaHandler))
	router.Get("/alumni", commonHandlers.ThenFunc(appC.AlumniHandler))
	router.Get("/listalumni", commonHandlers.ThenFunc(appC.ListAlumniHandler))
	router.Get("/teas/:query", commonHandlers.ThenFunc(appC.TeaHandler))
	return router
}
