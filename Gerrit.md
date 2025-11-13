# Go FastCGI
See Gerrit [Change #715680](https://go.dev/cl/715680).

## Examples for Go Standard Library Change
There are two examples in this repo meant to show the use case for SSI, Server Side Includes, and FCGI. The entire example can be run with docker compose.

1. cloning the repository and running `docker compose up --build`
2. Open example to see SSI and FCGI in action
    - Working Example: [localhost](http://localhost/)
    - Pre-change Example: [localhost:8500](http://localhost:8500/)

## Explanation of the Example

The intent is for the majority of content to be static. To prevent duplicating code on every page, an include can be used to pull in the shared pieces. An include can also be used to pull in small dynamic pieces server side rather than client side.

In the index.html file, [login](sftp_content/index.html), server side includes make a call to the Go FCGI server. The go server will make a call the the json feed requested and return it for inclusion in the static page.

This is an overly simplistic example. It may not be clear from this example how it could be used in a "Real World" or production environment. For a more extensive example, see the [main branch](https://github.com/gibriil/enterprise_portal_example/tree/main).