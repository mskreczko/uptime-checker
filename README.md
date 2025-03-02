# Uptime Checker

CLI application that helps you monitor health of your services.

## Config example
```yaml
applications:
  - name: test
    targetGroups:
      - name: group1
        targets:
          - http://localhost:8000/health
          - http://localhost:8001/health
        healthCheckInterval: 10
        healthCheckStrategy: raw

notifications:
  notification_settings:
    - channel: EMAIL
      receivers:
        - test@test.com
        - test2@test.com

smtp:
  region: us-east-1
  access_key: 
  secret_key: 
  token:
  sender: uptimechecker@yourdomain.com
```