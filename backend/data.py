import pandas as pd
from datetime import date
from music import Artist
from backend.spotify import *

class SongPoller():
    
    def __init__(self):
        self.df = pd.read_csv("data/data.csv", index_col=0)

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


    def pollForTracks(self):
        allTracks = []
        df = self.getData()
        artists = self.getDataAsList(df)
        for artist in artists:
            albums = self.getArtistAlbumsAfterVisitDate(artist.id,artist.lastVisited)
            for album in albums:
                tracks = self.getAlbumTracks(album)
                for track in tracks:
                    allTracks.append(track)
        return allTracks