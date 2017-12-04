{
  "type": "object",
  "$schema": "http://json-schema.org/draft-03/schema",
  "id": "http://jsonschema.net",
  "required": true,
  "properties": {
    "eventType": {
      "type": "string",
      "required": true
    },
    "playlistId": {
      "type": "uuid",
      "required": true
    },
    "songId": {
      "type": "uuid",
      "required": true
    },
    "userId": {
      "type": "uuid",
      "required": true
    }
  }
}
