# mb-playlistevents
Service to handle outbound and inbound events for playlists.

# Description

This service will handle direct communication with each of the playlist and listen for events and send events depending on conditions.
- Feeds on Playlist changes events from a queue (populated by PlaylistManager).
- Caches Playlist manager information to prevent overloading it with READ opeartions.
- Makes decisions based on events and updates playlist manager for persistent status and the UI for ongoing listeners. 

# Technologies

Go
