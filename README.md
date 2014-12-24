webhog
======

**webhog** is a package that stores and downloads a given URL (including js, css, and images) for offline use and uploads it to a given AWS-S3 account (more persistance options to come).

##Installation
`go get github.com/johnernaut/webhog`

##Usage
Make a `POST` request to `http://localhost:3000/scrape` with a header set to value `X-API-KEY: SCRAPEAPI`.  Pass in a JSON value of the URL you'd like to fetch: `{ "url": "http://facebook.com"}` (as an example).  You'll notice an `Ent dir: /blah/blah/blah` printed to the console - your assets are saved there.  To test, open the given `index.html` file.

##Configuration
Create a `webhog.yml` file in the running directory.  The following options are supported:
```
development:
  mongodb: mongodb://127.0.0.1:27017/webhog
  api_key: SCRAPEAPI
  aws_key: <env var name for s3 key>
  aws_secret: <env var name for s3 secret>
  bucket: mybucket
production:
staging:
```
The setting root-key is established via a `MARTINI_ENV` environment variable that you should set.

### Usage Docs

Start with a POST request to Webhog, like so:

`POST http://webcache.mybigcampus.com/api/scrape`

Be sure to include the `X-API-KEY` = `SCRAPEAPI` as a header and the form data of `url` = `http://www.twitter.com`. The JSON response will look like this:

```json
{
    "id": "549adfd30a97ec9119d60bff",
    "uuid": "4b848173-78d0-4f9f-4e34-21e7652edec9",
    "url": "http://www.twitter.com",
    "aws_link": "",
    "status": "parsing",
    "created_at": "2014-12-24T15:43:39.392Z"
}
```

Note the status key. If the url has been parsed in the past , it will be cached with the status of `complete` and the `aws_link` will have the tar.gz file to download. Otherwise, you can use the `entity` endpoint to get a status of the parsing progress like so:

`GET http://webcache.mybigcampus.com/api/entity/4b848173-78d0-4f9f-4e34-21e7652edec9`

Note that the param in the url resource is the `uuid` that was returned from the initial `POST` request above. Feel free to query the route above as many times you like until the parsing is complete. Once finished, you should see a JSON response like so:

```json
{
    "id": "549adfd30a97ec9119d60bff",
    "uuid": "4b848173-78d0-4f9f-4e34-21e7652edec9",
    "url": "http://www.twitter.com",
    "aws_link": "https://s3.amazonaws.com/mbc-webcache-production/iDwUM6CnfI10401333421234110778778.tar.gz",
    "status": "complete",
    "created_at": "2014-12-24T15:43:39.392Z"
}
```

Use the `aws_link` to pull a tar'ed and gzipped file of the webpage, it's assets and such to use locally in your stashed content.
