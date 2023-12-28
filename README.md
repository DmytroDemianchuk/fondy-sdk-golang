# Example of using fondy in Golang

## Run (Locally)
1.Need to instal ngrok in computer
```
https://ngrok.com/download
```
2.Download ZIP file

3.Unzip ngrok from the terminal

4.You should register to find out your token

![token-image](../v1/assets/token-image.png)

5.Add authtoken

6.Start a tunnel
```ngrok http 8080```

7.You must copy forwarding link

![ngrok-image](../v1/assets/ngrok-image.png)

8.Past forwarding lik to ServerCallbackURL with callback method

![serverURL-image](../v1/assets/serverURL-image.png)

9.run app
```
go run client/main.go
```