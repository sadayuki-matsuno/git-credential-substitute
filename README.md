# git-helper-substitute

git credential helper, which switches git credentials for the username/organization.

If you are using multiple github accounts, it uses the information of the organization to determine the authentication information.

[日本語記事](https://qiita.com/sadayuki-matsuno/items/8df9469d2914078827b4)

# Dependency

- go version: < 1.13

# Installation

Download from [release](https://github.com/sadayuki-matsuno/git-credential-substitute/releases)

or 

```
go get -u github.com/sadayuki-matsuno/git-credential-substitute
```

# Features

- Store credentials for each github organization and use the appropriate credentials
- If not applicable to any organization, the default credentials are used.
- If you're in the git directory, it will determine the credential based on the organization information in the origin and use it.

# Caution

- You can't use this `git-helper-substitute` with the `osxkeychain`. You should disable all settings of the `osxkeychain`.

# Quick Start

- install git-credential-substitute

```
go get -u github.com/sadayuki-matsuno/git-credential-substitute
```

- make `$HOME/.git-secret.json`

```
cat $HOME/.git-secret.json

{
  "another-organization": {
    "username": "another-github-username",
    "password": "xxxxxxxxxxxxxxxxxxxxxxx"
  },
  "default": {
    "username": "sadayuki-matsuno",
    "password": "xxxxxxxxxxxxxxxx"
  }
}
```

- make sure the git-credential-substitute works

```
git-credential-substitute

It's working fine.
```

- change git-credential-helper

```
git config --global credential.helper substitute
```

- check your git-credential-healper and disable `osxkeychain`

```
git config --show-origin --get credential.helper

file:$HOME/.gitconfig    substitute
```

```
git config --system --show-origin --get credential.helper
```

**If you see osxkeychain, you need to remove their settings or they may not work properly.**

- All settings are complete. All that's left to do is to use the git command as usual

# Usage

- git clone
    - Because the current directory is used to determine the credential information, when you do `git clone`, create a directory of the organization name first, and then `git clone` in the directory.

```
mkdir `${organization_name}`
cd `${organization_name}`
git clone https://github.com/${organization_name}/xxxx
```

- other git command
    - If you're in the git directory, your credentials are automatically determined using the organization of the origin

# Debug

- To check which directory uses which credentials, try running `git-credential-substitute` in the target directory.
- Putting `GIT_CURL_VERBOSE=1` at the beginning of git command 
    - ex.  `GIT_CURL_VERBOSE=1 git clone xxx/xxxx`

# Author

[sadayuki-matsuno](https://github.com/sadayuki-matsuno)
