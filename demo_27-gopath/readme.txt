// GOPATH เป็น path สำหรับ set ไว้เพื่อให้ folder นั้นๆ เป็น default workspaces

----- การ Set GOPATH -----
บน Mac :
export GOPATH="D:\Works\Learning\Golang\golang_fundamental\demo_27-gopath"
บน Windows :
SET GOPATH="D:\Works\Learning\Golang\golang_fundamental\demo_27-gopath"

echo %GOPATH%

default folder :
บน Mac : export GOPATH="C:\Users\anapa\go"
บน Windows : SET GOPATH="%USERPROFILE%\go"

----- ตัวอย่างการ import dependency -----
https://gorm.io/docs/