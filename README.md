# Health Record Repository - Vulnerable Go Application

This is a deliberately vulnerable Go application designed for SBOM scanning and security testing.

## Features

- HTTP server running on port 5655
- Multiple vulnerable dependencies for SBOM testing
- REST endpoints: `/sbom`, `/login`, `/upload`
- Hardcoded secrets and insecure practices

## Local Development

```bash
# Run locally
go run main.go

# Build Docker image
docker build -t healthrecord-repository .

# Run Docker container
docker run -p 5655:5655 healthrecord-repository
```

## AWS Deployment (ECR + ECS)

### Prerequisites

1. AWS CLI installed and configured
2. Docker installed
3. Appropriate IAM permissions for ECR and ECS

### Deployment Steps

1. **Update Configuration**:
   - Edit `deploy.sh` and replace:
     - `YOUR_ACCOUNT_ID` with your AWS account ID
     - `us-east-1` with your preferred region
   - Edit `ecs-task-definition.json` and update:
     - Subnet IDs in the network configuration
     - Security group IDs

2. **Deploy**:
   ```bash
   chmod +x deploy.sh
   ./deploy.sh
   ```

### Manual Steps

If you prefer manual deployment:

```bash
# 1. Create ECR repository
aws ecr create-repository --repository-name healthrecord-repository

# 2. Login to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin YOUR_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com

# 3. Build and push
docker build -t healthrecord-repository .
docker tag healthrecord-repository:latest YOUR_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/healthrecord-repository:latest
docker push YOUR_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/healthrecord-repository:latest

# 4. Create ECS cluster
aws ecs create-cluster --cluster-name healthrecord-cluster

# 5. Register task definition (after updating the JSON file)
aws ecs register-task-definition --cli-input-json file://ecs-task-definition.json

# 6. Create ECS service (update network configuration)
aws ecs create-service --cluster healthrecord-cluster --service-name healthrecord-service --task-definition healthrecord-repository --desired-count 1 --launch-type FARGATE --network-configuration "awsvpcConfiguration={subnets=[subnet-xxx],securityGroups=[sg-xxx],assignPublicIp=ENABLED}"
```

## Endpoints

- `GET /sbom` - Returns vulnerability information
- `GET /login?username=user` - Insecure login endpoint
- `POST /upload` - File upload endpoint

## Security Vulnerabilities (Intentional)

This application contains intentional vulnerabilities for testing:
- Vulnerable JWT library (github.com/dgrijalva/jwt-go)
- Hardcoded secrets
- Insecure file uploads
- Multiple packages with known CVEs

Perfect for SBOM scanning tools like Grype, Snyk, or Trivy.