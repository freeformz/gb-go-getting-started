
# gb-go-getting-started

A barebones Go app, which can easily be deployed to Heroku.  

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```console
# Install gb to your $GOPATH
$ go get github.com/constabulary/gb/...
$ cd ~/devel               #or wherever you keep your projects
$ git clone https://github.com/freeformz/gb-go-getting-started.git
$ cd gb-go-getting-started
$ gb build all
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```console
$ heroku create
$ git push heroku master
$ heroku open
```
or 

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
- [Go Dependencies via gb](https://devcenter.heroku.com/articles/go-dependencies-via-gb?preview=1)
