## Simple Job Scheduler
It is a simple job scheduler. 

### Create a job
```
# Via Post request 

curl --location --request POST 'localhost:8080/api/v1/job' \
--header 'Content-Type: application/json' \
--data-raw '{
    "jobTime": "2022-04-03T09:17:03.013Z",
    "jobType": "webhook calling",
    "jobMeta": "{\"json\":\"body\"}",
    "lockedAt": null
}'
```
