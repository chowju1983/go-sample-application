# go-sample-application
Sample GO Application for Performance Benchmark for GO Application.
This application is deployed in AWS Lambda using the native binary that is generated as part of the Docker build process. The subsequet Docker image is then pushed to AWS ECR(Elastic Container Registry) which is then loaded into the Lambda function.

The Lambda Function is then fronted by an API Gateway REST Endpoint. The request_id needs to be passed as the Path Parameter in the request. Current implementation supports only the path parameter 1. 

Data is pre-loaded in a PostGres Database instance running in EC2. 

Please build the Docker Image using docker build .... and then push the same to an ECR repo. 

