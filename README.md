# Hex Plugin - Response

Simple plugin that responds with a string.

```
{
  "rule": "response rule example",
  "match": "hello",
  "actions": [
    {
      "type": "hex-response",
      "command": "Hello ${hex.user}!"
    }
  ]
}
```
