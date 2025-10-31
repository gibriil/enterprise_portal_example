# Enterprise Portal with OIDC SSO
This is an example infrastructure for delivering password protected and personalized static content on Apache with a Go backend replacement for PHP. This model is perfect for Universities wanting a password protected portal/intranet with the Modern Campus CMS.

**_This is should not be run as a production environment. This is only an example to inform how a production environment could be setup._**

## Go Standard Library Change
There are three example in this repo meant to show the use case working in a PHP environment, the same setup using Go, but containing errors, and the same setup in Go working after accepting [Change #715680](https://go.dev/cl/715680).

### Comparing Differences

## Running the Example Locally
Running the example should be as simple as cloning the repository and running `docker compose up --build -d`

## Modern Campus CMS Specifics
You're probably wondering how the Modern Campus CMS plays into this. Modern Campus SFTPs static files to an SFTP location. That SFTP location could easily be mounted to a container in an environment like Kubernetes. That is the model this example is intending to mimic.

## Questions and Answers

- **Q:** Can I use any IdP or do I have to use Keycloak
    - **A:** You should be able to use any IdP. Keycloak is used here for the sole purpose of having one all inclusive setup to run and test. Keycloak was the best option for an easy download and run example. Preferred enterprise options like Azure or Okta would not have been an easy setup nor secure. Accounts and Secrets would be required and would be problematic for user testing.
- **Q:** Are there not easier approaches?
    - **A:** Absolutely! There are usually many options for solving any problem. This example is meant to demonstrate a principle that can be applied to many situations. This example is most beneficial for academic institutions utilizing the Modern Campus CMS.
- **Q:** Do I have to use Apache?
    - **A:** No, you should be able to use any server that works with SSI (Server Side Includes) following the format `<!--#include virtual="<some path>" -->` and FCGI/CGI. I believe this would at least be Apache, ASP.Net, and NGINX.
- **Q:** Why not do the whole server in Go, why use FCGI Proxy?
    - You could certainly do that. This model is really dependant on how your organization is structured. An example may be that you have a department of server admins who really don't code, a department that handles server side scripting, and a department that handles front-end web. If you have specialized teams, it likely wouldn't make sense to have one tightly coupled system.
- **Q:** Do I have to use Go?
    - **A:** No. Any language that supports FCGI would work in this setup. PHP-FPM is a common use for this setup. Many organizations are migrating their PHP code to Go. This example shows how to do that with a Modern Campus environment. Pearl and Python are also common alternatives to PHP and could be replaced with Go.
- **Q:** I don't use Modern Campus but I think I could benefit from this approach. What options do I have?
    - **A:** The Modern Campus CMS is my expertise. I'm sure there is a way to utilize this concept in other environments. You could likely take this type of approach with Wordpress, Joomla, or Drupal. Any Apache environment with the correct modules installed would work and could possibly be done at an htaccess level and not the server config.


