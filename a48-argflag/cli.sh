# WINDOWS
#!/bin/bash

$FilesName = Get-ChildItem -Filter "*.go"

go run $FilesName.FullName