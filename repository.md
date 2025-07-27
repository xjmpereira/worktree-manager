# Valid Repository URLs

Regex: `^(?P<scheme>git@|git:\/\/|file:\/\/|ssh:\/\/|ftps?:\/\/|https?:\/\/)?(?P<user>[a-zA-Z0-9]\w*(:.*)?@)?(?P<host>[a-zA-Z0-9][\w\.]*(:[0-9]*)?)?(?P<path>(\.?\/?~?|~).+)\/(?P<name>([\w\.-]+(\.git)|[\w\.-]+))\/?$`

[regex101](https://regex101.com/r/2IfeQ0/1)

```bash
# Match scheme
(?P<scheme>git@|git:\/\/|file:\/\/|ssh:\/\/|ftps?:\/\/|https?:\/\/)?

# Match user (including optional password)
(?P<user>[a-zA-Z0-9]\w*(:.*)?@)?

# Match host (including optional port)
(?P<host>[a-zA-Z0-9][^\/ \n]*(:[0-9]+)?)

# Match path
(?P<path>(\.?\/~?|~).+)\/

# Match repository name
#  try to first match with '.git' suffix
#  if there is no '.git' match until EOL
(?P<name>([\w\.-]+(\.git)|[\w\.-]+))\/?
```


```bash

# ssh://[<user>@]<host>[:<port>]/<path-to-git-repo>
ssh://host.xz/path/to/repo
ssh://host.xz/path/to/repo.git
ssh://host.xz/path/to/repo.git/

ssh://host.xz:8000/path/to/repo
ssh://host.xz:8000/path/to/repo.git
ssh://host.xz:8000/path/to/repo.git/

ssh://user@host.xz/path/to/repo
ssh://user@host.xz/path/to/repo.git
ssh://user@host.xz/path/to/repo.git/

ssh://user@host.xz:8000/path/to/repo
ssh://user@host.xz:8000/path/to/repo.git
ssh://user@host.xz:8000/path/to/repo.git/

ssh://user:pass@host.xz/path/to/repo
ssh://user:pass@host.xz/path/to/repo.git
ssh://user:pass@host.xz/path/to/repo.git/

ssh://user:pass@host.xz:8000/path/to/repo
ssh://user:pass@host.xz:8000/path/to/repo.git
ssh://user:pass@host.xz:8000/path/to/repo.git/

# ssh://[<user>@]<host>[:<port>]/~<user>/<path-to-git-repo>
ssh://host.xz/~user/path/to/repo
ssh://host.xz/~user/path/to/repo/
ssh://host.xz/~user/path/to/repo.git
ssh://host.xz/~user/path/to/repo.git/

ssh://host.xz:8000/~user/path/to/repo
ssh://host.xz:8000/~user/path/to/repo.git
ssh://host.xz:8000/~user/path/to/repo.git/

ssh://user@host.xz/~user/path/to/repo
ssh://user@host.xz/~user/path/to/repo.git
ssh://user@host.xz/~user/path/to/repo.git/

ssh://user@host.xz:8000/~user/path/to/repo
ssh://user@host.xz:8000/~user/path/to/repo.git
ssh://user@host.xz:8000/~user/path/to/repo.git/

ssh://user:pass@host.xz/~user/path/to/repo
ssh://user:pass@host.xz/~user/path/to/repo.git
ssh://user:pass@host.xz/~user/path/to/repo.git/

ssh://user:pass@host.xz:8000/~user/path/to/repo
ssh://user:pass@host.xz:8000/~user/path/to/repo.git
ssh://user:pass@host.xz:8000/~user/path/to/repo.git/

# [<user>@]<host>:/<path-to-git-repo>
host.xz:/path/to/repo
host.xz:/path/to/repo.git
host.xz:/path/to/repo.git/

host.xz:8000/path/to/repo
host.xz:8000/path/to/repo.git
host.xz:8000/path/to/repo.git/

user@host.xz:/path/to/repo
user@host.xz:/path/to/repo.git
user@host.xz:/path/to/repo.git/

user@host.xz:8000/path/to/repo
user@host.xz:8000/path/to/repo.git
user@host.xz:8000/path/to/repo.git/

user:pass@host.xz:/path/to/repo
user:pass@host.xz:/path/to/repo.git
user:pass@host.xz:/path/to/repo.git/

user:pass@host.xz:8000/path/to/repo
user:pass@host.xz:8000/path/to/repo.git
user:pass@host.xz:8000/path/to/repo.git/

# [<user>@]<host>:~<user>/<path-to-git-repo>
host.xz:~user/path/to/repo
host.xz:~user/path/to/repo.git
host.xz:~user/path/to/repo.git/

host.xz:8000/~user/path/to/repo
host.xz:8000/~user/path/to/repo.git
host.xz:8000/~user/path/to/repo.git/

user@host.xz:~user/path/to/repo
user@host.xz:~user/path/to/repo.git
user@host.xz:~user/path/to/repo.git/

user@host.xz:8000/~user/path/to/repo
user@host.xz:8000/~user/path/to/repo.git
user@host.xz:8000/~user/path/to/repo.git/

user:pass@host.xz:~user/path/to/repo
user:pass@host.xz:~user/path/to/repo.git
user:pass@host.xz:~user/path/to/repo.git/

user:pass@host.xz:8000/~user/path/to/repo
user:pass@host.xz:8000/~user/path/to/repo.git
user:pass@host.xz:8000/~user/path/to/repo.git/

# git://<host>[:<port>]/<path-to-git-repo>
git://host.xz/path/to/repo
git://host.xz/path/to/repo.git
git://host.xz/path/to/repo.git/

git://host.xz:8000/path/to/repo
git://host.xz:8000/path/to/repo.git
git://host.xz:8000/path/to/repo.git/

# git://<host>[:<port>]/~<user>/<path-to-git-repo>
git://host.xz/~user/path/to/repo
git://host.xz/~user/path/to/repo.git
git://host.xz/~user/path/to/repo.git/

git://host.xz:8000/~user/path/to/repo
git://host.xz:8000/~user/path/to/repo.git
git://host.xz:8000/~user/path/to/repo.git/

# git://<host>:<user>/<path-to-git-repo>
git://host.xz:user/repo
git://host.xz:user/repo.git
git://host.xz:user/repo.git/

# git@<host>[:<port>]/<path-to-git-repo>
git@host.xz/path/to/repo
git@host.xz/path/to/repo.git
git@host.xz/path/to/repo.git/

git@host.xz:8000/path/to/repo
git@host.xz:8000/path/to/repo.git
git@host.xz:8000/path/to/repo.git/

# git@<host>[:<port>]/~<user>/<path-to-git-repo>
git@host.xz/~user/path/to/repo
git@host.xz/~user/path/to/repo.git
git@host.xz/~user/path/to/repo.git/

git@host.xz:8000/~user/path/to/repo
git@host.xz:8000/~user/path/to/repo.git
git@host.xz:8000/~user/path/to/repo.git/

# git@<host>:<user>/<path-to-git-repo>
git@host.xz:user/repo
git@host.xz:user/repo.git
git@host.xz:user/repo.git/

# http[s]://<host>[:<port>]/<path-to-git-repo>
http://host.xz/path/to/repo
http://host.xz/path/to/repo.git
http://host.xz/path/to/repo.git/
https://host.xz/path/to/repo
https://host.xz/path/to/repo.git
https://host.xz/path/to/repo.git/

http://host.xz:8000/path/to/repo
http://host.xz:8000/path/to/repo.git
http://host.xz:8000/path/to/repo.git/
https://host.xz:8000/path/to/repo
https://host.xz:8000/path/to/repo.git
https://host.xz:8000/path/to/repo.git/

http://user@host.xz/path/to/repo
http://user@host.xz/path/to/repo.git
http://user@host.xz/path/to/repo.git/
https://user@host.xz/path/to/repo
https://user@host.xz/path/to/repo.git
https://user@host.xz/path/to/repo.git/

http://user@host.xz:8000/path/to/repo
http://user@host.xz:8000/path/to/repo.git
http://user@host.xz:8000/path/to/repo.git/
https://user@host.xz:8000/path/to/repo
https://user@host.xz:8000/path/to/repo.git
https://user@host.xz:8000/path/to/repo.git/

http://user:pass@host.xz/path/to/repo
http://user:pass@host.xz/path/to/repo.git
http://user:pass@host.xz/path/to/repo.git/
https://user:pass@host.xz/path/to/repo
https://user:pass@host.xz/path/to/repo.git
https://user:pass@host.xz/path/to/repo.git/

http://user:pass@host.xz:8000/path/to/repo
http://user:pass@host.xz:8000/path/to/repo.git
http://user:pass@host.xz:8000/path/to/repo.git/
https://user:pass@host.xz:8000/path/to/repo
https://user:pass@host.xz:8000/path/to/repo.git
https://user:pass@host.xz:8000/path/to/repo.git/

# ftp[s]://<host>[:<port>]/<path-to-git-repo>
ftp://host.xz/path/to/repo
ftp://host.xz/path/to/repo.git
ftp://host.xz/path/to/repo.git/
ftps://host.xz/path/to/repo
ftps://host.xz/path/to/repo.git
ftps://host.xz/path/to/repo.git/

ftp://host.xz:8000/path/to/repo
ftp://host.xz:8000/path/to/repo.git
ftp://host.xz:8000/path/to/repo.git/
ftps://host.xz:8000/path/to/repo
ftps://host.xz:8000/path/to/repo.git
ftps://host.xz:8000/path/to/repo.git/

# Local files
/path/to/repo
/path/to/repo.git
/path/to/repo.git/
./path/to/repo
./path/to/repo.git
./path/to/repo.git/
file:///path/to/repo
file:///path/to/repo.git
file:///path/to/repo.git/

# Uncommon Repository names
https://host.xz/path/to/re-po
https://host.xz/path/to/re-po.git
https://host.xz/path/to/re-po.git/
https://host.xz/path/to/re_po
https://host.xz/path/to/re_po.git
https://host.xz/path/to/re_po.git/
https://host.xz/path/to/re.po
https://host.xz/pa%th/to/re.po.git
https://host.xz/pa%th/to/re.po.git/
https://host.xz/pa.-th/to/re.po
https://host.xz/pa.-th/to/re.po.git
https://host.xz/pa.-th/to/re.po.git/
```

Clone documentation: [git-clone](https://git-scm.com/docs/git-clone)
