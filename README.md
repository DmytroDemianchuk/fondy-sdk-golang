# Example of using fondy in Golang

## Run (Locally)
1.Need to instal ngrok in computer
```
https://ngrok.com/download
```
2.Download ZIP file
3.Unzip ngrok from the terminal
4.You should register to find out your token
![website-image](../main/assets/token-image.png)
5.Add authtoken
6.Start a tunnel
7.You must copy forwarding link
![website-image](../main/assets/ngrok-image.png)
8.Past forwarding lik to ServerCallbackURL with callback method
![website-image](../main/assets/serverURL-image.png)

```
ngrok http 8080
```
and run app
```
go run client/main.go
```