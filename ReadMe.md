
# GoLang and AWS Lambda
The project structure sets up the build executable to be used as a deployed and used in the AWS cloud.
The files in this project are *not* run locally, they are currently trial-and-error tested as manual deploys to AWS Lambda.
**FUTURE IMPLEMENTATIONS WILL INCLUDE STRATEGIES FOR LOCAL DEVELOPMENT AND DEBUGGING**

## GO Basics
### go.mod
- lists project configurations
- lists external dependencies (i.e. AWS SDK Libraries)

### go.sum
- Simliar to a package-lock type file


---

## Steps to deploy Go code to existing AWS Lambda (Windows OS)
### 1. Ensure that you have removed unused any modules by running ```go mod tidy``` on the command line
 - This command updates your go.mod, go.sum file and removes unused imports

### 2. Ensure that lambda function runs locally
 - An error response such as "AWS needs credentials" is fine, **compilation errors should be addressed before pushing (imports, module not found, etc**)

### 3. Update Go project to run on linux with command (for windows users): `$Env:GOOS = "linux"`
 - This ensures that AWS is able to compile the Go file (AWS Lambda requires Linux with GoLang)

### 4. Build executable for function 
```bash
go build ${filename}.go
```
- This will create an executable file (no file extension) with the same name as the Go file

### 5. Compress Go executable
- Find your created executable from the previous step in your File Explorer.
- Right click and Send To => Compressed Folder

### 6. Deploy Go Executable to AWS
- Log into the AWS Console and Navigate to Lambda
- Search for your existing Lambda Function and Select Name
- In the selected Lambda Function dashboard, select `Upload From` in the Code Source section
- Select `.zip file` and search for your compressed file from the previous step
- Upload and Save

### **Note: Change environment back to windows for local development: ```$Env:GOOS="windows" ```
Failure to do so may result in 'file not found' type error when attempting to run Go files locally