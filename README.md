# Enterprise Portal with OIDC SSO
This is an example infrastructure for delivering password protected and personalized static content on Apache with a Go backend replacement for PHP-FPM. This model is perfect for Universities wanting a password protected portal/intranet with the Modern Campus CMS.

**_This is should not be run as a production environment. This is only an example to inform how a production environment could be setup._**

## Running the Example Locally
Running the example should be as simple as cloning the repository and running `docker compose up --build`

Wait until the terminal stops outputting for keycloak and apache outputs "✅ Keycloak is ready!" and runs the Apache process. Just open [localhost](http://localhost/) to see the full working example.

### Portal Users and Credentials
I have pre-defined 4 users to show how groups/roles can be used by FCGI to manipulate what content is displayed. Go serves the included header for the static html pages. The navigation in the header changes based on the groups/roles attached to the logged in user.

All users share the password `password123` to easily login/logout and change between users.

1. **username:** csr
    - csr will see the *Employee Resources* and *CSR Resources* navigation items
2. **username:** sales
    - sales will see the *Employee Resources* and *Sales Resources* navigation items
3. **username:** lead
    - lead will see the *Employee Resources* and *People Management* navigation items
4. **username:** admin
    - admin will see the ALL navigation items

Users, Groups, and Roles can be modified, prior to running the example, in the [relm-myportal.json](idp/relm-myportal.json) file.

## Modern Campus CMS Specifics
Modern Campus SFTPs static files to an SFTP location. That SFTP location could easily be mounted to a container in an environment like Kubernetes. That is the model this example is intending to mimic. the sftp_content folder is mounted to the Apache container to simulate this type of setup.

## Questions and Answers

- **Q:** Why does SSI (Server Side Includes) matter?
    - **A:** From Apache documentation on SSI, "They let you add dynamically generated content to an existing HTML page, without having to serve the entire page via a CGI program, or other dynamic technology."<br /><br />
    This is extremely handy when you want to primarily serve static content but you want a little bit of server side scripting. Instead of rendering the whole page though PHP or Go, you can serve it statically and just include the little pieces of scripting as needed.<br /><br />
    This approach allows you to not rely on client side JavaScript, which may be desired for many reasons.

- **Q:** Why OpenID and not SAML 2.0?
    - **A:** OpenID is a more universal approach. This example is meant to show how GO FCGI and Server Side Includes can be used together. OpenID provides a case for wider use than Higher Education. SAML 2.0 is most certainly an option.

- **Q:** Can I use any IdP or do I have to use Keycloak
    - **A:** You should be able to use any IdP. Keycloak is used here for the sole purpose of having one all inclusive setup to run and test. Keycloak was the best option for an easy download and run example. Preferred enterprise options like Azure or Okta would not have been an easy setup nor secure. Accounts and Secrets would be required and would be problematic for user testing.

- **Q:** Are there not easier approaches?
    - **A:** There are usually many options for solving any problem. This example is meant to demonstrate a principle that can be applied to many situations. This example is most beneficial for academic institutions utilizing the Modern Campus CMS. "Easier" is relative to what your needs are.

- **Q:** Do I have to use Apache?
    - **A:** No, you should be able to use any server that works with SSI (Server Side Includes) following the format `<!--#include virtual="<some path>" -->` and FCGI/CGI. I believe this would at least be Apache, ASP.Net, and NGINX.

- **Q:** Why not do the whole server in Go, why use FCGI Proxy?
    - **A:** You could certainly do that. This model is really dependant on how your organization is structured. An example may be that you have a department of server admins who really don't code, a department that handles server side scripting, a department that handles front-end web, and/or CMS users across your organization. If you have specialized teams, it likely wouldn't make sense to have one tightly coupled system.<br /><br />
    This also allows you to use servers like Apache and Nginx to serve static content, which they are designed to do and do very well.

- **Q:** Do I have to use Go?
    - **A:** No. Any language that supports FCGI would work in this setup. PHP-FPM is a common use for this setup. Many organizations are migrating their PHP code to Go. This example shows how to do that with a Modern Campus environment. Pearl and Python are also common alternatives to PHP and could be replaced with Go.

- **Q:** I don't use Modern Campus but I think I could benefit from this approach. What options do I have?
    - **A:** The Modern Campus CMS is my expertise. You can take this type of approach with Wordpress, Joomla, or Drupal. Any Apache environment with the correct modules installed would work. You can use htaccess files as to not adjust the server config. I have provided a simplified example of using Go FCGI through Apache SSI while utilizing htaccess files. [See htaccess example](https://github.com/gibriil/enterprise_portal_example/tree/simple-ssi-fcgi).
