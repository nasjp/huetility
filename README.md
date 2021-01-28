# huetility

huetility is [Philips Hue](https://www.philips-hue.com/en-us)'s utility command line tool.

```sh
$ cat $HOME/.huetility.json # please make .huetility.json
{
  "ip_address": "192.168.11.22",
  "user_name": "Cj4emVG69juW2A63ELkE861hA2x46mXwfvsqSuTW",
  "scenes" : [
    {
      "id": "1",
      "light_id": "3",
      "name": "Concentration",
      "state": {
        "bri": 254,
        "ct": 230,
        "alert": "none",
        "on": true
      }
    },
    {
      "id": "2",
      "light_id": "3",
      "name": "Savannah",
      "state": {
        "bri": 123,
        "ct": 322,
        "alert": "none",
        "on": true
      }
    },
    {
      "id": "3",
      "light_id": "3",
      "name": "Off",
      "state": {
        "alert": "none",
        "on": false
      }
    }
  ]
}

$ huetility get # get current state
id: 3
  name: Hue ambiance lamp 1
  state:
    bri       : 123
    ct        : 322
    alert     : "none"
    colormode : "ct"
    mode      : "homeautomation"
    reachable : true
    on        : false

$ huetility set 1 # set state with 'Concentration' scene
id: 3
  name: Hue ambiance lamp 1
  state:
    bri       : 254
    ct        : 230
    alert     : "none"
    colormode : "ct"
    mode      : "homeautomation"
    reachable : true
    on        : true
```
