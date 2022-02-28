# Template Fetcher

- [x] CLI
- [x] Working with github
- [ ] Enough tests (more tests should be created)

Template fetcher is a project that allows you to store your project templates in a shared repo and fetch then into your local folder when is needed. 

It assimilates in functionality to cli tools like `aws lambda new` which creates a new project from a template.

## With template fetcher you will be able to:

- Store your microservice template and share with your team 
- Avoid to have to remember all libraries and configurations you need to have in order to create that kind of project  

## How to use: 

```go
// Sync with remote repo
tfetch sync 
// List available template
tfetch list 
// Fetch and download template in current repo
tfetch fetch [path to template repo or doc] 
```  

## How to configure:

```go
// The git token will be encrypted and stored locally in ~/.template.fetcher.config file  
tfetch config --repo="[your repo name]" --repo-owner="[repo owner name]" --git-token="[your git token]"  
```  
