carbon-relay-ng
===============

A relay for carbon streams, in go.
Like carbon-relay from the graphite project, except it:


 * should perform better (benchmark currently reports 50k metrics/s through a table with 1 endpoint. in prod i see 20k metrics/s before maxing out a single core.  still much room for optimization)
 * you can adjust the routing table at runtime, in real time using the web or telnet interface (interfaces need more work)
 * has aggregator functionality built-in
 * supports a per-route spooling policy.
   (i.e. in case of an endpoint outage, we can temporarily queue the data up to disk and resume later)
 * you can choose between plaintext or pickle output, per route.
 * can be restarted without dropping packets (needs testing)
 

This makes it easy to fanout to other tools that feed in on the metrics.
Or balance/split load, or provide redundancy, or partition the data, etc.
This pattern allows alerting and event processing systems to act on the data as it is received (which is much better than repeated reading from your storage)


![screenshot](https://raw.githubusercontent.com/graphite-ng/carbon-relay-ng/master/screenshot.png)


Future work aka what's missing
------------------------------

* multi-node clustered high availability (open for discussion whether it's worth it)
* pub-sub interface, maybe
* consistent hashing across different endpoints, if it can be implemented in an elegant way.  (note that this would still be a hack and mostly aimed for legacy setups, [decent storage has redundancy and distribution built in properly ](http://dieter.plaetinck.be/on-graphite-whisper-and-influxdb.html).


Releases & versions
-------------------

* master (work in progress): refactored version with non-blocking operations, extensive internal stats and a more extensive routing system (which can support round robin & hashing) (see #23).  the admin interfaces need more work.
* v0.5 extended version with config file, http and telnet interfaces, statsd for internal instrumentation, disk spooling support, but still blocking sends. (no longer supported)
* v0.1 initial, simple version that used commandline args to configure. no admin interfaces. blocking sends (1 endpoint down blocks the program)(no longer supported)


Instrumentation
---------------

* All performance variables are available in json at http://localhost:8081/debug/vars2 (update port if you change it in config)
* You can also send metrics to graphite (or feed back into the relay), see config.
* Comes with a [grafana dashboard template](https://github.com/graphite-ng/carbon-relay-ng/blob/master/grafana-dashboard.json) so you get up and running in no time.

![grafana dashboard](https://raw.githubusercontent.com/graphite-ng/carbon-relay-ng/master/grafana-screenshot.png)


Building
--------

    export GOPATH=/some/path/
    export PATH="$PATH:$GOPATH/bin"
    go get -d github.com/graphite-ng/carbon-relay-ng
    go get github.com/jteeuwen/go-bindata/...
    cd "$GOPATH/src/github.com/graphite-ng/carbon-relay-ng"
    # optional: check out an older version: git checkout v0.5
    make


Installation
------------

    You only need the compiled binary and a config.  Put them whereever you want.

Usage
-----

<pre><code>carbon-relay-ng [-cpuprofile <em>cpuprofile-file</em>] <em>config-file</em></code></pre>


Concepts
--------

You have 1 master routing table.  This table contains 0-N routes.  Each route can contain 0-M destinations (tcp endpoints)

First: "matching": you can match metrics on one or more of: prefix, substring, or regex.  All 3 default to "" (empty string, i.e. allow all).
The conditions are AND-ed.  Regexes are more resource intensive and hence should, and often can be avoided.

* All incoming matrics are validated, filtered through the blacklist and then go into the table.
* The table sends the metric to:
  * the aggregators, who match the metrics against their rules, compute aggregations and feed results back into the table. see Aggregation section below.
  * any routes that matches
* The route can have different behaviors, based on its type:

  * sendAllMatch: send all metrics to all the defined endpoints (possibly, and commonly only 1 endpoint).
  * sendFirstMatch: send the metrics to the first endpoint that matches it.
  * consistent hashing: the route is a CH pool (not implemented)
  * round robin: the route is a RR pool (not implemented)


carbon-relay-ng (for now) focuses on staying up and not consuming much resources.

if connection is up but slow, we drop the data
if connection is down and spooling enabled.  we try to spool but if it's slow we drop the data
if connection is down and spooling disabled -> drop the data



Aggregation
-----------

As discussed in concepts above, we can combine, at each point in time, the points of multiple series into a new series.
Note:
* The interval parameter let's you quantize ("fix") timestamps, for example with an interval of 60 seconds, if you have incoming metrics for times that differ from each other, but all fall within the same minute, they will be counted together.
* The wait parameter allows up to the specified amount of seconds to wait for values, With a wait of 120, metrics can come 2 minutes late and still be included in the aggregation results.
* The fmt parameter dictates what the metric key of the aggregated metric will be.  use $1, $2 to refer to groups in the regex
* Note that we direct incoming values to an aggregation bucket based on the interval the timestamp is in, and the output key it generates.
  This means that you can have 3 aggregation cases, based on how you set your regex, interval and fmt string.
  - aggregation of points with different metric keys, but with the same, or similar timestamps) into one outgoing value (~ carbon-aggregator).
    if you set the interval to the period between each incoming packet of a given key, and the fmt yields the same key for different input metric keys
  - aggregation of individual metrics, i.e. packets for the same key, with different timestamps.  For example if you receive values for the same key every second, you can aggregate into minutely buckets by setting interval to 60, and have the fmt yield a unique key for every input metric key.  (~ graphite rollups)
  - the combination: compute aggregates from values seen with different keys, and at multiple points in time.
* functions currently available: avg and sum
* aggregation output goes back into the table so you can route it however you want.
* since the output metrics go back into the table, this means you can create new rules that leverage results from other rules.
  (this is probably not useful and will probably be changed). Either way, make sure you don't accidentally match your own output.
* see the included ini for examples


Configuration
-------------


Look at the included carbon-relay-ng.ini, it should be self describing.
In the init option you can create routes, populate the blacklist, etc using the same command as the telnet interface, detailed below.
This mechanism is choosen so we can reuse the code, instead of doing much configuration boilerplate code which would have to execute on
a declarative specification.  We can just use the same imperative commands since we just set up the initial state here.


TCP interface
-------------

commands:

    help                                         show this menu
    view                                         view full current routing table

    addBlack <substring>                         blacklist (drops the metric matching this as soon as it is received)

    addAgg <func> <regex> <fmt> <interval> <wait>  add a new aggregation rule.
             <func>:                             aggregation function to use
               sum
               avg
             <regex>                             regex to match incoming metrics. supports groups (numbered, see fmt)
             <fmt>                               format of output metric. you can use $1, $2, etc to refer to numbered groups
             <interval>                          align odd timestamps of metrics into buckets by this interval in seconds.
             <wait>                              amount of seconds to wait for "late" metric messages before computing and flushing final result.


    addRoute <type> <key> [opts]   <dest>  [<dest>[...]] add a new route. note 2 spaces to separate destinations
             <type>:
               sendAllMatch                      send metrics in the route to all destinations
               sendFirstMatch                    send metrics in the route to the first one that matches it
             <opts>:
               prefix=<str>                      only take in metrics that have this prefix
               sub=<str>                         only take in metrics that match this substring
               regex=<regex>                     only take in metrics that match this regex (expensive!)
             <dest>: <addr> <opts>
               <addr>                            a tcp endpoint. i.e. ip:port or hostname:port
               <opts>:
                   prefix=<str>                  only take in metrics that have this prefix
                   sub=<str>                     only take in metrics that match this substring
                   regex=<regex>                 only take in metrics that match this regex (expensive!)
                   flush=<int>                   flush interval in ms
                   reconn=<int>                  reconnection interval in ms
                   pickle={true,false}           pickle output format instead of the default text protocol
                   spool={true,false}            enable spooling for this endpoint

    addDest <routeKey> <dest>                    not implemented yet

    modDest <routeKey> <dest> <opts>:            modify dest by updating one or more space separated option strings
                   addr=<addr>                   new tcp address
                   prefix=<str>                  new matcher prefix
                   sub=<str>                     new matcher substring
                   regex=<regex>                 new matcher regex

    modRoute <routeKey> <opts>:                  modify route by updating one or more space separated option strings
                   prefix=<str>                  new matcher prefix
                   sub=<str>                     new matcher substring
                   regex=<regex>                 new matcher regex

