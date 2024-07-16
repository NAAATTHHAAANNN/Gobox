<h1>Gobox 💬</h1>


<strong>Gobox</strong> is a simple library used for creating popup dialogs on <strong>Windows</strong>,<strong>Linux</strong> and <strong>Mac</strong>.

<strong>Getting started:</strong>

```bash
go get github.com/jexlor/gobox
```

```go
import( "github.com/jexlor/gobox" ) //import library
```

<h1>Usage:</h1> 

```go
func main() {
	gobox.DialogBox() //this is the function which popups dialog box. let's configure it..
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error") //title of a box
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error", gobox.Error, ) //add icon (there are also Question,Info and Warning icons)
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error", gobox.Error, "Error while loading files!") //add message
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error", gobox.Error, "Error while loading files!", "Ok") //add button text
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error", gobox.Error, "Error while loading files!", "Ok", gobox.StandardSize) //standard fontsize for your os (you can change it manually though)
}
```
-----------
```go
func main() {
	gobox.DialogBox("Error", gobox.Error, "Error while loading files!", "Ok", gobox.StandardSize, 0, 0) //add height and width(0 0 is a default size)
}
```
<h2>Result:</h2> 

![Screenshot from 2024-07-16 17-37-23](https://github.com/user-attachments/assets/d6dfed41-62f2-46e4-a14f-2ccd48cddb60)




