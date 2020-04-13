package xjs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pubgo/g/xerror"
	dom "github.com/siongui/godom"
	"log"
	"net/url"
	"regexp"
	"strings"
)

// HandleFunc will cause the router to call f whenever window.location.pathname
// (or window.location.hash, if history.pushState is not supported) matches path.
// path can contain any number of parameters which are denoted with curly brackets.
// So, for example, a path argument of "users/{id}" will be triggered when the user
// visits users/123 and will call the handler function with params["id"] = "123".
func (t *Vapper) Route(path string, handler Handler) {
	t.routes = append(t.routes, newRoute(path, handler))
}

func (t *Vapper) NotFound(handler Handler) {
	t.notFoundRoute = &route{
		handler: handler,
	}
}

// browserSupportsPushState will be true if the current browser
// supports history.pushState and the onpopstate event.
var browserSupportsPushState bool

func init() {
	// We only want to initialize certain things if we are running
	// inside a browser. Otherwise, they will cause the program to
	// panic.
	dom.Window.History()
	browserSupportsPushState = (dom.Window.Get("onpopstate") != js.Undefined) &&
		(dom.Window.Get("history") != js.Undefined) &&
		(dom.Window.Get("history").Get("pushState") != js.Undefined)
}

// Context is used as an argument to Handlers
type Context struct {
	// Params is the parameters from the url as a map of names to values.
	Params map[string]string
	// Path is the path that triggered this particular route. If the hash
	// fallback is being used, the value of path does not include the '#'
	// symbol.
	Path string
	// InitialLoad is true iff this route was triggered during the initial
	// page load. I.e. it is true if this is the first path that the browser
	// was visiting when the javascript finished loading.
	InitialLoad bool
	// QueryParams is the query params from the URL. Because params may be
	// repeated with different values, the value part of the map is a slice
	QueryParams map[string][]string
}

// Handler is a function which is run in response to a specific
// route. A Handler takes a Context as an argument, which gives
// handler functions access to path parameters and other important
// information.
type Handler interface {
	Handle(ctx *Context)
}

// route is a representation of a specific route
type route struct {
	// regex is a regex pattern that matches route
	regex *regexp.Regexp
	// paramNames is an ordered list of parameters expected
	// by route handler
	paramNames []string
	// handler called when route is matched
	handler Handler
}

// newRoute returns a route with the given arguments. paramNames and regex
// are calculated from the path
func newRoute(path string, handler Handler) *route {
	route := &route{
		handler: handler,
	}

	strs := strings.Split(path, "/")
	strs = removeEmptyStrings(strs)
	pattern := `^`
	for _, str := range strs {
		if str[0] == '{' && str[len(str)-1] == '}' {
			pattern += `/`
			pattern += `([^/]*)`
			route.paramNames = append(route.paramNames, str[1:(len(str) - 1)])
		} else {
			pattern += `/`
			pattern += str
		}
	}
	pattern += `/?$`
	route.regex = regexp.MustCompile(pattern)
	return route
}

// Navigate will trigger the handler associated with the given path
// and update window.location accordingly. If the browser supports
// history.pushState, that will be used. Otherwise, Navigate will
// set the hash component of window.location to the given path.
func (t *Vapper) Navigate(path string) {
	if browserSupportsPushState && !t.ForceHashURL {
		pushState(path)
		t.pathChanged(path, false)
	} else {
		setHash(path)
	}
	if t.ShouldInterceptLinks {
		t.InterceptLinks()
	}
}

// CanNavigate returns true if the specified path can be navigated by the
// router, and false otherwise
func (t *Vapper) CanNavigate(path string) bool {
	if bestRoute, _, _ := t.findBestRoute(path); bestRoute != nil {
		return true
	}
	return false
}

// Back will cause the browser to go back to the previous page.
// It has the same effect as the user pressing the back button,
// and is just a wrapper around history.back()
func (t *Vapper) Back() {
	dom.Window.History().Back()
	if t.ShouldInterceptLinks {
		t.InterceptLinks()
	}
}

// InterceptLinks intercepts click events on links of the form <a href="/foo"></a>
// and calls router.Navigate("/foo") instead, which triggers the appropriate Handler
// instead of requesting a new page from the server. Since InterceptLinks works by
// setting event listeners in the DOM, you must call this function whenever the DOM
// is changed. Alternatively, you can set r.ShouldInterceptLinks to true, which will
// trigger this function whenever Start, Navigate, or Back are called, or when the
// onpopstate event is triggered. Even with r.ShouldInterceptLinks set to true, you
// may still need to call this function if you change the DOM manually without
// triggering a route.
func (t *Vapper) InterceptLinks() {
	for _, link := range dom.Document.QuerySelectorAll("links") {
		href := link.GetAttribute("href")
		switch {
		case href == "":
			continue

		case strings.HasPrefix(href, "http://"), strings.HasPrefix(href, "https://"), strings.HasPrefix(href, "//"):
			// These are external links and should behave normally.
			continue

		case strings.HasPrefix(href, "#"):
			// These are anchor links and should behave normally.
			// Recall that even when we are using the hash trick, href
			// attributes should be relative paths without the "#" and
			// router will handle them appropriately.
			continue

		case strings.HasPrefix(href, "/"):
			// These are relative links. The kind that we want to intercept.
			if t.listener != nil {
				// Remove the old listener (if any)
				link.RemoveEventListener("click", t.listener)
				continue
			}

			t.listener = t.interceptLink
			link.AddEventListener("click", t.listener)
		}
	}
}

// interceptLink is intended to be used as a callback function. It stops
// the default behavior of event and instead calls r.Navigate, passing through
// the link's href property.
func (t *Vapper) interceptLink(event dom.Event) {
	path := event.Target().GetAttribute("href")

	// Only intercept the click event if we have a route which matches
	// Otherwise, just do the default.
	if bestRoute, _, _ := t.findBestRoute(path); bestRoute != nil {
		event.PreventDefault() // TODO - don't think this will work?
		go t.Navigate(path)
	}
}

// setInitialHash will set hash to / if there is currently no hash.
func (t *Vapper) setInitialHash() {
	if getHash() == "" {
		setHash("/")
	} else {
		t.pathChanged(getPathFromHash(getHash()), true)
	}
}

// pathChanged should be called whenever the path changes and will trigger
// the appropriate handler. initial should be true iff this is the first
// time the javascript is loaded on the page.
func (t *Vapper) pathChanged(path string, initial bool) {
	bestRoute, tokens, params := t.findBestRoute(path)
	// If no routes match, we throw console error and no handlers are called
	if bestRoute == nil {
		if t.Verbose {
			log.Println("Could not find route to match: " + path)
		}
		t.notFoundRoute.handler.Handle(&Context{
			Path:        path,
			InitialLoad: initial,
			Params:      map[string]string{},
			QueryParams: params,
		})
		return
	}
	// Create the context and pass it through to the handler
	c := &Context{
		Path:        path,
		InitialLoad: initial,
		Params:      map[string]string{},
		QueryParams: params,
	}
	for i, token := range tokens {
		c.Params[bestRoute.paramNames[i]] = token
	}
	bestRoute.handler.Handle(c)
}

// findBestRoute compares the given path against regex patterns of routes.
// Preference given to routes with most literal (non-parameter) matches. For
// example if we have the following:
//   Route 1: /todos/work
//   Route 2: /todos/{category}
// And the path argument is "/todos/work", the bestRoute would be todos/work
// because the string "work" matches the literal in Route 1.
func (t Vapper) findBestRoute(path string) (bestRoute *route, tokens []string, params map[string][]string) {
	parts := strings.SplitN(path, "?", 2)
	leastParams := -1
	for _, route := range t.routes {
		matches := route.regex.FindStringSubmatch(parts[0])
		if matches != nil {
			if (leastParams == -1) || (len(matches) < leastParams) {
				leastParams = len(matches)
				bestRoute = route
				tokens = matches[1:]
			}
		}
	}
	if len(parts) > 1 {
		params = t.parseQueryPart(parts[1])
	}
	return bestRoute, tokens, params
}

// parseQueryPart extracts query params from the query part of the URL
func (t Vapper) parseQueryPart(queryPart string) map[string][]string {
	params, err := url.ParseQuery(queryPart)
	xerror.PanicT(err != nil && t.Verbose, "Error parsing query %s", queryPart)
	return params
}

// watchHash listens to the onhashchange event and calls r.pathChanged when
// it changes
func (t *Vapper) watchHash() {
	dom.Window.AddEventListener("onhashchange", func(event dom.Event) {
		go func() {
			t.pathChanged(getPathFromHash(getHash()), false)
		}()
	})
}

// watchHistory listens to the onpopstate event and calls r.pathChanged when
// it changes
func (t *Vapper) watchHistory() {
	dom.Window.AddEventListener("onpopstate", func(event dom.Event) {
		go func() {
			t.pathChanged(getPath(), false)
			if t.ShouldInterceptLinks {
				t.InterceptLinks()
			}
		}()
	})
}

// getPathFromHash returns everything after the "#" character in hash.
func getPathFromHash(hash string) string {
	return strings.SplitN(hash, "#", 2)[1]
}

// getHash is an alias for js.Global.Get("location").Get("hash").String()
func getHash() string {
	return dom.Window.Location().Get("hash").String()
}

// setHash is an alias for js.Global.Get("location").Set("hash", hash)
func setHash(hash string) {
	dom.Window.Location().Set("hash", hash)
}

// getPath is an alias for js.Global.Get("location").Get("pathname").String()
func getPath() string {
	return dom.Window.Location().Get("pathname").String()
}

// pushState is an alias for js.Global.Get("history").Call("pushState", nil, "", path)
func pushState(path string) {
	dom.Window.History().PushState(nil, "", path)
}

// removeEmptyStrings removes any empty strings from strings
func removeEmptyStrings(strings []string) []string {
	var result []string
	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
