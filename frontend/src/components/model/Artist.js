import './Artist.css';

export default function Artist(
    {
        Name = "Placeholder name",
        URI = "Placeholder uri"
    }
) {
    return (
        <div className="artist-container">
            <p className="artist-field bold">{Name}</p>
            <p className="artist-field">{URI}</p>
        </div>
    )
}