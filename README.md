# Config before run
- Change port: You can change value of the Port variable
![/blob/main/CRM/evidences/create_fail.png](port.png)
# How to run
## Using Visual Code
- Open Visual Code IDE
- Click Terminal in Menubar -> Select New Terminar(Or using hotkey(MacOS): "control + `")
## Run using command
- Run the `go get github.com/gorilla/mux`
- Run the `go run .`
## Result
![./CRM/envidences/run.png](run.png)

# Testing
- You can test using Postman tool
- First step: Import GoLang_Udacity.postman_collection.json to Postman tool
- Result:
  - ![./CRM/envidences/get_success.png](get_success.png)
  - ![./CRM/envidences/get_fail.png](get_fail.png)
  - ![./CRM/envidences/get_all_success.png](get_all_success.png)
  - ![./CRM/envidences/create_success.png](create_success.png)
  - ![./CRM/envidences/create_fail.png](create_fail.png)
  - ![./CRM/envidences/update_success.png](update_success.png)
  - ![./CRM/envidences/update_fail.png](update_fail.png)
  - ![./CRM/envidences/delete_success.png](delete_success.png)
  - ![./CRM/envidences/delete_fail.png](delete_fail.png)

