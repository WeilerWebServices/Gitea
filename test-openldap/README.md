# OpenLDAP Docker Image for testing

This image provides an OpenLDAP Server for testing LDAP applications, i.e. unit tests. The server is initialized with the example domain `planetexpress.com` with data from the [Futurama Wiki][futuramawikia].

Based on Rafael Römhild [docker-test-openldap][dockertestopenldap].

Parts of the image are based on the work from Nick Stenning [docker-slapd][slapd] and Bertrand Gouny [docker-openldap][openldap].

The Flask extension [flask-ldapconn][flaskldapconn] use this image for unit tests.

[slapd]: https://github.com/nickstenning/docker-slapd
[openldap]: https://github.com/osixia/docker-openldap
[flaskldapconn]: https://github.com/rroemhild/flask-ldapconn
[futuramawikia]: http://futurama.wikia.com
[dockertestopenldap]: https://github.com/rroemhild/docker-test-openldap


## Features

* Support for TLS (snake oil cert on build)
* Initialized with data from Futurama
* ~124MB images size (~40MB compressed)


## Usage

```
docker pull gitea/test-openldap
docker run --privileged -d -p 389:389 gitea/test-openldap
```

## Exposed ports

* 389
* 636

## Exposed volumes

* /etc/ldap/slapd.d
* /etc/ldap/ssl
* /var/lib/ldap
* /run/slapd


## LDAP structure

### dc=planetexpress,dc=com

| Admin            | Secret           |
| ---------------- | ---------------- |
| cn=admin,dc=planetexpress,dc=com | GoodNewsEveryone |

### ou=service,dc=planetexpress,dc=com

#### uid=gitea,ou=service,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | account |
| description      | Service account for integration with Gitea |
| mail             | git@planetexpress.com |
| uid              | gitea |
| userPassword     | password |

### ou=people,dc=planetexpress,dc=com

#### cn=Hubert J. Farnsworth,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | Hubert J. Farnsworth |
| sn               | Farnsworth |
| description      | Human |
| displayName      | Professor Farnsworth |
| employeeType     | Owner |
| employeeType     | Founder |
| givenName        | Hubert |
| jpegPhoto        | JPEG-Photo (630x507 Pixel, 26780 Bytes) |
| mail             | professor@planetexpress.com |
| mail             | hubert@planetexpress.com |
| ou               | Office Management |
| title            | Professor |
| uid              | professor |
| userPassword     | professor |


### cn=Philip J. Fry,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | Philip J. Fry |
| sn               | Fry |
| description      | Human |
| displayName      | Fry |
| employeeType     | Delivery boy |
| givenName        | Philip |
| jpegPhoto        | JPEG-Photo (429x350 Pixel, 22132 Bytes) |
| mail             | fry@planetexpress.com |
| ou               | Delivering Crew |
| uid              | fry |
| userPassword     | fry |


### cn=John A. Zoidberg,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | John A. Zoidberg |
| sn               | Zoidberg |
| description      | Decapodian |
| displayName      | Zoidberg |
| employeeType     | Doctor |
| givenName        | John |
| jpegPhoto        | JPEG-Photo (343x280 Pixel, 26438 Bytes) |
| mail             | zoidberg@planetexpress.com |
| ou               | Staff |
| title            | Ph. D. |
| uid              | zoidberg |
| userPassword     | zoidberg |

### cn=Hermes Conrad,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | Hermes Conrad |
| sn               | Conrad |
| description      | Human |
| employeeType     | Bureaucrat |
| employeeType     | Accountant |
| givenName        | Hermes |
| mail             | hermes@planetexpress.com |
| ou               | Office Management |
| uid              | hermes |
| userPassword     | hermes |
| sshPublicKey     | ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC8Fk93z/DLVBj4gHUw3+LOhaIAwCmmXHSfCOlD9Pa1NUTDgURf32m//tRBSDn6o9BTsaHkXyOdTYUF6mXxfwdHaGx3bfXDnUQEX/1J88x0LL6p+sigrGc9/2OUZtL5Af4lNPgbUl1U15U2hhh/Nv9URObSPIbAxURIIArrfYMgDNcUoA/BA4dxnk2lc9Mc/Fozkx7N7bNVT1GOAtosR5Y+ukdTwJFzKmrH6hBAzRKVIxfa4XT+cOmOYW4kL/ca/owHQURxqI4VMKcLEkEYdwANRt1/HJL5kJkpDPQF2gmrH+XNs46E3zicMIzEzKa53gks9hiyRt/AT7UMwXbmSbGv hermes@pc |
| sshPublicKey     | ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOajx6+YcZCdu97aJiC0cIoOkk63kp8mv3ZAOrC0m0g6kibiPmCVmo/ITegtu58boy0CMpmwD/thk36FFdy4Ig5ZNMCRv+9m6PtIE06pvUB6rtYfgzwMc+G3Wibs/zsb7XUwl6Cl/JtsYyeXpdlzIgSt7SzrWf+BVQxOAtJD7bjLXtvQhjHbBO21Zh/Xp0kKlMWlzhSSs5AEjUY0jRAXMPkXPzf8dqhd6JKpdxvk1fjc14BWISZqgeiLpUUL+hk+md5SB31NEKtTkx6nQ5QbXe9BSwu0CtxjAq1vaHwL3bLNL7NFArrltIfOh0BDSADOVfMJ6nSnRgv4uk50gnUr2F hermes@laptop |
| sshPublicKey     | ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH+EtG8OXSsgX9sRldi86H0lBEY/vZAeejxHxDjG1rF+SYF7piCA4C6vLmLKAYz6aWueK0QNvMVlR7SOXESnXEs= hermes@server |

Private key `hermes@pc`:
```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAvBZPd8/wy1QY+IB1MN/izoWiAMApplx0nwjpQ/T2tTVEw4FE
X99pv/7UQUg5+qPQU7Gh5F8jnU2FBepl8X8HR2hsd231w51EBF/9SfPMdCy+qfrI
oKxnPf9jlGbS+QH+JTT4G1JdVNeVNoYYfzb/VETm0jyGwMVESCAK632DIAzXFKAP
wQOHcZ5NpXPTHPxaM5Meze2zVU9RjgLaLEeWPrpHU8CRcypqx+oQQM0SlSMX2uF0
/nDpjmFuJC/3Gv6MB0FEcaiOFTCnCxJBGHcADUbdfxyS+ZCZKQz0BdoJqx/lzbOO
hN84nDCMxMymud4JLPYYskbfwE+1DMF25kmxrwIDAQABAoIBAFJH+E/TueVZTus3
Vw0ghFoQf8SQTCgo/iOshE9kVKEFQqW8YSsH031Yf4ZnkGWjUDms1cPQEwZ3qR/j
YOF6zrZgakL86ay/mcfILkgxzVltUaOOwEH56dhnZyq+qMCiLIoeWAOrzoSVIwON
Oh488wnscoW6UMD4C1z6F4zZhYl5E82pQpwacoBmxa9VjCuY+zLhzJKVfZ6N0KWt
d08XcS5rG+ZGpmF4g4LTLFxuBWUJw+3HUHzlolW+II2g6LYNs98DmYdQp+5d/wyD
nPpIamO4ta69AL1qcCNCRGVnUByOUF/WpHM++f3XFcLhbJ4kVbrtQMG53SZ8mCER
GfuSfzECgYEA73GeVAb2BlmrbQpp6PXaFG80quswXcSTExaZhIMQotwb2+Q3PG9a
CUXmtgx8fY95l1EVrCJvW2LsbJIqe0kXwNXBuHDLtFbBTIlsOVrvmzw3qHuCmj3l
VcaPXDZV655NzDNvHsufJjIghfPhLYJn23/+38G1hrTRFmHKbMjicXMCgYEAyReh
YoehC30NMyn775t4kzPHTdydFIZUNUqxa4znA56n+JgIzKWsp8THjUhrfwv90h6U
vh/nD3PC9JJZARs5cWmveVVLppachcpl0OXuEtPadW8oCi1PK+dV5ZrM++NeGfns
9qUOvGs7TRJAyZnjeQ//4vXwBnJceBQxiMKyP9UCgYB90z/3Of05Ew/xagKDyAYE
rPeJRbQR4kXDRyH/L3yjiHUfVUrteDphGxmE2wTkWmvz50kzPpkz9cT1vM2UIbHY
xLta9/Mj3l8PoDt93FqDQd6hq93Svenw7DnTpD38ZiDNyM2A6lHEmZzbp2/SIXAA
Ob+ux7Vjh0tErrjX96x/HwKBgQCjPaSLyJAqNFSP0SgsRSgnTuQex9VYAQfAzyah
qRsrbBLtEfYqst8qvepEPaLN2p0sghi7EkjO2WlMgrTv2frSnzmMJHqp/B+J2Fi4
sL6H2CTCKf716/wWJtAq+HQoklUkfycmvazttZrJIOUpAtyOvTc7NeyoPxPjgnkH
jQ2IFQKBgHENrLBtC1AS5EGsw67z7Dhn0lLu6y5lorXXpYbpdIjGLpI4XaUqIHdV
+5IEa3vT+tpppA0lOy2g8s5b89kgBaXI1h21Bi42q7SL8ZDQnZvzKP6XMstFX0GJ
gUMPD/s2KNO+mwusxyPadtEy+D30VvUwXEBUvXalKsLXg3gZTcre
-----END RSA PRIVATE KEY-----
```
Private key `hermes@laptop`:
```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzmo8evmHGQnbve2iYgtHCKDpJOt5KfJr92QDqwtJtIOpIm4j
5glZqPyE3oLbufG6MtAjKZsA/7YZN+hRXcuCIOWTTAkb/vZuj7SBNOqb1Aeq7WH4
M8DHPht1om7P87G+11MJegpfybbGMnl6XZcyIEre0s61n/gVUMTgLSQ+24y17b0I
Yx2wTttWYf16dJCpTFpc4UkrOQBI1GNI0QFzD5Fz83/HaoXeiSqXcb5NX43NeAVi
EmaoHoi6VFC/oZPpneUgd9TRCrU5Mep0OUG13vQUsLtArcYwKtb2h8C92yzS+zRQ
K65bSHzodAQ0gAzlXzCep0p0YL+LpOdIJ1K9hQIDAQABAoIBAAXbtgO3eUIYqYfm
aqllsIpqJrPJixLJso6+4+vC0kCxS+eSQKqSsVy+bfbyt9G7LpGqnpTbtVeMj4Kq
sUR8NiFA4sFRsN23mMOzV8dssSd/YDaEhUrSudTlap1Fj9lWvhfWX3p7OOS8rztX
f0WQuPd5qIFLJJR5sgEs1T/yu2X3t0uctiGqHKe3Bm+5T74/1DCBPvW+3HCESiFs
ePlDzK0DuHBNIS116Ff//4ZFSNjClf/Sld61TrJQtdwNLSID3WcofN57wr+2HFO8
Bghljz9uap1m/MarDe6HaHGjlIfc33XOtBybftE2/tPz3bFbpR3lHuKwfikbAunF
80/Bh4ECgYEA+H4NIctq6Oni17I/8YaaRBp7fzYYFmgq009CqamwnxLVvJ8X/iPf
qnI0b5WypIFEj7vzPbuaIrRySy8GVQke2eGsiZrcgvHCwwcmceE484BHh7N93556
d/RXqBnMd6nT+c6In5rPLj4cJfqGgD2QQ13MmII/ID3k0guPc8VSNN0CgYEA1Ka9
bJjEE6ru1vHW7tbmGBLXyFYohcL4VECAosSpv4FbuPpzscwBWCmh/Sfoi/V8nKVr
GJmG9cao2mVJJv2ebo/q9bJ96oXV1A4W5ah4BExG+F524sI+bJJAr7sGsZ/P247r
SbXP1w8VVIMHqaGLtQ/ygX3EMTt3oCqQ76zCbMkCgYBRivvANUJ2ABpCcent1h/V
bWNNUXECGVjEUuQrTNX6vXGKHiL/cMI66pMucs9WkFzxVdnyThe8f75p0ZqgWQfY
s+esmDb4eWFSIoyJHkFUFqpia5JIyXai2nnRXfXs5rv6472Nsn1+TT9rbxSoDIvE
r4kuGr+gUu89xjFi6kOZuQKBgQDHGr+tWHPuPlOWaaVmx0t1Kt9jMliKtXyx1hsb
S6vsJQBueAGvbWWs2H5Ve/JeaSGdwbw+sjENGk6q/b66hSi8OIA0QEVpOpp1DCQg
L9b/nzOsBTanJlwwZ9etMh4YXZvO5UgkIdlScUr1cCHSj/ExPJdA6zKxLg7ZpkFC
R61bEQKBgCJTuCZzEXN6AZ721QgJWlJakcqV62NTXlRxPR9kW4msRHodkXMVjUhF
IAS9tJObLcVpOiL2TaZ5jvrjP/9u8Zq7AVLmY36oaEz3Uw2abGGn0+lM0Ai019p7
g8Jnx5PIr5hLY4qCMOY2ItdP5n+Kne9WwZaQETULVUV+m2X+aC/i
-----END RSA PRIVATE KEY-----
```

### cn=Turanga Leela,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | Turanga Leela |
| sn               | Turanga |
| description      | Mutant |
| employeeType     | Captain |
| employeeType     | Pilot |
| givenName        | Leela |
| jpegPhoto        | JPEG-Photo (429x350 Pixel, 26526 Bytes) |
| mail             | leela@planetexpress.com |
| ou               | Delivering Crew |
| uid              | leela |
| userPassword     | leela |

### cn=Bender Bending Rodríguez,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | inetOrgPerson |
| cn               | Bender Bending Rodríguez |
| sn               | Rodríguez |
| description      | Robot |
| employeeType     | Ship's Robot |
| givenName        | Bender |
| jpegPhoto        | JPEG-Photo (436x570 Pixel, 26819 Bytes) |
| mail             | bender@planetexpress.com |
| ou               | Delivering Crew |
| uid              | bender |
| userPassword     | bender |

### cn=admin_staff,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | Group |
| cn               | admin_staff |
| member           | cn=Hubert J. Farnsworth,ou=people,dc=planetexpress,dc=com |
| member           | cn=Hermes Conrad,ou=people,dc=planetexpress,dc=com |

### cn=ship_crew,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | Group |
| cn               | ship_crew |
| member           | cn=Turanga Leela,ou=people,dc=planetexpress,dc=com |
| member           | cn=Philip J. Fry,ou=people,dc=planetexpress,dc=com |
| member           | cn=Bender Bending Rodríguez,ou=people,dc=planetexpress,dc=com |

### cn=git,ou=people,dc=planetexpress,dc=com

| Attribute        | Value            |
| ---------------- | ---------------- |
| objectClass      | Group |
| cn               | git |
| member           | cn=admin_staff,ou=people,dc=planetexpress,dc=com |
| member           | cn=ship_crew,ou=people,dc=planetexpress,dc=com |
