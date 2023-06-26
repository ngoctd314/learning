# Cross origin resource sharing (CORS)

Cross-Origin Resource Sharing (CORS) is an HTTP-header based mechnism that allows a server to indicate any origins (domain, scheme, or port) other that its own from which a browser should permit loading resources. CORS also relies on a mechanism by which browsers make a preflight request to the server hosting the cross-origin resource, in order to check that the server will permit the actual request. In that preflight, the browser sends headers that indicate the HTTP method and headers that will be used to the actual request.

For security reasons, browsers restrict cross-origin HTTP requests initiated from scripts. XMLHttpRequest and the Fetch API follow the same-origin policy. This means that a web application using those APIs can only request resources from the same origin the application was loaded from unless the response from the other origins includes the right CORS headers.

