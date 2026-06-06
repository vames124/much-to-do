# StartTech Operations Runbook

## 1. CI/CD Pipeline Failures

### Scenario: Frontend Deployment Fails
* **Symptom:** GitHub Actions `deploy` step fails with "Access Denied" or "NoSuchBucket".
* **Resolution:** 1. Verify the `S3_BUCKET_NAME` and `CLOUDFRONT_DIST_ID` GitHub repository secrets match the current Terraform outputs.
  2. Verify the OIDC IAM Role has permissions for `s3:PutObject` and `cloudfront:CreateInvalidation`.

### Scenario: Backend Deployment Fails at ECR Push
* **Symptom:** Pipeline fails during "Login to Amazon ECR" or "Push Docker Image".
* **Resolution:**
  1. Ensure the ECR repository (`starttech-backend`) exists in the target AWS region.
  2. Verify the AWS Region environment variable in `backend-ci-cd.yml` matches the deployed infrastructure.

## 2. Infrastructure Troubleshooting

### Scenario: ALB Health Checks are Failing (502 Bad Gateway)
* **Symptom:** The ALB marks EC2 instances as "Unhealthy" and the application is unreachable.
* **Resolution:**
  1. **Check Application Logs:** Navigate to CloudWatch Logs > `/var/log/*-backend-app` to verify if the Golang application crashed on startup.
  2. **Check Security Groups:** Verify the Backend Security Group allows inbound traffic from the ALB Security Group on port 8080.
  3. **Check Route to Internet:** If the Go app is failing to start because it cannot reach MongoDB, ensure the NAT Gateway is active and properly associated with the Private Route Table.

## 3. Manual Rollback Procedures

### Rolling Back the Frontend
If a critical bug is deployed to the React frontend:
1. In the GitHub repository, navigate to the **Actions** tab.
2. Select the `Frontend CI/CD` workflow.
3. Find the previous, known-good workflow run.
4. Click **Re-run all jobs** to rebuild and redeploy the stable commit over the broken one.

### Rolling Back the Backend
If a new Docker image causes API failures:
1. Navigate to the AWS EC2 Console > Auto Scaling Groups.
2. Edit the Launch Template associated with the ASG.
3. Update the User Data script (or the deployment parameter) to pull the previous, stable Docker image tag from ECR instead of `latest`.
4. Trigger an Instance Refresh to cycle out the broken containers.