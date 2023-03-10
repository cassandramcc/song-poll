import pandas as pd
from datetime import date
from music import Artist
from spotify import *

def getData() -> pd.DataFrame:
    df = pd.read_csv("data/data.csv", index_col=0)
    return df


def updateLastVisited(df: pd.DataFrame, when=date.today()):
    """
    Puts a specified date in the last visited column of the data, or today's date if left unspecified
    """
    df['last visited'] = df['last visited'].map(lambda x: when)


def setData(df: pd.DataFrame):
    df.to_csv("data/data.csv")


def getDataAsList(df :pd.DataFrame):
    result = [Artist(uri,name,lastVisited) for uri,name,lastVisited in zip(df['uri'], df['name'], df['last visited'])]
    return result


def songpoll():
    allTracks = []
    df = getData()
    artists = getDataAsList(df)
    for artist in artists:
        albums = getArtistAlbumsAfterVisitDate(artist.id,artist.lastVisited)
        for album in albums:
            tracks = getAlbumTracks(album)
            for track in tracks:
                spAdd.playlist_add_items("6fIMpwbZo8wrKCDNmGjiS3",[track.id])
                allTracks.append(track)
    return allTracks

songpoll()