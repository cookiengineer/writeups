
# Irish Name Repo 1

SQL Injection Attack


# Problem

On the server-side, there's probably a PHP script running which
builds an SQL query.

The `support.html` file hints to an SQL problem when using the
name `Brian O' Connor`, which means that the server-side script
probably looks similar to this:

```php
$username = $_POST['username'];
$password = $_POST['password'];
$result   = mysql_query("SELECT whatever FROM admin WHERE username = '$username' and password = '$password'");

if ($result) {
	// ...
}
```

The above script, however, doesn't sanitize malicious input which
means that the `login.php` script is vulnerable to an SQL injection.

In the username, we can enter `admin`.
In the password, we can enter `' OR '1'='1';`.

Et voila, admin access.


# Solution

- Always sanitize your inputs
- Always remove special characters if necessary
- Use static query templates, in PHP via `$connection->prepare(...)` and using `?` as placeholders, which can be bound later via `bind_param(type, value)`.

