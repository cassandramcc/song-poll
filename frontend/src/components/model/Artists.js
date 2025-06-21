import Artist from './Artist';

export default function Artists(
    {data}
) {
    return (
        data.map((artist, index) => (
            <Artist 
                key={index}
                Name={artist.Name}
                URI={artist.URI}
            />)
        )
    )
}