jsonjson
========

Yo dawg, I heard you liked json. Here is a jsonpp for your json containing json.

The bolt command line utility in https://github.com/boltdb/bolt has an export command that exports the contents of the bolt database as json.

If you, like I do, store json objects in your boltdb, this utility can be handy.

Just pipe your json through jsonjson:

```
bolt export bolt.db | jsonjson | less
```
