
# SOAP

Never use a standard XML serializer that supports external entity resources.


# Problem

The server-side implementation uses a standard XML serializer which supports
the references of external `<!ENTITY ...>` resources, and which can be
referenced in the request template for the returned response.

At first I thought the descriptions looked way too similar to the `/etc/passwd`
file format (with all the `:::::` in the description strings), which is why
I assumed that the `/data` API returns the description of a specific user or
group via their UIDs/GIDs. I let this code stay in the `bruteforce.mjs` file
for reference.

The second attack I had in mind was that I know XML parsers have huge security
issues when they're used in a simple manner on the server-side, and without
explicit feature deactivations of external entity resources. There's even an
OWASP article about [XML External Entity Processing](https://owasp.org/www-community/vulnerabilities/XML_External_Entity_\(XXE\)_Processing)
in case you're curious about it.


# Solution

1. The API accepts XML payloads requests and returns the content of `<data><ID>...</ID></data>`
2. Create an `<!ENTITY ...>` with a specific name, and reference it in the request payload.
3. Et voila, you have the `/etc/passwd` file contents.

