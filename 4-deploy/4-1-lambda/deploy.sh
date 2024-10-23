GOARCH=arm64 GOOS=linux  CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w" -o bootstrap .
sudo npm i serverless-dotenv-plugin serverless-prune-plugin
serverless deploy
