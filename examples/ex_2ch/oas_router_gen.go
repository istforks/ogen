// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/istforks/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [3]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "api/"

				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "captcha/"

					if l := len("captcha/"); len(elem) >= l && elem[0:l] == "captcha/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '2': // Prefix: "2chcaptcha/"

						if l := len("2chcaptcha/"); len(elem) >= l && elem[0:l] == "2chcaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptcha2chcaptchaIDGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						case 's': // Prefix: "show"

							if l := len("show"); len(elem) >= l && elem[0:l] == "show" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptcha2chcaptchaShowGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						}

					case 'a': // Prefix: "app/id/"

						if l := len("app/id/"); len(elem) >= l && elem[0:l] == "app/id/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "public_key"
						// Leaf parameter, slashes are prohibited
						idx := strings.IndexByte(elem, '/')
						if idx >= 0 {
							break
						}
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAPICaptchaAppIDPublicKeyGetRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

					case 'i': // Prefix: "invisible_recaptcha/"

						if l := len("invisible_recaptcha/"); len(elem) >= l && elem[0:l] == "invisible_recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaInvisibleRecaptchaIDGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						case 'm': // Prefix: "mobile"

							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaInvisibleRecaptchaMobileGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						}

					case 'r': // Prefix: "recaptcha/"

						if l := len("recaptcha/"); len(elem) >= l && elem[0:l] == "recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaRecaptchaIDGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						case 'm': // Prefix: "mobile"

							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPICaptchaRecaptchaMobileGetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						}

					}

				case 'd': // Prefix: "dislike"

					if l := len("dislike"); len(elem) >= l && elem[0:l] == "dislike" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPIDislikeGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 'l': // Prefix: "like"

					if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAPILikeGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

				case 'm': // Prefix: "mobile/v2/"

					if l := len("mobile/v2/"); len(elem) >= l && elem[0:l] == "mobile/v2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "after/"

						if l := len("after/"); len(elem) >= l && elem[0:l] == "after/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "thread"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[1] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "num"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "GET":
										s.handleAPIMobileV2AfterBoardThreadNumGetRequest([3]string{
											args[0],
											args[1],
											args[2],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "GET")
									}

									return
								}

							}

						}

					case 'b': // Prefix: "boards"

						if l := len("boards"); len(elem) >= l && elem[0:l] == "boards" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAPIMobileV2BoardsGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

					case 'i': // Prefix: "info/"

						if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "thread"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPIMobileV2InfoBoardThreadGetRequest([2]string{
										args[0],
										args[1],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						}

					case 'p': // Prefix: "post/"

						if l := len("post/"); len(elem) >= l && elem[0:l] == "post/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "num"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleAPIMobileV2PostBoardNumGetRequest([2]string{
										args[0],
										args[1],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}

						}

					}

				}

			case 'u': // Prefix: "user/"

				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "p"

					if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "asslogin"

						if l := len("asslogin"); len(elem) >= l && elem[0:l] == "asslogin" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleUserPassloginPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

					case 'o': // Prefix: "osting"

						if l := len("osting"); len(elem) >= l && elem[0:l] == "osting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleUserPostingPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}

					}

				case 'r': // Prefix: "report"

					if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "POST":
							s.handleUserReportPostRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "POST")
						}

						return
					}

				}

			}

		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [3]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"

			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "api/"

				if l := len("api/"); len(elem) >= l && elem[0:l] == "api/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "captcha/"

					if l := len("captcha/"); len(elem) >= l && elem[0:l] == "captcha/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '2': // Prefix: "2chcaptcha/"

						if l := len("2chcaptcha/"); len(elem) >= l && elem[0:l] == "2chcaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptcha2chcaptchaIDGetOperation
									r.summary = "Получение ид для использования 2chcaptcha."
									r.operationID = ""
									r.pathPattern = "/api/captcha/2chcaptcha/id"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 's': // Prefix: "show"

							if l := len("show"); len(elem) >= l && elem[0:l] == "show" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptcha2chcaptchaShowGetOperation
									r.summary = "Отображение 2chcaptcha по id."
									r.operationID = ""
									r.pathPattern = "/api/captcha/2chcaptcha/show"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						}

					case 'a': // Prefix: "app/id/"

						if l := len("app/id/"); len(elem) >= l && elem[0:l] == "app/id/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "public_key"
						// Leaf parameter, slashes are prohibited
						idx := strings.IndexByte(elem, '/')
						if idx >= 0 {
							break
						}
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = APICaptchaAppIDPublicKeyGetOperation
								r.summary = "Получение app_response_id для отправки поста."
								r.operationID = ""
								r.pathPattern = "/api/captcha/app/id/{public_key}"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

					case 'i': // Prefix: "invisible_recaptcha/"

						if l := len("invisible_recaptcha/"); len(elem) >= l && elem[0:l] == "invisible_recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptchaInvisibleRecaptchaIDGetOperation
									r.summary = "Получение публичного ключа invisible recaptcha."
									r.operationID = ""
									r.pathPattern = "/api/captcha/invisible_recaptcha/id"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'm': // Prefix: "mobile"

							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptchaInvisibleRecaptchaMobileGetOperation
									r.summary = "Получение html страницы для решения капчи, CORS отключён."
									r.operationID = ""
									r.pathPattern = "/api/captcha/invisible_recaptcha/mobile"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						}

					case 'r': // Prefix: "recaptcha/"

						if l := len("recaptcha/"); len(elem) >= l && elem[0:l] == "recaptcha/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'i': // Prefix: "id"

							if l := len("id"); len(elem) >= l && elem[0:l] == "id" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptchaRecaptchaIDGetOperation
									r.summary = "Получение публичного ключа recaptcha v2."
									r.operationID = ""
									r.pathPattern = "/api/captcha/recaptcha/id"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						case 'm': // Prefix: "mobile"

							if l := len("mobile"); len(elem) >= l && elem[0:l] == "mobile" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APICaptchaRecaptchaMobileGetOperation
									r.summary = "Получение html страницы для решения капчи, CORS отключён."
									r.operationID = ""
									r.pathPattern = "/api/captcha/recaptcha/mobile"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

						}

					}

				case 'd': // Prefix: "dislike"

					if l := len("dislike"); len(elem) >= l && elem[0:l] == "dislike" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = APIDislikeGetOperation
							r.summary = "Добавление дизлайка на пост."
							r.operationID = ""
							r.pathPattern = "/api/dislike"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'l': // Prefix: "like"

					if l := len("like"); len(elem) >= l && elem[0:l] == "like" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = APILikeGetOperation
							r.summary = "Добавление лайка на пост."
							r.operationID = ""
							r.pathPattern = "/api/like"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				case 'm': // Prefix: "mobile/v2/"

					if l := len("mobile/v2/"); len(elem) >= l && elem[0:l] == "mobile/v2/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "after/"

						if l := len("after/"); len(elem) >= l && elem[0:l] == "after/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "thread"
							// Match until "/"
							idx := strings.IndexByte(elem, '/')
							if idx < 0 {
								idx = len(elem)
							}
							args[1] = elem[:idx]
							elem = elem[idx:]

							if len(elem) == 0 {
								break
							}
							switch elem[0] {
							case '/': // Prefix: "/"

								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "num"
								// Leaf parameter, slashes are prohibited
								idx := strings.IndexByte(elem, '/')
								if idx >= 0 {
									break
								}
								args[2] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "GET":
										r.name = APIMobileV2AfterBoardThreadNumGetOperation
										r.summary = "Получение постов в треде >= указанного. Не рекомендуется использовать для получения треда целиком, только для проверки новых постов."
										r.operationID = ""
										r.pathPattern = "/api/mobile/v2/after/{board}/{thread}/{num}"
										r.args = args
										r.count = 3
										return r, true
									default:
										return
									}
								}

							}

						}

					case 'b': // Prefix: "boards"

						if l := len("boards"); len(elem) >= l && elem[0:l] == "boards" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = APIMobileV2BoardsGetOperation
								r.summary = "Получение списка досок и их настроек."
								r.operationID = ""
								r.pathPattern = "/api/mobile/v2/boards"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case 'i': // Prefix: "info/"

						if l := len("info/"); len(elem) >= l && elem[0:l] == "info/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "thread"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APIMobileV2InfoBoardThreadGetOperation
									r.summary = "Получение информации о треде."
									r.operationID = ""
									r.pathPattern = "/api/mobile/v2/info/{board}/{thread}"
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}

						}

					case 'p': // Prefix: "post/"

						if l := len("post/"); len(elem) >= l && elem[0:l] == "post/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "board"
						// Match until "/"
						idx := strings.IndexByte(elem, '/')
						if idx < 0 {
							idx = len(elem)
						}
						args[0] = elem[:idx]
						elem = elem[idx:]

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case '/': // Prefix: "/"

							if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "num"
							// Leaf parameter, slashes are prohibited
							idx := strings.IndexByte(elem, '/')
							if idx >= 0 {
								break
							}
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "GET":
									r.name = APIMobileV2PostBoardNumGetOperation
									r.summary = "Получение информации о посте."
									r.operationID = ""
									r.pathPattern = "/api/mobile/v2/post/{board}/{num}"
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}

						}

					}

				}

			case 'u': // Prefix: "user/"

				if l := len("user/"); len(elem) >= l && elem[0:l] == "user/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'p': // Prefix: "p"

					if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "asslogin"

						if l := len("asslogin"); len(elem) >= l && elem[0:l] == "asslogin" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = UserPassloginPostOperation
								r.summary = "Авторизация пасскода."
								r.operationID = ""
								r.pathPattern = "/user/passlogin"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					case 'o': // Prefix: "osting"

						if l := len("osting"); len(elem) >= l && elem[0:l] == "osting" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "POST":
								r.name = UserPostingPostOperation
								r.summary = "Создание нового поста или треда."
								r.operationID = ""
								r.pathPattern = "/user/posting"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

					}

				case 'r': // Prefix: "report"

					if l := len("report"); len(elem) >= l && elem[0:l] == "report" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "POST":
							r.name = UserReportPostOperation
							r.summary = "Отправка жалобы."
							r.operationID = ""
							r.pathPattern = "/user/report"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

				}

			}

		}
	}
	return r, false
}
