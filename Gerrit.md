# Go FastCGI
The net/http package has everything needed for cgi, fcgi, and http servers. When trying to utilize the CGI/FCGI servers with Apache, I discovered a bug that prevented SSI (Server Side Includes).

I opened an [issue on github](https://github.com/golang/go/issues/70416) reporting the findings. Because FCGI utilizes the `RequestFromMap` function of the net/http/cgi package, the fix needed to be implemented in that package to fix both CGI and FCGI use cases.

Prior to the inclusion of my submitted change, the existing protocol check for fcgi/cgi requests did not properly
account for Apache SSI SERVER_PROTOCOL value of
INCLUDED. The CGI Spec as specified in [RFC 3875 - section 4.1.16](https://www.rfc-editor.org/rfc/rfc3875.html#section-4.1.16) states:
>  A well-known value for SERVER_PROTOCOL which the server MAY use is "INCLUDED", which signals that the current document is being included as part of a composite document, rather than being the direct target of the client request. The script should treat this as an HTTP/1.0 request.

See Gerrit [Change #715680](https://go.dev/cl/715680).

## Examples for Go Standard Library Change
There are two examples in this repo meant to show the use case for SSI, Server Side Includes, and FCGI. The entire example can be run with docker compose.

1. cloning the repository and running `docker compose up --build`
2. Wait until the terminal stops outputting for keycloak and apache outputs "✅ Keycloak is ready!" and runs the Apache process.
3. Modify user, roles, and/or groups. See [README](README.md) for more information
4. Open example to see SSI and FCGI in action
    - Working Example: [localhost](http://localhost/)
    - Pre-change Example: [localhost:8500](http://localhost:8500/)

*Login credentials are detailed in the [README](README.md)*

## Explanation of the Example
*For a more simplified example please see the [simple-ssi-fcgi branch](https://github.com/gibriil/enterprise_portal_example/tree/simple-ssi-fcgi)*

The intent is for the majority of content to be static. To prevent duplicating code on every page, an include can be used to pull in the shared pieces. Because of the authentication aspect of this example, we can setup Go to handle the changes and the authorizing of content pieces.

In the index.html files, [login](sftp_content/index.html) and [dashboard](sftp_content/dashboard/index.html), the header and footer use server side include to make a call to the Go FCGI server. The Go server responds with the respective [header](golang-fcgi/server//internal//html/header.tmpl) and [footer](golang-fcgi/server//internal//html/footer.tmpl) templates. If authenticated, the user is passed to the template and the header template will add navigation links based on the user groups/roles. The template contains additional SSI directives to pull in the drop down navigation menus for those navigation items.

The Go FCGI server also handles the login redirect to redirect the authenticated user to the dashboard from the login screen.