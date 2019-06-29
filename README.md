# BCrypt-Wrapper
BCrypt-Wrapper is a small wrapper around Go's standard BCrypt-Implementation with the goal of increasing the cost-factor when its needed.
To do this, everytime you use this wrapper when you verify its password, there is a check if the used cost-factor is not already to low.
When the cost-factor is to low, the password will be hashed again with the focused cost-factor and returned.
