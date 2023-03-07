import pandas as pd
from datetime import datetime
from datetime import date

class Artist:
    def __init__(self, uri, name, lastVisited):
        self.uri = uri
        self.name = name
        self.lastVisited = lastVisited

    def __str__(self):
        return f"{self.name} - {self.uri} - {self.lastVisited}"

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

df = getData()

artists = getDataAsList(df)

for artist in artists:
    print(artist)

