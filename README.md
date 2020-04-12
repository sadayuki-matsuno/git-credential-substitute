# git-helper-substitute

git credential helper, which switches git credentials for the username/organization.

If you are using multiple github accounts, it uses the information of the organization to determine the authentication information.

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

# Quick Start

- install git-credential-substitute

```
go get -u github.com/sadayuki-matsuno/git-credential-substitute
```

- make `.git-secret.json`

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

- change git-credential-helper

```
git config --global credential.helper substitute
```

- check your git-credential-healper

```
git config --show-origin --get credential.helper

file:$HOME/.gitconfig    substitute
```

**If you see other helpers, you need to remove their settings or they may not work properly.**

- All settings are complete. All that's left to do is to use the git command as usual

# Debug

To check which directory uses which credentials, try running `git-credential-substitute` in the target directory.

# Author

[sadayuki-matsuno](https://github.com/sadayuki-matsuno)
