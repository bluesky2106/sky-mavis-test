# IP Location Record

## How to deploy on a VM (Virtual machine)

1. Access to directory `part-3`

2. Run following command:

```
docker-compose up
```

3. Access to `http://<server>:4200` where server is the ip of your VM.

In my case, I also deployed the app on a VM and you can access the following url for testing:

- http://103.160.5.142:4200/

However, the frontend app is running in debug mode. I haven't built frontend source code and use nginx / apache server to serve frontend files.
## Gloud App Engine & Firebase deployment

Currently, backend is deployed by using google cloud app engine. Please look at the `backend/app.yaml` for more details.

In addition, the frontend is deployed onto firebase.

Finally, the mysql DB is deployed in a VM. The reason why I don't use Google Cloud SQL for this DB is that the cost of Cloud SQL is not cheap at all.

So, in order to test the app, please access to the following links (both of them link to the same app):
- https://sky-mavis-94446.firebaseapp.com/
- https://sky-mavis-94446.web.app/
