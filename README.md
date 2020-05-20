# GoIgnore

A CLI tool for generating .gitignore files. Still in development.

Takes a passed in list of the technologies or languages you want your ignore file configured for. 
Based on this list it fetches the necessary .gitignore config from this 
[GitHub provided repo](https://github.com/github/gitignore) and creates a `.gitignore` file.

Usage:
```shell
Syntax:
goignore new [tech1, tech2]

Example:
# Generates a .gitignore for the languages Ruby, Node & Go
goignore new node ruby go
```


## To do
- [ ] Better 'tech' name handling (repo directories)
- [ ] Implement user config file for default techs
- [ ] Tests
- [ ] CI/CD to package and upload to S3 for user downloads
