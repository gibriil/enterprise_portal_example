# Go FCGI behind Apache SSI
This is an example of how Apache, or Nginx, can be used in front of a Go FCGI server. It goes beyond just reverse proxying to allow SSI (Server Side Includes) to use the FCGI proxy for including dynamic content into static pages.

From Apache documentation on SSI, "They let you add dynamically generated content to an existing HTML page, without having to serve the entire page via a CGI program, or other dynamic technology."

**_This is should not be run as a production environment. This is only an example to inform how a production environment could be setup._**

This example utilizes htaccess files to setup the proxy instead of having all the configuration in the server config. This is a good example of how this approach can be plugged into an existing Wordpress, Joomla, or Drupal site.

## Running the Example Locally
Running the example should be as simple as cloning the repository and running `docker compose up --build`

Wait until the terminal stops outputting and the apache container is running the Apache process. Just open [localhost](http://localhost/) to see the full working example.

## Questions and Answers

- **Q:** Why does SSI (Server Side Includes) matter?
    - **A:** From Apache documentation on SSI, "They let you add dynamically generated content to an existing HTML page, without having to serve the entire page via a CGI program, or other dynamic technology."<br /><br />
    This is extremely handy when you want to primarily serve static content but you want a little bit of server side scripting. Instead of rendering the whole page though PHP or Go, you can serve it statically and just include the little pieces of scripting as needed.<br /><br />
    This approach allows you to not rely on client side JavaScript, which may be desired for many reasons.

- **Q:** Are there not easier approaches?
    - **A:** There are usually many options for solving any problem. This example is meant to demonstrate a principle that can be applied to many situations. This example is most beneficial for academic institutions utilizing the Modern Campus CMS. "Easier" is relative to what your needs are.

- **Q:** Do I have to use Apache?
    - **A:** No, you should be able to use any server that works with SSI (Server Side Includes) following the format `<!--#include virtual="<some path>" -->` and FCGI/CGI. I believe this would at least be Apache, ASP.Net, and NGINX.

- **Q:** Why not do the whole server in Go, why use FCGI Proxy?
    - **A:** You could certainly do that. This model is really dependant on how your organization is structured. An example may be that you have a department of server admins who really don't code, a department that handles server side scripting, a department that handles front-end web, and/or CMS users across your organization. If you have specialized teams, it likely wouldn't make sense to have one tightly coupled system.<br /><br />
    This also allows you to use servers like Apache and Nginx to serve static content, which they are designed to do and do very well.

- **Q:** Do I have to use Go?
    - **A:** No. Any language that supports FCGI would work in this setup. PHP-FPM is a common use for this setup. Many organizations are migrating their PHP code to Go. This example shows how to do that with a Modern Campus environment. Pearl and Python are also common alternatives to PHP and could be replaced with Go.
