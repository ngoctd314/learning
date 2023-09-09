# How web server work

By its simplest definition, a web server is a computer program that sends back HTML pages in response to HTTP requests. Modern web servers encompass a much broader range of functionality than this suggests.

## Static and Dynamic Resources

Web servers serve two types of content in response to HTTP requests: static resources and dynamic resources. A static resource is an HTML file, image file, or other type of file that the web server returns unaltered in HTTP responses. A dynamic resource is code, a script, or a template that the web server executes or interprets in response to an HTTP request. Modern web servers are capable of hosting both static and dynamic resources. Which resource the server executes or returns depends on the URL in the HTTP request.

### Static Resources

In the early days on the internet, websites consisted mostly of static resources. Developers coded HTML files by hand, and websites consisted of individual HTML files that were deployed to the web server. The deployment of a web-site required the developer to copy all the HTML files to the web server and restart the server process. The browser would make an HTTP request to the web server hosting the website, which would implement the incoming URL as a request for a file on disk.

The browser would make an HTTP request to the web server hosting the website, which would interpret the incoming URL as a request for a file on disk.

When returning a static resource, modern web servers often add data to the HTTP responses or process the static resource before returning it. For example, web servers often dynamically compress large resource files by using the gzip algorithm to reduce the bandwidth used in the response, or add caching headers in HTTP response to instruct the browser to cache and use a local copy of a static resource.

Because static resources are simple files of one form or another, they don't by themselves, exhibit much in the way of security vulnerabilities.

#### CDN (Content Delivery Networks)

A modern innovation designed to improve the delivery speeds of static files is the content delivery network (CDN), which will store duplicated copies of static resources in data centers around the world, and quickly deliver those resources to browsers from the nearest physical location. CDNs like Cloudflare, Akamai, or Amazon CloudFront offload the burder of serving large resource files, such as images, to a third party. Intergrating CDN into your site is usually straight forward, and the CDN service charges a monthly fee depending on the amount of resources you deploy.

#### Content Management Systems

Plenty of websites still consist of mostly static content. Rather than being coded by hand, these sites are generally built using content management systems (CMSs) that provide authoring tools requiring little to no technical knowledge to write the content. CMSs generally impose a uniform style on the pages and allow administrators to update content directly in the browser.

CMS plug-ins can also provide analytics to track visitors, and appoint-ment management or customer support functions, and even create online stores.

### Dynamic Resources

Most modern websites instead use dynamic resources. Often the dynamic resource's code load data from a database in order to populate the HTTP response.

### Templates

### Databases

### Summary

Web servers serve two types of content in response to HTTP requests: static resources, such as images, and dynamic resources, which execute custom code.

Statis resources are resources that we can serve directly from a file-system or a content deliver network to increase the responsiveness of the site.

Dynamic resources on the other hand, are resources that we often define in the form of templates, HTML that's interspersed with programmatic instructions to be implement.
