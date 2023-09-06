# How browsers work

## Web page rendering

### The rendering pipeline: An overview

When the browser receives an HTTP response, it parses the HTML in the body of the response into a DOM: an in-memory data structure that represents the browser's understanding of the way the page is structured. In modern HTML, the layout of the page can't determined until the whole of the HTML is parsed.

Once the browser generates the DOM, but before anything can be drawn onscreen, styling rules must be applied to each DOM element. Last, after the browser finalizes the structure of the page and breaks down how to apply styling information, it draws the web page onscreen. All of this happens in a fraction of a second, and repeats on a loop as the user interacts with the page.

The browser also loads and executes any JS it comes across as it constructs the DOM. JS code can dynamically make changes to DOM and styling rules, either before the page is rendered or in response to user actions.

### The document object model

Some HTML tags like <script>, <style>, <image>, <font> and <video> tags, can reference an external URL in an attribute. When they're parsed into the DOM, these tags cause the browser to import the external resources, meaning that the browser must initiate a further HTTP request. Modern browsers perform these requests in parallel.

The construction of the DOM from HTML is designed to be as robust as possible. Browsers are forgiving about malformed HTML; they close unclosed tags, insert missing tags, and ignore corrupted tags as needed.

### Styling information

Once the browser has constructed the DOM tree, it needs to determine which DOM nodes correspond to onscreen elements, how to lay out those elements relative to each other, and what styling information to apply to them.

The constructor of the DOM tree and the application of styling rules occur in parallel to the processing of any JS code contained in the web page. This Javascript code can change the structure and layout of the page even before it's rendered.

### Javascript

Modern web pages use Javascript to respond to user actions. Javascript it a fully fledged programming language that is executed by the browser's Javascript engine when web pages are rendered.

By default, any JS code is executed by the browser as soon as the relevant <script> tag is parsed into a DOM node. For JavaScript code loaded from an external URL, this means the code is executed as soon as it is loaded.

This default behavior causes problems if the rendering pipeline hasn't finished parsing the HTML document; the Javascript code will attempt to interact with page elements that may not yet exist in the DOM. To allow for this, <script> tags are offen marked with a defer attribute. This causes the Javascript to execute only when the entire DOM has been constructed.

JS code must be executed within a sandbox, where it's not permitted to perform any of the following actions:

- Start new processes or access other existing processes.
- Read arbitrary chunks of system memory. As a managed memory lanaguage , JS can't read memory outside its sandbox.
- Access the local disk. Modern browsers allow websites to store small amounts of data locally, but this storage is abstracted from the file system itself.
- Access the operating system's network layer.
- Call os functions.

JS executing in the browser sandbox is permitted to do the following actions:

- Read and manipulate the DOM of the current web page.
- Listen to and respond to user actions on the current page by registering event listeners.
- Make HTTP calls on behalf of the user
- Open new web pages or refresh the URL of the current page, but only in response to a user action.
- Write new entries to the browser history and go backward and forward in history
- Ask for user's location. For example, "Google Maps would like to use your location."
- Ask permission to send desktop notifications.

Even with these restrictions, an attacker who can inject malicious JavaScript into your web page can still do a lot of harm by using xss to read credit card details or credentials as a user enters them.

### Before and After Rendering: Everything Else the Browser Does

A browser is much more than a rendering pipeline and a JavaScript engine. In addition to rendering HTML and executing JavaScript, modern browsers contain logic for many other responsibilities. Browsers connect with the os to resolve and cache DNS addresses, interpret and verify security certificates, encode requests in HTTPS if needed, and store and transmit cookies according to the web server's instructions.

1. The use visits www.amazon.com in their favorite browser.
2. The browser attempts to resolve the domain (amazon.com) to an IP address. First, the browser consults the os's DNS cache. If it finds no results, it asks the internet service provider to look in the provider's DNS cache. In the unlikely event that nobody on the ISP has visited the Amazon website before, the ISP will resolve the domain at an authoritative DNS server.
3. Now that it has resolved the IP address, the browser attempts to initiate a TCP
   handshake with the server corresponding to the IP address in order to establish a secure connection.
4. Once the TCP session has been established, the browser constructs an HTTP GET request to www.amazon.com. TCP splits the HTTP request into packets and sends them to the server to be reassembled.
5. At this point, the HTTP conversation upgrades to HTTPS to ensure secure communication. The browser and server undertake a TLS hand-shake, agree on an encryption cypher, and exchange encryption keys.
6. The server uses the secure channel to send back an HTTP response containing HTML of the Amazon front page. The browser parses and displays the page, typically triggering many other HTTP GET requests.

### Summary
