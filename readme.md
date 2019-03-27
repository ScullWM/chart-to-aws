# Chart-to-aws (screener)

Microservice to generate screenshot from a webpage and upload it to a AWS S3 Bucket.

### Usage
Simple call with default values using curl:
`curl http://localhost:8000/\?path\=/mypage/id\&selector\=my_chartjs_id\&output\=screen/my_chartjs_id.png`

Generate a screenshot from a javascript chart
![](http://www.updemia.com/static/g/b/xl/5c5acbd09a056.png)
Can be include in email or Slack notification

### Configuration
Service can be configured by the config.yaml file. Here's the default values:


```
httpserver:
    port: '127.0.0.1:8000'
    query: 'path'
    selector: 'selector'
    output: 'ouput'
domain: 'https://pepperreport.io/'
aws:
  id: ''
  secret: ''
  token: ''
  bucket: ''
  region: 'eu-central-1'

```

Query terms can be updated using the config file.
`domain` used to restrict scope of screenshot source
`path` is corresponding to the domain path
`selector` CSS id selector of the element
`output` Destination of the file in the bucket


### Contribute
To contribute just open a Pull Request with your new code. Feel free to contribute.
