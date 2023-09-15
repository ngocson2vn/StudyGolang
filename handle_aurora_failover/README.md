# Lambda function that handles Amazon Aurora failover event
![image](https://user-images.githubusercontent.com/1695690/63017065-3f8a1600-bed0-11e9-8476-6183de82745c.png)


## Setup
```
# Install nvm
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.2/install.sh | bash
nvm install node
nvm use node

#Install serverless package
npm install -g serverless
```

## Deploy:
### Step 1:
```
make
```
### Step 2:
```
serverless deploy --stage production
```
