# terseurl

[![Go Report Card](https://goreportcard.com/badge/github.com/MicahParks/terseurl)](https://goreportcard.com/report/github.com/MicahParks/terseurl) [![Go Reference](https://pkg.go.dev/badge/github.com/MicahParks/terseurl.svg)](https://pkg.go.dev/github.com/MicahParks/terseurl)

Currently under development.

## Terms

* *client*: Something using editor access to an instance of terseurl through the REST API.
* *user*: Something visiting an instance of terseurl with a web browser without any special access.
* *Visits data*: Data about *users* that have visited a particular shortened URL.
* *Terse data*: Data that describes the process of how to get from a shortened URL to an original URL. This data
  includes the shortened URL and original URL.

## Features

### Redirection with shortened URLs

Create web browser redirections to original URLs through shortened URLs. Shortened URLs are unique URL safe strings. A
*client* may provide one, or the server can generate one.

*Example:* A *client* created *Terse data* with the shortened URL, `myblog`. The *Terse data* has the original URL of
`http://example.com/blogs/my/1`. The link `https://terseurl.com/myblog` is shared with other *users*. When a *user*
visits `https://terseurl.com/myblog`, their web browser will redirect them to `http://example.com/blogs/my/1`.

### Multiple redirection types

Currently, the project supports the following redirection types:

* HTTP 301
* HTTP 302
* HTML `<meta>`
* JavaScript

If there are more redirection types (that are widely accepted by web browsers) suggest them to the developers.

### Social media link previews

If *Terse data* is configured to perform a redirect via HTML `<meta>` tags or JavaScript, there is the option to add
social media link previews. This is done by adding HTML `<meta>` tags for [Open Graph](https://ogp.me) and
[Twitter](https://developer.twitter.com/en/docs/twitter-for-websites/cards/overview/markup).

This can be added manually to *Terse data*. It can also be inherited from the original URL by using an API endpoint, or
a button on the frontend.

### *Visits data*

By default, the project will not keep track of *Visits data*. If the project is configured to, it can track visits to
shortened URLs. All gathered *Visits data* is placed in backend storage and accessible via the web frontend or API.

The types of *Visits data* collected can vary. It can include IP address, HTTP headers, and information gathered form
JavaScript.

### Control *Terse data* and *Visits data*

*Terse data* and *Visits data* is accessible through the web interface and API. Data can easily be imported and exported
in JSON format via the frontend and API endpoints. Data can also be interacted with directly via the frontend or other
API endpoints.

### Customizable storage options

Currently, the project natively supports these storage backends:

* memory
* bbolt (file on disk)

However, the project can support any storage backend that implements its respective storage interface. TODO

## Configuration

Environment variable table:

|Name                |Description                                                                                                                                                                        |Default Value                  |Example Value                                |
|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------|---------------------------------------------|
|`DEFAULT_TIMEOUT`   |The amount of time to wait before timing out for an incoming (client) or an outgoing (database) request in seconds.                                                                |`60`                           |`180`                                        |
|`HTTP_PREFIX`       |The HTTP prefix all shortened URLs will have. This is used by the frontend.                                                                                                        |`https://terseurl.com/`        |`https://example.com/`                       |
|`INVALID_PATHS`     |A comma separated list of paths that cannot be assigned to a shortened URL. Whitespace prefixes and suffixes are trimmed. All swagger endpoints like `api` are invalid.            |swagger endpoints and frontend |`ready ,live, v2`                            |
|`SHORTID_PARANOID`  |Indicate whether randomly generated short URLs should be checked to see if they are already in use. Any value sets the boolean to true. Empty for false.                           |blank                          |`true`                                       |
|`SHORTID_SEED`      |The seed to give the random shortened URL generator. Unsigned 64 bit integer. It is recommend to set this in a production setting.                                                 |System clock                   |`2301015`                                    |
|`TEMPLATE_PATH`     |The full or relative path to the HTML template to use when a shortened URL is requested and JavaScript fingerprinting or social media link previews are on.                        |`redirect.gohtml`              |`customTemplate.gohtml`                      |
|`SUMMARY_STORE_JSON`|The JSON formatted storage configuration for the SummaryStore. If empty, it will try to read the file at `summaryStore.json`. If not found it will use an in memory implementation.|blank                          |`{"type":"memory"}`                          |
|`TERSE_STORE_JSON`  |The JSON formatted storage configuration for the TerseStore. If empty, it will try to read the file at `terseStore.json`. If not found it will use an in memory implementation.    |blank                          |`{"type":"bbolt","bboltPath":"terse.bbolt"}` |
|`VISITS_STORE_JSON` |The JSON formatted storage configuration for the VisitsStore. If empty, it will try to read the file at `visitsStore.json`. If not found, visits will not be tracked.              |blank                          |`{"type":"bbolt","bboltPath":"visits.bbolt"}`|
|`WORKER_COUNT`      |The quantity of workers to use for incoming request async tasks like performing async database calls. Not the number of incoming requests that can be handled at one time.         |`4`                            |`10`                                         |

### JSON formatted storage configuration

TODO

## Deployment

```bash
touch terse.bbolt visits.bbolt
docker-compose up
```

## TODO

- [ ] Address source code TODOs.
- [ ] Inherit HTML title.
- [ ] Copy to clipboard button for shortened URL.
- [ ] Show full shortened URL in table data.
- [ ] Hyperlinks for shortened URL and original URL.
- [ ] Embed `redirect.gohtml` add a flag that can switch between that one and one read in at runtime.
- [ ] Add referer URL to query parameters?
- [ ] Terse export reported as not found if there are no visits.
- [ ] Implement `SHORTID_PARANOID` environment variable.
- [ ] Make buttons work on `table.html`.
- [ ] Implement JavaScript tracking.
- [ ] Implement JavaScript fingerprinting.
- [ ] Create HTML navigation.
- [ ] Customizable visit data tracking.
- [ ] Visit data middleware for data purging or whatever before it goes to backend storage.
- [ ] SummaryStore not used when VisitsStore is `nil`?
- [ ] Outgoing validation of spec with model methods?
- [ ] Make sure the frontend has a way to interact with visits data when the terse data has been deleted.
- [ ] Make endpoints for reading multiple Terse/Visits.
- [ ] Change endpoints like `/api/delete` to `/api/delete/all`.
- [ ] Write a utility that will export `.bbolt` to JSON.
- [ ] Implement `SHORTID_PARANOID`.
- [ ] Allow for shortened URLs of the form `{owner}/{shortened}` in `/api/write/{operation}` endpoint.
  - [ ] Only allow for random shortened URLs in top level.
- [ ] Implement fingerprinting with fingerprintjs, but remove HTML canvas extraction. Embed minified in single HTML
  template.
- [ ] Implement Redis storage backend?
- [ ] Implement pluggable store interface implementations. (Using custom `main`.)
- [ ] Reimplement Mongo storage.
- [ ] Write tests.
- [ ] Write a good README.md.
- [ ] Implement JWT + JWKS authentication?
