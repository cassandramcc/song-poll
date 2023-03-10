import spotipy
from spotipy.oauth2 import SpotifyClientCredentials
from music import Album, Artist, Track

import config.keys as keys
from datetime import datetime

sp = spotipy.Spotify(auth_manager=SpotifyClientCredentials(client_id=keys.clientID,
                                                           client_secret=keys.clientSecret))

def getArtistID(artist):
    searchResult = sp.search(artist,1,type='artist')
    id = searchResult['artists']['items'][0]['id']
    return id

def getArtist(id):
    artistResult = sp.artist(id)
    return Artist(artistResult['name'],artistResult['id'],None)

def getArtistAlbums(id):
    offset = 0
    albums = []
    while True:
        albumsResult = sp.artist_albums(id,limit=50,offset=offset)
        if len(albumsResult['items']) == 0:
            break
        for _, album in enumerate(albumsResult['items']):
            albums.append(Album(album['name'],[getArtist(id)],album['release_date'],album['id']))
        offset += 50
    return albums

def sortAlbums(albums):
    sortedAlbums = sorted(albums, key=lambda a: a.releaseDate, reverse=True)
    return sortedAlbums

def getAlbumTracks(album):
    tracks = []
    result = sp.album_tracks(album.id)
    for track in result['items']:
        artists = track['artists']
        artistsObj = []
        for artist in artists:
            artistsObj.append(Artist(artist['name'],artist['id'],'2023-01-01'))
        tracks.append(Track(track['name'],artistsObj,track['id']))
    return tracks

def getAlbumsAfterVisitDate(albums,visitDate):
    newAlbums = []
    for album in albums:
        try:
            if datetime.strptime(album.releaseDate,'%Y-%m-%d') >= datetime.strptime(visitDate,'%Y-%m-%d'):
                newAlbums.append(album)
        except:
            pass
    return newAlbums

def getPlaylistTracks(id):
    offset = 0
    tracks = []
    while True:
        playlistResult = sp.playlist_items(id, limit=100,offset=offset)
        if len(playlistResult['items']) == 0:
            break
        for _, track in enumerate(playlistResult['items']):
            artists = track['track']['artists']
            artistsObj = []
            for artist in artists:
                artistsObj.append(Artist(artist['name'],artist['id'],'2023-01-01'))
            tracks.append(Track(track['track']['name'],artistsObj, track['track']['id']))
        offset += 100
    return tracks

def getArtistAlbumsAfterVisitDate(id,visitDate):
    albums =  getArtistAlbums(id)
    newAlbums = getAlbumsAfterVisitDate(albums,visitDate)
    return newAlbums    

def artistExists(artist, allArtists):
    for a in allArtists:
        if a.name == artist.name:
            return True
    return False
