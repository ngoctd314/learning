# Remote code execution

Attackers can inject malicious code to be executed in the language of the web server itself. A tactic called remote code execution. Remote code execution attacks on websites are rarer than the injection attacks, but every bit as dangerous.

## Anatomy of a remote code execution attack

The exploit script incorporates malicious code in the body of the HTTP request, encoded in such a way that the server will read and execute that code when the request is handled. Security researchers will analyze codebases for common web servers, looking for vulnerabilities that permit malicious code to be injected.

In early 2013, researchers discovered a vulnerability in Ruby on Rails that permitted attackers to inject their own Ruby code into the server process.

## Mitigation: disable code execution during deserialization

Remote code execution vulnerabilities usually occur when web server software insecure serialization. Serialization is the process of converting an in-memory data structure into a stream of binary data, usually for the purpose of passing the data structure across a network. Deserialization refers to the reverse process that occurs at the other end, when the binary data is converted back into a data structure.

Serialization libraries exist in every major programming language and are widely used. Some serialization libraries, such as the YAML parser, allow data structures to execute code as they reinitialize themselves in memory. This is a useful feature if you trust the source of the serialized data, but can be very dangerous if you don't, because it can permit arbitrary code execution.
