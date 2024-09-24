GET {{host}}/results/string_value

HTTP *
[Asserts]

status < 400
jsonpath "$" isCollection