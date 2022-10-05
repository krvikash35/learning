
* `HttpURLConnection`: java1.1+. Not recommended. syncOnly. Most of the feature unavailable.
* `Java HttpClient`: java9+(exp), java11+(GA). Official. No Websocket/forms/multi-part/caching/compression. FutureStyleAPI
* `Appache HttpClient`: java7+. No Websocket. widespread. vulnerable. FutureStyleAPI.
* `OkHttp`: java8+. recommended. Yes sync-async/http2/form/multi-part/cookie/auth/compression/cache/ws. callbackStyleAPI