
# Some Assembly Required

WebASM isn't a good obfuscation method


# Problem

The client-side validates the password/flag on the client-side,
but it's using WebASM as an intermediary language.

# Solution

The generated JavaScript code is hard to read, and was obfuscated
a little. With a step-by-step approach it's doable, and the initial
dictionary used an offset to lookup the static string values.

